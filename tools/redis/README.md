# Redis 安装和使用说明

## 概述

本项目使用 Redis 作为数据存储服务。**启动脚本专门使用 Docker 方式运行 Redis**，确保跨平台一致性和易于管理。

## 配置文件

- **配置文件**: `redis.conf`
- **默认端口**: 6379
- **默认密码**: 123456
- **容器名称**: `redis-server`

## 前置要求

### 安装 Docker Desktop

启动脚本需要 Docker 环境，请先安装 Docker Desktop：

1. 访问 https://www.docker.com/products/docker-desktop
2. 下载并安装 Docker Desktop for Windows
3. 启动 Docker Desktop
4. 确保 Docker 正在运行（系统托盘显示 Docker 图标）

## 使用脚本（推荐）

### 启动 Redis

```bash
cd tools/redis
bash redis_start.sh
```

**脚本功能**：
- 自动检测 Docker 是否安装和运行
- 检查 Redis 容器是否已存在
- 如果容器已停止，自动启动
- 如果容器不存在，创建新容器
- 自动验证 Redis 是否正常启动

**启动方式**：
- 如果存在 `redis.conf` 配置文件，使用配置文件启动
- 如果不存在配置文件，使用默认配置（仅设置密码）

### 停止 Redis

```bash
cd tools/redis
bash redis_stop.sh
```

**停止方式**：
- 正常停止：会询问是否删除容器
- 强制停止：`bash redis_stop.sh -f` 或 `bash redis_stop.sh --force`（直接停止并删除容器）

## 手动使用 Docker 命令

### 启动 Redis 容器

**方式一：使用配置文件（推荐）**

```bash
cd tools/redis
docker run -d \
    --name redis-server \
    -p 6379:6379 \
    -v "$(pwd)/redis.conf:/usr/local/etc/redis/redis.conf:ro" \
    redis:latest \
    redis-server /usr/local/etc/redis/redis.conf \
    --requirepass 123456
```

**方式二：仅设置密码（不使用配置文件）**

```bash
docker run -d \
    --name redis-server \
    -p 6379:6379 \
    redis:latest \
    redis-server --requirepass 123456
```

### 查看容器状态

```bash
# 查看运行中的容器
docker ps

# 查看所有容器（包括已停止的）
docker ps -a

# 查看 Redis 容器日志
docker logs redis-server

# 实时查看日志
docker logs -f redis-server
```

### 停止和删除容器

```bash
# 停止容器
docker stop redis-server

# 删除容器
docker rm redis-server

# 停止并删除（一步完成）
docker stop redis-server && docker rm redis-server
```

### 重启容器

```bash
# 如果容器已存在但已停止
docker start redis-server

# 重启运行中的容器
docker restart redis-server
```

## 验证安装

### 使用 Docker 连接 Redis CLI（推荐）

```bash
# 进入 Redis CLI
docker exec -it redis-server redis-cli -a 123456

# 测试连接
127.0.0.1:6379> PING
PONG

# 查看 Redis 信息
127.0.0.1:6379> INFO server

# 退出
127.0.0.1:6379> exit
```

### 测试 Redis 连接

```bash
# 直接执行命令测试
docker exec redis-server redis-cli -a 123456 PING

# 设置和获取值测试
docker exec redis-server redis-cli -a 123456 SET test "Hello Redis"
docker exec redis-server redis-cli -a 123456 GET test
```

### 使用本地 redis-cli 连接（如果已安装）

如果本地已安装 redis-cli（如通过 WSL），可以直接连接：

```bash
redis-cli -h 127.0.0.1 -p 6379 -a 123456

# 测试连接
127.0.0.1:6379> PING
PONG
```

## 配置说明

### 主要配置项

- **bind**: 127.0.0.1（仅本地访问）
- **port**: 6379（默认端口）
- **requirepass**: 123456（访问密码）
- **appendonly**: yes（启用 AOF 持久化）
- **databases**: 16（数据库数量）

### 修改配置

编辑 `redis.conf` 文件，修改相应配置项后重启 Redis。

## 常见问题

### 1. Docker 未安装或未运行

**错误信息**：`错误: 未找到 Docker` 或 `错误: Docker 未运行`

**解决方法**：
- 安装 Docker Desktop：https://www.docker.com/products/docker-desktop
- 启动 Docker Desktop
- 等待 Docker 完全启动（系统托盘图标显示运行中）

### 2. 端口被占用

**错误信息**：`Error response from daemon: driver failed programming external connectivity`

**解决方法**：
- 检查端口占用：`netstat -ano | findstr 6379`（Windows）或 `lsof -i :6379`（Linux/WSL）
- 修改 `redis.conf` 中的 `port` 配置为其他端口（如 6380）
- 启动时指定新端口：`docker run -p 6380:6380 ...`

### 3. 容器名称已存在

**错误信息**：`Error response from daemon: Conflict. The container name "redis-server" is already in use`

**解决方法**：
```bash
# 删除已存在的容器
docker rm -f redis-server

# 或使用不同的容器名称启动
docker run -d --name redis-server-new ...
```

### 4. 密码验证失败

**确保**：
- 配置文件中的 `requirepass` 设置为 `123456`
- 连接时使用正确的密码：`redis-cli -a 123456`
- 如果使用配置文件，确保 `--requirepass 123456` 参数已传递

### 5. 配置文件挂载失败

**错误信息**：配置文件相关错误

**解决方法**：
- 确保 `redis.conf` 文件存在于 `tools/redis/` 目录
- 检查文件路径是否正确（使用绝对路径或相对路径）
- 如果配置文件有问题，可以删除配置文件，脚本会使用默认配置

### 6. 容器启动后无法连接

**检查步骤**：
```bash
# 1. 检查容器是否运行
docker ps | grep redis-server

# 2. 查看容器日志
docker logs redis-server

# 3. 检查端口映射
docker port redis-server

# 4. 测试容器内连接
docker exec redis-server redis-cli -a 123456 PING
```

## 项目中的使用

项目代码中 Redis 连接配置位于：

- `src/server/redisdb/internal/handler.go`

当前配置：
- Host: 192.168.2.174
- Port: 9001
- Password: wtredis@123

如需使用本地 Redis，请修改为：
- Host: 127.0.0.1
- Port: 6379
- Password: 123456

## 日志查看

### 查看容器日志

```bash
# 查看所有日志
docker logs redis-server

# 实时查看日志（类似 tail -f）
docker logs -f redis-server

# 查看最近 100 行日志
docker logs --tail 100 redis-server
```

### 日志位置

- **容器日志**：通过 `docker logs` 命令查看
- **Redis 日志**：如果配置了 `logfile`，日志在容器内的指定路径（需要挂载卷才能持久化）

## 数据持久化

Redis 配置了两种持久化方式：

1. **RDB 快照**: 定期保存数据快照（配置了 `save` 规则）
2. **AOF 日志**: 记录所有写操作（`appendonly yes`）

### 当前配置说明

**注意**：当前 Docker 启动方式**未挂载数据卷**，容器删除后数据会丢失。

### 启用数据持久化（推荐）

如果需要数据持久化，可以修改启动命令挂载数据卷：

```bash
docker run -d \
    --name redis-server \
    -p 6379:6379 \
    -v "$(pwd)/redis.conf:/usr/local/etc/redis/redis.conf:ro" \
    -v "$(pwd)/data:/data" \
    redis:latest \
    redis-server /usr/local/etc/redis/redis.conf \
    --requirepass 123456
```

这样数据会保存在 `tools/redis/data/` 目录中，即使容器删除，数据也不会丢失。

### 数据备份

```bash
# 进入容器执行备份
docker exec redis-server redis-cli -a 123456 SAVE

# 或使用 BGSAVE（后台保存）
docker exec redis-server redis-cli -a 123456 BGSAVE

# 复制数据文件（如果挂载了数据卷）
# 数据在挂载的目录中，直接复制目录即可
```

## 其他安装方式（参考）

虽然脚本使用 Docker 方式，但以下方式也可以手动安装使用：

### 使用 WSL

如果已安装 WSL（Windows Subsystem for Linux），可以在 WSL 中安装 Redis：

```bash
# 在 WSL 中执行
sudo apt-get update
sudo apt-get install redis-server -y

# 启动 Redis（使用配置文件）
redis-server /path/to/redis.conf
```

### 使用 Memurai（Windows 原生）

Memurai 是 Redis 的 Windows 兼容版本：

1. 下载 Memurai: https://www.memurai.com/
2. 安装 Memurai
3. 手动配置和启动

### 使用 Redis for Windows 端口版本

可以从以下位置下载 Windows 版本的 Redis：

- tporadowski/redis: https://github.com/tporadowski/redis/releases
- 下载后解压并手动配置启动

