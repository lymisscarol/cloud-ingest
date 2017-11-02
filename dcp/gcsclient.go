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
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"io"
)

// Pass-through wrapper for Google Cloud Storage client.
type GCS interface {
	GetAttrs(bucketName string, objectName string) (*storage.ObjectAttrs, error)
	NewReader(bucketName string, objectName string) (io.ReadCloser, error)
	NewWriter(bucketName string, objectName string) io.WriteCloser
}

type GCSClient struct {
	client *storage.Client
}

func NewGCSClient(client *storage.Client) *GCSClient {
	return &GCSClient{client}
}

func (gcs *GCSClient) GetAttrs(bucketName string, objectName string) (*storage.ObjectAttrs, error) {
	return gcs.client.Bucket(bucketName).Object(objectName).Attrs(context.Background())
}

func (gcs *GCSClient) NewReader(bucketName string, objectName string) (io.ReadCloser, error) {
	return gcs.client.Bucket(bucketName).Object(objectName).NewReader(context.Background())
}

func (gcs *GCSClient) NewWriter(bucketName string, objectName string) io.WriteCloser {
	return gcs.client.Bucket(bucketName).Object(objectName).NewWriter(context.Background())
}
