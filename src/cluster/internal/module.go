package internal

import (
	"io/ioutil"

	"github.com/zhanglifan/leaf_server/leaf/cluster"
	lconf "github.com/zhanglifan/leaf_server/leaf/conf"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/leaf/module"
	"github.com/zhanglifan/leaf_server/src/server/base"
	"github.com/zhanglifan/leaf_server/src/server/conf"
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

	// Rpc 配置加载（不阻塞）
	cluster.Load(data, lconf.Node)
}

func (m *Module) Run(closeSig chan bool) {
	// 在独立的goroutine中启动HTTP RPC服务器
	go cluster.Start()
	// 运行Skeleton的事件循环
	m.Skeleton.Run(closeSig)
}

func (m *Module) OnDestroy() {
}
