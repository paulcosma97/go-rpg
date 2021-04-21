package srv

import (
	"game/src/client"
	"game/src/msg/cmsg"
	"log"
	"net/http"
)

func (s *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicf(`Upgrade request failed. %s %s %s`, r.Method, r.Host, err)
		return
	}

	client := client.New(ws)
	s.Clients.Put(ws, client)
	go s.exhaustClientMessages(client)
	go s.exhaustServerMessagesForClient(client)

	for !client.IsClosed() {

		var m cmsg.Message
		err := ws.ReadJSON(&m)

		if client.IsClosed() {
			break
		}

		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			break
		}

		client.ClientMessage <- &m
		client.ResetDeadlines()
	}
}

func (s *Server) exhaustClientMessages(c *client.Client) {
	for m := range c.ClientMessage {
		err := s.cmsgToHandler(c, *m)

		if err != nil {
			log.Printf("Could not handle message %v for client %v.\n\tError: %v", m, c.Id, err)
			c.Close()
			return
		}
	}
}

func (s *Server) exhaustServerMessagesForClient(c *client.Client) {
	for m := range c.ServerMessage {
		err := c.Connection.WriteJSON(m)
		if err != nil {
			log.Printf(`Could not respond to client %v. Closing connection.`, c.Connection.RemoteAddr())
			return
		}
	}
}

func (s *Server) exhaustBroadcastMessages() {
	for message := range s.Broadcast {
		for _, client := range s.Clients.Safe() {
			client.ServerMessage <- message
		}
	}
}
