package internal

import (
	"reflect"

	"github.com/zhanglifan/leaf_server/leaf/cluster"
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
	// 参考 Heartbeat 的调用方式，使用 cluster.Call 进行 RPC 调用
	rpcArgs := &CFriend.CFriendAdd{
		RoleUID:   "test_role_uid",   // 这里需要根据实际业务从 m 中获取
		FriendUID: "test_friend_uid", // 这里需要根据实际业务从 m 中获取
		ListType:  "friend",
	}
	var reply friend.CFriendAddReply
	err := cluster.Call("friend", "FriendMsg.FriendAdd", rpcArgs, &reply)
	if err != nil {
		log.Error("[Game] 调用friend服务失败: %v", err)
	} else {
		log.Debug("[Game] 好友添加结果: Success=%v, Message=%s", reply.Success, reply.Message)
	}

	// 发送响应
	a.WriteMsgBase(&Friend.RspFriendTest{
		ErrorCode: Enum.EErrorCode_SUCCESS,
	}, "Friend.RspFriendTest")
}
