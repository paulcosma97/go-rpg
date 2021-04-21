package main

import (
	srv "game/src/server"
	"game/src/service"
)

func main() {
	srv := srv.New()
	game := service.New(srv)

	game.Start(`localhost`, `3000`)
}
