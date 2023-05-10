#!/bin/bash
# 进程名称
app_name="go-admin"
# 代码目录
code_dir="/wwwroot/compile/go-payment/cmd/go-admin"
# git分支
branch="master"
# 当前时间
now=$(date -d today +'%Y-%m-%d-%H-%M-%S')
# 工作目录
work_dir=$(pwd)

##编译程序
cd $code_dir || exit
git reset --hard HEAD
git pull
git checkout $branch
go mod tidy -compat=1.17
go build -o $app_name
##备份旧程序
cd "$work_dir" || exit
if [ ! -d "old_app" ]; then
  mkdir -p old_app
fi
if [ ! -d "logs" ]; then
  mkdir -p logs
fi

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
mv $app_name old_app/${app_name}-"${now}"
mv logs/$app_name.nohup logs/$app_name."${now}".nohup
cp $code_dir/$app_name ./$app_name
nohup ./$app_name >logs/$app_name.nohup 2>&1 &
