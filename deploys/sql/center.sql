-- 创建数据库
CREATE DATABASE IF NOT EXISTS `center`;
USE `center`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `phone` bigint(20) NOT NULL COMMENT '手机号',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint(4) NOT NULL DEFAULT '1' COMMENT '性别',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
  `sign` varchar(255) NOT NULL DEFAULT '' COMMENT '签名',
  `area` varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
  `created_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `updated_at` int(11) unsigned NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_username` (`username`),
  UNIQUE KEY `idx_user_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
