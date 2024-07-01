package cmd

import (
	"fmt"
	"net"
	"strings"
)

func Login(conn net.Conn, message string) {
	id := strings.TrimSpace(strings.Join(strings.Split(message, " ")[1:], ""))
	user := AllUsers.Create(conn, id)
	user.Conn.Write([]byte(fmt.Sprintf("%s welcome to the chat!\n\n%s > ", user.Id, user.Id)))
}
