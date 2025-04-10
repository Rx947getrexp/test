-- 推广dns域名映射表
CREATE TABLE IF NOT EXISTS `speed`.`t_promotion_dns` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `dns` varchar(64) DEFAULT NULL COMMENT '域名',
  `ip` varchar(64) DEFAULT NULL COMMENT 'ip地址',
  `mac_channel` varchar(64) DEFAULT NULL COMMENT '苹果电脑渠道',
  `win_channel` varchar(64) DEFAULT NULL COMMENT 'windows电脑渠道',
  `android_channel` varchar(64) DEFAULT NULL COMMENT '安卓渠道',
  `promoter` varchar(64) DEFAULT NULL COMMENT '推广人员',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推广dns域名映射表';