package cmd

import (
	"bufio"
	"fmt"
	"net"

	"github.com/meanii/tcp.chat/utils"
)

var AllUsers = UsersInit()

func HandleRawTCPConnections(conn net.Conn) {
	WelcomeHandler(conn)
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}

		command := utils.CommandParser(message)
		switch command {
		case "":
			Message(conn, message)
		case "login":
			Login(conn, message)
		case "users":
			AvailableUsers(conn)
		case "broadcast":
			boradcast(conn, message)
		default:
			conn.Write([]byte(fmt.Sprintf("\n'%s' command not found!\n", command)))
		}
	}
}
