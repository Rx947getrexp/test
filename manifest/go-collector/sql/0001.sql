create database speed_collector;

CREATE TABLE `t_v2ray_user_traffic`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `email`      varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
    `date`       int(10) unsigned NOT NULL COMMENT '数据日期, 20230101',
    `ip`         varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
    `uplink`     bigint unsigned NOT NULL COMMENT '上行流量',
    `downlink`   bigint unsigned NOT NULL COMMENT '下行流量',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_k` (`email`, `date`, `ip`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户流量采集记录';


CREATE TABLE `t_v2ray_user_traffic_log`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `email`      varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
    `ip`         varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
    `date_time` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '数据采集时间',
    `uplink`     bigint unsigned NOT NULL COMMENT '上行流量',
    `downlink`   bigint unsigned NOT NULL COMMENT '下行流量',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `k_1` (`email`, `ip`, `date_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户流量采集流水记录';
