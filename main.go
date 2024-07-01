package main

import (
	"github.com/meanii/tcp.chat/server"
)

func main() {
	tcpServer := server.NewTCPServer(server.Server{
		Port: 3000,
		Host: "0.0.0.0",
	})
	tcpServer.Start()
}
