module github.com/zhanglifan/leaf_server

go 1.20

require (
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.5.0
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/zhanglifan/proto v0.0.0
	go.uber.org/zap v1.24.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

replace github.com/zhanglifan/proto => ../../pkg/mod/github.com/zhanglifan/proto
