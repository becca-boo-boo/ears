// Copyright 2021 Comcast Cable Communications Management, LLC
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

package http

import (
	"github.com/rs/zerolog"
	"github.com/xmidt-org/ears/pkg/errs"
	pkgplugin "github.com/xmidt-org/ears/pkg/plugin"
	"github.com/xmidt-org/ears/pkg/receiver"
	"github.com/xmidt-org/ears/pkg/sender"
	"go.opentelemetry.io/otel/metric"
	"net/http"
)

var _ sender.Sender = (*Sender)(nil)
var _ receiver.Receiver = (*Receiver)(nil)

var (
	Name     = "http"
	Version  = "v0.0.0"
	CommitID = ""
)

func NewPlugin() (*pkgplugin.Plugin, error) {
	return NewPluginVersion(Name, Version, CommitID)
}

func NewPluginVersion(name string, version string, commitID string) (*pkgplugin.Plugin, error) {
	return pkgplugin.NewPlugin(
		pkgplugin.WithName(name),
		pkgplugin.WithVersion(version),
		pkgplugin.WithCommitID(commitID),
		pkgplugin.WithNewReceiver(NewReceiver),
		pkgplugin.WithNewSender(NewSender),
	)
}

type ReceiverConfig struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Port   *int   `json:"port"`
	Trace  *bool  `json:"trace,omitempty"`
}

type Receiver struct {
	logger              *zerolog.Logger
	srv                 *http.Server
	config              ReceiverConfig
	eventSuccessCounter metric.BoundFloat64Counter
	eventFailureCounter metric.BoundFloat64Counter
	eventBytesCounter   metric.BoundInt64Counter
}

type SenderConfig struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type Sender struct {
	client *http.Client
	config SenderConfig
}

type BadHttpStatusError struct {
	statusCode int
}

func (e *BadHttpStatusError) Error() string {
	return errs.String("BadHttpStatusError", map[string]interface{}{"statusCode": e.statusCode}, nil)
}
