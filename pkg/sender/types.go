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

package sender

import (
	"context"
	"github.com/xmidt-org/ears/internal/pkg/syncer"
	"github.com/xmidt-org/ears/pkg/event"
	"github.com/xmidt-org/ears/pkg/secret"
	"github.com/xmidt-org/ears/pkg/tenant"
)

//go:generate rm -f testing_mock.go
//go:generate moq -out testing_mock.go . Hasher NewSenderer Sender

type InvalidConfigError struct {
	Err error
}

type Hasher interface {
	// SenderHash calculates the hash of a sender based on the given configuration
	SenderHash(config interface{}) (string, error)
}

type NewSenderer interface {
	Hasher
	// Returns an objec that implements the Sender interface
	NewSender(tid tenant.Id, plugin string, name string, config interface{}, secrets secret.Vault, tableSyncer syncer.DeltaSyncer) (Sender, error)
}

// or Outputter[√] or Producer[x] or Publisher[√]
type Sender interface {
	// Send consumes and event and sends it to the target
	Send(e event.Event)
	// Unwrap
	Unwrap() Sender
	// StopSending will stop any long running maintenance threads in the sender plugin.
	// Often this function does nothing.
	StopSending(ctx context.Context)
	//
	Config() interface{}
	Name() string
	Plugin() string
	Tenant() tenant.Id
	EventSuccessCount() int
	EventSuccessVelocity() int
	EventErrorCount() int
	EventErrorVelocity() int
	EventTs() int64
}
