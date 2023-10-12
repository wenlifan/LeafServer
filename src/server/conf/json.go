package conf

import (
	"encoding/json"
	"flag"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"io/ioutil"
)

var Serverstruct {
LogLevel    string
LogPath     string
WSAddr      string
CertFile    string
KeyFile     string
TCPAddr     string
MaxConnNum  int
ConsolePort int
ProfilePath string
Node        string
ClusterPath string
}

func init() {
	var configPath = ""
	var clusterPath = ""
	flag.StringVar(&configPath, "config", "", "配置")
	flag.StringVar(&clusterPath, "cluster", "", "集群")
	flag.Parse()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
	Server.ClusterPath = clusterPath
}
