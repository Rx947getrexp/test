CREATE TABLE `t_doc` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
 `type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
 `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
 `desc` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
 `content` MEDIUMTEXT DEFAULT NULL COMMENT 'content',
 `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
 `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `uiq_t_n` (`type`,`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='官网文档';