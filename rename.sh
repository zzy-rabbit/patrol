#!/usr/bin/env bash

# 遇到错误退出（可选）
set -e

# 遍历所有 .go1 文件
find . -type f -name "*.go1" | while read -r file; do
    # 去掉扩展名
    base="${file%.go1}"

    # 如果 .go 存在，删除
    if [ -f "${base}.go" ]; then
        echo "删除 ${base}.go"
        rm -f "${base}.go"
    fi

    # 重命名
    echo "重命名 ${file} -> ${base}.go"
    mv -f "${file}" "${base}.go"
done

echo "完成"
