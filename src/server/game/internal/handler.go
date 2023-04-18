package internal

import (
	"github.com/zhanglifan/leaf_server/leaf/gate"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/proto/PreLobby"
	"reflect"
)

func init() {
	handler(&PreLobby.ReqLogin{}, handlePreLobbyReqLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePreLobbyReqLogin(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*PreLobby.ReqLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("AccountName: %v", m.GetAccountName())

	a.WriteMsg(&PreLobby.ReqLogin{})
}
