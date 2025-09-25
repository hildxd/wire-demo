package message

type Message string

type Greeter struct {
	// ... TBD
}

type Event struct {
	// ... TBD
}

func NewMessage() Message {
	return Message("Hi there!")
}
