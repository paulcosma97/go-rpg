package types

import (
	"github.com/gorilla/websocket"
)

type GameConnection interface {
	Id() string
	WebSocket() *websocket.Conn
	WriteMessage(*Message)
	ReadChannel() chan *Message
	WriteChannel() chan *Message
	IsClosed() bool
	ResetDeadlines()
	Close()
}
