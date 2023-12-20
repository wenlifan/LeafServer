package redisdb

import "github.com/zhanglifan/leaf_server/src/server/redisdb/internal"

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
