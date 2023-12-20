package base

import (
	"github.com/zhanglifan/leaf_server/leaf/chanrpc"
	"github.com/zhanglifan/leaf_server/leaf/module"
	"github.com/zhanglifan/leaf_server/src/server/conf"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
