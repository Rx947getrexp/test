import sys
import logging
import fcntl
import traceback
from logging.handlers import RotatingFileHandler
from util import get_previous_months
from log import init_logging
from datetime import datetime, timedelta
import pymysql
from common import load_config

TASK_NAME = "user_monthly_retention_task"
start_month = get_previous_months(2) #获取当前月份的前几个月

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
    config = load_config("/shell/report/config.yaml")
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

def get_registered_users_by_month(month, db_conn):
    query = """SELECT DATE_FORMAT(tu.created_at, '%%Y-%%m') AS stat_month, tu.email, COUNT(*) AS user_count FROM t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s GROUP BY tu.email, stat_month;"""
    return execute_query(db_conn, query, (month,))

def get_device_types_and_counts(month, db_conn, device_mapping):
    query = """SELECT  tud.os AS os, COUNT(*) AS device_count FROM t_user_device tud JOIN t_user tu ON tud.user_id = tu.id WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s GROUP BY tud.os;"""
    results = execute_query(db_conn, query, (month,))
    
    categorized_counts = {os_type: 0 for os_type in device_mapping.keys()}
    
    for row in results:
        original_os = row[0]
        categorized_os = categorize_os(original_os)
        categorized_counts[categorized_os] += row[1]
    
    return [(os_type, count) for os_type, count in categorized_counts.items()]

def get_new_users_in_month(month, db_conn):
    query = """SELECT tud.os, tu.email, COUNT(*) AS new_users FROM t_user tu JOIN t_user_device tud ON tu.id = tud.user_id WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s GROUP BY tud.os, tu.email;"""
    return execute_query(db_conn, query, (month,))

def get_registered_emails_by_month(month, db_conn):
    query = """SELECT tu.email FROM t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s"""
    results = execute_query(db_conn, query, (month,))
    return [row[0] for row in results]

def get_active_emails_in_next_month(month, collector_db_conn):
    month_date = datetime.strptime(month, '%Y-%m')
    next_month_date = month_date + timedelta(days=32)
    next_month = next_month_date.strftime('%Y-%m')
    
    next_month_start = next_month_date.replace(day=1)
    next_month_end = (next_month_start + timedelta(days=32)).replace(day=1) - timedelta(days=1)
    
    query = """SELECT tut.email FROM t_v2ray_user_traffic tut WHERE tut.date >= %s AND tut.date <= %s"""
    results = execute_query(collector_db_conn, query, (int(next_month_start.strftime('%Y%m%d')), int(next_month_end.strftime('%Y%m%d'))))
    return [row[0] for row in results]

def calculate_retention_of_next_month(month, speed_db_conn, collector_db_conn):
    registered_emails = get_registered_emails_by_month(month, speed_db_conn)
    active_emails = set(get_active_emails_in_next_month(month, collector_db_conn))  # 转换成集合
    
    if not registered_emails:
        return {os_type: set() for os_type in os_types}  # 如果没有注册用户，直接返回空结果
    
    registered_device_emails = {}
    query = """SELECT tud.os, tu.email FROM t_user_device tud JOIN t_user tu ON tud.user_id = tu.id WHERE tu.email IN %s;"""
    results = execute_query(speed_db_conn, query, (tuple(registered_emails),))
    registered_device_emails = {email: os for os, email in results}

    retained_users_by_os = {os_type: set() for os_type in os_types}
    
    for email in registered_emails:
        if email in active_emails:
            original_os = registered_device_emails.get(email, 'Others')
            categorized_os = categorize_os(original_os)
            retained_users_by_os[categorized_os].add(email)
    
    return retained_users_by_os

def clear_t_user_report_monthly(cursor):
    query = """TRUNCATE TABLE t_user_report_monthly;"""
    cursor.execute(query)

def insert_into_report(stat_month, os, user_count, new_users, retained_users, cursor):
    stat_month_date = int(stat_month.replace('-', ''))
    query = """INSERT INTO t_user_report_monthly (stat_month, os, user_count, new_users, retained_users) VALUES (%s, %s, %s, %s, %s);"""
    cursor.execute(query, (stat_month_date, os, user_count, new_users, retained_users))

def process_monthly_data(start_month):
    current_month = datetime.strptime(start_month, '%Y-%m')
    end_month = datetime.now().strftime('%Y-%m')  # 终止于当前月份
    
    # 连接数据库
    speed_db_conn = get_connection('speed-db')
    collector_db_conn = get_connection('collector-speed-db')
    report_db_conn = get_connection('report-speed-db')
    
    with report_db_conn.cursor() as cursor:
        clear_t_user_report_monthly(cursor)
        report_db_conn.commit()  # 提交事务以清空 t_user_report_monthly 表

    while current_month.strftime('%Y-%m') <= end_month:
        month_str = current_month.strftime('%Y-%m')
        
        with speed_db_conn.cursor() as cursor:
            # 获取注册用户的统计信息
            registered_results = get_registered_users_by_month(month_str, speed_db_conn)
            
            # 获取设备类型和计数
            device_results = get_device_types_and_counts(month_str, speed_db_conn, device_mapping)
            
            # 获取新用户统计信息
            new_users_results = get_new_users_in_month(month_str, speed_db_conn)
            
            # 新用户按设备类型分组
            new_users_by_os = {os_type: 0 for os_type in os_types}
            
            for row in new_users_results:
                original_os, email, new_users = row
                categorized_os = categorize_os(original_os)
                new_users_by_os[categorized_os] += new_users
            
            # 计算次月留存用户数
            retained_users_by_os = calculate_retention_of_next_month(month_str, speed_db_conn, collector_db_conn)
            
            # 处理results并插入到t_user_report_monthly中
            if registered_results:
                stat_month = registered_results[0][0]  # 获取正确的stat_month
                user_count = sum([row[2] for row in registered_results])  # 计算总用户数
            else:
                stat_month = month_str
                user_count = 0
            
            # 设备类型和计数
            device_os_counts = {row[0]: row[1] for row in device_results}
            
            # 插入数据
            with report_db_conn.cursor() as cursor:
                for os in os_types:
                    total_new_users = new_users_by_os.get(os, 0)
                    total_device_count = device_os_counts.get(os, 0)
                    retained_users_count = len(retained_users_by_os.get(os, set()))
                    insert_into_report(stat_month, os, user_count, total_new_users, retained_users_count, cursor)
                report_db_conn.commit()  # 提交事务
        
        # 移动到下一个月
        current_month += timedelta(days=32)
        current_month = current_month.replace(day=1)

if __name__ == '__main__':
    lock_file = f"/tmp/{TASK_NAME}.lock"
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.error(f"已经有一个 {TASK_NAME} 进程在运行，本进程将退出")
        sys.exit(1)

    init_logging(f"/shell/report/log/{TASK_NAME}.log")
    logging.info(f"\n\n\n start {TASK_NAME}")
    try:
        process_monthly_data(start_month)
    except Exception as e:
        logging.error(f"捕获到异常：{type(e).__name__}")
        logging.error(f"异常信息：{str(e)}")
        logging.error(traceback.format_exc())
    finally:
        fp.close()

    logging.info(f"end {TASK_NAME}")