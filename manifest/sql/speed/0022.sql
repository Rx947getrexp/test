CREATE TABLE `t_user_report_monthly` (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    stat_month INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '统计月份，格式YYYYMM',
    os VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '设备类型，如Android, iPhone等',
    user_count INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户总数',
    new_users INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '新增用户量',
    retained_users INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '次月留存用户数',
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户月度留存报表';