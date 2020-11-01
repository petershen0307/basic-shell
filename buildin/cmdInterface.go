package buildin

import "errors"

// ErrorBreak is the error for break command chain
var ErrorBreak error = errors.New("break")

// ICmd is the command interface
type ICmd interface {
	Handle(args []string) error
}
