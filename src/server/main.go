package main

import (
	"zlf_leaf/src/frame/leaf"
	lconf "zlf_leaf/src/frame/leaf/conf"
	"zlf_leaf/src/server/conf"
	"zlf_leaf/src/server/game"
	"zlf_leaf/src/server/gate"
	"zlf_leaf/src/server/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
