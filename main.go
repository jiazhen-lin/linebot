package main

import (
	"github.com/jiazhen-lin/linebot/api"
	"github.com/jiazhen-lin/linebot/command"
	"github.com/jiazhen-lin/linebot/config"
	"github.com/jiazhen-lin/linebot/server"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	config, err := config.New()
	if err != nil {
		log.Errorf("%s", err)
	}
	srv := server.New(log, config)
	cmd := command.New(log, config)

	api.NewBotAPIs(srv, cmd)
	api.NewIndexAPIs(srv, log)

	srv.Run()
}
