package main

import (
	"accounting-bot/api"
	"accounting-bot/command"
	"accounting-bot/server"
)

func main() {
	srv := server.New()
	cmd := command.New()

	api.NewBotAPIs(srv, cmd)
	api.NewIndexAPIs(srv)

	srv.Run()
}
