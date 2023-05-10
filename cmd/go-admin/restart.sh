#!/bin/bash
# 进程名称
app_name="go-admin"

##服务PID
service_PID=$(ps -ef | grep ./$app_name | grep -v "systemd" | grep -v grep | grep -v "sd-pam" | awk '{print $2}')
##kill服务
if [ -n "$service_PID" ]; then
  kill "${service_PID}"
  for i in $(seq 1 300); do
    echo "$(date -d today +'%Y-%m-%d %H:%M:%S')::::: wait $i ，sleep 100mS"
    service_PID=$(ps -ef | grep ./$app_name | grep -v "systemd" | grep -v grep | grep -v "sd-pam" | awk '{print $2}')
    if [ -z "$service_PID" ]; then
      break
    fi
    kill "${service_PID}"
    sleep 0.1
  done
fi
if [ -n "$service_PID" ]; then
  echo "程序无法kill，需要手动停止"
  return
fi
nohup ./$app_name >logs/$app_name.nohup 2>&1 &
