// Copyright 2020 Comcast Cable Communications Management, LLC
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

package main

import (
	"github.com/xmidt-org/ears/internal/pkg/syncer"
	"github.com/xmidt-org/ears/pkg/filter"
	pkgplugin "github.com/xmidt-org/ears/pkg/plugin"
	"github.com/xmidt-org/ears/pkg/secret"
	"github.com/xmidt-org/ears/pkg/tenant"

	"github.com/xmidt-org/ears/pkg/event"
)

func main() {
	// required for `go build` to not fail
}

var Plugin, PluginErr = NewPlugin()

// for golangci-lint
var _ = Plugin
var _ = PluginErr

var _ filter.Filterer = (*plugin)(nil)

type plugin struct{}

const (
	Name     = "name"
	Version  = "version"
	CommitID = "commitID"
)

// ===================================================

func NewPlugin() (*pkgplugin.Plugin, error) {
	return NewPluginVersion(Name, Version, CommitID)
}

func NewPluginVersion(name string, version string, commitID string) (*pkgplugin.Plugin, error) {
	return pkgplugin.NewPlugin(
		pkgplugin.WithName(name),
		pkgplugin.WithVersion(version),
		pkgplugin.WithCommitID(commitID),
		pkgplugin.WithNewFilterer(NewFilterer),
	)
}

// Filterer ============================================================

func NewFilterer(tid tenant.Id, pluginType string, name string, config interface{}, secrets secret.Vault, tableSyncer syncer.DeltaSyncer) (filter.Filterer, error) {

	return &plugin{}, nil
}

func (p *plugin) Filter(e event.Event) []event.Event {
	return nil
}

func (p *plugin) Config() interface{} {
	return nil
}

func (p *plugin) Name() string {
	return ""
}

func (p *plugin) Plugin() string {
	return "filter"
}

func (p *plugin) Tenant() tenant.Id {
	return tenant.Id{OrgId: "myorg", AppId: "myapp"}
}

func (p *plugin) EventSuccessCount() int {
	return 0
}

func (p *plugin) EventSuccessVelocity() int {
	return 0
}

func (p *plugin) EventFilterCount() int {
	return 0
}

func (p *plugin) EventFilterVelocity() int {
	return 0
}

func (p *plugin) EventErrorCount() int {
	return 0
}

func (p *plugin) EventErrorVelocity() int {
	return 0
}

func (p *plugin) EventTs() int64 {
	return 0
}
