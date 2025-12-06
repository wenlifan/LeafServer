#!/bin/bash
# ============================================
# Redis 启动脚本（Docker 方式）
# 用于通过 Docker 启动 Redis 服务
# ============================================

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Redis 配置
REDIS_CONTAINER_NAME="redis-server"
REDIS_PORT="6379"
REDIS_PASSWORD="123456"
REDIS_CONF="redis.conf"

# 检查 Docker 是否安装
check_docker() {
    if ! command -v docker &> /dev/null; then
        echo "错误: 未找到 Docker"
        echo ""
        echo "请先安装 Docker Desktop:"
        echo "1. 访问 https://www.docker.com/products/docker-desktop"
        echo "2. 下载并安装 Docker Desktop for Windows"
        echo "3. 启动 Docker Desktop"
        echo "4. 重新运行此脚本"
        exit 1
    fi
    
    # 检查 Docker 是否运行
    if ! docker info &> /dev/null; then
        echo "错误: Docker 未运行"
        echo ""
        echo "请启动 Docker Desktop 后重试"
        exit 1
    fi
}

# 检查 Redis 容器是否已运行
check_redis_running() {
    # 检查容器是否存在且正在运行
    if docker ps --format '{{.Names}}' | grep -q "^${REDIS_CONTAINER_NAME}$"; then
        echo "Redis 容器已经在运行中"
        echo "容器名称: ${REDIS_CONTAINER_NAME}"
        echo "端口: ${REDIS_PORT}"
        echo "密码: ${REDIS_PASSWORD}"
        return 0
    fi
    
    # 检查容器是否存在但已停止
    if docker ps -a --format '{{.Names}}' | grep -q "^${REDIS_CONTAINER_NAME}$"; then
        echo "发现已停止的 Redis 容器，正在启动..."
        docker start "${REDIS_CONTAINER_NAME}" > /dev/null 2>&1
        if [ $? -eq 0 ]; then
            echo "Redis 容器已启动"
            echo "容器名称: ${REDIS_CONTAINER_NAME}"
            echo "端口: ${REDIS_PORT}"
            echo "密码: ${REDIS_PASSWORD}"
            return 0
        else
            echo "警告: 启动已存在的容器失败，将删除并重新创建"
            docker rm -f "${REDIS_CONTAINER_NAME}" > /dev/null 2>&1
        fi
    fi
    
    return 1
}

# 启动 Redis
start_redis() {
    # 检查 Docker
    check_docker
    
    # 检查是否已运行
    if check_redis_running; then
        return 0
    fi
    
    echo "正在启动 Redis (Docker)..."
    
    # 检查配置文件是否存在
    if [ ! -f "$REDIS_CONF" ]; then
        echo "警告: 配置文件 $REDIS_CONF 不存在，使用默认配置（仅设置密码）"
        # 使用简单方式启动，只设置密码
        docker run -d \
            --name "${REDIS_CONTAINER_NAME}" \
            -p "${REDIS_PORT}:6379" \
            redis:latest \
            redis-server --requirepass "${REDIS_PASSWORD}"
    else
        echo "使用配置文件: $REDIS_CONF"
        # 使用配置文件启动
        docker run -d \
            --name "${REDIS_CONTAINER_NAME}" \
            -p "${REDIS_PORT}:6379" \
            -v "${SCRIPT_DIR}/${REDIS_CONF}:/usr/local/etc/redis/redis.conf:ro" \
            redis:latest \
            redis-server /usr/local/etc/redis/redis.conf \
            --requirepass "${REDIS_PASSWORD}"
    fi
    
    if [ $? -eq 0 ]; then
        # 等待容器启动
        sleep 2
        
        # 验证 Redis 是否正常启动
        if docker exec "${REDIS_CONTAINER_NAME}" redis-cli -a "${REDIS_PASSWORD}" ping &> /dev/null; then
            echo ""
            echo "✓ Redis 已成功启动"
            echo "  容器名称: ${REDIS_CONTAINER_NAME}"
            echo "  端口: ${REDIS_PORT}"
            echo "  密码: ${REDIS_PASSWORD}"
            echo ""
            echo "常用命令:"
            echo "  查看日志: docker logs ${REDIS_CONTAINER_NAME}"
            echo "  连接 Redis: docker exec -it ${REDIS_CONTAINER_NAME} redis-cli -a ${REDIS_PASSWORD}"
            echo "  停止 Redis: bash redis_stop.sh"
        else
            echo "警告: Redis 容器已启动，但连接测试失败"
            echo "请检查日志: docker logs ${REDIS_CONTAINER_NAME}"
        fi
    else
        echo "错误: Docker 启动失败"
        echo "请检查 Docker 是否正常运行，或查看错误信息"
        exit 1
    fi
}

# 主函数
main() {
    start_redis
}

main "$@"

