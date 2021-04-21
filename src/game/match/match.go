package match

import (
	"errors"
	"game/src/client"
	char "game/src/game/character"
	"game/src/msg/servmsg"
	"log"
	"math"
	"math/rand"
)

type Player struct {
	Client    *client.Client
	Character *char.Character
}

func (p *Player) ToOutbound() servmsg.PlayerPayload {
	return servmsg.PlayerPayload{
		Id:          p.Client.Id,
		Character:   p.Character,
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
	if m.IsEmpty() || !m.IsFull() {
		log.Printf(`Cannot get enemy of %v as match %v is not yet full.`, p.Client.Id, m.Id)
		return nil, errors.New(`Not enough players.`)
	}

	pIdx := int(math.Round(rand.Float64()))
	return m.Players[pIdx], nil
}

func (m *Match) ToOutbound() servmsg.MatchPayload {
	p1 := m.Players[0].ToOutbound()
	p2 := m.Players[1].ToOutbound()


	return servmsg.MatchPayload{
		Id:       m.Id,
		Player1:  &p1,
		Player2:  &p2,
		Turn:     &m.Turn.Client.Id,
		NextTurn: &m.NextTurn.Client.Id,
	}
}
