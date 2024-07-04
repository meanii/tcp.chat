package pkg

import (
	"fmt"
	"net"
)

func NewMessage(conn net.Conn, userId string, message string) {
	if userId == "" {
		userId = "login #"
	}
	conn.Write([]byte(fmt.Sprintf("\n%s > %s\n%s > ", userId, message, userId)))
}

func NewGroupMessage(conn net.Conn, groupId string, userId string, message string) {
	conn.Write([]byte(fmt.Sprintf("\n%s ( %s ) > %s\n%s ( %s ) >", groupId, userId, message, groupId, userId)))
}
