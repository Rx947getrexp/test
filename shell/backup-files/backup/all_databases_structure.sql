-- MySQL dump 10.13  Distrib 8.0.39, for Linux (x86_64)
--
-- Host: localhost    Database: speed
-- ------------------------------------------------------
-- Server version	8.0.39-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `speed`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `speed` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `speed`;

--
-- Table structure for table `admin_res`
--

DROP TABLE IF EXISTS `admin_res`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_res` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资源名称',
  `res_type` int NOT NULL COMMENT '类型：1-菜单；2-接口；3-按钮',
  `pid` int NOT NULL COMMENT '上级id（没有默认为0）',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'url地址',
  `sort` int NOT NULL COMMENT '排序',
  `icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '图标',
  `is_del` int NOT NULL COMMENT '软删状态：0-未删（默认）；1-已删',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `res_path` (`url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='资源表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `admin_role`
--

DROP TABLE IF EXISTS `admin_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `is_del` int NOT NULL DEFAULT '0' COMMENT '0-正常；1-软删',
  `is_used` int NOT NULL COMMENT '1-已启用；2-未启用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `admin_role_res`
--

DROP TABLE IF EXISTS `admin_role_res`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_role_res` (
  `role_id` int NOT NULL COMMENT '角色id',
  `res_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资源id列表',
  `res_tree` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '资源菜单json树',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色资源表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `admin_user`
--

DROP TABLE IF EXISTS `admin_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `pwd2` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '二级密码',
  `authkey` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '谷歌验证码私钥',
  `status` int DEFAULT '0' COMMENT '冻结状态：0-正常；1-冻结',
  `is_del` int DEFAULT '0' COMMENT '0-正常；1-软删',
  `is_reset` int DEFAULT NULL COMMENT '0-否；1-代表需要重置两步验证码',
  `is_first` int DEFAULT NULL COMMENT '0-否；1-代表首次登录需要修改密码',
  `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '可查看渠道范围,为空则可查看所有范围',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uname_index` (`uname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100031 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='后台用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `admin_user_role`
--

DROP TABLE IF EXISTS `admin_user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_user_role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增序列',
  `uid` int NOT NULL COMMENT '用户id',
  `role_id` int NOT NULL COMMENT '角色id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `is_del` int NOT NULL COMMENT '软删：0-未删；1-已删',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `u_role_user` (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_activity`
--

DROP TABLE IF EXISTS `t_activity`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_activity` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `status` int NOT NULL COMMENT '状态:1-success；2-fail',
  `gift_sec` int NOT NULL COMMENT '赠送时间（失败为0）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1826307487830118401 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='免费领会员活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_ad`
--

DROP TABLE IF EXISTS `t_ad`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_ad` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `status` int NOT NULL COMMENT '状态:1-上架；2-下架',
  `sort` int NOT NULL COMMENT '排序',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '广告名称',
  `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '广告logo',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '广告链接',
  `ad_type` int NOT NULL COMMENT '广告分类：1-社交；2-游戏；3-漫画；4-视频...',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标签标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '正文介绍',
  `author` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='广告表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_app_dns`
--

DROP TABLE IF EXISTS `t_app_dns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_app_dns` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `site_type` int NOT NULL COMMENT '站点类型:1-appapi域名；2-offsite域名；3-onlineservice在线客服；4-管理后台...',
  `dns` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '域名',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
  `level` int NOT NULL COMMENT '线路级别:1,2,3...用于白名单机制',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app域名表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_app_version`
--

DROP TABLE IF EXISTS `t_app_version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_app_version` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `app_type` int NOT NULL COMMENT '1-ios;2-安卓；3-h5zip',
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本号',
  `link` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '超链地址',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=237 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app版本管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_channel`
--

DROP TABLE IF EXISTS `t_channel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_channel` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '渠道名称',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '渠道编号',
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '渠道链接',
  `status` int DEFAULT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推广渠道表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_country`
--

DROP TABLE IF EXISTS `t_country`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_country` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '国家名称',
  `name_cn` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '国家名称中文',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k1` (`name`) USING BTREE,
  UNIQUE KEY `uiq_k2` (`name_cn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='国家名称表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_dev`
--

DROP TABLE IF EXISTS `t_dev`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_dev` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `client_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `os` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户端设备系统os',
  `is_send` int DEFAULT NULL COMMENT '1-已赠送时间；2-未赠送',
  `network` int DEFAULT '1' COMMENT '网络模式（1-自动；2-手动）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1826307620055552001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='客户端设备表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_dict`
--

DROP TABLE IF EXISTS `t_dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_dict` (
  `key_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '键',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '值',
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `is_del` int DEFAULT NULL COMMENT '0-正常；1-软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`key_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_doc`
--

DROP TABLE IF EXISTS `t_doc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_doc` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `content` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'content',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_t_n` (`type`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='官网文档';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_gift`
--

DROP TABLE IF EXISTS `t_gift`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_gift` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `op_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '业务id',
  `op_uid` bigint DEFAULT NULL COMMENT '业务uid',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '赠送标题',
  `gift_sec` int NOT NULL COMMENT '赠送时间（单位s）',
  `g_type` int NOT NULL COMMENT '赠送类别（1-注册；2-推荐；3-日常活动；4-充值）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29860 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='赠送用户记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_goods`
--

DROP TABLE IF EXISTS `t_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_goods` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `m_type` int NOT NULL COMMENT '会员类型：1-vip1；2-vip2',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '套餐标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '套餐标题（英文）',
  `title_rus` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '套餐标题（俄文）',
  `price` decimal(10,6) NOT NULL COMMENT '单价(U)',
  `price_unit` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '价格单位',
  `period` int NOT NULL COMMENT '有效期（天）',
  `dev_limit` int NOT NULL COMMENT '设备限制数',
  `flow_limit` bigint NOT NULL COMMENT '流量限制数；单位：字节；0-不限制',
  `is_discount` int DEFAULT NULL COMMENT '是否优惠:1-是；2-否',
  `low` int DEFAULT NULL COMMENT '最低赠送(天)',
  `high` int DEFAULT NULL COMMENT '最高赠送(天)',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  `usd_pay_price` decimal(10,6) NOT NULL COMMENT 'usd_pay价格(U)',
  `usd_price_unit` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'USD支付的价格单位',
  `webmoney_pay_price` decimal(10,2) NOT NULL COMMENT 'webmoney价格(wmz)',
  `price_rub` decimal(10,2) NOT NULL COMMENT '卢布价格(RUB)',
  `price_wmz` decimal(10,2) NOT NULL COMMENT 'WMZ价格(WMZ)',
  `price_usd` decimal(10,2) NOT NULL COMMENT 'USD价格(USD)',
  `price_uah` decimal(10,2) NOT NULL COMMENT 'UAH价格(UAH)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品套餐表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_ios_account`
--

DROP TABLE IF EXISTS `t_ios_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_ios_account` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号id',
  `account` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ios账号',
  `pass` varchar(64) CHARACTER SET utf16 COLLATE utf16_general_ci NOT NULL COMMENT '密码',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '别名',
  `account_type` int DEFAULT NULL COMMENT '1-国区；2-海外',
  `status` int DEFAULT NULL COMMENT '1-正常；2-下架',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `ios_account` (`account`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='ios账号管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_node`
--

DROP TABLE IF EXISTS `t_node`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_node` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点名称',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '节点标题（英文）',
  `title_rus` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '节点标题（俄文)',
  `country` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '国家',
  `country_en` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '国家（英文）',
  `country_rus` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '国家（俄文)',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内网IP',
  `server` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公网域名',
  `node_type` int DEFAULT NULL COMMENT '节点类别:1-常规；2-高带宽...(根据情况而定)',
  `port` int NOT NULL COMMENT '公网端口',
  `cpu` int NOT NULL COMMENT 'cpu核数量（单位个）',
  `flow` bigint NOT NULL COMMENT '流量带宽',
  `disk` bigint NOT NULL COMMENT '磁盘容量（单位B）',
  `memory` bigint NOT NULL COMMENT '内存大小（单位B）',
  `min_port` int DEFAULT NULL COMMENT '最小端口',
  `max_port` int DEFAULT NULL COMMENT '最大端口',
  `path` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ws路径',
  `is_recommend` int DEFAULT NULL COMMENT '推荐节点1-是；2-否',
  `channel_id` int DEFAULT NULL COMMENT '市场渠道（默认0）-优选节点有效',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  `weight` int unsigned NOT NULL DEFAULT '0' COMMENT '权重',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100022 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='节点表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_node_dns`
--

DROP TABLE IF EXISTS `t_node_dns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_node_dns` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `node_id` bigint NOT NULL COMMENT '节点id',
  `dns` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '域名',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
  `level` int NOT NULL COMMENT '线路级别:1,2,3...用于白名单机制',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  `is_machine` int NOT NULL DEFAULT '0' COMMENT '是否为真实机器',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='节点域名表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_node_uuid`
--

DROP TABLE IF EXISTS `t_node_uuid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_node_uuid` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `node_id` bigint NOT NULL COMMENT '节点id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '节点邮箱，用于区分流量',
  `v2ray_uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点UUID',
  `server` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '公网域名',
  `port` int NOT NULL COMMENT '公网端口',
  `used_flow` bigint NOT NULL COMMENT '已使用流量（单位B）',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='节点UUID表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_notice`
--

DROP TABLE IF EXISTS `t_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_notice` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `title_en` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标题（英文）',
  `title_rus` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标题（俄文）',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标签',
  `tag_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标签（英文）',
  `tag_rus` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标签（俄文）',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '正文内容',
  `content_en` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '正文内容（英文）',
  `content_rus` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '正文内容（俄文）',
  `author` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `status` int DEFAULT NULL COMMENT '状态:1-发布；2-软删',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推荐关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_order`
--

DROP TABLE IF EXISTS `t_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `goods_id` bigint NOT NULL COMMENT '商品id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品标题',
  `price` decimal(10,6) NOT NULL COMMENT '单价(U)',
  `price_cny` decimal(10,2) NOT NULL COMMENT '折合RMB单价(CNY)',
  `status` int NOT NULL COMMENT '订单状态:1-init；2-success；3-cancel',
  `finished_at` timestamp NULL DEFAULT NULL COMMENT '完成时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1784410600542048257 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='商品订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_pay_order`
--

DROP TABLE IF EXISTS `t_pay_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_pay_order` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint unsigned NOT NULL COMMENT '用户uid',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户邮箱',
  `order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
  `order_amount` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易金额',
  `currency` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '交易币种',
  `pay_type_code` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付类型编码',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '状态:1-正常；2-已软删',
  `return_status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付平台返回的结果',
  `status_mes` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '状态:1-正常；2-已软删',
  `order_data` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建订单时支付平台返回的信息',
  `result_status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '查询结果，实际订单状态',
  `order_reality_amount` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '实际交易金额',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `version` int DEFAULT '1' COMMENT '数据版本号',
  `payment_channel_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付通道ID',
  `goods_id` int NOT NULL COMMENT '套餐ID',
  `payment_proof` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '支付凭证地址',
  `commission` decimal(10,2) DEFAULT NULL COMMENT '手续费',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k1` (`order_no`) USING BTREE,
  KEY `k1` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='支付订单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_payment_channel`
--

DROP TABLE IF EXISTS `t_payment_channel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_payment_channel` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `channel_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付通道ID',
  `channel_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付通道名称',
  `is_active` tinyint(1) NOT NULL DEFAULT '1' COMMENT '支付通道是否可用，1表示可用,2表示不可用',
  `free_trial_days` int NOT NULL DEFAULT '3' COMMENT '赠送的免费时长（以天为单位）',
  `timeout_duration` int NOT NULL DEFAULT '30' COMMENT '订单未支付超时关闭时间（单位分钟）',
  `payment_qr_code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '支付收款码. eg: U支付收款码',
  `payment_qr_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '支付收款链接',
  `bank_card_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '银行卡信息',
  `customer_service_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '客服信息',
  `mer_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'mer_no',
  `pay_type_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'pay_type_code',
  `weight` int NOT NULL DEFAULT '0' COMMENT '权重，排序使用',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `usd_network` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'USD支付网络',
  `currency_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '支付渠道币种',
  `freekassa_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'freekassa支付通道',
  `commission_rate` decimal(10,2) NOT NULL COMMENT '手续费比例',
  `commission` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `min_pay_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '最低支付金额',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_id` (`channel_id`) USING BTREE,
  UNIQUE KEY `uiq_name` (`channel_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='支付通道表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_serving_country`
--

DROP TABLE IF EXISTS `t_serving_country`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_serving_country` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '国家名称，不可以修改，作为ID用',
  `display` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用于在用户侧展示的国家名称',
  `logo_link` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '国家图片地址',
  `ping_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '前端使用',
  `is_recommend` int DEFAULT NULL COMMENT '推荐节点1-是；2-否',
  `weight` int NOT NULL DEFAULT '0' COMMENT '权重',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态:1-正常；2-已软删',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `level` int DEFAULT NULL COMMENT '等级：0-所有用户都可以选择；1-青铜、铂金会员可选择；2-铂金会员可选择',
  `is_free` int DEFAULT '0' COMMENT '是否为免费站点，0: 不免费,1: 免费',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k1` (`name`) USING BTREE,
  UNIQUE KEY `uiq_k2` (`display`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='上架国家表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_site`
--

DROP TABLE IF EXISTS `t_site`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_site` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `site` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '域名',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip',
  `status` int DEFAULT NULL COMMENT '1-正常；2-软删',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作者',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='域名表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_success_record`
--

DROP TABLE IF EXISTS `t_success_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_success_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `order_id` bigint NOT NULL COMMENT '订单id',
  `start_time` bigint NOT NULL COMMENT '本次计费开始时间戳',
  `end_time` bigint NOT NULL COMMENT '本次计费结束时间戳',
  `surplus_sec` bigint NOT NULL COMMENT '剩余时长(s)',
  `total_sec` bigint DEFAULT NULL COMMENT '订单总时长(s）',
  `goods_day` int DEFAULT NULL COMMENT '套餐天数',
  `send_day` int DEFAULT NULL COMMENT '赠送天数',
  `pay_type` int NOT NULL COMMENT '订单状态:1-银行卡；2-支付宝；3-微信支付',
  `status` int DEFAULT NULL COMMENT '1-using使用中；2-wait等待; 3-end已结束',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户购买成功记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_upload_log`
--

DROP TABLE IF EXISTS `t_upload_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_upload_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日志内容',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=139 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='上传日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user`
--

DROP TABLE IF EXISTS `t_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '电话',
  `level` int DEFAULT NULL COMMENT '等级：0-vip0；1-vip1；2-vip2',
  `expired_time` bigint DEFAULT NULL COMMENT 'vip到期时间',
  `v2ray_uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '节点UUID',
  `v2ray_tag` int DEFAULT NULL COMMENT 'v2ray存在UUID标签:1-有；2-无',
  `channel` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `channel_id` int DEFAULT NULL COMMENT '渠道id',
  `status` int DEFAULT NULL COMMENT '冻结状态：0-正常；1-冻结',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  `client_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '最近一次登录的ip',
  `last_login_country` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '最近一次登录的国家',
  `preferred_country` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '用户选择的国家（国家名称）',
  `version` int DEFAULT '1' COMMENT '数据版本号',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `p_uname_index` (`uname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=219159672 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_cancelled`
--

DROP TABLE IF EXISTS `t_user_cancelled`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_cancelled` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户密码',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '电话',
  `level` int DEFAULT NULL COMMENT '等级：0-vip0；1-vip1；2-vip2',
  `expired_time` bigint DEFAULT NULL COMMENT 'vip到期时间',
  `v2ray_uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '节点UUID',
  `v2ray_tag` int DEFAULT NULL COMMENT 'v2ray存在UUID标签:1-有；2-无',
  `channel` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `channel_id` int DEFAULT NULL COMMENT '渠道id',
  `status` int DEFAULT NULL COMMENT '冻结状态：0-正常；1-冻结',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  `client_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=219159564 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='已经注销的用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_config`
--

DROP TABLE IF EXISTS `t_user_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_config` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint NOT NULL COMMENT 'id',
  `node_id` bigint NOT NULL COMMENT 'ID',
  `status` int NOT NULL COMMENT ':1-2-',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=682 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_dev`
--

DROP TABLE IF EXISTS `t_user_dev`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_dev` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `status` int NOT NULL COMMENT '状态:1-正常；2-已踢',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23436 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户设备表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_device`
--

DROP TABLE IF EXISTS `t_user_device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_device` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint unsigned NOT NULL COMMENT '用户uid',
  `client_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `os` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户端设备系统os',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_id` (`user_id`,`client_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=959149 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='客户端设备表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_team`
--

DROP TABLE IF EXISTS `t_user_team`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_team` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `direct_id` bigint NOT NULL COMMENT '上级id',
  `direct_tree` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '上级列表',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35963 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推荐关系表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_traffic`
--

DROP TABLE IF EXISTS `t_user_traffic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_traffic` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `uplink` bigint unsigned NOT NULL COMMENT '上行流量',
  `downlink` bigint unsigned NOT NULL COMMENT '下行流量',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`email`,`ip`,`date`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=86135 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户流量采集记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_traffic_log`
--

DROP TABLE IF EXISTS `t_user_traffic_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_traffic_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
  `date_time` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '数据采集时间',
  `uplink` bigint unsigned NOT NULL COMMENT '上行流量',
  `downlink` bigint unsigned NOT NULL COMMENT '下行流量',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `k_1` (`email`,`ip`,`date_time`) USING BTREE,
  KEY `k_2` (`email`,`date_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1524238 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户流量采集流水记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_vip_attr_record`
--

DROP TABLE IF EXISTS `t_user_vip_attr_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_vip_attr_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户邮箱',
  `level_from` int DEFAULT NULL COMMENT '等级：0-vip0；1-vip1；2-vip2',
  `level_to` int DEFAULT NULL COMMENT '等级：0-vip0；1-vip1；2-vip2',
  `source` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '来源',
  `order_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单号',
  `expired_time_from` int DEFAULT NULL COMMENT '会员到期时间-原值',
  `expired_time_to` int DEFAULT NULL COMMENT '会员到期时间-新值',
  `desc` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '记录描述',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `is_revert` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被回滚',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k1` (`email`,`order_no`),
  KEY `k1` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=870 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户会员属性变更记录';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_work_log`
--

DROP TABLE IF EXISTS `t_work_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_work_log` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `mode_type` int NOT NULL COMMENT '模式类别:1-智能；2-手选',
  `node_id` bigint NOT NULL COMMENT '工作节点',
  `flow` bigint NOT NULL COMMENT '使用流量（字节）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='工作日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_work_mode`
--

DROP TABLE IF EXISTS `t_work_mode`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_work_mode` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `dev_id` bigint NOT NULL COMMENT '设备id',
  `mode_type` int NOT NULL COMMENT '模式类别:1-智能；2-手选',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `u_dev` (`dev_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='工作模式';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_logs`
--

DROP TABLE IF EXISTS `user_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `datestr` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期',
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'IP地址',
  `user_agent` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求头user-agent',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注信息',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `log_user_date` (`user_id`,`datestr`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1391507 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日志表（仅记录第一次事件)';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Current Database: `go_fly2`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `go_fly2` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `go_fly2`;

--
-- Table structure for table `about`
--

DROP TABLE IF EXISTS `about`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `about` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
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
  UNIQUE KEY `page` (`page`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `config`
--

DROP TABLE IF EXISTS `config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `conf_name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `conf_key` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `conf_value` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `conf_key` (`conf_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ipblack`
--

DROP TABLE IF EXISTS `ipblack`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ipblack` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `kefu_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `ip` (`ip`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `land_page`
--

DROP TABLE IF EXISTS `land_page`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `land_page` (
  `id` int NOT NULL,
  `title` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `keyword` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `language` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `page_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `language`
--

DROP TABLE IF EXISTS `language`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `language` (
  `id` int NOT NULL,
  `country` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `short_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
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
  KEY `kefu_id` (`kefu_id`) USING BTREE,
  KEY `visitor_id` (`visitor_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1304 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `reply_group`
--

DROP TABLE IF EXISTS `reply_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reply_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `user_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `reply_item`
--

DROP TABLE IF EXISTS `reply_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reply_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `group_id` int NOT NULL DEFAULT '0',
  `user_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `item_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `method` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `path` varchar(2048) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `password` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `avator` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_client`
--

DROP TABLE IF EXISTS `user_client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_client` (
  `id` int NOT NULL AUTO_INCREMENT,
  `kefu` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `client_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user` (`kefu`,`client_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_role`
--

DROP TABLE IF EXISTS `user_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL DEFAULT '0',
  `role_id` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `visitor`
--

DROP TABLE IF EXISTS `visitor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visitor` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `avator` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `source_ip` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `to_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `visitor_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT '0',
  `refer` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `city` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `client_ip` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `extra` varchar(2048) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `visitor_id` (`visitor_id`) USING BTREE,
  KEY `to_id` (`to_id`) USING BTREE,
  KEY `idx_update` (`updated_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1297 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `welcome`
--

DROP TABLE IF EXISTS `welcome`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `welcome` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `keyword` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `content` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '',
  `is_default` tinyint unsigned NOT NULL DEFAULT '0',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `keyword` (`keyword`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Current Database: `speed_report`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `speed_report` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `speed_report`;

--
-- Table structure for table `t_user_channel_day`
--

DROP TABLE IF EXISTS `t_user_channel_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_channel_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '推广渠道id',
  `total` int unsigned NOT NULL COMMENT '推广渠道用户总量',
  `new` int unsigned NOT NULL COMMENT '推广渠道新增用户',
  `retained` int unsigned NOT NULL COMMENT '留存',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  `total_recharge` int unsigned NOT NULL COMMENT '充值总人数',
  `total_recharge_money` decimal(10,2) unsigned NOT NULL COMMENT '充值总金额',
  `new_recharge_money` decimal(10,2) unsigned NOT NULL COMMENT '新增充值金额',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`channel`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6963 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_channel_recharge_day`
--

DROP TABLE IF EXISTS `t_user_channel_recharge_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_channel_recharge_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '渠道id',
  `goods_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品套餐名称',
  `usd_total` int unsigned NOT NULL COMMENT 'USD充值次数',
  `usd_new` int unsigned NOT NULL COMMENT '新增USD充值次数',
  `rub_total` int unsigned NOT NULL COMMENT 'RUB充值次数',
  `rub_new` int unsigned NOT NULL COMMENT '新增RUB充值次数',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`channel`,`goods_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3828 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_device_action_day`
--

DROP TABLE IF EXISTS `t_user_device_action_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_device_action_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `device` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备类型',
  `total_clicks` int unsigned NOT NULL COMMENT '总点击次数',
  `yesterday_day_clicks` int unsigned NOT NULL COMMENT '次日点击次数',
  `weekly_clicks` int unsigned NOT NULL COMMENT '周点击次数',
  `total_users_clicked` int unsigned NOT NULL COMMENT '总点击人数',
  `yesterday_day_users_clicked` int unsigned NOT NULL COMMENT '次日点击人数',
  `weekly_users_clicked` int unsigned NOT NULL COMMENT '周点击人数',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`device`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=169 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户设备前端点击日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_device_day`
--

DROP TABLE IF EXISTS `t_user_device_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_device_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `device` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备类型',
  `total` int unsigned NOT NULL COMMENT '设备类型用户总量',
  `new` int unsigned NOT NULL COMMENT '设备类型新增用户',
  `retained` int unsigned NOT NULL COMMENT '设备类型留存',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  `total_recharge` int unsigned NOT NULL COMMENT '设备类型充值总人数',
  `total_recharge_money` decimal(10,2) unsigned NOT NULL COMMENT '设备类型充值总金额',
  `new_recharge_money` decimal(10,2) unsigned NOT NULL COMMENT '设备类型新增充值金额',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`device`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户设备日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_node_day`
--

DROP TABLE IF EXISTS `t_user_node_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_node_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点',
  `total` int unsigned NOT NULL COMMENT '节点用户使用总量',
  `new` int NOT NULL COMMENT '推广渠道新增用户',
  `retained` int unsigned NOT NULL COMMENT '节点用户使用次日留存',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`ip`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=757 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_node_online_day`
--

DROP TABLE IF EXISTS `t_user_node_online_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_node_online_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
  `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `online_duration` int unsigned NOT NULL COMMENT '在线时间长度',
  `uplink` bigint unsigned NOT NULL COMMENT '上行流量',
  `downlink` bigint unsigned NOT NULL COMMENT '下行流量',
  `node` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '节点',
  `register_date` timestamp NOT NULL COMMENT '用户注册日期',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`email`,`node`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=74518 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户登陆统计日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_online_day`
--

DROP TABLE IF EXISTS `t_user_online_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_online_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮件',
  `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `online_duration` int unsigned NOT NULL COMMENT '在线时间长度',
  `uplink` bigint unsigned NOT NULL COMMENT '上行流量',
  `downlink` bigint unsigned NOT NULL COMMENT '下行流量',
  `last_login_country` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=71058 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户登陆统计日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_op_log`
--

DROP TABLE IF EXISTS `t_user_op_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_op_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户账号',
  `device_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备ID',
  `device_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备类型',
  `page_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'page_name',
  `result` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'result',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'content',
  `version` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '提交时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `key1` (`email`) USING BTREE,
  KEY `key2` (`device_id`) USING BTREE,
  KEY `key4` (`created_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1277470 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户操作轨迹日志';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_ping`
--

DROP TABLE IF EXISTS `t_user_ping`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_ping` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `host` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '节点host, ip or dns',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ping的结果',
  `cost` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ping耗时',
  `time` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '上报时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `key1` (`email`,`created_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=363632 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户ping节点结果上报数据';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_recharge_report_day`
--

DROP TABLE IF EXISTS `t_user_recharge_report_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_recharge_report_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `goods_id` int unsigned NOT NULL COMMENT '商品套餐id',
  `total` int unsigned NOT NULL COMMENT '用户充值总量',
  `new` int unsigned NOT NULL COMMENT '新增用户充值数量',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`goods_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=914 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_recharge_times_report_day`
--

DROP TABLE IF EXISTS `t_user_recharge_times_report_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_recharge_times_report_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `goods_id` int unsigned NOT NULL COMMENT '商品套餐id',
  `total` int unsigned NOT NULL COMMENT '用户充值次数总量',
  `new` int unsigned NOT NULL COMMENT '新增用户充值次数总量',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`goods_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=912 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户日报表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `t_user_report_day`
--

DROP TABLE IF EXISTS `t_user_report_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `t_user_report_day` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `date` int unsigned NOT NULL COMMENT '数据日期, 20230101',
  `channel_id` int unsigned NOT NULL COMMENT '渠道id',
  `total` int unsigned NOT NULL COMMENT '用户总量',
  `new` int unsigned NOT NULL COMMENT '新增用户',
  `retained` int unsigned NOT NULL COMMENT '留存',
  `month_retained` int unsigned NOT NULL COMMENT '月留存',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uiq_k` (`date`,`channel_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=276 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日报表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-08-21 20:17:42
