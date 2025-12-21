package internal

import (
	"reflect"

	"github.com/zhanglifan/leaf_server/leaf/gate"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/src/proto/CFriend"
	"github.com/zhanglifan/leaf_server/src/proto/Enum"
	"github.com/zhanglifan/leaf_server/src/proto/Friend"
	"github.com/zhanglifan/leaf_server/src/proto/PreLobby"
	"github.com/zhanglifan/leaf_server/src/server/friend"
)

func init() {
	handler(&PreLobby.ReqLogin{}, handlePreLobbyReqLogin)
	// ReqCreateRole
	handler(&PreLobby.ReqCreateRole{}, handlePreLobbyReqCreateRole)
	handler(&PreLobby.ReqEnterLobby{}, handlerPreLobbyReqEnterLobby)
	handler(&Friend.ReqFriendTest{}, handleFriendReqFriendTest)
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

func handleFriendReqFriendTest(args []interface{}) {
	// 收到的 ReqFriendTest 消息
	m := args[0].(*Friend.ReqFriendTest)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("[Friend] ReqFriendTest: %+v", m)

	// 调用 friend服务的 CFriend.CFriendAdd
	// 注意：需要使用类型而不是字符串，因为friend server注册时使用的是类型
	friend.ChanRPC.Go(reflect.TypeOf(&CFriend.CFriendAdd{}), a, m)

	// 发送响应
	a.WriteMsgBase(&Friend.RspFriendTest{
		ErrorCode: Enum.EErrorCode_SUCCESS,
	}, "Friend.RspFriendTest")
}
