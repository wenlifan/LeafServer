package msg

import (
	"github.com/zhanglifan/leaf_server/leaf/network/protobuf"
	"github.com/zhanglifan/leaf_server/src/proto/Friend"
	"github.com/zhanglifan/leaf_server/src/proto/PreLobby"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&PreLobby.ReqLogin{}, "PreLobby.ReqLogin")
	Processor.Register(&PreLobby.ReqCreateRole{}, "PreLobby.ReqCreateRole")
	Processor.Register(&PreLobby.ReqEnterLobby{}, "PreLobby.ReqEnterLobby")
	Processor.Register(&Friend.ReqFriendTest{}, "Friend.ReqFriendTest")
}
