package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type joinCMD struct {
}

func (c *joinCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	return nil, nil
}

// NewJoinCommand returns command concrete instance
func NewJoinCommand() Interface {
	return &joinCMD{}
}
