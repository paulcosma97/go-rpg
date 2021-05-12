package match

import (
	"errors"
	"game/src/client"
	smsg "game/src/messages/server_messages"
	"game/src/types"
	"log"
	"math"
	"math/rand"
)

type Player struct {
	Client     *client.Client
	Characters []*types.Character
}

func (p *Player) ToOutbound() smsg.PlayerPayload {
	return smsg.PlayerPayload{
		Id:          p.Client.Id,
		Characters:  p.Characters,
		DisplayName: p.Client.Profile.DisplayName,
	}
}

type Match struct {
	Id         string
	Players    []*Player
	InProgress bool
	Turn       *Player
	NextTurn   *Player
}

func (m *Match) IsEmpty() bool {
	return m.Players[0] == nil && m.Players[1] == nil
}

func (m *Match) IsFull() bool {
	return m.Players[0] != nil && m.Players[1] != nil
}

func (m *Match) EnemyOf(p *Player) (*Player, error) {
	if !m.IsFull() {
		log.Printf(`Cannot get enemy of %v as match %v is not yet full.`, p.Client.Id, m.Id)
		return nil, errors.New(`Not enough players.`)
	}

	if m.Players[0] == p {
		return m.Players[1], nil
	}

	return m.Players[0], nil
}

func (m *Match) RandomPlayer() (*Player, error) {
	if !m.IsFull() {
		log.Printf(`Cannot get random player as match %v is not yet full.`, m.Id)
		return nil, errors.New(`Not enough players.`)
	}

	pIdx := int(math.Round(rand.Float64()))
	return m.Players[pIdx], nil
}

func (m *Match) ToOutbound() smsg.MatchPayload {
	p1 := m.Players[0].ToOutbound()
	p2 := m.Players[1].ToOutbound()

	return smsg.MatchPayload{
		Id:       m.Id,
		Player1:  &p1,
		Player2:  &p2,
		Turn:     &m.Turn.Client.Id,
		NextTurn: &m.NextTurn.Client.Id,
	}
}
