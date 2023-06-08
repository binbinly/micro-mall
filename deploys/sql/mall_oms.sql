-- 创建数据库
CREATE DATABASE IF NOT EXISTS `mall_oms`;
USE `mall_oms`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for oms_config
-- ----------------------------
DROP TABLE IF EXISTS `oms_config`;
CREATE TABLE `oms_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` char(15) NOT NULL COMMENT '配置键',
  `value` varchar(1000) NOT NULL COMMENT '配置值',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_order
-- ----------------------------
DROP TABLE IF EXISTS `oms_order`;
CREATE TABLE `oms_order` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `order_no` char(30) NOT NULL COMMENT '订单号',
  `coupon_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '优惠券id',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `total_amount` bigint(20) NOT NULL COMMENT '订单总额/分',
  `freight_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '运费/分',
  `promotion_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '促销优惠金额（促销价、满减、阶梯价）',
  `integration_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '积分抵扣金额',
  `coupon_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '优惠券折扣金额',
  `discount_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '后台调整订单使用的折扣金额',
  `amount` bigint(20) NOT NULL COMMENT '优惠后的真实金额',
  `pay_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付金额/分',
  `pay_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '支付方式',
  `source_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单来源[0->PC订单；1->app订单]',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单状态',
  `delivery_company` varchar(64) NOT NULL DEFAULT '' COMMENT '物流公司(配送方式)',
  `delivery_no` varchar(20) NOT NULL DEFAULT '' COMMENT '物流单号',
  `auto_confirm_day` bigint(20) NOT NULL DEFAULT '15' COMMENT '自动确认天数',
  `integration` bigint(20) NOT NULL DEFAULT '0' COMMENT '可以获得的积分',
  `growth` bigint(20) NOT NULL DEFAULT '0' COMMENT '可以获得的成长值',
  `address_name` varchar(64) NOT NULL COMMENT '收货人姓名',
  `address_phone` char(11) NOT NULL COMMENT '收货人手机',
  `address_province` varchar(30) NOT NULL COMMENT '省',
  `address_city` varchar(30) NOT NULL COMMENT '市',
  `address_county` varchar(30) NOT NULL COMMENT '县/区',
  `address_detail` varchar(255) NOT NULL COMMENT '详情',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `is_confirm` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否确认收货',
  `use_integration` bigint(20) NOT NULL DEFAULT '0' COMMENT '下单时使用的积分',
  `trade_no` varchar(100) NOT NULL COMMENT '交易号',
  `trans_hash` varchar(100) NOT NULL COMMENT '交易hash',
  `pay_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付时间',
  `delivery_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '发货时间',
  `receive_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '确认收货时间',
  `comment_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '评价时间',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_oms_order_order_no` (`order_no`),
  KEY `idx_oms_order_member_id` (`member_id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_order_bill
-- ----------------------------
DROP TABLE IF EXISTS `oms_order_bill`;
CREATE TABLE `oms_order_bill` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `type` tinyint(4) NOT NULL COMMENT '类型：0:个人,1:企业',
  `name` varchar(120) NOT NULL COMMENT '名称/公司名称',
  `phone` char(11) NOT NULL COMMENT '手机号',
  `email` varchar(120) NOT NULL COMMENT '邮箱',
  `code` varchar(60) NOT NULL DEFAULT '' COMMENT '税号',
  `address` varchar(120) NOT NULL COMMENT '地址',
  `bank_name` varchar(60) NOT NULL COMMENT '开户行',
  `bank_account` varchar(30) NOT NULL COMMENT '银行账号',
  `is_finish` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否开票',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oms_order_bill_member_id` (`member_id`),
  KEY `idx_oms_order_bill_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order_bill
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_order_item
-- ----------------------------
DROP TABLE IF EXISTS `oms_order_item`;
CREATE TABLE `oms_order_item` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `order_no` char(30) NOT NULL COMMENT '订单号',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `sku_name` varchar(60) NOT NULL COMMENT '商品sku名',
  `sku_img` varchar(128) NOT NULL COMMENT '商品sku图片',
  `sku_price` bigint(20) NOT NULL COMMENT '商品单价/分',
  `sku_attrs` varchar(255) NOT NULL COMMENT '商品销售属性组合（JSON）',
  `sku_num` bigint(20) NOT NULL COMMENT '购买数量',
  `promotion_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '促销优惠金额（促销价、满减、阶梯价）',
  `integration_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '积分抵扣金额',
  `coupon_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '优惠券折扣金额',
  `real_amount` bigint(20) NOT NULL COMMENT '优惠后的真实金额',
  `give_integration` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送积分',
  `give_growth` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送成长值',
  PRIMARY KEY (`id`),
  KEY `idx_oms_order_item_order_id` (`order_id`),
  KEY `idx_oms_order_item_order_no` (`order_no`),
  KEY `idx_oms_order_item_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order_item
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_order_operate
-- ----------------------------
DROP TABLE IF EXISTS `oms_order_operate`;
CREATE TABLE `oms_order_operate` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `order_status` tinyint(4) NOT NULL COMMENT '订单状态',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `operate_name` varchar(60) NOT NULL COMMENT '操作人',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_oms_order_operate_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order_operate
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `oms_order_refund`;
CREATE TABLE `oms_order_refund` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `amount` bigint(20) NOT NULL COMMENT '退款金额',
  `trade_no` varchar(32) NOT NULL DEFAULT '' COMMENT '交易号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单状态',
  `channel` tinyint(4) NOT NULL DEFAULT '0' COMMENT '退款渠道[1-支付宝，2-微信，3-银联，4-汇款]',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '退款内容',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oms_order_refund_order_id` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order_refund
-- ----------------------------
BEGIN;
INSERT INTO `oms_order_refund` VALUES (1, 16, 0, '', 0, 0, 'aaaaaaa', 1635776365, 1635776365);
COMMIT;

-- ----------------------------
-- Table structure for oms_order_return_apply
-- ----------------------------
DROP TABLE IF EXISTS `oms_order_return_apply`;
CREATE TABLE `oms_order_return_apply` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `order_no` char(30) NOT NULL COMMENT '订单号',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `amount` bigint(20) NOT NULL COMMENT '退款金额',
  `return_name` varchar(100) NOT NULL COMMENT '退款人姓名',
  `return_phone` char(11) NOT NULL COMMENT '退款人手机号',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单状态',
  `sku_name` varchar(60) NOT NULL COMMENT '商品sku名',
  `sku_img` varchar(128) NOT NULL COMMENT '商品sku图片',
  `sku_price` bigint(20) NOT NULL COMMENT '商品单价/分',
  `sku_attrs` varchar(255) NOT NULL COMMENT '商品销售属性组合（JSON）',
  `pay_amount` bigint(20) NOT NULL COMMENT '支付金额',
  `sku_num` bigint(20) NOT NULL COMMENT '退货数量',
  `spu_brand` varchar(60) NOT NULL COMMENT '品牌',
  `reason` varchar(255) NOT NULL COMMENT '原因',
  `desc` varchar(255) NOT NULL COMMENT '描述',
  `desc_pics` varchar(2000) NOT NULL DEFAULT '' COMMENT '凭证图片多个逗号连接',
  `handle_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '处理时间',
  `handle_note` varchar(255) NOT NULL DEFAULT '' COMMENT '处理备注',
  `handle_by` varchar(50) NOT NULL DEFAULT '' COMMENT '处理人',
  `receive_name` varchar(64) NOT NULL COMMENT '收货人姓名',
  `receive_phone` char(11) NOT NULL COMMENT '收货人手机',
  `receive_note` varchar(255) NOT NULL DEFAULT '' COMMENT '收货备注',
  `receive_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '收货时间',
  `address_detail` varchar(255) NOT NULL COMMENT '收货地址',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oms_order_return_apply_order_id` (`order_id`),
  KEY `idx_oms_order_return_apply_order_no` (`order_no`),
  KEY `idx_oms_order_return_apply_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order_return_apply
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_order_return_reason
-- ----------------------------
DROP TABLE IF EXISTS `oms_order_return_reason`;
CREATE TABLE `oms_order_return_reason` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL COMMENT '退货原因名',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_order_return_reason
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for oms_payment
-- ----------------------------
DROP TABLE IF EXISTS `oms_payment`;
CREATE TABLE `oms_payment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `order_id` bigint(20) NOT NULL COMMENT '订单id',
  `order_no` char(30) NOT NULL COMMENT '订单号',
  `trade_no` varchar(32) NOT NULL COMMENT '交易号',
  `amount` bigint(20) NOT NULL COMMENT '交易金额/分',
  `subject` varchar(255) NOT NULL COMMENT '交易内容',
  `confirm_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '确认时间',
  `callback_content` varchar(3000) NOT NULL DEFAULT '' COMMENT '回调内容',
  `callback_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '回调时间',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_oms_payment_order_no` (`order_no`),
  KEY `idx_oms_payment_order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of oms_payment
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
