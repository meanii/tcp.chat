package plugins

import (
	"bufio"
	"net"

	"github.com/meanii/tcp.chat/internal/pkg"
	"github.com/meanii/tcp.chat/utils"
)

type pluginFuncArgs struct {
	conn    net.Conn
	user    pkg.User
	message string
	users   *pkg.Users
}

type plugin struct {
	name          string
	description   string
	priorityIndex int
	function      func(pluginFuncArgs) error
}

var plugins map[string]plugin = make(map[string]plugin)
var users = pkg.InitUsersInstance()

func registerPlugin(plugin plugin) {
	plugins[plugin.name] = plugin
}

func RegisterPluginsWithTCPConnnections(conn net.Conn) {
	reader := bufio.NewReader(conn)
	plugins["welcome"].function(pluginFuncArgs{conn: conn})
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}

		cmd := utils.CommandParser(message)
		if cmd.Command == plugins[cmd.Command].name {
			plugins[cmd.Command].function(pluginFuncArgs{
				conn:    conn,
				message: cmd.Message,
				user:    *users.GetUser(conn),
				users:   users,
			})
		}
	}
}
