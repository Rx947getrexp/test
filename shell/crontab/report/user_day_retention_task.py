import sys
import logging
# import fcntl
import traceback
from logging.handlers import RotatingFileHandler
from util import get_previous_days
from log import init_logging
from datetime import datetime, timedelta
import pymysql
from common import load_config

TASK_NAME = "user_daily_retention_task"

device_mapping = {
    'Android': ['Android'],
    'iPhone': ['iPhone', 'iOS'],
    'Mac': ['Mac', 'macOS', 'Mac OS'],
    'Windows': ['Windows'],
    'Others': []  # 任何未匹配的都归为 Others
}

os_types = list(device_mapping.keys())

def categorize_os(original_os):
    if 'Mac OS' in original_os:
        if 'iPhone' in original_os:
            return 'iPhone'
        else:
            return 'Mac'
    elif 'Android' in original_os:
        return 'Android'
    elif 'Windows' in original_os:
        return 'Windows'
    else:
        return 'Others'

def get_connection(db_name):
    config = load_config("./config.yaml")
    config = config.get(db_name)
    return pymysql.connect(
        host=config["host"],
        port=config["port"],
        user=config["user"],
        passwd=config["pswd"],
        db=config["db"],
        charset=config["charset"]
    )

def execute_query(db_conn, query, params=None):
    with db_conn.cursor() as cursor:
        cursor.execute(query, params)
        return cursor.fetchall()

def get_registered_users_by_day(date, db_conn):
    query = """SELECT DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') AS stat_day, tu.email, COUNT(*) AS user_count FROM t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') = %s GROUP BY tu.email, stat_day;"""
    return execute_query(db_conn, query, (date,))

def get_device_types_and_counts(date, db_conn, device_mapping):
    query = """SELECT  tud.os AS os, COUNT(*) AS device_count FROM t_user_device tud JOIN t_user tu ON tud.user_id = tu.id WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') = %s GROUP BY tud.os;"""
    results = execute_query(db_conn, query, (date,))
    
    categorized_counts = {os_type: 0 for os_type in device_mapping.keys()}
    
    for row in results:
        original_os = row[0]
        categorized_os = categorize_os(original_os)
        categorized_counts[categorized_os] += row[1]
    
    return [(os_type, count) for os_type, count in categorized_counts.items()]

def get_registered_emails_by_day(date, db_conn):
    query = """SELECT tu.email FROM t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') = %s"""
    results = execute_query(db_conn, query, (date,))
    return [row[0] for row in results]

def get_active_emails_in_next_days(date, days, collector_db_conn):
    date_obj = datetime.strptime(date, '%Y-%m-%d')
    start_date = date_obj + timedelta(days=1)
    # end_date = date_obj + timedelta(days=days)
    end_date = (date_obj + timedelta(days=days + 1)) - timedelta(seconds=1)
    query = """SELECT tut.email FROM t_v2ray_user_traffic tut WHERE tut.date >= %s AND tut.date <= %s"""
    results = execute_query(collector_db_conn, query, (int(start_date.strftime('%Y%m%d')), int(end_date.strftime('%Y%m%d'))))
    return [row[0] for row in results]

def get_registered_device_emails(registered_emails, db_conn):
    if not registered_emails:
        return {}
    query = """SELECT tud.os, tu.email FROM t_user_device tud JOIN t_user tu ON tud.user_id = tu.id WHERE tu.email IN %s;"""
    results = execute_query(db_conn, query, (tuple(registered_emails),))
    return {email: os for os, email in results}

def calculate_retention_of_next_days(date, days, speed_db_conn, collector_db_conn):
    registered_emails = get_registered_emails_by_day(date, speed_db_conn)
    active_emails = set(get_active_emails_in_next_days(date, days, collector_db_conn))
    if not registered_emails:
        return {os_type: set() for os_type in os_types}
    registered_device_emails = get_registered_device_emails(registered_emails, speed_db_conn)
    retained_users_by_os = {os_type: set() for os_type in os_types}
    for email in registered_emails:
        if email in active_emails:
            original_os = registered_device_emails.get(email, 'Others')
            categorized_os = categorize_os(original_os)
            retained_users_by_os[categorized_os].add(email)
    return retained_users_by_os

# def insert_or_update_report(stat_day, os, user_count, new_users, next_day_retention, seven_days_retention, fifteen_days_retention, cursor):
#     stat_day_date = int(stat_day.replace('-', ''))
#     query = """
#     INSERT INTO t_user_report_daily (stat_day, os, user_count, new_users, next_day_retention, seven_days_retention, fifteen_days_retention)
#     VALUES (%s, %s, %s, %s, %s, %s, %s)
#     ON DUPLICATE KEY UPDATE
#     user_count = VALUES(user_count),
#     new_users = VALUES(new_users),
#     next_day_retention = VALUES(next_day_retention),
#     seven_days_retention = VALUES(seven_days_retention),
#     fifteen_days_retention = VALUES(fifteen_days_retention);
#     """
#     cursor.execute(query, (stat_day_date, os, user_count, new_users, next_day_retention, seven_days_retention, fifteen_days_retention))

def insert_or_update_report2(date, device, new, retained, day7_retained, day15_retained, cursor):
    date = int(date.replace('-', ''))
    current_time = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
    query = """
    INSERT INTO t_user_device_retention (date, device, new, retained, day7_retained, day15_retained, created_at)
    VALUES (%s, %s, %s, %s, %s, %s, %s)
    ON DUPLICATE KEY UPDATE
    new = VALUES(new),
    retained = VALUES(retained),
    day7_retained = VALUES(day7_retained),
    day15_retained = VALUES(day15_retained);
    """
    cursor.execute(query, (date, device, new, retained, day7_retained, day15_retained, current_time))

def process_daily_data(start_day):
    current_day = datetime.strptime(start_day, '%Y-%m-%d')
    end_day = datetime.now().strftime('%Y-%m-%d')
    
    # 连接数据库
    speed_db_conn = get_connection('speed-db')
    collector_db_conn = get_connection('collector-speed-db')
    report_db_conn = get_connection('report-speed-db')
    
    while current_day.strftime('%Y-%m-%d') <= end_day:
        day_str = current_day.strftime('%Y-%m-%d')
        
        with speed_db_conn.cursor() as cursor:
            # 获取注册用户的统计信息
            registered_results = get_registered_users_by_day(day_str, speed_db_conn)
            
            # 获取设备类型和计数
            #device_results = get_device_types_and_counts(day_str, speed_db_conn, device_mapping)
            
            # 获取新用户统计信息
            new_users_results = get_registered_users_by_day(day_str, speed_db_conn)
            
            # 获取注册用户的设备类型
            registered_emails = get_registered_emails_by_day(day_str, speed_db_conn)
            registered_device_emails = get_registered_device_emails(registered_emails, speed_db_conn)
            
            # 新用户按设备类型分组
            new_users_by_os = {os_type: 0 for os_type in os_types}
            
            for row in new_users_results:
                _, email, _ = row
                original_os = registered_device_emails.get(email, 'Others')
                categorized_os = categorize_os(original_os)
                new_users_by_os[categorized_os] += 1
            
            # 计算次日留存用户数
            next_day_retention_by_os = calculate_retention_of_next_days(day_str, 1, speed_db_conn, collector_db_conn)
            
            # 计算7日留存用户数
            seven_days_retention_by_os = calculate_retention_of_next_days(day_str, 7, speed_db_conn, collector_db_conn)
            
            # 计算15日留存用户数
            fifteen_days_retention_by_os = calculate_retention_of_next_days(day_str, 15, speed_db_conn, collector_db_conn)
            
            # 处理results并插入到t_user_report_daily中
            if registered_results:
                stat_day = registered_results[0][0]
                #user_count = sum([row[2] for row in registered_results])
            else:
                stat_day = day_str
                #user_count = 0
            
            # 插入或更新数据
            with report_db_conn.cursor() as cursor:
                for os in os_types:
                    total_new_users = new_users_by_os.get(os, 0)
                    next_day_retention_count = len(next_day_retention_by_os.get(os, set()))
                    seven_days_retention_count = len(seven_days_retention_by_os.get(os, set()))
                    fifteen_days_retention_count = len(fifteen_days_retention_by_os.get(os, set()))
                    insert_or_update_report2(stat_day, os, total_new_users, next_day_retention_count, seven_days_retention_count, fifteen_days_retention_count, cursor)
                report_db_conn.commit()
        
        # 移动到下一天
        current_day += timedelta(days=1)
if __name__ == '__main__':
    start_day = get_previous_days(15)  # 获取当前日期的前两天
    process_daily_data(start_day)
    # lock_file = f"/tmp/{TASK_NAME}.lock"
    # fp = open(lock_file, "w")
    # try:
    #     fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    # except IOError:
    #     logging.error(f"已经有一个 {TASK_NAME} 进程在运行，本进程将退出")
    #     sys.exit(1)

    # init_logging(f"/shell/report/log/{TASK_NAME}.log")
    # logging.info(f"\n\n\n start {TASK_NAME}")
    # try:
    #     process_daily_data(start_day)
    # except Exception as e:
    #     logging.error(f"捕获到异常：{type(e).__name__}")
    #     logging.error(f"异常信息：{str(e)}")
    #     logging.error(traceback.format_exc())
    # finally:
    #     fp.close()

    # logging.info(f"end {TASK_NAME}")