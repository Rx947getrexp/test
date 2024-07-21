insert into t_payment_channel set channel_id='freekassa-12',channel_name='МИР',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,created_at=now(),updated_at=now(),usd_network='',freekassa_code='12';
insert into t_payment_channel set channel_id='freekassa-36',channel_name='Card RUB',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,created_at=now(),updated_at=now(),usd_network='',freekassa_code='36';
insert into t_payment_channel set channel_id='freekassa-43',channel_name='СБЕР Pay',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,created_at=now(),updated_at=now(),usd_network='',freekassa_code='43';
insert into t_payment_channel set channel_id='freekassa-44',channel_name='СБП (НСПК)',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,created_at=now(),updated_at=now(),usd_network='',freekassa_code='44';
insert into t_payment_channel set channel_id='freekassa-7',channel_name='VISA/MasterCard',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,created_at=now(),updated_at=now(),usd_network='',freekassa_code='7';



-- begin;
--     delete from t_payment_channel where channel_id like 'freekassa-%';
-- commit;

ALTER TABLE `t_payment_channel`ADD COLUMN `currency_type` varchar(64) default NULL COMMENT '支付渠道币种';
ALTER TABLE `t_payment_channel`ADD COLUMN `freekassa_code` varchar(64) default NULL COMMENT 'freekassa支付通道';
ALTER TABLE `t_payment_channel`ADD COLUMN `commission_rate` decimal(10,2) NOT NULL COMMENT '手续费比例';
ALTER TABLE `t_payment_channel`ADD COLUMN `commission` decimal(10,2) NOT NULL default 0.0 COMMENT '手续费';


--  PayChannelPnSafePay    = "pnsafepay" // RUB
-- 	PayChannelUPay         = "usd"
-- 	PayChannelBankCardPay  = "bankcard"
-- 	PayChannelWebMoneyPay  = "webmoney"
-- 	PayChannelFreekassa_7  = "freekassa-7"
-- 	PayChannelFreekassa_12 = "freekassa-12"
-- 	PayChannelFreekassa_36 = "freekassa-36"
-- 	PayChannelFreekassa_43 = "freekassa-43"
-- 	PayChannelFreekassa_44 = "freekassa-44"

update t_payment_channel set currency_type = 'RUB' where channel_id in ('bankcard', 'pnsafepay', 'freekassa-12', 'freekassa-36', 'freekassa-43', 'freekassa-44');
update t_payment_channel set currency_type = 'WMZ' where channel_id in ('webmoney');
update t_payment_channel set currency_type = 'USD' where channel_id in ('usd');
update t_payment_channel set currency_type = 'UAH' where channel_id in ('freekassa-7');


-- PriceRUB         float64 `json:"price_rub" dc:"卢布价格"`
-- PriceWMZ         float64 `json:"price_wmz" dc:"WMZ价格"`
-- PriceUSD         float64 `json:"price_usd" dc:"USD价格"`
-- PriceUAH         float64 `json:"price_uah" dc:"UAH价格"`

ALTER TABLE `t_goods` ADD COLUMN `price_rub` decimal(10,2) NOT NULL COMMENT '卢布价格(RUB)';
ALTER TABLE `t_goods` ADD COLUMN `price_wmz` decimal(10,2) NOT NULL COMMENT 'WMZ价格(WMZ)';
ALTER TABLE `t_goods` ADD COLUMN `price_usd` decimal(10,2) NOT NULL COMMENT 'USD价格(USD)';
ALTER TABLE `t_goods` ADD COLUMN `price_uah` decimal(10,2) NOT NULL COMMENT 'UAH价格(UAH)';

-- select id, period, title, price,usd_pay_price,webmoney_pay_price from t_goods;
--
-- update t_goods set  price_rub = 500, price_wmz = 0.01, price_usd = 0.02, price_uah = 1000 where id = 1;
-- update t_goods set  price_rub = 1250, price_wmz = 19.38, price_usd = 14, price_uah = 2000 where id = 2;
-- update t_goods set  price_rub = 500, price_wmz = 37.84, price_usd = 29, price_uah = 3000 where id = 3;
-- update t_goods set  price_rub = 600, price_wmz = 66.15, price_usd = 51, price_uah = 4000 where id = 4;
-- update t_goods set  price_rub = 601, price_wmz = 0.02, price_usd = 7, price_uah = 1500 where id = 5;
-- update t_goods set  price_rub = 1600, price_wmz = 24.61, price_usd = 19, price_uah = 2500 where id = 6;
-- update t_goods set  price_rub = 3100, price_wmz = 47.69, price_usd = 37, price_uah = 3500 where id = 7;
-- update t_goods set  price_rub = 6200, price_wmz = 0.03, price_usd = 73, price_uah = 4500 where id = 8;
--
-- select id, period, title, price,usd_pay_price,webmoney_pay_price, price_rub, price_wmz, price_usd, price_uah from t_goods;


ALTER TABLE `t_pay_order`ADD COLUMN `commission` decimal(10,2) DEFAULT NULL COMMENT '手续费';