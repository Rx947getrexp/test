DROP TABLE `t_payment_channels`;
CREATE TABLE `t_payment_channels`
(
    `id`              bigint      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name`            varchar(32) NOT NULL COMMENT '支付通道名称',
    `is_active`       tinyint(1)  NOT NULL DEFAULT '1' COMMENT '支付通道是否可用，1表示可用,2表示不可用',
    `free_trial_days` int         NOT NULL DEFAULT '3' COMMENT '赠送的免费时长（以天为单位）',
    `created_at`      timestamp   NULL     DEFAULT NULL COMMENT '创建时间',
    `updated_at`      timestamp   NULL     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_name` (`name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='支付通道表';


ALTER TABLE `t_pay_order`ADD COLUMN `payment_channel_name` varchar(32) NOT NULL COMMENT '支付通道名称';