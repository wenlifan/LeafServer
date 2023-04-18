package main

import (
	"github.com/zhanglifan/leaf_server/leaf"
	lconf "github.com/zhanglifan/leaf_server/leaf/conf"
	"github.com/zhanglifan/leaf_server/src/server/conf"
	"github.com/zhanglifan/leaf_server/src/server/game"
	"github.com/zhanglifan/leaf_server/src/server/gate"
	"github.com/zhanglifan/leaf_server/src/server/login"
	"github.com/zhanglifan/leaf_server/src/server/web"
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
		web.Module,
	)
}
