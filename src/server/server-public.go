package srv

import (
	"fmt"
	"game/src/client"
	"game/src/client/cmap"
	"game/src/msg/cmsg"
	"game/src/msg/servmsg"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Clients       *cmap.ConcurrentClientMap
	Broadcast     chan *servmsg.Message
	Upgrader      *websocket.Upgrader
	cmsgToHandler func(c *client.Client, m cmsg.Message) error
}

func New() *Server {
	return &Server{
		Clients:   cmap.New(),
		Broadcast: make(chan *servmsg.Message),
		Upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  0,
			WriteBufferSize: 0,
		},
	}
}

func (s *Server) Serve(addr, port string, cmsgToHandler func(c *client.Client, m cmsg.Message) error) {
	completeAddr := fmt.Sprintf(`%s:%s`, addr, port)

	s.cmsgToHandler = cmsgToHandler

	go s.exhaustBroadcastMessages()
	http.HandleFunc(`/ws`, s.handleConnections)

	err := http.ListenAndServe(completeAddr, nil)
	if err != nil {
		log.Panicf(`Cannot listen on %s:%s\n`, addr, port)
	}
}
