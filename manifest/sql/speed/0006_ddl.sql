CREATE TABLE `t_pay_order`
(
    `id`                   bigint          NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `user_id`              bigint unsigned NOT NULL COMMENT '用户uid',
    `email`                varchar(64)     NOT NULL COMMENT '用户邮箱',
    `order_no`             varchar(64)     NOT NULL COMMENT '订单号',
    `order_amount`         varchar(64)     NOT NULL COMMENT '交易金额',
    `currency`             varchar(16)     NOT NULL COMMENT '交易币种',
    `pay_type_code`        varchar(16)     NOT NULL COMMENT '支付类型编码',
    `status`               varchar(32)     NOT NULL DEFAULT '' COMMENT '状态:1-正常；2-已软删',
    `return_status`        varchar(32)     NOT NULL DEFAULT '' COMMENT '支付平台返回的结果',
    `status_mes`           varchar(256)    NOT NULL DEFAULT '' COMMENT '状态:1-正常；2-已软删',
    `order_data`           varchar(512)    NOT NULL DEFAULT '' COMMENT '创建订单时支付平台返回的信息',
    `result_status`        varchar(32)     NOT NULL DEFAULT '' COMMENT '查询结果，实际订单状态',
    `order_reality_amount` varchar(32)     NOT NULL DEFAULT '' COMMENT '实际交易金额',
    `created_at`           timestamp       NULL     DEFAULT NULL COMMENT '创建时间',
    `updated_at`           timestamp       NULL     DEFAULT NULL COMMENT '更新时间',
    `version`              int                      DEFAULT NULL default 1 COMMENT '数据版本号',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_k1` (`order_no`) USING BTREE,
    KEY `k1` (`user_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='支付订单表';


CREATE TABLE `t_user_vip_attr_records`
(
    `id`           bigint       NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `email`        varchar(64)  NOT NULL COMMENT '用户邮箱',
    `source`       varchar(64)  NOT NULL COMMENT '来源',
    `order_no`     varchar(64)  NOT NULL COMMENT '订单号',
    `expired_time` int COMMENT '会员到期时间',
    `desc`         varchar(512) NOT NULL DEFAULT '' COMMENT '记录描述',
    `created_at`   timestamp    NULL     DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_k1` (`order_no`) USING BTREE,
    KEY `k1` (`email`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='用户会员属性变更记录';

ALTER TABLE `speed`.`t_user`
    ADD COLUMN `version` int DEFAULT NULL default 1 COMMENT '数据版本号';
