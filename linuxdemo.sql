/*
 Navicat Premium Dump SQL

 Source Server         : 我的数据库
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : linuxdemo

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 28/10/2024 22:51:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dormitory_info
-- ----------------------------
DROP TABLE IF EXISTS `dormitory_info`;
CREATE TABLE `dormitory_info`  (
  `dormitory_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `floor` int NULL DEFAULT NULL,
  `dormitory_number` int NULL DEFAULT NULL,
  PRIMARY KEY (`dormitory_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dormitory_info
-- ----------------------------

-- ----------------------------
-- Table structure for employee_is_on
-- ----------------------------
DROP TABLE IF EXISTS `employee_is_on`;
CREATE TABLE `employee_is_on`  (
  `employee_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '职工编号',
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户编号',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `is_on` enum('是','否') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否在职',
  `school_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '学校编号',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`employee_id`) USING BTREE,
  INDEX `employee_user_id`(`user_id` ASC) USING BTREE,
  INDEX `employee_school_id`(`school_id` ASC) USING BTREE,
  CONSTRAINT `school_id` FOREIGN KEY (`school_id`) REFERENCES `school_info` (`school_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee_is_on
-- ----------------------------

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info`  (
  `order_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `dormitory_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` enum('100','200') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '100为照明，200为空调',
  `is_pay` int NOT NULL COMMENT '0为未支付，1为已支付',
  `price` double NOT NULL,
  `create_at` datetime NULL DEFAULT NULL,
  `update_at` datetime NULL DEFAULT NULL,
  `delete_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `order_dormitory_id_fk`(`dormitory_id` ASC) USING BTREE,
  INDEX `order_user_user_id_fk`(`user_id` ASC) USING BTREE,
  CONSTRAINT `order_dormitory_id` FOREIGN KEY (`dormitory_id`) REFERENCES `dormitory_info` (`dormitory_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `order_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_info
-- ----------------------------

-- ----------------------------
-- Table structure for rotation_info
-- ----------------------------
DROP TABLE IF EXISTS `rotation_info`;
CREATE TABLE `rotation_info`  (
  `rotation_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '轮播图编号',
  `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '轮播图路径',
  `rotation_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '轮播图图片名',
  `create_at` datetime NULL DEFAULT NULL COMMENT '轮播图创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '轮播图修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '轮播图软删除',
  PRIMARY KEY (`rotation_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rotation_info
-- ----------------------------

-- ----------------------------
-- Table structure for school_info
-- ----------------------------
DROP TABLE IF EXISTS `school_info`;
CREATE TABLE `school_info`  (
  `school_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学校编号',
  `school_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学校名称',
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学校省份',
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学校城市',
  `county` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学号市区',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学校详细地址',
  PRIMARY KEY (`school_id`) USING BTREE,
  UNIQUE INDEX `schoolName`(`school_name` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of school_info
-- ----------------------------
INSERT INTO `school_info` VALUES ('10000', '牛马大学', '', '', '', '');

-- ----------------------------
-- Table structure for student_is_on
-- ----------------------------
DROP TABLE IF EXISTS `student_is_on`;
CREATE TABLE `student_is_on`  (
  `student_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学号',
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户编号',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '姓名',
  `is_on` enum('是','否') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '是否在校',
  `school_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学校编号',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间\r\n',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`student_id` DESC) USING BTREE,
  INDEX `student_is_on_user_id`(`user_id` ASC) USING BTREE,
  INDEX `student_is_on_school_id`(`student_id` ASC) USING BTREE,
  CONSTRAINT `student_is_on_school_id` FOREIGN KEY (`student_id`) REFERENCES `school_info` (`school_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `student_is_on_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of student_is_on
-- ----------------------------

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户ID\r\n',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户电话',
  `wechat` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户微信',
  `user_salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '加密盐',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `user_pk_2`(`phone` ASC) USING BTREE,
  UNIQUE INDEX `user_pk_3`(`wechat` ASC) USING BTREE,
  UNIQUE INDEX `user_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户名' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_info
-- ----------------------------
INSERT INTO `user_info` VALUES ('10000000000001', '992098059', '913bef72708cbc2b4573360f74c8ebce', '13132313112', '123123111', 'Vca55U85up', '2024-10-28 22:48:15', '2024-10-28 22:48:15', NULL);
INSERT INTO `user_info` VALUES ('12313323', '312312', '', '1234', '123123', NULL, '2024-09-24 07:52:06', '2024-09-24 07:52:06', '2024-10-27 15:24:16');
INSERT INTO `user_info` VALUES ('1231332333', '3123123', '', '1234333', '123123113', NULL, '2024-09-24 16:00:23', '2024-09-24 16:00:23', NULL);
INSERT INTO `user_info` VALUES ('12313323331', '13123123', '', '12343331', '1231231131', NULL, '2024-09-24 21:17:03', '2024-09-24 21:17:03', NULL);
INSERT INTO `user_info` VALUES ('992098059', '你好吗223211111', '', '13431855812', '7465165', NULL, '2024-09-25 15:55:16', '2024-09-25 15:56:12', '2024-09-25 16:10:31');

SET FOREIGN_KEY_CHECKS = 1;
