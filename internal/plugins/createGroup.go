package plugins

import (
	"github.com/meanii/tcp.chat/internal/pkg"
)

func init() {
	registerPlugin(plugin{
		name:          "create_group",
		description:   "create a new group",
		priorityIndex: 1,
		function: func(pfa pluginFuncArgs) error {
			if pfa.user.Id == "" {
				pkg.NewMessage(pfa.conn, pfa.user.Id, "login first, in order to create new group.")
				return nil
			}
			room := pfa.rooms.CreateGroup(pfa.message)
			pfa.rooms.JoinGroup(room.Id, pfa.user)
			pkg.NewMessage(pfa.conn, room.Id, "welcome to the group.")
			return nil
		},
	})
}
