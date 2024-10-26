# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import time
import fcntl
from datetime import datetime

import logging
# from logging.handlers import TimedRotatingFileHandler
from logging.handlers import RotatingFileHandler

import os
import glob

# 配置信息
backup_dir = '/shell/database_backup'
max_files = 6

# mysqldump -u root -p --no-data --databases speed go_fly2 speed_report > all_databases_structure.sql

def cleanup_old_files():
    # 获取目录下所有文件
    files = glob.glob(f'{backup_dir}/*.sql')

    # 按创建时间排序
    files.sort(key=lambda x: os.path.getctime(x), reverse=True)

    # 删除多余的文件
    for file in files[max_files:]:
        os.remove(file)
        logging.info(f'删除文件：{file}')


# * * * * * /usr/bin/python3 /shell/monitor_api_redis.py &
def init_logging(file_name):
    log = logging.getLogger()
    log.setLevel(logging.DEBUG)
    formatter = logging.Formatter('%(asctime)s - [line]:%(lineno)d - %(levelname)s - %(message)s',
                                  '%Y-%m-%d %H:%M:%S')

    ch = logging.StreamHandler()  # 输出到控制台的handler
    ch.setFormatter(formatter)
    ch.setLevel(logging.DEBUG)  # 也可以不设置，不设置就默认用logger的level

    # log_file_handler = TimedRotatingFileHandler(filename=file_name, when="D", interval=1, backupCount=14)
    # log_file_handler.setFormatter(formatter)
    # log.addHandler(log_file_handler)

    filehandler = RotatingFileHandler(filename=file_name, mode='a', maxBytes=1024 * 1024 * 200, backupCount=2)
    filehandler.setFormatter(formatter)
    log.addHandler(filehandler)
    logging.info("init_logging success")



# 配置信息
export_host = '127.0.0.1'
export_user = 'root'
export_password = 'IUY*&^*^!12312HGJHG886!32'

import_host = '185.22.154.21'
import_user = 'speed_backup'
import_password = 'bakIUY*&^*^!12H6!326oihjh*(78712YH129-,IUTCJGFZA6761HGqw[ooooPPPP'

databases_to_export = ['speed', 'go_fly2'] #, 'speed_report']

ignore_tables = ["--ignore-table=speed.t_user_traffic_log",
                "--ignore-table=speed.user_logs",
                 "--ignore-table=speed.t_user_traffic",
                 ]

def export_databases(export_file_path):
    # 导出数据库
    command = f'mysqldump --user={export_user} --single-transaction --databases {" ".join(databases_to_export)} {" ".join(ignore_tables)} > {export_file_path}'
    logging.info(command)
    result = subprocess.run(command, shell=True, check=True)
    logging.info(result)
    # 检查命令返回状态
    if result.returncode == 0:
        logging.info(f'数据库已成功导出到 {export_file_path}')
        return True
    else:
        logging.error(f'导出数据库时出错，错误代码：{result.returncode}')
        return False

def import_databases(export_file_path):
    # 导入数据库
    if is_between_2_and_4():
        logging.info("当前时间在凌晨3点到早上12点之间，不更新DB，避免影响报表")
        return
    else:
        logging.info("当前时间不在凌晨3点到临晨12点之间，同步 speed 库的数据")

    command = f'mysql --host={import_host} --user={import_user} --password=\'{import_password}\' < {export_file_path}'
    logging.info(command)
    result = subprocess.run(command, shell=True, check=True)
    logging.info(result)

def execute_cmd(command):
    logging.info(command)
    result = subprocess.run(command, shell=True, check=True)
    logging.info(result)


def gen_backup_filename():
    now = datetime.now()
    timestamp = now.strftime("%Y-%m-%d_%H-%M-%S")
    return f'/shell/database_backup/speed-{timestamp}.sql'

def is_between_2_and_4():
    now = datetime.now()
    current_hour = now.hour
    return 3 <= current_hour < 12


if __name__ == '__main__':
    task_name = "database_backup_speed"
    lock_file = "/tmp/%s.lock" % task_name
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.info("已经有一个 %s 进程在运行，本进程将退出" % task_name)
        sys.exit(1)

    execute_cmd("mkdir -p /shell/database_backup")

    init_logging("/shell/log/%s.log" % task_name)
    logging.info("start backup")
    while True:
        logging.info("-"*20+"开始备份")
        cleanup_old_files()
        file_name = gen_backup_filename()

        ret = export_databases(file_name)
        # rsync -avz -e "ssh -i /root/.ssh/id_rsa" /shell/database_backup/speed-2024-08-12_02-37-48.sql user@remote_host:/path/to/remote/backup
        if ret:
            execute_cmd("scp %s root@185.22.152.47:/shell/sql_backup/" % file_name)
            execute_cmd("scp %s root@185.22.154.21:/shell/sql_backup/" % file_name)
            execute_cmd("scp %s root@45.251.243.140:/shell/sql_backup/" % file_name)
            import_databases(file_name)
            logging.info("-" * 20 + "备份成功")

            logging.info("begin to sleep 2h")
            time.sleep(60*60*2)
            logging.info("\n\n\n")
