package internal

import (
	"reflect"

	"github.com/zhanglifan/leaf_server/leaf/log"
	"github.com/zhanglifan/leaf_server/src/proto/CFriend"
)

func init() {
	handler(&CFriend.CFriendAdd{}, handleFriendAdd)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleFriendAdd(args []interface{}) {
	log.Debug("handleFriendAdd")
}

func TestAdd() {
	log.Debug("handleFriendAdd")
}
