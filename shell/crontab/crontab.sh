# 主节点crontab配置
* * * * * /usr/bin/python3 /shell/monitor_api_redis.py &
* * * * * /usr/bin/python3 /shell/database_backup_speed.py &
#12 0 * * * /usr/bin/python3 /shell/report/report_task.py
* * * * * /usr/bin/python3 /shell/report/clean_report_data.py &
* * * * * /usr/bin/python3 /shell/app_port.py &
* * * * * /usr/bin/python3 /shell/monitor_ng.py &
