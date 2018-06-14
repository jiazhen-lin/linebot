package command

import (
	"github.com/jiazhen-lin/linebot/config"
	"github.com/sirupsen/logrus"
)

// Interface is command interface
type Interface interface {
	// Command executes cmd and return result
	Command(cmd string) error
	// Register message command
	Register(cmd string, handler func() error) error
}

type command struct {
	log    *logrus.Logger
	config *config.Config
}

func (c *command) Register(cmd string, handler func() error) error {
	return nil
}
func (c *command) Command(cmd string) error {
	return nil
}

// New return command concrete instance
func New(log *logrus.Logger, config *config.Config) Interface {
	return &command{
		log:    log,
		config: config,
	}
}
