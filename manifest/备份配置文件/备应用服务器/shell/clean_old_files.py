# -*- coding: utf-8 -*-
import os
import glob
import time
import sys
import fcntl

import logging
from logging.handlers import RotatingFileHandler

def init_logging(file_name):
    log = logging.getLogger()
    log.setLevel(logging.DEBUG)
    formatter = logging.Formatter('%(asctime)s - [line]:%(lineno)d - %(levelname)s - %(message)s',
                                  '%Y-%m-%d %H:%M:%S')

    ch = logging.StreamHandler()  # 输出到控制台的handler
    ch.setFormatter(formatter)
    ch.setLevel(logging.DEBUG)  # 也可以不设置，不设置就默认用logger的level

    filehandler = RotatingFileHandler(filename=file_name, mode='a', maxBytes=1024 * 1024 * 200, backupCount=2)
    filehandler.setFormatter(formatter)
    log.addHandler(filehandler)
    logging.info("init_logging success")

def clean_old_sql_invalid_files():
    # 获取 /shell/sql_backup 路径下的所有 .sql 文件
    file_list = glob.glob('/shell/sql_backup/*.sql')

    # 获取当前时间戳
    current_time = time.time()

    # 遍历文件列表
    for f in file_list:
        logging.info("file: %s" % f)
        file_size = os.path.getsize(f)
        file_mtime = os.path.getmtime(f)

        # 计算文件的修改时间与当前时间的差异（单位：秒）
        time_diff = current_time - file_mtime

        # 判断文件是否小于 100M 且创建时间在 7 天前
        if file_size < 100 * 1024 * 1024 and time_diff > 7 * 24 * 60 * 60:
            logging.info("remove invalid file: %s" % f)
            os.remove(f)  # 删除文件

def clean_old_sql_files():
    # 获取 /shell/sql_backup 路径下的所有 .sql 文件
    file_list = glob.glob('/shell/sql_backup/*.sql')

    # 筛选出大于 200M 的文件
    # file_list = [f for f in file_list if os.path.getsize(f) > 200 * 1024 * 1024]

    # 按照修改时间降序排序
    file_list.sort(key=os.path.getmtime, reverse=True)

    logging.info("sql file len(files): %d" % len(file_list))
    # 保留创建时间最近的 100 个文件
    if len(file_list) > 100:
        for f in file_list[100:]:
            logging.info("remove old sql file: %s" % f)
            os.remove(f)  # 删除文件


if __name__ == '__main__':
    task_name = "clean_old_files"
    lock_file = "/tmp/%s.lock" % task_name
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.info("已经有一个 %s 进程在运行，本进程将退出" % task_name)
        sys.exit(1)

    init_logging("/shell/log/%s.log" % task_name)
    logging.info("start clean")
    # clean_old_sql_invalid_files()
    clean_old_sql_files()
