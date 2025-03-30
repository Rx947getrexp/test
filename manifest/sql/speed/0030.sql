ALTER TABLE `speed_report`.`t_user_op_log`
    ADD COLUMN `app_name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'app_name';

ALTER TABLE `speed_report`.`t_user_ad_log`
    ADD COLUMN `app_name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'app_name';

