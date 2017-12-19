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
	"context"
	"errors"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/cloud-ingest/gcloud"
)

const MTIME_ATTR_NAME string = "goog-reserved-file-mtime"

// Type encapsulating the object metadata we care about.
type ObjectMetadata struct {
	Size             int64
	Mtime            int64
	GenerationNumber int64
	MD5              []byte
}

// ObjectMetadataReader is a simple interface around reading object metadata from GCS objects.
type ObjectMetadataReader interface {
	GetMetadata(ctx context.Context, bucketName string, objectName string) (*ObjectMetadata, error)
}

type GCSObjectMetadataReader struct {
	gcs gcloud.GCS
}

// NewGCSObjectMetadataReader constructs an ObjectMetadataReader that calls into GCS.
func NewGCSObjectMetadataReader(gcs gcloud.GCS) *GCSObjectMetadataReader {
	return &GCSObjectMetadataReader{gcs}
}

func GCSAttrToObjectMetadata(attrs *storage.ObjectAttrs) (*ObjectMetadata, error) {
	if attrs == nil {
		return nil, errors.New("nil object attributes passed")
	}

	mtimeStr, ok := attrs.Metadata[MTIME_ATTR_NAME]
	var mtime int64
	if ok {
		var err error
		mtime, err = strconv.ParseInt(mtimeStr, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &ObjectMetadata{Size: attrs.Size, GenerationNumber: attrs.Generation, Mtime: mtime, MD5: attrs.MD5}, nil
}

// GetMetadata retrieves metadata for an object.
// When an object does not exist, the "not found" error is propagated, as with all GCS errors.
func (r *GCSObjectMetadataReader) GetMetadata(ctx context.Context, bucketName string, objectName string) (*ObjectMetadata, error) {
	attr, err := r.gcs.GetAttrs(ctx, bucketName, objectName)
	if err != nil {
		return nil, err
	}

	return GCSAttrToObjectMetadata(attr)
}
