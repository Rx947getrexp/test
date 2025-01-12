set names utf8;

CREATE TABLE IF NOT EXISTS `t_user_node`
(
    `id`                   bigint  unsigned  NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `email`                varchar(64)     NOT NULL COMMENT '用户邮箱',
    `ip`                   varchar(64)     NOT NULL COMMENT '节点IP',
    `v2ray_uuid`           varchar(128)    NOT NULL COMMENT 'uuid',
    `created_at`           timestamp       NULL     DEFAULT NULL COMMENT '创建时间',
    `user_id`              bigint unsigned NOT NULL COMMENT '用户uid',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uiq_k1` (`email`,`ip`)
    ) ENGINE = InnoDB AUTO_INCREMENT = 1 COMMENT ='用户节点配置状态管理表';