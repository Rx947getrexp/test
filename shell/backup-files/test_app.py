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

def execute_cmd(command):
    # 记录开始时间
    start_time = time.time()
    logging.info(subprocess.run(command, shell=True, check=True, capture_output=True))
    # 计算耗时并打印
    elapsed_time = time.time() - start_time
    logging.info(f"调用函数 f1 的耗时为：{elapsed_time:.6f} 秒\n\n")


if __name__ == '__main__':
    task_name = "test_hs_app"
    lock_file = "/tmp/%s.lock" % task_name
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.info("已经有一个 %s 进程在运行，本进程将退出" % task_name)
        sys.exit(1)

    init_logging("/shell/log/%s.log" % task_name)
    logging.info("start test")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://eigrrht.xyz/app-api/dns_list""")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://siaax.xyz/app-api/dns_list""")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://beiyo.xyz/app-api/dns_list""")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://thertee.xyz/app-api/dns_list""")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://weechat.xyz/app-api/dns_list""")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://2yiny.xyz/app-api/dns_list""")
    execute_cmd("""curl -X GET -H "Content-Type: application/json" -H "Lang: cn" https://yinyong.xyz/app-api/dns_list""")
