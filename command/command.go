package command

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Interface is command interface
type Interface interface {
	// Command executes cmd and return result
	Command(ctx context.Context, event *linebot.Event) ([]linebot.Message, error)
}
