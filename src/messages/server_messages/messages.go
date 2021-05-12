package smsg

import (
	"game/src/types"
)

const (
	TWelcome     = `[Server] Welcome`
	TJoinMatch   = `[Server] Join Match`
	TMatchUpdate = `[Server] Join Match`
)

func Welcome(cId string) *Message {
	return &Message{
		Kind:    TWelcome,
		Payload: cId,
	}
}

type PlayerPayload struct {
	Id          string             `json:"id"`
	DisplayName string             `json:"displayName"`
	Characters  []*types.Character `json:"characters"`
}

type MatchPayload struct {
	Id       string         `json:"id"`
	Player1  *PlayerPayload `json:"player1"`
	Player2  *PlayerPayload `json:"player2"`
	Turn     *string        `json:"turn"`
	NextTurn *string        `json:"nextTurn"`
}

func JoinMatch(m MatchPayload) *Message {
	return &Message{
		Kind:    TJoinMatch,
		Payload: m,
	}
}

func MatchUpdate(m MatchPayload) *Message {
	return &Message{
		Kind:    TMatchUpdate,
		Payload: m,
	}
}
