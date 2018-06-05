package main

import (
	"github.com/jiazhen-lin/linebot/api"
	"github.com/jiazhen-lin/linebot/command"
	"github.com/jiazhen-lin/linebot/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	srv := server.New(log)
	cmd := command.New(log)

	api.NewBotAPIs(srv, cmd)
	api.NewIndexAPIs(srv)

	srv.Run()
}
