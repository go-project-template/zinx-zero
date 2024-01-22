/*
 Navicat Premium Data Transfer

 Source Server         : docker
 Source Server Type    : MySQL
 Source Server Version : 80200 (8.2.0)
 Source Host           : localhost:33069
 Source Schema         : gamex

 Target Server Type    : MySQL
 Target Server Version : 80200 (8.2.0)
 File Encoding         : 65001

 Date: 22/01/2024 14:22:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for redis_cache_key
-- ----------------------------
CREATE TABLE `redis_cache_key`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `prefix1` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '一级模块/包名',
  `prefix2` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '二级模块/文件名',
  `prefix3` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '额外字段',
  `prefix4` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `prefix5` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `prefix6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `arg_int64` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `arg_string` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `ret_struct` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '返回的数据结构(需支持json.Unmarshal)',
  `expiry` int NOT NULL DEFAULT 604800 COMMENT '有效期（秒）默认7天',
  `not_fount_expiry` int NOT NULL DEFAULT 60 COMMENT '空数据占位有效期（秒）默认1分钟',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_cache_key`(`prefix1` ASC, `prefix2` ASC, `prefix3` ASC, `prefix4` ASC, `prefix5` ASC, `prefix6` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of redis_cache_key
-- ----------------------------
INSERT INTO `redis_cache_key` VALUES (1, 'GenRoleId', 'UserIdPool', '', '', '', '', '', '', 'int64', 604800, 60, '2024-01-22 14:17:57', '2024-01-22 14:21:28');

SET FOREIGN_KEY_CHECKS = 1;
