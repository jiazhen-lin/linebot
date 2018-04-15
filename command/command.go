package command

// Interface is command interface
type Interface interface {
	// Command executes cmd and return result
	Command(cmd string) error
	// Register message command
	Register(cmd string, handler func() error) error
}

type command struct {
}

func (c *command) Register(cmd string, handler func() error) error {
	return nil
}
func (c *command) Command(cmd string) error {
	return nil
}

// New return command concrete instance
func New() Interface {
	return &command{}
}
