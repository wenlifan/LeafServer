package internal

import (
	"github.com/zhanglifan/leaf_server/leaf/module"
	"github.com/zhanglifan/leaf_server/src/server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {
}
