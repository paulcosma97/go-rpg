package srv

import (
	"game/src/msg/cmsg"
	"game/src/msg/servmsg"
	"log"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id            string
	server        *Server
	Connection    *Connection
	ClientMessage chan *cmsg.Message
	ServerMessage chan *servmsg.Message
	closed        bool
}

func (c *Client) Close() {
	if c.closed {
		return
	}

	c.closed = true
	conn := c.Connection
	conn.Close()
	delete(c.server.Clients, conn)
	close(c.ClientMessage)
	close(c.ServerMessage)
	c.server = nil
	c.Connection = nil
	c.ClientMessage = nil
	c.ServerMessage = nil

	log.Printf(`Client %v has been disconnected.`, c.Id)
}

func NewClient(s *Server, conn *Connection) *Client {
	client := &Client{
		Connection:    conn,
		server:        s,
		ClientMessage: make(chan *cmsg.Message),
		ServerMessage: make(chan *servmsg.Message),
		Id:            uuid.NewString(),
		closed:        false,
	}

	go func() {
		client.ClientMessage <- cmsg.Connect()
	}()

	client.ResetDeadlines()
	return client
}

func (c *Client) ResetDeadlines() {
	c.Connection.SetReadDeadline(time.Now().Add(15 * time.Second))
	c.Connection.SetWriteDeadline(time.Now().Add(15 * time.Second))
}
