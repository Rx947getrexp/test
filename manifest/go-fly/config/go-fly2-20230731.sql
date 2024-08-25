/*
 Navicat Premium Data Transfer

 Source Server         : 加速器测试库
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33-0ubuntu0.22.04.2)
 Source Host           : localhost:3306
 Source Schema         : go-fly2

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33-0ubuntu0.22.04.2)
 File Encoding         : 65001

 Date: 31/07/2023 14:34:56
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
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO `config` VALUES (1, 'Отправить уведомление по электронной почте (адрес SMTP)', 'NoticeEmailSmtp', '');
INSERT INTO `config` VALUES (2, 'Отправить уведомление по электронной почте (адрес электронной почты)', 'NoticeEmailAddress', '');
INSERT INTO `config` VALUES (3, 'Отправить уведомление по электронной почте (пароль)', 'NoticeEmailPassword', '');

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
INSERT INTO `language` VALUES (5, '俄语', 'ru');

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
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, 'kefu001', 'b82a0238-fb9b-4558-9c10-0934e247044c', '1', '2023-07-13 10:01:01', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (2, 'kefu001', 'b82a0238-fb9b-4558-9c10-0934e247044c', 'Я сейчас оффлайн, отвечу вам позже!', '2023-07-13 10:01:02', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (3, 'kefu001', '22a2f80e-1213-4133-9307-4ae45d3f0aff', '1', '2023-07-19 09:11:57', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (4, 'kefu001', '22a2f80e-1213-4133-9307-4ae45d3f0aff', 'Вы видите, что я говорю?', '2023-07-19 09:12:06', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (5, 'kefu001', '22a2f80e-1213-4133-9307-4ae45d3f0aff', 'Можно', '2023-07-19 09:12:22', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (6, 'kefu001', '22a2f80e-1213-4133-9307-4ae45d3f0aff', 'OK', '2023-07-19 09:12:27', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (7, 'kefu001', '22a2f80e-1213-4133-9307-4ae45d3f0aff', 'face[嘻嘻]', '2023-07-19 09:12:33', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (8, 'kefu001', '22a2f80e-1213-4133-9307-4ae45d3f0aff', 'face[嘻嘻]', '2023-07-19 09:12:36', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (9, 'kefu001', '37519c3a-9d97-49d7-9b30-3c2a27c002eb', '111', '2023-07-19 11:39:28', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (10, 'kefu001', '37519c3a-9d97-49d7-9b30-3c2a27c002eb', '11', '2023-07-19 11:39:33', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (11, 'kefu001', '37519c3a-9d97-49d7-9b30-3c2a27c002eb', 'В чем дело?', '2023-07-19 11:39:48', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (12, 'kefu001', '37519c3a-9d97-49d7-9b30-3c2a27c002eb', 'face[哈哈]', '2023-07-19 11:40:00', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (13, 'kefu001', '37519c3a-9d97-49d7-9b30-3c2a27c002eb', 'img[/static/upload/2023July/e9f0f709b9b3b00837703b16e919a751.png]', '2023-07-19 11:40:16', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (14, 'kefu001', '37519c3a-9d97-49d7-9b30-3c2a27c002eb', 'attachment[{\"name\":\"E64DFD2D-EAA4-4468-96F3-F0292A77F44B.png\",\"ext\":\".png\",\"size\":159546,\"path\":\"/static/upload/2023July/4a3fd5cf1ab82cbb10462ca5ff151ef9.png\"}]', '2023-07-19 11:40:30', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (15, 'kefu001', '811619bd-6ce9-43ca-9d67-b75dee8cc7d4', 'face[微笑]', '2023-07-20 02:36:05', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (16, 'kefu001', '811619bd-6ce9-43ca-9d67-b75dee8cc7d4', '666', '2023-07-20 02:36:12', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (17, 'kefu001', '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', 'face[嘻嘻]', '2023-07-20 03:44:02', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (18, 'kefu001', '853ba72a-3dbf-4adf-b2e4-c0e9d26e14bb', '1', '2023-07-20 06:32:50', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (19, 'kefu001', '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', 'face[哈哈]', '2023-07-20 12:06:50', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (20, 'kefu001', '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', 'img[/static/upload/2023July/94bedc1da9701e90478a8d9304a4acf8.png]', '2023-07-20 12:30:46', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (21, 'kefu001', '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', 'img[/static/upload/2023July/4dde603ed7566131a9a23599724092cd.png]', '2023-07-20 12:30:55', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (22, 'kefu001', '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', 'img[/static/upload/2023July/8164e75945094807e8617b66eb00f44f.png]', '2023-07-20 12:32:15', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (23, 'kefu001', '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', '666', '2023-07-21 01:51:27', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (24, 'kefu001', '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', 'В чем дело?', '2023-07-21 09:49:02', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (25, 'kefu001', '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', 'face[微笑]', '2023-07-21 09:49:19', '2023-07-23 05:37:47', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (26, 'kefu001', '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', 'img[/static/upload/2023July/a7dec441bb2e1c1ce306702b2b86c80e.png]', '2023-07-21 09:50:32', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (27, 'kefu001', '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', 'img[/static/upload/2023July/3c636d679972bdd3b332748f602dede0.png]', '2023-07-21 09:50:39', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (28, 'kefu001', '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', 'attachment[{\"name\":\"71162593640__D2EA7754-4FB1-46AF-BE1B-8A9F359DC9D3.MOV\",\"ext\":\".mov\",\"size\":340623,\"path\":\"/static/upload/2023July/b141bc72779a077f6aa16b90cb52956d.mov\"}]', '2023-07-21 09:52:26', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (29, 'kefu001', '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', '71162593640__D2EA7754-4FB1-46AF-BE1B-8A9F359DC9D3.MOV', '2023-07-21 09:54:04', '2023-07-23 05:37:47', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (30, 'kefu001', '1ccfbbea-ee1f-4173-b930-439258f2fb29', 'face[微笑]', '2023-07-24 09:22:09', '2023-07-24 09:22:09', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (31, 'kefu001', '1ccfbbea-ee1f-4173-b930-439258f2fb29', 'Я сейчас оффлайн, отвечу вам позже!', '2023-07-24 09:22:10', '2023-07-24 09:22:10', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (32, 'kefu001', '9ecd8bb8-21b5-4fcb-95aa-1af6fa2ea509', 'face[嘻嘻]', '2023-07-25 03:39:00', '2023-07-25 03:39:00', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (33, 'kefu001', '9ecd8bb8-21b5-4fcb-95aa-1af6fa2ea509', 'Я сейчас оффлайн, отвечу вам позже!', '2023-07-25 03:39:01', '2023-07-25 03:39:01', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (34, 'kefu001', '3eb0ca86-0435-4a74-b277-de3a63fd9e87', '111', '2023-07-25 07:16:43', '2023-07-25 07:16:43', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (35, 'kefu001', '3eb0ca86-0435-4a74-b277-de3a63fd9e87', 'Я сейчас оффлайн, отвечу вам позже!', '2023-07-25 07:16:44', '2023-07-25 07:16:44', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (36, 'kefu001', '3eb0ca86-0435-4a74-b277-de3a63fd9e87', '666', '2023-07-25 09:28:49', '2023-07-25 09:28:49', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (37, 'kefu001', 'b82a0238-fb9b-4558-9c10-0934e247044c', '666', '2023-07-25 11:45:28', '2023-07-25 11:45:28', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (38, 'kefu001', '8d7901d4-db77-4508-8b66-eee9126bec9c', 'face[哈哈]', '2023-07-25 11:52:55', '2023-07-25 11:52:55', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (39, 'kefu001', '8d7901d4-db77-4508-8b66-eee9126bec9c', 'face[微笑]', '2023-07-25 11:53:02', '2023-07-25 11:53:02', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (40, 'kefu001', '3eb0ca86-0435-4a74-b277-de3a63fd9e87', '666777', '2023-07-25 11:53:44', '2023-07-25 11:53:44', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (41, 'kefu001', '3eb0ca86-0435-4a74-b277-de3a63fd9e87', 'face[哈哈]', '2023-07-26 02:49:33', '2023-07-26 02:49:33', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (42, 'kefu001', 'a67a6d6b-e435-4bd5-bc70-86b821c00d21', 'face[哈哈]', '2023-07-27 02:15:26', '2023-07-27 02:15:26', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (43, 'kefu001', 'a67a6d6b-e435-4bd5-bc70-86b821c00d21', 'Я сейчас оффлайн, отвечу вам позже!', '2023-07-27 02:15:27', '2023-07-27 02:15:27', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (44, 'kefu001', '088e3251-2207-4012-8d12-7132ea104203', 'face[微笑]', '2023-07-27 08:53:16', '2023-07-27 08:53:16', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (45, 'kefu001', '088e3251-2207-4012-8d12-7132ea104203', '666', '2023-07-27 08:53:52', '2023-07-27 08:53:52', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (46, 'kefu001', '088e3251-2207-4012-8d12-7132ea104203', 'face[哈哈]', '2023-07-27 08:53:59', '2023-07-27 08:53:59', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (47, 'kefu001', '088e3251-2207-4012-8d12-7132ea104203', 'face[可爱]', '2023-07-27 08:54:05', '2023-07-27 08:54:05', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (48, 'kefu001', '3eb0ca86-0435-4a74-b277-de3a63fd9e87', '666', '2023-07-27 12:11:16', '2023-07-27 12:11:16', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (49, 'kefu001', 'f675a15e-d25e-4b1c-a831-6e1779e06b69', 'face[微笑]', '2023-07-28 07:17:41', '2023-07-28 07:17:41', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (50, 'kefu001', 'f675a15e-d25e-4b1c-a831-6e1779e06b69', 'face[酷]', '2023-07-28 07:17:58', '2023-07-28 07:17:58', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (51, 'kefu001', '3eb10275-ae99-4fd7-aa09-99c56914c0d7', 'В чем дело?', '2023-07-29 03:59:20', '2023-07-29 03:59:20', NULL, 'kefu', 'unread');
INSERT INTO `message` VALUES (52, 'kefu001', '3eb10275-ae99-4fd7-aa09-99c56914c0d7', 'face[熊猫]', '2023-07-29 03:59:33', '2023-07-29 03:59:33', NULL, 'visitor', 'unread');
INSERT INTO `message` VALUES (53, 'kefu001', '3eb10275-ae99-4fd7-aa09-99c56914c0d7', 'face[给力]', '2023-07-29 04:00:05', '2023-07-29 04:00:05', NULL, 'kefu', 'unread');

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
INSERT INTO `reply_group` VALUES (1, 'Часто задаваемые вопросы', 'kefu2');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of reply_item
-- ----------------------------

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
INSERT INTO `role` VALUES (1, 'Обычный оператор', '*', '*');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `password` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
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
INSERT INTO `user` VALUES (1, 'kefu001', 'e10adc3949ba59abbe56e057f20f883e', 'интеллектуальная онлайн система обслуживания клиентов', '2020-06-27 19:32:41', '2020-07-04 09:32:20', NULL, '/static/images/4.jpg');

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
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of visitor
-- ----------------------------
INSERT INTO `visitor` VALUES (1, '日本日本网友', '/static/images/2.png', '172.105.235.23', 'kefu001', '2023-07-13 10:00:53', '2023-07-27 08:56:11', NULL, 'b82a0238-fb9b-4558-9c10-0934e247044c', 0, '直接访问', '新加坡新加坡', '172.105.235.23', '');
INSERT INTO `visitor` VALUES (2, '日本日本网友', '/static/images/2.png', '172.105.235.23', 'kefu001', '2023-07-13 10:01:31', '2023-07-27 09:11:03', NULL, '0fca9f0a-b445-4100-8542-e0d95ef6db95', 0, '直接访问', '新加坡新加坡', '172.105.235.23', '');
INSERT INTO `visitor` VALUES (3, '中国河南驻马店网友', '/static/images/2.png', '111.7.100.21', 'kefu001', '2023-07-13 22:29:18', '2023-07-23 05:37:47', NULL, 'ec439d65-c30f-433a-a156-c56d50698af4', 0, '直接访问', '中国河南驻马店', '111.7.100.21', '');
INSERT INTO `visitor` VALUES (4, '中国河南驻马店网友', '/static/images/2.png', '111.7.100.23', 'kefu001', '2023-07-13 22:29:21', '2023-07-23 05:37:47', NULL, '2065c099-1981-4452-bc89-3ca3c2ed38bc', 0, '直接访问', '中国河南驻马店', '111.7.100.23', '');
INSERT INTO `visitor` VALUES (5, '中国广东深圳网友', '/static/images/1.png', '112.97.52.132', 'kefu001', '2023-07-19 08:26:33', '2023-07-23 05:37:47', NULL, 'fae4c9b8-0b70-41b4-9087-4de75d0808b1', 0, '直接访问', '中国广东深圳', '112.97.52.132', '');
INSERT INTO `visitor` VALUES (6, '中国北京北京网友', '/static/images/1.png', '49.7.199.230', 'kefu001', '2023-07-19 08:27:33', '2023-07-23 05:37:47', NULL, 'b933099b-d160-4fd7-a0cc-50d0c3d4ab83', 0, '直接访问', '中国北京北京', '49.7.199.230', '');
INSERT INTO `visitor` VALUES (7, '中国安徽宣城网友', '/static/images/1.png', '36.56.102.158', 'kefu001', '2023-07-19 08:27:34', '2023-07-23 05:37:47', NULL, 'f5955bbc-f3a9-40e9-8dfe-1dcaa3929507', 0, '直接访问', '中国安徽宣城', '36.56.102.158', '');
INSERT INTO `visitor` VALUES (8, '中国广东深圳网友', '/static/images/1.png', '183.15.204.177', 'kefu001', '2023-07-19 08:27:35', '2023-07-23 05:37:47', NULL, '05e1612b-c56c-4494-9344-a6cf91541c47', 0, 'http://nativeapp.toutiao.com/', '中国广东深圳', '183.15.204.177', '');
INSERT INTO `visitor` VALUES (9, '美国美国网友', '/static/images/2.png', '67.21.82.245', 'kefu001', '2023-07-19 09:11:46', '2023-07-23 05:37:47', NULL, '22a2f80e-1213-4133-9307-4ae45d3f0aff', 0, '直接访问', '美国美国', '67.21.82.245', '');
INSERT INTO `visitor` VALUES (10, '中国香港网友', '/static/images/1.png', '139.5.108.71', 'kefu001', '2023-07-19 11:39:09', '2023-07-23 05:37:47', NULL, '37519c3a-9d97-49d7-9b30-3c2a27c002eb', 0, '直接访问', '中国广东深圳', '139.5.108.71', '');
INSERT INTO `visitor` VALUES (11, '中国广东深圳网友', '/static/images/1.png', '113.118.117.75', 'kefu001', '2023-07-19 13:33:53', '2023-07-23 05:37:47', NULL, 'e1befa1a-c60b-4ae3-abf6-aaafb4bfced3', 0, '直接访问', '中国广东深圳', '113.118.117.75', '');
INSERT INTO `visitor` VALUES (12, '中国香港网友', '/static/images/1.png', '139.5.108.71', 'kefu001', '2023-07-20 02:35:56', '2023-07-23 05:37:47', NULL, '811619bd-6ce9-43ca-9d67-b75dee8cc7d4', 0, '直接访问', '中国香港', '139.5.108.71', '');
INSERT INTO `visitor` VALUES (13, '中国香港网友', '/static/images/1.png', '139.5.108.135', 'kefu001', '2023-07-20 03:43:53', '2023-07-23 05:37:47', NULL, '7ca48046-b4a6-4e27-b8ee-67eebec0bc5b', 0, '直接访问', '印度印度', '139.5.108.135', '');
INSERT INTO `visitor` VALUES (14, '美国美国网友', '/static/images/1.png', '107.148.239.239', 'kefu001', '2023-07-20 05:25:05', '2023-07-23 05:37:47', NULL, 'd7692cf4-45bc-42ca-a0d8-9806969bdab7', 0, '直接访问', '美国美国', '107.148.239.239', '');
INSERT INTO `visitor` VALUES (15, '中国香港网友', '/static/images/1.png', '159.138.54.221', 'kefu001', '2023-07-20 06:32:27', '2023-07-23 05:37:47', NULL, '853ba72a-3dbf-4adf-b2e4-c0e9d26e14bb', 0, '直接访问', '中国广东', '159.138.54.221', '');
INSERT INTO `visitor` VALUES (16, '美国美国网友', '/static/images/1.png', '107.148.239.239', 'kefu001', '2023-07-21 01:51:19', '2023-07-27 05:04:19', NULL, '03558bcc-7b48-44a9-b42d-307e6859eb4b', 0, '直接访问', '日本日本', '107.148.239.239', '');
INSERT INTO `visitor` VALUES (17, '蒙古蒙古网友', '/static/images/1.png', '49.0.245.9', 'kefu001', '2023-07-21 09:48:42', '2023-07-24 08:23:00', NULL, '803f21b5-87f1-4b51-bc28-cdd2cbd59a6e', 0, '直接访问', '中国香港', '49.0.245.9', '');
INSERT INTO `visitor` VALUES (18, '中国广东深圳网友', '/static/images/1.png', '183.15.206.81', 'kefu001', '2023-07-24 09:22:04', '2023-07-24 09:22:13', NULL, '1ccfbbea-ee1f-4173-b930-439258f2fb29', 0, '直接访问', '中国广东深圳', '183.15.206.81', '');
INSERT INTO `visitor` VALUES (19, '中国广东深圳网友', '/static/images/1.png', '183.15.206.81', 'kefu001', '2023-07-25 03:38:46', '2023-07-25 09:29:31', NULL, '9ecd8bb8-21b5-4fcb-95aa-1af6fa2ea509', 0, '直接访问', '中国广东深圳', '183.15.206.81', '');
INSERT INTO `visitor` VALUES (20, '中国香港网友', '/static/images/1.png', '103.7.29.109', 'kefu001', '2023-07-25 05:16:14', '2023-07-25 11:21:22', NULL, '9bd1463c-4b61-4c3f-9a15-bfd62701af82', 0, '直接访问', '中国广东', '103.7.29.109', '');
INSERT INTO `visitor` VALUES (21, '日本日本网友', '/static/images/1.png', '43.132.98.39', 'kefu001', '2023-07-25 06:36:48', '2023-07-25 06:49:48', NULL, '6c22a29a-a003-4d73-a046-c6365b7e37aa', 0, '直接访问', '日本日本', '43.132.98.39', '');
INSERT INTO `visitor` VALUES (22, '中国广东深圳网友', '/static/images/1.png', '183.15.206.81', 'kefu001', '2023-07-25 06:45:12', '2023-07-25 06:46:45', NULL, 'e296367f-9839-4585-b91e-b2ab44044198', 0, '直接访问', '中国广东深圳', '183.15.206.81', '');
INSERT INTO `visitor` VALUES (23, '中国广东深圳网友', '/static/images/1.png', '183.15.207.234', 'kefu001', '2023-07-25 07:16:31', '2023-07-28 06:28:18', NULL, '3eb0ca86-0435-4a74-b277-de3a63fd9e87', 0, '直接访问', '中国广东深圳', '183.15.207.234', '');
INSERT INTO `visitor` VALUES (24, '中国广东深圳网友', '/static/images/1.png', '183.15.206.81', 'kefu001', '2023-07-25 09:38:18', '2023-07-25 09:38:24', NULL, '600eef6e-eed5-47bb-9026-b164e0c800c4', 0, '直接访问', '中国广东深圳', '183.15.206.81', '');
INSERT INTO `visitor` VALUES (25, '中国广东深圳网友', '/static/images/1.png', '183.15.206.81', 'kefu001', '2023-07-25 11:51:00', '2023-07-25 11:53:13', NULL, '8d7901d4-db77-4508-8b66-eee9126bec9c', 0, '直接访问', '中国广东深圳', '183.15.206.81', '');
INSERT INTO `visitor` VALUES (26, '中国香港网友', '/static/images/1.png', '139.5.108.209', 'kefu001', '2023-07-25 11:58:32', '2023-07-25 11:59:28', NULL, '000fb1ff-f2a7-4c84-b2fb-22709eb30552', 0, '直接访问', '中国香港', '139.5.108.209', '');
INSERT INTO `visitor` VALUES (27, '日本日本网友', '/static/images/1.png', '43.132.98.31', 'kefu001', '2023-07-26 02:49:15', '2023-07-26 02:52:14', NULL, 'a095507f-c86e-45b0-a7cb-ef7111aa9da1', 0, '直接访问', '日本日本', '43.132.98.31', '');
INSERT INTO `visitor` VALUES (28, '中国广东深圳网友', '/static/images/1.png', '183.15.207.234', 'kefu001', '2023-07-27 01:29:36', '2023-07-27 01:30:03', NULL, 'c0140744-3e0b-40fb-b920-15c793248c1f', 0, '直接访问', '中国广东深圳', '183.15.207.234', '');
INSERT INTO `visitor` VALUES (29, '美国美国网友', '/static/images/1.png', '107.148.239.239', 'kefu001', '2023-07-27 01:33:21', '2023-07-27 07:50:32', NULL, 'a67a6d6b-e435-4bd5-bc70-86b821c00d21', 0, '直接访问', '美国美国', '107.148.239.239', '');
INSERT INTO `visitor` VALUES (30, '日本日本网友', '/static/images/1.png', '43.132.98.32', 'kefu001', '2023-07-27 02:06:38', '2023-07-28 06:59:08', NULL, '6460f21f-7f9a-4990-9d95-4d0673718d02', 0, '直接访问', '日本日本', '43.132.98.32', '');
INSERT INTO `visitor` VALUES (31, '日本日本网友', '/static/images/1.png', '43.132.98.33', 'kefu001', '2023-07-27 08:19:11', '2023-07-28 06:58:05', NULL, 'e3423c39-6e60-4e50-8dfe-1d3d2563aeeb', 0, '直接访问', '日本日本', '43.132.98.33', '');
INSERT INTO `visitor` VALUES (32, '中国广东深圳网友', '/static/images/1.png', '183.15.207.234', 'kefu001', '2023-07-27 08:53:09', '2023-07-27 12:10:31', NULL, '088e3251-2207-4012-8d12-7132ea104203', 0, '直接访问', '中国广东深圳', '183.15.207.234', '');
INSERT INTO `visitor` VALUES (33, '中国广东深圳网友', '/static/images/1.png', '183.15.207.234', 'kefu001', '2023-07-27 12:21:17', '2023-07-27 12:26:07', NULL, '6280c7e0-0dbb-4475-84af-5fbe558241e9', 0, '直接访问', '中国广东深圳', '183.15.207.234', '');
INSERT INTO `visitor` VALUES (34, '中国广东深圳网友', '/static/images/1.png', '219.133.101.53', 'kefu001', '2023-07-28 01:47:34', '2023-07-29 02:49:43', NULL, '564dc11c-6cb4-40f7-8cca-98fa2ef3df66', 0, '直接访问', '日本日本', '219.133.101.53', '');
INSERT INTO `visitor` VALUES (35, '美国美国网友', '/static/images/1.png', '107.148.239.239', 'kefu001', '2023-07-28 05:52:37', '2023-07-28 05:53:34', NULL, 'a55d7ece-c032-4e81-b231-588a2c458d6e', 0, '直接访问', '美国美国', '107.148.239.239', '');
INSERT INTO `visitor` VALUES (36, '美国美国网友', '/static/images/1.png', '107.148.239.239', 'kefu001', '2023-07-28 06:31:55', '2023-07-28 12:11:00', NULL, 'f675a15e-d25e-4b1c-a831-6e1779e06b69', 0, '直接访问', '美国美国', '107.148.239.239', '');
INSERT INTO `visitor` VALUES (37, '中国广东深圳网友', '/static/images/1.png', '183.15.207.234', 'kefu001', '2023-07-29 03:58:06', '2023-07-29 04:43:15', NULL, '3eb10275-ae99-4fd7-aa09-99c56914c0d7', 0, '直接访问', '中国广东深圳', '183.15.207.234', '');
INSERT INTO `visitor` VALUES (38, '日本日本网友', '/static/images/1.png', '43.132.98.38', 'kefu001', '2023-07-31 02:33:09', '2023-07-31 02:33:46', NULL, '5531c5e9-f203-475a-a4d4-edfcab571b50', 0, '直接访问', '日本日本', '43.132.98.38', '');

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
INSERT INTO `welcome` VALUES (1, 'kefu2', 'offline', 'Я сейчас оффлайн, ваше сообщение было отправлено на мой адрес электронной почты, и я отвечу позже ~', 1, '2020-08-24 02:57:49');
INSERT INTO `welcome` VALUES (2, 'kefu2', 'welcome', 'Чем могу помочь ?', 0, '2020-08-24 02:57:49');

SET FOREIGN_KEY_CHECKS = 1;
