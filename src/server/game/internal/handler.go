package internal

import (
	"reflect"

	"github.com/zhanglifan/leaf_server/leaf/gate"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/src/proto/PreLobby"
)

func init() {
	handler(&PreLobby.ReqLogin{}, handlePreLobbyReqLogin)
	// ReqCreateRole
	handler(&PreLobby.ReqCreateRole{}, handlePreLobbyReqCreateRole)
	handler(&PreLobby.ReqEnterLobby{}, handlerPreLobbyReqEnterLobby)
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
	log.Debug("[PreLobby] ReqLogin AccountName: %+v", m)

	a.WriteMsgBase(&PreLobby.RspRoleInfo{
		RoleName: "testName",
		RoleUID:  12341345,
	}, "PreLobby.RspRoleInfo")
}

func handlePreLobbyReqCreateRole(args []interface{}) {
	m := args[0].(*PreLobby.ReqCreateRole)
	a := args[1].(gate.Agent)

	log.Debug("[PreLobby] ReqCreateRole RoleName: %+v", m)

	a.WriteMsgBase(&PreLobby.RspCreateRole{
		Succeed:  true,
		RoleName: m.GetRoleName(),
	}, "PreLobby.RspCreateRole")
}

func handlerPreLobbyReqEnterLobby(args []interface{}) {
	m := args[0].(*PreLobby.ReqEnterLobby)
	a := args[1].(gate.Agent)

	log.Debug("[PreLobby] ReqEnterLobby AccountName: %+v", m)

	a.WriteMsgBase(&PreLobby.RspEnterLobby{}, "PreLobby.RspEnterLobby")
}
