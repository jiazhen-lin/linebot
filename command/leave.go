package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type leaveCMD struct {
}

func (c *leaveCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	return nil, nil
}

// NewLeaveCommand returns command concrete instance
func NewLeaveCommand() Interface {
	return &leaveCMD{}
}
