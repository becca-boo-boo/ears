// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sender

import (
	"context"
	"github.com/xmidt-org/ears/pkg/event"
	"github.com/xmidt-org/ears/pkg/secret"
	"github.com/xmidt-org/ears/pkg/tenant"
	"sync"
)

// Ensure, that HasherMock does implement Hasher.
// If this is not the case, regenerate this file with moq.
var _ Hasher = &HasherMock{}

// HasherMock is a mock implementation of Hasher.
//
// 	func TestSomethingThatUsesHasher(t *testing.T) {
//
// 		// make and configure a mocked Hasher
// 		mockedHasher := &HasherMock{
// 			SenderHashFunc: func(config interface{}) (string, error) {
// 				panic("mock out the SenderHash method")
// 			},
// 		}
//
// 		// use mockedHasher in code that requires Hasher
// 		// and then make assertions.
//
// 	}
type HasherMock struct {
	// SenderHashFunc mocks the SenderHash method.
	SenderHashFunc func(config interface{}) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// SenderHash holds details about calls to the SenderHash method.
		SenderHash []struct {
			// Config is the config argument value.
			Config interface{}
		}
	}
	lockSenderHash sync.RWMutex
}

// SenderHash calls SenderHashFunc.
func (mock *HasherMock) SenderHash(config interface{}) (string, error) {
	if mock.SenderHashFunc == nil {
		panic("HasherMock.SenderHashFunc: method is nil but Hasher.SenderHash was just called")
	}
	callInfo := struct {
		Config interface{}
	}{
		Config: config,
	}
	mock.lockSenderHash.Lock()
	mock.calls.SenderHash = append(mock.calls.SenderHash, callInfo)
	mock.lockSenderHash.Unlock()
	return mock.SenderHashFunc(config)
}

// SenderHashCalls gets all the calls that were made to SenderHash.
// Check the length with:
//     len(mockedHasher.SenderHashCalls())
func (mock *HasherMock) SenderHashCalls() []struct {
	Config interface{}
} {
	var calls []struct {
		Config interface{}
	}
	mock.lockSenderHash.RLock()
	calls = mock.calls.SenderHash
	mock.lockSenderHash.RUnlock()
	return calls
}

// Ensure, that NewSendererMock does implement NewSenderer.
// If this is not the case, regenerate this file with moq.
var _ NewSenderer = &NewSendererMock{}

// NewSendererMock is a mock implementation of NewSenderer.
//
// 	func TestSomethingThatUsesNewSenderer(t *testing.T) {
//
// 		// make and configure a mocked NewSenderer
// 		mockedNewSenderer := &NewSendererMock{
// 			NewSenderFunc: func(tid tenant.Id, plugin string, name string, config interface{}, secrets secret.Vault) (Sender, error) {
// 				panic("mock out the NewSender method")
// 			},
// 			SenderHashFunc: func(config interface{}) (string, error) {
// 				panic("mock out the SenderHash method")
// 			},
// 		}
//
// 		// use mockedNewSenderer in code that requires NewSenderer
// 		// and then make assertions.
//
// 	}
type NewSendererMock struct {
	// NewSenderFunc mocks the NewSender method.
	NewSenderFunc func(tid tenant.Id, plugin string, name string, config interface{}, secrets secret.Vault) (Sender, error)

	// SenderHashFunc mocks the SenderHash method.
	SenderHashFunc func(config interface{}) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// NewSender holds details about calls to the NewSender method.
		NewSender []struct {
			// Tid is the tid argument value.
			Tid tenant.Id
			// Plugin is the plugin argument value.
			Plugin string
			// Name is the name argument value.
			Name string
			// Config is the config argument value.
			Config interface{}
			// Secrets is the secrets argument value.
			Secrets secret.Vault
		}
		// SenderHash holds details about calls to the SenderHash method.
		SenderHash []struct {
			// Config is the config argument value.
			Config interface{}
		}
	}
	lockNewSender  sync.RWMutex
	lockSenderHash sync.RWMutex
}

// NewSender calls NewSenderFunc.
func (mock *NewSendererMock) NewSender(tid tenant.Id, plugin string, name string, config interface{}, secrets secret.Vault) (Sender, error) {
	if mock.NewSenderFunc == nil {
		panic("NewSendererMock.NewSenderFunc: method is nil but NewSenderer.NewSender was just called")
	}
	callInfo := struct {
		Tid     tenant.Id
		Plugin  string
		Name    string
		Config  interface{}
		Secrets secret.Vault
	}{
		Tid:     tid,
		Plugin:  plugin,
		Name:    name,
		Config:  config,
		Secrets: secrets,
	}
	mock.lockNewSender.Lock()
	mock.calls.NewSender = append(mock.calls.NewSender, callInfo)
	mock.lockNewSender.Unlock()
	return mock.NewSenderFunc(tid, plugin, name, config, secrets)
}

// NewSenderCalls gets all the calls that were made to NewSender.
// Check the length with:
//     len(mockedNewSenderer.NewSenderCalls())
func (mock *NewSendererMock) NewSenderCalls() []struct {
	Tid     tenant.Id
	Plugin  string
	Name    string
	Config  interface{}
	Secrets secret.Vault
} {
	var calls []struct {
		Tid     tenant.Id
		Plugin  string
		Name    string
		Config  interface{}
		Secrets secret.Vault
	}
	mock.lockNewSender.RLock()
	calls = mock.calls.NewSender
	mock.lockNewSender.RUnlock()
	return calls
}

// SenderHash calls SenderHashFunc.
func (mock *NewSendererMock) SenderHash(config interface{}) (string, error) {
	if mock.SenderHashFunc == nil {
		panic("NewSendererMock.SenderHashFunc: method is nil but NewSenderer.SenderHash was just called")
	}
	callInfo := struct {
		Config interface{}
	}{
		Config: config,
	}
	mock.lockSenderHash.Lock()
	mock.calls.SenderHash = append(mock.calls.SenderHash, callInfo)
	mock.lockSenderHash.Unlock()
	return mock.SenderHashFunc(config)
}

// SenderHashCalls gets all the calls that were made to SenderHash.
// Check the length with:
//     len(mockedNewSenderer.SenderHashCalls())
func (mock *NewSendererMock) SenderHashCalls() []struct {
	Config interface{}
} {
	var calls []struct {
		Config interface{}
	}
	mock.lockSenderHash.RLock()
	calls = mock.calls.SenderHash
	mock.lockSenderHash.RUnlock()
	return calls
}

// Ensure, that SenderMock does implement Sender.
// If this is not the case, regenerate this file with moq.
var _ Sender = &SenderMock{}

// SenderMock is a mock implementation of Sender.
//
// 	func TestSomethingThatUsesSender(t *testing.T) {
//
// 		// make and configure a mocked Sender
// 		mockedSender := &SenderMock{
// 			ConfigFunc: func() interface{} {
// 				panic("mock out the Config method")
// 			},
// 			NameFunc: func() string {
// 				panic("mock out the Name method")
// 			},
// 			PluginFunc: func() string {
// 				panic("mock out the Plugin method")
// 			},
// 			SendFunc: func(e event.Event)  {
// 				panic("mock out the Send method")
// 			},
// 			StopSendingFunc: func(ctx context.Context)  {
// 				panic("mock out the StopSending method")
// 			},
// 			TenantFunc: func() tenant.Id {
// 				panic("mock out the Tenant method")
// 			},
// 			UnwrapFunc: func() Sender {
// 				panic("mock out the Unwrap method")
// 			},
// 		}
//
// 		// use mockedSender in code that requires Sender
// 		// and then make assertions.
//
// 	}
type SenderMock struct {
	// ConfigFunc mocks the Config method.
	ConfigFunc func() interface{}

	// NameFunc mocks the Name method.
	NameFunc func() string

	// PluginFunc mocks the Plugin method.
	PluginFunc func() string

	// SendFunc mocks the Send method.
	SendFunc func(e event.Event)

	// StopSendingFunc mocks the StopSending method.
	StopSendingFunc func(ctx context.Context)

	// TenantFunc mocks the Tenant method.
	TenantFunc func() tenant.Id

	// UnwrapFunc mocks the Unwrap method.
	UnwrapFunc func() Sender

	// calls tracks calls to the methods.
	calls struct {
		// Config holds details about calls to the Config method.
		Config []struct {
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Plugin holds details about calls to the Plugin method.
		Plugin []struct {
		}
		// Send holds details about calls to the Send method.
		Send []struct {
			// E is the e argument value.
			E event.Event
		}
		// StopSending holds details about calls to the StopSending method.
		StopSending []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Tenant holds details about calls to the Tenant method.
		Tenant []struct {
		}
		// Unwrap holds details about calls to the Unwrap method.
		Unwrap []struct {
		}
	}
	lockConfig      sync.RWMutex
	lockName        sync.RWMutex
	lockPlugin      sync.RWMutex
	lockSend        sync.RWMutex
	lockStopSending sync.RWMutex
	lockTenant      sync.RWMutex
	lockUnwrap      sync.RWMutex
}

// Config calls ConfigFunc.
func (mock *SenderMock) Config() interface{} {
	if mock.ConfigFunc == nil {
		panic("SenderMock.ConfigFunc: method is nil but Sender.Config was just called")
	}
	callInfo := struct {
	}{}
	mock.lockConfig.Lock()
	mock.calls.Config = append(mock.calls.Config, callInfo)
	mock.lockConfig.Unlock()
	return mock.ConfigFunc()
}

// ConfigCalls gets all the calls that were made to Config.
// Check the length with:
//     len(mockedSender.ConfigCalls())
func (mock *SenderMock) ConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockConfig.RLock()
	calls = mock.calls.Config
	mock.lockConfig.RUnlock()
	return calls
}

func (mock *SenderMock) EventSuccessCount() int {
	return 0
}

func (mock *SenderMock) EventSuccessVelocity() int {
	return 0
}

func (mock *SenderMock) EventErrorCount() int {
	return 0
}

func (mock *SenderMock) EventErrorVelocity() int {
	return 0
}

func (mock *SenderMock) EventTs() int64 {
	return 0
}

// Name calls NameFunc.
func (mock *SenderMock) Name() string {
	if mock.NameFunc == nil {
		panic("SenderMock.NameFunc: method is nil but Sender.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedSender.NameCalls())
func (mock *SenderMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Plugin calls PluginFunc.
func (mock *SenderMock) Plugin() string {
	if mock.PluginFunc == nil {
		panic("SenderMock.PluginFunc: method is nil but Sender.Plugin was just called")
	}
	callInfo := struct {
	}{}
	mock.lockPlugin.Lock()
	mock.calls.Plugin = append(mock.calls.Plugin, callInfo)
	mock.lockPlugin.Unlock()
	return mock.PluginFunc()
}

// PluginCalls gets all the calls that were made to Plugin.
// Check the length with:
//     len(mockedSender.PluginCalls())
func (mock *SenderMock) PluginCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockPlugin.RLock()
	calls = mock.calls.Plugin
	mock.lockPlugin.RUnlock()
	return calls
}

// Send calls SendFunc.
func (mock *SenderMock) Send(e event.Event) {
	if mock.SendFunc == nil {
		panic("SenderMock.SendFunc: method is nil but Sender.Send was just called")
	}
	callInfo := struct {
		E event.Event
	}{
		E: e,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	mock.SendFunc(e)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//     len(mockedSender.SendCalls())
func (mock *SenderMock) SendCalls() []struct {
	E event.Event
} {
	var calls []struct {
		E event.Event
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}

// StopSending calls StopSendingFunc.
func (mock *SenderMock) StopSending(ctx context.Context) {
	if mock.StopSendingFunc == nil {
		panic("SenderMock.StopSendingFunc: method is nil but Sender.StopSending was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockStopSending.Lock()
	mock.calls.StopSending = append(mock.calls.StopSending, callInfo)
	mock.lockStopSending.Unlock()
	mock.StopSendingFunc(ctx)
}

// StopSendingCalls gets all the calls that were made to StopSending.
// Check the length with:
//     len(mockedSender.StopSendingCalls())
func (mock *SenderMock) StopSendingCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockStopSending.RLock()
	calls = mock.calls.StopSending
	mock.lockStopSending.RUnlock()
	return calls
}

// Tenant calls TenantFunc.
func (mock *SenderMock) Tenant() tenant.Id {
	if mock.TenantFunc == nil {
		panic("SenderMock.TenantFunc: method is nil but Sender.Tenant was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTenant.Lock()
	mock.calls.Tenant = append(mock.calls.Tenant, callInfo)
	mock.lockTenant.Unlock()
	return mock.TenantFunc()
}

// TenantCalls gets all the calls that were made to Tenant.
// Check the length with:
//     len(mockedSender.TenantCalls())
func (mock *SenderMock) TenantCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTenant.RLock()
	calls = mock.calls.Tenant
	mock.lockTenant.RUnlock()
	return calls
}

// Unwrap calls UnwrapFunc.
func (mock *SenderMock) Unwrap() Sender {
	if mock.UnwrapFunc == nil {
		panic("SenderMock.UnwrapFunc: method is nil but Sender.Unwrap was just called")
	}
	callInfo := struct {
	}{}
	mock.lockUnwrap.Lock()
	mock.calls.Unwrap = append(mock.calls.Unwrap, callInfo)
	mock.lockUnwrap.Unlock()
	return mock.UnwrapFunc()
}

// UnwrapCalls gets all the calls that were made to Unwrap.
// Check the length with:
//     len(mockedSender.UnwrapCalls())
func (mock *SenderMock) UnwrapCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockUnwrap.RLock()
	calls = mock.calls.Unwrap
	mock.lockUnwrap.RUnlock()
	return calls
}
