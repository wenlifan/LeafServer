module github.com/zhanglifan/leaf_server

go 1.20

require google.golang.org/protobuf v1.30.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/websocket v1.5.0
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

require (
	github.com/go-redis/redis/v8 v8.11.5
	go.uber.org/zap v1.24.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
