package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type postbackCMD struct {
}

func (c *postbackCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	return nil, nil
}

// NewPostbackCommand returns command concrete instance
func NewPostbackCommand() Interface {
	return &postbackCMD{}
}
