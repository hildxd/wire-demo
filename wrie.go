//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"go-wire/pkg/event"
	"go-wire/pkg/greeting"
	"go-wire/pkg/message"

	"github.com/google/wire"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeEvent(phrase string) (event.Event, error) {
	wire.Build(event.NewEvent, greeting.NewGreeter, message.NewMessage)
	return event.Event{}, nil
}
