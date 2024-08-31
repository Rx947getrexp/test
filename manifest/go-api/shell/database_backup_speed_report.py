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

import_host = '43.131.37.240'
import_user = 'speed_backup'
import_password = 'bakIUY*&^*^!12H6!326'

# databases_to_export = ['speed', 'go_fly2'] #, 'speed_report']
databases_to_export = ['speed', 'go_fly2', 'speed_report']

def export_databases(export_file_path):
    # 导出数据库
    command = f'mysqldump --user={export_user} --databases {" ".join(databases_to_export)} > {export_file_path}'
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


if __name__ == '__main__':
    task_name = "database_backup_speed_report"
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
            # import_databases(file_name)
            logging.info("-" * 20 + "备份成功")

        logging.info("begin to sleep 10min")
        # 间隔10分钟
        time.sleep(10*60)
        logging.info("\n\n\n")
        break