/*
Copyright 2017 Google Inc. All Rights Reserved.
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

package dcp

import (
	"encoding/json"
	"testing"
)

/*******************************************************************************
GetJobStatus Tests
*******************************************************************************/

func TestGetJobStatusNotStarted(t *testing.T) {
	progressObj := JobCounters{
		counter: map[string]int64{
			KeyTotalTasks:     0,
			KeyTasksCompleted: 0,
			KeyTasksFailed:    0,
		},
	}
	status := progressObj.GetJobStatus()
	if status != JobNotStarted {
		t.Errorf("expected job status for %+v to be %d, instead found %d",
			progressObj, JobNotStarted, status)
	}
}

func TestGetJobStatusInProgressNoFailures(t *testing.T) {
	progressObj := JobCounters{
		counter: map[string]int64{
			KeyTotalTasks:     5,
			KeyTasksCompleted: 3,
			KeyTasksFailed:    0,
		},
	}
	status := progressObj.GetJobStatus()
	if status != JobInProgress {
		t.Errorf("expected job status for %+v to be %d, instead found %d",
			progressObj, JobInProgress, status)
	}
}

func TestGetJobStatusInProgressWithFailures(t *testing.T) {
	progressObj := JobCounters{
		counter: map[string]int64{
			KeyTotalTasks:     5,
			KeyTasksCompleted: 3,
			KeyTasksFailed:    1,
		},
	}
	status := progressObj.GetJobStatus()
	if status != JobInProgress {
		t.Errorf("expected job status for %+v to be %d, instead found %d",
			progressObj, JobInProgress, status)
	}
}

func TestGetJobStatusSuccess(t *testing.T) {
	progressObj := JobCounters{
		counter: map[string]int64{
			KeyTotalTasks:     5,
			KeyTasksCompleted: 5,
			KeyTasksFailed:    0,
		},
	}
	status := progressObj.GetJobStatus()
	if status != JobSuccess {
		t.Errorf("expected job status for %+v to be %d, instead found %d",
			progressObj, JobSuccess, status)
	}
}

func TestGetJobStatusFailureMixture(t *testing.T) {
	progressObj := JobCounters{
		counter: map[string]int64{
			KeyTotalTasks:     5,
			KeyTasksCompleted: 4,
			KeyTasksFailed:    1,
		},
	}
	status := progressObj.GetJobStatus()
	if status != JobFailed {
		t.Errorf("expected job status for %+v to be %d, instead found %d",
			progressObj, JobFailed, status)
	}
}

func TestGetJobStatusFailureAllFails(t *testing.T) {
	progressObj := JobCounters{
		counter: map[string]int64{
			KeyTotalTasks:     5,
			KeyTasksCompleted: 0,
			KeyTasksFailed:    5,
		},
	}
	status := progressObj.GetJobStatus()
	if status != JobFailed {
		t.Errorf("expected job status for %+v to be %d, instead found %d",
			progressObj, JobFailed, status)
	}
}

/*******************************************************************************
updateForTaskUpdate tests, inserting new tasks
*******************************************************************************/

func assertOtherDeltaFieldsUnchangedForInsert(t *testing.T,
	delta *JobCounters) {
	if delta.counter[KeyTasksCompleted] != 0 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 0, found %d",
			delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksFailed] != 0 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be 0, found %d",
			delta.counter[KeyTasksFailed])
	}
}

func TestUpdateForTaskUpdateOneInsertSingleJob(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    listTaskType,
	}
	tu := &TaskUpdate{nil, nil, []*Task{task}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Unqueued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTotalTasks] != 1 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 1, found %d", delta.counter[KeyTotalTasks])
	}
	assertOtherDeltaFieldsUnchangedForInsert(t, delta)
}

func TestUpdateForTaskUpdateMultipleInsertsSingleJob(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task1 := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    listTaskType,
	}
	task2 := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    uploadGCSTaskType,
	}
	task3 := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    loadBQTaskType,
	}
	tu := &TaskUpdate{nil, nil, []*Task{task1, task2, task3}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Unqueued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTotalTasks] != 3 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 3, found %d", delta.counter[KeyTotalTasks])
	}
	assertOtherDeltaFieldsUnchangedForInsert(t, delta)
}

func TestUpdateForTaskUpdateMultipleInsertsMultipleJobs(t *testing.T) {
	id1 := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	id2 := JobRunFullId{
		JobConfigId: "C",
		JobRunId:    "B",
	}
	task1 := &Task{
		JobConfigId: id1.JobConfigId,
		JobRunId:    id1.JobRunId,
		TaskType:    listTaskType,
	}
	task2 := &Task{
		JobConfigId: id1.JobConfigId,
		JobRunId:    id1.JobRunId,
		TaskType:    uploadGCSTaskType,
	}
	task3 := &Task{
		JobConfigId: id1.JobConfigId,
		JobRunId:    id1.JobRunId,
		TaskType:    loadBQTaskType,
	}
	task4 := &Task{
		JobConfigId: id2.JobConfigId,
		JobRunId:    id2.JobRunId,
		TaskType:    listTaskType,
	}
	task5 := &Task{
		JobConfigId: id2.JobConfigId,
		JobRunId:    id2.JobRunId,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{nil, nil, []*Task{task1, task2, task3, task4, task5}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Unqueued)

	if len(counters.deltas) != 2 {
		t.Errorf("expected counters.deltas to contain 2 deltas, found %d",
			len(counters.deltas))
	}
	delta1, exists := counters.deltas[id1]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", id1)
	}
	if delta1.counter[KeyTotalTasks] != 3 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 3, found %d", delta1.counter[KeyTotalTasks])
	}
	assertOtherDeltaFieldsUnchangedForInsert(t, delta1)

	delta2, exists := counters.deltas[id2]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", id2)
	}
	if delta2.counter[KeyTotalTasks] != 2 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 2, found %d", delta2.counter[KeyTotalTasks])
	}
	assertOtherDeltaFieldsUnchangedForInsert(t, delta2)
}

/*******************************************************************************
updateForTaskUpdate tests, updating existing tasks
*******************************************************************************/

func assertOtherDeltaFieldsUnchangedForUpdate(t *testing.T,
	delta *JobCounters) {
	if delta.counter[KeyTotalTasks] != 0 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 0, found %d",
			delta.counter[KeyTotalTasks])
	}
}

func TestUpdateForTaskUpdateQueuedToSuccess(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Success,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{task, nil, []*Task{}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Queued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTasksCompleted] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 1, found %d", delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksFailed] != 0 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be 0, found %d", delta.counter[KeyTasksFailed])
	}
	assertOtherDeltaFieldsUnchangedForUpdate(t, delta)
}

func TestUpdateForTaskUpdateQueuedToSuccessDeltaObjAlreadyExists(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Success,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{task, nil, []*Task{}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Queued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTasksCompleted] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 1, found %d", delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksFailed] != 0 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be 0, found %d", delta.counter[KeyTasksFailed])
	}
	assertOtherDeltaFieldsUnchangedForUpdate(t, delta)
}

func TestUpdateForTaskUpdateFailedToSuccess(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Success,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{task, nil, []*Task{}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Failed)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTasksCompleted] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 1, found %d", delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksFailed] != -1 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be -1, found %d", delta.counter[KeyTasksFailed])
	}
	assertOtherDeltaFieldsUnchangedForUpdate(t, delta)
}

func TestUpdateForTaskUpdateUnqueuedToSuccess(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Success,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{task, nil, []*Task{}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Unqueued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTasksCompleted] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 1, found %d", delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksFailed] != 0 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be 0, found %d", delta.counter[KeyTasksFailed])
	}
	assertOtherDeltaFieldsUnchangedForUpdate(t, delta)
}

func TestUpdateForTaskUpdateUnqueuedToFailed(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Failed,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{task, nil, []*Task{}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Unqueued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	if delta.counter[KeyTasksCompleted] != 0 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 0, found %d", delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksFailed] != 1 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be 1, found %d", delta.counter[KeyTasksFailed])
	}
	assertOtherDeltaFieldsUnchangedForUpdate(t, delta)
}

func TestUpdateForTaskUpdateUnqueuedToQueued(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	task := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Queued,
		TaskType:    uploadGCSTaskType,
	}
	tu := &TaskUpdate{task, nil, []*Task{}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Unqueued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
}

/*******************************************************************************
updateForTaskUpdate tests, updating a task and inserting new tasks
*******************************************************************************/

func TestUpdateForTaskUpdateListTaskNewCopyTasks(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	updatedListTask := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Success,
		TaskType:    listTaskType,
	}
	newCopyTask1 := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    uploadGCSTaskType,
	}
	newCopyTask2 := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    uploadGCSTaskType,
	}
	logEntryData := make(map[string]interface{})
	logEntryData["files_found"] = json.Number("2")
	logEntryData["bytes_found"] = json.Number("12345678")
	logEntryData["file_stat_errors"] = json.Number("1")
	logEntry := NewLogEntry(logEntryData)
	tu := &TaskUpdate{updatedListTask, logEntry, []*Task{newCopyTask1, newCopyTask2}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Queued)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	// Expect one completed list task.
	if delta.counter[KeyTasksCompleted] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 1, found %d",
		delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksCompleted + KeySuffixList] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted + KeySuffixList] to be 1, found %d",
		delta.counter[KeyTasksCompleted + KeySuffixList])
	}
	if delta.counter[KeyTasksQueued] != -1 {
		t.Errorf("expected delta.counter[KeyTasksQueued] to be -1, found %d",
		delta.counter[KeyTasksQueued])
	}
	if delta.counter[KeyTasksQueued + KeySuffixList] != -1 {
		t.Errorf("expected delta.counter[KeyTasksQueued + KeySuffixList] to be -1, found %d",
		delta.counter[KeyTasksQueued + KeySuffixList])
	}
	// Expect the listing counters to exist.
	if delta.counter[KeyListFilesFound] != 2 {
		t.Errorf("expected delta.counter[KeyListFilesFound] to be 12345, found %d",
		delta.counter[KeyListFilesFound])
	}
	if delta.counter[KeyListBytesFound] != 12345678 {
		t.Errorf("expected delta.counter[KeyListBytesFound] to be 12345, found %d",
		delta.counter[KeyListBytesFound])
	}
	if delta.counter[KeyListFileStatErrors] != 1 {
		t.Errorf("expected delta.counter[KeyListFileStatErrors] to be 12345, found %d",
		delta.counter[KeyListFileStatErrors])
	}
	// Expect two new copy tasks.
	if delta.counter[KeyTotalTasks] != 2 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 2, found %d",
		delta.counter[KeyTotalTasks])
	}
	if delta.counter[KeyTotalTasks + KeySuffixCopy] != 2 {
		t.Errorf("expected delta.counter[KeyTotalTasks + KeySuffixCopy] to be 2, found %d",
		delta.counter[KeyTotalTasks + KeySuffixCopy])
	}
}

func TestUpdateForTaskUpdateCopyTaskNewLoadTask(t *testing.T) {
	fullJobId := JobRunFullId{
		JobConfigId: "A",
		JobRunId:    "B",
	}
	updatedCopyTask := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		Status:      Success,
		TaskType:    uploadGCSTaskType,
	}
	newLoadTask := &Task{
		JobConfigId: fullJobId.JobConfigId,
		JobRunId:    fullJobId.JobRunId,
		TaskType:    loadBQTaskType,
	}
	logEntryData := make(map[string]interface{})
	logEntryData["src_bytes"] = json.Number("12345")
	logEntry := NewLogEntry(logEntryData)
	tu := &TaskUpdate{updatedCopyTask, logEntry, []*Task{newLoadTask}}

	var counters JobCountersCollection
	counters.deltas = make(map[JobRunFullId]*JobCounters)
	counters.updateForTaskUpdate(tu, Failed)

	if len(counters.deltas) != 1 {
		t.Errorf("expected counters.deltas to contain 1 delta, found %d",
			len(counters.deltas))
	}
	delta, exists := counters.deltas[fullJobId]
	if !exists {
		t.Errorf("expected counters.deltas to contain a delta for id %+v", fullJobId)
	}
	// Expect one completed copy task.
	if delta.counter[KeyTasksCompleted] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted] to be 1, found %d",
		delta.counter[KeyTasksCompleted])
	}
	if delta.counter[KeyTasksCompleted + KeySuffixCopy] != 1 {
		t.Errorf("expected delta.counter[KeyTasksCompleted + KeySuffixCopy] to be 1, found %d",
		delta.counter[KeyTasksCompleted + KeySuffixCopy])
	}
	if delta.counter[KeyTasksFailed] != -1 {
		t.Errorf("expected delta.counter[KeyTasksFailed] to be -1, found %d",
		delta.counter[KeyTasksFailed])
	}
	if delta.counter[KeyTasksFailed + KeySuffixCopy] != -1 {
		t.Errorf("expected delta.counter[KeyTasksFailed + KeySuffixCopy] to be -1, found %d",
		delta.counter[KeyTasksFailed + KeySuffixCopy])
	}
	// Expect the bytes copied counter to exist.
	if delta.counter[KeyBytesCopied] != 12345 {
		t.Errorf("expected delta.counter[KeyBytesCopied] to be 12345, found %d",
		delta.counter[KeyBytesCopied])
	}
	// Expect two new copy tasks.
	if delta.counter[KeyTotalTasks] != 1 {
		t.Errorf("expected delta.counter[KeyTotalTasks] to be 1, found %d",
		delta.counter[KeyTotalTasks])
	}
	if delta.counter[KeyTotalTasks + KeySuffixLoad] != 1 {
		t.Errorf("expected delta.counter[KeyTotalTasks + KeySuffixLoad] to be 1, found %d",
		delta.counter[KeyTotalTasks + KeySuffixLoad])
	}
}
