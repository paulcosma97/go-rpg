package service

import (
	"errors"
	"fmt"
	cmsg "game/src/messages/client_messages"
	"game/src/types"
)

func (g *GameService) mapClientMessage(c *types.GameConnection, m cmsg.Message) error {
	switch m.Kind() {
	case cmsg.TPing:
		return g.OnPing(*c)
	case cmsg.TConnect:
		return g.onConnect(*c)
	case cmsg.TSetProfile:
		return g.onSetProfile(*c, m.Payload().(cmsg.SetProfilePayload))
	}

	return errors.New(fmt.Sprintf(`Could not find suitable handler for message type "%v".`, m.Kind))
}
