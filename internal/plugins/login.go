package plugins

import (
	"github.com/meanii/tcp.chat/internal/pkg"
)

func init() {
	registerPlugin(plugin{
		name:          "login",
		description:   "login to the app",
		priorityIndex: 1,
		function: func(args pluginFuncArgs) error {
			user := args.users.Create(args.message, args.conn)
			pkg.NewMessage(args.conn, user.Id, "welcome to the chat!")
			return nil
		},
	})
}
