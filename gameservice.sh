#!/bin/bash
# ============================================
# Game Service 启动脚本
# 用于启动游戏服务
# ============================================

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# 启动服务
go run ./cmd/gameserver/main.go \
    -config="bin/conf/gameserver.json" \
    -cluster="bin/conf/cluster.json"
