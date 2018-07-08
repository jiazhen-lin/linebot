package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type followCMD struct {
}

func (c *followCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	return nil, nil
}

// NewFollowCommand returns command concrete instance
func NewFollowCommand() Interface {
	return &followCMD{}
}
