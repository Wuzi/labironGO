package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var oldState *terminal.State

func init() {
	var err error
	oldState, err = terminal.MakeRaw(0)
	if err != nil {
		log.Fatalf("Unable to activate raw mode terminal: %v\n", err)
	}
}

func exit() {
	terminal.Restore(0, oldState)
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}
