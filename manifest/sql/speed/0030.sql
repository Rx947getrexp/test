ALTER TABLE `speed_report`.`t_user_op_log`
    ADD COLUMN `app_name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'app_name';

ALTER TABLE `speed_report`.`t_user_ad_log`
    ADD COLUMN `app_name` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'app_name';


insert into speed.t_payment_channel set
    channel_id='russ-new-pay-card',channel_name='russ-new-pay-card',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,
    created_at=now(),updated_at=now(),usd_network='',currency_type='RUB',commission_rate=0;

insert into t_payment_channel set
    channel_id='russ-new-pay-sbp',channel_name='russ-new-pay-sbp',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,
    created_at=now(),updated_at=now(),usd_network='',currency_type='RUB',commission_rate=0;

