package cluster

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/zhanglifan/leaf_server/leaf/log"
)

var (
	server      net.Listener
	clients     map[string]*rpc.Client
	clusterConf map[string]map[string]string
)

func init() {
	clients = make(map[string]*rpc.Client)
	clusterConf = make(map[string]map[string]string)
}

// 开启监听
func listen(addr string) error {
	rpc.HandleHTTP()
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

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Error("connect rpc server fail node:%s addr:%s", node, addr)
		return true
	}

	clients[node] = client
	log.Warning("connect rpc server success node:%s addr:%s", node, addr)
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
	err = listen(nodeServer)
	if err != nil {
		log.Fatal("listen rpc server err: %s", err)
	}

	go func() {
		// 连接其他节点
		for node, addr := range clusterConf["node"] {
			// 连接
			go func(node string, addr string, nodeName string) {
				if node != nodeName {
					log.Warning("connect rpc server node:%s, addr:%s", node, addr)
					for connect(node, addr) {
						// 连接服务器失败, 延时再次连接
						log.Warning("Wait for reconnection: ", node, addr)
						time.Sleep(time.Duration(3) * time.Second)
					}
				}
			}(node, addr, nodeName)
		}
	}()

	// TODO: 重连
	go func() {
		for {
			// 间隔3秒 发起心跳检测
			time.Sleep(time.Duration(3) * time.Second)
			for node, rpcConnect := range clients {
				var reply HeartbeatMsg
				args := HeartbeatMsg{Msg: "ping"}
				err := rpcConnect.Call("ConnMsg.Heartbeat", args, &reply)
				if err != nil {
					log.Error("重连 Heartbeat err: %s", err)
					// 连接失败, 重新连接
					for connect(node, clusterConf["node"][node]) {
						// 连接服务器失败, 延时再次连接
						log.Warning("重新连接失败: ", node, clusterConf["node"][node])
					}
					//} else {
					//	log.Warning("Heartbeat reply: %s", reply.Msg)
				}
			}
		}
	}()
}

// Call 调用指定节点的RPC方法
// node: 目标节点名称（如"friend"、"game"等）
// method: RPC方法名（如"FriendMsg.FriendAdd"）
// args: 请求参数
// reply: 回复参数（必须是指针）
func Call(node string, method string, args interface{}, reply interface{}) error {
	client, ok := clients[node]
	if !ok {
		return fmt.Errorf("node %s not connected", node)
	}
	return client.Call(method, args, reply)
}

// Start 启动HTTP RPC服务器，应该在模块的Run方法中调用
func Start() {
	if server == nil {
		log.Fatal("cluster server not initialized, call Load first")
	}
	http.Serve(server, nil)
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
