/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.8.76
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 192.168.8.76:3306
 Source Schema         : mall_wms

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 22/11/2021 21:51:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for wms_purchase
-- ----------------------------
DROP TABLE IF EXISTS `wms_purchase`;
CREATE TABLE `wms_purchase` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `assignee_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '采购人id',
  `assignee_name` varchar(64) NOT NULL DEFAULT '' COMMENT '采购人',
  `phone` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `priority` tinyint(4) NOT NULL DEFAULT '0' COMMENT '优先级',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `ware_id` bigint(20) NOT NULL COMMENT '仓库id',
  `amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '总金额/分',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wms_purchase
-- ----------------------------
BEGIN;
INSERT INTO `wms_purchase` VALUES (1, 1, '张三', '13333333333', 1, 3, 1, 0, 1633179593, 1633182541);
COMMIT;

-- ----------------------------
-- Table structure for wms_purchase_detail
-- ----------------------------
DROP TABLE IF EXISTS `wms_purchase_detail`;
CREATE TABLE `wms_purchase_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `purchase_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '采购单id',
  `sku_id` bigint(20) NOT NULL COMMENT '采购商品id',
  `sku_num` bigint(20) NOT NULL COMMENT '采购数量',
  `sku_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '采购金额/分',
  `ware_id` bigint(20) NOT NULL COMMENT '仓库id',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态[0新建，1已分配，2正在采购，3已完成，4采购失败]',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_wms_purchase_detail_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wms_purchase_detail
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wms_ware
-- ----------------------------
DROP TABLE IF EXISTS `wms_ware`;
CREATE TABLE `wms_ware` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL COMMENT '仓库名',
  `address` varchar(255) NOT NULL COMMENT '仓库地址',
  `area_code` int(10) unsigned NOT NULL COMMENT '最后一级地区编码',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wms_ware
-- ----------------------------
BEGIN;
INSERT INTO `wms_ware` VALUES (1, '1号仓库', '北京', 1, 1633161615, 1633161615);
INSERT INTO `wms_ware` VALUES (2, '2号仓库', '上海', 2, 1633161628, 1633161628);
INSERT INTO `wms_ware` VALUES (3, '3号仓库', '广州', 3, 1633161643, 1633161643);
INSERT INTO `wms_ware` VALUES (4, '4号仓库', '深圳', 4, 1633161658, 1633161658);
COMMIT;

-- ----------------------------
-- Table structure for wms_ware_sku
-- ----------------------------
DROP TABLE IF EXISTS `wms_ware_sku`;
CREATE TABLE `wms_ware_sku` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sku_id` bigint(20) NOT NULL COMMENT '采购商品id',
  `ware_id` bigint(20) NOT NULL COMMENT '仓库id',
  `sku_name` varchar(255) NOT NULL DEFAULT '' COMMENT '采购商品名',
  `stock` bigint(20) NOT NULL DEFAULT '0' COMMENT '库存',
  `stock_lock` bigint(20) NOT NULL DEFAULT '0' COMMENT '锁定库存',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_wms_ware_sku_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2667 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wms_ware_sku
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wms_ware_task
-- ----------------------------
DROP TABLE IF EXISTS `wms_ware_task`;
CREATE TABLE `wms_ware_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` bigint(20) NOT NULL COMMENT 'order_id',
  `order_no` varchar(32) NOT NULL COMMENT '单号',
  `consignee` varchar(64) NOT NULL COMMENT '收货人',
  `phone` char(11) NOT NULL COMMENT '收货人手机号',
  `address` varchar(255) NOT NULL COMMENT '配送地址',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '任务备注',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 1=锁定 2=解锁 3=扣减',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_wms_ware_task_order_id` (`order_id`),
  KEY `idx_wms_ware_task_order_no` (`order_no`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wms_ware_task
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wms_ware_task_detail
-- ----------------------------
DROP TABLE IF EXISTS `wms_ware_task_detail`;
CREATE TABLE `wms_ware_task_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `task_id` bigint(20) NOT NULL COMMENT '工作单id',
  `ware_id` bigint(20) NOT NULL COMMENT '仓库id',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `sku_name` varchar(255) NOT NULL COMMENT '采购商品名',
  `sku_num` bigint(20) NOT NULL COMMENT '采购数量',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_wms_ware_task_detail_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of wms_ware_task_detail
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
