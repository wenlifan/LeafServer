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
# 输出到项目根目录，配合 module 选项使用
PROJECT_ROOT="$PROTO_ROOT/../.."
OUTPUT_DIR="$PROJECT_ROOT/src/proto"

# 创建输出目录（如果不存在）
if [ ! -d "$OUTPUT_DIR" ]; then
    echo "创建输出目录: $OUTPUT_DIR"
    mkdir -p "$OUTPUT_DIR"
fi

# 获取输出目录和项目根目录的绝对路径
OUTPUT_DIR_ABS="$(cd "$OUTPUT_DIR" && pwd)"
PROJECT_ROOT_ABS="$(cd "$PROJECT_ROOT" && pwd)"

echo "Proto 根目录: $PROTO_ROOT"
echo "项目根目录: $PROJECT_ROOT_ABS"
echo "输出目录: $OUTPUT_DIR_ABS"
echo

# 切换到 proto 根目录
cd "$PROTO_ROOT"

# ============================================
# 定义生成函数：扫描目录并生成所有 .proto 文件
# ============================================
generate_proto_files() {
    local dir_name=$1
    local step_num=$2
    local step_name=$3
    local proto_paths=$4  # 额外的 proto_path，用空格分隔
    
    if [ ! -d "$dir_name" ]; then
        echo "[$step_num] 跳过 $step_name (目录不存在: $dir_name)"
        return 0
    fi
    
    # 查找目录下所有 .proto 文件（兼容 Windows 和 Linux）
    local proto_files=()
    # 使用通配符展开，兼容所有平台
    for file in "$dir_name"/*.proto; do
        # 检查文件是否存在（避免通配符未匹配时的情况）
        [ -f "$file" ] && proto_files+=("$file")
    done
    
    # 对文件列表进行排序
    if [ ${#proto_files[@]} -gt 0 ]; then
        IFS=$'\n' proto_files=($(printf '%s\n' "${proto_files[@]}" | sort))
    fi
    
    if [ ${#proto_files[@]} -eq 0 ]; then
        echo "[$step_num] 跳过 $step_name (目录中没有 .proto 文件: $dir_name)"
        return 0
    fi
    
    local file_count=${#proto_files[@]}
    echo "[$step_num] 生成 $step_name (找到 $file_count 个文件)..."
    
    # 构建 proto_path 参数（使用数组，避免字符串展开问题）
    # protoc 的导入解析规则：
    # - 导入语句如 import "Enum.proto" 会在 proto_path 中查找
    # - 如果 proto_path 是 ./client，它会查找 ./client/Enum.proto
    # - 因此，对于同一目录下的文件，proto_path 应该设置为该目录
    local proto_path_args=()
    # 首先添加当前目录的 proto_path（包含该目录下的文件）
    proto_path_args+=("--proto_path=./$dir_name")
    # 添加 proto 根目录，用于跨目录导入（如 internal 导入 client）
    proto_path_args+=("--proto_path=.")
    # 添加额外的 proto_path（用于依赖的其他目录）
    if [ -n "$proto_paths" ]; then
        for path in $proto_paths; do
            proto_path_args+=("--proto_path=./$path")
        done
    fi
    
    # 遍历每个 .proto 文件并生成
    local success_count=0
    local fail_count=0
    
    for proto_file in "${proto_files[@]}"; do
        local file_name=$(basename "$proto_file")
        echo "  正在生成: $file_name"
        
        # 使用 module 选项，从 go_package 中移除模块路径
        # go_package: github.com/zhanglifan/leaf_server/src/proto/PreLobby
        # module: github.com/zhanglifan/leaf_server
        # 结果: src/proto/PreLobby (相对于项目根目录)
        # 输出目录设置为项目根目录，这样文件会生成到 src/proto/XXX
        # 构建完整的 protoc 命令
        # 注意：proto_path 的顺序很重要，当前目录应该放在最前面
        # 文件路径使用相对于 proto 根目录的路径
        # 使用数组展开 "${proto_path_args[@]}" 来正确传递多个参数
        if "$SCRIPT_DIR/$PROTOC" "${proto_path_args[@]}" \
            --plugin=protoc-gen-go="$SCRIPT_DIR/$PROTOC_GEN_GO" \
            --go_out="$PROJECT_ROOT_ABS" \
            --go_opt=module=github.com/zhanglifan/leaf_server \
            "$proto_file" 2>&1; then
            success_count=$((success_count + 1))
        else
            echo "  错误: $file_name 生成失败"
            fail_count=$((fail_count + 1))
        fi
    done
    
    if [ $fail_count -gt 0 ]; then
        echo "错误: $step_name 中有 $fail_count 个文件生成失败"
        exit 1
    fi
    
    echo "  完成: 成功生成 $success_count 个文件"
}

# ============================================
# 按顺序生成各目录的协议文件
# ============================================

# 步骤1: 生成共享协议文件（基础协议，其他协议可能依赖）
generate_proto_files "share" "1/5" "共享协议文件" ""

# 步骤2: 生成客户端协议文件
generate_proto_files "client" "2/5" "客户端协议文件" ""

# 步骤3: 生成服务端协议文件（如果存在）
generate_proto_files "server" "3/5" "服务端协议文件" ""

# 步骤4: 生成集群协议文件
generate_proto_files "cluster" "4/5" "集群协议文件" ""

# 步骤5: 生成内部协议文件（需要 client 目录的 proto_path）
generate_proto_files "internal" "5/5" "内部协议文件" "client"

# ============================================
# 清理可能存在的嵌套目录（如果使用了 module 选项仍生成了嵌套目录）
# ============================================
echo "[最后一步] 检查并整理文件目录结构..."
PROTO_SRC_DIR="$OUTPUT_DIR_ABS/github.com/zhanglifan/leaf_server/src/proto"

if [ -d "$PROTO_SRC_DIR" ]; then
    echo "检测到嵌套目录，正在移动文件到正确位置..."
    cd "$PROTO_SRC_DIR"
    
    # 获取所有子目录
    for dir in */; do
        if [ -d "$dir" ]; then
            dirName="${dir%/}"
            srcPath="$PROTO_SRC_DIR/$dirName"
            dstPath="$OUTPUT_DIR_ABS/$dirName"
            
            if [ ! -d "$dstPath" ]; then
                # 目标目录不存在，直接移动
                echo "  正在移动: $dirName"
                if mv "$srcPath" "$dstPath" 2>/dev/null; then
                    echo "  已移动: $dirName"
                else
                    echo "  警告: 移动 $dirName 失败，尝试复制..."
                    if cp -r "$srcPath" "$dstPath" 2>/dev/null && [ -d "$dstPath" ]; then
                        echo "  已复制: $dirName"
                        rm -rf "$srcPath" 2>/dev/null
                    else
                        echo "  错误: 无法移动或复制 $dirName"
                    fi
                fi
            else
                # 目标目录已存在，先删除旧目录再移动新目录（更新文件）
                echo "  更新: $dirName (目标目录已存在，先删除旧文件)"
                if rm -rf "$dstPath" 2>/dev/null; then
                    if mv "$srcPath" "$dstPath" 2>/dev/null; then
                        echo "  已更新: $dirName"
                    else
                        echo "  警告: 移动 $dirName 失败，尝试复制..."
                        if cp -r "$srcPath" "$dstPath" 2>/dev/null && [ -d "$dstPath" ]; then
                            echo "  已复制更新: $dirName"
                            rm -rf "$srcPath" 2>/dev/null
                        else
                            echo "  错误: 无法更新 $dirName"
                        fi
                    fi
                else
                    echo "  警告: 无法删除旧目录 $dirName，尝试直接复制覆盖..."
                    if cp -r "$srcPath"/* "$dstPath" 2>/dev/null; then
                        echo "  已复制覆盖: $dirName"
                        rm -rf "$srcPath" 2>/dev/null
                    else
                        echo "  错误: 无法更新 $dirName"
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
    echo "文件已生成在正确位置，无需整理"
fi

# ============================================
# 完成
# ============================================
echo
echo "============================================"
echo "所有 Protocol Buffers 代码已生成到: $OUTPUT_DIR_ABS"
echo "============================================"
echo
