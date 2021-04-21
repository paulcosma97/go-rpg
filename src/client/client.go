package client

import (
	"game/src/msg/cmsg"
	"game/src/msg/servmsg"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ClientProfile struct {
	DisplayName string
}

type Client struct {
	Id            string
	Connection    *websocket.Conn
	ClientMessage chan *cmsg.Message
	ServerMessage chan *servmsg.Message
	Profile       *ClientProfile
	closed        bool
}

func (c *Client) Close() {
	if c.closed {
		return
	}

	c.closed = true
	conn := c.Connection
	conn.Close()
	close(c.ClientMessage)
	close(c.ServerMessage)
	c.Connection = nil
	c.ClientMessage = nil
	c.ServerMessage = nil

	log.Printf(`Client %v has been disconnected.`, c.Id)
}

func (c *Client) IsClosed() bool {
	return c.closed
}

func New(conn *websocket.Conn) *Client {
	client := &Client{
		Connection:    conn,
		ClientMessage: make(chan *cmsg.Message, 100),
		ServerMessage: make(chan *servmsg.Message, 100),
		Id:            uuid.NewString(),
		closed:        false,
		Profile:       &ClientProfile{},
	}

	go func() {
		client.ClientMessage <- &cmsg.Message{Kind: cmsg.TConnect}
	}()

	client.ResetDeadlines()
	return client
}

func (c *Client) ResetDeadlines() {
	c.Connection.SetReadDeadline(time.Now().Add(15 * time.Second))
	c.Connection.SetWriteDeadline(time.Now().Add(15 * time.Second))
}
