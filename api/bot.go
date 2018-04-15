package api

import (
	"accounting-bot/command"
	"accounting-bot/server"

	"net/http"

	"github.com/gin-gonic/gin"
)

// NewBotAPIs registers bot api
func NewBotAPIs(s server.Server, cmd command.Interface) {
	b := &botAPI{
		cmd: cmd,
	}
	s.RegisterAPI("/linebot", http.MethodPost, b.Handler)
}

type botAPI struct {
	cmd command.Interface
}

func (api *botAPI) Handler(c *gin.Context) {
}
