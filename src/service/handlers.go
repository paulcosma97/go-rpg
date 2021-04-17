package service

import (
	"game/src/msg/servmsg"
	srv "game/src/server"
	"log"
)

func (g *GameService) OnPing(c *srv.Client) (*servmsg.Message, error) {
	return nil, nil
}

func (g *GameService) onConnect(c *srv.Client) (*servmsg.Message, error) {
	log.Printf(`A new client has just connected! { Id '%v', Addr '%v' }`, c.Id, c.Connection().RemoteAddr())
	return nil, nil
}
