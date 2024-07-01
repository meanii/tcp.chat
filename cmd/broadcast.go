package cmd

import "net"

func boradcast(conn net.Conn, message string) {
	AllUsers.BroadcastMessage(conn, message)
}
