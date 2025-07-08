/*
 Navicat Premium Dump SQL

 Source Server         : 127.0.01
 Source Server Type    : MySQL
 Source Server Version : 80035 (8.0.35)
 Source Host           : 127.0.01:3306
 Source Schema         : gin-api

 Target Server Type    : MySQL
 Target Server Version : 80035 (8.0.35)
 File Encoding         : 65001

 Date: 03/07/2025 19:15:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `post_id` int DEFAULT NULL COMMENT '岗位id',
  `dept_id` int DEFAULT NULL COMMENT '部门id',
  `username` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '账号',
  `password` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '昵称',
  `icon` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '头像',
  `email` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '手机',
  `note` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '备注信息',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `status` int NOT NULL DEFAULT '1' COMMENT '帐号启用状态：1->启用,2->禁用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='后台管理员表';

-- ----------------------------
-- Records of sys_admin
-- ----------------------------
BEGIN;
INSERT INTO `sys_admin` (`id`, `post_id`, `dept_id`, `username`, `password`, `nickname`, `icon`, `email`, `phone`, `note`, `create_time`, `status`) VALUES (89, 1, 15, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'admin', 'http://127.0.0.1:8000/upload/20250620/645132000.png', '123456789@qq.com', '13754354536', '后端研发', '2023-05-23 22:15:50', 1);
INSERT INTO `sys_admin` (`id`, `post_id`, `dept_id`, `username`, `password`, `nickname`, `icon`, `email`, `phone`, `note`, `create_time`, `status`) VALUES (98, 11, 15, 'lisi', 'e10adc3949ba59abbe56e057f20f883e', '李四', 'http://127.0.0.1:8000/upload/20250629/500333800.png', '123@qq.com', '13826541566', 'ops', '2025-06-20 11:54:02', 1);
INSERT INTO `sys_admin` (`id`, `post_id`, `dept_id`, `username`, `password`, `nickname`, `icon`, `email`, `phone`, `note`, `create_time`, `status`) VALUES (99, 10, 6, 'zhangfan', 'e10adc3949ba59abbe56e057f20f883e', '大白同学', 'http://127.0.0.1:8000/upload/20250629/341737900.jpg', 'zhangfan@lockin.com', '13826541566', '123', '2025-06-29 15:02:48', 1);
INSERT INTO `sys_admin` (`id`, `post_id`, `dept_id`, `username`, `password`, `nickname`, `icon`, `email`, `phone`, `note`, `create_time`, `status`) VALUES (101, 5, 14, 'wangwu', 'e10adc3949ba59abbe56e057f20f883e', '王五', 'http://127.0.0.1:8000/upload/20250629/341737900.jpg', 'zf@123.com', '13826541566', '123', '2025-07-01 14:34:18', 2);
COMMIT;

-- ----------------------------
-- Table structure for sys_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_role`;
CREATE TABLE `sys_admin_role` (
  `admin_id` int NOT NULL COMMENT '管理员id',
  `role_id` int NOT NULL COMMENT '角色id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='管理员和角色关系表';

-- ----------------------------
-- Records of sys_admin_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (99, 12);
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (89, 1);
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (101, 10);
INSERT INTO `sys_admin_role` (`admin_id`, `role_id`) VALUES (98, 13);
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int NOT NULL COMMENT '父id',
  `dept_type` int NOT NULL COMMENT '部门类型（1->公司, 2->中心，3->部门）',
  `dept_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '部门名称',
  `dept_status` int NOT NULL DEFAULT '1' COMMENT '部门状态（1->正常 2->停用）',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `dept_name` (`dept_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_bin ROW_FORMAT=DYNAMIC COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (1, 0, 1, '鹿客科技有限公司', 1, '2023-06-14 17:53:23');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (2, 1, 2, '深圳区研发中心', 1, '2023-06-14 17:53:55');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (3, 2, 3, '架构设计部门', 1, '2023-06-14 17:54:15');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (4, 2, 3, '前端研发部门', 1, '2023-06-14 17:55:17');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (5, 2, 3, '后端研发部门', 1, '2023-06-14 17:55:25');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (6, 2, 3, '系统测试部门', 1, '2023-06-14 17:55:31');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (8, 2, 3, '产品体验部门', 1, '2023-06-14 17:55:46');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (12, 1, 2, '北京研发中心', 1, '2025-06-28 23:42:46');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (13, 1, 2, '重庆研发中心', 1, '2025-06-28 23:43:15');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (14, 12, 3, '运维1部', 1, '2025-06-28 23:43:34');
INSERT INTO `sys_dept` (`id`, `parent_id`, `dept_type`, `dept_name`, `dept_status`, `create_time`) VALUES (15, 13, 3, '运维2部', 1, '2025-06-28 23:44:15');
COMMIT;

-- ----------------------------
-- Table structure for sys_login_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_info`;
CREATE TABLE `sys_login_info` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '用户账号',
  `ip_address` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '登录地点',
  `browser` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '操作系统',
  `login_status` int DEFAULT NULL COMMENT '登录状态（1-成功 2-失败）',
  `message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='登录日志记录';

-- ----------------------------
-- Records of sys_login_info
-- ----------------------------
BEGIN;
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (1, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-26 15:10:07');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (2, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '验证码不正确', '2025-06-27 10:11:15');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (3, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 10:11:22');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (4, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 10:35:36');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (5, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 10:39:36');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (6, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '验证码已过期', '2025-06-27 10:51:17');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (7, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 10:51:22');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (8, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 11:06:40');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (9, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 13:56:34');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (10, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 14:56:28');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (11, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-27 19:19:50');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (12, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 2, '密码不正确', '2025-06-28 14:29:35');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (13, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-28 14:29:40');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (14, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-28 15:57:37');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (15, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-28 15:58:50');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (16, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-28 16:01:59');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (17, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-28 19:06:18');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (18, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-28 23:36:22');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (23, 'lisi', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-29 15:23:15');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (24, 'test', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 2, '密码不正确', '2025-06-29 15:23:55');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (25, 'test', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-29 15:24:00');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (26, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-29 15:24:42');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (27, 'lisi', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-29 15:25:26');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (28, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 2, '密码不正确', '2025-06-29 15:26:54');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (31, 'zhangfan', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-29 16:16:15');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (32, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-30 00:06:13');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (33, 'admin', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-30 00:53:23');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (34, 'lisi', '192.168.3.40', '局域网', 'Chrome/137.0.0.0', 'Windows 10', 1, '登录成功', '2025-06-30 00:53:59');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (35, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-06-30 10:28:30');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (36, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-01 10:30:54');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (37, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-02 11:58:08');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (38, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 11:59:20');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (39, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 14:12:11');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (40, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 14:27:21');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (41, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 16:53:19');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (42, 'zhangfan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '密码不正确', '2025-07-03 17:26:51');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (43, 'zhangfan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '密码不正确', '2025-07-03 17:26:57');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (44, 'zhangfan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '密码不正确', '2025-07-03 17:27:02');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (45, 'zhangfan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '密码不正确', '2025-07-03 17:27:11');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (46, 'zhangsan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 2, '密码不正确', '2025-07-03 17:27:18');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (47, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 17:27:29');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (48, 'zhangfan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 17:28:28');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (49, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 17:29:34');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (50, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 17:47:19');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (51, 'zhangfan', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 17:52:11');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (52, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 17:54:00');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (53, 'lisi', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 18:54:26');
INSERT INTO `sys_login_info` (`id`, `username`, `ip_address`, `login_location`, `browser`, `os`, `login_status`, `message`, `login_time`) VALUES (54, 'admin', '10.7.16.22', '局域网', 'Chrome/137.0.0.0', 'Mac OS X 10_15_7', 1, '登录成功', '2025-07-03 18:55:35');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int DEFAULT NULL COMMENT '父级菜单id',
  `menu_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '菜单名称',
  `icon` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '图标',
  `value` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '接口权限值',
  `menu_type` int DEFAULT NULL COMMENT '菜单类型：1->目录；2->菜单；3->按钮（接口绑定权限）',
  `url` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '菜单url',
  `menu_status` int DEFAULT '2' COMMENT '启用状态；1->禁用；2->启用',
  `sort` int DEFAULT NULL COMMENT '排序',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='菜单表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (4, 0, '基础管理', 'Tools', '', 1, '', 2, 4, '2022-09-04 13:57:39');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (6, 4, '用户信息', 'Avatar', 'base:admin:list', 2, 'system/admin', 2, 1, '2022-09-04 13:59:39');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (7, 4, '角色信息', 'InfoFilled', 'base:role:list', 2, 'system/role', 2, 2, '2022-09-04 14:00:12');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (8, 4, '菜单信息', 'Histogram', 'base:menu:list', 2, 'system/menu', 2, 3, '2022-09-04 14:00:17');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (9, 4, '部门信息', 'Menu', 'base:dept:list', 2, 'system/dept', 2, 4, '2022-09-04 14:01:58');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (10, 4, '岗位信息', 'Promotion', 'base:post:list', 2, 'system/post', 2, 5, '2022-09-04 14:02:06');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (16, 6, '新增用户', '', 'base:admin:add', 3, '', 2, 1, '2022-09-04 18:32:55');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (17, 6, '修改用户', '', 'base:admin:edit', 3, '', 2, 2, '2022-09-04 18:33:29');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (18, 6, '删除用户', '', 'base:admin:delete', 3, '', 2, 3, '2022-09-04 18:33:51');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (21, 7, '新增角色', '', 'base:role:add', 3, '', 2, 1, '2022-09-04 18:44:47');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (22, 7, '修改角色', '', 'base:role:edit', 3, '', 2, 2, '2022-09-04 18:45:10');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (23, 7, '删除角色', '', 'base:role:delete', 3, '', 2, 3, '2022-09-04 18:45:46');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (24, 7, '分配权限', '', 'base:role:assign', 3, '', 2, 4, '2022-09-04 18:46:20');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (26, 8, '新增菜单', '', 'base:menu:add', 3, '', 2, 1, '2022-09-04 18:49:51');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (27, 8, '修改菜单', '', 'base:menu:edit', 3, '', 2, 2, '2022-09-04 18:50:24');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (28, 8, '删除菜单', '', 'base:menu:delete', 3, '', 2, 3, '2022-09-04 18:50:53');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (29, 9, '新增部门', '', 'base:dept:add', 3, '', 2, 1, '2022-09-04 18:52:16');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (30, 9, '修改部门', '', 'base:dept:edit', 3, '', 2, 2, '2022-09-04 18:52:37');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (31, 9, '删除部门', '', 'base:dept:delete', 3, '', 2, 3, '2022-09-04 18:52:50');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (32, 10, '新增岗位', '', 'base:post:add', 3, '', 2, 1, '2022-09-04 18:53:28');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (33, 10, '修改岗位', '', 'base:post:edit', 3, '', 2, 2, '2022-09-04 18:53:48');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (34, 10, '删除岗位', '', 'base:post:delete', 3, '', 2, 3, '2022-09-04 18:54:00');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (44, 0, '日志管理', 'BellFilled', '', 1, '', 2, 5, '2022-09-05 11:06:57');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (45, 44, '操作日志', 'User', 'monitor:operator:list', 2, 'monitor/operator', 2, 1, '2022-09-05 11:10:54');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (46, 44, '登录日志', 'DocumentRemove', 'monitor:loginLog:list', 2, 'monitor/loginlog', 2, 2, '2022-09-05 11:11:31');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (47, 45, '清空操作日志', '', 'monitor:operator:clean', 3, '', 2, 1, '2022-09-05 11:12:36');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (49, 46, '清空登录日志', '', 'monitor:loginLog:clean', 3, '', 2, 1, '2022-09-05 11:16:01');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (60, 6, '重置密码', NULL, 'base:admin:reset', 3, NULL, 2, 6, '2022-12-01 16:33:34');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (62, 46, '删除登录日志', '', 'monitor:loginLog:delete', 3, '', 2, 4, '2022-12-02 15:41:56');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (72, 0, '仪表盘', 'HomeFilled', '', 1, 'dashboard', 2, 1, '2023-05-24 22:11:13');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (73, 45, '删除操作日志', '', 'monitor:operator:delete', 3, '', 2, 3, '2023-06-02 10:09:38');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (78, 80, '主机管理', 'Platform', 'cmdb:ecs:list', 2, 'cmdb/ecs', 2, 1, '2025-06-29 00:30:35');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (80, 0, '资产管理', 'TrendCharts', '', 1, '', 2, 2, '2025-07-03 11:47:07');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (81, 0, '容器管理', 'UploadFilled', '', 1, '', 2, 3, '2025-07-03 11:50:47');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (82, 81, '集群管理', 'Menu', 'cloud:k8s:list', 2, 'k8s/list', 2, 1, '2025-07-03 11:56:44');
INSERT INTO `sys_menu` (`id`, `parent_id`, `menu_name`, `icon`, `value`, `menu_type`, `url`, `menu_status`, `sort`, `create_time`) VALUES (83, 81, '节点管理', 'Help', 'cloud:k8s:node', 2, 'k8s/node', 2, 2, '2025-07-03 12:04:59');
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_log`;
CREATE TABLE `sys_operation_log` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int NOT NULL COMMENT '管理员id',
  `username` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '管理员账号',
  `method` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '请求方式',
  `ip` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'IP',
  `url` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT 'URL',
  `create_time` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=174 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='操作日志记录表';

-- ----------------------------
-- Records of sys_operation_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (1, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 20:40:41');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (2, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 20:51:55');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (3, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 20:58:31');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (4, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 20:58:34');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (5, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 22:35:25');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (6, 89, 'admin', 'post', '192.168.3.40', '/api/post/add', '2025-06-28 22:46:32');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (7, 89, 'admin', 'put', '192.168.3.40', '/api/post/update', '2025-06-28 22:48:09');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (8, 89, 'admin', 'delete', '192.168.3.40', '/api/post/delete', '2025-06-28 22:48:24');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (9, 89, 'admin', 'post', '192.168.3.40', '/api/post/add', '2025-06-28 22:50:29');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (10, 89, 'admin', 'post', '192.168.3.40', '/api/post/add', '2025-06-28 22:52:57');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (11, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 23:04:14');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (12, 89, 'admin', 'put', '192.168.3.40', '/api/post/updateStatus', '2025-06-28 23:08:01');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (13, 89, 'admin', 'put', '192.168.3.40', '/api/post/update', '2025-06-28 23:08:24');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (14, 89, 'admin', 'put', '192.168.3.40', '/api/dept/update', '2025-06-28 23:41:00');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (15, 89, 'admin', 'put', '192.168.3.40', '/api/dept/update', '2025-06-28 23:41:41');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (16, 89, 'admin', 'put', '192.168.3.40', '/api/dept/update', '2025-06-28 23:42:31');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (17, 89, 'admin', 'post', '192.168.3.40', '/api/dept/add', '2025-06-28 23:42:46');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (18, 89, 'admin', 'post', '192.168.3.40', '/api/dept/add', '2025-06-28 23:43:15');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (19, 89, 'admin', 'post', '192.168.3.40', '/api/dept/add', '2025-06-28 23:43:34');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (20, 89, 'admin', 'post', '192.168.3.40', '/api/dept/add', '2025-06-28 23:43:50');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (21, 89, 'admin', 'post', '192.168.3.40', '/api/dept/add', '2025-06-28 23:44:11');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (22, 89, 'admin', 'post', '192.168.3.40', '/api/dept/add', '2025-06-28 23:44:15');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (23, 89, 'admin', 'delete', '192.168.3.40', '/api/dept/delete', '2025-06-28 23:54:32');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (24, 89, 'admin', 'put', '192.168.3.40', '/api/dept/update', '2025-06-28 23:54:40');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (25, 89, 'admin', 'put', '192.168.3.40', '/api/dept/update', '2025-06-28 23:54:45');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (26, 89, 'admin', 'post', '192.168.3.40', '/api/menu/add', '2025-06-29 00:30:35');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (27, 89, 'admin', 'put', '192.168.3.40', '/api/menu/update', '2025-06-29 00:49:46');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (28, 89, 'admin', 'put', '192.168.3.40', '/api/menu/update', '2025-06-29 13:18:50');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (29, 89, 'admin', 'put', '192.168.3.40', '/api/role/updateStatus', '2025-06-29 13:55:46');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (30, 89, 'admin', 'put', '192.168.3.40', '/api/role/updateStatus', '2025-06-29 13:59:53');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (31, 89, 'admin', 'post', '192.168.3.40', '/api/role/add', '2025-06-29 14:01:02');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (32, 89, 'admin', 'post', '192.168.3.40', '/api/role/add', '2025-06-29 14:01:15');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (33, 89, 'admin', 'post', '192.168.3.40', '/api/role/add', '2025-06-29 14:01:20');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (34, 89, 'admin', 'delete', '192.168.3.40', '/api/role/delete', '2025-06-29 14:01:32');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (35, 89, 'admin', 'post', '192.168.3.40', '/api/role/add', '2025-06-29 14:01:49');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (36, 89, 'admin', 'put', '192.168.3.40', '/api/role/update', '2025-06-29 14:02:14');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (37, 89, 'admin', 'post', '192.168.3.40', '/api/role/add', '2025-06-29 14:02:29');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (38, 89, 'admin', 'put', '192.168.3.40', '/api/role/update', '2025-06-29 14:10:23');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (39, 89, 'admin', 'put', '192.168.3.40', '/api/role/assignPermissions', '2025-06-29 14:10:35');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (40, 89, 'admin', 'put', '192.168.3.40', '/api/role/updateStatus', '2025-06-29 14:11:56');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (41, 89, 'admin', 'put', '192.168.3.40', '/api/role/assignPermissions', '2025-06-29 14:12:13');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (42, 89, 'admin', 'put', '192.168.3.40', '/api/role/assignPermissions', '2025-06-29 14:29:21');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (43, 89, 'admin', 'put', '192.168.3.40', '/api/role/updateStatus', '2025-06-29 14:37:06');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (44, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updateStatus', '2025-06-29 14:37:30');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (45, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updateStatus', '2025-06-29 14:37:33');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (46, 89, 'admin', 'put', '192.168.3.40', '/api/admin/update', '2025-06-29 14:37:58');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (47, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePassword', '2025-06-29 14:55:31');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (48, 89, 'admin', 'post', '192.168.3.40', '/api/admin/add', '2025-06-29 15:02:48');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (49, 89, 'admin', 'put', '192.168.3.40', '/api/admin/update', '2025-06-29 15:06:23');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (50, 89, 'admin', 'post', '192.168.3.40', '/api/admin/add', '2025-06-29 15:07:28');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (51, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePassword', '2025-06-29 15:07:41');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (52, 99, 'zhangfan', 'post', '127.0.0.1', '/api/upload', '2025-06-29 15:13:26');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (53, 99, 'zhangfan', 'put', '192.168.3.40', '/api/admin/updatePersonal', '2025-06-29 15:13:35');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (54, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePassword', '2025-06-29 15:15:17');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (55, 89, 'admin', 'put', '192.168.3.40', '/api/admin/update', '2025-06-29 15:17:10');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (56, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePassword', '2025-06-29 15:22:48');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (57, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePassword', '2025-06-29 15:22:56');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (58, 100, 'test', 'post', '127.0.0.1', '/api/upload', '2025-06-29 15:24:12');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (59, 100, 'test', 'put', '192.168.3.40', '/api/admin/updatePersonal', '2025-06-29 15:24:18');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (60, 89, 'admin', 'put', '192.168.3.40', '/api/role/assignPermissions', '2025-06-29 15:25:09');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (61, 98, 'lisi', 'post', '127.0.0.1', '/api/upload', '2025-06-29 15:26:13');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (62, 98, 'lisi', 'put', '192.168.3.40', '/api/admin/updatePersonal', '2025-06-29 15:26:18');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (63, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePersonalPassword', '2025-06-29 15:28:16');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (64, 89, 'admin', 'delete', '192.168.3.40', '/api/sysLoginInfo/batch/delete', '2025-06-29 15:31:44');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (65, 89, 'admin', 'delete', '192.168.3.40', '/api/sysLoginInfo/batch/delete', '2025-06-29 15:31:50');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (66, 89, 'admin', 'delete', '192.168.3.40', '/api/sysLoginInfo/batch/delete', '2025-06-29 15:32:08');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (67, 89, 'admin', 'delete', '192.168.3.40', '/api/sysLoginInfo/delete', '2025-06-29 15:32:14');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (68, 89, 'admin', 'put', '192.168.3.40', '/api/role/assignPermissions', '2025-06-29 16:15:39');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (69, 99, 'zhangfan', 'delete', '192.168.3.40', '/api/admin/delete', '2025-06-29 16:16:31');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (70, 89, 'admin', 'put', '192.168.3.40', '/api/role/assignPermissions', '2025-06-30 00:52:42');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (71, 89, 'admin', 'put', '192.168.3.40', '/api/admin/updatePassword', '2025-06-30 00:52:59');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (72, 89, 'admin', 'put', '192.168.3.40', '/api/admin/update', '2025-06-30 00:53:07');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (73, 89, 'admin', 'put', '10.7.16.22', '/api/admin/update', '2025-06-30 13:16:08');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (74, 89, 'admin', 'put', '10.7.16.22', '/api/admin/update', '2025-06-30 14:05:25');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (75, 89, 'admin', 'put', '10.7.16.22', '/api/admin/update', '2025-06-30 14:05:48');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (76, 89, 'admin', 'put', '10.7.16.22', '/api/admin/update', '2025-06-30 14:13:13');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (77, 89, 'admin', 'put', '10.7.16.22', '/api/role/update', '2025-06-30 17:23:18');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (78, 89, 'admin', 'post', '10.7.16.22', '/api/admin/add', '2025-07-01 14:34:18');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (79, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 14:45:39');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (80, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 14:45:50');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (81, 89, 'admin', 'put', '10.7.16.22', '/api/role/updateStatus', '2025-07-01 15:49:01');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (82, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 15:49:07');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (83, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 16:20:45');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (84, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 16:20:55');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (85, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 16:22:42');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (86, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 16:24:06');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (87, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updateStatus', '2025-07-01 16:24:13');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (88, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updatePassword', '2025-07-02 14:06:36');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (89, 89, 'admin', 'put', '10.7.16.22', '/api/post/updateStatus', '2025-07-02 14:13:30');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (90, 89, 'admin', 'put', '10.7.16.22', '/api/post/updateStatus', '2025-07-02 14:13:33');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (91, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 10:22:59');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (92, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 10:23:57');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (93, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 10:32:20');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (94, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 10:37:08');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (95, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 10:40:08');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (96, 89, 'admin', 'post', '10.7.16.22', '/api/menu/add', '2025-07-03 10:42:22');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (97, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 11:12:17');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (98, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 11:31:47');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (99, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 11:33:03');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (100, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 11:33:21');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (101, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:33:34');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (102, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 11:33:40');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (103, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:33:54');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (104, 89, 'admin', 'delete', '10.7.16.22', '/api/menu/delete', '2025-07-03 11:34:02');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (105, 89, 'admin', 'post', '10.7.16.22', '/api/menu/add', '2025-07-03 11:47:07');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (106, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:47:41');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (107, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:48:46');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (108, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:48:57');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (109, 89, 'admin', 'post', '10.7.16.22', '/api/menu/add', '2025-07-03 11:50:47');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (110, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:50:55');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (111, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 11:51:22');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (112, 89, 'admin', 'post', '10.7.16.22', '/api/menu/add', '2025-07-03 11:56:44');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (113, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 12:00:01');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (114, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 12:03:21');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (115, 89, 'admin', 'post', '10.7.16.22', '/api/menu/add', '2025-07-03 12:04:59');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (116, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 12:58:09');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (117, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 12:58:24');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (118, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 12:58:54');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (119, 89, 'admin', 'put', '10.7.16.22', '/api/role/updateStatus', '2025-07-03 12:59:09');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (120, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 12:59:27');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (121, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:06:15');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (122, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:06:47');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (123, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:10:08');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (124, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:10:18');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (125, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:10:27');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (126, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:12:41');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (127, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:13:39');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (128, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:15:30');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (129, 89, 'admin', 'put', '10.7.16.22', '/api/admin/update', '2025-07-03 13:21:25');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (130, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:24:38');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (131, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:26:24');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (132, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:32:33');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (133, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:36:36');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (134, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:40:00');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (135, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:41:15');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (136, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:41:29');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (137, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:41:48');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (138, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:42:03');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (139, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:46:17');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (140, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:46:28');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (141, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:47:02');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (142, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:47:16');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (143, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:47:31');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (144, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:50:07');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (145, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 13:50:14');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (146, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 14:03:26');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (147, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 14:11:50');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (148, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 14:25:37');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (149, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 14:25:44');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (150, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 14:26:12');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (151, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 14:26:49');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (152, 89, 'admin', 'put', '10.7.16.22', '/api/menu/update', '2025-07-03 14:26:54');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (153, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 14:30:15');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (154, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 14:30:34');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (155, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 14:30:40');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (156, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 15:00:48');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (157, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 16:52:40');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (158, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:23:16');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (159, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:23:40');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (160, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:24:27');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (161, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:26:06');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (162, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updatePassword', '2025-07-03 17:27:46');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (163, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:28:03');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (164, 99, 'zhangfan', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:29:23');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (165, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:39:30');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (166, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:39:44');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (167, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:40:23');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (168, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:40:49');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (169, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 17:47:01');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (170, 89, 'admin', 'post', '10.7.16.22', '/api/role/add', '2025-07-03 18:47:25');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (171, 89, 'admin', 'put', '10.7.16.22', '/api/role/assignPermissions', '2025-07-03 18:53:44');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (172, 89, 'admin', 'put', '10.7.16.22', '/api/admin/update', '2025-07-03 18:54:06');
INSERT INTO `sys_operation_log` (`id`, `admin_id`, `username`, `method`, `ip`, `url`, `create_time`) VALUES (173, 89, 'admin', 'put', '10.7.16.22', '/api/admin/updatePassword', '2025-07-03 18:54:13');
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL COMMENT '岗位名称',
  `post_status` int NOT NULL DEFAULT '1' COMMENT '状态（1->正常 2->停用）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_bin ROW_FORMAT=DYNAMIC COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (1, 'AAA', '总监', 1, '2023-06-14 20:08:22', '主管各个部门');
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (3, 'CCC', '主管', 1, '2023-06-14 20:08:59', '主管部门');
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (4, 'DDD', '组长', 2, '2023-06-14 20:09:30', '部门组长');
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (5, 'EEE', '组员', 2, '2023-06-14 20:09:51', '部门组员');
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (10, 'ops', '运维工程师', 1, '2025-06-28 22:46:33', '运维工程师');
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (11, 'dev', '研发工程师', 1, '2025-06-28 22:50:29', '研发工程师');
INSERT INTO `sys_post` (`id`, `post_code`, `post_name`, `post_status`, `create_time`, `remark`) VALUES (12, 'test', '测试工程师', 1, '2025-06-28 22:52:57', '测试工程师');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_name` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '角色权限字符串',
  `status` int NOT NULL DEFAULT '1' COMMENT '启用状态：1->启用；2->禁用',
  `description` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '描述',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `name` (`role_name`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='后台角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `status`, `description`, `create_time`) VALUES (1, '超级管理员', 'admin', 1, '最大权限', '2023-06-12 20:04:53');
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `status`, `description`, `create_time`) VALUES (10, '研发同学', 'dev', 1, '研发同学', '2025-06-29 14:01:02');
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `status`, `description`, `create_time`) VALUES (11, '测试同学', 'test', 1, '测试同学', '2025-06-29 14:01:49');
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `status`, `description`, `create_time`) VALUES (12, '运维同学', 'ops', 1, '运维同学', '2025-06-29 14:02:29');
INSERT INTO `sys_role` (`id`, `role_name`, `role_key`, `status`, `description`, `create_time`) VALUES (13, '游客', 'test1', 1, 'test1', '2025-07-03 18:47:25');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int DEFAULT NULL COMMENT '角色ID',
  `menu_id` int DEFAULT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='角色和菜单关系表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 44);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 45);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 47);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 73);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 46);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 49);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (11, 62);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (10, 72);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 4);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 6);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 16);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 17);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 18);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 60);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 7);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 21);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 22);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 23);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 24);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 8);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 26);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 27);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 28);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 9);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 29);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 30);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 31);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 10);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 32);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 33);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 34);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 44);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 45);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 47);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 73);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 46);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 49);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 62);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 72);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 80);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 78);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 81);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 82);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (1, 83);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (12, 80);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (12, 78);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (12, 83);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (12, 81);
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES (13, 72);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
