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

package sampledata

import (
	"context"
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/store"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	id    string
	store store.SampleDataStatusDescriber
}

var watchTemplate = "\nSample data available.\n"

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (bool, error) {
	result, err := opts.store.SampleDataStatus(opts.ConfigProjectID(), opts.id)
	if err != nil {
		return false, err
	}
	return result.State != nil && *result.State == "COMPLETED", nil
}

func (opts *WatchOpts) Run() error {
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// atlas cluster(s) sampledata watch <id> [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <id>",
		Short: "Watch the specified sample data job in your cluster until it completes.",
		Long: `This command checks the sample data job's status periodically until it reaches an COMPLETED state. 
If you run the command in the terminal, it blocks the terminal session until the resource state changes to COMPLETED.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: fmt.Sprintf(`  # Watch for the sample data job with ID 5e2211c17a3e5a48f5497de3 to complete:
  %s clusters sampledata watch 5e2211c17a3e5a48f5497de3`, cli.ExampleAtlasEntryPoint()),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"idDesc": "Unique identifier of the sample data job.",
			"output": watchTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), watchTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
