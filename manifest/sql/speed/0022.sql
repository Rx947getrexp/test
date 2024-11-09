ALTER TABLE `speed`.`t_user`
    ADD UNIQUE KEY `uniq_uuid` (`v2ray_uuid`) USING BTREE,
    ADD UNIQUE KEY `uniq_email` (`email`) USING BTREE,
    ADD KEY `ind_cid` (`client_id`) USING BTREE;

ALTER TABLE `speed`.`t_user_cancelled`
    ADD COLUMN `last_login_ip` varchar(32) DEFAULT NULL COMMENT '最近一次登录的ip',
    ADD COLUMN `last_login_country` varchar(32) DEFAULT NULL COMMENT '最近一次登录的国家',
    ADD COLUMN `preferred_country` varchar(64) DEFAULT '' COMMENT '用户选择的国家（国家名称）',
    ADD COLUMN `version` int(11) DEFAULT '0' COMMENT '数据版本号',
    ADD KEY `ind_cid` (`client_id`) USING BTREE;

ALTER TABLE `speed`.`t_user_device`  ADD KEY `ind_ut_t` (`updated_at`) USING BTREE;
    ADD COLUMN `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    ADD KEY `ind_ut` (`updated_at`) USING BTREE;
