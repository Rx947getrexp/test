CREATE TABLE `t_user_device` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
     `user_id` bigint unsigned NOT NULL COMMENT '用户uid',
     `client_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
     `os` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户端设备系统os',
     `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
     PRIMARY KEY (`id`) USING BTREE,
     UNIQUE KEY `uiq_id` (`user_id`,`client_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='客户端设备表';

INSERT INTO t_user_device (user_id, client_id, os, created_at)
SELECT u.id AS user_id, u.client_id, d.os, d.created_at
FROM t_user u
         JOIN t_dev d ON u.client_id = d.client_id
WHERE u.client_id IS NOT NULL and u.client_id != '';