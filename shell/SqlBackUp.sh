#!/bin/bash

# 定义本地和远程备份目录
LOCAL_BACKUP_DIR="/shell/sql_backup"
REMOTE_BACKUP_DIR="root@45.251.243.140:/shell/ru_sqlback"

# 备份数据库到本地目录
/usr/bin/mysqldump -u root -p"$MYSQL_PASSWORD" speed > "$LOCAL_BACKUP_DIR/speed.sql"
/usr/bin/mysqldump -u root -p"$MYSQL_PASSWORD" go_fly2 > "$LOCAL_BACKUP_DIR/go_fly2.sql"
/usr/bin/mysqldump -u root -p"$MYSQL_PASSWORD" speed_report > "$LOCAL_BACKUP_DIR/speed_report.sql"

# 同步备份文件到远程服务器
rsync -avz -e "ssh -i /root/.ssh/id_rsa" "$LOCAL_BACKUP_DIR/" "$REMOTE_BACKUP_DIR"

