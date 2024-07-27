# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import time
import fcntl
import traceback

import logging

import log
import db_util_speed
import util

TASK_NAME = "clean_report_data_task"


# * * * * * /usr/bin/python3 /shell/clean_report_data.py &

from datetime import datetime, timedelta


def run():
    """ 清理report-log """
    speed_report = db_util_speed.SpeedReport()
    speed_report.clean_data("t_user_op_log", 7)
    speed_report.clean_data("t_user_ping", 7)
    speed_report.clean_data("t_user_online_day", 60)
    speed_report.clean_data("t_user_node_online_day", 60)

    speed = db_util_speed.Speed()
    speed.clean_data("t_user_traffic_log", 60)
    speed.clean_data("t_user_traffic", 60)


if __name__ == '__main__':
    lock_file = "/tmp/%s.lock" % TASK_NAME
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.error("已经有一个 %s 进程在运行，本进程将退出" % TASK_NAME)
        sys.exit(1)

    log.init_logging("/shell/report/log/%s.log" % TASK_NAME)
    logging.info("\n\n\n start %s" % TASK_NAME)
    try:
        run()
    except Exception as e:
        # 这里处理异常
        logging.error(f"捕获到异常：{type(e).__name__}")
        logging.error(f"异常信息：{str(e)}")
        logging.error(traceback.format_exc())

    logging.info("end %s" % TASK_NAME)
