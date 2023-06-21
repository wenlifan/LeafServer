package db

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"time"
)

// 定义一个RedisSingleObj结构体
type RedisSingleObj struct {
	RedisHost string
	RedisPort uint16
	RedisAuth string
	Database  int
	Db        *redis.Client
}

// 结构体InitSingleRedis方法: 用于初始化redis数据库
func (r *RedisSingleObj) ConnectDB() (err error) {
	// Redis连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", r.RedisHost, r.RedisPort)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器。
	r.Db = redis.NewClient(&redis.Options{
		Addr:        redisAddr,   // redis服务ip:port
		Password:    r.RedisAuth, // redis的认证密码
		DB:          r.Database,  // 连接的database库
		IdleTimeout: 300,         // 默认Idle超时时间
		PoolSize:    100,         // 连接池
	})
	fmt.Printf("Connecting Redis : %v\n", redisAddr)

	// 需要使用context库
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 验证是否连接到redis服务端
	res, err := r.Db.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Connect Failed! Err: %v\n", err)
		return err
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
		return nil
	}
}

func (r *RedisSingleObj) DisConnectDB() {
	r.Db.Close()
}

// Redis 命令执行
type BaseCommandArgs struct {
	Cmd  string   `json:"cmd"`
	Key  string   `json:"key"`
	Args []string `json:"args"`
}
