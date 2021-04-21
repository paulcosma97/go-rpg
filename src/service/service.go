package service

import (
	"game/src/game/match"
	srv "game/src/server"
	"log"
)

type GameService struct {
	server       *srv.Server
	matchManager *match.MatchManager
}

func New(server *srv.Server) *GameService {
	g := &GameService{
		server:       server,
		matchManager: match.New(),
	}

	return g
}

func (g *GameService) Start(addr, port string) {
	log.Default().SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf(`Game server started listening on %v:%v and serving /ws.`, addr, port)
	g.server.Serve(addr, port, g.mapClientMessage)

}
