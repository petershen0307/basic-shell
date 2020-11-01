package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/petershen0307/basic-shell/buildin"
)

var cmdList []buildin.ICmd = []buildin.ICmd{
	new(buildin.CD),
	new(buildin.Exit),
	new(execInput),
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			keyboardInput := scanner.Text()

			if err := run(keyboardInput); err != nil && buildin.ErrorBreak != err {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func run(input string) error {
	// trim trailing space and new line
	input = strings.TrimRight(input, "\n ")
	// separate the command and arguments
	args := strings.Split(input, " ")
	var err error
	for _, cmd := range cmdList {
		err = cmd.Handle(args)
		if nil != err {
			break
		}
	}
	return err
}
