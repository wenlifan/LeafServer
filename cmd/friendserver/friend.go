package main

import (
	"github.com/zhanglifan/leaf_server/leaf"
	lconf "github.com/zhanglifan/leaf_server/leaf/conf"
	"github.com/zhanglifan/leaf_server/src/server/conf"
	"github.com/zhanglifan/leaf_server/src/server/game"
	"github.com/zhanglifan/leaf_server/src/server/gate"
	"github.com/zhanglifan/leaf_server/src/server/redisdb"
)

func main() {
	lconf.LogLevel = conf.FriendServer.LogLevel
	lconf.LogPath = conf.FriendServer.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.FriendServer.ConsolePort
	lconf.ProfilePath = conf.FriendServer.ProfilePath

	leaf.Run(
		game.Module,
		gate.Module,
		redisdb.Module,
	)
}
