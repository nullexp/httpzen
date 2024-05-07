package protocol

import "github.com/nullexp/httpzen/misc"

type (
	DuplexHandler           func(msg DuplexMessage)
	DuplexHandlerDefinition struct {
		Topic          string
		Dto            any
		Handler        DuplexHandler
		AnyPermissions []string
		FreeRoute      bool // Free route require neither authentication nor authorization
	}
)

type DuplexConnection interface {
	Publish(topic string, message any) error
	GetCaller() (misc.Caller, bool)
	MustGetCaller() misc.Caller
}

type DuplexMessage interface {
	GetDto() any
	MustGetCaller() misc.Caller
	GetTopic() string
}

type DuplexModule interface {
	GetDuplexHandlers() []*DuplexHandlerDefinition
	GetBaseURL() string
	OnDuplexConnected(DuplexConnection)
	OnDuplexDisconnected(DuplexConnection)
}
