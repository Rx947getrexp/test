use speed_report;
set names utf8;

CREATE TABLE `t_user_ping`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `email` varchar(64) DEFAULT NULL COMMENT '用户邮箱',
    `host` varchar(64) DEFAULT NULL COMMENT '节点host, ip or dns',
    `code` varchar(64) DEFAULT NULL COMMENT 'ping的结果',
    `cost` mediumtext DEFAULT NULL COMMENT 'ping耗时',
    `time` varchar(64) DEFAULT NULL COMMENT '上报时间',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `key1` (`email`, `created_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户ping节点结果上报数据';




ALTER TABLE `speed_report`.`t_user_online_day` ADD COLUMN `last_login_country` varchar(64) DEFAULT "" COMMENT '最后登陆的国家';
