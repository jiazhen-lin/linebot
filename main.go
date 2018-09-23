package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"

	"github.com/jiazhen-lin/linebot/api"
	"github.com/jiazhen-lin/linebot/command"
	"github.com/jiazhen-lin/linebot/config"
	"github.com/jiazhen-lin/linebot/server"
)

func main() {
	// initialize log, config, bot, server, db
	log := logrus.New()
	config, err := config.New()
	if err != nil {
		log.Panic(err)
	}
	bot, err := linebot.New(config.LineConfig.Secret, config.LineConfig.Token)
	if err != nil {
		log.Panic(err)
	}
	srv := server.New(log, config)
	dbConnectionString := fmt.Sprintf(
		"%v:%v@tcp(%v:3306)/%v",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Database)
	db, err := sqlx.Connect("mysql", dbConnectionString)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	// create database tables
	tableSchema, err := ioutil.ReadFile("./sql_script/create_tables.sql")
	if err != nil {
		log.Panic(err)
	}
	script := strings.Split(string(tableSchema), ";")
	for _, s := range script {
		// skip last string
		if s == "" {
			continue
		}
		db.MustExec(s)
	}

	// Linebot command handler
	follow := command.NewFollowCommand()
	unFollow := command.NewUnfollowCommand()
	join := command.NewJoinCommand()
	leave := command.NewLeaveCommand()
	postback := command.NewPostbackCommand()
	message := command.NewMessageCommand(db)

	api.NewBotAPIs(srv, bot, log, follow, unFollow, join, leave, postback, message)
	api.NewIndexAPIs(srv, log)

	srv.Run()
}
