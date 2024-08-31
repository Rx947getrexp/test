curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:23001/report_user_op_log -d '{
    "user_id": 219122692,
    "device_id": "1770423763070881792",
    "device_type": "ios",
    "page_name": "1",
    "content": "12311231321232-collector",
    "create_time": "2024-01-01 00:00:01",
    "result": "success"
}'


curl -X POST -H 'Content-Type: application/json' -H 'Lang: cn' http://localhost:23001/report_node_ping_result  -d '{
"user_id":219122692, "report_time":"2024-04-03 00:00:00", "items":[{"ip":"2.2.1.3", "code":"success", "cost":"12ms"}]}'

curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:23001/report_user_op_log -d '{
"user_id": 219122692,
"device_id": "1770423763070881792",
"device_type": "ios",
"page_name": "1",
"content": "12311231321232-collector",
"create_time": "2024-01-01 00:00:01",
"result": "success"
}'


curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/internal/describe_user_info -d '{"user_id": 219122692}'