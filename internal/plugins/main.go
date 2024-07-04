package plugins

import (
	"bufio"
	"fmt"
	"net"

	"github.com/meanii/tcp.chat/internal/pkg"
	"github.com/meanii/tcp.chat/utils"
)

type pluginFuncArgs struct {
	conn    net.Conn
	message string
	user    pkg.User
	room    pkg.Room
	users   *pkg.Users
	rooms   *pkg.Rooms
}

type plugin struct {
	name          string
	description   string
	priorityIndex int
	function      func(pluginFuncArgs) error
}

var plugins map[string]plugin = make(map[string]plugin)
var users = pkg.InitUsersInstance()
var rooms = pkg.InitRoomsInstance()

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
		plugin, ok := plugins[cmd.Command]

		user := users.GetUser(conn)
		room := rooms.GetRoomByUser(user)

		if !ok {
			var activeSession string
			if room.Id != "" {
				activeSession = fmt.Sprintf("%s ( %s )", room.Id, user.Id)
				room.RoomBroadcast(conn, room.Id, user.Id, message, pkg.NewGroupMessage)
			}
			if activeSession == "" && user.Id != "" {
				activeSession = user.Id
			}
			pkg.NewMessage(conn, activeSession, "\n")
			continue
		}

		if cmd.Command == plugin.name {
			plugin.function(pluginFuncArgs{
				conn:    conn,
				message: cmd.Message,
				user:    *user,
				room:    *room,
				users:   users,
				rooms:   rooms,
			})
		}
	}
}
