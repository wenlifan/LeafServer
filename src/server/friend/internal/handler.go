package internal

import (
	"net/rpc"
	"reflect"

	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/src/proto/CFriend"
)

// RPC: 好友类管理消息
type FriendMsg struct{}

// FriendAddResult 好友添加处理结果
type FriendAddResult struct {
	Success bool   // 是否成功
	Message string // 返回消息
}

func init() {
	handler(&CFriend.CFriendAdd{}, handleFriendAdd)
	myFriendMsg := new(FriendMsg)
	rpc.Register(myFriendMsg)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// handleFriendAdd 处理好友添加请求
// 返回处理结果：FriendAddResult结构，包含Success和Message
func handleFriendAdd(args []interface{}) interface{} {
	// 提取消息
	m := args[0].(*CFriend.CFriendAdd)

	log.Debug("[Friend] 收到好友添加请求: RoleUID=%s, FriendUID=%s, ListType=%s", m.RoleUID, m.FriendUID, m.ListType)

	// 业务处理逻辑
	// TODO: 在这里实现具体的好友添加业务逻辑
	// 例如：检查好友关系、添加到数据库等
	log.Debug("[Friend] 处理好友添加: RoleUID=%s 添加 FriendUID=%s 到 ListType=%s", m.RoleUID, m.FriendUID, m.ListType)

	// 业务逻辑处理
	// 例如：检查是否已经是好友、添加到数据库等
	// 这里只是示例，实际需要根据业务需求实现
	// 如果处理失败，返回错误信息
	// if err := doAddFriend(m); err != nil {
	//     return &FriendAddResult{
	//         Success: false,
	//         Message: err.Error(),
	//     }
	// }

	log.Debug("[Friend] 好友添加处理完成")
	return &FriendAddResult{
		Success: true,
		Message: "好友添加成功",
	}
}

func TestAdd() {
	log.Debug("handleFriendAdd")
}

// 好友添加回复结构体
type CFriendAddReply struct {
	Success bool   // 是否成功
	Message string // 返回消息
}

// 好友类: 处理好友添加消息
// 该方法只负责接收RPC请求并转发到friend模块的handler，所有业务逻辑在handler中处理
func (f *FriendMsg) FriendAdd(args *CFriend.CFriendAdd, reply *CFriendAddReply) error {

	handleFriendAdd([]interface{}{args})

	reply.Success = true
	reply.Message = "好友添加成功"
	return nil
}
