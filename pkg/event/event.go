package event

import (
	"fmt"
	"go-wire/pkg/greeting"
)

func NewEvent(g greeting.Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter greeting.Greeter // <- adding a Greeter field
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
