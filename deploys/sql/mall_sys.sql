-- 创建数据库
CREATE DATABASE IF NOT EXISTS `mall_sys`;
USE `mall_sys`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `order` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `permission` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 0, 1, 'Dashboard', 'fa-bar-chart', '/', NULL, NULL, '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (2, 0, 32, '系统管理', 'fa-tasks', '', NULL, NULL, '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (3, 2, 33, '用户管理', 'fa-users', 'auth/users', NULL, NULL, '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (4, 2, 34, '角色管理', 'fa-user', 'auth/roles', NULL, NULL, '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (5, 2, 35, '权限管理', 'fa-ban', 'auth/permissions', NULL, NULL, '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (6, 2, 36, '菜单管理', 'fa-bars', 'auth/menu', NULL, NULL, '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (7, 2, 37, '操作日志', 'fa-history', 'auth/logs', NULL, NULL, '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (8, 2, 38, '小工具', 'fa-gears', '', NULL, '2021-08-17 16:44:42', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (9, 8, 39, 'Scaffold', 'fa-keyboard-o', 'helpers/scaffold', NULL, '2021-08-17 16:44:42', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (10, 8, 40, 'Database terminal', 'fa-database', 'helpers/terminal/database', NULL, '2021-08-17 16:44:42', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (11, 8, 41, 'Laravel artisan', 'fa-terminal', 'helpers/terminal/artisan', NULL, '2021-08-17 16:44:42', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (12, 8, 42, 'Routes', 'fa-list-alt', 'helpers/routes', NULL, '2021-08-17 16:44:42', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (13, 2, 43, '异常报告', 'fa-bug', 'exceptions', NULL, '2021-08-17 16:49:14', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (14, 2, 44, '日志信息', 'fa-database', 'logs', NULL, '2021-08-17 16:52:04', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (16, 2, 45, '定时任务', 'fa-clock-o', 'scheduling', NULL, '2021-08-17 17:23:36', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (17, 0, 16, '产品管理', 'fa-shopping-basket', '', '', '2021-08-18 08:03:32', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (18, 17, 21, 'SPU管理', 'fa-bars', 'product/spu', '', '2021-08-18 08:03:46', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (20, 23, 4, '页面设置', 'fa-bullseye', 'market/setting', '', '2021-08-22 16:25:12', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (21, 17, 17, '商品分类', 'fa-certificate', 'product/category', '', '2021-08-22 17:05:51', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (23, 0, 2, '营销管理', 'fa-coffee', '', '', '2021-08-31 06:51:57', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (26, 23, 10, '配置中心', 'fa-gavel', 'market/config', '', '2021-09-04 06:04:12', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (27, 17, 18, '商品品牌管理', 'fa-bars', 'product/brand', '', '2021-09-29 07:38:23', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (28, 17, 20, '属性管理', 'fa-bars', 'product/attr', '', '2021-09-29 07:39:15', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (30, 17, 19, '属性分组', 'fa-object-ungroup', 'product/attr_group', '', '2021-09-29 07:40:19', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (31, 17, 22, '商品管理', 'fa-columns', 'product/sku', '', '2021-10-02 06:05:01', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (32, 0, 23, '仓储管理', 'fa-cc-jcb', '', '', '2021-10-02 07:55:29', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (33, 32, 24, '仓库管理', 'fa-bars', 'ware/list', '', '2021-10-02 07:55:46', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (34, 32, 25, '商品库存', 'fa-bars', 'ware/sku', '', '2021-10-02 07:56:06', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (35, 32, 26, '采购单', 'fa-bars', 'ware/purchase', '', '2021-10-02 07:56:43', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (36, 32, 27, '采购需求', 'fa-child', 'ware/purchase_detail', '', '2021-10-02 07:57:19', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (37, 32, 28, '库存工作单', 'fa-bars', 'ware/task', '', '2021-10-02 07:57:46', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (39, 23, 3, '公告管理', 'fa-circle-o-notch', 'market/notice', '', '2021-10-08 09:00:09', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (40, 23, 5, '优惠券管理', 'fa-columns', 'market/coupon', '', '2021-10-13 17:23:46', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (41, 23, 6, '会员优惠券', 'fa-hand-paper-o', 'market/coupon_member', '', '2021-10-13 17:24:27', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (42, 23, 7, '会员价格优惠', 'fa-heart', 'market/member_price', '', '2021-10-13 17:25:12', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (43, 23, 8, '满减优惠', 'fa-foursquare', 'market/full_reduction', '', '2021-10-13 17:25:42', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (44, 23, 9, '商品积分', 'fa-genderless', 'market/bounds', '', '2021-10-13 17:26:26', '2021-10-13 17:27:10');
INSERT INTO `admin_menu` VALUES (45, 0, 11, '秒杀管理', 'fa-user-secret', '', '', '2021-11-07 13:22:00', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (46, 45, 12, '活动管理', 'fa-git-square', 'seckill/activity', '', '2021-11-07 13:22:34', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (47, 45, 13, '场次管理', 'fa-compress', 'seckill/session', '', '2021-11-07 13:22:57', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (48, 45, 14, '商品管理', 'fa-indent', 'seckill/sku', '', '2021-11-07 13:23:24', '2021-11-07 13:24:17');
INSERT INTO `admin_menu` VALUES (49, 45, 15, '商品订阅', 'fa-circle-o-notch', 'seckill/notice', '', '2021-11-07 13:23:52', '2021-11-07 13:24:17');
COMMIT;

-- ----------------------------
-- Table structure for admin_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_operation_log`;
CREATE TABLE `admin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `admin_operation_log_user_id_index` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=562 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_operation_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_permissions`;
CREATE TABLE `admin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `admin_permissions_name_unique` (`name`) USING BTREE,
  UNIQUE KEY `admin_permissions_slug_unique` (`slug`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_permissions` VALUES (1, 'All permission', '*', '', '*', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (2, 'Dashboard', 'dashboard', 'GET', '/', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (3, 'Login', 'auth.login', '', '/auth/login\r\n/auth/logout', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (4, 'User setting', 'auth.setting', 'GET,PUT', '/auth/setting', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (5, 'Auth management', 'auth.management', '', '/auth/roles\r\n/auth/permissions\r\n/auth/menu\r\n/auth/logs', NULL, NULL);
INSERT INTO `admin_permissions` VALUES (6, 'Admin helpers', 'ext.helpers', '', '/helpers/*', '2021-08-17 16:44:42', '2021-08-17 16:44:42');
INSERT INTO `admin_permissions` VALUES (7, 'Exceptions reporter', 'ext.reporter', '', '/exceptions*', '2021-08-17 16:49:14', '2021-08-17 16:49:14');
INSERT INTO `admin_permissions` VALUES (8, 'Logs', 'ext.log-viewer', '', '/logs*', '2021-08-17 16:52:04', '2021-08-17 16:52:04');
INSERT INTO `admin_permissions` VALUES (9, 'Scheduling', 'ext.scheduling', '', '/scheduling*', '2021-08-17 17:23:36', '2021-08-17 17:23:36');
COMMIT;

-- ----------------------------
-- Table structure for admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menu`;
CREATE TABLE `admin_role_menu` (
  `role_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_menu` VALUES (1, 2, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permissions`;
CREATE TABLE `admin_role_permissions` (
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_role_permissions_role_id_permission_id_index` (`role_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_permissions
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_permissions` VALUES (1, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_role_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_users`;
CREATE TABLE `admin_role_users` (
  `role_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_role_users_role_id_user_id_index` (`role_id`,`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_role_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_role_users` VALUES (1, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `admin_roles_name_unique` (`name`) USING BTREE,
  UNIQUE KEY `admin_roles_slug_unique` (`slug`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_roles` VALUES (1, 'Administrator', 'administrator', '2021-08-17 16:18:26', '2021-08-17 16:18:26');
COMMIT;

-- ----------------------------
-- Table structure for admin_user_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_user_permissions`;
CREATE TABLE `admin_user_permissions` (
  `user_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  KEY `admin_user_permissions_user_id_permission_id_index` (`user_id`,`permission_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_user_permissions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(190) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `admin_users_username_unique` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_users` VALUES (1, 'admin', '$2y$10$rx24L.R15vHK.Qpx4qDXZeuMlK59kCwu8A9Vqvh/nj9pTYUBBW6mu', 'Administrator', '', 'LLlwquk3bg4XwWftdpDRch5A8UGe8g9ZUY6NdOF5nTHgK4tezuKHLeQ0Up0J', '2021-08-17 16:18:26', '2021-08-23 06:54:42');
COMMIT;

-- ----------------------------
-- Table structure for laravel_exceptions
-- ----------------------------
DROP TABLE IF EXISTS `laravel_exceptions`;
CREATE TABLE `laravel_exceptions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `line` int(11) NOT NULL,
  `trace` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `query` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `body` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `cookies` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `headers` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=266 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of laravel_exceptions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of migrations
-- ----------------------------
BEGIN;
INSERT INTO `migrations` VALUES (1, '2014_10_12_000000_create_users_table', 1);
INSERT INTO `migrations` VALUES (2, '2014_10_12_100000_create_password_resets_table', 1);
INSERT INTO `migrations` VALUES (3, '2016_01_04_173148_create_admin_tables', 1);
INSERT INTO `migrations` VALUES (4, '2017_07_17_040159_create_exceptions_table', 2);
INSERT INTO `migrations` VALUES (5, '2018_06_19_180205_create_stores_table', 3);
INSERT INTO `migrations` VALUES (6, '2018_06_20_151344_create_goods_attrs_table', 3);
INSERT INTO `migrations` VALUES (7, '2018_06_21_101948_create_goods_attr_values_table', 3);
INSERT INTO `migrations` VALUES (8, '2018_06_22_093416_create_media_category_table', 3);
INSERT INTO `migrations` VALUES (9, '2018_06_22_095600_create_media_table', 3);
INSERT INTO `migrations` VALUES (10, '2019_01_08_031443_create_goods_table', 3);
INSERT INTO `migrations` VALUES (11, '2019_01_08_032622_create_goods_medias_table', 3);
INSERT INTO `migrations` VALUES (12, '2019_02_21_070346_create_goods_attr_map', 3);
INSERT INTO `migrations` VALUES (13, '2019_02_22_165313_create_goods_sku_table', 3);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
