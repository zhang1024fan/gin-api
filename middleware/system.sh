#!/bin/bash

# 获取私有 IP（默认网卡）
privateIp=$(hostname -I | awk '{print $1}')
# 获取公网 IP（如果存在）
publicIp=$(curl -s ipinfo.io/ip || echo "")
# 获取操作系统版本
os=$(cat /etc/os-release | grep PRETTY_NAME | cut -d= -f2 | tr -d '"' | sed 's/ LTS//;s/ //g')
# 获取 CPU 核心数
cpu=$(nproc)核
# 获取内存大小（单位为 MB，转换为 G）
memory=$(free -m | awk '/^Mem:/{printf "%.0fG\n",$2/1024}')
# 获取磁盘总容量（单位为 G）
disk=$(lsblk -b | grep 'disk' | awk '{print $4}' | head -1 | awk '{printf "%.0fG\n",$1/1024/1024/1024}')
# 输出 JSON 格式
cat <<EOF
{"privateIp":"$privateIp","publicIp":"$publicIp","os":"$os","cpu":"$cpu","memory":"$memory","disk":"$disk"}
EOF
