package service

import (
	"errors"
	"fmt"
	"game/src/client"
	"game/src/msg/cmsg"
	"game/src/msg/servmsg"
	"log"
)

func send(c *client.Client, msg *servmsg.Message) {
	go func() {
		c.ServerMessage <- msg
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

func (g *GameService) OnPing(c *client.Client) error {
	return nil
}

func (g *GameService) onConnect(c *client.Client) error {
	log.Printf(`A new client has just connected! { Id '%v', Addr '%v' }`, c.Id, c.Connection.RemoteAddr())

	send(c, servmsg.Welcome(c.Id))

	return nil
}

func (g *GameService) onSetProfile(c *client.Client, p cmsg.SetProfilePayload) error {
	c.Profile.DisplayName = p.DisplayName
	return nil
}

func (g *GameService) onCreateMatch(c *client.Client) error {
	m, err := g.matchManager.CreateMatch(c)
	if err != nil {
		return err
	}

	payload := servmsg.MatchPayload{
		Id: m.Id,
		Player1: &servmsg.PlayerPayload{
			Id:          c.Id,
			Character:   m.Players[0].Character,
			DisplayName: c.Profile.DisplayName,
		},
		Player2:  nil,
		Turn:     nil,
		NextTurn: nil,
	}

	send(c, servmsg.JoinMatch(payload))

	return nil
}

func (g *GameService) onJoinMatch(c *client.Client, mId string) error {
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

	send(c, servmsg.JoinMatch(m.ToOutbound()))

	for _, player := range m.Players {
		send(player.Client, servmsg.MatchUpdate(m.ToOutbound()))
	}

	return nil
}
