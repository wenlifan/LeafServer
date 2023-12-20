package cluster

import (
	"net/rpc"
)

func init() {
	myConnMsg := new(ConnMsg)
	rpc.Register(myConnMsg)
}

// RPC: 连接类管理消息
type ConnMsg struct{}

// 心跳消息结构体
type HeartbeatMsg struct {
	Msg string
}

// 连接类: 处理心跳消息
func (c *ConnMsg) Heartbeat(args *HeartbeatMsg, reply *HeartbeatMsg) error {
	reply.Msg = "pong"
	return nil
}
