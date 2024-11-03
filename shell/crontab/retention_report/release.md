#创建t_user_report_monthly月留存统计报表数据表
CREATE TABLE t_user_report_monthly (
    id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    stat_month  INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '统计月份',
    os VARCHAR(128) NOT NULL COMMENT '设备类型',
    user_count INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户总数',
    new_users INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '新增用户量',
    retained_users INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '次月留存',
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    PRIMARY KEY (id)
)