package buildin

import (
	"os"
)

const exitStr = "exit"

// Exit is "change directory"
type Exit struct {
}

// Handle the cd command
func (cmd *Exit) Handle(args []string) error {
	if exitStr != args[0] {
		return nil
	}
	os.Exit(0)
	return ErrorBreak
}
