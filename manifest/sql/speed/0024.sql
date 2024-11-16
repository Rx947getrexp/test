CREATE TABLE `t_ad_slot` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `location` varchar(255) NOT NULL COMMENT '广告位的位置，相当于ID',
    `name` varchar(255) NOT NULL COMMENT '广告位名称',
    `desc` varchar(512) DEFAULT NULL COMMENT '广告位描述',
    `status` int NOT NULL COMMENT '状态:1-上架；2-下架',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX uiq_l (`location`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告位置表';

CREATE TABLE `t_ad` (
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `advertiser` varchar(128) DEFAULT NULL COMMENT '广告主，客户名称',
    `name` varchar(128) NOT NULL COMMENT '广告名称',
    `type` varchar(32) DEFAULT NULL COMMENT '广告类型. enum: text,image,video',
    `url` varchar(255) DEFAULT NULL COMMENT '广告内容地址',
    `logo` varchar(255) DEFAULT NULL COMMENT 'logo',
    `slot_locations` varchar(256) DEFAULT NULL COMMENT '广告位的位置，包括权重',
    `devices` varchar(256) DEFAULT NULL COMMENT '广告位的位置，包括权重',
    `target_urls` text DEFAULT NULL COMMENT '跳转地址，包括：pc,ios,android',
    `labels` varchar(255) DEFAULT NULL COMMENT '标签',
    `exposure_time` int DEFAULT 0 COMMENT '单次曝光时间，单位秒',
    `user_levels` varchar(255) DEFAULT NULL COMMENT '用户等级',
    `start_time` timestamp NULL DEFAULT NULL COMMENT '广告生效时间',
    `end_time` timestamp NULL DEFAULT NULL COMMENT '广告失效时间',
    `status` int DEFAULT 2 COMMENT '状态:1-上架；2-下架',
    `gift_duration` int DEFAULT 0 COMMENT '赠送时间',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX uiq_name (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告表';


CREATE TABLE `t_ad_gift` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户id',
  `ad_id` bigint UNSIGNED NOT NULL COMMENT '广告ID',
  `ad_name` varchar(128) NOT NULL COMMENT '广告名称',
  `exposure_time` int DEFAULT 0 COMMENT '单次曝光时间，单位秒',
  `gift_duration` int DEFAULT 0 COMMENT '赠送时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `ind_1` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='观看广告赠送记录';


ALTER TABLE `t_user`
    ADD COLUMN kicked tinyint(1)  NOT NULL DEFAULT 0 COMMENT '被踢标记，0: 未被踢, 1: 已经被踢';

ALTER TABLE `t_user`
    ADD COLUMN `last_kicked_at` timestamp NULL DEFAULT NULL COMMENT '最近一次踢出时间';