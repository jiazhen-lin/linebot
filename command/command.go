package command

// Interface is command interface
type Interface interface {
	// Command executes cmd and return result
	Command(cmd string) error
}
