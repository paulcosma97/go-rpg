package srv

import (
	"fmt"
	"game/src/msg/cmsg"
	"game/src/msg/servmsg"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
)

type Connection = websocket.Conn

type Server struct {
	Clients       map[*Connection]*Client
	Broadcast     chan *servmsg.Message
	Upgrader      *websocket.Upgrader
	cmsgToHandler func(c *Client, m *cmsg.Message) (*servmsg.Message, error)
}

func New() *Server {
	return &Server{
		Clients:   make(map[*Connection]*Client),
		Broadcast: make(chan *servmsg.Message, 100),
		Upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicf(`Upgrade request. %s %s %s\n`, r.Method, r.Host, err)
		return
	}

	client := NewClient(s, ws)
	s.Clients[ws] = client
	go s.exhaustClientMessages(client)
	go s.exhaustServerMessagesForClient(client)

	for !client.closed {

		var m cmsg.Message
		err := ws.ReadJSON(&m)

		if client.closed {
			return
		}

		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			return
		}

		client.ClientMessage <- &m
		client.ResetDeadlines()
	}
}

func (s *Server) exhaustClientMessages(c *Client) {
	for m := range c.ClientMessage {
		res, err := s.cmsgToHandler(c, m)

		if err != nil {
			log.Printf("Client %v rejected response.\n\tError: %v", c.Id, err)
			c.Close()
			return
		}

		if res != nil {
			c.ServerMessage <- res
		}
	}
}

func (s *Server) exhaustServerMessagesForClient(c *Client) {
	for m := range c.ServerMessage {
		err := c.Connection().WriteJSON(m)
		if err != nil {
			log.Printf(`Could not respond to client %v. Closing connection.`, c.Connection().RemoteAddr())
			return
		}
	}
}

func (s *Server) exhaustBroadcastMessages() {
	for message := range s.Broadcast {
		for _, client := range s.Clients {
			client.ServerMessage <- message
		}
	}
}

func (s *Server) printStatsPeriodaically() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		log.Printf("Alloc = %.2f TotalAlloc = %.2f Sys = %.2f NumGC = %v", float32(m.Alloc)/1024.0/1024.0, float32(m.TotalAlloc)/1024.0/1024.0, float32(m.Sys)/1024.0/1024.0, m.NumGC)

		log.Printf(`A total of %v clients are connected.`, len(s.Clients))
		time.Sleep(20 * time.Second)
	}
}

func (s *Server) Serve(addr, port string, cmsgToHandler func(c *Client, m *cmsg.Message) (*servmsg.Message, error)) {
	completeAddr := fmt.Sprintf(`%s:%s`, addr, port)

	s.cmsgToHandler = cmsgToHandler

	go s.exhaustBroadcastMessages()
	go s.printStatsPeriodaically()
	http.HandleFunc(`/ws`, s.handleConnections)

	err := http.ListenAndServe(completeAddr, nil)
	if err != nil {
		log.Panicf(`Cannot listen on %s:%s\n`, addr, port)
	}
}
