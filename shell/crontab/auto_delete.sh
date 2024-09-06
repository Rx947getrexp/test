#!/bin/bash

# 设置备份文件所在的目录
BACKUP_DIR="/shell/sql_backup/"

# 获取7天前的日期
CUTOFF_DATE=$(date -d "7 days ago" +%Y-%m-%d)

# 删除7天前的文件
find "$BACKUP_DIR" -type f -name "*.sql" -not -newermt "$CUTOFF_DATE" -delete

echo "已删除 $CUTOFF_DATE 之前的所有SQL备份文件"
