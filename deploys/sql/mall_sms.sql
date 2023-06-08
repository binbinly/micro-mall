-- 创建数据库
CREATE DATABASE IF NOT EXISTS `mall_sms`;
USE `mall_sms`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sms_app_notice
-- ----------------------------
DROP TABLE IF EXISTS `sms_app_notice`;
CREATE TABLE `sms_app_notice` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `content` varchar(1000) NOT NULL COMMENT '内容',
  `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_app_notice
-- ----------------------------
BEGIN;
INSERT INTO `sms_app_notice` VALUES (1, '测试', '阿斯顿发斯蒂芬阿道夫撒旦大发放阿斯蒂芬发生发发是非得失阿斯顿发斯蒂芬', 1, 1, 1633683688, 1633683699);
COMMIT;

-- ----------------------------
-- Table structure for sms_app_setting
-- ----------------------------
DROP TABLE IF EXISTS `sms_app_setting`;
CREATE TABLE `sms_app_setting` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `cat_id` int(10) unsigned NOT NULL COMMENT '产品分类',
  `page` tinyint(3) unsigned NOT NULL COMMENT '页面',
  `type` tinyint(3) unsigned NOT NULL COMMENT '类型',
  `data` varchar(5000) NOT NULL DEFAULT '' COMMENT '数据',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_app_setting
-- ----------------------------
BEGIN;
INSERT INTO `sms_app_setting` VALUES (1, 0, 1, 1, '[\"images\\/20211106\\/f94724e0dc32171058af703d0989b6ef.jpeg\",\"images\\/20211106\\/895b61aadaf08f3b412151a5b937c414.jpeg\",\"images\\/20211106\\/2b4ca1507f039706dd42c5654cbc5e15.jpeg\"]', 50, 1, 1, 1633685383, 1636201862);
INSERT INTO `sms_app_setting` VALUES (2, 0, 1, 2, '[{\"title\":\"\\u65b0\\u54c1\\u53d1\\u5e03\",\"icon\":\"images\\/20211106\\/a63f32c50a1527b417f0ca212f3bc670.png\"},{\"title\":\"\\u5546\\u57ce\\u4f17\\u7b79\",\"icon\":\"images\\/20211106\\/ee6bbb5c159f8d575a4e93bb9b1daade.png\"},{\"title\":\"\\u4ee5\\u65e7\\u6362\\u65b0\",\"icon\":\"images\\/20211106\\/be6db792abda63884f12e70177259eb0.png\"},{\"title\":\"\\u4e00\\u5206\\u6362\\u56e2\",\"icon\":\"images\\/20211106\\/553a99c0da970429d1fd601bcee0697b.png\"},{\"title\":\"\\u8d85\\u503c\\u7279\\u5356\",\"icon\":\"images\\/20211106\\/47d8bcbb619c99f3f15d8e3eb7d7b2c2.png\"},{\"title\":\"\\u5546\\u57ce\\u79d2\\u6740\",\"icon\":\"images\\/20211106\\/fb9aed4fb77ff029d2d9304686f9b4a4.png\"},{\"title\":\"\\u771f\\u5fc3\\u60f3\\u8981\",\"icon\":\"images\\/20211106\\/db8bd586d5b372fb20962f824f91cada.png\"},{\"title\":\"\\u7535\\u89c6\\u7279\\u5356\",\"icon\":\"images\\/20211106\\/5f583636c8ccbf9ee230ca9d2e793b36.png\"},{\"title\":\"\\u5bb6\\u7535\\u70ed\\u5356\",\"icon\":\"images\\/20211106\\/53a6298811691e99ca02165f4abadb79.png\"},{\"title\":\"\\u5176\\u4ed6\\u9009\\u9879\",\"icon\":\"images\\/20211106\\/072babc281abbac7150555c9ddb4771f.png\"}]', 50, 1, 1, 1635250853, 1636201908);
INSERT INTO `sms_app_setting` VALUES (3, 0, 1, 3, '[\"images\\/20211106\\/55985da9c51f1389d7385808cedfb630.jpeg\",\"images\\/20211106\\/7a95292e160c8041b430c6ec6abc07e9.png\",\"images\\/20211106\\/913b6eb5ed088122dee7e43bd47750eb.png\"]', 50, 1, 1, 1635251140, 1636201977);
INSERT INTO `sms_app_setting` VALUES (4, 0, 1, 4, '{\"title\":\"\\u6bcf\\u65e5\\u7cbe\\u9009\",\"cover\":\"images\\/20211106\\/c8ca207f2eae19b78b66c31364dc11bc.jpeg\"}', 50, 1, 1, 1635251150, 1636202003);
INSERT INTO `sms_app_setting` VALUES (5, 0, 1, 5, '{\"router\":\"\\/product\\/list\"}', 50, 1, 1, 1635251170, 1635401996);
INSERT INTO `sms_app_setting` VALUES (6, 0, 2, 4, '{\"title\":\"\\u70ed\\u95e8\\u641c\\u7d22\",\"cover\":\"images\\/20211106\\/5ef2f71aebb6d92794766d4d11959f59.jpeg\"}', 50, 1, 1, 1635251400, 1636202018);
COMMIT;

-- ----------------------------
-- Table structure for sms_config
-- ----------------------------
DROP TABLE IF EXISTS `sms_config`;
CREATE TABLE `sms_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` char(15) NOT NULL COMMENT '配置键',
  `value` varchar(1000) NOT NULL COMMENT '配置值',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_config
-- ----------------------------
BEGIN;
INSERT INTO `sms_config` VALUES (1, 'app_home_cat', '[{\"id\":1220,\"name\":\"\\u5367\\u5ba4\\u5bb6\\u5177\"},{\"id\":1221,\"name\":\"\\u5ba2\\u9910\\u5385\\u5bb6\\u5177\"},{\"id\":1222,\"name\":\"\\u4e66\\u623f\\/\\u513f\\u7ae5\\u623f\"},{\"id\":1225,\"name\":\"\\u5bb6\\u9970\\u9986\"},{\"id\":1224,\"name\":\"\\u706f\\u9970\\u9986\"},{\"id\":1227,\"name\":\"\\u5b9a\\u5236\\u9986\"}]', '', 1, 1, 1633685008, 1636202082);
INSERT INTO `sms_config` VALUES (2, 'app_pay_list', '[{\"id\":1,\"name\":\"\\u4ee5\\u592a\\u5e01-local\",\"address\":\"0xaabE55929c247eC3413B9467be8FdE6AC69e1e12\",\"_remove_\":\"0\"},{\"id\":2,\"name\":\"\\u4ee5\\u592a\\u5e01-test\",\"address\":\"0x0BBCE2CE94E15b8f954a342B2A2A449705694dB7\",\"_remove_\":\"0\"}]', '', 1, 1, 1633685038, 1635826597);
INSERT INTO `sms_config` VALUES (3, 'test', '1', '', 0, 0, 13, 23);
COMMIT;

-- ----------------------------
-- Table structure for sms_coupon
-- ----------------------------
DROP TABLE IF EXISTS `sms_coupon`;
CREATE TABLE `sms_coupon` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(128) NOT NULL COMMENT '优惠券名',
  `cover` varchar(128) NOT NULL DEFAULT '' COMMENT '优惠券封面',
  `type` tinyint(4) NOT NULL COMMENT '优惠卷类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]',
  `num` bigint(20) NOT NULL COMMENT '数量',
  `amount` bigint(20) NOT NULL COMMENT '金额/分',
  `per_limit` bigint(20) NOT NULL DEFAULT '1' COMMENT '每人限领张数',
  `min_point` bigint(20) NOT NULL COMMENT '使用门槛',
  `start_at` bigint(20) NOT NULL COMMENT '开始时间',
  `end_at` bigint(20) NOT NULL COMMENT '结束时间',
  `use_type` tinyint(4) NOT NULL COMMENT '使用类型[0->全场通用；1->指定分类；2->指定商品]',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `publish_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '发行数量',
  `use_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '已使用数量',
  `receive_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '已领取数',
  `enable_start_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '可以领取的开始时间',
  `enable_end_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '可以领取的结束时间',
  `code` varchar(20) NOT NULL COMMENT '优惠码',
  `member_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可以领取的会员等级[0->不限等级，其他-对应等级]',
  `is_release` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否发布上线',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_coupon
-- ----------------------------
BEGIN;
INSERT INTO `sms_coupon` VALUES (1, '满100减5', '', 0, 100, 500, 1, 10000, 1634572800, 1638201600, 0, '', 100, 0, 0, 1634572800, 1638201600, '', 0, 1, 1634632170, 1635702231);
INSERT INTO `sms_coupon` VALUES (2, '满200减15', '', 0, 100, 1500, 1, 20000, 1634572800, 1638201600, 0, '', 100, 0, 0, 1634572800, 1638201600, '', 0, 1, 1634632221, 1635702219);
INSERT INTO `sms_coupon` VALUES (3, '满300减40', '', 0, 100, 4000, 1, 30000, 1634572800, 1638201600, 0, '', 100, 0, 0, 1634572800, 1638201600, '', 0, 1, 1634632300, 1635702205);
INSERT INTO `sms_coupon` VALUES (4, '满500减60', '', 0, 100, 6000, 1, 50000, 1634572800, 1638201600, 0, '', 100, 0, 0, 1634572800, 1638201600, '', 0, 1, 1634632354, 1635702193);
INSERT INTO `sms_coupon` VALUES (5, '满1000减188', '', 0, 100, 18800, 1, 100000, 1634572800, 1638201600, 0, '满1000减188', 100, 0, 0, 1634572800, 1638201600, '', 0, 1, 1634632398, 1635702450);
COMMIT;

-- ----------------------------
-- Table structure for sms_coupon_member
-- ----------------------------
DROP TABLE IF EXISTS `sms_coupon_member`;
CREATE TABLE `sms_coupon_member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `member_id` int(11) unsigned NOT NULL COMMENT '用户id',
  `coupon_id` bigint(20) NOT NULL COMMENT '优惠券id',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '会员昵称',
  `get_type` tinyint(4) NOT NULL COMMENT '获取方式[0->后台赠送；1->主动领取]',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '使用状态[0->未使用；1->已使用；2->已过期]',
  `used_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '使用时间',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单id',
  `order_no` varchar(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_sms_coupon_member_member_id` (`member_id`),
  KEY `idx_sms_coupon_member_coupon_id` (`coupon_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_coupon_member
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_coupon_rel_cat
-- ----------------------------
DROP TABLE IF EXISTS `sms_coupon_rel_cat`;
CREATE TABLE `sms_coupon_rel_cat` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `coupon_id` bigint(20) NOT NULL COMMENT '优惠券id',
  `cat_id` int(10) unsigned NOT NULL COMMENT '产品分类',
  `cat_name` varchar(255) NOT NULL COMMENT '产品分类名',
  PRIMARY KEY (`id`),
  KEY `idx_sms_coupon_rel_cat_coupon_id` (`coupon_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_coupon_rel_cat
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_coupon_rel_spu
-- ----------------------------
DROP TABLE IF EXISTS `sms_coupon_rel_spu`;
CREATE TABLE `sms_coupon_rel_spu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `coupon_id` bigint(20) NOT NULL COMMENT '优惠券id',
  `spu_id` bigint(20) NOT NULL COMMENT 'spu_id',
  `spu_name` varchar(255) NOT NULL COMMENT '产品名',
  PRIMARY KEY (`id`),
  KEY `idx_sms_coupon_rel_spu_coupon_id` (`coupon_id`),
  KEY `idx_sms_coupon_rel_spu_spu_id` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_coupon_rel_spu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_home_adv
-- ----------------------------
DROP TABLE IF EXISTS `sms_home_adv`;
CREATE TABLE `sms_home_adv` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `img` varchar(128) NOT NULL COMMENT '图片',
  `start_at` bigint(20) NOT NULL COMMENT '开始时间',
  `end_at` bigint(20) NOT NULL COMMENT '结束时间',
  `click_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '点击数',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT '链接地址',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `is_release` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否发布上线',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_home_adv
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_home_subject
-- ----------------------------
DROP TABLE IF EXISTS `sms_home_subject`;
CREATE TABLE `sms_home_subject` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `subtitle` varchar(128) NOT NULL COMMENT '副标题',
  `img` varchar(128) NOT NULL COMMENT '图片',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT '链接地址',
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_home_subject
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_home_subject_spu
-- ----------------------------
DROP TABLE IF EXISTS `sms_home_subject_spu`;
CREATE TABLE `sms_home_subject_spu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(128) NOT NULL COMMENT '名称',
  `subject_id` bigint(20) NOT NULL COMMENT '专题id',
  `spu_id` bigint(20) NOT NULL COMMENT 'spu_id',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sms_home_subject_spu_spu_id` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_home_subject_spu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_member_price
-- ----------------------------
DROP TABLE IF EXISTS `sms_member_price`;
CREATE TABLE `sms_member_price` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `level_id` bigint(20) NOT NULL COMMENT '会员等级',
  `level_name` varchar(64) NOT NULL COMMENT '会员等级名',
  `price` bigint(20) NOT NULL COMMENT '会员对应价格',
  `is_super` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否叠加优惠',
  PRIMARY KEY (`id`),
  KEY `idx_sms_member_price_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_member_price
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_seckill_activity
-- ----------------------------
DROP TABLE IF EXISTS `sms_seckill_activity`;
CREATE TABLE `sms_seckill_activity` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(128) NOT NULL COMMENT '标题',
  `cover` varchar(128) NOT NULL DEFAULT '' COMMENT '活动封面',
  `start_at` bigint(20) NOT NULL COMMENT '开始时间',
  `end_at` bigint(20) NOT NULL COMMENT '结束时间',
  `is_release` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否发布上线',
  `create_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_seckill_activity
-- ----------------------------
BEGIN;
INSERT INTO `sms_seckill_activity` VALUES (1, '每日秒杀', '', 1636214400, 1638201600, 1, 1, 0, 1636272174, 1636272174);
COMMIT;

-- ----------------------------
-- Table structure for sms_seckill_session
-- ----------------------------
DROP TABLE IF EXISTS `sms_seckill_session`;
CREATE TABLE `sms_seckill_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(128) NOT NULL COMMENT '场次名',
  `start_at` bigint(20) NOT NULL COMMENT '开始时间',
  `end_at` bigint(20) NOT NULL COMMENT '结束时间',
  `is_release` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否发布上线',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_seckill_session
-- ----------------------------
BEGIN;
INSERT INTO `sms_seckill_session` VALUES (1, '九点', 1636462800, 1636466399, 1, 1636272265, 1636365418);
INSERT INTO `sms_seckill_session` VALUES (2, '十点', 1636466400, 1636469999, 1, 1636274795, 1636365410);
INSERT INTO `sms_seckill_session` VALUES (3, '八点', 1636459200, 1636462799, 1, 1636274823, 1636365402);
COMMIT;

-- ----------------------------
-- Table structure for sms_seckill_sku
-- ----------------------------
DROP TABLE IF EXISTS `sms_seckill_sku`;
CREATE TABLE `sms_seckill_sku` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sku_id` int(10) unsigned NOT NULL COMMENT 'sku_id',
  `activity_id` int(10) unsigned NOT NULL COMMENT '活动id',
  `session_id` int(10) unsigned NOT NULL COMMENT '场次id',
  `price` int(10) unsigned NOT NULL COMMENT '秒杀价格',
  `count` int(10) unsigned NOT NULL COMMENT '秒杀总量',
  `limit` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '每人限购数量',
  `sort` int(10) unsigned NOT NULL DEFAULT '50' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sms_seckill_sku_sku_id` (`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_seckill_sku
-- ----------------------------
BEGIN;
INSERT INTO `sms_seckill_sku` VALUES (2, 1300, 1, 1, 100, 100, 1, 50);
INSERT INTO `sms_seckill_sku` VALUES (3, 2900, 1, 1, 100, 100, 1, 50);
INSERT INTO `sms_seckill_sku` VALUES (4, 4300, 1, 2, 100, 100, 1, 50);
INSERT INTO `sms_seckill_sku` VALUES (5, 7101, 1, 3, 100, 100, 1, 50);
COMMIT;

-- ----------------------------
-- Table structure for sms_seckill_sku_notice
-- ----------------------------
DROP TABLE IF EXISTS `sms_seckill_sku_notice`;
CREATE TABLE `sms_seckill_sku_notice` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `member_id` bigint(20) NOT NULL COMMENT '会员id',
  `session_id` bigint(20) NOT NULL COMMENT '场次id',
  `send_at` bigint(20) NOT NULL COMMENT '发送时间',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通知方式[0-短信，1-邮件]',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_sms_seckill_sku_notice_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_seckill_sku_notice
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_sku_full_reduction
-- ----------------------------
DROP TABLE IF EXISTS `sms_sku_full_reduction`;
CREATE TABLE `sms_sku_full_reduction` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `full_price` bigint(20) NOT NULL COMMENT '满多少',
  `reduce_price` bigint(20) NOT NULL COMMENT '减多少',
  `is_super` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否叠加优惠',
  PRIMARY KEY (`id`),
  KEY `idx_sms_sku_full_reduction_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_sku_full_reduction
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_sku_ladder
-- ----------------------------
DROP TABLE IF EXISTS `sms_sku_ladder`;
CREATE TABLE `sms_sku_ladder` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sku_id` bigint(20) NOT NULL COMMENT 'sku_id',
  `full_count` bigint(20) NOT NULL COMMENT '满几件',
  `discount` bigint(20) NOT NULL COMMENT '打几折',
  `price` bigint(20) NOT NULL COMMENT '折后价',
  `is_super` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否叠加优惠',
  PRIMARY KEY (`id`),
  KEY `idx_sms_sku_ladder_sku_id` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_sku_ladder
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sms_spu_bounds
-- ----------------------------
DROP TABLE IF EXISTS `sms_spu_bounds`;
CREATE TABLE `sms_spu_bounds` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `spu_id` bigint(20) NOT NULL COMMENT 'spu_id',
  `grow_bounds` bigint(20) NOT NULL COMMENT '成长积分',
  `buy_bounds` bigint(20) NOT NULL COMMENT '购物积分',
  `work` tinyint(4) NOT NULL DEFAULT '0' COMMENT '优惠生效情况[0 - 无优惠，成长积分是否赠送',
  PRIMARY KEY (`id`),
  KEY `idx_sms_spu_bounds_spu_id` (`spu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sms_spu_bounds
-- ----------------------------
BEGIN;
INSERT INTO `sms_spu_bounds` VALUES (1, 95404, 0, 0, 0);
INSERT INTO `sms_spu_bounds` VALUES (2, 95405, 0, 0, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
