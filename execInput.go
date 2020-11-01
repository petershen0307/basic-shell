package main

import (
	"os"
	"os/exec"
)

type execInput struct {
}

func (*execInput) Handle(args []string) error {
	// put the command and argument
	execCmd := exec.Command(args[0], args[1:]...)
	execCmd.Stderr = os.Stderr
	execCmd.Stdout = os.Stdout
	// execute command and return error
	return execCmd.Run()
}
