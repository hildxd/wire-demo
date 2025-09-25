package greeting

import "go-wire/pkg/message"

func NewGreeter(m message.Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message message.Message // <- adding a Message field
}

func (g Greeter) Greet() message.Message {
	return g.Message
}
