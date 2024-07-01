package utils

import (
	"strings"
)

// CommandParser parse the command
func CommandParser(message string) string {
	if !strings.HasPrefix(message, "\\") {
		return ""
	}
	return strings.TrimSpace(strings.Split(strings.Replace(string(message), "\\", "", 1), " ")[0])
}
