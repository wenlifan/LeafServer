package conf

import (
	"encoding/json"
	"github.com/zhanglifan/leaf_server/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

var FriendServer struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	data, err := ioutil.ReadFile("bin/conf/gameserver.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}

	dataFriend, err := ioutil.ReadFile("bin/conf/friendserver.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(dataFriend, &FriendServer)
	if err != nil {
		log.Fatal("%v", err)
	}
}
