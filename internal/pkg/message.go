package pkg

import (
	"fmt"
	"net"
)

func NewMessage(conn net.Conn, userId string, message string) {
	conn.Write([]byte(fmt.Sprintf("\n%s > %s\n%s > ", userId, message, userId)))
}
