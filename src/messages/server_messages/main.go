package smsg

import "game/src/types"

type MessageOut struct {
	Kind    string      `json:"kind"`
	Payload interface{} `json:"payload"`
}

type Message struct {
	kind    string
	payload interface{}
}

func (m Message) Kind() string {
	return m.kind
}

func (m Message) Payload() interface{} {
	return m.payload
}

func (m Message) Outbound() types.Message {
	msg := Message{
		m.Kind(),
		m.Payload(),
	}
	return msg
}
