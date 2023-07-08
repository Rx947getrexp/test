/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : speed

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 08/07/2023 16:17:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_res
-- ----------------------------
DROP TABLE IF EXISTS `admin_res`;
CREATE TABLE `admin_res`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资源名称',
  `res_type` int NOT NULL COMMENT '类型：1-菜单；2-接口；3-按钮',
  `pid` int NOT NULL COMMENT '上级id（没有默认为0）',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'url地址',
  `sort` int NOT NULL COMMENT '排序',
  `icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `is_del` int NOT NULL COMMENT '软删状态：0-未删（默认）；1-已删',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `res_path`(`url` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 129 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_res
-- ----------------------------
INSERT INTO `admin_res` VALUES (34, '首页', 1, 0, '/home', 0, 'home', 0, '2023-02-28 20:27:19', '2023-02-28 20:27:11', 'root');
INSERT INTO `admin_res` VALUES (35, '系统管理', 1, 0, '/sys', 0, 'international', 0, '2023-02-28 20:28:16', '2023-02-28 20:28:10', 'root');
INSERT INTO `admin_res` VALUES (36, '管理员管理', 1, 0, '/admin', 0, 'user', 0, '2023-02-28 20:30:02', '2023-05-10 15:30:42', 'root');
INSERT INTO `admin_res` VALUES (37, '资源菜单', 1, 35, '/sys/menu', 0, 'example', 0, '2023-02-28 20:33:10', '2023-02-28 20:32:22', 'root');
INSERT INTO `admin_res` VALUES (39, '管理员信息', 1, 36, '/admin/user', 0, 'example', 0, '2023-02-28 20:38:32', '2023-02-28 20:38:26', 'root');
INSERT INTO `admin_res` VALUES (40, '角色权限', 1, 36, '/admin/role', 0, 'example', 0, '2023-02-28 20:39:19', '2023-02-28 20:39:15', 'root');
INSERT INTO `admin_res` VALUES (41, '密码修改', 1, 35, '/sys/passwd', 0, 'example', 0, '2023-03-02 15:27:45', '2023-03-02 15:27:45', 'root');
INSERT INTO `admin_res` VALUES (44, '谷歌验证器', 1, 35, '/sys/auth2', 0, 'example', 0, '2023-03-02 15:57:24', '2023-03-02 15:57:24', 'root');
INSERT INTO `admin_res` VALUES (104, '应用中心', 1, 0, '/app', 0, 'edit', 0, '2023-06-02 16:09:06', '2023-06-02 16:12:33', 'root');
INSERT INTO `admin_res` VALUES (105, '数据报表', 1, 0, '/report2', 0, 'chart', 0, '2023-06-02 16:10:31', '2023-06-02 16:11:29', 'root');
INSERT INTO `admin_res` VALUES (106, '节点列表', 1, 104, '/app/node_list', 0, 'example', 0, '2023-06-02 16:14:40', '2023-06-02 16:14:40', 'root');
INSERT INTO `admin_res` VALUES (107, '会员列表', 1, 104, '/app/member_list', 0, 'example', 0, '2023-06-02 16:15:14', '2023-06-02 16:15:14', 'root');
INSERT INTO `admin_res` VALUES (108, '节点uuid', 1, 104, '/app/node_uuid', 0, 'example', 0, '2023-06-02 16:15:43', '2023-06-02 16:15:43', 'root');
INSERT INTO `admin_res` VALUES (109, '套餐列表', 1, 104, '/app/goods_list', 0, 'example', 0, '2023-06-02 16:16:24', '2023-06-02 16:16:24', 'root');
INSERT INTO `admin_res` VALUES (110, '订单列表', 1, 104, '/app/order_list', 0, 'example', 0, '2023-06-02 16:16:46', '2023-06-02 16:16:46', 'root');
INSERT INTO `admin_res` VALUES (111, '赠送记录', 1, 104, '/app/gift_list', 0, 'example', 0, '2023-06-02 16:17:15', '2023-06-02 16:17:15', 'root');
INSERT INTO `admin_res` VALUES (112, '免费领会员', 1, 104, '/app/free_member', 0, 'example', 0, '2023-06-02 16:17:49', '2023-06-02 16:17:49', 'root');
INSERT INTO `admin_res` VALUES (113, '通知消息', 1, 104, '/app/notice_list', 0, 'example', 0, '2023-06-02 16:18:10', '2023-06-02 16:18:10', 'root');
INSERT INTO `admin_res` VALUES (114, '用户设备', 1, 104, '/app/user_dev_list', 0, 'example', 0, '2023-06-02 16:18:44', '2023-06-02 16:18:44', 'root');
INSERT INTO `admin_res` VALUES (115, '设备日志', 1, 104, '/app/dev_log_list', 0, 'example', 0, '2023-06-02 16:19:46', '2023-06-02 16:19:46', 'root');
INSERT INTO `admin_res` VALUES (116, '推广渠道', 1, 104, '/app/channel_list', 0, 'example', 0, '2023-06-02 16:20:10', '2023-06-02 16:20:10', 'root');
INSERT INTO `admin_res` VALUES (117, '加速日志', 1, 104, '/app/speed_log', 0, 'example', 0, '2023-06-02 16:20:50', '2023-06-02 16:21:00', 'root');
INSERT INTO `admin_res` VALUES (118, '广告列表', 1, 104, '/app/ad_list', 0, 'example', 0, '2023-06-02 16:21:22', '2023-06-02 16:21:22', 'root');
INSERT INTO `admin_res` VALUES (119, '域名列表', 1, 104, '/app/dns_list', 0, 'example', 0, '2023-06-02 16:21:58', '2023-06-02 16:21:58', 'root');
INSERT INTO `admin_res` VALUES (120, '平台用户按日', 1, 105, '/report2/plant_member_day', 0, 'example', 0, '2023-06-02 16:23:12', '2023-06-02 16:23:12', 'root');
INSERT INTO `admin_res` VALUES (121, '平台用户按月', 1, 105, '/report2/plant_member_month', 0, 'example', 0, '2023-06-02 16:23:40', '2023-06-02 16:23:40', 'root');
INSERT INTO `admin_res` VALUES (122, '下载渠道按日', 1, 105, '/report2/download_channel_day', 0, 'example', 0, '2023-06-02 16:24:13', '2023-06-02 16:24:13', 'root');
INSERT INTO `admin_res` VALUES (123, '下载渠道按月', 1, 105, '/report2/download_channel_month', 0, 'example', 0, '2023-06-02 16:24:40', '2023-06-02 16:24:40', 'root');
INSERT INTO `admin_res` VALUES (124, '平台收益按日', 1, 105, '/report2/plant_profit_day', 0, 'example', 0, '2023-06-02 16:25:13', '2023-06-02 16:25:13', 'root');
INSERT INTO `admin_res` VALUES (125, '平台收益按月', 1, 105, '/report2/plant_profit_month', 0, 'example', 0, '2023-06-02 16:25:38', '2023-06-02 16:25:38', 'root');
INSERT INTO `admin_res` VALUES (126, '广告统计按日', 1, 105, '/report2/ad_summary_day', 0, 'example', 0, '2023-06-02 16:26:10', '2023-06-02 16:26:46', 'root');
INSERT INTO `admin_res` VALUES (127, '广告统计按月', 1, 105, '/report2/ad_summary_month', 0, 'example', 0, '2023-06-02 16:26:32', '2023-06-02 16:26:32', 'root');
INSERT INTO `admin_res` VALUES (128, 'App配置', 1, 104, '/app/app_config', 0, 'example', 0, '2023-06-16 19:04:10', '2023-07-07 17:16:07', 'root');

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `is_del` int NOT NULL DEFAULT 0 COMMENT '0-正常；1-软删',
  `is_used` int NOT NULL COMMENT '1-已启用；2-未启用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (15, '超级管理员', 0, 1, '所有权限', '2023-02-10 21:02:15', '2023-06-02 16:12:59');
INSERT INTO `admin_role` VALUES (16, '测试管理员', 0, 1, '测试专用', '2023-02-10 20:50:57', '2023-03-02 18:48:03');
INSERT INTO `admin_role` VALUES (17, '财务', 0, 1, '负责公司财务', '2023-03-02 13:58:23', '2023-03-03 16:19:47');
INSERT INTO `admin_role` VALUES (18, '运维', 0, 1, '运维人员', '2023-03-02 19:15:48', '2023-03-27 20:49:59');

-- ----------------------------
-- Table structure for admin_role_res
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_res`;
CREATE TABLE `admin_role_res`  (
  `role_id` int NOT NULL COMMENT '角色id',
  `res_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资源id列表',
  `res_tree` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '资源菜单json树',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色资源表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role_res
-- ----------------------------
INSERT INTO `admin_role_res` VALUES (15, '34,35,37,41,44,36,39,40,104,105', NULL, '2023-02-10 21:02:15', '2023-06-02 16:12:59', 'root');
INSERT INTO `admin_role_res` VALUES (16, '34,35,37,41,44,45,36,39,40', NULL, '2023-02-10 20:50:57', '2023-03-02 18:48:03', 'root');
INSERT INTO `admin_role_res` VALUES (17, '34,35,37,41,44,36,39,40,83,84,85,86,87,88,89,90,91,92,93,94,95', '', '2023-03-02 13:58:23', '2023-03-03 16:19:47', 'root');
INSERT INTO `admin_role_res` VALUES (18, '34,41,44', '', '2023-03-02 19:15:48', '2023-03-27 20:49:59', 'root');

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `pwd2` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '二级密码',
  `authkey` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '谷歌验证码私钥',
  `status` int NULL DEFAULT 0 COMMENT '冻结状态：0-正常；1-冻结',
  `is_del` int NULL DEFAULT 0 COMMENT '0-正常；1-软删',
  `is_reset` int NULL DEFAULT NULL COMMENT '0-否；1-代表需要重置两步验证码',
  `is_first` int NULL DEFAULT NULL COMMENT '0-否；1-代表首次登录需要修改密码',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uname_index`(`uname` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100013 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (100001, 'zs', 'e10adc3949ba59abbe56e057f20f883e', 'zs1001', '2023-02-09 19:56:25', '2023-03-15 12:14:19', NULL, NULL, 0, 0, NULL, NULL);
INSERT INTO `admin_user` VALUES (100004, 'root', 'e9bc0e13a8a16cbb07b175d92a113126', 'root', '2023-02-10 20:30:44', '2023-03-02 18:27:57', '', 'B29D590954DF52E288ED4483DA332E0DD2D348B261086B69E3A19A03E1D22B5CFA0B455E1A0AAD75869DFC15E9F076F2', 0, 0, 0, NULL);
INSERT INTO `admin_user` VALUES (100005, 'zhaoliu', 'e10adc3949ba59abbe56e057f20f883e', '赵六', '2023-03-01 17:05:07', '2023-03-03 16:21:13', '', '', 0, 0, NULL, NULL);
INSERT INTO `admin_user` VALUES (100006, 'at', 'e9bc0e13a8a16cbb07b175d92a113126', 'aaa', '2023-03-15 12:16:45', '2023-03-15 12:18:52', '', '', 0, 1, 1, 0);
INSERT INTO `admin_user` VALUES (100008, 'df', 'e9bc0e13a8a16cbb07b175d92a113126', 'aaa', '2023-03-15 12:19:07', '2023-03-15 12:19:15', '', '', 0, 1, 1, 0);
INSERT INTO `admin_user` VALUES (100009, 'yyy', 'e9bc0e13a8a16cbb07b175d92a113126', '尼古拉斯', '2023-03-17 14:42:21', '2023-03-17 14:42:32', '', '', 0, 1, 1, 0);
INSERT INTO `admin_user` VALUES (100011, '123', '12b6231cc3459f0b9b7a799c73bac61f', '尼古拉斯', '2023-03-18 09:37:29', '2023-03-18 09:37:38', '', '', 0, 1, 1, 0);
INSERT INTO `admin_user` VALUES (100012, 'rrr', 'e9bc0e13a8a16cbb07b175d92a113126', '尼古拉斯', '2023-03-18 11:50:49', '2023-03-18 11:50:49', '', '', 0, 0, 1, 0);

-- ----------------------------
-- Table structure for admin_user_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_role`;
CREATE TABLE `admin_user_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增序列',
  `uid` int NOT NULL COMMENT '用户id',
  `role_id` int NOT NULL COMMENT '角色id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `is_del` int NOT NULL COMMENT '软删：0-未删；1-已删',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `u_role_user`(`uid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user_role
-- ----------------------------
INSERT INTO `admin_user_role` VALUES (1, 100001, 16, '2023-02-10 21:02:15', '2023-03-01 17:01:44', 0, 'root');
INSERT INTO `admin_user_role` VALUES (2, 100004, 15, '2023-02-10 21:02:15', '2023-02-28 17:11:22', 0, 'zs');
INSERT INTO `admin_user_role` VALUES (6, 100005, 17, '2023-03-01 17:05:07', '2023-03-03 16:21:13', 0, 'root');
INSERT INTO `admin_user_role` VALUES (7, 100006, 15, '2023-03-15 12:16:45', '2023-03-15 12:17:32', 0, 'root');
INSERT INTO `admin_user_role` VALUES (8, 100008, 15, '2023-03-15 12:19:07', '2023-03-15 12:19:07', 0, 'root');
INSERT INTO `admin_user_role` VALUES (9, 100009, 15, '2023-03-17 14:42:21', '2023-03-17 14:42:21', 0, 'root');
INSERT INTO `admin_user_role` VALUES (10, 100011, 15, '2023-03-18 09:37:29', '2023-03-18 09:37:29', 0, 'root');
INSERT INTO `admin_user_role` VALUES (11, 100012, 15, '2023-03-18 11:50:49', '2023-03-18 11:50:49', 0, 'root');

-- ----------------------------
-- Table structure for t_activity
-- ----------------------------
DROP TABLE IF EXISTS `t_activity`;
CREATE TABLE `t_activity`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `status` int NOT NULL COMMENT '状态:1-success；2-fail',
  `gift_sec` int NOT NULL COMMENT '赠送时间（失败为0）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1677224386010025985 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '免费领会员活动' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_activity
-- ----------------------------
INSERT INTO `t_activity` VALUES (1677224386010025984, 219122276, 2, 0, '2023-07-07 15:53:49', '2023-07-07 15:53:49', '');

-- ----------------------------
-- Table structure for t_ad
-- ----------------------------
DROP TABLE IF EXISTS `t_ad`;
CREATE TABLE `t_ad`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `status` int NOT NULL COMMENT '状态:1-上架；2-下架',
  `sort` int NOT NULL COMMENT '排序',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告名称',
  `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告logo',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '广告链接',
  `ad_type` int NOT NULL COMMENT '广告分类：1-社交；2-游戏；3-漫画；4-视频...',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '正文介绍',
  `author` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '广告表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_ad
-- ----------------------------
INSERT INTO `t_ad` VALUES (1, 1, 0, '环宇app1', '/public/upload/img/5c55ab8ddfcc1cf26f03e9929ba81b45.jpg', '12', 1, '环宇,app1', '2223', 'root', '2023-06-16 16:35:03', '2023-06-16 17:07:27', '');
INSERT INTO `t_ad` VALUES (2, 1, 0, '魔兽世界', '/public/upload/img/1f3d55461ba66f41eb7928cb351fa1ad.png', 'http://localhost', 2, '暴雪,九城', '网易代理魔兽世界', 'root', '2023-06-16 20:07:31', '2023-06-16 20:07:31', '');

-- ----------------------------
-- Table structure for t_channel
-- ----------------------------
DROP TABLE IF EXISTS `t_channel`;
CREATE TABLE `t_channel`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '渠道名称',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '渠道编号',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '渠道链接',
  `status` int NULL DEFAULT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推广渠道表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_channel
-- ----------------------------
INSERT INTO `t_channel` VALUES (1, '知乎', 'zhihu', 'http://127.0.0.1', 1, '2023-06-30 16:14:01', '2023-06-30 16:14:01', 'root', '');
INSERT INTO `t_channel` VALUES (2, '百度', 'baidu', 'http://localhost', 1, '2023-06-30 16:14:23', '2023-06-30 16:14:23', 'root', '');

-- ----------------------------
-- Table structure for t_dev
-- ----------------------------
DROP TABLE IF EXISTS `t_dev`;
CREATE TABLE `t_dev`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `os` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户端设备系统os',
  `network` int NULL DEFAULT 1 COMMENT '网络模式（1-自动；2-手动）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1677211938733428738 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户端设备表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_dev
-- ----------------------------
INSERT INTO `t_dev` VALUES (1677211938733428736, 'PostmanRuntime/7.32.2', 1, '2023-07-07 15:04:21', '2023-07-07 15:04:21', '');
INSERT INTO `t_dev` VALUES (1677211938733428737, 'PostmanRuntime/7.32.3', 1, '2023-07-07 17:08:37', '2023-07-07 17:08:40', NULL);

-- ----------------------------
-- Table structure for t_dict
-- ----------------------------
DROP TABLE IF EXISTS `t_dict`;
CREATE TABLE `t_dict`  (
  `key_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '键',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '值',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `is_del` int NULL DEFAULT NULL COMMENT '0-正常；1-软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`key_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_dict
-- ----------------------------
INSERT INTO `t_dict` VALUES ('app_js_zip', 'http://localhost/jszip', 'app静态包链接', 0, '2023-07-07 17:19:57', '2023-07-07 17:58:27');
INSERT INTO `t_dict` VALUES ('app_link', 'http://localhost/download', 'app下载链接', 0, '2023-06-16 19:11:40', '2023-07-07 17:58:27');
INSERT INTO `t_dict` VALUES ('app_version', 'v1.0.3', 'app最新版', 0, '2023-07-07 17:18:30', '2023-07-07 17:58:27');

-- ----------------------------
-- Table structure for t_gift
-- ----------------------------
DROP TABLE IF EXISTS `t_gift`;
CREATE TABLE `t_gift`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `op_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务id',
  `op_uid` bigint NULL DEFAULT NULL COMMENT '业务uid',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '赠送标题',
  `gift_sec` int NOT NULL COMMENT '赠送时间（单位s）',
  `g_type` int NOT NULL COMMENT '赠送类别（1-注册；2-推荐；3-日常活动；4-充值）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '赠送用户记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_gift
-- ----------------------------

-- ----------------------------
-- Table structure for t_goods
-- ----------------------------
DROP TABLE IF EXISTS `t_goods`;
CREATE TABLE `t_goods`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `m_type` int NOT NULL COMMENT '会员类型：1-vip1；2-vip2',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '套餐标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '套餐标题（英文）',
  `price` decimal(10, 6) NOT NULL COMMENT '单价(U)',
  `period` int NOT NULL COMMENT '有效期（天）',
  `dev_limit` int NOT NULL COMMENT '设备限制数',
  `flow_limit` bigint NOT NULL COMMENT '流量限制数；单位：字节；0-不限制',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品套餐表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_goods
-- ----------------------------
INSERT INTO `t_goods` VALUES (1, 1, 'vip1普通会员1周', '', 2.000000, 7, 2, 2000000000, 1, '2023-06-30 16:15:30', '2023-06-30 16:22:09', 'root', '');
INSERT INTO `t_goods` VALUES (2, 1, 'vip1普通会员1个月', '', 5.000000, 30, 2, 200000000, 1, '2023-06-30 16:18:04', '2023-06-30 16:21:55', 'root', '');
INSERT INTO `t_goods` VALUES (3, 1, 'vip1普通会员3个月', '', 11.000000, 90, 2, 3000000000, 1, '2023-06-30 16:19:37', '2023-06-30 16:21:10', 'root', '');
INSERT INTO `t_goods` VALUES (4, 1, 'vip1普通会员12个月', '', 40.000000, 365, 5, 9000000000000, 1, '2023-06-30 16:20:15', '2023-06-30 16:20:57', 'root', '');
INSERT INTO `t_goods` VALUES (5, 2, 'vip2超级会员1周', '', 3.000000, 7, 2, 500000000000, 1, '2023-06-30 16:23:25', '2023-06-30 16:23:31', 'root', '');
INSERT INTO `t_goods` VALUES (6, 2, 'vip2超级会员1个月', '', 7.000000, 30, 3, 5000000000000, 1, '2023-06-30 16:24:01', '2023-06-30 16:24:01', 'root', '');
INSERT INTO `t_goods` VALUES (7, 2, 'vip2超级会员3个月', '', 16.000000, 90, 5, 5000000000000000, 1, '2023-06-30 16:24:50', '2023-06-30 16:24:50', 'root', '');
INSERT INTO `t_goods` VALUES (8, 2, 'vip2超级会员12个月', '', 54.000000, 365, 5, 5000000000000000, 1, '2023-06-30 16:25:25', '2023-06-30 16:25:25', 'root', '');

-- ----------------------------
-- Table structure for t_node
-- ----------------------------
DROP TABLE IF EXISTS `t_node`;
CREATE TABLE `t_node`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点名称',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点标题（英文）',
  `country` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '国家',
  `country_en` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '国家（英文）',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内网IP',
  `server` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公网域名',
  `port` int NOT NULL COMMENT '公网端口',
  `cpu` int NOT NULL COMMENT 'cpu核数量（单位个）',
  `flow` bigint NOT NULL COMMENT '流量带宽',
  `disk` bigint NOT NULL COMMENT '磁盘容量（单位B）',
  `memory` bigint NOT NULL COMMENT '内存大小（单位B）',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_node
-- ----------------------------
INSERT INTO `t_node` VALUES (1, '线路1', '加拿大VIP1', '', '加拿大1', '', '10.10.10.111', '10.10.10.111', 8888, 0, 0, 0, 0, 1, '2023-07-01 10:13:09', '2023-07-01 10:13:35', 'root', '');

-- ----------------------------
-- Table structure for t_node_uuid
-- ----------------------------
DROP TABLE IF EXISTS `t_node_uuid`;
CREATE TABLE `t_node_uuid`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `node_id` bigint NOT NULL COMMENT '节点id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点邮箱，用于区分流量',
  `v2ray_uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点UUID',
  `server` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公网域名',
  `port` int NOT NULL COMMENT '公网端口',
  `used_flow` bigint NOT NULL COMMENT '已使用流量（单位B）',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '节点UUID表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_node_uuid
-- ----------------------------

-- ----------------------------
-- Table structure for t_notice
-- ----------------------------
DROP TABLE IF EXISTS `t_notice`;
CREATE TABLE `t_notice`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '正文内容',
  `author` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `status` int NULL DEFAULT NULL COMMENT '状态:1-发布；2-软删',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推荐关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_notice
-- ----------------------------
INSERT INTO `t_notice` VALUES (1, '欢迎新用户', '新用户,注册', '欢迎新用户来到本平台', 'root', '2023-06-16 17:07:17', '2023-06-16 17:08:22', 1, '');

-- ----------------------------
-- Table structure for t_order
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `goods_id` bigint NOT NULL COMMENT '商品id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品标题',
  `price` decimal(10, 6) NOT NULL COMMENT '单价(U)',
  `price_cny` decimal(10, 2) NOT NULL COMMENT '折合RMB单价(CNY)',
  `status` int NOT NULL COMMENT '订单状态:1-init；2-success；3-cancel',
  `finished_at` timestamp NULL DEFAULT NULL COMMENT '完成时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1677241462653194241 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品订单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_order
-- ----------------------------
INSERT INTO `t_order` VALUES (1677241462653194240, 219122276, 1, 'vip1普通会员1周', 2.000000, 2.00, 1, NULL, '2023-07-07 17:01:40', '2023-07-07 17:01:40', '');

-- ----------------------------
-- Table structure for t_site
-- ----------------------------
DROP TABLE IF EXISTS `t_site`;
CREATE TABLE `t_site`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `site` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '域名',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip',
  `status` int NULL DEFAULT NULL COMMENT '1-正常；2-软删',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '域名表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_site
-- ----------------------------
INSERT INTO `t_site` VALUES (1, 'locahost', '127.0.0.2', 1, 'root', '2023-06-16 17:19:42', '2023-06-16 17:23:30', '');
INSERT INTO `t_site` VALUES (2, 'localhost', '127.0.0.1', 1, 'root', '2023-06-16 17:23:22', '2023-06-16 17:23:22', '');

-- ----------------------------
-- Table structure for t_success_record
-- ----------------------------
DROP TABLE IF EXISTS `t_success_record`;
CREATE TABLE `t_success_record`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `order_id` bigint NOT NULL COMMENT '订单id',
  `start_time` bigint NOT NULL COMMENT '套餐开始时间（时间戳）',
  `end_time` bigint NOT NULL COMMENT '套餐结束时间（时间戳）',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品标题',
  `price` decimal(10, 6) NOT NULL COMMENT '单价(U)',
  `price_cny` decimal(10, 2) NOT NULL COMMENT '折合RMB单价(CNY)',
  `m_type` int NOT NULL COMMENT '会员类型：1-vip1；2-vip2',
  `pay_type` int NOT NULL COMMENT '订单状态:1-银行卡；2-支付宝；3-微信支付',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户购买成功记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_success_record
-- ----------------------------

-- ----------------------------
-- Table structure for t_upload_log
-- ----------------------------
DROP TABLE IF EXISTS `t_upload_log`;
CREATE TABLE `t_upload_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日志内容',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '上传日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_upload_log
-- ----------------------------
INSERT INTO `t_upload_log` VALUES (1, 219122276, 1677211938733428736, 'aaa error', '2023-07-07 16:24:21', '2023-07-07 16:24:21', '');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮件',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '电话',
  `level` int NULL DEFAULT NULL COMMENT '等级：0-vip0；1-vip1；2-vip2',
  `expired_time` bigint NULL DEFAULT NULL COMMENT 'vip到期时间',
  `v2ray_uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点UUID',
  `channel_id` int NULL DEFAULT NULL COMMENT '渠道id',
  `status` int NULL DEFAULT NULL COMMENT '冻结状态：0-正常；1-冻结',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `p_uname_index`(`uname` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 219122277 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (219122276, 'win2023bar@outlook.com', 'c33367701511b4f6020ec61ded352059', 'win2023bar@outlook.com', '', 0, 0, 'b736923a1c9711ee99370c9d92c013fb', 0, 0, '2023-07-07 15:27:23', '2023-07-07 16:58:57', '');

-- ----------------------------
-- Table structure for t_user_dev
-- ----------------------------
DROP TABLE IF EXISTS `t_user_dev`;
CREATE TABLE `t_user_dev`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已踢',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `u_dev`(`dev_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户设备表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user_dev
-- ----------------------------
INSERT INTO `t_user_dev` VALUES (1, 219122276, 1677211938733428736, 1, '2023-07-07 15:42:36', '2023-07-07 15:42:36', '');
INSERT INTO `t_user_dev` VALUES (2, 219122276, 1677211938733428737, 2, '2023-07-07 17:09:11', '2023-07-07 17:10:34', NULL);

-- ----------------------------
-- Table structure for t_user_team
-- ----------------------------
DROP TABLE IF EXISTS `t_user_team`;
CREATE TABLE `t_user_team`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `direct_id` bigint NOT NULL COMMENT '上级id',
  `direct_tree` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '上级列表',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推荐关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user_team
-- ----------------------------
INSERT INTO `t_user_team` VALUES (1, 219122276, 0, '', '2023-07-07 15:27:23', '2023-07-07 15:27:23', '');

-- ----------------------------
-- Table structure for t_work_log
-- ----------------------------
DROP TABLE IF EXISTS `t_work_log`;
CREATE TABLE `t_work_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `mode_type` int NOT NULL COMMENT '模式类别:1-智能；2-手选',
  `node_id` bigint NOT NULL COMMENT '工作节点',
  `flow` bigint NOT NULL COMMENT '使用流量（字节）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '工作日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_work_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_work_mode
-- ----------------------------
DROP TABLE IF EXISTS `t_work_mode`;
CREATE TABLE `t_work_mode`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `mode_type` int NOT NULL COMMENT '模式类别:1-智能；2-手选',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `u_dev`(`dev_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '工作模式' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_work_mode
-- ----------------------------

-- ----------------------------
-- Table structure for user_logs
-- ----------------------------
DROP TABLE IF EXISTS `user_logs`;
CREATE TABLE `user_logs`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `datestr` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期',
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'IP地址',
  `user_agent` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求头user-agent',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `log_user_date`(`user_id` ASC, `datestr` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2807 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户日志表（仅记录第一次事件)' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_logs
-- ----------------------------
INSERT INTO `user_logs` VALUES (4, 1641398917893459968, '2023-04-04', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-04 14:11:44', '2023-04-04 14:11:44', NULL);
INSERT INTO `user_logs` VALUES (8, 100004, '2023-04-04', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-04 14:16:37', '2023-04-04 14:16:37', NULL);
INSERT INTO `user_logs` VALUES (13, 100004, '2023-04-06', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-06 20:09:23', '2023-04-06 20:09:23', NULL);
INSERT INTO `user_logs` VALUES (23, 1641398917893459968, '2023-04-06', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-06 20:32:46', '2023-04-06 20:32:46', NULL);
INSERT INTO `user_logs` VALUES (30, 1641398917893459968, '2023-04-07', '10.10.10.18', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1', '2023-04-07 10:10:23', '2023-04-07 10:10:23', NULL);
INSERT INTO `user_logs` VALUES (32, 1641346214014226432, '2023-04-07', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-07 14:24:18', '2023-04-07 14:24:18', NULL);
INSERT INTO `user_logs` VALUES (60, 100004, '2023-04-07', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-07 15:01:05', '2023-04-07 15:01:05', NULL);
INSERT INTO `user_logs` VALUES (302, 1640976539392675840, '2023-04-07', '10.10.10.81', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.96 Mobile Safari/537.36', '2023-04-07 20:09:35', '2023-04-07 20:09:35', NULL);
INSERT INTO `user_logs` VALUES (345, 100004, '2023-04-08', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-08 09:36:53', '2023-04-08 09:36:53', NULL);
INSERT INTO `user_logs` VALUES (361, 1641398917893459968, '2023-04-08', '10.10.10.18', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1', '2023-04-08 18:54:49', '2023-04-08 18:54:49', NULL);
INSERT INTO `user_logs` VALUES (376, 1641398917893459968, '2023-04-10', '10.10.10.222', 'Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36', '2023-04-10 15:11:43', '2023-04-10 15:11:43', NULL);
INSERT INTO `user_logs` VALUES (403, 100004, '2023-04-10', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-10 15:33:53', '2023-04-10 15:33:53', NULL);
INSERT INTO `user_logs` VALUES (406, 1641346214014226432, '2023-04-10', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-10 15:33:58', '2023-04-10 15:33:58', NULL);
INSERT INTO `user_logs` VALUES (784, 100004, '2023-04-11', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-11 09:30:58', '2023-04-11 09:30:58', NULL);
INSERT INTO `user_logs` VALUES (787, 1641398917893459968, '2023-04-11', '10.10.10.222', 'Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36', '2023-04-11 09:32:07', '2023-04-11 09:32:07', NULL);
INSERT INTO `user_logs` VALUES (793, 1641346214014226432, '2023-04-11', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-04-11 09:34:01', '2023-04-11 09:34:01', NULL);
INSERT INTO `user_logs` VALUES (823, 12, '2023-04-11', '10.10.10.111', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.34', '2023-04-11 17:22:31', '2023-04-11 17:22:31', NULL);
INSERT INTO `user_logs` VALUES (828, 100004, '2023-05-10', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-05-10 15:24:36', '2023-05-10 15:24:36', NULL);
INSERT INTO `user_logs` VALUES (924, 100004, '2023-05-11', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-05-11 09:12:39', '2023-05-11 09:12:39', NULL);
INSERT INTO `user_logs` VALUES (945, 100004, '2023-05-12', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-05-12 14:24:27', '2023-05-12 14:24:27', NULL);
INSERT INTO `user_logs` VALUES (1001, 100004, '2023-05-15', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36', '2023-05-15 14:14:39', '2023-05-15 14:14:39', NULL);
INSERT INTO `user_logs` VALUES (1047, 100004, '2023-05-18', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-05-18 10:09:49', '2023-05-18 10:09:49', NULL);
INSERT INTO `user_logs` VALUES (1057, 100004, '2023-05-19', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-05-19 09:59:47', '2023-05-19 09:59:47', NULL);
INSERT INTO `user_logs` VALUES (1060, 100004, '2023-06-02', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-02 14:24:26', '2023-06-02 14:24:26', NULL);
INSERT INTO `user_logs` VALUES (1250, 100004, '2023-06-05', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-05 10:11:28', '2023-06-05 10:11:28', NULL);
INSERT INTO `user_logs` VALUES (1281, 100004, '2023-06-06', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-06 10:21:11', '2023-06-06 10:21:11', NULL);
INSERT INTO `user_logs` VALUES (1316, 100004, '2023-06-07', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-07 17:53:00', '2023-06-07 17:53:00', NULL);
INSERT INTO `user_logs` VALUES (1324, 100004, '2023-06-12', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-12 09:39:20', '2023-06-12 09:39:20', NULL);
INSERT INTO `user_logs` VALUES (1337, 100004, '2023-06-15', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-15 10:10:03', '2023-06-15 10:10:03', NULL);
INSERT INTO `user_logs` VALUES (1619, 100004, '2023-06-16', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-16 14:56:04', '2023-06-16 14:56:04', NULL);
INSERT INTO `user_logs` VALUES (1810, 100004, '2023-06-17', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-17 11:23:01', '2023-06-17 11:23:01', NULL);
INSERT INTO `user_logs` VALUES (1816, 100004, '2023-06-20', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-20 15:03:00', '2023-06-20 15:03:00', NULL);
INSERT INTO `user_logs` VALUES (1820, 100004, '2023-06-21', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-21 14:13:41', '2023-06-21 14:13:41', NULL);
INSERT INTO `user_logs` VALUES (1939, 100004, '2023-06-25', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-25 09:17:44', '2023-06-25 09:17:44', NULL);
INSERT INTO `user_logs` VALUES (2124, 100004, '2023-06-26', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-26 14:41:33', '2023-06-26 14:41:33', NULL);
INSERT INTO `user_logs` VALUES (2140, 100004, '2023-06-27', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-27 10:36:29', '2023-06-27 10:36:29', NULL);
INSERT INTO `user_logs` VALUES (2375, 100004, '2023-06-28', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-28 16:10:15', '2023-06-28 16:10:15', NULL);
INSERT INTO `user_logs` VALUES (2383, 100004, '2023-06-30', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-06-30 10:53:32', '2023-06-30 10:53:32', NULL);
INSERT INTO `user_logs` VALUES (2508, 100004, '2023-07-01', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-07-01 10:04:43', '2023-07-01 10:04:43', NULL);
INSERT INTO `user_logs` VALUES (2722, 100004, '2023-07-03', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-07-03 10:50:46', '2023-07-03 10:50:46', NULL);
INSERT INTO `user_logs` VALUES (2723, 219122276, '2023-07-07', '127.0.0.1', 'PostmanRuntime/7.32.2', '2023-07-07 15:39:22', '2023-07-07 15:39:22', NULL);
INSERT INTO `user_logs` VALUES (2746, 100004, '2023-07-07', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-07-07 17:14:35', '2023-07-07 17:14:35', NULL);

SET FOREIGN_KEY_CHECKS = 1;
