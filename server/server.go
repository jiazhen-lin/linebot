package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jiazhen-lin/linebot/config"
	"github.com/sirupsen/logrus"
)

var (
	// ErrRegisterAPI means api register failed if method is invalid
	ErrRegisterAPI = errors.New("register error")
)

// Server interface
type Server interface {
	// RegisterAPI registers api handler
	RegisterAPI(route string, method string, handler func(c *gin.Context))
	// Run runs server
	Run()
}

// New return server implement
func New(log *logrus.Logger, config *config.Config) Server {
	router := gin.New()
	return &botServer{
		router: router,
		log:    log,
		config: config,
	}
}

type botServer struct {
	router *gin.Engine
	log    *logrus.Logger
	config *config.Config

	// TODO: middleware
	// verify bot request
}

func (s *botServer) RegisterAPI(route string, method string, handler func(c *gin.Context)) {
	// TODO: we can make api groups in the future
	s.router.Handle(method, route, handler)
}

func (s *botServer) Run() {
	srv := &http.Server{
		Addr:    ":" + s.config.ServerConfig.Port,
		Handler: s.router,
	}
	go func() {
		if err := srv.ListenAndServeTLS("ssl/bundle.crt", "ssl/private.key"); err != nil {
			//if err := srv.ListenAndServe(); err != nil {
			s.log.Error("server error: ", err)
		}
	}()
	// Graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.log.Info("ready to shutdown... ")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.log.Error("server shutdown error: ", err)
	}
	s.log.Info("server shutdown")
}
