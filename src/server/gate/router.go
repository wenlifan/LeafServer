package gate

import (
	"github.com/zhanglifan/leaf_server/src/proto/PreLobby"
	"github.com/zhanglifan/leaf_server/src/server/game"
	"github.com/zhanglifan/leaf_server/src/server/msg"
)

func init() {
	msg.Processor.SetRouter(&PreLobby.ReqLogin{}, game.ChanRPC)
	msg.Processor.SetRouter(&PreLobby.ReqCreateRole{}, game.ChanRPC)
	msg.Processor.SetRouter(&PreLobby.ReqEnterLobby{}, game.ChanRPC)
}
