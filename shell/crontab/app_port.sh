#!/bin/bash

# 定义要检查和开放的端口数组
PORTS=(80 443 3306 6379)

# 指定使用的协议，通常是tcp或udp
PROTOCOL="tcp"

# 定义日志文件路径
LOGFILE="/shell/log/port_opening.log"

# 记录执行日期和时间
echo "执行日期: $(date '+%Y-%m-%d %H:%M:%S')" >> $LOGFILE

for PORT in "${PORTS[@]}"
do
    # 检查端口是否已经开放
    firewall-cmd --list-all | grep "$PORT/$PROTOCOL" > /dev/null

    if [ $? -ne 0 ]; then
        echo "端口 $PORT 未开放。尝试开放端口..." | tee -a $LOGFILE

        # 尝试开放端口，这里默认操作在默认区域（zone）
        firewall-cmd --zone=public --add-port=$PORT/$PROTOCOL --permanent >> $LOGFILE 2>&1

        # 重新加载firewalld的规则，使改动生效
        firewall-cmd --reload >> $LOGFILE 2>&1

        if [ $? -eq 0 ]; then
            echo "端口 $PORT 已成功开放。" | tee -a $LOGFILE
        else
            echo "尝试开放端口 $PORT 时遇到错误。请检查你的firewalld配置或运行此脚本的权限。" | tee -a $LOGFILE
        fi
    else
        echo "端口 $PORT 已经是开放状态，无需操作。" | tee -a $LOGFILE
    fi
done

