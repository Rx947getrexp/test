ALTER TABLE t_goods ADD COLUMN price_btc decimal(10,8) NOT NULL COMMENT "BTC价格(BTC)" AFTER price_usd;
update t_goods set price_btc=0.0001008 WHERE id=1;
update t_goods set price_btc=0.0001848 WHERE id=2;
update t_goods set price_btc=0.0003528 WHERE id=3;
update t_goods set price_btc=0.000672  WHERE id=4;
update t_goods set price_btc=0.0001176 WHERE id=5;
update t_goods set price_btc=0.0002688 WHERE id=6;
update t_goods set price_btc=0.000504  WHERE id=7;
update t_goods set price_btc=0.0009072 WHERE id=8;

INSERT INTO t_payment_channel VALUES (6,'btc','BTC',1,3,30,'/public/upload/payment/c03323f1c40c84cee9b77607384c7c5e.png','1NdtEswb2RA4A4e4h2J9GHBx2NjS14ZpvL',NULL,'{"phone":"\",\"email\":\"\",\"working_hours\":\"\"}','','',80,now(),now(),'bitcoin','BTC',NULL,0.00,0.00)