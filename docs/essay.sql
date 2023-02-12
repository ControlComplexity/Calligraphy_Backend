/*
Navicat MySQL Data Transfer

Source Server         : 101.43.39.188
Source Server Version : 80031
Source Host           : 101.43.39.188:3306
Source Database       : scenery

Target Server Type    : MYSQL
Target Server Version : 80031
File Encoding         : 65001

Date: 2023-02-12 08:31:15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for essay
-- ----------------------------
DROP TABLE IF EXISTS `essay`;
CREATE TABLE `essay` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(128) NOT NULL DEFAULT '' COMMENT '文章编号',
  `title` varchar(128) NOT NULL DEFAULT '' COMMENT '文章标题',
  `abstract` varchar(255) DEFAULT NULL,
  `content` varchar(3000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文章内容',
  `url` varchar(128) NOT NULL DEFAULT '' COMMENT 'URL',
  `time` datetime(3) NOT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '文章内容',
  `type` varchar(128) NOT NULL DEFAULT '' COMMENT '文章类型',
  `image` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `home_image` varchar(255) DEFAULT NULL,
  `hits` int(8) unsigned zerofill DEFAULT NULL COMMENT '访问量',
  `like` int(8) unsigned zerofill DEFAULT NULL COMMENT '喜欢次数',
  `display_home_page` tinyint(1) unsigned zerofill NOT NULL COMMENT '是否在首页展示。1为是，0为否',
  `created_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '删除日期',
  PRIMARY KEY (`id`,`display_home_page`),
  KEY `idx_essay_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;
