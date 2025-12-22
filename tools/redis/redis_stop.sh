#!/bin/bash
# ============================================
# Redis 停止脚本（Docker 方式）
# 用于停止 Docker 中的 Redis 服务
# ============================================

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

REDIS_CONTAINER_NAME="redis-server"

# 停止 Redis
stop_redis() {
    # 检查 Docker 容器是否存在
    if docker ps -a --format '{{.Names}}' | grep -q "^${REDIS_CONTAINER_NAME}$"; then
        # 检查容器是否正在运行
        if docker ps --format '{{.Names}}' | grep -q "^${REDIS_CONTAINER_NAME}$"; then
            echo "正在停止 Redis 容器..."
            docker stop "${REDIS_CONTAINER_NAME}" > /dev/null 2>&1
            if [ $? -eq 0 ]; then
                echo "✓ Redis 容器已停止"
            else
                echo "错误: 停止容器失败"
                return 1
            fi
        else
            echo "Redis 容器已停止"
        fi
        
        # 询问是否删除容器
        echo "是否删除容器? (y/N)"
        read -r response
        if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
            docker rm "${REDIS_CONTAINER_NAME}" > /dev/null 2>&1
            if [ $? -eq 0 ]; then
                echo "✓ Redis 容器已删除"
            else
                echo "错误: 删除容器失败"
                return 1
            fi
        else
            echo "容器已保留，可以使用 'docker start ${REDIS_CONTAINER_NAME}' 重新启动"
        fi
        return 0
    else
        echo "Redis 容器不存在，无需停止"
        return 1
    fi
}

# 强制停止并删除（不询问）
force_stop_redis() {
    if docker ps -a --format '{{.Names}}' | grep -q "^${REDIS_CONTAINER_NAME}$"; then
        echo "正在强制停止并删除 Redis 容器..."
        docker stop "${REDIS_CONTAINER_NAME}" > /dev/null 2>&1
        docker rm "${REDIS_CONTAINER_NAME}" > /dev/null 2>&1
        if [ $? -eq 0 ]; then
            echo "✓ Redis 容器已停止并删除"
            return 0
        else
            echo "错误: 操作失败"
            return 1
        fi
    else
        echo "Redis 容器不存在"
        return 1
    fi
}

# 主函数
main() {
    # 如果传入 -f 或 --force 参数，强制停止并删除
    if [ "$1" = "-f" ] || [ "$1" = "--force" ]; then
        force_stop_redis
    else
        stop_redis
    fi
}

main "$@"

