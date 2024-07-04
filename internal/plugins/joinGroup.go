package plugins

import (
	"fmt"

	"github.com/meanii/tcp.chat/internal/pkg"
)

func init() {
	registerPlugin(plugin{
		name:          "join_group",
		description:   "join any group by id",
		priorityIndex: 1,
		function: func(pfa pluginFuncArgs) error {
			roomId := pfa.message
			room := pfa.rooms.JoinGroup(roomId, pfa.user)
			pkg.NewMessage(pfa.conn, room.Id, "welcome to the room.")
			room.RoomBroadcastNotification(pfa.conn, fmt.Sprintf("%s has joined the room.", pfa.user.Id), pkg.NewMessage)
			return nil
		},
	})
}
