package friend

import (
	"github.com/zhanglifan/leaf_server/src/server/friend/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

// CFriendAddReply 好友添加回复结构体（从internal包导出）
type CFriendAddReply = internal.CFriendAddReply
