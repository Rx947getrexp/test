# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import fcntl
import traceback

import logging

import log
import user
import util

TASK_NAME = "report_task"


# * * * * * /usr/bin/python3 /shell/report_user.py &


def run():
    """ 补充历史数据 """
    # 将字符串转换为datetime对象
    # start_date_str = "2023-12-05"
    # end_date_str = "2024-01-19"
    # start_date = datetime.strptime(start_date_str, "%Y-%m-%d")
    # end_date = datetime.strptime(end_date_str, "%Y-%m-%d")
    #
    # # 遍历日期范围并打印每一天的日期
    # current_date = start_date
    # while current_date <= end_date:
    #     t = util.Time(current_date.strftime("%Y-%m-%d"))
    #     user.ReportUser(t.date, t.get_start_time(), t.get_end_time()).run()
    #     current_date += timedelta(days=1)

    # """ 统计昨天的数据 """
    t = util.Time(util.get_yesterday_date())
    user.ReportUser(t.date, t.get_start_time(), t.get_end_time()).run()


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
