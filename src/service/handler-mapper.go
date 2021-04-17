package service

import (
	"errors"
	"fmt"
	"game/src/msg/cmsg"
	"game/src/msg/servmsg"
	srv "game/src/server"
)

func (g *GameService) mapClientMessage(c *srv.Client, m *cmsg.Message) (*servmsg.Message, error) {
	switch m.Kind {
	case cmsg.TPing:
		return g.OnPing(c)
	case cmsg.TConnect:
		return g.onConnect(c)
	}

	return nil, errors.New(fmt.Sprintf(`Could not find suitable handler for message type "%v"`, m.Kind))
}
