


insert into t_country set name = 'China-HK',name_cn='中国香港',created_at=now(),updated_at=now();
insert into t_country set name = 'Russia',name_cn='俄罗斯',created_at=now(),updated_at=now();

# 2
curl -X GET -H "Lang: cn" http://localhost:13001/country/list

curl -X GET -H "Lang: cn" http://localhost:13001/country_list


# 3
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/serving_country_add -d '{"name":"China-HK", "name_display":"hk", "logo_link":"http://123", "ping_url":"http://123", "is_recommend":1, "weight":100}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/serving_country_add -d '{"name":"Russia", "name_display":"russia", "logo_link":"http://123", "ping_url":"http://123", "is_recommend":0, "weight":50}'

# 4
curl -X GET -H "Lang: cn" http://localhost:13001/serving_country_list

# 5
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/serving_country_edit -d '{"name":"Russia", "name_display":"russia1", "logo_link":"http://1234", "ping_url":"http://1234", "is_recommend":0, "weight":50, "status":2}'

# 6
curl -X GET -H "Lang: cn" http://localhost:13001/machine_list


# 7
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/machine_add -d '{
    "country_name": "Russia",
    "ip": "1.1.1.1",
    "server": "www.xxx",
    "port": 443,
    "min_port": 13001,
    "max_port": 13005,
    "weight": 50,
    "comment": "测试测试"
}'

# 8
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/machine/edit -d '{
    "id": 100004,
    "ip": "1.1.1.2",
    "server": "1.www.xxx",
    "port":    1443,
    "min_port": 23001,
    "max_port": 23005,
    "weight": 60,
    "comment": "测试测试1",
    "status": 2
}'

# 9
curl -X GET -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/get_user_op_log_list?device_type=11@qq.com&result=success  -d '{
    "email": "11@qq.com",
    "device_type": "1.1.1.2",
    "page_name": "1.www.xxx",
    "result": "1.www.xxx",
    "order_by": "create_time",
    "order_type": "desc",
    "page":    1443,
    "size": 2
}'


curl -X GET -H "Content-Type: application/json" -H "Lang: cn" 'http://localhost:13001/get_user_op_log_list?device_type=11@qq.com&result=success&order_by=create_time&page=2&order_type=desc'

curl -X GET -H "Content-Type: application/json" -H "Lang: cn" 'http://localhost:13001/get_user_op_log_list?order_by=created_at&page=0&order_type=desc'



/////////////
# 1
curl -X GET -H "Lang: cn" http://localhost:13002/get_serving_country_list?user_id=219122623

# 2
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/report_user_op_log -d '{
    "user_id": 219122692,
    "device_id": "1770423763070881792",
    "device_type": "ios",
    "page_name": "1",
    "content": "xx",
    "create_time": "2024-01-01 00:00:01",
    "result": "success"
}'


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://www.yyy360.xyz/app-api/report_user_op_log -d '{
    "user_id": 219122692,
    "device_id": "1770423763070881792",
    "device_type": "ios",
    "page_name": "1",
    "content": "xx",
    "create_time": "2024-01-01 00:00:01",
    "result": "success"
}'



curl -X GET -H "Lang: cn" http://localhost:13002/get_server_config?user_id=219122623

curl -X GET -H "Lang: cn" http://localhost:13002/get_server_config_without_rules?user_id=219122623
curl -X GET -H "Lang: cn" http://localhost:13002/get_rules?user_id=219122623

# 3
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/set_default_country -d '{
    "user_id": 219122623,
    "country_name": "China-HK"
}'

# 4
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/connect_server -d '{
"user_id": 219122692,
"country_name": "China-HK"
}'


### order

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/create_order -d '{
    "user_id": 219122692,
    "product_no": "vip-month",
    "currency": "rub",
    "order_amount": 30
}'


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/create_order -d '{
"channel_id": "freekassa-44",
"goods_id": 1
}'


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/pay_notify -d '{"order_no": "20240505184228571983"}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://www.hongkongs.xyz/app-api/send_email -d '{"email": "VPNHERO@outlook.com"}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://www.hongkongs.xyz/app-api/forget_passwd -d '{"account": "VPNHERO@outlook.com","verify_code":"828247","passwd":"123456","enter_passwd":"123456"}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/send_email -d '{"email": "VPNHERO@outlook.com"}'


curl -X GET -H "Content-Type: application/json" -H "Lang: cn" http://localhost:23001/get_user_op_log_list
curl -X GET -H "Content-Type: application/json" -H "Lang: cn" http://localhost:23001/list_node_for_report

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:23001/report_user_op_log -d '{
    "user_id": 219122692,
    "device_id": "1770423763070881792",
    "device_type": "ios",
    "page_name": "1",
    "content": "xx",
    "create_time": "2024-01-01 00:00:01",
    "result": "success"
}'

curl -X POST -H "Content-Type: application/json" http://api.pnsafepay.com/gateway.aspx -d '{
    "currency": "RUB",
    "mer_no": "1082775",
    "method": "trade.create",
    "order_amount": "500",
    "order_no": "20240429231417116748",
    "payemail": "2233@gmail.com",
    "payname": "hsfly",
    "payphone": "18818811881",
    "paytypecode": "29001",
    "returnurl": "http://www.wuwuwu360.xyz/app-api/pay_notify",
    "sign": "d2ce6618b4fbd94d3c99a9f1a057a0a5"
}'


{"currency":"RUB","mer_no":"1082775","method":"trade.create","order_amount":"500","order_no":"20240429231417116748","payemail":"2233@gmail.com","payname":"hsfly","payphone":"18818811881","paytypecode":"29001","returnurl":"http://www.wuwuwu360.xyz/app-api/pay_notify","sign":"d2ce6618b4fbd94d3c99a9f1a057a0a5"}

curl -X GET -H "Lang: cn" https://www.baodu.xyz/app-api/get_rules?user_id=219122623

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://www.baodu.xyz/app-api/create_order -d '{
    "user_id": 219122692,
    "product_no": "pro-vip-month",
    "currency": "RUB",
    "order_amount": 500
}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://www.baodu.xyz/app-api/pay_notify -d '{"order_no": "20240505121440989398"}'

curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://thertee.xyz/app-api/dns_list
curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://thertee.xyz/app-api/notice_list

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://thertee.xyz/app-api/get_official_docs


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://beiyo.xyz/app-api/get_official_docs


# 支付相关
## 管理后台
### 支付渠道列表
curl -X POST -H "Lang: cn" http://localhost:13001/payment_channel/list
curl -XPOST  -H "Lang: cn" https://keeperpro.xyz/app-api/app_version

### 修改支付渠道配置接口
curl -X POST -H "Lang: cn" -H "Content-Type: application/json" http://localhost:13001/payment_channel/edit -d '{"ChannelId":"usd","payment_qr_code":"qr-123","customer_service_info":{"phone":"18118811881","working_hours":"10:00~20:00"}}'


curl -X POST -H "Lang: cn" -H "Content-Type: application/json" http://localhost:13001/payment_channel/edit -d '{
    "ChannelId": "usd",
    "PaymentQRCode": "qr-123",
    "CustomerServiceInfo": {
        "Phone": "18118811881",
        "WorkingHours": "10:00~20:00"
    }
}'

curl -X POST -H "Lang: cn" -H "Content-Type: application/json"  http://localhost:13001/order/pay_order_list -d '{"order_no":"20240506024451529790","page":1,"size":10,"email":"zzz@qq.com"}'


curl -X POST -H "Lang: cn" -H "Content-Type: application/json"  http://localhost:13001/confirm_order -d '{"order_no":"100725013401746"}'



curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/list

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/add -d '{
    "type": "help",
    "name": "help-question",
    "desc": "官方文档",
    "content": "12344\n123123\n12312312\n123123sfasjdlakjdlakj\n"
}'


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/edit -d '{
    "id": 1,
    "type": "help2",
    "name": "help-question2",
    "desc": "官方文档2",
    "content": "12344\n123123\n12312312\n123123sfasjdlakjdlakj222222222\n"
}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/get_official_docs


select * from t_v2ray_user_traffic where email in ('alt.j9-co173v7i@yopmail.com','alt.j9-co173v7i@yopmail.com','kakyoin2929@gmail.com','cry69gry@mail.ru','vlad.ku4erov2015@gmail.com','sergey-gamzik11@mail.ru','slavamihailov222@gmail.com','Kapshanova14@gmail.com')

grep 'alt.j9-co173v7i@yopmail.com\|kakyoin2929@gmail.com\|cry69gry@mail.ru\|vlad.ku4erov2015@gmail.com\|sergey-gamzik11@mail.ru\|slavamihailov222@gmail.com\|Kapshanova14@gmail.com' config.json


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/device_list

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/kick_device -d '{"client_id":"123234"}'


1）查询设备列表的接口
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://www.yyy360.xyz/app-api/device_list
返回示例：
{
    "code": 200,
    "message": "成功",
    "data": {
        "items": [
            {
                "client_id": "d3c2a8be175ef59375de3001e9cb24703852bc4f3a2f4cf269110d2f96bd4ee3",
                "os": "Android 12",
                "created_at": "2025-02-16 13:48:09",
                "updated_at": "2025-02-16 13:48:09"
            }
        ]
    }
}

2）剔除设备的接口：
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://www.yyy360.xyz/app-api/kick_device -d '{"client_id":"d3c2a8be175ef59375de3001e9cb24703852bc4f3a2f4cf269110d2f96bd4ee3"}'

