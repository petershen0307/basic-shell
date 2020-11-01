package buildin

import (
	"errors"
	"os"
)

const cdStr = "cd"

// CD is "change directory"
type CD struct {
}

// Handle the cd command
func (cmd *CD) Handle(args []string) error {
	if cdStr != args[0] {
		return nil
	}
	if len(args) != 2 {
		return errors.New("path required")
	}
	err := os.Chdir(args[1])
	if nil == err {
		err = ErrorBreak
	}
	return err
}
