use speed;
set names utf8;

CREATE TABLE `t_country`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name` varchar(64) DEFAULT NULL COMMENT '国家名称',
    `name_cn` varchar(64) DEFAULT NULL COMMENT '国家名称中文',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uiq_k1` (`name`) USING BTREE,
    UNIQUE KEY `uiq_k2` (`name_cn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='国家名称表';

insert into t_country set name = 'China', name_cn = '中国', created_at = now(), updated_at = now();
insert into t_country set name = 'Russia', name_cn = '俄罗斯', created_at = now(), updated_at = now();
insert into t_country set name = 'Latvia', name_cn = '拉脱维亚', created_at = now(), updated_at = now();
insert into t_country set name = 'Germany', name_cn = '德国', created_at = now(), updated_at = now();

CREATE TABLE `t_serving_country` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) NOT NULL COMMENT '国家名称，不可以修改，作为ID用',
  `display` varchar(64) NOT NULL COMMENT '用于在用户侧展示的国家名称',
  `logo_link` varchar(1024) DEFAULT '' COMMENT '国家图片地址',
  `ping_url` varchar(512) DEFAULT '' COMMENT '前端使用',
  `is_recommend` int DEFAULT NULL COMMENT '推荐节点1-是；2-否',
  `weight` int NOT NULL DEFAULT 0 COMMENT '权重',
  `status` int NOT NULL DEFAULT 1 COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k1` (`name`) USING BTREE,
  UNIQUE KEY `uiq_k2` (`display`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='上架国家表';

ALTER TABLE `speed`.`t_node` ADD COLUMN `weight` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '权重';
ALTER TABLE `speed`.`t_user` ADD COLUMN `preferred_country` varchar(64) DEFAULT "" COMMENT '用户选择的国家（国家名称）';

use speed_report;

CREATE TABLE `t_user_op_log`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `email` varchar(64) DEFAULT NULL COMMENT '用户账号',
    `device_id` varchar(64) DEFAULT NULL COMMENT '设备ID',
    `device_type` varchar(64) DEFAULT NULL COMMENT '设备类型',
    `page_name` varchar(64) DEFAULT NULL COMMENT 'page_name',
    `result` varchar(128) DEFAULT NULL COMMENT 'result',
    `content` text DEFAULT NULL COMMENT 'content',
    `create_time` varchar(64) DEFAULT NULL COMMENT '提交时间',
    `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `key1` (`email`) USING BTREE,
    KEY `key2` (`device_id`) USING BTREE,
    KEY `key4` (`created_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户操作轨迹日志';





