ALTER TABLE `t_user_op_log` ADD COLUMN version VARCHAR(64) AFTER content;
ALTER TABLE admin_user ADD COLUMN `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '可查看渠道范围,为空则可查看所有范围'
