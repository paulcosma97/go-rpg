package cmap

import (
	"game/src/types"
	"sync"

	"github.com/gorilla/websocket"
)

type ConcurrentClientMap struct {
	mutex   *sync.Mutex
	clients map[*websocket.Conn]*types.GameConnection
}

func (m *ConcurrentClientMap) Put(key *websocket.Conn, value *types.GameConnection) {
	m.mutex.Lock()
	m.clients[key] = value
	m.mutex.Unlock()
}

func (m *ConcurrentClientMap) Get(key *websocket.Conn) *types.GameConnection {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.clients[key]
}

// Prefer using Put and Get over this as it is expensive
func (m *ConcurrentClientMap) Safe() map[*websocket.Conn]*types.GameConnection {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	copy := make(map[*websocket.Conn]*types.GameConnection)

	for key, value := range m.clients {
		copy[key] = value
	}

	return copy

}

func New() *ConcurrentClientMap {
	return &ConcurrentClientMap{
		mutex:   &sync.Mutex{},
		clients: make(map[*websocket.Conn]*types.GameConnection),
	}
}
