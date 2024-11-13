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
    `advertiser` varchar(128) NOT NULL COMMENT '广告主，客户名称',
    `title` varchar(255) NOT NULL COMMENT '广告标题',
    `desc` text NOT NULL COMMENT '广告描述',
    `category` varchar(32) NOT NULL COMMENT '广告分类. enum: game,commerce,food,social,eg.',
    `type` varchar(32) NOT NULL COMMENT '广告类型. enum: text,image,video',
    `slot_location` varchar(255) NOT NULL COMMENT '广告位的位置',
    `sort` int NOT NULL COMMENT '在广告位置中的排序，序号越小越靠前，若序号一样就按end_time越靠后的放到前面',
    `url` varchar(255) DEFAULT NULL COMMENT '广告内容地址',
    `target_url` text DEFAULT NULL COMMENT '跳转地址，包括：pc,ios,android',
    `exposure_time` int NOT NULL COMMENT '单次曝光时间，单位秒',
    `user_level` int NOT NULL COMMENT '用户等级',
    `start_time` timestamp NULL DEFAULT NULL COMMENT '广告生效时间',
    `end_time` timestamp NULL DEFAULT NULL COMMENT '广告失效时间',
    `status` int NOT NULL COMMENT '状态:1-上架；2-下架',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告表';