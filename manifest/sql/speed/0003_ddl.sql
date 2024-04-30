create database speed_report;
use speed_report;

CREATE TABLE `t_user_report_day`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `date`       int(10) unsigned NOT NULL COMMENT '数据日期, 20230101',
    `channel_id` int(10) unsigned NOT NULL COMMENT '渠道id',
    `total` int(10) unsigned NOT NULL COMMENT '用户总量',
    `new` int(10) unsigned NOT NULL COMMENT '新增用户',
    `retained` int(10) unsigned NOT NULL COMMENT '留存',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_k` (`date`, `channel_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日报表';
