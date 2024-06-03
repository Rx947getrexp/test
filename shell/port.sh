#!/bin/bash

# 定义要检查和开放的端口数组
PORTS=(9100 80 443 15003 10085 "13001-13005")

# 指定使用的协议，通常是tcp或udp
PROTOCOL="tcp"

# 定义日志文件路径
LOGFILE="/shell/log/port_opening.log"

# 确保日志文件的目录存在
mkdir -p $(dirname $LOGFILE)

# 记录执行日期和时间
echo "执行日期: $(date '+%Y-%m-%d %H:%M:%S')" >> $LOGFILE

for PORT in "${PORTS[@]}"
do
    if [[ $PORT == *-* ]]; then
        # 处理端口范围
        IFS=- read start end <<< "$PORT"
        for ((PORT=$start; PORT<=$end; PORT++)); do
            # 对端口范围内的每个端口进行操作
            firewall-cmd --list-all | grep "$PORT/$PROTOCOL" > /dev/null
            if [ $? -ne 0 ]; then
                echo "端口 $PORT 未开放。尝试开放端口..." | tee -a $LOGFILE
                firewall-cmd --zone=public --add-port=$PORT/$PROTOCOL --permanent >> $LOGFILE 2>&1
                firewall-cmd --reload >> $LOGFILE 2>&1
                echo "端口 $PORT 已尝试开放。" | tee -a $LOGFILE
            else
                echo "端口 $PORT 已经是开放状态，无需操作。" | tee -a $LOGFILE
            fi
        done
    else
        # 处理单个端口
        firewall-cmd --list-all | grep "$PORT/$PROTOCOL" > /dev/null
        if [ $? -ne 0 ]; then
            echo "端口 $PORT 未开放。尝试开放端口..." | tee -a $LOGFILE
            firewall-cmd --zone=public --add-port=$PORT/$PROTOCOL --permanent >> $LOGFILE 2>&1
            firewall-cmd --reload >> $LOGFILE 2>&1
            echo "端口 $PORT 已尝试开放。" | tee -a $LOGFILE
        else
            echo "端口 $PORT 已经是开放状态，无需操作。" | tee -a $LOGFILE
        fi
    fi
done
