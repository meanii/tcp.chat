package server

import (
	"fmt"
	"log"
	"net"

	"github.com/meanii/tcp.chat/internal/plugins"
)

var Message = make(chan net.Conn)

type Server struct {
	Port        int
	Host        string
	Name        string
	Connections Connections
}

type Connections struct {
	net.Conn
}

func NewTCPServer(config Server) *Server {
	return &config
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		log.Fatalf("tcp.chat::Server::start::Error: failed to start the TCP server!\nDetails: \n%s", err)
	}

	defer listener.Close()

	log.Printf("tcp.chat::Server::start::Info: server listening on %s:%d", s.Host, s.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("tcp.chat::Server::start::Error: couldn't read the message!\nDetails: %s", err)
			continue
		}

		s.Connections = Connections{conn}
		client := s.Connections
		go client.handlerRawTCPRequest()
	}
}

func (c *Connections) handlerRawTCPRequest() {
	plugins.RegisterPluginsWithTCPConnnections(c)
}
