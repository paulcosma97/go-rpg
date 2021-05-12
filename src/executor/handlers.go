package service

import (
	"errors"
	"fmt"
	cmsg "game/src/messages/client_messages"
	smsg "game/src/messages/server_messages"
	"game/src/types"
	"log"
)

func send(c types.GameConnection, msg *types.Message) {
	go func() {
		c.WriteMessage(msg)
	}()
}

func indexOf(arr []interface{}, item interface{}) (int, error) {
	for idx, it := range arr {
		if it == item {
			return idx, nil
		}
	}

	return 0, errors.New(fmt.Sprintf(`Could not find element %v in array %v.`, item, arr))
}

func (g *GameService) OnPing(c types.GameConnection) error {
	return nil
}

func (g *GameService) onConnect(c types.GameConnection) error {
	log.Printf(`A new client has just connected! { Id '%v', Addr '%v' }`, c.Id, c.WebSocket().RemoteAddr())

	msg := *smsg.Welcome(c.Id())
	send(c, msg)

	return nil
}

func (g *GameService) onSetProfile(c types.GameConnection, p cmsg.SetProfilePayload) error {
	c.Profile.DisplayName = p.DisplayName
	return nil
}

func (g *GameService) onCreateMatch(c types.GameConnection) error {
	m, err := g.matchManager.CreateMatch(c)
	if err != nil {
		return err
	}

	payload := smsg.MatchPayload{
		Id: m.Id,
		Player1: &smsg.PlayerPayload{
			Id:          c.Id,
			Characters:  m.Players[0].Characters,
			DisplayName: c.Profile.DisplayName,
		},
		Player2:  nil,
		Turn:     nil,
		NextTurn: nil,
	}

	send(c, smsg.JoinMatch(payload))

	return nil
}

func (g *GameService) onJoinMatch(c types.GameConnection, mId string) error {
	m, _, err := g.matchManager.JoinMatch(c, mId)
	if err != nil {
		return err
	}

	m.InProgress = true

	m.Turn, err = m.RandomPlayer()
	if err != nil {
		log.Printf(`Could not assign first turn to a random player in match %v.`, m.Id)
		return err
	}

	m.NextTurn, err = m.EnemyOf(m.Turn)
	if err != nil {
		log.Printf(`Could not assign second turn to a random player in match %v.`, m.Id)
	}

	send(c, smsg.JoinMatch(m.ToOutbound()))

	for _, player := range m.Players {
		send(player.Client, smsg.MatchUpdate(m.ToOutbound()))
	}

	return nil
}

func (g *GameService) onChooseCharacters(c types.GameConnection, cIds []uint8) error {

	return nil
}
