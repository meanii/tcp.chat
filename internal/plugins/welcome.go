package plugins

func init() {
	registerPlugin(plugin{
		name:          "welcome",
		description:   "welcome message",
		priorityIndex: 0,
		function: func(args pluginFuncArgs) error {
			msg := "Welcome to tcp.chat!\n"
			msg += "This is a raw TCP chat application designed for easy communication and public chat participation.\n\n"
			msg += "To get started, please log in using the following command: \\login {username}\n"
			msg += "For more information, use the \\help command.\n\nlogin # > "
			args.conn.Write([]byte(msg))
			return nil
		},
	})
}
