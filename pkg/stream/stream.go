/*
Copyright 2019 The Kubernetes Authors.

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

package stream

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"

	"github.com/WanLinghao/coredump-detector/pkg/backend"
	"github.com/WanLinghao/coredump-detector/pkg/backend/types"
)

// CoredumpStreamer is a resource that streams the contents of a particular
// location URL.
type CoredumpStreamer struct {
	Namespace     string
	PodUID        string
	ContainerName string

	// storage handles core file download
	storage types.Storage
}

func NewCoredumpStreamer(ns, podUID, containerName string) (*CoredumpStreamer, error) {
	return &CoredumpStreamer{
		Namespace:     ns,
		PodUID:        podUID,
		ContainerName: containerName,
		storage:       backend.GetBackendStorage(),
	}, nil
}

// a CoredumpStreamer must implement a rest.ResourceStreamer
var _ rest.ResourceStreamer = &CoredumpStreamer{}

func (obj *CoredumpStreamer) GetObjectKind() schema.ObjectKind {
	return schema.EmptyObjectKind
}
func (obj *CoredumpStreamer) DeepCopyObject() runtime.Object {
	panic("rest.LocationStreamer does not implement DeepCopyObject")
}

func (s *CoredumpStreamer) InputStream(ctx context.Context, apiVersion, acceptHeader string) (stream io.ReadCloser, flush bool, contentType string, err error) {
	tarFilePath, err := s.storage.GetCoreFiles(s.Namespace, s.PodUID, s.ContainerName)
	if os.IsNotExist(err) {
		return nil, true, "text/plain", fmt.Errorf("container %s in pod %s/%s has no core dump files stored", s.ContainerName, s.Namespace, s.PodUID)
	}
	if err != nil {
		return nil, true, "text/plain", err
	}

	f, err := os.Open(tarFilePath)
	if err != nil {
		return nil, true, "text/plain", err
	}
	stream = ioutil.NopCloser(bufio.NewReader(f)) // r type is io.ReadCloser
	flush = true
	contentType = "application/x-gtar"

	err = nil
	return
}
