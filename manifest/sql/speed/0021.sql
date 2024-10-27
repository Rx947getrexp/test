ALTER TABLE `speed_report`.`t_user_op_log`
    ADD COLUMN `interface_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '接口地址',
    ADD COLUMN `server_code` varchar(8) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '后端状态码',
    ADD COLUMN `http_code` varchar(8) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'HTTP状态码',
    ADD COLUMN `trace_id` varchar(48) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'TraceId',
    ADD COLUMN `user_id` bigint unsigned DEFAULT NULL COMMENT '用户uid';
