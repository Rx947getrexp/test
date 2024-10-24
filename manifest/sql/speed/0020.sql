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