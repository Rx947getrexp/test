DROP TABLE `t_payment_channel`;
CREATE TABLE `t_payment_channel`
(
    `id`              bigint      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `channel_id`            varchar(32) NOT NULL COMMENT '支付通道ID',
    `channel_name`            varchar(32) NOT NULL COMMENT '支付通道名称',
    `is_active`       tinyint(1)  NOT NULL DEFAULT '1' COMMENT '支付通道是否可用，1表示可用,2表示不可用',
    `free_trial_days` int         NOT NULL DEFAULT '3' COMMENT '赠送的免费时长（以天为单位）',
    `timeout_duration` int         NOT NULL DEFAULT '30' COMMENT '订单未支付超时关闭时间（单位分钟）',
    `payment_qr_code`            text default NULL COMMENT '支付收款码. eg: U支付收款码',
    `bank_card_info`            text default NULL COMMENT '银行卡信息',
    `customer_service_info`            text default NULL COMMENT '客服信息',
    `mer_no`            varchar(64) default NULL COMMENT 'mer_no',
    `pay_type_code`            varchar(64) default NULL COMMENT 'pay_type_code',
    `weight` int NOT NULL DEFAULT 0 COMMENT '权重，排序使用',
    `created_at`      timestamp   NULL     DEFAULT NULL COMMENT '创建时间',
    `updated_at`      timestamp   NULL     DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_id` (`channel_id`) USING BTREE,
    UNIQUE KEY `uiq_name` (`channel_name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='支付通道表';


ALTER TABLE `t_pay_order`ADD COLUMN `payment_channel_id` varchar(32) NOT NULL COMMENT '支付通道ID';
ALTER TABLE `t_goods`ADD COLUMN `usd_pay_price` decimal(10,6) NOT NULL COMMENT 'usd_pay价格(U)';