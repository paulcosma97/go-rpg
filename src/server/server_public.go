package srv

import (
	"fmt"
	cmap "game/src/client/concurrent_map"
	"game/src/types"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Clients       *cmap.ConcurrentClientMap
	Broadcast     chan *types.Message
	Upgrader      *websocket.Upgrader
	cmsgToHandler func(c *types.GameConnection, m types.Message) error
}

func New() *Server {
	return &Server{
		Clients:   cmap.New(),
		Broadcast: make(chan *types.Message),
		Upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  0,
			WriteBufferSize: 0,
		},
	}
}

func (s *Server) Serve(addr, port string, cmsgToHandler func(c *types.GameConnection, m types.Message) error) {
	completeAddr := fmt.Sprintf(`%s:%s`, addr, port)

	s.cmsgToHandler = cmsgToHandler

	go s.exhaustBroadcastMessages()
	http.HandleFunc(`/ws`, s.handleConnections)

	err := http.ListenAndServe(completeAddr, nil)
	if err != nil {
		log.Panicf(`Cannot listen on %s:%s\n`, addr, port)
	}
}
