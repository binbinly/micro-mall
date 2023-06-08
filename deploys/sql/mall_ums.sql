-- 创建数据库
CREATE DATABASE IF NOT EXISTS `mall_ums`;
USE `mall_ums`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ums_collect_spu
-- ----------------------------
DROP TABLE IF EXISTS `ums_collect_spu`;
CREATE TABLE `ums_collect_spu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `spu_id` bigint(20) NOT NULL COMMENT 'spu_id',
  `spu_name` varchar(255) NOT NULL COMMENT 'spu_name',
  `spu_img` varchar(255) NOT NULL COMMENT 'spu_img',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_ums_collect_spu_member_id` (`member_id`),
  KEY `idx_ums_collect_spu_spu_id` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_collect_spu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_collect_subject
-- ----------------------------
DROP TABLE IF EXISTS `ums_collect_subject`;
CREATE TABLE `ums_collect_subject` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `subject_id` bigint(20) NOT NULL COMMENT '活动id',
  `subject_name` varchar(255) NOT NULL COMMENT '活动名',
  `subject_img` varchar(255) NOT NULL COMMENT '活动封面',
  `subject_url` varchar(255) NOT NULL COMMENT '活动url',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_ums_collect_subject_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_collect_subject
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_growth_log
-- ----------------------------
DROP TABLE IF EXISTS `ums_growth_log`;
CREATE TABLE `ums_growth_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `change_count` bigint(20) NOT NULL COMMENT '改变的值（正负计数）',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `source_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '积分来源[0-购物，1-管理员修改]',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_ums_growth_log_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_growth_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_integration_log
-- ----------------------------
DROP TABLE IF EXISTS `ums_integration_log`;
CREATE TABLE `ums_integration_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `change_count` bigint(20) NOT NULL COMMENT '改变的值（正负计数）',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `source_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '积分来源[0-购物，1-管理员修改]',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_ums_integration_log_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_integration_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_login_log
-- ----------------------------
DROP TABLE IF EXISTS `ums_login_log`;
CREATE TABLE `ums_login_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `ip` varchar(15) NOT NULL COMMENT 'ip地址',
  `city` varchar(64) NOT NULL DEFAULT '' COMMENT '城市',
  `type` tinyint(4) NOT NULL COMMENT '登录类型[1-web，2-app]',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_ums_login_log_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_member
-- ----------------------------
DROP TABLE IF EXISTS `ums_member`;
CREATE TABLE `ums_member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '会员等级id',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',
  `phone` bigint(20) NOT NULL COMMENT '手机号',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(4) NOT NULL DEFAULT '1' COMMENT '性别',
  `birth` date DEFAULT NULL COMMENT '生日',
  `area` varchar(255) NOT NULL DEFAULT '' COMMENT '城市',
  `job` varchar(255) NOT NULL DEFAULT '' COMMENT '职业',
  `source_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '用户来源',
  `integration` bigint(20) NOT NULL DEFAULT '0' COMMENT '积分',
  `growth` bigint(20) NOT NULL DEFAULT '0' COMMENT '成长值',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
  `sign` varchar(255) NOT NULL DEFAULT '' COMMENT '签名',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_ums_member_username` (`username`),
  UNIQUE KEY `idx_ums_member_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_member
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_member_address
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_address`;
CREATE TABLE `ums_member_address` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `name` varchar(30) NOT NULL COMMENT '收货人姓名',
  `phone` char(11) NOT NULL COMMENT '收货人手机号',
  `province` varchar(30) NOT NULL COMMENT '省',
  `city` varchar(30) NOT NULL COMMENT '市',
  `county` varchar(30) NOT NULL COMMENT '区/县',
  `area_code` int(10) unsigned NOT NULL COMMENT '最后一级地区编码',
  `detail` varchar(255) NOT NULL COMMENT '详细地址',
  `is_default` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否默认',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_ums_member_address_member_id` (`member_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_member_address
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_member_level
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_level`;
CREATE TABLE `ums_member_level` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(100) NOT NULL COMMENT '等级名称',
  `growth` bigint(20) NOT NULL COMMENT '等级需要的成长值',
  `free_freight` bigint(20) NOT NULL COMMENT '免运费标准',
  `comment_growth` bigint(20) NOT NULL COMMENT '评论获取的成长值',
  `is_freight` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有免邮特权',
  `is_member` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有会员价格特权',
  `is_birthday` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有生日特权',
  `is_default` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否默认等级',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_member_level
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ums_member_stat
-- ----------------------------
DROP TABLE IF EXISTS `ums_member_stat`;
CREATE TABLE `ums_member_stat` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `total_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '累计消费金额/分',
  `coupon_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '累计优惠金额/分',
  `order_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单数',
  `coupon_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '优惠券数',
  `comment_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '评价数',
  `return_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '退货数',
  `login_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `follow_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '关注数',
  `fens_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '粉丝数',
  `collect_product_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '收藏商品数',
  `collect_subject_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '收藏专题活动数',
  `collect_comment_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '收藏评论数',
  `invite_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀请好友数',
  PRIMARY KEY (`id`),
  KEY `idx_ums_member_stat_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ums_member_stat
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
