# ~ 紧急恢复
# 1. 问题说明
由于业务的特殊性，无法购买大厂云服务。小厂云服务无法保证服务的稳定性，例如服务器宕机、域名被禁、IP、端口号被禁。一旦发生以上情况，可能会导致整个业务停摆，影响可能是毁灭性的。

然而，系统要打造成具备容灾、自主切换的能力，难度很大，以目前的人力很难达到，主要难点有：
- 元数据库自己搭建，一旦所在机器故障后，无自主切换能力；
- 域名、IP、证书强关联，域名绑定的IP切换后，证书需要重新生成；
- 一旦数据库发生迁移后，服务都需要修改配置并重启服务；
- 上传服务故障发生迁移后，也需要修改相关配置并重启服务；
- 执行自动切换的工具本身要具备高可靠、高可用才能实现自主切换，这个存在难度，也没有测试人员来验证；

所以目前的预案是，遇到服务故障后，人工解决。

# 2. 目标
人工恢复服务的时间要尽量短，期望在发现故障后，在30分钟内恢复。

# 3. 现状说明
## 3.1. 机器情况
目前暴露给前端的机器有四台，分别是域名的形式：
- nginx机器 `185.139.69.160`
  - thertee.xyz
  - weechat.xyz
- nginx机器 `185.22.152.9`
  - 2yiny.xyz
  - yinyong.xyz
- 主应用服务器 `31.128.41.86` (重要！！！ 部署了mysql、redis、go-upload、go-job)
  - eigrrht.xyz
  - siaax.xyz
- 备应用服务器 `185.22.152.47`
  - beiyo.xyz

## 3.2. 机器角色
1）nginx机器只做转发，将请求转发到应用服务器上。
- 配置文件参见 `go-speed/manifest/备份配置文件/nginx机器`

2）主应用服务器最为核心，上面部署了mysql、redis、go-upload、go-job服务，迁移成本最高。
- 服务列表：
  - go-admin
  - go-api
  - go-fly
  - go-job
  - go-upload
  - shell脚本：/shell
- mysql配置文件参见 `go-speed/manifest/备份配置文件/主应用服务器/mysql配置`
- redis配置文件参见 `go-speed/manifest/备份配置文件/主应用服务器/redis配置`
- crontab配置：
```shell
* * * * * /usr/bin/python3 /shell/monitor_api_redis.py &
* * * * * /usr/bin/python3 /shell/database_backup_speed.py &
* * * * * /usr/bin/python3 /shell/report/clean_report_data.py &
* * * * * /usr/bin/python3 /shell/app_port.py &
* * * * * /usr/bin/python3 /shell/monitor_ng.py &
```

3）备应用服务器，上面部署的后台服务主要是作为备机的角色。
- 服务列表：
  - go-admin
  - go-api
  - go-fly
- 说明：
  - 没有部署go-upload是因为：上传服务器对文件的保存都是在本地磁盘，目前看来只能部署一份。
  - 没有部署go-job是因为：定时器主要是下线已过期的用户、支付账单的对账，只需要部署一份。
  - 后台服务依赖元数据库都是连接的主服务器上的mysql和redis。
- crontab配置：
```shell
* * * * * /usr/bin/python3 /shell/monitor_api_redis.py &
* * * * * /usr/bin/python3 /shell/clean_old_files.py &
* * * * * /usr/bin/python3 /shell/port.py &
* * * * * /usr/bin/python3 /shell/monitor_ng.py &
```

# 4. 故障处理
故障处理需要分情况而定，挂掉的机器不同，采取的恢复方式不同，复杂度和恢复时间也不一样。

## 4.1. nginx机器故障

> 人工恢复说明：
> 1. 确认应用服务器主和备机正常；
> 2. 将故障机器从域名列表中下线（管理后台可以操作）；
> 3. 验证业务是否恢复正常；
> 4. 重新上架新的nginx机器；


## 4.2. 备应用服务器故障

> 人工恢复说明：
> 1. 确认主应用服务器正常；
> 2. 确认nginx机器正常；
> 3. 将故障机器从域名列表中下线（管理后台可以操作）；
> 4. 把nginx上转发列表剔掉挂掉的机器;
> 4. 立即找一台新的机器重新部署成新的备机；很重要！！！！ 
>    1. 可以把主应用服务器上的 `/wwwroot/`服务拷贝过来，修改配置文件的mysql、redis的连接信息； 
>    2. 参见应用服务器部署文档，初始化机器和部署redis和mysql； 
>    3. 只需要启动go-api、go-fly、go-admin三个后台服务； 
>    4. 备应用服务器上的`/shell`路径脚本见：`go-speed/manifest/备份配置文件/备应用服务器/shell`
> 5. 修改主节点上数据库备份的脚本 `database_backup_speed.py`，确保数据库继续备份；


## 4.3. 主应用服务器故障

> 人工恢复说明： 
>   - 备应用服务器上已经部署过mysql和redis！
>   - 后台服务也部署过。
>   - 优先恢复：
>     - go-admin
>     - go-api
>     - go-fly
>   - 等服务恢复后再恢复：
>     - go-job
>     - go-upload
>     - /shell各种任务

### 4.3.1. 确认备应用服务器节点正常
### 4.3.2. 确认nginx机器正常
### 4.3.3. 在备应用服务器上恢复数据库
> 1. 在备应用服务器的 `/shell/sql_backup/`路径找到最新的数据库备份文件；
> 2. 将数据导入本机mysql
>    1. ``mysql < /shell/sql_backup/speed-找到最新的文件.sql``
> 3. 因为备份文件会忽略日志类的文件，所以等数据导入后，还需要新建缺失的表：
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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日志表（仅记录第一次事件)
```

### 4.3.4. 将应用服务器主节点的域名下掉（sql操作）
```sql
begin;
update speed.t_app_dns set status=2 where ip='挂掉机器的IP';

-- 确认没有问题后再执行commit
commit;
```

### 4.3.5. 修改备应用服务器上的后台服务配置文件
- 配置文件列表：
> - `/wwwroot/go/go-api/config.yaml`
> - `/wwwroot/go/go-admin/config.yaml`
> - `/wwwroot/go/go-fly/config/mysql.json`

- 修改内容：
> - redis 连接地址改为本机 `localhost`
> - mysql 连接地址改为本机 `localhost`

### 4.3.6. 重启服务
- `go-api`
- `go-admin`
- `go-fly`

### 4.3.7. 把nginx上转发列表剔掉挂掉的机器

### 4.3.8. 验证业务是否恢复

### 4.3.9. 赶紧找一台新的机器顶上，防止应用服务器主挂掉后没有备份

### 4.3.10. shell任务、crontab配置对齐到主节点

### 4.3.11. 恢复其他服务
- go-upload
- go-job

## 4.4. 两台应用服务器同时挂
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


