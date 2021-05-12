package client

import (
	cmsg "game/src/messages/client_messages"
	"game/src/types"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id            string
	connection    *websocket.Conn
	clientMessage chan *types.Message
	serverMessage chan *types.Message
	closed        bool
}

func (c *Client) Close() {
	if c.closed {
		return
	}

	c.closed = true
	conn := c.connection
	conn.Close()
	close(c.clientMessage)
	close(c.serverMessage)
	c.connection = nil
	c.clientMessage = nil
	c.serverMessage = nil

	log.Printf(`Client %v has been disconnected.`, c.id)
}

func (c *Client) Id() string {
	return c.id
}

func (c *Client) WebSocket() *websocket.Conn {
	return c.connection
}

func (c *Client) WriteMessage(m *types.Message) {
	c.serverMessage <- m
}

func (c *Client) ReadChannel() chan *types.Message {
	return c.clientMessage
}

func (c *Client) WriteChannel() chan *types.Message {
	return c.serverMessage
}

func (c *Client) IsClosed() bool {
	return c.closed
}

func New(conn *websocket.Conn) *types.GameConnection {
	client := &Client{
		connection:    conn,
		clientMessage: make(chan *types.Message, 100),
		serverMessage: make(chan *types.Message, 100),
		id:            uuid.NewString(),
		closed:        false,
	}

	go func() {
		msg := (&cmsg.MessageIn{Kind: cmsg.TConnect}).Inbound()
		client.clientMessage <- &msg
	}()

	client.ResetDeadlines()
	return client
}

func (c *Client) ResetDeadlines() {
	c.connection.SetReadDeadline(time.Now().Add(15 * time.Second))
	c.connection.SetWriteDeadline(time.Now().Add(15 * time.Second))
}
