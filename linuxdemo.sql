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

 Date: 28/11/2024 14:32:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity_info
-- ----------------------------
DROP TABLE IF EXISTS `activity_info`;
CREATE TABLE `activity_info`  (
  `uuid` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `activity_posters` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '活动海报',
  `activity_title` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '活动标题',
  `activity_number` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '活动编号',
  `activity_type_id` tinyint NULL DEFAULT NULL COMMENT '活动类型id',
  `keywords` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关键词',
  `release_time` datetime NULL DEFAULT NULL COMMENT '发布时间',
  `check_status` tinyint NULL DEFAULT NULL COMMENT '审核状态（0：草稿 1：待审核 2：报名中 3：审核失败 4：待举办 5：进行中 6：已结束）',
  `check_time` datetime NULL DEFAULT NULL COMMENT '审核时间',
  `check_person_contact` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '审核人联系方式',
  `check_remark` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '审核备注',
  `registration_start_time` datetime NULL DEFAULT NULL COMMENT '报名开始时间',
  `activity_start_time` datetime NULL DEFAULT NULL COMMENT '活动开始时间',
  `addr` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '地址',
  `person_limit` int NULL DEFAULT NULL COMMENT '人数限制',
  `registration_fee` decimal(10, 2) NULL DEFAULT NULL COMMENT '报名费用',
  `wx_customer_code` varchar(256) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '客服微信二维码',
  `payment_method` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '收款方式',
  `registration_end_time` datetime NULL DEFAULT NULL COMMENT '报名结束时间',
  `activity_end_time` datetime NULL DEFAULT NULL COMMENT '活动结束时间',
  `check_need` tinyint NULL DEFAULT NULL COMMENT '是否需要审核（1：需要 0：不需要）',
  `enter_activists_id` int NULL DEFAULT NULL COMMENT '组织活动id',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`uuid`) USING BTREE,
  INDEX `activity_number`(`activity_number` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '活动' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for activity_order
-- ----------------------------
DROP TABLE IF EXISTS `activity_order`;
CREATE TABLE `activity_order`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态（1：待审核 2：待参加 3：已检票 4：已取消）',
  `activity_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '活动id',
  `common_user_id` int NULL DEFAULT NULL COMMENT '普通用户id',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `foreign_key_activity_id_activity_order`(`activity_id` ASC) USING BTREE,
  CONSTRAINT `foreign_key_activity_id_activity_order` FOREIGN KEY (`activity_id`) REFERENCES `activity_info` (`activity_number`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '活动订单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for dormitory_info
-- ----------------------------
DROP TABLE IF EXISTS `dormitory_info`;
CREATE TABLE `dormitory_info`  (
  `dormitory_id` int NOT NULL AUTO_INCREMENT,
  `floor` int NULL DEFAULT NULL,
  `dormitory_number` enum('学生宿舍','教师宿舍') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `dormitory_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`dormitory_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
  CONSTRAINT `order_user_id` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for rotation_info
-- ----------------------------
DROP TABLE IF EXISTS `rotation_info`;
CREATE TABLE `rotation_info`  (
  `rotation_id` int NOT NULL AUTO_INCREMENT COMMENT '轮播图编号',
  `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '轮播图路径',
  `rotation_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '轮播图图片名',
  `create_at` datetime NULL DEFAULT NULL COMMENT '轮播图创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '轮播图修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '轮播图软删除',
  PRIMARY KEY (`rotation_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1000002 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for student_is_on
-- ----------------------------
DROP TABLE IF EXISTS `student_is_on`;
CREATE TABLE `student_is_on`  (
  `student_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '学号',
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户编号',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '姓名',
  `is_on` enum('在校','毕业','休学','退学','实习') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '是否在校',
  `school_id` int NOT NULL COMMENT '学校编号',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间\r\n',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`student_id` DESC) USING BTREE,
  INDEX `student_is_on_school_id`(`student_id` ASC) USING BTREE,
  INDEX `user_id_key_forkey`(`user_id` ASC) USING BTREE,
  INDEX `school_id_key_for`(`school_id` ASC) USING BTREE,
  CONSTRAINT `school_id_key_for` FOREIGN KEY (`school_id`) REFERENCES `linuxdemo_network`.`school_info` (`school_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_id_key_forkey` FOREIGN KEY (`user_id`) REFERENCES `user_info` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for teacher_is_on
-- ----------------------------
DROP TABLE IF EXISTS `teacher_is_on`;
CREATE TABLE `teacher_is_on`  (
  `teacher_id` int NOT NULL COMMENT '职工编号',
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户编号',
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '姓名',
  `is_on` enum('是','否') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '是否在职',
  `school_id` int NOT NULL COMMENT '学校编号',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '软删除',
  PRIMARY KEY (`teacher_id`) USING BTREE,
  INDEX `school_id_key_teacher_for`(`school_id` ASC) USING BTREE,
  CONSTRAINT `school_id_key_teacher_for` FOREIGN KEY (`school_id`) REFERENCES `linuxdemo_network`.`school_info` (`school_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户ID\r\n',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户昵称',
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号',
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
-- Triggers structure for table activity_info
-- ----------------------------
DROP TRIGGER IF EXISTS `before_insert_release_time`;
delimiter ;;
CREATE TRIGGER `before_insert_release_time` BEFORE INSERT ON `activity_info` FOR EACH ROW begin if new.release_time>new.registration_start_time then signal sqlstate '45000' set message_text='发布时间不能超过报名时间'; end if; end
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table activity_info
-- ----------------------------
DROP TRIGGER IF EXISTS `before_insert_registration`;
delimiter ;;
CREATE TRIGGER `before_insert_registration` BEFORE INSERT ON `activity_info` FOR EACH ROW begin if new.registration_start_time>new.registration_end_time then signal sqlstate '45000' set message_text='报名开始时间不能超过报名结束时间'; end if; end
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
