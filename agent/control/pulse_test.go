/*
Copyright 2018 Google Inc. All Rights Reserved.
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

package control

import (
	"context"
	"sync"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/common"
	pubsubinternal "github.com/GoogleCloudPlatform/cloud-ingest/agent/pubsub"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/stats"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"

	pulsepb "github.com/GoogleCloudPlatform/cloud-ingest/proto/pulse_go_proto"
)

func TestPulseSender(t *testing.T) {
	tests := []int{1, 3, 10, 100}
	for _, numPulses := range tests {
		ctx := context.Background()

		// Set up the PubSub mock.
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockPublishResult := pubsubinternal.NewMockPSPublishResult(ctrl)
		mockPublishResult.EXPECT().Get(ctx).MaxTimes(numPulses).MinTimes(numPulses).Return("serverid", nil)
		mockPulseTopic := pubsubinternal.NewMockPSTopic(ctrl)

		st := stats.NewTracker(ctx)
		st.RecordBytesSent(123)
		st.RecordBytesSent(456)
		st.RecordBytesSent(789)

		logsDir := "/tmp/mylogs"
		ps := NewPulseSender(ctx, mockPulseTopic, logsDir, st)
		ps.hostname = "hostname"
		ps.pid = 1234
		ps.version = "1.2.3"

		// Complete the mock pulse topic.
		msg, err := proto.Marshal(ps.pulseMsg())
		if err != nil {
			t.Fatalf("proto.Marshal(%v) got err: %v", ps.pulseMsg(), err)
		}
		mockPulseTopic.EXPECT().Publish(ctx, &pubsub.Message{Data: msg}).MaxTimes(numPulses).MinTimes(numPulses).Return(mockPublishResult)

		// Set up the test hooks and send the pulses.
		var wg sync.WaitGroup
		ps.selectDone = func() { wg.Done() }
		mockSendTicker := common.NewMockTicker()
		ps.sendTicker = mockSendTicker
		for i := 0; i < numPulses; i++ {
			wg.Add(1)
			mockSendTicker.Tick()
			wg.Wait()
		}
	}
}

func TestPulseMsg(t *testing.T) {
	tests := []struct {
		hostname string
		pid      int
		logsDir  string
		version  string
		want     *pulsepb.Msg
	}{
		{
			"hostname", 1234, "/tmp/mylogs", "1.2.3",
			&pulsepb.Msg{
				AgentId: &pulsepb.AgentId{
					HostName:  "hostname",
					ProcessId: "1234",
				},
				AgentVersion: "1.2.3",
				AgentLogsDir: "/tmp/mylogs",
			},
		},
	}
	for _, tc := range tests {
		ps := &PulseSender{
			hostname: tc.hostname,
			pid:      tc.pid,
			logsDir:  tc.logsDir,
			version:  tc.version,
		}
		if got := ps.pulseMsg(); !proto.Equal(got, tc.want) {
			t.Errorf("ps.pulseMsg() = %v, want %v", got, tc.want)
		}
	}
}
