// Copyright 2023 MongoDB Inc
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

// This code was autogenerated at 2023-05-04T11:20:25+01:00. Note: Manual updates are allowed, but may be overwritten.

package datasets

import (
	"context"
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	store "github.com/andreaangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store        store.PipelineDatasetDeleter
	pipelineName string
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeletePipelineDataset, opts.ConfigProjectID(), opts.pipelineName)
}

// atlas dataLakePipelines datasets delete <pipelineRunId> [--projectId projectId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("'%s' deleted\n", "Not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <pipelineRunId>",
		Short: "Remove the specified data lake pipeline dataset from your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"pipelineRunIdDesc": "Unique identifier for the data lake pipeline run",
			"output":            opts.SuccessMessage(),
		},
		Example: `# delete dataset '507f1f77bcf86cd799439011' for data lake pipeline called 'Pipeline1':
  atlas dataLakePipelines datasets delete 507f1f77bcf86cd799439011 --pipeline Pipeline1
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.pipelineName, flag.Pipeline, "", usage.Pipeline)
	_ = cmd.MarkFlagRequired(flag.Pipeline)

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
