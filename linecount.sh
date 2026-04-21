#!/bin/bash

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 初始化变量
total_lines=0
file_count=0
declare -a results

# 使用说明
usage() {
    echo "用法: $0 <通配符模式>"
    echo "示例:"
    echo "  $0 '*.txt'        # 统计所有 txt 文件"
    echo "  $0 '*.sh'         # 统计所有 shell 脚本"
    echo "  $0 '*.{c,h}'      # 统计所有 c 和 h 文件"
    echo "  $0 'test*'        # 统计所有以 test 开头的文件"
    echo "  $0 'src/*.py'     # 统计 src 目录下的 py 文件"
    exit 1
}

# 检查参数
if [ $# -ne 1 ]; then
    echo "${RED}错误: 请提供通配符模式${NC}"
    usage
fi

pattern="$1"

echo "${BLUE}========================================${NC}"
echo "${BLUE}统计文件行数${NC}"
echo "${BLUE}匹配模式: $pattern${NC}"
echo "${BLUE}========================================${NC}\n"

# 使用 find 命令递归查找文件并统计行数
# 使用 while read 循环处理每个文件
find . -type f -path "./$pattern" -o -name "$pattern" 2>/dev/null | while IFS= read -r file; do
    # 去掉开头的 "./"
    file="${file#./}"
    
    # 检查文件是否为普通文本文件（跳过二进制文件）
    if file "$file" | grep -q text; then
        # 获取行数
        lines=$(wc -l < "$file" 2>/dev/null)
        
        # 检查 wc -l 是否成功
        if [ $? -eq 0 ]; then
            # 格式化输出
            printf "${GREEN}%s${NC}:${YELLOW}%d${NC}\n" "$file" "$lines"
            
            # 累加总行数（在子 shell 中需要特殊处理）
            echo "$lines" >> /tmp/total_lines_temp.$$
            echo "$file" >> /tmp/files_list_temp.$$
        else
            echo "${RED}无法读取文件: $file${NC}"
        fi
    fi
done

# 计算总行数
if [ -f /tmp/total_lines_temp.$$ ]; then
    total_lines=$(awk '{sum+=$1} END {print sum+0}' /tmp/total_lines_temp.$$)
    file_count=$(wc -l < /tmp/files_list_temp.$$ 2>/dev/null)
    
    echo "\n${BLUE}========================================${NC}"
    echo "${BLUE}统计结果:${NC}"
    echo "${GREEN}文件数量:${NC} $file_count"
    echo "${GREEN}总行数:${NC} $total_lines"
    echo "${BLUE}========================================${NC}"
    
    # 清理临时文件
    rm -f /tmp/total_lines_temp.$$ /tmp/files_list_temp.$$
else
    echo "${RED}未找到匹配的文件${NC}"
fi
