package utils

import (
	"strings"
)

type commands struct {
	Command string
	Message string
	Args    []string
}

// CommandParser parse the command
func CommandParser(message string) *commands {
	messageString := strings.TrimSpace(strings.Join(strings.Split(message, " ")[1:], " "))
	MessageArgs := strings.Split(message[1:], " ")
	if !strings.HasPrefix(message, "\\") {
		return &commands{Command: "", Message: messageString, Args: MessageArgs}
	}
	return &commands{
		Command: strings.TrimSpace(strings.Split(strings.Replace(string(message), "\\", "", 1), " ")[0]),
		Message: messageString,
		Args:    MessageArgs,
	}
}
