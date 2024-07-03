package plugins

import (
	"fmt"

	"github.com/meanii/tcp.chat/internal/pkg"
)

func helpMessage() string {
	msg := "\n"
	for _, plugin := range plugins {
		if plugin.priorityIndex != 0 {
			msg += fmt.Sprintf("\\%s - %s\n", plugin.name, plugin.description)
		}
	}
	return msg
}

func init() {
	registerPlugin(plugin{
		name:          "help",
		description:   "for this message",
		priorityIndex: 1,
		function: func(args pluginFuncArgs) error {
			user := args.user
			var name string
			name = user.Id
			if user.Id == "" {
				name = "login #"
			}

			pkg.NewMessage(args.conn, name, helpMessage())
			return nil
		},
	})
}
