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
	if !strings.HasPrefix(message, "\\") {
		return &commands{Command: "", Message: message[1:], Args: strings.Split(message[1:], " ")}
	}
	return &commands{
		Command: strings.TrimSpace(strings.Split(strings.Replace(string(message), "\\", "", 1), " ")[0]),
		Message: message[1:],
		Args:    strings.Split(message[1:], " "),
	}
}
