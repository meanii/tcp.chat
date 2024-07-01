package cmd

import (
	"fmt"
	"net"
	"strings"
)

func AvailableUsers(conn net.Conn) {
	users := AllUsers.AvailableUserNames()
	user := AllUsers.GetUserByConn(conn)
	conn.Write([]byte(fmt.Sprintf("List of public users you can chat with:\n%s\n\n%s > ", strings.Join(users, "\n"), user.Id)))
}
