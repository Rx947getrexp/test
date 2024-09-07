# 紧急恢复
## 背景
当应用服务器挂掉后，如何恢复业务。 

### 现状：

- 前端连接会连接以下四台机器（后端通过放出的应用服务器域名来控制）
    - nginx 1
    - nginx 2
    - 应用服务器主（部署了 redis 和 mysql）
    - 应用服务器备（连接应用服务器主机器上的 redis 和 mysql）

### 故障场景：
#### 场景1：nginx机器挂
1. 要确认应用服务器主和备机正常；
2. 把nginx机器从域名列表下掉（通过管理后台可以操作）。
3. 再上架新的nginx机器顶上；



#### 场景2：应用服务器备挂
1. 确认nginx和应用服务器主节点正常；
2. 将应用服务器备节点的域名下掉（通过管理后台可以操作）；
3. 把nginx上转发列表剔掉挂掉的机器；
4. 要赶紧找一台新的机器顶上，防止应用服务器主挂掉后没有备份；
5. 修改主节点上 database_backup_speed.py 脚本，保证数据库备份任务正常；

#### 场景3. 应用服务器主挂

1. 确认nginx和应用服务器备节点正常；
2. 在备节点上恢复数据
   - 在备机的 /shell/sql_backup 路径下找到最新的备份sql文件
   - 将备份数据导入本地 mysql
    ```sql
    mysql < /shell/sql_backup/speed-找到最新的文件.sql
    ```
   - 在备机的mysql上创建缺失的表
```sql
    CREATE TABLE `user_logs` (
`id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
`user_id` bigint NOT NULL COMMENT '用户id',
`datestr` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期',
`ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'IP地址',
`user_agent` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求头user-agent',
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
`comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `log_user_date` (`user_id`,`datestr`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2217919 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日志表（仅记录第一次事件)
```

3. 将应用服务器主节点的域名下掉（sql操作）；
```sql
begin;
update speed.t_app_dns set status=2 where ip='挂掉机器的IP';

-- 确认没有问题后再执行commit
commit;
```
4. 修改备机上go-api的配置文件 /wwwroot/go/go-api/config.yaml
   - redis 连接地址改为本机 `localhost`
   - mysql 连接地址改为本机 `localhost`
5. 重启 go-api
6. 把nginx上转发列表剔掉挂掉的机器；
7. 严重服务是否已经恢复;
8. 要赶紧找一台新的机器顶上，防止应用服务器主挂掉后没有备份；
9. 修改主节点上 database_backup_speed.py 脚本，保证数据库备份任务正常；
10. 稍后快速把其他的几个服务恢复：
    - go-upload
    - go-fly
    - go-admin
    - go-job

#### 场景4. 两台应用服务器同时挂
1. 赶紧买一台新机器搭建起来，这个步骤就是一个全新的搭建步骤了。
   - 数据可以从上报服务器上去找到最新的备份sql文件（/shell/sql_backup 路径下找到最新的备份sql文件）

#### 场景5. 两台应用服务器同时挂且上报服务器也挂了
@@@@@ 那就完蛋了，数据都都丢了，等于要从头再来了！

# 备份文件
- 备份的机器列表：
  - 185.22.152.47 （备机）
  - 185.22.154.21 （上报机器）
  - 45.251.243.140 (香港机器)
- 备份路径：
  - 安装文件备份：/root/workdir/
  - SQL文件备份：/shell/sql_backup/


