#!/bin/bash
# ============================================
# Protocol Buffers 代码生成脚本 (跨平台版本)
# 生成目录: ../../src/proto
# 支持 Windows 和 Linux 环境
# ============================================

set -e  # 遇到错误立即退出

echo "============================================"
echo "开始生成 Protocol Buffers 代码..."
echo "============================================"
echo

# 获取脚本所在目录
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# 检测操作系统
OS="$(uname -s)"
case "${OS}" in
    Linux*)     PLATFORM="linux";;
    Darwin*)    PLATFORM="darwin";;
    CYGWIN*)    PLATFORM="win";;
    MINGW*)     PLATFORM="win";;
    MSYS*)      PLATFORM="win";;
    *)          PLATFORM="win";;
esac

echo "检测到操作系统: $OS (平台: $PLATFORM)"
echo

# 设置工具路径
if [ "$PLATFORM" = "win" ]; then
    PROTOC="win/protoc.exe"
    PROTOC_GEN_GO="win/protoc-gen-go.exe"
    PROTOC_GEN_GO_GRPC="win/protoc-gen-go-grpc.exe"
else
    PROTOC="protoc"
    PROTOC_GEN_GO="protoc-gen-go"
    PROTOC_GEN_GO_GRPC="protoc-gen-go-grpc"
fi

# 检查工具是否存在
if [ "$PLATFORM" = "win" ]; then
    if [ ! -f "$SCRIPT_DIR/$PROTOC" ]; then
        echo "错误: 找不到 protoc 工具: $SCRIPT_DIR/$PROTOC"
        echo "请确保 Windows 工具文件存在于 tool/win/ 目录"
        exit 1
    fi
    if [ ! -f "$SCRIPT_DIR/$PROTOC_GEN_GO" ]; then
        echo "错误: 找不到 protoc-gen-go 工具: $SCRIPT_DIR/$PROTOC_GEN_GO"
        exit 1
    fi
else
    if ! command -v "$PROTOC" >/dev/null 2>&1; then
        echo "错误: 找不到 protoc 工具"
        echo "Linux 环境请确保 protoc 已安装并在 PATH 中"
        exit 1
    fi
    if ! command -v "$PROTOC_GEN_GO" >/dev/null 2>&1; then
        echo "错误: 找不到 protoc-gen-go 工具"
        echo "请运行: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
        exit 1
    fi
fi

# 设置输出目录（相对于 bin/proto 目录）
PROTO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
OUTPUT_DIR="$PROTO_ROOT/../../src/proto"

# 创建输出目录（如果不存在）
if [ ! -d "$OUTPUT_DIR" ]; then
    echo "创建输出目录: $OUTPUT_DIR"
    mkdir -p "$OUTPUT_DIR"
fi

# 获取输出目录的绝对路径
OUTPUT_DIR_ABS="$(cd "$OUTPUT_DIR" && pwd)"

echo "Proto 根目录: $PROTO_ROOT"
echo "输出目录: $OUTPUT_DIR_ABS"
echo

# 切换到 proto 根目录
cd "$PROTO_ROOT"

# ============================================
# 生成共享协议文件
# ============================================
echo "[1/6] 生成共享协议文件..."
"$SCRIPT_DIR/$PROTOC" --proto_path=./share --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" share/Base.proto || {
    echo "错误: Base.proto 生成失败"
    exit 1
}

# ============================================
# 生成客户端协议文件
# ============================================
echo "[2/6] 生成客户端协议文件..."
"$SCRIPT_DIR/$PROTOC" --proto_path=./client --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" client/Struct.proto || {
    echo "错误: Struct.proto 生成失败"
    exit 1
}

"$SCRIPT_DIR/$PROTOC" --proto_path=./client --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" client/Enum.proto || {
    echo "错误: Enum.proto 生成失败"
    exit 1
}

"$SCRIPT_DIR/$PROTOC" --proto_path=./client --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" client/PreLobby.proto || {
    echo "错误: PreLobby.proto 生成失败"
    exit 1
}

"$SCRIPT_DIR/$PROTOC" --proto_path=./client --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" client/Lobby.proto || {
    echo "错误: Lobby.proto 生成失败"
    exit 1
}

"$SCRIPT_DIR/$PROTOC" --proto_path=./client --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" client/Chat.proto || {
    echo "错误: Chat.proto 生成失败"
    exit 1
}

# ============================================
# 生成服务端协议文件
# ============================================
echo "[3/6] 生成服务端协议文件..."
"$SCRIPT_DIR/$PROTOC" --proto_path=./server --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" server/Level.proto || {
    echo "错误: Level.proto 生成失败"
    exit 1
}

# ============================================
# 生成集群协议文件
# ============================================
echo "[4/6] 生成集群协议文件..."
"$SCRIPT_DIR/$PROTOC" --proto_path=./cluster --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" cluster/CFriend.proto || {
    echo "错误: CFriend.proto 生成失败"
    exit 1
}

# ============================================
# 生成内部协议文件
# ============================================
echo "[5/6] 生成内部协议文件..."
"$SCRIPT_DIR/$PROTOC" --proto_path=./client --proto_path=./internal --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" --go_out="$OUTPUT_DIR_ABS" internal/DBData.proto || {
    echo "错误: DBData.proto 生成失败"
    exit 1
}

# ============================================
# 移动文件到正确的目录结构
# ============================================
echo "[6/6] 整理文件目录结构..."
# 新的 go_package 路径: github.com/zhanglifan/leaf_server/src/proto/XXX
PROTO_SRC_DIR="$OUTPUT_DIR_ABS/github.com/zhanglifan/leaf_server/src/proto"

if [ -d "$PROTO_SRC_DIR" ]; then
    echo "移动文件从嵌套目录到直接目录..."
    cd "$PROTO_SRC_DIR"
    
    # 获取所有子目录
    for dir in */; do
        if [ -d "$dir" ]; then
            dirName="${dir%/}"
            srcPath="$PROTO_SRC_DIR/$dirName"
            dstPath="$OUTPUT_DIR_ABS/$dirName"
            
            if [ ! -d "$dstPath" ]; then
                # 目标目录不存在，直接移动
                echo "正在移动: $dirName"
                if mv "$srcPath" "$dstPath" 2>/dev/null; then
                    echo "已移动: $dirName"
                else
                    echo "警告: 移动 $dirName 失败，尝试复制..."
                    if cp -r "$srcPath" "$dstPath" 2>/dev/null && [ -d "$dstPath" ]; then
                        echo "已复制: $dirName"
                        rm -rf "$srcPath" 2>/dev/null
                    else
                        echo "错误: 无法移动或复制 $dirName"
                    fi
                fi
            else
                # 目标目录已存在，先删除旧目录再移动新目录（更新文件）
                echo "更新: $dirName (目标目录已存在，先删除旧文件)"
                if rm -rf "$dstPath" 2>/dev/null; then
                    if mv "$srcPath" "$dstPath" 2>/dev/null; then
                        echo "已更新: $dirName"
                    else
                        echo "警告: 移动 $dirName 失败，尝试复制..."
                        if cp -r "$srcPath" "$dstPath" 2>/dev/null && [ -d "$dstPath" ]; then
                            echo "已复制更新: $dirName"
                            rm -rf "$srcPath" 2>/dev/null
                        else
                            echo "错误: 无法更新 $dirName"
                        fi
                    fi
                else
                    echo "警告: 无法删除旧目录 $dirName，尝试直接复制覆盖..."
                    if cp -r "$srcPath"/* "$dstPath" 2>/dev/null; then
                        echo "已复制覆盖: $dirName"
                        rm -rf "$srcPath" 2>/dev/null
                    else
                        echo "错误: 无法更新 $dirName"
                    fi
                fi
            fi
        fi
    done
    
    cd "$OUTPUT_DIR_ABS"
    echo "删除临时嵌套目录..."
    if [ -d "$OUTPUT_DIR_ABS/github.com" ]; then
        rm -rf "$OUTPUT_DIR_ABS/github.com" 2>/dev/null
        echo "临时目录已删除"
    fi
    echo "文件目录整理完成！"
else
    echo "未找到嵌套目录，跳过整理步骤"
fi

# ============================================
# 完成
# ============================================
echo
echo "============================================"
echo "所有 Protocol Buffers 代码已生成到: $OUTPUT_DIR_ABS"
echo "============================================"
echo
