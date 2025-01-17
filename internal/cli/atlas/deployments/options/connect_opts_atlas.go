// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package options

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
)

type ConnectToAtlasOpts struct {
	cli.GlobalOpts
	cli.InputOpts
	DBUsername           string
	DBUserPassword       string
	ConnectionStringType string
	Store                store.AtlasClusterDescriber
}

func (opts *ConnectToAtlasOpts) InitAtlasStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.Store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ConnectOpts) validateAndPromptAtlasOpts() error {
	requiresAuth := opts.ConnectWith == MongoshConnect || opts.ConnectWith == CompassConnect
	if requiresAuth && opts.DBUsername == "" {
		if err := opts.promptDBUsername(); err != nil {
			return err
		}
	}

	if requiresAuth && opts.DBUserPassword == "" {
		if err := opts.promptDBUserPassword(); err != nil {
			return err
		}
	}

	return opts.validateAndPromptConnectionStringType()
}

func (opts *ConnectToAtlasOpts) validateAndPromptConnectionStringType() error {
	if opts.ConnectionStringType == "" {
		p := &survey.Select{
			Message: promptConnectionStringType,
			Options: connectionStringTypeOptions,
			Help:    usage.ConnectionStringType,
		}

		err := telemetry.TrackAskOne(p, &opts.ConnectionStringType, nil)
		if err != nil {
			return err
		}
	}

	if !search.StringInSliceFold(connectionStringTypeOptions, opts.ConnectionStringType) {
		return fmt.Errorf("%w: %s", errConnectionStringTypeNotImplemented, opts.ConnectionStringType)
	}

	return nil
}

func (opts *ConnectToAtlasOpts) promptDBUsername() error {
	p := &survey.Input{
		Message: "Username for authenticating to MongoDB deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DBUsername)
}

func (opts *ConnectToAtlasOpts) promptDBUserPassword() error {
	if !opts.IsTerminalInput() {
		_, err := fmt.Fscanln(opts.InReader, &opts.DBUserPassword)
		return err
	}

	p := &survey.Password{
		Message: "Password for authenticating to MongoDB deployment",
	}
	return telemetry.TrackAskOne(p, &opts.DBUserPassword)
}

func (opts *ConnectOpts) connectToAtlas() error {
	r, err := opts.Store.AtlasCluster(opts.ConfigProjectID(), opts.DeploymentName)
	if err != nil {
		return err
	}

	if opts.ConnectionStringType == connectionStringTypePrivate {
		if r.ConnectionStrings.PrivateSrv == nil {
			return errNetworkPeeringConnectionNotConfigured
		}
		return opts.connectToDeployment(*r.ConnectionStrings.PrivateSrv)
	}

	return opts.connectToDeployment(*r.ConnectionStrings.StandardSrv)
}
