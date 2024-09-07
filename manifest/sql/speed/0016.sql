ALTER TABLE speed_report.t_user_report_day ADD COLUMN `channel` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '';


ALTER TABLE speed_report.t_user_report_day
DROP INDEX uiq_k;

ALTER TABLE speed_report.t_user_report_day
    ADD UNIQUE INDEX uiq_k (`date`, `channel`);