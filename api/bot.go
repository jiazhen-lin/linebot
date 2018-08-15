package api

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/jiazhen-lin/linebot/command"
	"github.com/jiazhen-lin/linebot/server"
)

// NewBotAPIs registers bot api
func NewBotAPIs(s server.Server,
	bot *linebot.Client,
	db *sqlx.DB,
	log *logrus.Logger,
	follow command.Interface,
	unFollow command.Interface,
	join command.Interface,
	leave command.Interface,
	postback command.Interface,
	message command.Interface,
) {
	b := &botAPI{
		bot:      bot,
		db:       db,
		log:      log,
		follow:   follow,
		unfollow: unFollow,
		join:     join,
		leave:    leave,
		postback: postback,
		message:  message,
	}
	s.RegisterAPI("/linebot", http.MethodPost, b.handler)
}

type botAPI struct {
	// Linebot object helps parse request and reply/push response to clients
	bot *linebot.Client
	db  *sqlx.DB
	log *logrus.Logger
	// Line request action handlers
	follow   command.Interface
	unfollow command.Interface
	join     command.Interface
	leave    command.Interface
	postback command.Interface
	message  command.Interface
}

func (api *botAPI) handler(c *gin.Context) {
	bodyByte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		api.log.Error("read body error: ", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("parse body error: %s", err))
		return
	}
	bodyString := string(bodyByte)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
	api.log.Info("read body: ", bodyString)

	// Dispatch request by event.Type
	events, err := api.bot.ParseRequest(c.Request)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("parse request error: %s", err))
	}
	api.log.Info("linebot.Request events: ", events)
	for _, event := range events {
		ctx := context.Background()
		messages, err := api.handle(ctx, event)
		if err != nil {
			c.String(http.StatusBadRequest, "")
		}
		_, err = api.bot.ReplyMessage(event.ReplyToken, messages...).Do()
		if err != nil {
			c.String(http.StatusServiceUnavailable, fmt.Sprintf("reply error: %s", err))
		}
	}
	c.String(http.StatusOK, "")
}

func (api *botAPI) handle(ctx context.Context, event *linebot.Event) (messages []linebot.Message, err error) {
	switch event.Type {
	case linebot.EventTypeFollow:
		messages, err = api.follow.Command(ctx, event)
	case linebot.EventTypeUnfollow:
		messages, err = api.unfollow.Command(ctx, event)
	case linebot.EventTypeJoin:
		messages, err = api.join.Command(ctx, event)
	case linebot.EventTypeLeave:
		messages, err = api.leave.Command(ctx, event)
	case linebot.EventTypePostback:
		messages, err = api.postback.Command(ctx, event)
	case linebot.EventTypeMessage:
		messages, err = api.message.Command(ctx, event)
	default:
		// Error
		return nil, fmt.Errorf(fmt.Sprintf("event type not support: %s", event))
	}
	if err != nil {
		return nil, err
	}
	return messages, nil
}
