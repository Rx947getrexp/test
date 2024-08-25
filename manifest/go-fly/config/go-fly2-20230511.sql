/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : go-fly2

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 11/05/2023 14:52:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for about
-- ----------------------------
DROP TABLE IF EXISTS `about`;
CREATE TABLE `about`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title_cn` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `title_en` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `keywords_cn` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `keywords_en` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `desc_cn` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `desc_en` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `css_js` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `html_cn` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `html_en` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `page` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `page`(`page` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of about
-- ----------------------------

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `conf_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `conf_key` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `conf_value` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `conf_key`(`conf_key` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES (1, '公告信息', 'AllNotice', '开源智能客服系统为您服务');
INSERT INTO `config` VALUES (2, '离线消息', 'OfflineMessage', '我现在离线，稍后回复您！');
INSERT INTO `config` VALUES (3, '欢迎消息', 'WelcomeMessage', '有什么可以帮您？');
INSERT INTO `config` VALUES (4, '邮箱地址(SMTP地址)', 'NoticeEmailSmtp', '');
INSERT INTO `config` VALUES (5, '邮箱账户', 'NoticeEmailAddress', '');
INSERT INTO `config` VALUES (6, '邮箱密码(SMTP密码)', 'NoticeEmailPassword', '');

-- ----------------------------
-- Table structure for ipblack
-- ----------------------------
DROP TABLE IF EXISTS `ipblack`;
CREATE TABLE `ipblack`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `kefu_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ip`(`ip` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ipblack
-- ----------------------------

-- ----------------------------
-- Table structure for land_page
-- ----------------------------
DROP TABLE IF EXISTS `land_page`;
CREATE TABLE `land_page`  (
  `id` int NOT NULL,
  `title` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `keyword` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `language` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `page_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of land_page
-- ----------------------------

-- ----------------------------
-- Table structure for language
-- ----------------------------
DROP TABLE IF EXISTS `language`;
CREATE TABLE `language`  (
  `id` int NOT NULL,
  `country` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `short_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of language
-- ----------------------------
INSERT INTO `language` VALUES (1, '中文简体', 'zh-cn');
INSERT INTO `language` VALUES (2, '正體中文', 'zh-tw');
INSERT INTO `language` VALUES (3, 'English', 'en_us');
INSERT INTO `language` VALUES (4, '日本語', 'ja_jp');

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `kefu_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `visitor_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `mes_type` enum('kefu','visitor') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'visitor',
  `status` enum('read','unread') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'unread',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `kefu_id`(`kefu_id` ASC) USING BTREE,
  INDEX `visitor_id`(`visitor_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '1', '2023-05-11 10:38:05', '2023-05-11 10:38:05', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (2, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '我现在离线，稍后回复您！', '2023-05-11 10:38:06', '2023-05-11 10:38:06', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (3, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '666', '2023-05-11 10:42:11', '2023-05-11 10:42:11', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (4, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '你好，请问有什么可以帮助您', '2023-05-11 10:42:28', '2023-05-11 10:42:28', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (5, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', 'face[哈哈]', '2023-05-11 10:44:03', '2023-05-11 10:44:03', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (6, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', '你好，在吗', '2023-05-11 10:47:24', '2023-05-11 10:47:24', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (7, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', '？？', '2023-05-11 10:47:33', '2023-05-11 10:47:33', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (8, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', 'attachment[{\"name\":\"-5797289432884293611_109.jpg\",\"ext\":\".jpg\",\"size\":3871,\"path\":\"/static/upload/2023May/e28e832d182d1cf939e353af5e2d66b7.jpg\"}]', '2023-05-11 10:49:04', '2023-05-11 10:49:04', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (9, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', 'img[/static/upload/2023May/a7626c861e7527ffe14b32e6302e402a.png]', '2023-05-11 10:49:35', '2023-05-11 10:49:35', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (10, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '99', '2023-05-11 10:52:01', '2023-05-11 10:52:01', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (11, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', '666', '2023-05-11 10:54:35', '2023-05-11 10:54:35', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (12, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', 'img[/static/upload/2023May/e616e09a80bebbc61cb3068dbfde88f9.jpg]', '2023-05-11 10:58:25', '2023-05-11 10:58:25', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (13, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', 'face[鼓掌]', '2023-05-11 10:59:33', '2023-05-11 10:59:33', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (14, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '123', '2023-05-11 11:07:08', '2023-05-11 11:07:08', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (15, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '689', '2023-05-11 11:07:24', '2023-05-11 11:07:24', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (16, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '555', '2023-05-11 11:07:51', '2023-05-11 11:07:51', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (17, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '123', '2023-05-11 11:08:00', '2023-05-11 11:08:00', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (18, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '889', '2023-05-11 11:08:13', '2023-05-11 11:08:13', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (19, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '999', '2023-05-11 11:33:36', '2023-05-11 11:33:36', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (20, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', '666', '2023-05-11 11:44:52', '2023-05-11 11:44:52', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (21, 'kefu2', '7d27761c-2470-4679-897d-5fddfb51945c', 'img[/static/upload/2023May/adf6221db45e70d2a78187e711bc34a2.jpg]', '2023-05-11 14:29:24', '2023-05-11 14:29:24', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (22, 'kefu2', '432c2270-0844-4b10-b7eb-2536f3496566', '还在么', '2023-05-11 14:30:52', '2023-05-11 14:30:52', NULL, 'kefu', 'unread');

-- ----------------------------
-- Table structure for reply_group
-- ----------------------------
DROP TABLE IF EXISTS `reply_group`;
CREATE TABLE `reply_group`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `user_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of reply_group
-- ----------------------------
INSERT INTO `reply_group` VALUES (1, '常见问题', 'kefu2');

-- ----------------------------
-- Table structure for reply_item
-- ----------------------------
DROP TABLE IF EXISTS `reply_item`;
CREATE TABLE `reply_item`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `group_id` int NOT NULL DEFAULT 0,
  `user_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `item_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `group_id`(`group_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of reply_item
-- ----------------------------
INSERT INTO `reply_item` VALUES (1, '在这里[官网]link[https://gofly.sopans.com]!', 1, 'kefu2', '官方地址在哪?');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `method` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `path` varchar(2048) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '普通客服', '*', '*');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `password` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `avator` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_name`(`name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'kefu2', '202cb962ac59075b964b07152d234b70', '智能客服系统', '2020-06-27 19:32:41', '2020-07-04 09:32:20', NULL, '/static/images/4.jpg');

-- ----------------------------
-- Table structure for user_client
-- ----------------------------
DROP TABLE IF EXISTS `user_client`;
CREATE TABLE `user_client`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `kefu` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `client_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user`(`kefu` ASC, `client_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_client
-- ----------------------------

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL DEFAULT 0,
  `role_id` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for visitor
-- ----------------------------
DROP TABLE IF EXISTS `visitor`;
CREATE TABLE `visitor`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `avator` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `source_ip` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `to_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `visitor_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT 0,
  `refer` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `city` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `client_ip` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `extra` varchar(2048) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `visitor_id`(`visitor_id` ASC) USING BTREE,
  INDEX `to_id`(`to_id` ASC) USING BTREE,
  INDEX `idx_update`(`updated_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visitor
-- ----------------------------
INSERT INTO `visitor` VALUES (1, '匿名网友', '/static/images/2.png', '::1', 'kefu2', '2023-05-11 10:37:42', '2023-05-11 14:29:06', NULL, '7d27761c-2470-4679-897d-5fddfb51945c', 1, '独立私有化智能在线客服系统', '未识别地区', '::1', 'eyJyZWZlciI6IuaXoCIsImhvc3QiOiJodHRwOi8vbG9jYWxob3N0OjgwODEvIn0');
INSERT INTO `visitor` VALUES (2, '局域网局域网网友', '/static/images/1.png', '10.10.10.6', 'kefu2', '2023-05-11 10:46:41', '2023-05-11 11:21:18', NULL, '432c2270-0844-4b10-b7eb-2536f3496566', 0, 'http://10.10.10.222:8081/', '局域网局域网', '10.10.10.6', '');

-- ----------------------------
-- Table structure for welcome
-- ----------------------------
DROP TABLE IF EXISTS `welcome`;
CREATE TABLE `welcome`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `keyword` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `content` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `is_default` tinyint UNSIGNED NOT NULL DEFAULT 0,
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `keyword`(`keyword` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of welcome
-- ----------------------------
INSERT INTO `welcome` VALUES (1, 'kefu2', 'offline', '我暂时离线，留言已转发到我的邮箱，稍后回复~', 1, '2020-08-24 02:57:49');
INSERT INTO `welcome` VALUES (2, 'kefu2', 'welcome', '请问有什么可以帮您？', 0, '2020-08-24 02:57:49');

SET FOREIGN_KEY_CHECKS = 1;
