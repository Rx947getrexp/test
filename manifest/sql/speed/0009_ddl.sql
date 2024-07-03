ALTER TABLE `t_goods` ADD COLUMN `webmoney_pay_price` decimal(10,2) NOT NULL COMMENT 'webmoney价格(wmz)';

insert into t_payment_channel set channel_id='webmoney',channel_name='webmoney',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,created_at=now(),updated_at=now(),usd_network='';

/*
    mkdir config

webmoneyconfig:
  wmid: 283361774557
  purse: Z113494876653
  rand_code: pingguoqm23
*/


