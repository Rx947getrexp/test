/*
 Navicat Premium Data Transfer

 Source Server         : 加速器测试库
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33-0ubuntu0.22.04.2)
 Source Host           : localhost:3306
 Source Schema         : speed

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33-0ubuntu0.22.04.2)
 File Encoding         : 65001

 Date: 31/07/2023 14:32:47
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
) ENGINE = InnoDB AUTO_INCREMENT = 131 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资源表' ROW_FORMAT = Dynamic;

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
INSERT INTO `admin_res` VALUES (119, 'App域名', 1, 104, '/app/app_dns_list', 0, 'example', 0, '2023-06-02 16:21:58', '2023-07-17 20:32:50', 'root');
INSERT INTO `admin_res` VALUES (120, '平台用户按日', 1, 105, '/report2/plant_member_day', 0, 'example', 0, '2023-06-02 16:23:12', '2023-06-02 16:23:12', 'root');
INSERT INTO `admin_res` VALUES (121, '平台用户按月', 1, 105, '/report2/plant_member_month', 0, 'example', 0, '2023-06-02 16:23:40', '2023-06-02 16:23:40', 'root');
INSERT INTO `admin_res` VALUES (122, '下载渠道按日', 1, 105, '/report2/download_channel_day', 0, 'example', 0, '2023-06-02 16:24:13', '2023-06-02 16:24:13', 'root');
INSERT INTO `admin_res` VALUES (123, '下载渠道按月', 1, 105, '/report2/download_channel_month', 0, 'example', 0, '2023-06-02 16:24:40', '2023-06-02 16:24:40', 'root');
INSERT INTO `admin_res` VALUES (124, '平台收益按日', 1, 105, '/report2/plant_profit_day', 0, 'example', 0, '2023-06-02 16:25:13', '2023-06-02 16:25:13', 'root');
INSERT INTO `admin_res` VALUES (125, '平台收益按月', 1, 105, '/report2/plant_profit_month', 0, 'example', 0, '2023-06-02 16:25:38', '2023-06-02 16:25:38', 'root');
INSERT INTO `admin_res` VALUES (126, '广告统计按日', 1, 105, '/report2/ad_summary_day', 0, 'example', 0, '2023-06-02 16:26:10', '2023-06-02 16:26:46', 'root');
INSERT INTO `admin_res` VALUES (127, '广告统计按月', 1, 105, '/report2/ad_summary_month', 0, 'example', 0, '2023-06-02 16:26:32', '2023-06-02 16:26:32', 'root');
INSERT INTO `admin_res` VALUES (128, 'App配置', 1, 104, '/app/app_config', 0, 'example', 0, '2023-06-16 19:04:10', '2023-07-07 17:16:07', 'root');
INSERT INTO `admin_res` VALUES (129, 'App版本', 1, 104, '/app/app_version', 0, 'example', 0, '2023-07-17 15:30:04', '2023-07-17 19:13:27', 'root');
INSERT INTO `admin_res` VALUES (130, 'Ios账号', 1, 104, '/app/ios_account', 0, 'example', 0, '2023-07-17 19:16:16', '2023-07-17 19:16:16', 'root');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1685133650045177857 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '免费领会员活动' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_activity
-- ----------------------------
INSERT INTO `t_activity` VALUES (1677224386010025984, 219122276, 2, 0, '2023-07-07 15:53:49', '2023-07-07 15:53:49', '');
INSERT INTO `t_activity` VALUES (1681618631932252160, 219122277, 2, 0, '2023-07-19 10:54:59', '2023-07-19 10:54:59', '');
INSERT INTO `t_activity` VALUES (1681630239475634176, 219122277, 2, 0, '2023-07-19 11:41:06', '2023-07-19 11:41:06', '');
INSERT INTO `t_activity` VALUES (1681630800195358720, 219122279, 1, 3600, '2023-07-19 11:43:20', '2023-07-19 11:43:20', '');
INSERT INTO `t_activity` VALUES (1681890360827056128, 219122280, 1, 3600, '2023-07-20 04:54:44', '2023-07-20 04:54:44', '');
INSERT INTO `t_activity` VALUES (1681932501020315648, 219122279, 1, 3600, '2023-07-20 07:42:11', '2023-07-20 07:42:11', '');
INSERT INTO `t_activity` VALUES (1681998850782400512, 219122279, 2, 0, '2023-07-20 12:05:50', '2023-07-20 12:05:50', '');
INSERT INTO `t_activity` VALUES (1682311107827470336, 219122277, 2, 0, '2023-07-21 08:46:38', '2023-07-21 08:46:38', '');
INSERT INTO `t_activity` VALUES (1682330463558963200, 219122279, 2, 0, '2023-07-21 10:03:33', '2023-07-21 10:03:33', '');
INSERT INTO `t_activity` VALUES (1682584590595657728, 219122277, 2, 0, '2023-07-22 02:53:21', '2023-07-22 02:53:21', '');
INSERT INTO `t_activity` VALUES (1682586131201265664, 219122279, 1, 3600, '2023-07-22 02:59:29', '2023-07-22 02:59:29', '');
INSERT INTO `t_activity` VALUES (1682699880247595008, 219122277, 2, 0, '2023-07-22 10:31:29', '2023-07-22 10:31:29', '');
INSERT INTO `t_activity` VALUES (1683100674218266624, 219122277, 1, 3600, '2023-07-23 13:04:05', '2023-07-23 13:04:05', '');
INSERT INTO `t_activity` VALUES (1683103039457595392, 219122279, 2, 0, '2023-07-23 13:13:29', '2023-07-23 13:13:29', '');
INSERT INTO `t_activity` VALUES (1683288201172619264, 219122279, 1, 3600, '2023-07-24 01:29:15', '2023-07-24 01:29:15', '');
INSERT INTO `t_activity` VALUES (1683680509349072896, 219122279, 2, 0, '2023-07-25 03:28:09', '2023-07-25 03:28:09', '');
INSERT INTO `t_activity` VALUES (1683727644723515392, 219122279, 1, 3600, '2023-07-25 06:35:27', '2023-07-25 06:35:27', '');
INSERT INTO `t_activity` VALUES (1683731957860536320, 219122276, 1, 3600, '2023-07-25 06:52:35', '2023-07-25 06:52:35', '');
INSERT INTO `t_activity` VALUES (1683732910047236096, 219122277, 1, 3600, '2023-07-25 06:56:22', '2023-07-25 06:56:22', '');
INSERT INTO `t_activity` VALUES (1683733513477558272, 219122284, 1, 3600, '2023-07-25 06:58:46', '2023-07-25 06:58:46', '');
INSERT INTO `t_activity` VALUES (1683751437881839616, 219122284, 2, 0, '2023-07-25 08:10:00', '2023-07-25 08:10:00', '');
INSERT INTO `t_activity` VALUES (1683756080745680896, 219122279, 2, 0, '2023-07-25 08:28:26', '2023-07-25 08:28:26', '');
INSERT INTO `t_activity` VALUES (1683799043962048512, 219122282, 1, 3600, '2023-07-25 11:19:10', '2023-07-25 11:19:10', '');
INSERT INTO `t_activity` VALUES (1684130886091542528, 219122277, 1, 3600, '2023-07-26 09:17:47', '2023-07-26 09:17:47', '');
INSERT INTO `t_activity` VALUES (1684132709233856512, 219122284, 1, 3600, '2023-07-26 09:25:02', '2023-07-26 09:25:02', '');
INSERT INTO `t_activity` VALUES (1684178698762194944, 219122284, 2, 0, '2023-07-26 12:27:46', '2023-07-26 12:27:46', '');
INSERT INTO `t_activity` VALUES (1684377084908015616, 219122284, 2, 0, '2023-07-27 01:36:05', '2023-07-27 01:36:05', '');
INSERT INTO `t_activity` VALUES (1684404379735560192, 219122282, 1, 3600, '2023-07-27 03:24:33', '2023-07-27 03:24:33', '');
INSERT INTO `t_activity` VALUES (1684468325079322624, 219122282, 1, 3600, '2023-07-27 07:38:39', '2023-07-27 07:38:39', '');
INSERT INTO `t_activity` VALUES (1684542307329642496, 219122282, 1, 3600, '2023-07-27 12:32:38', '2023-07-27 12:32:38', '');
INSERT INTO `t_activity` VALUES (1684900680835272704, 219122282, 1, 3600, '2023-07-28 12:16:40', '2023-07-28 12:16:40', '');
INSERT INTO `t_activity` VALUES (1685133650045177856, 219122305, 2, 0, '2023-07-29 03:42:25', '2023-07-29 03:42:25', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '广告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_ad
-- ----------------------------
INSERT INTO `t_ad` VALUES (1, 1, 0, '环宇app', '/public/upload/img/5ccb6c5bb3c38d15e609d462e5080fdc.jpg', '12', 1, '环宇,app', '2223', 'root', '2023-06-16 16:35:03', '2023-07-25 11:52:57', '');
INSERT INTO `t_ad` VALUES (2, 1, 0, '魔兽世界', '/public/upload/img/8b9ac785adedd6f39402df9968cc8494.png', 'http://localhost', 2, '暴雪,九城', '网易代理魔兽世界', 'root', '2023-06-16 20:07:31', '2023-07-25 11:52:18', '');

-- ----------------------------
-- Table structure for t_app_dns
-- ----------------------------
DROP TABLE IF EXISTS `t_app_dns`;
CREATE TABLE `t_app_dns`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `site_type` int NOT NULL COMMENT '站点类型:1-appapi域名；2-offsite域名；3-onlineservice在线客服；4-管理后台...',
  `dns` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '域名',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `level` int NOT NULL COMMENT '线路级别:1,2,3...用于白名单机制',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'app域名表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_app_dns
-- ----------------------------
INSERT INTO `t_app_dns` VALUES (1, 1, 'www.wuwuwu360.xyz', '107.148.68.118', 1, 1, '2023-07-17 21:06:37', '2023-07-18 19:34:28', 'root', '');
INSERT INTO `t_app_dns` VALUES (2, 1, 'www.wandan8.xyz ', '107.148.68.117', 2, 1, '2023-07-17 21:06:57', '2023-07-18 19:34:19', 'root', '');
INSERT INTO `t_app_dns` VALUES (3, 1, 'www.yyy360.xyz', '107.148.68.116', 1, 1, '2023-07-18 19:32:45', '2023-07-18 19:32:45', 'root', '');
INSERT INTO `t_app_dns` VALUES (4, 3, 'im.yyy360.xyz', '107.148.68.116', 1, 1, '2023-07-19 09:41:45', '2023-07-19 09:41:45', 'root', '');

-- ----------------------------
-- Table structure for t_app_version
-- ----------------------------
DROP TABLE IF EXISTS `t_app_version`;
CREATE TABLE `t_app_version`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `app_type` int NOT NULL COMMENT '1-ios;2-安卓；3-h5zip',
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '版本号',
  `link` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '超链地址',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'app版本管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_app_version
-- ----------------------------
INSERT INTO `t_app_version` VALUES (1, 3, 'v1.1.0', '/public/upload/other/f170a86fc6bfa6872bef474b0de73106.zip', 1, '2023-07-17 17:45:28', '2023-07-17 18:59:03', 'root', '');
INSERT INTO `t_app_version` VALUES (2, 3, 'v1.1.1', '/public/upload/other/20ad21cb1c8579c35e57fed83a6c71d7.zip', 1, '2023-07-17 19:04:04', '2023-07-17 19:04:04', 'root', '');
INSERT INTO `t_app_version` VALUES (3, 1, 'v1.0.0', 'http://localhost', 2, '2023-07-17 19:04:46', '2023-07-20 03:59:30', 'root', '');
INSERT INTO `t_app_version` VALUES (4, 3, 'v1.1.2', '/public/upload/other/28ec45e634adb839e51640f2877451d7.zip', 1, '2023-07-19 01:56:58', '2023-07-19 01:56:58', 'root', '');
INSERT INTO `t_app_version` VALUES (5, 3, 'v1.1.3', '/public/upload/other/450e433db39eae7aadeb9bfd75f33cab.zip', 1, '2023-07-19 06:14:55', '2023-07-19 06:14:55', 'root', '');
INSERT INTO `t_app_version` VALUES (6, 3, 'v1.1.4', '/public/upload/other/d605f98a6f46cfcee4bb46c695cd959b.zip', 1, '2023-07-19 11:35:53', '2023-07-19 11:35:53', 'root', '');
INSERT INTO `t_app_version` VALUES (7, 3, 'v1.1.5', '/public/upload/other/1a9079697d967ffb0d0e0a2ceb9f1cdf.zip', 1, '2023-07-20 03:42:28', '2023-07-20 03:42:28', 'root', '');
INSERT INTO `t_app_version` VALUES (8, 3, 'v1.1.6', '/public/upload/other/990948076532e845d9cbe357e36b0c1c.zip', 1, '2023-07-20 09:55:22', '2023-07-20 09:55:22', 'root', '');
INSERT INTO `t_app_version` VALUES (9, 3, 'v1.1.7', '/public/upload/other/0c1f3c68159dd53986f4a3e5017c0ef0.zip', 1, '2023-07-20 11:46:54', '2023-07-20 11:46:54', 'root', '');
INSERT INTO `t_app_version` VALUES (10, 3, 'v1.1.8', '/public/upload/other/cf40aefbf10c30a33aca4e24b82c0069.zip', 1, '2023-07-21 02:43:43', '2023-07-21 02:43:43', 'root', '');
INSERT INTO `t_app_version` VALUES (11, 3, 'v1.1.9', '/public/upload/other/8f2e5cb8156c6ae0494e9914c611a6b3.zip', 1, '2023-07-21 08:03:01', '2023-07-21 08:03:01', 'root', '');
INSERT INTO `t_app_version` VALUES (12, 3, 'v1.1.10', '/public/upload/other/6751770142734b475698546e13b80886.zip', 1, '2023-07-21 08:43:34', '2023-07-21 08:43:34', 'root', '');
INSERT INTO `t_app_version` VALUES (13, 3, 'v1.1.11', '/public/upload/other/a3dbf24d853e105d35e7eb6afb298f9c.zip', 1, '2023-07-22 01:23:11', '2023-07-22 01:23:11', 'root', '');
INSERT INTO `t_app_version` VALUES (14, 3, 'v1.1.12', '/public/upload/other/49dde1f9add12673794d78eee31b26be.zip', 2, '2023-07-22 02:33:59', '2023-07-22 02:37:56', 'root', '');
INSERT INTO `t_app_version` VALUES (15, 3, 'v1.1.13', '/public/upload/other/326e437715f99a5ce64d9fe1af21447e.zip', 1, '2023-07-22 02:44:10', '2023-07-22 02:44:10', 'root', '');
INSERT INTO `t_app_version` VALUES (16, 3, 'v0.0.1', '/public/upload/other/a1d11774ed74884f8bcb9c8bcc42359c.zip', 1, '2023-07-24 08:50:11', '2023-07-24 08:50:11', 'root', '');
INSERT INTO `t_app_version` VALUES (17, 3, 'v1.1.14', '/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip', 1, '2023-07-25 01:30:56', '2023-07-25 01:30:56', 'root', '');
INSERT INTO `t_app_version` VALUES (18, 3, 'v1.1.15', '/public/upload/other/e5ac32e358596301db8c84bc019d496f.zip', 1, '2023-07-25 06:22:23', '2023-07-25 06:22:23', 'root', '');
INSERT INTO `t_app_version` VALUES (19, 3, 'v1.1.16', '/public/upload/other/f915f5b2b07a4b3a73ee2333389e4dce.zip', 1, '2023-07-25 07:06:26', '2023-07-25 07:06:26', 'root', '');
INSERT INTO `t_app_version` VALUES (20, 3, 'v1.1.17', '/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip', 1, '2023-07-25 09:06:59', '2023-07-25 09:06:59', 'root', '');
INSERT INTO `t_app_version` VALUES (21, 3, 'v1.1.18', '/public/upload/other/27faddb133b27204959071ec0832b929.zip', 1, '2023-07-26 09:26:57', '2023-07-26 09:26:57', 'root', '');
INSERT INTO `t_app_version` VALUES (22, 3, 'v1.1.19', '/public/upload/other/vpn.zip', 1, '2023-07-26 20:09:47', '2023-07-26 20:09:53', 'root', NULL);
INSERT INTO `t_app_version` VALUES (23, 3, 'v1.1.23', '/public/upload/other/vpn0727.zip', 1, '2023-07-27 10:14:27', '2023-07-27 10:14:30', 'root', NULL);
INSERT INTO `t_app_version` VALUES (24, 3, 'v1.1.24', '/public/upload/other/vpn-v1.1.24.zip', 1, '2023-07-27 11:00:47', '2023-07-27 11:00:49', 'root', NULL);
INSERT INTO `t_app_version` VALUES (25, 3, 'v1.1.25', '/public/upload/other/1db7468b8042c218b33f3bbeefb3d7b6.zip', 1, '2023-07-27 11:37:16', '2023-07-27 11:37:16', 'root', '');
INSERT INTO `t_app_version` VALUES (26, 3, 'v1.1.26', '/public/upload/other/48cacf0f3c022dd92250a072c270ce99.zip', 1, '2023-07-28 02:39:37', '2023-07-28 02:39:37', 'root', '');
INSERT INTO `t_app_version` VALUES (27, 3, 'v1.1.30', '/public/upload/other/8e0fc5c713a0d9f8f765ed2f33d66aea.zip', 1, '2023-07-28 03:09:12', '2023-07-28 03:09:12', 'root', '');
INSERT INTO `t_app_version` VALUES (28, 3, 'v1.1.31', '/public/upload/other/5a6670e8640ed06aff6bf0c7bf890fba.zip', 1, '2023-07-28 03:33:27', '2023-07-28 03:33:27', 'root', '');
INSERT INTO `t_app_version` VALUES (29, 3, 'v1.1.32', '/public/upload/other/5f3d18419a00b60032736ebf4b4b6a14.zip', 1, '2023-07-28 09:24:58', '2023-07-28 09:24:58', 'root', '');
INSERT INTO `t_app_version` VALUES (30, 3, 'v1.1.33', '/public/upload/other/d1edc834d9a8b3942613b15886893b33.zip', 1, '2023-07-29 02:21:04', '2023-07-29 02:21:04', 'root', '');
INSERT INTO `t_app_version` VALUES (31, 3, 'v1.1.34', '/public/upload/other/b2ddc5efbcf2e99c9923121f535d9623.zip', 1, '2023-07-29 07:22:17', '2023-07-29 07:22:17', 'root', '');
INSERT INTO `t_app_version` VALUES (32, 3, 'v1.1.35', '/public/upload/other/d523aaef310e096b0f5859662a2db681.zip', 1, '2023-07-29 07:46:47', '2023-07-29 07:46:47', 'root', '');
INSERT INTO `t_app_version` VALUES (33, 3, 'v1.1.37', '/public/upload/other/343d9a96cd45a67582bdd70792e14d53.zip', 1, '2023-07-29 09:22:39', '2023-07-29 09:22:39', 'root', '');
INSERT INTO `t_app_version` VALUES (34, 3, 'v1.1.38', '/public/upload/other/f21c0b8803d13570e42f57fe5d8bc960.zip', 1, '2023-07-31 01:24:13', '2023-07-31 01:24:13', 'root', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推广渠道表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_channel
-- ----------------------------
INSERT INTO `t_channel` VALUES (1, '中国', 'china', 'http://127.0.0.1', 1, '2023-06-30 16:14:01', '2023-06-30 16:14:01', 'root', '');
INSERT INTO `t_channel` VALUES (2, '俄罗斯', 'russia', 'http://localhost', 1, '2023-06-30 16:14:23', '2023-06-30 16:14:23', 'root', '');
INSERT INTO `t_channel` VALUES (3, '英国', 'english', 'locahost', 1, '2023-07-29 07:26:34', '2023-07-29 07:26:34', 'root', '');

-- ----------------------------
-- Table structure for t_dev
-- ----------------------------
DROP TABLE IF EXISTS `t_dev`;
CREATE TABLE `t_dev`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `client_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户端自身设备ID',
  `os` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户端设备系统os',
  `is_send` int NULL DEFAULT NULL COMMENT '1-已赠送时间；2-未赠送',
  `network` int NULL DEFAULT 1 COMMENT '网络模式（1-自动；2-手动）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1685855271462637569 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '客户端设备表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_dev
-- ----------------------------
INSERT INTO `t_dev` VALUES (1677211938733428736, NULL, 'PostmanRuntime/7.32.2', 1, 1, '2023-07-07 15:04:21', '2023-07-07 15:04:21', '');
INSERT INTO `t_dev` VALUES (1677211938733428737, NULL, 'PostmanRuntime/7.32.3', 1, 1, '2023-07-07 17:08:37', '2023-07-07 17:08:40', NULL);
INSERT INTO `t_dev` VALUES (1681561840787656704, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-19 07:09:19', '2023-07-19 07:09:19', '');
INSERT INTO `t_dev` VALUES (1681570017864323072, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-19 07:41:48', '2023-07-19 07:41:48', '');
INSERT INTO `t_dev` VALUES (1681571524739338240, NULL, 'CPU iPhone OS 13_2_3 like Mac OS X', 1, 1, '2023-07-19 07:47:48', '2023-07-19 07:47:48', '');
INSERT INTO `t_dev` VALUES (1681630009988485120, NULL, 'CPU iPhone OS 14_2 like Mac OS X', 1, 1, '2023-07-19 11:40:12', '2023-07-19 11:40:12', '');
INSERT INTO `t_dev` VALUES (1681655699433590784, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-19 13:22:17', '2023-07-19 13:22:17', '');
INSERT INTO `t_dev` VALUES (1681855058557276160, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-20 02:34:27', '2023-07-20 02:34:27', '');
INSERT INTO `t_dev` VALUES (1681867380034113536, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-20 03:23:25', '2023-07-20 03:23:25', '');
INSERT INTO `t_dev` VALUES (1681867930939166720, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-20 03:25:36', '2023-07-20 03:25:36', '');
INSERT INTO `t_dev` VALUES (1681869950098083840, NULL, 'CPU iPhone OS 14_2 like Mac OS X', 1, 1, '2023-07-20 03:33:38', '2023-07-20 03:33:38', '');
INSERT INTO `t_dev` VALUES (1681877802594340864, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-20 04:04:50', '2023-07-20 04:04:50', '');
INSERT INTO `t_dev` VALUES (1681940927377051648, NULL, 'Android 9', 1, 1, '2023-07-20 08:15:40', '2023-07-20 08:15:40', '');
INSERT INTO `t_dev` VALUES (1682230987523624960, NULL, 'Android 9', 1, 1, '2023-07-21 03:28:16', '2023-07-21 03:28:16', '');
INSERT INTO `t_dev` VALUES (1682326372267069440, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-21 09:47:17', '2023-07-21 09:47:17', '');
INSERT INTO `t_dev` VALUES (1682557440404492288, NULL, 'Windows 10', 1, 1, '2023-07-22 01:05:28', '2023-07-22 01:05:28', '');
INSERT INTO `t_dev` VALUES (1682574427012730880, NULL, 'Intel Mac OS X 10_15_7', 1, 1, '2023-07-22 02:12:58', '2023-07-22 02:12:58', '');
INSERT INTO `t_dev` VALUES (1683322128373387264, NULL, 'CPU iPhone OS 16_6 like Mac OS X', 1, 1, '2023-07-24 03:44:04', '2023-07-24 03:44:04', '');
INSERT INTO `t_dev` VALUES (1683359980536729600, NULL, 'CPU iPhone OS 16_6 like Mac OS X', 1, 1, '2023-07-24 06:14:29', '2023-07-24 06:14:29', '');
INSERT INTO `t_dev` VALUES (1683361531191889920, NULL, 'CPU iPhone OS 16_6 like Mac OS X', 1, 1, '2023-07-24 06:20:39', '2023-07-24 06:20:39', '');
INSERT INTO `t_dev` VALUES (1683363454737453056, NULL, 'Android 9', 1, 1, '2023-07-24 06:28:17', '2023-07-24 06:28:17', '');
INSERT INTO `t_dev` VALUES (1683399907219607552, NULL, 'CPU iPhone OS 16_6 like Mac OS X', 1, 1, '2023-07-24 08:53:08', '2023-07-24 08:53:08', '');
INSERT INTO `t_dev` VALUES (1683404824428679168, NULL, 'Android 9', 1, 1, '2023-07-24 09:12:40', '2023-07-24 09:12:40', '');
INSERT INTO `t_dev` VALUES (1683405443369537536, NULL, 'Android 9', 1, 1, '2023-07-24 09:15:08', '2023-07-24 09:15:08', '');
INSERT INTO `t_dev` VALUES (1683406721491406848, NULL, 'Android 12', 1, 1, '2023-07-24 09:20:13', '2023-07-24 09:20:13', '');
INSERT INTO `t_dev` VALUES (1683407562147368960, NULL, 'Android 13', 1, 1, '2023-07-24 09:23:33', '2023-07-24 09:23:33', '');
INSERT INTO `t_dev` VALUES (1683408308884475904, NULL, 'Android 13', 1, 1, '2023-07-24 09:26:31', '2023-07-24 09:26:31', '');
INSERT INTO `t_dev` VALUES (1683409632028004352, NULL, 'Android 13', 1, 1, '2023-07-24 09:31:47', '2023-07-24 09:31:47', '');
INSERT INTO `t_dev` VALUES (1683409637799366656, NULL, 'Android 13', 1, 1, '2023-07-24 09:31:48', '2023-07-24 09:31:48', '');
INSERT INTO `t_dev` VALUES (1683409685522157568, NULL, 'Android 10', 1, 1, '2023-07-24 09:31:59', '2023-07-24 09:31:59', '');
INSERT INTO `t_dev` VALUES (1683411282608263168, NULL, 'Android 11', 1, 1, '2023-07-24 09:38:20', '2023-07-24 09:38:20', '');
INSERT INTO `t_dev` VALUES (1683422926776307712, NULL, 'Android 11', 1, 1, '2023-07-24 10:24:36', '2023-07-24 10:24:36', '');
INSERT INTO `t_dev` VALUES (1683426939466944512, NULL, 'CPU iPhone OS 16_6 like Mac OS X', 1, 1, '2023-07-24 10:40:33', '2023-07-24 10:40:33', '');
INSERT INTO `t_dev` VALUES (1683434186548973568, NULL, 'Android 13', 1, 1, '2023-07-24 11:09:21', '2023-07-24 11:09:21', '');
INSERT INTO `t_dev` VALUES (1683512069950803968, NULL, 'Intel Mac OS X 10_15_7', 1, 1, '2023-07-24 16:18:50', '2023-07-24 16:18:50', '');
INSERT INTO `t_dev` VALUES (1683650502971101184, NULL, 'Android 9', 1, 1, '2023-07-25 01:28:55', '2023-07-25 01:28:55', '');
INSERT INTO `t_dev` VALUES (1683660611679948800, NULL, 'CPU iPhone OS 16_6 like Mac OS X', 1, 1, '2023-07-25 02:09:05', '2023-07-25 02:09:05', '');
INSERT INTO `t_dev` VALUES (1683671523908390912, NULL, 'Android 9', 1, 1, '2023-07-25 02:52:27', '2023-07-25 02:52:27', '');
INSERT INTO `t_dev` VALUES (1683671936527241216, NULL, 'Android 9', 1, 1, '2023-07-25 02:54:05', '2023-07-25 02:54:05', '');
INSERT INTO `t_dev` VALUES (1683674043779125248, NULL, 'Android 13', 1, 1, '2023-07-25 03:02:27', '2023-07-25 03:02:27', '');
INSERT INTO `t_dev` VALUES (1683678852921954304, NULL, 'Android 13', 1, 1, '2023-07-25 03:21:34', '2023-07-25 03:21:34', '');
INSERT INTO `t_dev` VALUES (1683679854756630528, NULL, 'Android 12', 1, 1, '2023-07-25 03:25:33', '2023-07-25 03:25:33', '');
INSERT INTO `t_dev` VALUES (1683681627332415488, NULL, 'Android 13', 1, 1, '2023-07-25 03:32:35', '2023-07-25 03:32:35', '');
INSERT INTO `t_dev` VALUES (1683681627907035136, NULL, 'Android 13', 1, 1, '2023-07-25 03:32:36', '2023-07-25 03:32:36', '');
INSERT INTO `t_dev` VALUES (1683687695353647104, NULL, 'Android 13', 1, 1, '2023-07-25 03:56:42', '2023-07-25 03:56:42', '');
INSERT INTO `t_dev` VALUES (1683703053024235520, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-25 04:57:44', '2023-07-25 04:57:44', '');
INSERT INTO `t_dev` VALUES (1683703712901500928, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-25 05:00:21', '2023-07-25 05:00:21', '');
INSERT INTO `t_dev` VALUES (1683727763741085696, NULL, 'Android 9', 1, 1, '2023-07-25 06:35:55', '2023-07-25 06:35:55', '');
INSERT INTO `t_dev` VALUES (1683729239741829120, NULL, 'Android 11', 1, 1, '2023-07-25 06:41:47', '2023-07-25 06:41:47', '');
INSERT INTO `t_dev` VALUES (1683736170959212544, NULL, 'Android 9', 1, 1, '2023-07-25 07:09:20', '2023-07-25 07:09:20', '');
INSERT INTO `t_dev` VALUES (1683736718768869376, NULL, 'Android 13', 1, 1, '2023-07-25 07:11:30', '2023-07-25 07:11:30', '');
INSERT INTO `t_dev` VALUES (1683737355854286848, NULL, 'Android 10', 1, 1, '2023-07-25 07:14:02', '2023-07-25 07:14:02', '');
INSERT INTO `t_dev` VALUES (1683743613386756096, NULL, 'Android 13', 1, 1, '2023-07-25 07:38:54', '2023-07-25 07:38:54', '');
INSERT INTO `t_dev` VALUES (1683746051959296000, NULL, 'Android 13', 1, 1, '2023-07-25 07:48:35', '2023-07-25 07:48:35', '');
INSERT INTO `t_dev` VALUES (1683756993799524352, NULL, 'Android 13', 1, 1, '2023-07-25 08:32:04', '2023-07-25 08:32:04', '');
INSERT INTO `t_dev` VALUES (1683777087724326912, NULL, 'CPU iPhone OS 16_0_3 like Mac OS X', 1, 1, '2023-07-25 09:51:55', '2023-07-25 09:51:55', '');
INSERT INTO `t_dev` VALUES (1683783152318812160, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-25 10:16:01', '2023-07-25 10:16:01', '');
INSERT INTO `t_dev` VALUES (1683808014429065216, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-25 11:54:48', '2023-07-25 11:54:48', '');
INSERT INTO `t_dev` VALUES (1683827203457945600, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-25 13:11:03', '2023-07-25 13:11:03', '');
INSERT INTO `t_dev` VALUES (1683839055315341312, NULL, 'CPU iPhone OS 14_2 like Mac OS X', 1, 1, '2023-07-25 13:58:09', '2023-07-25 13:58:09', '');
INSERT INTO `t_dev` VALUES (1684021330074144768, NULL, 'Android 12', 1, 1, '2023-07-26 02:02:27', '2023-07-26 02:02:27', '');
INSERT INTO `t_dev` VALUES (1684026378766258176, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-26 02:22:31', '2023-07-26 02:22:31', '');
INSERT INTO `t_dev` VALUES (1684092087143763968, NULL, 'Android 13', 1, 1, '2023-07-26 06:43:37', '2023-07-26 06:43:37', '');
INSERT INTO `t_dev` VALUES (1684111615957405696, NULL, 'Android 13', 1, 1, '2023-07-26 08:01:13', '2023-07-26 08:01:13', '');
INSERT INTO `t_dev` VALUES (1684373826332266496, NULL, 'Android 13', 1, 1, '2023-07-27 01:23:09', '2023-07-27 01:23:09', '');
INSERT INTO `t_dev` VALUES (1684389925358669824, NULL, 'Android 13', 1, 1, '2023-07-27 02:27:07', '2023-07-27 02:27:07', '');
INSERT INTO `t_dev` VALUES (1684464079315406848, NULL, 'CPU iPhone OS 15_7_4 like Mac OS X', 1, 1, '2023-07-27 07:21:46', '2023-07-27 07:21:46', '');
INSERT INTO `t_dev` VALUES (1684484081770827776, NULL, 'CPU iPhone OS 16_5_1 like Mac OS X', 1, 1, '2023-07-27 08:41:15', '2023-07-27 08:41:15', '');
INSERT INTO `t_dev` VALUES (1684487002629607424, NULL, 'Android 12', 1, 1, '2023-07-27 08:52:52', '2023-07-27 08:52:52', '');
INSERT INTO `t_dev` VALUES (1684496326865195008, NULL, 'Android 13', 1, 1, '2023-07-27 09:29:55', '2023-07-27 09:29:55', '');
INSERT INTO `t_dev` VALUES (1684501142672773120, NULL, 'Android 10', 1, 1, '2023-07-27 09:49:03', '2023-07-27 09:49:03', '');
INSERT INTO `t_dev` VALUES (1684521730258767872, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-27 11:10:52', '2023-07-27 11:10:52', '');
INSERT INTO `t_dev` VALUES (1684522279972638720, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-27 11:13:03', '2023-07-27 11:13:03', '');
INSERT INTO `t_dev` VALUES (1684525565203189760, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-27 11:26:06', '2023-07-27 11:26:06', '');
INSERT INTO `t_dev` VALUES (1684753295664484352, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-28 02:31:01', '2023-07-28 02:31:01', '');
INSERT INTO `t_dev` VALUES (1684753436886700032, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-28 02:31:35', '2023-07-28 02:31:35', '');
INSERT INTO `t_dev` VALUES (1684759095699050496, NULL, 'CPU iPhone OS 14_2 like Mac OS X', 1, 1, '2023-07-28 02:54:04', '2023-07-28 02:54:04', '');
INSERT INTO `t_dev` VALUES (1684803646232989696, NULL, 'CPU iPhone OS 16_3_1 like Mac OS X', 1, 1, '2023-07-28 05:51:06', '2023-07-28 05:51:06', '');
INSERT INTO `t_dev` VALUES (1684807658973958144, NULL, 'CPU iPhone OS 16_0_3 like Mac OS X', 1, 1, '2023-07-28 06:07:02', '2023-07-28 06:07:02', '');
INSERT INTO `t_dev` VALUES (1684810635877027840, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-28 06:18:52', '2023-07-28 06:18:52', '');
INSERT INTO `t_dev` VALUES (1684814428836466688, NULL, 'CPU iPhone OS 15_7_4 like Mac OS X', 1, 1, '2023-07-28 06:33:56', '2023-07-28 06:33:56', '');
INSERT INTO `t_dev` VALUES (1684815749710876672, NULL, 'CPU iPhone OS 16_0_3 like Mac OS X', 1, 1, '2023-07-28 06:39:11', '2023-07-28 06:39:11', '');
INSERT INTO `t_dev` VALUES (1684819516325892096, NULL, 'Android 12', 1, 1, '2023-07-28 06:54:09', '2023-07-28 06:54:09', '');
INSERT INTO `t_dev` VALUES (1684823808927600640, NULL, 'CPU iPhone OS 16_0_3 like Mac OS X', 1, 1, '2023-07-28 07:11:13', '2023-07-28 07:11:13', '');
INSERT INTO `t_dev` VALUES (1684824801509642240, NULL, 'Android 13', 1, 1, '2023-07-28 07:15:09', '2023-07-28 07:15:09', '');
INSERT INTO `t_dev` VALUES (1684838129904652288, NULL, 'Android 13', 1, 1, '2023-07-28 08:08:07', '2023-07-28 08:08:07', '');
INSERT INTO `t_dev` VALUES (1684839812965601280, NULL, 'Android 13', 1, 1, '2023-07-28 08:14:48', '2023-07-28 08:14:48', '');
INSERT INTO `t_dev` VALUES (1684968511522213888, NULL, 'CPU iPhone OS 14_8_1 like Mac OS X', 1, 1, '2023-07-28 16:46:13', '2023-07-28 16:46:13', '');
INSERT INTO `t_dev` VALUES (1684968714077736960, NULL, 'CPU iPhone OS 16_0_3 like Mac OS X', 1, 1, '2023-07-28 16:47:01', '2023-07-28 16:47:01', '');
INSERT INTO `t_dev` VALUES (1684969498777489408, NULL, 'CPU iPhone OS 16_0_3 like Mac OS X', 1, 1, '2023-07-28 16:50:08', '2023-07-28 16:50:08', '');
INSERT INTO `t_dev` VALUES (1685114917293658112, NULL, 'Intel Mac OS X 10_15_7', 1, 1, '2023-07-29 02:27:58', '2023-07-29 02:27:58', '');
INSERT INTO `t_dev` VALUES (1685285207160131584, '', 'CPU iPhone OS 16_3_1 like Mac OS X', 2, 1, '2023-07-29 13:44:39', '2023-07-29 13:44:39', '');
INSERT INTO `t_dev` VALUES (1685288035249295360, '', 'CPU iPhone OS 16_3_1 like Mac OS X', 2, 1, '2023-07-29 13:55:53', '2023-07-29 13:55:53', '');
INSERT INTO `t_dev` VALUES (1685291172928425984, '', 'CPU iPhone OS 16_3_1 like Mac OS X', 2, 1, '2023-07-29 14:08:21', '2023-07-29 14:08:21', '');
INSERT INTO `t_dev` VALUES (1685831043220770816, '', 'Android 13', 2, 1, '2023-07-31 01:53:36', '2023-07-31 01:53:36', '');
INSERT INTO `t_dev` VALUES (1685833743207501824, '', 'CPU iPhone OS 16_3_1 like Mac OS X', 2, 1, '2023-07-31 02:04:20', '2023-07-31 02:04:20', '');
INSERT INTO `t_dev` VALUES (1685834344431620096, '', 'CPU iPhone OS 16_3_1 like Mac OS X', 2, 1, '2023-07-31 02:06:43', '2023-07-31 02:06:43', '');
INSERT INTO `t_dev` VALUES (1685840833791660032, '', 'CPU iPhone OS 16_3_1 like Mac OS X', 2, 1, '2023-07-31 02:32:30', '2023-07-31 02:32:30', '');
INSERT INTO `t_dev` VALUES (1685854340339732480, '', 'Android 13', 2, 1, '2023-07-31 03:26:11', '2023-07-31 03:26:11', '');
INSERT INTO `t_dev` VALUES (1685855271462637568, '', 'Android 13', 2, 1, '2023-07-31 03:29:53', '2023-07-31 03:29:53', '');

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_dict
-- ----------------------------
INSERT INTO `t_dict` VALUES ('app_js_zip', 'http://localhost/jszip', 'app静态包链接', 0, '2023-07-07 17:19:57', '2023-07-07 17:58:27');
INSERT INTO `t_dict` VALUES ('app_link', 'http://localhost/download1', 'app下载链接', 0, '2023-06-16 19:11:40', '2023-07-17 19:56:24');
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
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '赠送用户记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_gift
-- ----------------------------
INSERT INTO `t_gift` VALUES (1, 219122279, '1681630800195358720', 219122279, '免费领会员', 3600, 3, '2023-07-19 11:43:20', '2023-07-19 11:43:20', '');
INSERT INTO `t_gift` VALUES (2, 219122280, '1681890360827056128', 219122280, '免费领会员', 3600, 3, '2023-07-20 04:54:44', '2023-07-20 04:54:44', '');
INSERT INTO `t_gift` VALUES (3, 219122279, '1681932501020315648', 219122279, '免费领会员', 3600, 3, '2023-07-20 07:42:11', '2023-07-20 07:42:11', '');
INSERT INTO `t_gift` VALUES (4, 219122279, '1682586131201265664', 219122279, '免费领会员', 3600, 3, '2023-07-22 02:59:29', '2023-07-22 02:59:29', '');
INSERT INTO `t_gift` VALUES (5, 219122277, '1683100674218266624', 219122277, '免费领会员', 3600, 3, '2023-07-23 13:04:05', '2023-07-23 13:04:05', '');
INSERT INTO `t_gift` VALUES (6, 219122279, '1683288201172619264', 219122279, '免费领会员', 3600, 3, '2023-07-24 01:29:15', '2023-07-24 01:29:15', '');
INSERT INTO `t_gift` VALUES (7, 219122279, '1683727644723515392', 219122279, '免费领会员', 3600, 3, '2023-07-25 06:35:27', '2023-07-25 06:35:27', '');
INSERT INTO `t_gift` VALUES (8, 219122276, '1683731957860536320', 219122276, '免费领会员', 3600, 3, '2023-07-25 06:52:35', '2023-07-25 06:52:35', '');
INSERT INTO `t_gift` VALUES (9, 219122277, '1683732910047236096', 219122277, '免费领会员', 3600, 3, '2023-07-25 06:56:22', '2023-07-25 06:56:22', '');
INSERT INTO `t_gift` VALUES (10, 219122284, '1683733513477558272', 219122284, '免费领会员', 3600, 3, '2023-07-25 06:58:46', '2023-07-25 06:58:46', '');
INSERT INTO `t_gift` VALUES (11, 219122282, '1683799043962048512', 219122282, '免费领会员', 3600, 3, '2023-07-25 11:19:10', '2023-07-25 11:19:10', '');
INSERT INTO `t_gift` VALUES (12, 219122277, '1684130886091542528', 219122277, '免费领会员', 3600, 3, '2023-07-26 09:17:47', '2023-07-26 09:17:47', '');
INSERT INTO `t_gift` VALUES (13, 219122284, '1684132709233856512', 219122284, '免费领会员', 3600, 3, '2023-07-26 09:25:02', '2023-07-26 09:25:02', '');
INSERT INTO `t_gift` VALUES (14, 219122291, '1690363527', 219122291, '注册赠送', 3600, 1, '2023-07-26 09:25:27', '2023-07-26 09:25:27', '');
INSERT INTO `t_gift` VALUES (15, 219122292, '1690364060', 219122292, '注册赠送', 3600, 1, '2023-07-26 09:34:20', '2023-07-26 09:34:20', '');
INSERT INTO `t_gift` VALUES (16, 219122293, '1690424856', 219122293, '注册赠送', 3600, 1, '2023-07-27 02:27:36', '2023-07-27 02:27:36', '');
INSERT INTO `t_gift` VALUES (17, 219122282, '1684404379735560192', 219122282, '免费领会员', 3600, 3, '2023-07-27 03:24:33', '2023-07-27 03:24:33', '');
INSERT INTO `t_gift` VALUES (18, 219122282, '1684468325079322624', 219122282, '免费领会员', 3600, 3, '2023-07-27 07:38:39', '2023-07-27 07:38:39', '');
INSERT INTO `t_gift` VALUES (19, 219122295, '1690450218', 219122295, '注册赠送', 3600, 1, '2023-07-27 09:30:18', '2023-07-27 09:30:18', '');
INSERT INTO `t_gift` VALUES (20, 219122296, '1690451003', 219122296, '注册赠送', 3600, 1, '2023-07-27 09:43:23', '2023-07-27 09:43:23', '');
INSERT INTO `t_gift` VALUES (21, 219122297, '1690451412', 219122297, '注册赠送', 3600, 1, '2023-07-27 09:50:12', '2023-07-27 09:50:12', '');
INSERT INTO `t_gift` VALUES (22, 219122298, '1690452676', 219122298, '注册赠送', 3600, 1, '2023-07-27 10:11:16', '2023-07-27 10:11:16', '');
INSERT INTO `t_gift` VALUES (23, 219122299, '1690456431', 219122299, '注册赠送', 3600, 1, '2023-07-27 11:13:51', '2023-07-27 11:13:51', '');
INSERT INTO `t_gift` VALUES (24, 219122282, '1684542307329642496', 219122282, '免费领会员', 3600, 3, '2023-07-27 12:32:38', '2023-07-27 12:32:38', '');
INSERT INTO `t_gift` VALUES (25, 219122300, '1690507604', 219122300, '注册赠送', 3600, 1, '2023-07-28 01:26:44', '2023-07-28 01:26:44', '');
INSERT INTO `t_gift` VALUES (26, 219122302, '1690525513', 219122302, '注册赠送', 3600, 1, '2023-07-28 06:25:13', '2023-07-28 06:25:13', '');
INSERT INTO `t_gift` VALUES (27, 219122303, '1690528537', 219122303, '注册赠送', 3600, 1, '2023-07-28 07:15:37', '2023-07-28 07:15:37', '');
INSERT INTO `t_gift` VALUES (28, 219122304, '1690532117', 219122304, '注册赠送', 3600, 1, '2023-07-28 08:15:17', '2023-07-28 08:15:17', '');
INSERT INTO `t_gift` VALUES (29, 219122282, '1684900680835272704', 219122282, '免费领会员', 3600, 3, '2023-07-28 12:16:40', '2023-07-28 12:16:40', '');
INSERT INTO `t_gift` VALUES (30, 219122305, '1690562831', 219122305, '注册赠送', 3600, 1, '2023-07-28 16:47:11', '2023-07-28 16:47:11', '');
INSERT INTO `t_gift` VALUES (31, 219122306, '1690600996', 219122306, '注册赠送', 3600, 1, '2023-07-29 03:23:16', '2023-07-29 03:23:16', '');

-- ----------------------------
-- Table structure for t_goods
-- ----------------------------
DROP TABLE IF EXISTS `t_goods`;
CREATE TABLE `t_goods`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `m_type` int NOT NULL COMMENT '会员类型：1-vip1；2-vip2',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '套餐标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '套餐标题（英文）',
  `title_rus` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '套餐标题（俄文）',
  `price` decimal(10, 6) NOT NULL COMMENT '单价(U)',
  `period` int NOT NULL COMMENT '有效期（天）',
  `dev_limit` int NOT NULL COMMENT '设备限制数',
  `flow_limit` bigint NOT NULL COMMENT '流量限制数；单位：字节；0-不限制',
  `is_discount` int NULL DEFAULT NULL COMMENT '是否优惠:1-是；2-否',
  `low` int NULL DEFAULT NULL COMMENT '最低赠送(天)',
  `high` int NULL DEFAULT NULL COMMENT '最高赠送(天)',
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
INSERT INTO `t_goods` VALUES (1, 1, '1周', '', NULL, 2.000000, 7, 2, 2000000000, 2, NULL, NULL, 1, '2023-06-30 16:15:30', '2023-07-20 11:51:00', 'root', '');
INSERT INTO `t_goods` VALUES (2, 1, '1个月', '', NULL, 5.000000, 30, 2, 200000000, 2, NULL, NULL, 1, '2023-06-30 16:18:04', '2023-07-20 11:50:55', 'root', '');
INSERT INTO `t_goods` VALUES (3, 1, '3个月', '', NULL, 11.000000, 90, 2, 3000000000, 2, NULL, NULL, 1, '2023-06-30 16:19:37', '2023-07-20 11:50:50', 'root', '');
INSERT INTO `t_goods` VALUES (4, 1, '12个月', '', NULL, 40.000000, 365, 5, 9000000000000, 2, NULL, NULL, 1, '2023-06-30 16:20:15', '2023-07-20 11:50:46', 'root', '');
INSERT INTO `t_goods` VALUES (5, 2, '1周', '', NULL, 3.000000, 7, 2, 500000000000, 1, 1, 3, 1, '2023-06-30 16:23:25', '2023-07-20 11:50:40', 'root', '');
INSERT INTO `t_goods` VALUES (6, 2, '1个月', '', NULL, 7.000000, 30, 3, 5000000000000, 1, 3, 5, 1, '2023-06-30 16:24:01', '2023-07-20 11:50:33', 'root', '');
INSERT INTO `t_goods` VALUES (7, 2, '3个月', '', NULL, 16.000000, 90, 5, 5000000000000000, 1, 10, 15, 1, '2023-06-30 16:24:50', '2023-07-20 11:50:27', 'root', '');
INSERT INTO `t_goods` VALUES (8, 2, '12个月', '', NULL, 54.000000, 365, 5, 5000000000000000, 1, 30, 60, 1, '2023-06-30 16:25:25', '2023-07-20 11:50:19', 'root', '');

-- ----------------------------
-- Table structure for t_ios_account
-- ----------------------------
DROP TABLE IF EXISTS `t_ios_account`;
CREATE TABLE `t_ios_account`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号id',
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ios账号',
  `pass` varchar(64) CHARACTER SET utf16 COLLATE utf16_general_ci NOT NULL COMMENT '密码',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '别名',
  `account_type` int NULL DEFAULT NULL COMMENT '1-国区；2-海外',
  `status` int NULL DEFAULT NULL COMMENT '1-正常；2-下架',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `ios_account`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'ios账号管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_ios_account
-- ----------------------------
INSERT INTO `t_ios_account` VALUES (1, 'a123', 'a123', 'aaa', 2, 1, '2023-07-17 19:34:13', '2023-07-17 19:34:13', 'root', '');
INSERT INTO `t_ios_account` VALUES (2, 'b123', 'b123', 'bbb', 1, 1, '2023-07-17 19:36:14', '2023-07-17 19:36:14', 'root', '');

-- ----------------------------
-- Table structure for t_node
-- ----------------------------
DROP TABLE IF EXISTS `t_node`;
CREATE TABLE `t_node`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点名称',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点标题（英文）',
  `title_rus` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点标题（俄文)',
  `country` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '国家',
  `country_en` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '国家（英文）',
  `country_rus` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '国家（俄文)',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内网IP',
  `server` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公网域名',
  `node_type` int NULL DEFAULT NULL COMMENT '节点类别:1-常规；2-高带宽...(根据情况而定)',
  `port` int NOT NULL COMMENT '公网端口',
  `cpu` int NOT NULL COMMENT 'cpu核数量（单位个）',
  `flow` bigint NOT NULL COMMENT '流量带宽',
  `disk` bigint NOT NULL COMMENT '磁盘容量（单位B）',
  `memory` bigint NOT NULL COMMENT '内存大小（单位B）',
  `min_port` int NULL DEFAULT NULL COMMENT '最小端口',
  `max_port` int NULL DEFAULT NULL COMMENT '最大端口',
  `path` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ws路径',
  `is_recommend` int NULL DEFAULT NULL COMMENT '推荐节点1-是；2-否',
  `channel_id` int NULL DEFAULT NULL COMMENT '市场渠道（默认0）-优选节点有效',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1003 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '节点表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_node
-- ----------------------------
INSERT INTO `t_node` VALUES (1001, '日本线路1', '日本VIP1', '', NULL, '日本', '', NULL, '104.233.171.69', '10.10.10.111', 1, 443, 0, 0, 0, 0, 13001, 13005, '/work', 2, NULL, 1, '2023-07-01 10:13:09', '2023-07-21 12:37:48', 'root', '');
INSERT INTO `t_node` VALUES (1002, '香港线路1', '香港VIP1', '', NULL, '香港', '', NULL, '107.148.239.239', '10.10.10.111', 1, 443, 0, 0, 0, 0, 13001, 13005, '/work', 1, NULL, 1, '2023-07-01 10:13:09', '2023-07-21 12:37:42', 'root', '');

-- ----------------------------
-- Table structure for t_node_dns
-- ----------------------------
DROP TABLE IF EXISTS `t_node_dns`;
CREATE TABLE `t_node_dns`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `node_id` bigint NOT NULL COMMENT '节点id',
  `dns` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '域名',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'ip地址',
  `level` int NOT NULL COMMENT '线路级别:1,2,3...用于白名单机制',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '节点域名表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_node_dns
-- ----------------------------
INSERT INTO `t_node_dns` VALUES (1, 1001, 'node1.goudd.xyz', '104.233.171.69', 2, 1, '2023-07-18 17:07:45', '2023-07-18 19:12:01', 'root', '');
INSERT INTO `t_node_dns` VALUES (2, 1001, 'node1.wuwuwu360.xyz', '104.233.171.69', 2, 1, '2023-07-18 17:08:29', '2023-07-18 19:11:46', 'root', '');
INSERT INTO `t_node_dns` VALUES (3, 1002, 'node2.wuwuwu360.xyz', '107.148.239.239', 1, 1, '2023-07-18 17:57:28', '2023-07-18 19:14:31', 'root', '');
INSERT INTO `t_node_dns` VALUES (4, 1002, 'node2.thankw.xyz', '107.148.239.240', 1, 1, '2023-07-18 17:57:53', '2023-07-18 19:14:53', 'root', '');
INSERT INTO `t_node_dns` VALUES (5, 1001, 'node1.thankw.xyz', '104.233.171.69', 1, 1, '2023-07-18 18:01:26', '2023-07-18 18:01:26', 'root', '');
INSERT INTO `t_node_dns` VALUES (6, 1002, 'node2.goudd.xyz', '107.148.239.241', 2, 1, '2023-07-18 19:15:27', '2023-07-18 19:15:27', 'root', '');

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
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题（英文）',
  `title_rus` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题（俄文）',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签',
  `tag_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签（英文）',
  `tag_rus` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标签（俄文）',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '正文内容',
  `content_en` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '正文内容（英文）',
  `content_rus` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '正文内容（俄文）',
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
INSERT INTO `t_notice` VALUES (1, '欢迎新用户', NULL, NULL, '新用户,注册', NULL, NULL, '欢迎新用户来到本平台', NULL, NULL, 'root', '2023-06-16 17:07:17', '2023-06-16 17:08:22', 1, '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1684565700770795521 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '商品订单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_order
-- ----------------------------
INSERT INTO `t_order` VALUES (1677241462653194240, 219122276, 1, 'vip1普通会员1周', 2.000000, 2.00, 1, NULL, '2023-07-07 17:01:40', '2023-07-07 17:01:40', '');
INSERT INTO `t_order` VALUES (1681914560715427840, 219122277, 3, 'vip1普通会员3个月', 11.000000, 11.00, 1, NULL, '2023-07-20 06:30:54', '2023-07-20 06:30:54', '');
INSERT INTO `t_order` VALUES (1681983053699747840, 219122279, 7, 'vip2超级会员3个月', 16.000000, 16.00, 1, NULL, '2023-07-20 11:03:04', '2023-07-20 11:03:04', '');
INSERT INTO `t_order` VALUES (1681986037007519744, 219122280, 8, 'vip2超级会员12个月', 54.000000, 54.00, 1, NULL, '2023-07-20 11:14:55', '2023-07-20 11:14:55', '');
INSERT INTO `t_order` VALUES (1681986070180270080, 219122280, 3, 'vip1普通会员3个月', 11.000000, 11.00, 1, NULL, '2023-07-20 11:15:03', '2023-07-20 11:15:03', '');
INSERT INTO `t_order` VALUES (1681986085393010688, 219122280, 8, 'vip2超级会员12个月', 54.000000, 54.00, 1, NULL, '2023-07-20 11:15:07', '2023-07-20 11:15:07', '');
INSERT INTO `t_order` VALUES (1682330513722839040, 219122279, 1, '1周', 2.000000, 2.00, 1, NULL, '2023-07-21 10:03:45', '2023-07-21 10:03:45', '');
INSERT INTO `t_order` VALUES (1682330521457135616, 219122279, 1, '1周', 2.000000, 2.00, 1, NULL, '2023-07-21 10:03:47', '2023-07-21 10:03:47', '');
INSERT INTO `t_order` VALUES (1682330549940654080, 219122279, 1, '1周', 2.000000, 2.00, 1, NULL, '2023-07-21 10:03:53', '2023-07-21 10:03:53', '');
INSERT INTO `t_order` VALUES (1682330557641396224, 219122279, 5, '1周', 3.000000, 3.00, 1, NULL, '2023-07-21 10:03:55', '2023-07-21 10:03:55', '');
INSERT INTO `t_order` VALUES (1682357900263034880, 219122279, 5, '1周', 3.000000, 3.00, 1, NULL, '2023-07-21 11:52:34', '2023-07-21 11:52:34', '');
INSERT INTO `t_order` VALUES (1682357908861358080, 219122279, 5, '1周', 3.000000, 3.00, 1, NULL, '2023-07-21 11:52:36', '2023-07-21 11:52:36', '');
INSERT INTO `t_order` VALUES (1683038890283241472, 219122279, 8, '12个月', 54.000000, 54.00, 1, NULL, '2023-07-23 08:58:35', '2023-07-23 08:58:35', '');
INSERT INTO `t_order` VALUES (1683038908880785408, 219122279, 5, '1周', 3.000000, 3.00, 1, NULL, '2023-07-23 08:58:39', '2023-07-23 08:58:39', '');
INSERT INTO `t_order` VALUES (1683291275551313920, 219122279, 1, '1周', 2.000000, 2.00, 1, NULL, '2023-07-24 01:41:28', '2023-07-24 01:41:28', '');
INSERT INTO `t_order` VALUES (1683679940110716928, 219122279, 1, '1周', 2.000000, 2.00, 1, NULL, '2023-07-25 03:25:53', '2023-07-25 03:25:53', '');
INSERT INTO `t_order` VALUES (1683808263826575360, 219122282, 1, '1周', 2.000000, 2.00, 1, NULL, '2023-07-25 11:55:48', '2023-07-25 11:55:48', '');
INSERT INTO `t_order` VALUES (1683808287742496768, 219122282, 5, '1周', 3.000000, 3.00, 1, NULL, '2023-07-25 11:55:54', '2023-07-25 11:55:54', '');
INSERT INTO `t_order` VALUES (1683808299650125824, 219122282, 6, '1个月', 7.000000, 7.00, 1, NULL, '2023-07-25 11:55:56', '2023-07-25 11:55:56', '');
INSERT INTO `t_order` VALUES (1684565675428810752, 219122299, 8, '12个月', 54.000000, 54.00, 1, NULL, '2023-07-27 14:05:29', '2023-07-27 14:05:29', '');
INSERT INTO `t_order` VALUES (1684565700770795520, 219122299, 8, '12个月', 54.000000, 54.00, 1, NULL, '2023-07-27 14:05:35', '2023-07-27 14:05:35', '');

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
  `start_time` bigint NOT NULL COMMENT '本次计费开始时间戳',
  `end_time` bigint NOT NULL COMMENT '本次计费结束时间戳',
  `surplus_sec` bigint NOT NULL COMMENT '剩余时长(s)',
  `total_sec` bigint NULL DEFAULT NULL COMMENT '订单总时长(s）',
  `goods_day` int NULL DEFAULT NULL COMMENT '套餐天数',
  `send_day` int NULL DEFAULT NULL COMMENT '赠送天数',
  `pay_type` int NOT NULL COMMENT '订单状态:1-银行卡；2-支付宝；3-微信支付',
  `status` int NULL DEFAULT NULL COMMENT '1-using使用中；2-wait等待; 3-end已结束',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户购买成功记录' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '上传日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_upload_log
-- ----------------------------
INSERT INTO `t_upload_log` VALUES (1, 219122276, 1677211938733428736, 'aaa error', '2023-07-07 16:24:21', '2023-07-07 16:24:21', '');
INSERT INTO `t_upload_log` VALUES (2, 219122283, 1683650502971101184, '{\"log\":\"no log\"}', '2023-07-25 02:20:29', '2023-07-25 02:20:29', '');
INSERT INTO `t_upload_log` VALUES (3, 219122280, 1681877802594340864, '{\"log\":\"2023-07-24 20:00:53 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-24 20:00:54 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:00:56 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:00:59 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:02:53 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-24 20:02:56 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:02:58 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:03:16 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/EAB66917-4F00-4829-B8A2-A470A5E8D2F0/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-24 20:03:19 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/EAB66917-4F00-4829-B8A2-A470A5E8D2F0/Documents/web/vpn/index.html#/pages/country/index?type=1\\n2023-07-24 20:05:55 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-24 20:06:16 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connecting\\n2023-07-24 20:06:16 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connecting\\n2023-07-24 20:06:16 PacketTunnel/PacketTunnelProvider.swift:40:19 startTunnel(options:completionHandler:): start on main thread: false\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:161:19 start(completionHandler:): start on state stopped\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: stopped, net status: satisfied\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:273:23 processConfig(): \\n-------\\n[General]\\nloglevel = debug\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\nlogoutput = /private/var/mobile/Containers/Shared/AppGroup/2879856A-B8B2-4FE1-862C-9E69ADB3BF00/logs/leaf.log\\ntun-fd = 5\\nrouting-domain-resolve = true\\nalways-real-ip = tracker, apple.com\\n\\n[Proxy]\\nDirect = direct\\nReject = reject\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=/work, tls-insecure=true\\n\\n[Rule]\\nEXTERNAL, site:/private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/PlugIns/PacketTunnel.appex/site.dat:cn, Direct\\nEXTERNAL, mmdb:/private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/PlugIns/PacketTunnel.appex/geo.mmdb:cn, Direct\\nFINAL, VMessWSS\\n\\n-------\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to started\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:06:16 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 20:06:16 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:06:16 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:06:20 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:06:26 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:06:27 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:06:29 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:07:27 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:08:32 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:08:40 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:09:20 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:10:42 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:10:44 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:10:47 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:10:50 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:10:54 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:10:58 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:11:40 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:11:49 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:14:22 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:14:31 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:14:31 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:14:34 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:15:27 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:15:29 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:15:32 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:15:39 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:15:40 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 20:15:41 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:15:41 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:15:42 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:15:44 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:16:11 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:16:17 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:16:25 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:16:35 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:16:35 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:16:38 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:16:47 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:16:50 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:16:58 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:17:01 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:17:10 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:17:13 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:17:22 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:17:28 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:17:33 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:17:36 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:17:48 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:17:50 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:17:53 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:18:05 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:18:51 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:19:02 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:19:05 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:19:09 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:19:27 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:19:30 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:19:35 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:19:37 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:20:30 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:22:13 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:22:14 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:22:17 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:22:30 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:25:17 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:25:19 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:25:22 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:25:42 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:25:50 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:27:36 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:27:36 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:27:55 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:29:25 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:29:31 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:29:33 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:29:49 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:29:51 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:29:55 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:33:51 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:33:52 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:33:55 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:34:06 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:35:05 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:35:22 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:36:35 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:36:36 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:36:39 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:36:50 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:37:24 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:37:34 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:38:11 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:38:21 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:38:23 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:38:32 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:38:52 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:38:59 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:39:00 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:39:11 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:39:52 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:40:02 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:41:59 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:42:08 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:42:08 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:42:18 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:42:19 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:42:22 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:42:28 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:42:30 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:42:31 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 20:42:32 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:42:33 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:42:44 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-24 20:43:40 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:43:41 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:43:41 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:43:44 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:43:45 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:43:48 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:43:49 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:44:07 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:44:09 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:44:11 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:44:11 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:44:11 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:44:26 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:44:28 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:44:48 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:44:56 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:45:00 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:45:23 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:45:26 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:45:38 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:45:44 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:45:49 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:45:51 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:45:54 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:01 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:04 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:21 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:23 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:29 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:32 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:46:50 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:47:13 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:48:15 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:48:32 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:48:39 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:48:42 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:50:25 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:50:28 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:51:03 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:51:05 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:51:07 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:51:08 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:51:16 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 20:51:20 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/EAB66917-4F00-4829-B8A2-A470A5E8D2F0/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-24 20:51:29 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/EAB66917-4F00-4829-B8A2-A470A5E8D2F0/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-24 20:51:29 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:36 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:36 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:37 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:39 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:40 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:42 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:45 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:47 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:52:54 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:54:23 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:54:25 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:25 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:25 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:27 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:27 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:54:28 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:29 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:54:30 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:30 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:44 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:54:47 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 20:55:09 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:56:22 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:56:28 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:56:29 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:56:32 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:59:05 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 20:59:15 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 20:59:51 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 21:00:17 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:01:23 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:01:28 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:01:29 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-24 21:01:33 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-24 21:03:06 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 21:03:09 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 21:03:23 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 21:03:23 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:03:48 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:13:24 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 21:14:25 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 21:14:37 PacketTunnel/PacketTunnelProvider.swift:71:19 sleep(completionHandler:): sleep\\n2023-07-24 21:17:45 PacketTunnel/PacketTunnelProvider.swift:77:19 wake(): wake\\n2023-07-24 21:30:52 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 21:30:53 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:32:24 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 21:32:25 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:37:40 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 21:47:21 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-24 21:47:22 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 21:47:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnecting\\n2023-07-24 21:47:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnecting\\n2023-07-24 21:47:22 PacketTunnel/LeafAdapter.swift:222:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-24 21:47:22 PacketTunnel/PacketTunnelProvider.swift:60:19 stopTunnel(with:completionHandler:): stop\\n2023-07-24 21:47:22 PacketTunnel/LeafAdapter.swift:190:19 stop(completionHandler:): stop on state started\\n2023-07-24 21:47:22 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to stopped\\n2023-07-24 21:47:22 PacketTunnel/LeafAdapter.swift:179:27 start(completionHandler:): leaf shutdown on state stopped\\n2023-07-24 21:47:22 PacketTunnel/PacketTunnelProvider.swift:14:19 deinit: dealloc\\n2023-07-24 21:47:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-24 21:47:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-24 23:03:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-24 23:03:23 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/EAB66917-4F00-4829-B8A2-A470A5E8D2F0/Documents/web/vpn/index.html\\n2023-07-24 23:03:25 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed unknow\\n2023-07-24 23:03:47 JunJunAPP/WebViewController.swift:193:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/B914F5D3-5283-4675-A22F-8F00F5511C37/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-25 10:23:33 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 10:23:34 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip\\n2023-07-25 10:23:50 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 10:23:50 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip\\n2023-07-25 10:25:07 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip\\n2023-07-25 10:25:09 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/68C1411E-498B-45D4-AF52-FFA4F11DA6C2/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-25 10:25:09 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:25:09 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:25:09 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"data\\\":{\\\"loginInfo\\\":{\\\"id\\\":219122280,\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"surplus_flow\\\":0,\\\"expired_time\\\":0,\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"member_type\\\":0,\\\"lastDate\\\":\\\"已到期\\\"},\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"appVersion\\\":\\\"1.0\\\",\\\"deviceModal\\\":\\\"iPhone13,3\\\"},\\\"message\\\":\\\"OK\\\",\\\"code\\\":0}\\n2023-07-25 10:25:09 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:25:12 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {\\\"animated\\\":true}\\n2023-07-25 10:25:12 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{}}\\n2023-07-25 10:25:13 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:25:14 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/1F79C151-10B7-420B-BBC0-E44800A68669/Documents/web/vpn/index.html#/pages/country/index?type=1\\n2023-07-25 10:25:14 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:25:14 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:25:19 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pingHosts : {\\\"timeout\\\":6,\\\"hosts\\\":[\\\"10.10.10.111\\\",\\\"10.10.10.111\\\"]}\\n2023-07-25 10:25:25 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pingHosts : {\\\"data\\\":{\\\"10.10.10.111\\\":[{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"},{\\\"error\\\":\\\"timeout\\\"}]},\\\"code\\\":0,\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 10:25:37 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:26:39 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 10:26:40 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip\\n2023-07-25 10:26:40 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/68C1411E-498B-45D4-AF52-FFA4F11DA6C2/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-25 10:26:41 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:26:41 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:26:41 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{\\\"deviceModal\\\":\\\"iPhone13,3\\\",\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0},\\\"appVersion\\\":\\\"1.0\\\"}}\\n2023-07-25 10:26:41 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:26:43 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {\\\"animated\\\":true}\\n2023-07-25 10:26:43 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{}}\\n2023-07-25 10:26:43 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:26:45 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/1F79C151-10B7-420B-BBC0-E44800A68669/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 10:26:45 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:26:45 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:26:45 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"appVersion\\\":\\\"1.0\\\",\\\"deviceModal\\\":\\\"iPhone13,3\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0}}}\\n2023-07-25 10:26:45 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:26:46 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:26:47 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/1F79C151-10B7-420B-BBC0-E44800A68669/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 10:26:47 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:26:47 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:26:47 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"appVersion\\\":\\\"1.0\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0},\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"deviceModal\\\":\\\"iPhone13,3\\\",\\\"osVersion\\\":\\\"16.3.1\\\"},\\\"code\\\":0}\\n2023-07-25 10:26:48 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:26:54 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {}\\n2023-07-25 10:26:54 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 10:26:55 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:27:39 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip\\n2023-07-25 10:27:45 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///private/var/containers/Bundle/Application/68C1411E-498B-45D4-AF52-FFA4F11DA6C2/JunJunAPP.app/configs/bridgeTest.html\\n2023-07-25 10:27:45 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:27:45 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:27:45 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"code\\\":0,\\\"data\\\":{\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"appVersion\\\":\\\"1.0\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0},\\\"deviceModal\\\":\\\"iPhone13,3\\\"},\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 10:27:45 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:27:47 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.telnetHosts : {\\\"timeout\\\":3,\\\"hosts\\\":[\\\"www.google.com:443\\\",\\\"www.csdn.com:443\\\",\\\"node2.wuwuwu360.xyz:443\\\",\\\"104.233.171.69:9999\\\",\\\"104.233.171.69:8888\\\"]}\\n2023-07-25 10:27:48 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.telnetHosts : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"104.233.171.69:9999\\\":{\\\"delay\\\":0.11228203773498535},\\\"www.google.com:443\\\":{\\\"delay\\\":0.039106011390686035},\\\"104.233.171.69:8888\\\":{\\\"error\\\":\\\"connect error\\\"},\\\"www.csdn.com:443\\\":{\\\"delay\\\":0.13016796112060547},\\\"node2.wuwuwu360.xyz:443\\\":{\\\"delay\\\":0.47336196899414062}}}\\n2023-07-25 10:27:55 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {\\\"animated\\\":true}\\n2023-07-25 10:27:55 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"data\\\":{},\\\"message\\\":\\\"OK\\\",\\\"code\\\":0}\\n2023-07-25 10:27:55 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:27:58 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/1F79C151-10B7-420B-BBC0-E44800A68669/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 10:27:58 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:27:58 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:27:58 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0},\\\"appVersion\\\":\\\"1.0\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"deviceModal\\\":\\\"iPhone13,3\\\"}}\\n2023-07-25 10:27:58 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:28:19 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {}\\n2023-07-25 10:28:19 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 10:28:19 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:28:20 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/1F79C151-10B7-420B-BBC0-E44800A68669/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 10:28:20 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:28:20 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:28:20 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{\\\"deviceModal\\\":\\\"iPhone13,3\\\",\\\"appVersion\\\":\\\"1.0\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0}}}\\n2023-07-25 10:28:21 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:28:46 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/ae356b9188dd1ac3ac6aaf5171662d43.zip\\n2023-07-25 10:28:49 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {}\\n2023-07-25 10:28:49 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 10:28:50 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:28:50 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/1F79C151-10B7-420B-BBC0-E44800A68669/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 10:28:50 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:28:51 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 10:28:51 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"expired_time\\\":0,\\\"uname\\\":\\\"zhaoolin@126.com\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"已到期\\\",\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI4MCwiZXhwIjoxNjkyNDE3OTExfQ.P4NrIejnrZNCMVj1RpEnqexTNv8IVCfaI3Vs_o-11ec\\\",\\\"id\\\":219122280,\\\"surplus_flow\\\":0},\\\"deviceId\\\":\\\"66A97D92-1A9F-43B1-9292-24DC30EDB97A\\\",\\\"appVersion\\\":\\\"1.0\\\",\\\"osVersion\\\":\\\"16.3.1\\\",\\\"deviceModal\\\":\\\"iPhone13,3\\\"},\\\"code\\\":0}\\n2023-07-25 10:28:51 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:28:55 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.push : {\\\"hideNavBar\\\":false,\\\"url\\\":\\\"https:\\\\/\\\\/im.yyy360.xyz\\\\/chatIndex?kefu_id=kefu001\\\",\\\"animated\\\":true}\\n2023-07-25 10:28:55 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.push : {\\\"data\\\":{},\\\"message\\\":\\\"OK\\\",\\\"code\\\":0}\\n2023-07-25 10:28:55 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): https://im.yyy360.xyz/chatIndex?kefu_id=kefu001\\n2023-07-25 10:28:55 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:28:57 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:28:57 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:28:58 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.push : {\\\"hideNavBar\\\":false,\\\"url\\\":\\\"https:\\\\/\\\\/im.yyy360.xyz\\\\/chatIndex?kefu_id=kefu001\\\",\\\"animated\\\":true}\\n2023-07-25 10:28:58 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.push : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 10:28:58 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): https://im.yyy360.xyz/chatIndex?kefu_id=kefu001\\n2023-07-25 10:28:59 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:29:02 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 10:29:03 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 10:29:08 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:29:08 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 10:29:15 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppLog : {}\\n\"}', '2023-07-25 02:29:16', '2023-07-25 02:29:16', '');
INSERT INTO `t_upload_log` VALUES (4, 219122283, 1683674043779125248, '{\"log\":\"no log\"}', '2023-07-25 03:03:20', '2023-07-25 03:03:20', '');
INSERT INTO `t_upload_log` VALUES (5, 219122279, 1683679854756630528, '{\"log\":\"no log\"}', '2023-07-25 03:40:07', '2023-07-25 03:40:07', '');
INSERT INTO `t_upload_log` VALUES (6, 219122277, 1683729239741829120, '{\"log\":\"no log\"}', '2023-07-25 07:14:50', '2023-07-25 07:14:50', '');
INSERT INTO `t_upload_log` VALUES (7, 219122282, 1683737355854286848, '{\"log\":\"no log\"}', '2023-07-25 09:30:28', '2023-07-25 09:30:28', '');
INSERT INTO `t_upload_log` VALUES (8, 219122277, 1683839055315341312, '{\"log\":\"2023-07-25 21:58:06 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 21:58:07 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/2A20B0FE-9E38-4AA7-9125-DA93823EEED1/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 21:58:07 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 21:58:08 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 21:58:40 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.saveLoginInfo : {\\\"member_type\\\":0,\\\"surplus_flow\\\":0,\\\"id\\\":219122277,\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"expired_time\\\":1690271555,\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\"}\\n2023-07-25 21:58:40 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.saveLoginInfo : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 21:58:40 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 21:58:40 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"code\\\":0,\\\"data\\\":{\\\"appVersion\\\":\\\"1.0.2\\\",\\\"loginInfo\\\":{\\\"member_type\\\":0,\\\"surplus_flow\\\":0,\\\"id\\\":219122277,\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"expired_time\\\":1690271555,\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\"},\\\"deviceId\\\":\\\"8670F605-4B1E-429D-8EEE-F4E645C6EFDF\\\",\\\"deviceModal\\\":\\\"iPhone11,8\\\",\\\"osVersion\\\":\\\"14.2\\\"},\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 22:07:50 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:07:52 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/2A20B0FE-9E38-4AA7-9125-DA93823EEED1/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 22:07:52 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 22:07:52 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 22:07:54 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {}\\n2023-07-25 22:07:54 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"code\\\":0,\\\"data\\\":{},\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 22:07:55 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 22:07:56 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/2A20B0FE-9E38-4AA7-9125-DA93823EEED1/Documents/web/vpn/index.html#/pages/menu/index?type=1\\n2023-07-25 22:07:56 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 22:07:56 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 22:07:58 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.pop : {}\\n2023-07-25 22:07:58 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.pop : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{},\\\"code\\\":0}\\n2023-07-25 22:07:59 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 22:07:59 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): file:///var/mobile/Containers/Data/Application/2A20B0FE-9E38-4AA7-9125-DA93823EEED1/Documents/web/vpn/index.html\\n2023-07-25 22:07:59 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 22:07:59 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 22:08:04 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.saveLoginInfo : {\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"surplus_flow\\\":0,\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"id\\\":219122277,\\\"expired_time\\\":1690271555,\\\"member_type\\\":0}\\n2023-07-25 22:08:04 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.saveLoginInfo : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 22:08:04 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.setTunnelConfiguration : {\\\"nodeInfo\\\":{\\\"title\\\":\\\"自动\\\"},\\\"conf\\\":\\\"[General]\\\\nloglevel = {{logLevel}}\\\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\\\nlogoutput = {{leafLogFile}}\\\\ntun-fd = {{tunFd}}\\\\nrouting-domain-resolve = true\\\\nalways-real-ip = tracker, apple.com\\\\n\\\\n[Proxy]\\\\nDirect = direct\\\\nReject = reject\\\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=\\\\/work, tls-insecure=true\\\\n\\\\n[Rule]\\\\nEXTERNAL, site:{{dlcFile}}:cn, Direct\\\\nEXTERNAL, mmdb:{{geoFile}}:cn, Direct\\\\nFINAL, VMessWSS\\\\n\\\"}\\n2023-07-25 22:08:04 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.setTunnelConfiguration : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{},\\\"code\\\":0}\\n2023-07-25 22:08:04 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getStatus : {}\\n2023-07-25 22:08:04 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getStatus : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"status\\\":0}}\\n2023-07-25 22:08:07 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.setTunnelConfiguration : {\\\"conf\\\":\\\"[General]\\\\nloglevel = {{logLevel}}\\\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\\\nlogoutput = {{leafLogFile}}\\\\ntun-fd = {{tunFd}}\\\\nrouting-domain-resolve = true\\\\nalways-real-ip = tracker, apple.com\\\\n\\\\n[Proxy]\\\\nDirect = direct\\\\nReject = reject\\\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=\\\\/work, tls-insecure=true\\\\n\\\\n[Rule]\\\\nEXTERNAL, site:{{dlcFile}}:cn, Direct\\\\nEXTERNAL, mmdb:{{geoFile}}:cn, Direct\\\\nFINAL, VMessWSS\\\\n\\\",\\\"nodeInfo\\\":{\\\"title\\\":\\\"自动\\\"}}\\n2023-07-25 22:08:07 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.setTunnelConfiguration : {\\\"data\\\":{},\\\"code\\\":0,\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 22:08:07 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.open : {}\\n2023-07-25 22:08:09 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.open : {\\\"code\\\":5,\\\"message\\\":\\\"permission denied\\\"}\\n2023-07-25 22:08:27 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.setTunnelConfiguration : {\\\"conf\\\":\\\"[General]\\\\nloglevel = {{logLevel}}\\\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\\\nlogoutput = {{leafLogFile}}\\\\ntun-fd = {{tunFd}}\\\\nrouting-domain-resolve = true\\\\nalways-real-ip = tracker, apple.com\\\\n\\\\n[Proxy]\\\\nDirect = direct\\\\nReject = reject\\\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=\\\\/work, tls-insecure=true\\\\n\\\\n[Rule]\\\\nEXTERNAL, site:{{dlcFile}}:cn, Direct\\\\nEXTERNAL, mmdb:{{geoFile}}:cn, Direct\\\\nFINAL, VMessWSS\\\\n\\\",\\\"nodeInfo\\\":{\\\"title\\\":\\\"自动\\\"}}\\n2023-07-25 22:08:27 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.setTunnelConfiguration : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{}}\\n2023-07-25 22:08:27 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.open : {}\\n2023-07-25 22:08:28 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.open : {\\\"message\\\":\\\"permission denied\\\",\\\"code\\\":5}\\n2023-07-25 22:08:29 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.setTunnelConfiguration : {\\\"conf\\\":\\\"[General]\\\\nloglevel = {{logLevel}}\\\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\\\nlogoutput = {{leafLogFile}}\\\\ntun-fd = {{tunFd}}\\\\nrouting-domain-resolve = true\\\\nalways-real-ip = tracker, apple.com\\\\n\\\\n[Proxy]\\\\nDirect = direct\\\\nReject = reject\\\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=\\\\/work, tls-insecure=true\\\\n\\\\n[Rule]\\\\nEXTERNAL, site:{{dlcFile}}:cn, Direct\\\\nEXTERNAL, mmdb:{{geoFile}}:cn, Direct\\\\nFINAL, VMessWSS\\\\n\\\",\\\"nodeInfo\\\":{\\\"title\\\":\\\"自动\\\"}}\\n2023-07-25 22:08:29 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.setTunnelConfiguration : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{}}\\n2023-07-25 22:08:29 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.open : {}\\n2023-07-25 22:08:34 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 22:08:34 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"code\\\":0,\\\"data\\\":{\\\"status\\\":1}}\\n2023-07-25 22:08:34 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connecting\\n2023-07-25 22:08:34 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"code\\\":0,\\\"data\\\":{\\\"status\\\":2}}\\n2023-07-25 22:08:35 PacketTunnel/PacketTunnelProvider.swift:40:19 startTunnel(options:completionHandler:): start on main thread: false\\n2023-07-25 22:08:35 PacketTunnel/LeafAdapter.swift:161:19 start(completionHandler:): start on state stopped\\n2023-07-25 22:08:35 PacketTunnel/LeafAdapter.swift:224:19 didReceivePathUpdate(path:): net state changged on state: stopped, net status: satisfied\\n2023-07-25 22:08:35 PacketTunnel/LeafAdapter.swift:276:23 processConfig(): \\n-------\\n[General]\\nloglevel = error\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\nlogoutput = /private/var/mobile/Containers/Shared/AppGroup/6746B22E-A4E1-402C-A6CE-6750E62AC753/logs/leaf.log\\ntun-fd = 4\\nrouting-domain-resolve = true\\nalways-real-ip = tracker, apple.com\\n\\n[Proxy]\\nDirect = direct\\nReject = reject\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=/work, tls-insecure=true\\n\\n[Rule]\\nEXTERNAL, site:/private/var/containers/Bundle/Application/C1AFBAEB-1040-4321-B524-7A6A25956784/JunJunAPP.app/PlugIns/PacketTunnel.appex/site.dat:cn, Direct\\nEXTERNAL, mmdb:/private/var/containers/Bundle/Application/C1AFBAEB-1040-4321-B524-7A6A25956784/JunJunAPP.app/PlugIns/PacketTunnel.appex/geo.mmdb:cn, Direct\\nFINAL, VMessWSS\\n\\n-------\\n2023-07-25 22:08:35 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to started\\n2023-07-25 22:08:35 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-25 22:08:35 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.open : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{},\\\"code\\\":0}\\n2023-07-25 22:08:35 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":3},\\\"code\\\":0}\\n2023-07-25 22:08:35 PacketTunnel/LeafAdapter.swift:224:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-25 22:08:36 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:08:39 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.close : {}\\n2023-07-25 22:08:39 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnecting\\n2023-07-25 22:08:39 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"code\\\":0,\\\"data\\\":{\\\"status\\\":5}}\\n2023-07-25 22:08:39 PacketTunnel/PacketTunnelProvider.swift:60:19 stopTunnel(with:completionHandler:): stop\\n2023-07-25 22:08:39 PacketTunnel/LeafAdapter.swift:192:19 stop(completionHandler:): stop on state started\\n2023-07-25 22:08:39 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to stopped\\n2023-07-25 22:08:39 PacketTunnel/LeafAdapter.swift:181:27 start(completionHandler:): leaf shutdown on state stopped\\n2023-07-25 22:08:39 PacketTunnel/PacketTunnelProvider.swift:14:19 deinit: dealloc\\n2023-07-25 22:08:39 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 22:08:39 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.close : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{}}\\n2023-07-25 22:08:39 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"code\\\":0,\\\"data\\\":{\\\"status\\\":1}}\\n2023-07-25 22:08:40 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.setTunnelConfiguration : {\\\"conf\\\":\\\"[General]\\\\nloglevel = {{logLevel}}\\\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\\\nlogoutput = {{leafLogFile}}\\\\ntun-fd = {{tunFd}}\\\\nrouting-domain-resolve = true\\\\nalways-real-ip = tracker, apple.com\\\\n\\\\n[Proxy]\\\\nDirect = direct\\\\nReject = reject\\\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=\\\\/work, tls-insecure=true\\\\n\\\\n[Rule]\\\\nEXTERNAL, site:{{dlcFile}}:cn, Direct\\\\nEXTERNAL, mmdb:{{geoFile}}:cn, Direct\\\\nFINAL, VMessWSS\\\\n\\\",\\\"nodeInfo\\\":{\\\"title\\\":\\\"自动\\\"}}\\n2023-07-25 22:08:40 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.setTunnelConfiguration : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{},\\\"code\\\":0}\\n2023-07-25 22:08:40 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.open : {}\\n2023-07-25 22:08:40 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connecting\\n2023-07-25 22:08:40 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":2},\\\"code\\\":0}\\n2023-07-25 22:08:40 PacketTunnel/PacketTunnelProvider.swift:40:19 startTunnel(options:completionHandler:): start on main thread: false\\n2023-07-25 22:08:40 PacketTunnel/LeafAdapter.swift:161:19 start(completionHandler:): start on state stopped\\n2023-07-25 22:08:40 PacketTunnel/LeafAdapter.swift:224:19 didReceivePathUpdate(path:): net state changged on state: stopped, net status: satisfied\\n2023-07-25 22:08:40 PacketTunnel/LeafAdapter.swift:276:23 processConfig(): \\n-------\\n[General]\\nloglevel = error\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\nlogoutput = /private/var/mobile/Containers/Shared/AppGroup/6746B22E-A4E1-402C-A6CE-6750E62AC753/logs/leaf.log\\ntun-fd = 4\\nrouting-domain-resolve = true\\nalways-real-ip = tracker, apple.com\\n\\n[Proxy]\\nDirect = direct\\nReject = reject\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=/work, tls-insecure=true\\n\\n[Rule]\\nEXTERNAL, site:/private/var/containers/Bundle/Application/C1AFBAEB-1040-4321-B524-7A6A25956784/JunJunAPP.app/PlugIns/PacketTunnel.appex/site.dat:cn, Direct\\nEXTERNAL, mmdb:/private/var/containers/Bundle/Application/C1AFBAEB-1040-4321-B524-7A6A25956784/JunJunAPP.app/PlugIns/PacketTunnel.appex/geo.mmdb:cn, Direct\\nFINAL, VMessWSS\\n\\n-------\\n2023-07-25 22:08:40 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to started\\n2023-07-25 22:08:41 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-25 22:08:41 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.open : {\\\"code\\\":0,\\\"data\\\":{},\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 22:08:41 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":3},\\\"code\\\":0}\\n2023-07-25 22:08:41 PacketTunnel/LeafAdapter.swift:224:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-25 22:08:41 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:08:53 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.close : {}\\n2023-07-25 22:08:53 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnecting\\n2023-07-25 22:08:53 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":5},\\\"code\\\":0}\\n2023-07-25 22:08:53 PacketTunnel/PacketTunnelProvider.swift:60:19 stopTunnel(with:completionHandler:): stop\\n2023-07-25 22:08:53 PacketTunnel/LeafAdapter.swift:192:19 stop(completionHandler:): stop on state started\\n2023-07-25 22:08:53 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to stopped\\n2023-07-25 22:08:53 PacketTunnel/LeafAdapter.swift:181:27 start(completionHandler:): leaf shutdown on state stopped\\n2023-07-25 22:08:53 PacketTunnel/PacketTunnelProvider.swift:14:19 deinit: dealloc\\n2023-07-25 22:08:53 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 22:08:53 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.close : {\\\"message\\\":\\\"OK\\\",\\\"code\\\":0,\\\"data\\\":{}}\\n2023-07-25 22:08:53 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":1},\\\"code\\\":0}\\n2023-07-25 22:08:54 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:09:12 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.setTunnelConfiguration : {\\\"conf\\\":\\\"[General]\\\\nloglevel = {{logLevel}}\\\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\\\nlogoutput = {{leafLogFile}}\\\\ntun-fd = {{tunFd}}\\\\nrouting-domain-resolve = true\\\\nalways-real-ip = tracker, apple.com\\\\n\\\\n[Proxy]\\\\nDirect = direct\\\\nReject = reject\\\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=\\\\/work, tls-insecure=true\\\\n\\\\n[Rule]\\\\nEXTERNAL, site:{{dlcFile}}:cn, Direct\\\\nEXTERNAL, mmdb:{{geoFile}}:cn, Direct\\\\nFINAL, VMessWSS\\\\n\\\",\\\"nodeInfo\\\":{\\\"title\\\":\\\"自动\\\"}}\\n2023-07-25 22:09:12 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.setTunnelConfiguration : {\\\"code\\\":0,\\\"data\\\":{},\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 22:09:12 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.open : {}\\n2023-07-25 22:09:12 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connecting\\n2023-07-25 22:09:12 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"code\\\":0,\\\"data\\\":{\\\"status\\\":2}}\\n2023-07-25 22:09:12 PacketTunnel/PacketTunnelProvider.swift:40:19 startTunnel(options:completionHandler:): start on main thread: false\\n2023-07-25 22:09:12 PacketTunnel/LeafAdapter.swift:161:19 start(completionHandler:): start on state stopped\\n2023-07-25 22:09:12 PacketTunnel/LeafAdapter.swift:224:19 didReceivePathUpdate(path:): net state changged on state: stopped, net status: satisfied\\n2023-07-25 22:09:12 PacketTunnel/LeafAdapter.swift:276:23 processConfig(): \\n-------\\n[General]\\nloglevel = error\\ndns-server = 223.5.5.5, 114.114.114.114, 8.8.8.8\\nlogoutput = /private/var/mobile/Containers/Shared/AppGroup/6746B22E-A4E1-402C-A6CE-6750E62AC753/logs/leaf.log\\ntun-fd = 4\\nrouting-domain-resolve = true\\nalways-real-ip = tracker, apple.com\\n\\n[Proxy]\\nDirect = direct\\nReject = reject\\nVMessWSS = vmess, node2.wuwuwu360.xyz, 443, username=c541b521-17dd-11ee-bc4e-0c9d92c013fb, ws=true, tls=true, ws-path=/work, tls-insecure=true\\n\\n[Rule]\\nEXTERNAL, site:/private/var/containers/Bundle/Application/C1AFBAEB-1040-4321-B524-7A6A25956784/JunJunAPP.app/PlugIns/PacketTunnel.appex/site.dat:cn, Direct\\nEXTERNAL, mmdb:/private/var/containers/Bundle/Application/C1AFBAEB-1040-4321-B524-7A6A25956784/JunJunAPP.app/PlugIns/PacketTunnel.appex/geo.mmdb:cn, Direct\\nFINAL, VMessWSS\\n\\n-------\\n2023-07-25 22:09:12 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to started\\n2023-07-25 22:09:12 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Connected\\n2023-07-25 22:09:12 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.open : {\\\"data\\\":{},\\\"message\\\":\\\"OK\\\",\\\"code\\\":0}\\n2023-07-25 22:09:12 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":3},\\\"code\\\":0}\\n2023-07-25 22:09:12 PacketTunnel/LeafAdapter.swift:224:19 didReceivePathUpdate(path:): net state changged on state: started, net status: satisfied\\n2023-07-25 22:09:13 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:09:22 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:09:22 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.close : {}\\n2023-07-25 22:09:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnecting\\n2023-07-25 22:09:22 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":5},\\\"code\\\":0}\\n2023-07-25 22:09:22 PacketTunnel/PacketTunnelProvider.swift:60:19 stopTunnel(with:completionHandler:): stop\\n2023-07-25 22:09:22 PacketTunnel/LeafAdapter.swift:192:19 stop(completionHandler:): stop on state started\\n2023-07-25 22:09:22 PacketTunnel/LeafAdapter.swift:63:23 state: state changed to stopped\\n2023-07-25 22:09:22 PacketTunnel/LeafAdapter.swift:181:27 start(completionHandler:): leaf shutdown on state stopped\\n2023-07-25 22:09:22 PacketTunnel/PacketTunnelProvider.swift:14:19 deinit: dealloc\\n2023-07-25 22:09:22 JunJunAPP/VPNManager.swift:349:19 vpnStatusChannged(_:): Disconnected\\n2023-07-25 22:09:22 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.close : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{},\\\"code\\\":0}\\n2023-07-25 22:09:22 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.vpnStatusChanned to js : {\\\"data\\\":{\\\"status\\\":1},\\\"code\\\":0}\\n2023-07-25 22:09:24 JunJunAPP/AirportTool.swift:164:31 request(url:method:headers:completion:): request http://www.yyy360.xyz/app-api/app_info failed http://www.yyy360.xyz/app-upload/public/upload/other/a4f14b2c771ea34d442700320862a5d3.zip\\n2023-07-25 22:09:24 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 22:09:24 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"deviceId\\\":\\\"8670F605-4B1E-429D-8EEE-F4E645C6EFDF\\\",\\\"osVersion\\\":\\\"14.2\\\",\\\"appVersion\\\":\\\"1.0.2\\\",\\\"loginInfo\\\":{\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"surplus_flow\\\":0,\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"id\\\":219122277,\\\"expired_time\\\":1690271555,\\\"member_type\\\":0},\\\"deviceModal\\\":\\\"iPhone11,8\\\"},\\\"code\\\":0}\\n2023-07-25 22:09:36 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 22:09:36 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"code\\\":0,\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"deviceModal\\\":\\\"iPhone11,8\\\",\\\"osVersion\\\":\\\"14.2\\\",\\\"appVersion\\\":\\\"1.0.2\\\",\\\"loginInfo\\\":{\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"surplus_flow\\\":0,\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"id\\\":219122277,\\\"expired_time\\\":1690271555,\\\"member_type\\\":0},\\\"deviceId\\\":\\\"8670F605-4B1E-429D-8EEE-F4E645C6EFDF\\\"}}\\n2023-07-25 22:09:43 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 22:09:43 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{\\\"deviceId\\\":\\\"8670F605-4B1E-429D-8EEE-F4E645C6EFDF\\\",\\\"deviceModal\\\":\\\"iPhone11,8\\\",\\\"osVersion\\\":\\\"14.2\\\",\\\"appVersion\\\":\\\"1.0.2\\\",\\\"loginInfo\\\":{\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"surplus_flow\\\":0,\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"id\\\":219122277,\\\"expired_time\\\":1690271555,\\\"member_type\\\":0}},\\\"code\\\":0}\\n2023-07-25 22:09:45 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.push : {\\\"hideNavBar\\\":false,\\\"url\\\":\\\"https:\\\\/\\\\/im.yyy360.xyz\\\\/chatIndex?kefu_id=kefu001\\\",\\\"animated\\\":true}\\n2023-07-25 22:09:45 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.push : {\\\"message\\\":\\\"OK\\\",\\\"data\\\":{},\\\"code\\\":0}\\n2023-07-25 22:09:45 JunJunAPP/WebViewController.swift:200:19 webView(_:decidePolicyFor:decisionHandler:): https://im.yyy360.xyz/chatIndex?kefu_id=kefu001\\n2023-07-25 22:09:46 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 22:09:50 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppInfo : {}\\n2023-07-25 22:09:50 JunJunAPP/WebViewController.swift:121:31 viewDidLoad(): response to junjun.getAppInfo : {\\\"data\\\":{\\\"deviceId\\\":\\\"8670F605-4B1E-429D-8EEE-F4E645C6EFDF\\\",\\\"deviceModal\\\":\\\"iPhone11,8\\\",\\\"loginInfo\\\":{\\\"token\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIxOTEyMjI3NywiZXhwIjoxNjkyODg1NTE5fQ.IPBVDiMjZuv64hgGxblqoLP7yP9uqKLmCV_nMru4jyM\\\",\\\"uname\\\":\\\"wanglei@sina.com\\\",\\\"surplus_flow\\\":0,\\\"lastDate\\\":\\\"1970-01-20 09:31:11\\\",\\\"uuid\\\":\\\"c541b521-17dd-11ee-bc4e-0c9d92c013fb\\\",\\\"id\\\":219122277,\\\"expired_time\\\":1690271555,\\\"member_type\\\":0},\\\"appVersion\\\":\\\"1.0.2\\\",\\\"osVersion\\\":\\\"14.2\\\"},\\\"code\\\":0,\\\"message\\\":\\\"OK\\\"}\\n2023-07-25 22:09:50 JunJunAPP/WebUserContentController.swift:82:23 injectCookie(param:): COOKIE注入完成 nil\\n2023-07-25 22:09:50 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onHide to js : {\\\"data\\\":{\\\"animated\\\":true},\\\"code\\\":0}\\n2023-07-25 22:09:50 JunJunAPP/WebViewController.swift:193:19 fireEventToJS(eventName:eventData:): fire junjun.onShow to js : {\\\"code\\\":0,\\\"data\\\":{\\\"animated\\\":true}}\\n2023-07-25 22:09:53 JunJunAPP/V2RayBridge.swift:20:19 callApp(from:type:info:result:): revice JS message junjun.getAppLog : {}\\n\"}', '2023-07-25 14:09:53', '2023-07-25 14:09:53', '');
INSERT INTO `t_upload_log` VALUES (9, 219122284, 1684026378766258176, '{\"log\":\"\"}', '2023-07-27 01:33:18', '2023-07-27 01:33:18', '');
INSERT INTO `t_upload_log` VALUES (10, 219122284, 1684026378766258176, '{\"log\":\"\"}', '2023-07-27 02:15:30', '2023-07-27 02:15:30', '');

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
  `v2ray_tag` int NULL DEFAULT NULL COMMENT 'v2ray存在UUID标签:1-有；2-无',
  `channel_id` int NULL DEFAULT NULL COMMENT '渠道id',
  `status` int NULL DEFAULT NULL COMMENT '冻结状态：0-正常；1-冻结',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `p_uname_index`(`uname` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 219122308 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (219122276, 'win2023bar@outlook.com', 'c33367701511b4f6020ec61ded352059', 'win2023bar@outlook.com', '', 0, 1690271555, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-07 15:27:23', '2023-07-25 06:52:35', '');
INSERT INTO `t_user` VALUES (219122277, 'wanglei@sina.com', '5bacd9f25613659b2fbd2f3a58822e5c', 'wanglei@sina.com', '', 0, 1690366667, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-19 04:05:35', '2023-07-26 09:17:47', '');
INSERT INTO `t_user` VALUES (219122278, 'wanglei1@sina.com', '5bacd9f25613659b2fbd2f3a58822e5c', 'wanglei1@sina.com', '', 0, 1690271555, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-19 04:06:24', '2023-07-19 04:06:24', '');
INSERT INTO `t_user` VALUES (219122279, 'pujin@gmail.com', 'fcea920f7412b5da7be0cf42b8c93759', 'pujin@gmail.com', '', 0, 1690270527, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-19 07:42:32', '2023-07-25 06:35:27', '');
INSERT INTO `t_user` VALUES (219122280, 'zhaoolin@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'zhaoolin@126.com', '', 0, 1690271555, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-20 03:23:45', '2023-07-20 03:23:45', '');
INSERT INTO `t_user` VALUES (219122281, 'zhaoolin@sina.com', 'e10adc3949ba59abbe56e057f20f883e', 'zhaoolin@sina.com', '', 0, 1690179295, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-24 06:14:55', '2023-07-24 06:14:55', '');
INSERT INTO `t_user` VALUES (219122282, 'aaa@qq.com', 'e10adc3949ba59abbe56e057f20f883e', 'aaa@qq.com', '', 2, 1690550200, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-24 09:33:34', '2023-07-28 12:16:40', '');
INSERT INTO `t_user` VALUES (219122283, 'hshm20517@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'hshm20517@126.com', '', 0, 1690248474, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-25 01:27:54', '2023-07-25 01:27:54', '');
INSERT INTO `t_user` VALUES (219122284, 'habi@gmail.com', 'e10adc3949ba59abbe56e057f20f883e', 'habi@gmail.com', '', 0, 1690367102, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-25 03:57:44', '2023-07-26 09:25:02', '');
INSERT INTO `t_user` VALUES (219122285, 'wanglei0316@sina.com', '5bacd9f25613659b2fbd2f3a58822e5c', 'wanglei0316@sina.com', '', 0, 1690269375, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-25 07:16:15', '2023-07-25 07:16:15', '');
INSERT INTO `t_user` VALUES (219122286, 'bbb@qq.com', 'e10adc3949ba59abbe56e057f20f883e', 'bbb@qq.com', '', 0, 1690283919, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-25 11:18:39', '2023-07-25 11:18:39', '');
INSERT INTO `t_user` VALUES (219122287, 'zhaolin@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'zhaolin@126.com', '', 0, 1690290758, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-25 13:12:38', '2023-07-25 13:12:38', '');
INSERT INTO `t_user` VALUES (219122288, 'Hshm20517@163.com', 'e10adc3949ba59abbe56e057f20f883e', 'Hshm20517@163.com', '', 0, 1690339742, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-26 02:49:02', '2023-07-26 02:49:02', '');
INSERT INTO `t_user` VALUES (219122289, 'pujin@qq.com', 'e10adc3949ba59abbe56e057f20f883e', 'pujin@qq.com', '', 0, 1690342957, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-26 03:42:37', '2023-07-26 03:42:37', '');
INSERT INTO `t_user` VALUES (219122290, 'hh@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'hh@126.com', '', 0, 1690353871, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-26 06:44:31', '2023-07-26 06:44:31', '');
INSERT INTO `t_user` VALUES (219122291, 'ccc@qq.com', 'e10adc3949ba59abbe56e057f20f883e', 'ccc@qq.com', '', 0, 1690367127, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-26 09:25:27', '2023-07-26 09:25:27', '');
INSERT INTO `t_user` VALUES (219122292, 'wanglei0317@sina.com', '5bacd9f25613659b2fbd2f3a58822e5c', 'wanglei0317@sina.com', '', 0, 1690367660, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-26 09:34:20', '2023-07-26 09:34:20', '');
INSERT INTO `t_user` VALUES (219122293, 'aa@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'aa@126.com', '', 0, 1690428456, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-27 02:27:36', '2023-07-27 02:27:36', '');
INSERT INTO `t_user` VALUES (219122295, 'habi@qq.com', 'e10adc3949ba59abbe56e057f20f883e', 'habi@qq.com', '', 0, 1690453818, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-27 09:30:18', '2023-07-27 09:30:18', '');
INSERT INTO `t_user` VALUES (219122296, '2307823881@qq.com', '8e310df5da170452d5fe1bca11f1e5cc', '2307823881@qq.com', '', 0, 1690454603, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-27 09:43:23', '2023-07-27 09:43:23', '');
INSERT INTO `t_user` VALUES (219122297, '303468504@huawei.com', 'ca18025fa350e81f615bafaad25dadbb', '303468504@huawei.com', '', 0, 1690455012, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-27 09:50:12', '2023-07-27 09:50:12', '');
INSERT INTO `t_user` VALUES (219122298, 'wanlgei3@sina.com', '5bacd9f25613659b2fbd2f3a58822e5c', 'wanlgei3@sina.com', '', 0, 1690456276, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-27 10:11:16', '2023-07-27 10:11:16', '');
INSERT INTO `t_user` VALUES (219122299, 'zhaoolin@sina.com1', 'e10adc3949ba59abbe56e057f20f883e', 'zhaoolin@sina.com1', '', 0, 1690460031, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-27 11:13:51', '2023-07-27 11:13:51', '');
INSERT INTO `t_user` VALUES (219122300, 'zhaoolin@sina.com11', 'e10adc3949ba59abbe56e057f20f883e', 'zhaoolin@sina.com11', '', 0, 1690511204, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-28 01:26:44', '2023-07-28 01:26:44', '');
INSERT INTO `t_user` VALUES (219122302, 'pujin@outlook.com', 'e10adc3949ba59abbe56e057f20f883e', 'pujin@outlook.com', '', 0, 1690529113, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-28 06:25:13', '2023-07-28 06:25:13', '');
INSERT INTO `t_user` VALUES (219122303, 'hs@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'hs@126.com', '', 0, 1690532137, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-28 07:15:37', '2023-07-28 07:15:37', '');
INSERT INTO `t_user` VALUES (219122304, 'hm@126.com', 'e10adc3949ba59abbe56e057f20f883e', 'hm@126.com', '', 0, 1690535717, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-28 08:15:17', '2023-07-28 08:15:17', '');
INSERT INTO `t_user` VALUES (219122305, 'shako@qq.com', 'e10adc3949ba59abbe56e057f20f883e', 'shako@qq.com', '', 0, 1690566431, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-28 16:47:11', '2023-07-28 16:47:11', '');
INSERT INTO `t_user` VALUES (219122306, '123456@qq.com', 'e10adc3949ba59abbe56e057f20f883e', '123456@qq.com', '', 0, 1690604596, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 0, 0, '2023-07-29 03:23:16', '2023-07-29 03:23:16', '');
INSERT INTO `t_user` VALUES (219122307, 'zhaoolin@sina.com2', 'e10adc3949ba59abbe56e057f20f883e', 'zhaoolin@sina.com2', '', 0, 1690638308, 'c541b521-17dd-11ee-bc4e-0c9d92c013fb', 0, 1, 0, '2023-07-29 13:45:08', '2023-07-29 13:45:08', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 78 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户设备表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user_dev
-- ----------------------------
INSERT INTO `t_user_dev` VALUES (1, 219122276, 1677211938733428736, 1, '2023-07-07 15:42:36', '2023-07-07 15:42:36', '');
INSERT INTO `t_user_dev` VALUES (2, 219122276, 1677211938733428737, 2, '2023-07-07 17:09:11', '2023-07-07 17:10:34', NULL);
INSERT INTO `t_user_dev` VALUES (3, 219122279, 1681570017864323072, 2, '2023-07-19 07:42:33', '2023-07-20 12:15:33', '');
INSERT INTO `t_user_dev` VALUES (4, 219122277, 1681571524739338240, 1, '2023-07-19 07:47:56', '2023-07-28 10:42:18', '');
INSERT INTO `t_user_dev` VALUES (5, 219122277, 1681630009988485120, 2, '2023-07-19 11:40:35', '2023-07-20 07:28:49', '');
INSERT INTO `t_user_dev` VALUES (6, 219122279, 1681855058557276160, 2, '2023-07-20 02:35:24', '2023-07-20 12:15:37', '');
INSERT INTO `t_user_dev` VALUES (7, 219122280, 1681867380034113536, 1, '2023-07-20 03:23:48', '2023-07-20 03:23:48', '');
INSERT INTO `t_user_dev` VALUES (8, 219122279, 1681867930939166720, 2, '2023-07-20 03:26:09', '2023-07-21 10:04:20', '');
INSERT INTO `t_user_dev` VALUES (9, 219122277, 1681869950098083840, 2, '2023-07-20 03:34:11', '2023-07-25 06:42:38', '');
INSERT INTO `t_user_dev` VALUES (10, 219122280, 1681877802594340864, 1, '2023-07-20 04:05:11', '2023-07-20 04:05:11', '');
INSERT INTO `t_user_dev` VALUES (11, 219122279, 1682326372267069440, 2, '2023-07-21 09:47:52', '2023-07-25 09:10:43', '');
INSERT INTO `t_user_dev` VALUES (12, 219122277, 1682574427012730880, 2, '2023-07-22 02:17:27', '2023-07-25 09:25:28', '');
INSERT INTO `t_user_dev` VALUES (13, 219122281, 1683359980536729600, 2, '2023-07-24 06:14:56', '2023-07-25 05:15:54', '');
INSERT INTO `t_user_dev` VALUES (14, 219122281, 1683361531191889920, 2, '2023-07-24 06:20:52', '2023-07-25 05:15:52', '');
INSERT INTO `t_user_dev` VALUES (15, 219122281, 1683399907219607552, 2, '2023-07-24 08:53:21', '2023-07-25 05:15:51', '');
INSERT INTO `t_user_dev` VALUES (16, 219122279, 1683406721491406848, 2, '2023-07-24 09:20:59', '2023-07-25 08:11:38', '');
INSERT INTO `t_user_dev` VALUES (17, 219122279, 1683408308884475904, 2, '2023-07-24 09:27:04', '2023-07-25 08:11:34', '');
INSERT INTO `t_user_dev` VALUES (18, 219122279, 1683409637799366656, 2, '2023-07-24 09:32:14', '2023-07-25 06:36:51', '');
INSERT INTO `t_user_dev` VALUES (19, 219122282, 1683409685522157568, 2, '2023-07-24 09:33:38', '2023-07-25 11:54:57', '');
INSERT INTO `t_user_dev` VALUES (20, 219122277, 1683411282608263168, 2, '2023-07-24 09:38:53', '2023-07-25 06:42:41', '');
INSERT INTO `t_user_dev` VALUES (21, 219122277, 1683422926776307712, 2, '2023-07-24 10:25:08', '2023-07-25 06:42:45', '');
INSERT INTO `t_user_dev` VALUES (22, 219122281, 1683426939466944512, 2, '2023-07-24 10:40:48', '2023-07-25 05:15:49', '');
INSERT INTO `t_user_dev` VALUES (23, 219122279, 1683434186548973568, 2, '2023-07-24 11:09:38', '2023-07-25 08:11:21', '');
INSERT INTO `t_user_dev` VALUES (24, 219122283, 1683405443369537536, 1, '2023-07-25 01:27:57', '2023-07-25 01:27:57', '');
INSERT INTO `t_user_dev` VALUES (25, 219122283, 1683650502971101184, 1, '2023-07-25 01:29:08', '2023-07-25 01:29:08', '');
INSERT INTO `t_user_dev` VALUES (26, 219122288, 1683660611679948800, 1, '2023-07-25 02:09:16', '2023-07-26 02:49:04', '');
INSERT INTO `t_user_dev` VALUES (27, 219122283, 1683674043779125248, 1, '2023-07-25 03:03:01', '2023-07-25 03:03:01', '');
INSERT INTO `t_user_dev` VALUES (28, 219122283, 1683678852921954304, 1, '2023-07-25 03:21:58', '2023-07-25 03:21:58', '');
INSERT INTO `t_user_dev` VALUES (29, 219122279, 1683679854756630528, 2, '2023-07-25 03:25:44', '2023-07-26 07:31:49', '');
INSERT INTO `t_user_dev` VALUES (30, 219122279, 1683681627907035136, 2, '2023-07-25 03:32:47', '2023-07-25 06:55:18', '');
INSERT INTO `t_user_dev` VALUES (31, 219122279, 1683687695353647104, 2, '2023-07-25 03:57:47', '2023-07-25 10:17:41', '');
INSERT INTO `t_user_dev` VALUES (32, 219122281, 1683703712901500928, 1, '2023-07-25 05:01:05', '2023-07-25 05:01:05', '');
INSERT INTO `t_user_dev` VALUES (33, 219122279, 1683703053024235520, 2, '2023-07-25 06:33:18', '2023-07-25 09:10:49', '');
INSERT INTO `t_user_dev` VALUES (34, 219122283, 1683727763741085696, 1, '2023-07-25 06:36:06', '2023-07-25 06:36:06', '');
INSERT INTO `t_user_dev` VALUES (35, 219122277, 1683729239741829120, 2, '2023-07-25 06:48:36', '2023-07-25 09:27:09', '');
INSERT INTO `t_user_dev` VALUES (36, 219122283, 1683736170959212544, 1, '2023-07-25 07:09:30', '2023-07-25 07:09:30', '');
INSERT INTO `t_user_dev` VALUES (37, 219122283, 1683736718768869376, 1, '2023-07-25 07:11:49', '2023-07-25 07:11:49', '');
INSERT INTO `t_user_dev` VALUES (38, 219122282, 1683737355854286848, 1, '2023-07-25 07:14:41', '2023-07-28 11:55:06', '');
INSERT INTO `t_user_dev` VALUES (39, 219122283, 1683743613386756096, 1, '2023-07-25 07:39:11', '2023-07-25 07:39:11', '');
INSERT INTO `t_user_dev` VALUES (40, 219122283, 1683746051959296000, 1, '2023-07-25 07:49:16', '2023-07-25 07:49:16', '');
INSERT INTO `t_user_dev` VALUES (41, 219122283, 1683756993799524352, 1, '2023-07-25 08:32:34', '2023-07-25 08:32:34', '');
INSERT INTO `t_user_dev` VALUES (42, 219122279, 1683783152318812160, 1, '2023-07-25 10:24:19', '2023-07-26 02:15:24', '');
INSERT INTO `t_user_dev` VALUES (43, 219122284, 1683808014429065216, 2, '2023-07-25 11:57:02', '2023-07-27 01:29:28', '');
INSERT INTO `t_user_dev` VALUES (44, 219122287, 1683827203457945600, 1, '2023-07-25 13:12:40', '2023-07-25 13:12:40', '');
INSERT INTO `t_user_dev` VALUES (45, 219122277, 1683839055315341312, 1, '2023-07-25 13:58:39', '2023-07-27 13:08:49', '');
INSERT INTO `t_user_dev` VALUES (46, 219122279, 1684021330074144768, 1, '2023-07-26 02:02:43', '2023-07-27 01:43:20', '');
INSERT INTO `t_user_dev` VALUES (47, 219122284, 1684026378766258176, 1, '2023-07-26 02:23:24', '2023-07-26 02:23:24', '');
INSERT INTO `t_user_dev` VALUES (48, 219122289, 1683777087724326912, 1, '2023-07-26 03:42:39', '2023-07-26 03:42:39', '');
INSERT INTO `t_user_dev` VALUES (49, 219122290, 1684092087143763968, 1, '2023-07-26 06:44:32', '2023-07-26 06:44:32', '');
INSERT INTO `t_user_dev` VALUES (50, 219122290, 1684111615957405696, 1, '2023-07-26 08:01:29', '2023-07-26 08:01:29', '');
INSERT INTO `t_user_dev` VALUES (51, 219122284, 1684373826332266496, 2, '2023-07-27 01:33:36', '2023-07-27 08:18:04', '');
INSERT INTO `t_user_dev` VALUES (52, 219122290, 1684389925358669824, 1, '2023-07-27 02:27:37', '2023-07-28 06:57:56', '');
INSERT INTO `t_user_dev` VALUES (53, 219122295, 1684496326865195008, 1, '2023-07-27 09:30:20', '2023-07-27 12:13:42', '');
INSERT INTO `t_user_dev` VALUES (54, 219122296, 1684484081770827776, 1, '2023-07-27 09:43:25', '2023-07-27 09:43:25', '');
INSERT INTO `t_user_dev` VALUES (55, 219122297, 1684501142672773120, 1, '2023-07-27 09:50:16', '2023-07-27 09:50:16', '');
INSERT INTO `t_user_dev` VALUES (56, 219122295, 1684487002629607424, 1, '2023-07-27 09:56:21', '2023-07-27 09:56:21', '');
INSERT INTO `t_user_dev` VALUES (57, 219122281, 1684521730258767872, 1, '2023-07-27 11:11:08', '2023-07-27 11:11:08', '');
INSERT INTO `t_user_dev` VALUES (58, 219122299, 1684522279972638720, 2, '2023-07-27 11:13:52', '2023-07-28 01:39:11', '');
INSERT INTO `t_user_dev` VALUES (59, 219122299, 1684525565203189760, 1, '2023-07-27 11:26:29', '2023-07-28 07:21:06', '');
INSERT INTO `t_user_dev` VALUES (60, 219122300, 1684464079315406848, 1, '2023-07-28 01:26:45', '2023-07-28 01:26:45', '');
INSERT INTO `t_user_dev` VALUES (61, 219122284, 1684753436886700032, 1, '2023-07-28 02:52:32', '2023-07-28 02:52:32', '');
INSERT INTO `t_user_dev` VALUES (62, 219122278, 1684759095699050496, 1, '2023-07-28 03:00:33', '2023-07-29 10:05:54', '');
INSERT INTO `t_user_dev` VALUES (63, 219122289, 1684753295664484352, 1, '2023-07-28 03:40:53', '2023-07-28 03:40:53', '');
INSERT INTO `t_user_dev` VALUES (64, 219122299, 1684803646232989696, 1, '2023-07-28 06:00:58', '2023-07-28 12:41:51', '');
INSERT INTO `t_user_dev` VALUES (65, 219122302, 1684810635877027840, 1, '2023-07-28 06:25:15', '2023-07-28 08:32:06', '');
INSERT INTO `t_user_dev` VALUES (66, 219122285, 1684819516325892096, 1, '2023-07-28 06:54:39', '2023-07-28 06:54:39', '');
INSERT INTO `t_user_dev` VALUES (67, 219122300, 1684814428836466688, 1, '2023-07-28 07:00:20', '2023-07-28 07:00:20', '');
INSERT INTO `t_user_dev` VALUES (68, 219122302, 1684823808927600640, 1, '2023-07-28 07:12:05', '2023-07-28 07:12:05', '');
INSERT INTO `t_user_dev` VALUES (69, 219122303, 1684824801509642240, 1, '2023-07-28 07:15:39', '2023-07-28 07:34:12', '');
INSERT INTO `t_user_dev` VALUES (70, 219122303, 1684838129904652288, 1, '2023-07-28 08:08:27', '2023-07-28 08:08:27', '');
INSERT INTO `t_user_dev` VALUES (71, 219122304, 1684839812965601280, 1, '2023-07-28 08:15:19', '2023-07-28 08:15:19', '');
INSERT INTO `t_user_dev` VALUES (72, 219122289, 1684968511522213888, 1, '2023-07-28 16:47:13', '2023-07-29 04:15:25', '');
INSERT INTO `t_user_dev` VALUES (73, 219122277, 1685114917293658112, 1, '2023-07-29 02:28:27', '2023-07-30 12:06:47', '');
INSERT INTO `t_user_dev` VALUES (74, 219122306, 1684969498777489408, 1, '2023-07-29 03:23:22', '2023-07-29 03:23:22', '');
INSERT INTO `t_user_dev` VALUES (75, 219122307, 1685285207160131584, 2, '2023-07-29 13:45:10', '2023-07-29 13:56:20', '');
INSERT INTO `t_user_dev` VALUES (76, 219122307, 1685288035249295360, 1, '2023-07-29 13:56:11', '2023-07-29 13:56:11', '');
INSERT INTO `t_user_dev` VALUES (77, 219122307, 1685291172928425984, 1, '2023-07-29 14:08:36', '2023-07-29 14:08:36', '');

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
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '推荐关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_user_team
-- ----------------------------
INSERT INTO `t_user_team` VALUES (1, 219122276, 0, '', '2023-07-07 15:27:23', '2023-07-07 15:27:23', '');
INSERT INTO `t_user_team` VALUES (2, 219122277, 0, '', '2023-07-19 04:05:35', '2023-07-19 04:05:35', '');
INSERT INTO `t_user_team` VALUES (3, 219122278, 0, '', '2023-07-19 04:06:24', '2023-07-19 04:06:24', '');
INSERT INTO `t_user_team` VALUES (4, 219122279, 0, '', '2023-07-19 07:42:32', '2023-07-19 07:42:32', '');
INSERT INTO `t_user_team` VALUES (5, 219122280, 0, '', '2023-07-20 03:23:45', '2023-07-20 03:23:45', '');
INSERT INTO `t_user_team` VALUES (6, 219122281, 0, '', '2023-07-24 06:14:55', '2023-07-24 06:14:55', '');
INSERT INTO `t_user_team` VALUES (7, 219122282, 0, '', '2023-07-24 09:33:34', '2023-07-24 09:33:34', '');
INSERT INTO `t_user_team` VALUES (8, 219122283, 0, '', '2023-07-25 01:27:54', '2023-07-25 01:27:54', '');
INSERT INTO `t_user_team` VALUES (9, 219122284, 219122279, '219122279', '2023-07-25 03:57:44', '2023-07-25 03:57:44', '');
INSERT INTO `t_user_team` VALUES (10, 219122285, 219122277, '219122277', '2023-07-25 07:16:15', '2023-07-25 07:16:15', '');
INSERT INTO `t_user_team` VALUES (11, 219122286, 219122282, '219122282', '2023-07-25 11:18:39', '2023-07-25 11:18:39', '');
INSERT INTO `t_user_team` VALUES (12, 219122287, 0, '', '2023-07-25 13:12:38', '2023-07-25 13:12:38', '');
INSERT INTO `t_user_team` VALUES (13, 219122288, 0, '', '2023-07-26 02:49:02', '2023-07-26 02:49:02', '');
INSERT INTO `t_user_team` VALUES (14, 219122289, 0, '', '2023-07-26 03:42:37', '2023-07-26 03:42:37', '');
INSERT INTO `t_user_team` VALUES (15, 219122290, 0, '', '2023-07-26 06:44:31', '2023-07-26 06:44:31', '');
INSERT INTO `t_user_team` VALUES (16, 219122291, 219122276, '219122276', '2023-07-26 09:25:27', '2023-07-26 09:25:27', '');
INSERT INTO `t_user_team` VALUES (17, 219122292, 0, '', '2023-07-26 09:34:20', '2023-07-26 09:34:20', '');
INSERT INTO `t_user_team` VALUES (18, 219122293, 0, '', '2023-07-27 02:27:36', '2023-07-27 02:27:36', '');
INSERT INTO `t_user_team` VALUES (19, 219122295, 0, '', '2023-07-27 09:30:18', '2023-07-27 09:30:18', '');
INSERT INTO `t_user_team` VALUES (20, 219122296, 0, '', '2023-07-27 09:43:23', '2023-07-27 09:43:23', '');
INSERT INTO `t_user_team` VALUES (21, 219122297, 0, '', '2023-07-27 09:50:12', '2023-07-27 09:50:12', '');
INSERT INTO `t_user_team` VALUES (22, 219122298, 0, '', '2023-07-27 10:11:16', '2023-07-27 10:11:16', '');
INSERT INTO `t_user_team` VALUES (23, 219122299, 0, '', '2023-07-27 11:13:51', '2023-07-27 11:13:51', '');
INSERT INTO `t_user_team` VALUES (24, 219122300, 0, '', '2023-07-28 01:26:44', '2023-07-28 01:26:44', '');
INSERT INTO `t_user_team` VALUES (25, 219122302, 0, '', '2023-07-28 06:25:13', '2023-07-28 06:25:13', '');
INSERT INTO `t_user_team` VALUES (26, 219122303, 0, '', '2023-07-28 07:15:37', '2023-07-28 07:15:37', '');
INSERT INTO `t_user_team` VALUES (27, 219122304, 0, '', '2023-07-28 08:15:17', '2023-07-28 08:15:17', '');
INSERT INTO `t_user_team` VALUES (28, 219122305, 0, '', '2023-07-28 16:47:11', '2023-07-28 16:47:11', '');
INSERT INTO `t_user_team` VALUES (29, 219122306, 0, '', '2023-07-29 03:23:16', '2023-07-29 03:23:16', '');
INSERT INTO `t_user_team` VALUES (30, 219122307, 0, '', '2023-07-29 13:45:08', '2023-07-29 13:45:08', '');

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '工作模式' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 7185 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户日志表（仅记录第一次事件)' ROW_FORMAT = Dynamic;

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
INSERT INTO `user_logs` VALUES (2807, 100004, '2023-07-17', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-07-17 09:53:54', '2023-07-17 09:53:54', NULL);
INSERT INTO `user_logs` VALUES (3652, 100004, '2023-07-18', '10.10.10.222', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', '2023-07-18 09:28:54', '2023-07-18 09:28:54', NULL);
INSERT INTO `user_logs` VALUES (4014, 100004, '2023-07-19', '206.189.36.86', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-19 01:42:23', '2023-07-19 01:42:23', NULL);
INSERT INTO `user_logs` VALUES (4027, 219122307, '2023-07-19', '203.205.141.118', 'Mozilla/5.0 (iPhone; CPU iPhone OS 15_7_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-19 01:45:33', '2023-07-19 01:45:33', NULL);
INSERT INTO `user_logs` VALUES (4098, 219122296, '2023-07-19', '165.154.72.249', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-19 04:04:48', '2023-07-19 04:04:48', NULL);
INSERT INTO `user_logs` VALUES (4111, 219122303, '2023-07-19', '139.5.108.135', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-19 06:43:18', '2023-07-19 06:43:18', NULL);
INSERT INTO `user_logs` VALUES (4115, 219122279, '2023-07-19', '205.198.122.67', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-19 07:42:34', '2023-07-19 07:42:34', NULL);
INSERT INTO `user_logs` VALUES (4117, 219122277, '2023-07-19', '14.26.14.4', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-19 07:47:57', '2023-07-19 07:47:57', NULL);
INSERT INTO `user_logs` VALUES (4321, 219122277, '2023-07-20', '43.249.50.178', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-20 02:26:56', '2023-07-20 02:26:56', NULL);
INSERT INTO `user_logs` VALUES (4322, 219122279, '2023-07-20', '139.5.108.71', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-20 02:30:38', '2023-07-20 02:30:38', NULL);
INSERT INTO `user_logs` VALUES (4352, 219122280, '2023-07-20', '43.132.98.34', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-20 03:23:49', '2023-07-20 03:23:49', NULL);
INSERT INTO `user_logs` VALUES (4367, 100004, '2023-07-20', '188.166.233.207', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-20 03:41:48', '2023-07-20 03:41:48', NULL);
INSERT INTO `user_logs` VALUES (4450, 219122307, '2023-07-20', '203.205.141.118', 'Mozilla/5.0 (iPhone; CPU iPhone OS 15_7_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-20 07:34:10', '2023-07-20 07:34:10', NULL);
INSERT INTO `user_logs` VALUES (4646, 219122276, '2023-07-20', '188.166.233.207', 'PostmanRuntime/7.32.3', '2023-07-20 12:22:13', '2023-07-20 12:22:13', NULL);
INSERT INTO `user_logs` VALUES (4669, 219122277, '2023-07-21', '14.26.14.4', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36', '2023-07-21 02:11:37', '2023-07-21 02:11:37', NULL);
INSERT INTO `user_logs` VALUES (4693, 100004, '2023-07-21', '188.166.233.207', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-21 02:42:43', '2023-07-21 02:42:43', NULL);
INSERT INTO `user_logs` VALUES (4704, 219122279, '2023-07-21', '139.5.108.135', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-21 02:44:32', '2023-07-21 02:44:32', NULL);
INSERT INTO `user_logs` VALUES (4878, 219122277, '2023-07-22', '113.110.176.188', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-22 00:34:05', '2023-07-22 00:34:05', NULL);
INSERT INTO `user_logs` VALUES (4885, 100004, '2023-07-22', '188.166.233.207', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-22 01:22:42', '2023-07-22 01:22:42', NULL);
INSERT INTO `user_logs` VALUES (4966, 219122279, '2023-07-22', '49.0.245.9', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-22 02:58:18', '2023-07-22 02:58:18', NULL);
INSERT INTO `user_logs` VALUES (5055, 219122280, '2023-07-22', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-22 15:02:42', '2023-07-22 15:02:42', NULL);
INSERT INTO `user_logs` VALUES (5056, 219122279, '2023-07-23', '139.5.108.145', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-23 08:57:31', '2023-07-23 08:57:31', NULL);
INSERT INTO `user_logs` VALUES (5073, 219122280, '2023-07-23', '223.104.68.8', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-23 09:26:36', '2023-07-23 09:26:36', NULL);
INSERT INTO `user_logs` VALUES (5075, 219122277, '2023-07-23', '183.8.143.90', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-23 13:03:19', '2023-07-23 13:03:19', NULL);
INSERT INTO `user_logs` VALUES (5094, 219122277, '2023-07-24', '14.26.82.228', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-24 01:28:05', '2023-07-24 01:28:05', NULL);
INSERT INTO `user_logs` VALUES (5097, 219122279, '2023-07-24', '205.198.122.77', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-24 01:29:11', '2023-07-24 01:29:11', NULL);
INSERT INTO `user_logs` VALUES (5148, 219122281, '2023-07-24', '43.132.98.31', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-24 06:14:57', '2023-07-24 06:14:57', NULL);
INSERT INTO `user_logs` VALUES (5187, 100004, '2023-07-24', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-24 08:49:13', '2023-07-24 08:49:13', NULL);
INSERT INTO `user_logs` VALUES (5209, 219122276, '2023-07-24', '172.105.235.23', 'PostmanRuntime/7.32.3', '2023-07-24 09:19:57', '2023-07-24 09:19:57', NULL);
INSERT INTO `user_logs` VALUES (5235, 219122282, '2023-07-24', '107.148.239.239', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-24 09:33:38', '2023-07-24 09:33:38', NULL);
INSERT INTO `user_logs` VALUES (5275, 219122280, '2023-07-24', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-24 12:52:05', '2023-07-24 12:52:05', NULL);
INSERT INTO `user_logs` VALUES (5277, 219122283, '2023-07-25', '43.132.98.39', 'Mozilla/5.0 (Linux; Android 9; SM-N9500 Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/114.0.5735.196 Mobile Safari/537.36', '2023-07-25 01:27:57', '2023-07-25 01:27:57', NULL);
INSERT INTO `user_logs` VALUES (5282, 100004, '2023-07-25', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-25 01:30:48', '2023-07-25 01:30:48', NULL);
INSERT INTO `user_logs` VALUES (5305, 219122281, '2023-07-25', '43.132.98.31', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-25 02:09:16', '2023-07-25 02:09:16', NULL);
INSERT INTO `user_logs` VALUES (5311, 219122280, '2023-07-25', '103.7.29.100', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-25 02:26:49', '2023-07-25 02:26:49', NULL);
INSERT INTO `user_logs` VALUES (5313, 219122277, '2023-07-25', '203.91.85.81', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-25 02:49:03', '2023-07-25 02:49:03', NULL);
INSERT INTO `user_logs` VALUES (5326, 219122279, '2023-07-25', '183.15.206.81', 'Mozilla/5.0 (Linux; Android 12; V1981A Build/SP1A.210812.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Mobile Safari/537.36', '2023-07-25 03:25:46', '2023-07-25 03:25:46', NULL);
INSERT INTO `user_logs` VALUES (5352, 219122284, '2023-07-25', '183.15.206.81', 'Mozilla/5.0 (Linux; Android 13; V2243A Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/101.0.4951.74 Mobile Safari/537.36', '2023-07-25 03:57:49', '2023-07-25 03:57:49', NULL);
INSERT INTO `user_logs` VALUES (5419, 219122276, '2023-07-25', '172.105.235.23', 'PostmanRuntime/7.32.3', '2023-07-25 06:52:35', '2023-07-25 06:52:35', NULL);
INSERT INTO `user_logs` VALUES (5451, 219122282, '2023-07-25', '183.15.206.81', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-25 07:14:43', '2023-07-25 07:14:43', NULL);
INSERT INTO `user_logs` VALUES (5453, 219122285, '2023-07-25', '159.138.51.66', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-25 07:16:25', '2023-07-25 07:16:25', NULL);
INSERT INTO `user_logs` VALUES (5741, 219122287, '2023-07-25', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-25 13:12:40', '2023-07-25 13:12:40', NULL);
INSERT INTO `user_logs` VALUES (5752, 219122277, '2023-07-26', '183.8.142.48', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 00:12:15', '2023-07-26 00:12:15', NULL);
INSERT INTO `user_logs` VALUES (5753, 219122279, '2023-07-26', '183.15.206.81', 'Mozilla/5.0 (Linux; Android 12; V1981A Build/SP1A.210812.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Mobile Safari/537.36', '2023-07-26 01:55:47', '2023-07-26 01:55:47', NULL);
INSERT INTO `user_logs` VALUES (5756, 100004, '2023-07-26', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-26 01:57:44', '2023-07-26 01:57:44', NULL);
INSERT INTO `user_logs` VALUES (5767, 219122284, '2023-07-26', '183.15.206.81', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 02:16:59', '2023-07-26 02:16:59', NULL);
INSERT INTO `user_logs` VALUES (5769, 219122287, '2023-07-26', '43.132.98.37', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 02:19:25', '2023-07-26 02:19:25', NULL);
INSERT INTO `user_logs` VALUES (5777, 219122281, '2023-07-26', '43.132.98.31', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 02:47:48', '2023-07-26 02:47:48', NULL);
INSERT INTO `user_logs` VALUES (5778, 219122288, '2023-07-26', '43.132.98.31', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 02:49:04', '2023-07-26 02:49:04', NULL);
INSERT INTO `user_logs` VALUES (5788, 219122289, '2023-07-26', '203.91.85.247', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 03:42:39', '2023-07-26 03:42:39', NULL);
INSERT INTO `user_logs` VALUES (5792, 219122282, '2023-07-26', '183.15.206.81', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-26 03:53:42', '2023-07-26 03:53:42', NULL);
INSERT INTO `user_logs` VALUES (5800, 219122290, '2023-07-26', '43.132.98.34', 'Mozilla/5.0 (Linux; Android 13; 2106118C Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36', '2023-07-26 06:44:33', '2023-07-26 06:44:33', NULL);
INSERT INTO `user_logs` VALUES (5809, 219122280, '2023-07-26', '43.132.98.31', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-26 06:55:31', '2023-07-26 06:55:31', NULL);
INSERT INTO `user_logs` VALUES (5914, 219122291, '2023-07-26', '183.15.206.81', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-26 09:25:31', '2023-07-26 09:25:31', NULL);
INSERT INTO `user_logs` VALUES (5925, 219122292, '2023-07-26', '139.5.108.170', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-26 09:34:24', '2023-07-26 09:34:24', NULL);
INSERT INTO `user_logs` VALUES (6049, 219122279, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 12; V1981A Build/SP1A.210812.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Mobile Safari/537.36', '2023-07-27 01:13:42', '2023-07-27 01:13:42', NULL);
INSERT INTO `user_logs` VALUES (6056, 219122284, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 01:25:22', '2023-07-27 01:25:22', NULL);
INSERT INTO `user_logs` VALUES (6064, 219122291, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-27 01:30:39', '2023-07-27 01:30:39', NULL);
INSERT INTO `user_logs` VALUES (6066, 219122292, '2023-07-27', '159.138.56.94', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-27 01:32:41', '2023-07-27 01:32:41', NULL);
INSERT INTO `user_logs` VALUES (6080, 100004, '2023-07-27', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-27 01:37:24', '2023-07-27 01:37:24', NULL);
INSERT INTO `user_logs` VALUES (6104, 219122282, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-27 01:56:17', '2023-07-27 01:56:17', NULL);
INSERT INTO `user_logs` VALUES (6132, 219122307, '2023-07-27', '43.132.98.32', 'Mozilla/5.0 (iPhone; CPU iPhone OS 15_7_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 02:07:54', '2023-07-27 02:07:54', NULL);
INSERT INTO `user_logs` VALUES (6163, 219122293, '2023-07-27', '43.132.98.37', 'Mozilla/5.0 (Linux; Android 13; 2106118C Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36', '2023-07-27 02:27:37', '2023-07-27 02:27:37', NULL);
INSERT INTO `user_logs` VALUES (6210, 219122277, '2023-07-27', '156.59.197.10', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 03:02:33', '2023-07-27 03:02:33', NULL);
INSERT INTO `user_logs` VALUES (6265, 219122287, '2023-07-27', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 03:25:38', '2023-07-27 03:25:38', NULL);
INSERT INTO `user_logs` VALUES (6266, 219122280, '2023-07-27', '43.132.98.35', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 03:29:42', '2023-07-27 03:29:42', NULL);
INSERT INTO `user_logs` VALUES (6366, 219122278, '2023-07-27', '183.8.4.149', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-27 08:19:06', '2023-07-27 08:19:06', NULL);
INSERT INTO `user_logs` VALUES (6408, 219122295, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 13; V2243A Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/101.0.4951.74 Mobile Safari/537.36', '2023-07-27 09:30:21', '2023-07-27 09:30:21', NULL);
INSERT INTO `user_logs` VALUES (6424, 219122296, '2023-07-27', '116.30.143.106', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 09:43:26', '2023-07-27 09:43:26', NULL);
INSERT INTO `user_logs` VALUES (6431, 219122297, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; HLK-AL00 Build/HONORHLK-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-27 09:50:17', '2023-07-27 09:50:17', NULL);
INSERT INTO `user_logs` VALUES (6446, 219122298, '2023-07-27', '203.91.85.111', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-27 10:11:20', '2023-07-27 10:11:20', NULL);
INSERT INTO `user_logs` VALUES (6489, 219122281, '2023-07-27', '43.132.98.33', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 11:11:08', '2023-07-27 11:11:08', NULL);
INSERT INTO `user_logs` VALUES (6490, 219122299, '2023-07-27', '43.132.98.33', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 11:13:53', '2023-07-27 11:13:53', NULL);
INSERT INTO `user_logs` VALUES (6530, 219122289, '2023-07-27', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-27 12:24:25', '2023-07-27 12:24:25', NULL);
INSERT INTO `user_logs` VALUES (6579, 219122300, '2023-07-28', '43.132.98.32', 'Mozilla/5.0 (iPhone; CPU iPhone OS 15_7_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 01:26:46', '2023-07-28 01:26:46', NULL);
INSERT INTO `user_logs` VALUES (6583, 219122299, '2023-07-28', '43.132.98.32', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 01:39:05', '2023-07-28 01:39:05', NULL);
INSERT INTO `user_logs` VALUES (6586, 219122284, '2023-07-28', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 01:45:26', '2023-07-28 01:45:26', NULL);
INSERT INTO `user_logs` VALUES (6588, 219122298, '2023-07-28', '103.43.162.138', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-28 01:55:38', '2023-07-28 01:55:38', NULL);
INSERT INTO `user_logs` VALUES (6590, 219122277, '2023-07-28', '103.43.162.138', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-28 01:56:05', '2023-07-28 01:56:05', NULL);
INSERT INTO `user_logs` VALUES (6594, 219122282, '2023-07-28', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-28 02:12:37', '2023-07-28 02:12:37', NULL);
INSERT INTO `user_logs` VALUES (6599, 219122293, '2023-07-28', '43.132.98.33', 'Mozilla/5.0 (Linux; Android 13; 2106118C Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36', '2023-07-28 02:19:41', '2023-07-28 02:19:41', NULL);
INSERT INTO `user_logs` VALUES (6609, 100004, '2023-07-28', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-28 02:39:10', '2023-07-28 02:39:10', NULL);
INSERT INTO `user_logs` VALUES (6613, 219122295, '2023-07-28', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 13; V2243A Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/101.0.4951.74 Mobile Safari/537.36', '2023-07-28 02:46:30', '2023-07-28 02:46:30', NULL);
INSERT INTO `user_logs` VALUES (6618, 219122287, '2023-07-28', '43.132.98.32', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 02:55:25', '2023-07-28 02:55:25', NULL);
INSERT INTO `user_logs` VALUES (6624, 219122278, '2023-07-28', '103.43.162.138', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 03:00:34', '2023-07-28 03:00:34', NULL);
INSERT INTO `user_logs` VALUES (6642, 219122297, '2023-07-28', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; HLK-AL00 Build/HONORHLK-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-28 03:26:01', '2023-07-28 03:26:01', NULL);
INSERT INTO `user_logs` VALUES (6649, 219122289, '2023-07-28', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 03:40:55', '2023-07-28 03:40:55', NULL);
INSERT INTO `user_logs` VALUES (6693, 219122302, '2023-07-28', '203.91.85.91', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 06:25:15', '2023-07-28 06:25:15', NULL);
INSERT INTO `user_logs` VALUES (6713, 219122285, '2023-07-28', '223.104.68.192', 'Mozilla/5.0 (Linux; Android 12; OXF-AN00 Build/HUAWEIOXF-AN00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-28 06:54:39', '2023-07-28 06:54:39', NULL);
INSERT INTO `user_logs` VALUES (6717, 219122290, '2023-07-28', '43.132.98.33', 'Mozilla/5.0 (Linux; Android 13; 2106118C Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36', '2023-07-28 06:57:56', '2023-07-28 06:57:56', NULL);
INSERT INTO `user_logs` VALUES (6733, 219122303, '2023-07-28', '43.132.98.33', 'Mozilla/5.0 (Linux; Android 13; 2106118C Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36', '2023-07-28 07:15:39', '2023-07-28 07:15:39', NULL);
INSERT INTO `user_logs` VALUES (6772, 219122304, '2023-07-28', '43.132.98.33', 'Mozilla/5.0 (Linux; Android 13; 2106118C Build/TKQ1.220829.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36', '2023-07-28 08:15:19', '2023-07-28 08:15:19', NULL);
INSERT INTO `user_logs` VALUES (6890, 219122305, '2023-07-28', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-28 16:47:15', '2023-07-28 16:47:15', NULL);
INSERT INTO `user_logs` VALUES (6891, 219122278, '2023-07-29', '14.26.10.123', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 01:01:23', '2023-07-29 01:01:23', NULL);
INSERT INTO `user_logs` VALUES (6893, 219122296, '2023-07-29', '110.52.162.10', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 01:06:50', '2023-07-29 01:06:50', NULL);
INSERT INTO `user_logs` VALUES (6894, 219122299, '2023-07-29', '219.133.101.53', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 01:14:45', '2023-07-29 01:14:45', NULL);
INSERT INTO `user_logs` VALUES (6895, 219122305, '2023-07-29', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 01:44:42', '2023-07-29 01:44:42', NULL);
INSERT INTO `user_logs` VALUES (6897, 100004, '2023-07-29', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-29 01:46:04', '2023-07-29 01:46:04', NULL);
INSERT INTO `user_logs` VALUES (6931, 219122306, '2023-07-29', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 03:23:25', '2023-07-29 03:23:25', NULL);
INSERT INTO `user_logs` VALUES (6962, 219122289, '2023-07-29', '183.15.207.234', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 04:15:26', '2023-07-29 04:15:26', NULL);
INSERT INTO `user_logs` VALUES (7050, 219122282, '2023-07-29', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-29 07:52:03', '2023-07-29 07:52:03', NULL);
INSERT INTO `user_logs` VALUES (7051, 219122276, '2023-07-29', '172.105.235.23', 'PostmanRuntime/7.32.3', '2023-07-29 07:52:23', '2023-07-29 07:52:23', NULL);
INSERT INTO `user_logs` VALUES (7062, 219122297, '2023-07-29', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 10; HLK-AL00 Build/HONORHLK-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-29 08:40:50', '2023-07-29 08:40:50', NULL);
INSERT INTO `user_logs` VALUES (7088, 219122295, '2023-07-29', '183.15.207.234', 'Mozilla/5.0 (Linux; Android 13; V2243A Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/101.0.4951.74 Mobile Safari/537.36', '2023-07-29 09:33:29', '2023-07-29 09:33:29', NULL);
INSERT INTO `user_logs` VALUES (7101, 219122307, '2023-07-29', '219.133.101.53', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-29 13:45:11', '2023-07-29 13:45:11', NULL);
INSERT INTO `user_logs` VALUES (7112, 219122278, '2023-07-30', '139.5.108.135', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-30 00:49:29', '2023-07-30 00:49:29', NULL);
INSERT INTO `user_logs` VALUES (7118, 219122306, '2023-07-30', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-30 09:28:08', '2023-07-30 09:28:08', NULL);
INSERT INTO `user_logs` VALUES (7123, 219122277, '2023-07-30', '45.158.180.202', 'Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1', '2023-07-30 12:06:47', '2023-07-30 12:06:47', NULL);
INSERT INTO `user_logs` VALUES (7125, 219122296, '2023-07-30', '110.52.162.10', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-30 14:56:22', '2023-07-30 14:56:22', NULL);
INSERT INTO `user_logs` VALUES (7128, 100004, '2023-07-31', '172.105.235.23', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0', '2023-07-31 01:23:59', '2023-07-31 01:23:59', NULL);
INSERT INTO `user_logs` VALUES (7139, 219122282, '2023-07-31', '183.15.207.173', 'Mozilla/5.0 (Linux; Android 10; LRA-AL00 Build/HONORLRA-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.105 Mobile Safari/537.36', '2023-07-31 01:42:42', '2023-07-31 01:42:42', NULL);
INSERT INTO `user_logs` VALUES (7142, 219122276, '2023-07-31', '172.105.235.23', 'PostmanRuntime/7.32.3', '2023-07-31 02:05:03', '2023-07-31 02:05:03', NULL);
INSERT INTO `user_logs` VALUES (7157, 219122278, '2023-07-31', '14.26.10.123', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-31 02:15:54', '2023-07-31 02:15:54', NULL);
INSERT INTO `user_logs` VALUES (7158, 219122306, '2023-07-31', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-31 02:49:35', '2023-07-31 02:49:35', NULL);
INSERT INTO `user_logs` VALUES (7161, 219122295, '2023-07-31', '183.15.207.173', 'Mozilla/5.0 (Linux; Android 12; V1981A Build/SP1A.210812.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Mobile Safari/537.36', '2023-07-31 03:18:02', '2023-07-31 03:18:02', NULL);
INSERT INTO `user_logs` VALUES (7171, 219122289, '2023-07-31', '107.148.239.239', 'Mozilla/5.0 (iPhone; CPU iPhone OS 14_8_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) app/daxiaiOS', '2023-07-31 03:33:02', '2023-07-31 03:33:02', NULL);

SET FOREIGN_KEY_CHECKS = 1;
