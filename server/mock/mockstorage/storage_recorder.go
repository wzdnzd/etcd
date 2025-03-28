// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockstorage

import (
	"github.com/coreos/go-semver/semver"

	"go.etcd.io/etcd/client/pkg/v3/testutil"
	"go.etcd.io/raft/v3"
	"go.etcd.io/raft/v3/raftpb"
)

type StorageRecorder struct {
	testutil.Recorder
	dbPath string // must have '/' suffix if set
}

func NewStorageRecorder(db string) *StorageRecorder {
	return &StorageRecorder{&testutil.RecorderBuffered{}, db}
}

func NewStorageRecorderStream(db string) *StorageRecorder {
	return &StorageRecorder{testutil.NewRecorderStream(), db}
}

func (p *StorageRecorder) Save(st raftpb.HardState, ents []raftpb.Entry) error {
	p.Record(testutil.Action{Name: "Save"})
	return nil
}

func (p *StorageRecorder) SaveSnap(st raftpb.Snapshot) error {
	if !raft.IsEmptySnap(st) {
		p.Record(testutil.Action{Name: "SaveSnap"})
	}
	return nil
}

func (p *StorageRecorder) Release(st raftpb.Snapshot) error {
	if !raft.IsEmptySnap(st) {
		p.Record(testutil.Action{Name: "Release"})
	}
	return nil
}

func (p *StorageRecorder) Sync() error {
	p.Record(testutil.Action{Name: "Sync"})
	return nil
}

func (p *StorageRecorder) Close() error                        { return nil }
func (p *StorageRecorder) MinimalEtcdVersion() *semver.Version { return nil }
