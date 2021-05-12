package types

type Message interface {
	Kind() string
	Payload() interface{}
}
