package api

import (
	"accounting-bot/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewIndexAPIs registers index api
func NewIndexAPIs(s server.Server) {
	s.RegisterAPI("/", http.MethodGet, index)
	s.RegisterAPI("/test", http.MethodGet, test)
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "lionbot server")
}

func test(c *gin.Context) {
	data := c.Query("data")
	logrus.Info("test data: ", data)
	c.JSON(http.StatusOK, "test")
}
