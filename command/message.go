package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type messageCMD struct {
}

func (c *messageCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	m := linebot.NewTextMessage("test message")
	return []linebot.Message{m}, nil
}

func (c *messageCMD) Register(cmd string, handler func() error) error {
	return nil
}

// NewMessageCommand returns command concrete instance
func NewMessageCommand() Interface {
	return &messageCMD{}
}
