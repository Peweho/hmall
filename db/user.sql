/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.150.101
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 192.168.150.101:3306
 Source Schema         : hmall

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 31/07/2023 16:26:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码，加密存储',
  `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '注册手机号',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL,
  `status` int NULL DEFAULT 1 COMMENT '使用状态（1正常 2冻结）',
  `balance` int NULL DEFAULT NULL COMMENT '账户余额',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'Jack', '$2a$10$6ptTq3V9XfaJmFYwYT2W9ud377BUkEWk.whf.iQ.0sX5F.L497rAC', '13900112224', '2017-08-19 20:50:21', '2017-08-19 20:50:21', 1, 838500);
INSERT INTO `user` VALUES (2, 'Rose', '$2a$10$6ptTq3V9XfaJmFYwYT2W9ud377BUkEWk.whf.iQ.0sX5F.L497rAC', '13900112223', '2017-08-19 21:00:23', '2017-08-19 21:00:23', 1, 1000000);
INSERT INTO `user` VALUES (3, 'Hope', '$2a$10$6ptTq3V9XfaJmFYwYT2W9ud377BUkEWk.whf.iQ.0sX5F.L497rAC', '13900112222', '2017-08-19 22:37:44', '2017-08-19 22:37:44', 1, 1000000);
INSERT INTO `user` VALUES (4, 'Thomas', '$2a$10$6ptTq3V9XfaJmFYwYT2W9ud377BUkEWk.whf.iQ.0sX5F.L497rAC', '17701265258', '2017-08-19 23:44:45', '2017-08-19 23:44:45', 1, 1000000);

SET FOREIGN_KEY_CHECKS = 1;
