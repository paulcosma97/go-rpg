package match

import (
	"errors"
	"fmt"
	"game/src/client"
	"log"

	"github.com/google/uuid"
)

type MatchManager struct {
	Matches []*Match
}

func New() *MatchManager {
	return &MatchManager{
		Matches: []*Match{},
	}
}

func (mm *MatchManager) CreateMatch(c *client.Client) (*Match, error) {
	if found, _, err := mm.FindMatchByClient(c); err == nil {
		log.Printf(`Client %v attempted to create match while being part of match %v.`, c.Id, found.Id)
		return nil, errors.New(`You are already in a match.`)
	}

	m := &Match{
		Id:      uuid.NewString(),
		Players: make([]*Player, 2),
	}

	m.Players[0] = &Player{
		Client:    c,
		Character: nil,
	}

	mm.Matches = append(mm.Matches, m)
	return m, nil
}

func (mm *MatchManager) JoinMatch(c *client.Client, id string) (*Match, uint8, error) {
	m, _, err := mm.FindMatchByClient(c)

	if err == nil {
		log.Printf(`Client %v attempted to join match %v while being part of match %v.`, c.Id, id, m.Id)
		return nil, 0, errors.New(`You are already in a match.`)
	}

	p := &Player{
		Client:    c,
		Character: nil,
	}

	var idx uint8 = 0
	if m.Players[0] == nil {
		idx = 0
	} else if m.Players[1] == nil {
		idx = 1
	} else {
		log.Printf(`Client %v attempted to join full match %v.`, c.Id, m.Id)
		return nil, 0, errors.New(`Match is already full.`)
	}

	m.Players[idx] = p

	return m, idx, nil
}

func (mm *MatchManager) DestroyMatch(m *Match) {
	if m.Players[0] != nil {
		m.Players[0].Character = nil
		m.Players[0].Client = nil
		m.Players[0] = nil
	}

	if m.Players[1] != nil {
		m.Players[1].Character = nil
		m.Players[1].Client = nil
		m.Players[1] = nil
	}

	filteredMatches := []*Match{}
	for _, match := range mm.Matches {
		if match != m {
			filteredMatches = append(filteredMatches, match)
		}
	}

	mm.Matches = filteredMatches
}

func (mm *MatchManager) LeaveMatch(c *client.Client) error {
	m, pIdx, err := mm.FindMatchByClient(c)
	if err != nil {
		return err
	}

	m.Players[pIdx] = nil

	if m.IsEmpty() {
		mm.DestroyMatch(m)
	}

	return nil
}

func (mm *MatchManager) FindMatchByClient(c *client.Client) (*Match, uint8, error) {
	var m *Match = nil

	for _, match := range mm.Matches {
		if match.Players[0] != nil && match.Players[0].Client == c || match.Players[1] != nil && match.Players[1].Client == c {
			m = match
			break
		}
	}

	if m == nil {
		return nil, 0, errors.New(fmt.Sprintf(`Client %v is not part of any matches.`, c.Id))
	}

	var idx uint8 = 1
	if m.Players[0] != nil && m.Players[0].Client == c {
		idx = 0
	}

	return m, idx, nil
}
