package gate

import (
	"github.com/zhanglifan/leaf_server/src/server/game"
	"github.com/zhanglifan/leaf_server/src/server/msg"
	"github.com/zhanglifan/proto/PreLobby"
)

func init() {
	msg.Processor.SetRouter(&PreLobby.ReqLogin{}, game.ChanRPC)
}
