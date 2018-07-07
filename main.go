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
		log.Error(err)
	}
	bot, err := linebot.New(config.LineConfig.Secret, config.LineConfig.Token)
	if err != nil {
		log.Error(err)
	}
	srv := server.New(log, config)
	cmd := command.New(log, config)

	// Linebot command handler
	follow := command.NewFollowCommand()
	unFollow := command.NewUnfollowCommand()
	join := command.NewJoinCommand()
	leave := command.NewLeaveCommand()
	postback := command.NewPostbackCommand()
	message := command.NewMessageCommand()

	api.NewBotAPIs(srv, bot, follow, unFollow, join, leave, postback, message)
	api.NewIndexAPIs(srv, log)

	srv.Run()
}
