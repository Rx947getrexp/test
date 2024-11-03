import sys
import fcntl
import logging
from logging.handlers import RotatingFileHandler
import traceback
from db_util import get_connection
from datetime import datetime, timedelta
from dateutil.relativedelta import relativedelta

TASK_NAME = "user_monthly_retention_task"

def get_two_months_ago():
    two_months_ago = datetime.now() - relativedelta(months=2)
    return two_months_ago.strftime('%Y-%m')

# 指定起始月份
start_month = get_two_months_ago()

# 设备类型映射
device_mapping = {
    'Android': ['Android'],
    'iPhone': ['iPhone', 'iOS'],
    'Mac': ['Mac', 'macOS', 'Mac OS'],
    'Windows': ['Windows'],
    'Others': []  # 任何未匹配的都归为 Others
}

# 设备类型列表
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

def get_registered_users_by_month(month, db_conn):
    query = """
        SELECT 
            DATE_FORMAT(tu.created_at, '%%Y-%%m') AS stat_month,
            tu.email,
            COUNT(*) AS user_count
        FROM 
            t_user tu
        WHERE 
            DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s
        GROUP BY 
            tu.email, stat_month;
    """
    with db_conn.cursor() as cursor:
        cursor.execute(query, (month,))
        return cursor.fetchall()

def get_device_types_and_counts(month, db_conn, device_mapping):
    query = """
        SELECT 
            tud.os AS os,
            COUNT(*) AS device_count
        FROM 
            t_user_device tud
        JOIN 
            t_user tu ON tud.user_id = tu.id
        WHERE 
            DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s
        GROUP BY 
            tud.os;
    """
    with db_conn.cursor() as cursor:
        cursor.execute(query, (month,))
        results = cursor.fetchall()
    
    categorized_counts = {os_type: 0 for os_type in device_mapping.keys()}
    
    for row in results:
        original_os = row[0]
        categorized_os = categorize_os(original_os)
        categorized_counts[categorized_os] += row[1]
    
    return [(os_type, count) for os_type, count in categorized_counts.items()]

def get_new_users_in_month(month, db_conn):
    query = """
        SELECT 
            tud.os,
            tu.email,
            COUNT(*) AS new_users
        FROM 
            t_user tu
        JOIN 
            t_user_device tud ON tu.id = tud.user_id
        WHERE 
            DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s
        GROUP BY 
            tud.os, tu.email;
    """
    with db_conn.cursor() as cursor:
        cursor.execute(query, (month,))
        return cursor.fetchall()

def get_registered_emails_by_month(month, db_conn):
    query = """
        SELECT 
            tu.email
        FROM 
            t_user tu
        WHERE 
            DATE_FORMAT(tu.created_at, '%%Y-%%m') = %s
    """
    with db_conn.cursor() as cursor:
        cursor.execute(query, (month,))
        return [row[0] for row in cursor.fetchall()]

def get_active_emails_in_next_month(month, traffic_db_conn):
    month_date = datetime.strptime(month, '%Y-%m')
    next_month_date = month_date + timedelta(days=32)
    next_month = next_month_date.strftime('%Y-%m')
    
    next_month_start = next_month_date.replace(day=1)
    next_month_end = (next_month_start + timedelta(days=32)).replace(day=1) - timedelta(days=1)
    
    query = """
        SELECT 
            tut.email
        FROM 
            t_v2ray_user_traffic tut
        WHERE 
            tut.date >= %s
            AND tut.date <= %s
    """
    with traffic_db_conn.cursor() as cursor:
        cursor.execute(query, (int(next_month_start.strftime('%Y%m%d')), int(next_month_end.strftime('%Y%m%d'))))
        return [row[0] for row in cursor.fetchall()]

def calculate_retention_of_next_month(month, speed_db_conn, traffic_db_conn):
    registered_emails = get_registered_emails_by_month(month, speed_db_conn)
    active_emails = set(get_active_emails_in_next_month(month, traffic_db_conn))  # 转换成集合
    
    if not registered_emails:
        return {os_type: set() for os_type in os_types}  # 如果没有注册用户，直接返回空结果
    
    registered_device_emails = {}
    query = """
        SELECT 
            tud.os, tu.email
        FROM 
            t_user_device tud
        JOIN 
            t_user tu ON tud.user_id = tu.id
        WHERE 
            tu.email IN %s;
    """
    with speed_db_conn.cursor() as cursor:
        cursor.execute(query, (tuple(registered_emails),))
        registered_device_emails = {email: os for os, email in cursor.fetchall()}

    retained_users_by_os = {os_type: set() for os_type in os_types}
    
    for email in registered_emails:
        if email in active_emails:
            original_os = registered_device_emails.get(email, 'Others')
            categorized_os = categorize_os(original_os)
            retained_users_by_os[categorized_os].add(email)
    
    return retained_users_by_os

def clear_t_user_report_monthly(cursor):
    query = """
        TRUNCATE TABLE t_user_report_monthly;
    """
    cursor.execute(query)

def insert_into_report(stat_month, os, user_count, new_users, retained_users, cursor):
    stat_month_date = int(stat_month.replace('-', ''))
    query = """
        INSERT INTO t_user_report_monthly (stat_month, os, user_count, new_users, retained_users)
        VALUES (%s, %s, %s, %s, %s);
    """
    cursor.execute(query, (stat_month_date, os, user_count, new_users, retained_users))

def process_monthly_data(start_month):
    current_month = datetime.strptime(start_month, '%Y-%m')
    end_month = datetime.now().strftime('%Y-%m')  # 终止于当前月份
    
    # 连接数据库
    speed_db_conn = get_connection('speed')
    traffic_db_conn = get_connection('speed_collector')
    report_db_conn = get_connection('speed_report')
    
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
            retained_users_by_os = calculate_retention_of_next_month(month_str, speed_db_conn, traffic_db_conn)
            
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
def init_logging(file_name):
    log = logging.getLogger()
    log.setLevel(logging.DEBUG)
    formatter = logging.Formatter('%(asctime)s - [line]:%(lineno)d - %(levelname)s - %(message)s','%Y-%m-%d %H:%M:%S')

    ch = logging.StreamHandler()  # 输出到控制台的handler
    ch.setFormatter(formatter)
    ch.setLevel(logging.DEBUG)  # 也可以不设置，不设置就默认用logger的level

    handler = RotatingFileHandler(filename=file_name, mode='a', maxBytes=1024 * 1024 * 200, backupCount=2)
    handler.setFormatter(formatter)
    log.addHandler(handler)
    logging.info("init_logging success")


if __name__ == '__main__':
    lock_file = "/tmp/%s.lock" % TASK_NAME
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.error("已经有一个 %s 进程在运行，本进程将退出" % TASK_NAME)
        sys.exit(1)

    init_logging("/shell/retention_report/log/%s.log" % TASK_NAME)
    logging.info("\n\n\n start %s" % TASK_NAME)
    try:
        process_monthly_data(start_month)
    except Exception as e:
        # 这里处理异常
        logging.error(f"捕获到异常：{type(e).__name__}")
        logging.error(f"异常信息：{str(e)}")
        logging.error(traceback.format_exc())

    logging.info("end %s" % TASK_NAME)