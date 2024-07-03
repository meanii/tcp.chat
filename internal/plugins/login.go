package plugins

import (
	"fmt"
	"net"
	"strings"
)

func init() {
	registerPlugin(plugin{
		name:          "login",
		description:   "login to the app",
		priorityIndex: 1,
		function: func(conn net.Conn, message string) error {
			id := strings.TrimSpace(strings.Join(strings.Split(message, " ")[1:], ""))
			conn.Write([]byte(fmt.Sprintf("%s welcome to the chat!\n\n%s > ", id, id)))
			return nil
		},
	})
}
