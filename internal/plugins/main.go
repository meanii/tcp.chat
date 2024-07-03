package plugins

import (
	"bufio"
	"net"

	"github.com/meanii/tcp.chat/utils"
)

type plugin struct {
	name          string
	description   string
	priorityIndex int
	function      func(conn net.Conn, message string) error
}

var plugins map[string]plugin = make(map[string]plugin)

func registerPlugin(plugin plugin) {
	plugins[plugin.name] = plugin
}

func RegisterPluginsWithTCPConnnections(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}

		cmd := utils.CommandParser(message)
		if cmd.Command == plugins[cmd.Command].name {
			plugins[cmd.Command].function(conn, cmd.Message)
		}
	}
}
