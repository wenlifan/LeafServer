package internal

import (
	"github.com/zhanglifan/leaf_server/leaf/cluster"
	lconf "github.com/zhanglifan/leaf_server/leaf/conf"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/leaf/module"
	"github.com/zhanglifan/leaf_server/src/server/base"
	"github.com/zhanglifan/leaf_server/src/server/conf"
	"io/ioutil"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	//RouterMsg = cluster.NewRouter()
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	// 加载配置文件
	data, err := ioutil.ReadFile(conf.Server.ClusterPath)
	if err != nil {
		log.Fatal("%v", err)
	}

	// Rpc 启动
	cluster.Load(data, lconf.Node)
}

func (m *Module) OnDestroy() {
}
