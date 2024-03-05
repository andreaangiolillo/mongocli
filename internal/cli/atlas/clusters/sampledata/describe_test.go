// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package sampledata

import (
	"bytes"
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/mocks"
	"github.com/andreaangiolillo/mongocli-test/internal/pointer"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSampleDataStatusDescriber(ctrl)

	expected := &atlasv2.SampleDatasetStatus{
		Id:          pointer.Get("test"),
		ClusterName: pointer.Get("ClusterTest"),
		State:       pointer.Get("COMPLETED"),
	}

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		id:    "test",
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		SampleDataStatus(describeOpts.ProjectID, describeOpts.id).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID},
	)
}
