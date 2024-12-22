insert into t_payment_channel set
    channel_id='russpay-bankcard',channel_name='russpay-bankcard',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,
    created_at=now(),updated_at=now(),usd_network='',currency_type='RUB',commission_rate=0;

insert into t_payment_channel set
    channel_id='russpay-sbp',channel_name='russpay-sbp',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,
    created_at=now(),updated_at=now(),usd_network='',currency_type='RUB',commission_rate=0;

insert into t_payment_channel set
    channel_id='russpay-sber',channel_name='russpay-sber',is_active=1,free_trial_days=3,timeout_duration=60,weight=50,
    created_at=now(),updated_at=now(),usd_network='',currency_type='RUB',commission_rate=0;


-- https://www.3hks.xyz/app-api/russpay_callback
