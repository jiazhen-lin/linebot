package main

import (
	"github.com/jiazhen-lin/linebot/api"
	"github.com/jiazhen-lin/linebot/command"
	"github.com/jiazhen-lin/linebot/server"
)

func main() {
	srv := server.New()
	cmd := command.New()

	api.NewBotAPIs(srv, cmd)
	api.NewIndexAPIs(srv)

	srv.Run()
}
