package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			keyboardInput := scanner.Text()

			if err := execInput(keyboardInput); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func execInput(input string) error {
	// trim trailing space and new line
	input = strings.TrimRight(input, "\n ")
	// separate the command and arguments
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) != 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// put the command and argument
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	// execute command and return error
	return cmd.Run()
}
