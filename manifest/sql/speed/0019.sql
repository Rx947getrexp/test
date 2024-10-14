CREATE TABLE t_promotion_channels
(
    id               INT AUTO_INCREMENT PRIMARY KEY,
    promoter_name    VARCHAR(100) NOT NULL COMMENT '推广人姓名',
    promotion_domain VARCHAR(255) NOT NULL COMMENT '推广域名',
    channel          VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '推广域名对应渠道',
    created_at       TIMESTAMP             DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE KEY `uiq_promoter_domain` (`promoter_name`, `promotion_domain`, `channel`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='推广人渠道表';