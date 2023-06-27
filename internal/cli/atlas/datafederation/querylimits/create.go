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

// This code was autogenerated at 2023-06-23T15:50:57+01:00. Note: Manual updates are allowed, but may be overwritten.

package querylimits

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store         store.DataFederationQueryLimitCreator
	limitName     string
	tenantName    string
	value         int64
	overrunPolicy string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = `Data federation query limit {{.Name}} created.`

func (opts *CreateOpts) Run() error {
	createRequest := opts.newCreateRequest()

	r, err := opts.store.CreateDataFederationQueryLimit(opts.ConfigProjectID(), opts.tenantName, opts.limitName, createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCreateRequest() *admin.DataFederationTenantQueryLimit {
	req := &admin.DataFederationTenantQueryLimit{
		Name:       opts.limitName,
		TenantName: &opts.tenantName,
		Value:      opts.value,
	}
	if opts.overrunPolicy != "" {
		req.OverrunPolicy = &opts.overrunPolicy
	}
	return req
}

// atlas dataFederation queryLimits create <name> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Creates a new Data Federation query limit.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Identifier of the data federation query limit",
			"output":   createTemplate,
		},
		Example: `# create data federation query limit:
  atlas dataFederation queryLimit create bytesProcessed.query --value 1000 --dataFederation DataFederation1
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.limitName = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().Int64Var(&opts.value, flag.Value, 0, usage.Value)
	cmd.Flags().StringVar(&opts.overrunPolicy, flag.OverrunPolicy, "", usage.OverrunPolicy)
	cmd.Flags().StringVar(&opts.tenantName, flag.DataFederation, "", usage.DataFederation)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.DataFederation)
	_ = cmd.MarkFlagRequired(flag.Value)

	return cmd
}
