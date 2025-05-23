CREATE TABLE `t_user_device_retention`
(
    `id`              bigint UNSIGNED                                              NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `date`            int UNSIGNED                                                 NOT NULL COMMENT '数据日期, 20230101',
    `device`          varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '设备类型',
    `new`             int UNSIGNED                                                 NOT NULL COMMENT '设备类型新增用户',
    `retained`        int UNSIGNED                                                 NOT NULL COMMENT '设备类型次日留存',
    `day7_retained`  int  UNSIGNED                                                 NOT NULL COMMENT '7日留存',
    `day15_retained` int  UNSIGNED                                                 NOT NULL COMMENT '15日留存',
    `created_at`      timestamp                                                    NULL DEFAULT NULL COMMENT '记录创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `uiq_k` (`date`, `device`) USING BTREE
);

create database speed_status;
use speed_status;
set names utf8;

CREATE TABLE IF NOT EXISTS `t_user_node`
(
    `id`                   bigint          NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `user_id`              bigint unsigned NOT NULL COMMENT '用户uid',
    `email`                varchar(64)     NOT NULL COMMENT '用户邮箱',
    `ip`                   varchar(64)     NOT NULL COMMENT '节点IP',
    `v2ray_uuid`           varchar(128)    NOT NULL COMMENT 'uuid',
    `status`               tinyint(1) DEFAULT 0 COMMENT '状态：0-未写入节点配置；1-已经写入到节点配置',
    `created_at`           timestamp       NULL     DEFAULT NULL COMMENT '创建时间',
    `updated_at`           timestamp       NULL     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uiq_k1` (`email`,`ip`,`v2ray_uuid`),
    KEY `ind_1` (`ip`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT ='用户节点配置状态管理表';

CREATE TABLE IF NOT EXISTS `t_task`
(
    `id`                   bigint          NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `ip`                   varchar(64)     NOT NULL COMMENT '节点IP',
    `date`                 int unsigned    NOT NULL COMMENT '任务日期, 20230101',
    `user_cnt`             int unsigned    NOT NULL COMMENT '用户数量',
    `status`               tinyint  unsigned DEFAULT 0 COMMENT '状态：0-初始状态；1-完成',
    `type`                 varchar(64)     NOT NULL COMMENT '任务类型',
    `created_at`           timestamp       NULL     DEFAULT NULL COMMENT '创建时间',
    `updated_at`           timestamp       NULL     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uiq_k1` (`ip`,`date`,`type`),
    KEY `ind_1` (`date`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT ='任务表';
