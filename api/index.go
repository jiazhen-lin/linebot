package api

import (
	"net/http"

	"github.com/jiazhen-lin/linebot/server"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewIndexAPIs registers index api
func NewIndexAPIs(s server.Server, log *logrus.Logger) {
	h := index{log}
	s.RegisterAPI("/", http.MethodGet, h.index)
	s.RegisterAPI("/test", http.MethodGet, h.test)
}

type index struct {
	log *logrus.Logger
}

func (h *index) index(c *gin.Context) {
	c.String(http.StatusOK, "linebot server")
}

func (h *index) test(c *gin.Context) {
	data := c.Query("data")
	h.log.Info("test data: ", data)
	c.JSON(http.StatusOK, "test")
}
