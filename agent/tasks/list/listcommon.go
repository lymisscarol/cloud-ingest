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

package list

import (
	"context"
	"errors"
	"flag"
	"sort"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/gcloud"
	"github.com/GoogleCloudPlatform/cloud-ingest/agent/tasks/common"

	listfilepb "github.com/GoogleCloudPlatform/cloud-ingest/proto/listfile_go_proto"
	taskpb "github.com/GoogleCloudPlatform/cloud-ingest/proto/task_go_proto"
)

var (
	// NumberConcurrentListTasks is used to both determine listing memory constraints, and set the MaxOutstandingMessages for the PubSub List subscription.
	NumberConcurrentListTasks = flag.Int("number-concurrent-list-tasks", 4, "The maximum number of list tasks the agent will process at any given time.")

	listFileSizeThreshold          = flag.Int("list-file-size-threshold", 50000, "List tasks will keep listing directories until the number of listed files and directories exceeds this threshold, or until there are no more files/directories to list")
	listTaskChunkSize              = flag.Int("list-task-chunk-size", 8*1024*1024, "The resumable upload chunk size used for list tasks, defaults to 8MiB.")
	maxMemoryForListingDirectories = flag.Int("max-memory-for-listing-directories", 20, "Maximum amount of memory agent will use in total (not per task) to store directories before writing them to a list file. Value is in MiB.")
)

type listingFileMetadata struct {
	bytes, files, dirsDiscovered, dirsListed, dirsNotListed int64
	dirsNotFound                                            []string
}

func dirInfoEntry(path string) *listfilepb.ListFileEntry {
	return &listfilepb.ListFileEntry{
		Entry: &listfilepb.ListFileEntry_DirectoryInfo{
			DirectoryInfo: &listfilepb.DirectoryInfo{
				Path: path,
			},
		},
	}
}

func fileInfoEntry(path string, lastModTime, size int64) *listfilepb.ListFileEntry {
	return &listfilepb.ListFileEntry{
		Entry: &listfilepb.ListFileEntry_FileInfo{
			FileInfo: &listfilepb.FileInfo{
				Path:             path,
				LastModifiedTime: lastModTime,
				Size:             size,
			},
		},
	}
}

func setListLog(log *taskpb.Log, listMD *listingFileMetadata) {
	ll := log.GetListLog()
	ll.FilesFound = listMD.files
	ll.BytesFound = listMD.bytes
	ll.DirsFound = listMD.dirsDiscovered
	ll.DirsListed = listMD.dirsListed
	ll.DirsNotListed = listMD.dirsNotListed
	ll.DirsNotFound = listMD.dirsNotFound
}

func gcsWriterWithCondition(ctx context.Context, gcs gcloud.GCS, bucket, object string, generationNum int64, resumableChunkSize int) gcloud.WriteCloserWithError {
	w := gcs.NewWriterWithCondition(ctx, bucket, object, common.GetGCSGenerationNumCondition(generationNum))
	// Set the resumable upload chunk size.
	if t, ok := w.(*storage.Writer); ok {
		t.ChunkSize = resumableChunkSize
	}
	return w
}

func sortListFileEntries(entries []*listfilepb.ListFileEntry) error {
	// Readdir returns the entries in "directory order", so they must be sorted
	// to meet our expectations of lexicographical order.
	var err error
	sort.Slice(entries, func(i, j int) bool {
		iPath, e := getPath(entries[i])
		if e != nil {
			err = e
		}
		jPath, e := getPath(entries[j])
		if e != nil {
			err = e
		}
		return iPath < jPath
	})
	return err
}

func getPath(entry *listfilepb.ListFileEntry) (string, error) {
	if dirInfo := entry.GetDirectoryInfo(); dirInfo != nil {
		return dirInfo.Path, nil
	} else if fileInfo := entry.GetFileInfo(); fileInfo != nil {
		return fileInfo.Path, nil
	} else {
		return "", errors.New("unknown list file entry type")
	}
}
