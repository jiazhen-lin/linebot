package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

type unfollowCMD struct {
}

func (c *unfollowCMD) Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error) {
	return nil, nil
}

// NewUnfollowCommand returns command concrete instance
func NewUnfollowCommand() Interface {
	return &unfollowCMD{}
}
