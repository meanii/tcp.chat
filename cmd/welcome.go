package cmd

import "net"

func WelcomeHandler(conn net.Conn) {
	message := "Welcome to TCP.chat CLI, a robust TCP chat application designed for seamless communication. Effortlessly join public chats, built with Golang.\n\nAuthor: github.com/meanii\n\n"
	message += "Available commands:\n"
	message += "\\login - Log in to your existing account\n"
	message += "\\help - View a list of available commands\n"
	message += "\n\n"
	conn.Write([]byte(message))
}
