package service

import (
	srv "game/src/server"
	"log"
)

type GameService struct {
	server *srv.Server
}

func New(server *srv.Server) *GameService {
	g := &GameService{
		server: server,
	}

	return g
}

func (g *GameService) Start(addr, port string) {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf(`Game server started listening on %v:%v and serving /ws.`, addr, port)
	g.server.Serve(addr, port, g.mapClientMessage)

}
