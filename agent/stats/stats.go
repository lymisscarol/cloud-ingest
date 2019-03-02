/*
Copyright 2019 Google Inc. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stats

import (
	"context"
	"fmt"
	"io"
	"math"
	"sort"
	"time"

	"github.com/GoogleCloudPlatform/cloud-ingest/agent/stats/throughput"
	"github.com/GoogleCloudPlatform/cloud-ingest/helpers"
	"github.com/golang/glog"

	taskpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/task_go_proto"
)

const (
	statsLogFreq     = 3 * time.Minute // The frequency of logging periodic stats.
	statsDisplayFreq = 1 * time.Second // The frequency of displaying stats to stdout.
)

type taskDur struct {
	task string
	dur  time.Duration
}

type periodicStats struct {
	taskDurations    map[string][]time.Duration
	taskFailures     map[string]int
	bytesCopied      int64
	ctrlMsgsReceived int64
	pulseMsgsSent    int64
}

type lifetimeStats struct {
	taskDone    map[string]uint64
	bytesCopied int64

	ctrlMsgTime time.Time
	bwLimit     int64
}

// Tracker collects stats about the Agent, provides a display to STDOUT, and
// periodically writes to the INFO log. Stats are collected by calling the
// various Record* functions as appropriate.
type Tracker struct {
	taskDurChan   chan taskDur   // Channel to record task durations.
	taskFailChan  chan string    // Channel to record task failures.
	bytesSentChan chan int64     // Channel to record bytesSent.
	bwLimitChan   chan int64     // Channel to record the bandwidth limit.
	ctrlMsgChan   chan time.Time // Channel to record control message timing.
	pulseMsgChan  chan struct{}  // Channel to record send pulse messages.

	periodic  periodicStats       // Reset after every time they're INFO logged.
	lifetime  lifetimeStats       // Cumulative for the lifetime of this procces.
	tpTracker *throughput.Tracker // Measures outgoing copy throughput.

	spinnerIdx int // For displaying the mighty spinner.

	// Testing hooks.
	selectDone    func()
	logTicker     helpers.Ticker
	displayTicker helpers.Ticker
}

// NewTracker returns a new Tracker, which can then be used to record stats.
func NewTracker(ctx context.Context) *Tracker {
	t := &Tracker{
		// Large buffers to avoid blocking.
		taskDurChan:   make(chan taskDur, 100),
		taskFailChan:  make(chan string, 10),
		bytesSentChan: make(chan int64, 100),
		bwLimitChan:   make(chan int64, 10),
		ctrlMsgChan:   make(chan time.Time, 10),
		pulseMsgChan:  make(chan struct{}, 10),
		periodic: periodicStats{
			taskDurations: make(map[string][]time.Duration),
			taskFailures:  make(map[string]int),
		},
		lifetime: lifetimeStats{
			taskDone:    map[string]uint64{"copy": 0, "list": 0},
			ctrlMsgTime: time.Now(),
			bwLimit:     math.MaxInt32,
		},
		tpTracker:     throughput.NewTracker(ctx),
		selectDone:    func() {},
		logTicker:     helpers.NewClockTicker(statsLogFreq),
		displayTicker: helpers.NewClockTicker(statsDisplayFreq),
	}
	go t.track(ctx)
	return t
}

// RecordTaskResp tracks the count and duration of completed tasks.
func (t *Tracker) RecordTaskResp(resp *taskpb.TaskRespMsg, dur time.Duration) {
	task := ""
	if resp.ReqSpec.GetCopySpec() != nil {
		task = "copy"
	} else if resp.ReqSpec.GetListSpec() != nil {
		task = "list"
	} else if resp.ReqSpec.GetCopyBundleSpec() != nil {
		task = "copy"
	} else {
		glog.Errorf("resp.ReqSpec doesn't match any known spec type: %v", resp.ReqSpec)
	}

	if task != "" {
		t.taskDurChan <- taskDur{task, dur} // Record the task duration.

		if resp.FailureType != taskpb.FailureType_UNSET_FAILURE_TYPE {
			t.taskFailChan <- task // Record the failure.
		}
	}
}

// ByteTrackingReader is an io.Reader that wraps another io.Reader and
// performs byte tracking during the Read function.
type ByteTrackingReader struct {
	reader  io.Reader
	tracker *Tracker
}

// NewByteTrackingReader returns a ByteTrackingReader.
func (t *Tracker) NewByteTrackingReader(r io.Reader) io.Reader {
	return ByteTrackingReader{reader: r, tracker: t}
}

// Read implements the io.Reader interface.
func (btr ByteTrackingReader) Read(buf []byte) (n int, err error) {
	if n, err = btr.reader.Read(buf); err != nil {
		return 0, err
	}
	btr.tracker.RecordBytesSent(int64(n))
	return n, nil
}

// RecordBytesSent tracks the count of bytes sent, and enables throughput tracking.
// For accurate throughput measurement this function should be called every time
// bytes are sent on the wire. More frequent and granular calls to this function
// will provide a more accurate throughput measurement.
func (t *Tracker) RecordBytesSent(bytes int64) {
	t.tpTracker.RecordBytesSent(bytes)
	t.bytesSentChan <- bytes
}

// RecordBWLimit tracks the current bandwidth limit.
func (t *Tracker) RecordBWLimit(agentBW int64) {
	t.bwLimitChan <- agentBW
}

// RecordCtrlMsg tracks received control messages.
func (t *Tracker) RecordCtrlMsg(time time.Time) {
	t.ctrlMsgChan <- time
}

// RecordPulseMsg tracks sent pulse messages.
func (t *Tracker) RecordPulseMsg() {
	t.pulseMsgChan <- struct{}{}
}

// Throughput returns the current measured throughput.
func (t *Tracker) Throughput() int64 {
	return t.tpTracker.Throughput()
}

func (t *Tracker) track(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				glog.Infof("stats.Tracker track ctx ended with err: %v", err)
			}
			return
		case tr := <-t.taskDurChan:
			t.periodic.taskDurations[tr.task] = append(t.periodic.taskDurations[tr.task], tr.dur)
			t.lifetime.taskDone[tr.task]++
		case task := <-t.taskFailChan:
			t.periodic.taskFailures[task]++
		case bytes := <-t.bytesSentChan:
			t.periodic.bytesCopied += bytes
			t.lifetime.bytesCopied += bytes
		case agentBW := <-t.bwLimitChan:
			t.lifetime.bwLimit = agentBW
		case time := <-t.ctrlMsgChan:
			t.periodic.ctrlMsgsReceived++
			t.lifetime.ctrlMsgTime = time
		case <-t.pulseMsgChan:
			t.periodic.pulseMsgsSent++
		case <-t.logTicker.GetChannel():
			t.periodic.glogAndReset()
		case <-t.displayTicker.GetChannel():
			t.displayStats()
		}
		t.selectDone() // Testing hook.
	}
}

func (t *Tracker) displayStats() string {
	// Generate the transmission rate and sum.
	txRate := fmt.Sprintf("txRate:%v/s", byteCountBinary(t.tpTracker.Throughput(), 7))
	if txLim := t.lifetime.bwLimit; txLim > 0 && txLim < math.MaxInt32 {
		txRate += fmt.Sprintf(" (capped at %v/s)", byteCountBinary(t.lifetime.bwLimit, 7))
	}
	txSum := fmt.Sprintf("txSum:%v", byteCountBinary(t.lifetime.bytesCopied, 7))

	// Generate the task response counts.
	taskResps := "taskResps["
	var keys []string
	for k := range t.lifetime.taskDone {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i > 0 {
			taskResps += " "
		}
		taskResps += fmt.Sprintf("%v:%v", k, t.lifetime.taskDone[k])
	}
	taskResps += "]"

	// Generate the control message age and status.
	ctrlMsgAge := "-"
	ctrlMsgHealth := "-"
	if !t.lifetime.ctrlMsgTime.IsZero() {
		age := time.Now().Sub(t.lifetime.ctrlMsgTime).Truncate(time.Second)
		ctrlMsgAge = fmt.Sprintf("%v", age)
		ctrlMsgHealth = "ok"
		if age > 30*time.Second {
			ctrlMsgHealth = "??"
		}
	}
	ctrlMsg := fmt.Sprintf("ctrlMsgAge:%v (%v)", ctrlMsgAge, ctrlMsgHealth)

	// Generate the spinner.
	spinnerChars := `-\|/`
	t.spinnerIdx = (t.spinnerIdx + 1) % len(spinnerChars)
	spinner := spinnerChars[t.spinnerIdx]

	// Display the generated stats.
	// TODO(b/123023481): Ensure the Agent display works on Windows.
	fmt.Printf("\r%120s\r", "") // Overwrite the previous line and reset the cursor.
	displayLine := fmt.Sprintf("%v %v %v %v %c", txRate, txSum, taskResps, ctrlMsg, spinner)
	fmt.Print(displayLine)
	return displayLine // For testing.
}

func byteCountBinary(b int64, pad int) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%*dB", pad, b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%*.1f%ciB", pad-2, float64(b)/float64(div), "KMGTPE"[exp])
}
