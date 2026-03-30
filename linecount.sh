#!/usr/bin/env bash

# 用法检查
if [ $# -lt 1 ]; then
  echo "Usage: $0 <pattern>"
  echo "Example: $0 \"*.go\""
  exit 1
fi

pattern="$1"

total=0

# 遍历匹配文件
while IFS= read -r -d '' file; do
  lines=$(wc -l < "$file")
  echo "$file: $lines"
  total=$((total + lines))
done < <(find . -type f -name "$pattern" -print0)

echo "----------------------"
echo "Total lines: $total"
