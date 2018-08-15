package main

import (
	"fmt"

	"github.com/jiazhen-lin/linebot/api"
	"github.com/jiazhen-lin/linebot/command"
	"github.com/jiazhen-lin/linebot/config"
	"github.com/jiazhen-lin/linebot/server"
	"github.com/jmoiron/sqlx"
	"github.com/line/line-bot-sdk-go/linebot"
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
	dbConnectionString := fmt.Sprintf(
		"%v:%v@(%v:3306)/%v",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Database)
	db, err := sqlx.Connect("mysql", dbConnectionString)
	if err != nil {
		log.Error(err)
	}
	err = db.Ping()
	if err != nil {
		log.Error(err)
	}

	// Linebot command handler
	follow := command.NewFollowCommand()
	unFollow := command.NewUnfollowCommand()
	join := command.NewJoinCommand()
	leave := command.NewLeaveCommand()
	postback := command.NewPostbackCommand()
	message := command.NewMessageCommand()

	api.NewBotAPIs(srv, bot, db, log, follow, unFollow, join, leave, postback, message)
	api.NewIndexAPIs(srv, log)

	srv.Run()
}
