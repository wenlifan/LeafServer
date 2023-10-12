package cluster

import (
	"encoding/json"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"net"
	"net/rpc"
	"time"
)

var (
	server      net.Listener
	clients     map[string]*rpc.Client
	clusterConf map[string]map[string]string
)

func Init() {
}

// 开启监听
func listen(addr string) error {
	// 断开 old 监听
	//err := server.Close()
	//if err != nil {
	//	return err
	//}

	// 创建新的连接
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server = l
	return nil
}

// 连接
func connect(node string, addr string) bool {
	tempC, ok := clients[node]
	if ok {
		tempC.Close()
	}

	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Error("connect rpc server fail node:%s addr:%s", node, addr)
		return true
	}

	clients[node] = client
	return false
}

// Load jsonData: 集群配置地址
func Load(jsonData []byte, nodeName string) {
	err := json.Unmarshal(jsonData, &clusterConf)
	log.Debug("Cluster Config:%s nodeName:", string(jsonData), nodeName)
	if err != nil {
		log.Fatal("%v", err)
	}

	// 检查主节点
	nodeServer, ok := clusterConf["node"][nodeName]
	if !ok {
		log.Fatal("cluster config no find self: %s", nodeName)
	}
	listen(nodeServer)

	// 连接其他节点
	for node, addr := range clusterConf["node"] {
		// 连接
		go func() {
			if node == nodeName {
				return
			}
			log.Warning("connect rpc server node:%s, addr:%s", node, addr)
			for connect(node, addr) {
				// 连接服务器失败, 延时再次连接
				log.Warning("Wait for reconnection: ", node, addr)
				time.Sleep(time.Duration(3) * time.Second)
			}
		}()
	}
}

func Destroy() {
	// 服务器
	err := server.Close()
	if err != nil {
		log.Error("Destroy close rpc server err: %s", err)
		//return
	}

	// 客户端
	for node := range clusterConf["node"] {
		// 连接
		if clients[node] != nil {
			err := clients[node].Close()
			if err != nil {
				log.Error("Destroy close rpc client err: %s", err)
				//return
			}
		}
	}
}
