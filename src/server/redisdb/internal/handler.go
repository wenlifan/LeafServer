package internal

import (
	"context"
	"github.com/zhanglifan/leaf_server/leaf/db"
	"github.com/zhanglifan/leaf_server/leaf/log"
)

var (
	RedisDB = new(db.RedisSingleObj)
)

func init() {
	connectDB()

	skeleton.RegisterChanRPC("Set", rpcSet)
}

func connectDB() {
	RedisDB.RedisHost = "192.168.2.174"
	RedisDB.RedisPort = 9001
	RedisDB.RedisAuth = "wtredis@123"
	RedisDB.Database = 0

	RedisDB.ConnectDB()
}

/*
Redis命令调用模板:

redisArgs := db.BaseCommandArgs{
	Cmd:  "Get",
	Key:  "zhanglifan_tset",
	Args: []string{"test"},
}

error := redisdb.ChanRPC.Call0("Set", redisArgs)
if error != nil {
	log.Error("set error: ", error.Error())
}
*/

// set 命令使用
func rpcSet(args []interface{}) interface{} {
	ctx := context.Background()

	commandArgs := args[0].(db.BaseCommandArgs)
	log.Debug("ExeCommand Args: ", commandArgs)
	// StatusCmd
	err := RedisDB.Db.Set(ctx, commandArgs.Key, commandArgs.Args[0], 0).Err()
	if err != nil {
		log.Error("rpcSet: ", err.Error())
	}
	return err.(error)
}
