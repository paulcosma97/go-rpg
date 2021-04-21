package cmap

import (
	"game/src/client"
	"sync"

	"github.com/gorilla/websocket"
)

type ConcurrentClientMap struct {
	mutex   *sync.Mutex
	clients map[*websocket.Conn]*client.Client
}

func (m *ConcurrentClientMap) Put(key *websocket.Conn, value *client.Client) {
	m.mutex.Lock()
	m.clients[key] = value
	m.mutex.Unlock()
}

func (m *ConcurrentClientMap) Get(key *websocket.Conn) *client.Client {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.clients[key]
}

// Prefer using Put and Get over this as it is expensive
func (m *ConcurrentClientMap) Safe() map[*websocket.Conn]*client.Client {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	copy := make(map[*websocket.Conn]*client.Client)

	for key, value := range m.clients {
		copy[key] = value
	}

	return copy

}

func New() *ConcurrentClientMap {
	return &ConcurrentClientMap{
		mutex:   &sync.Mutex{},
		clients: make(map[*websocket.Conn]*client.Client),
	}
}
