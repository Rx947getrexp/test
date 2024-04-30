# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import time
import fcntl
from datetime import datetime, timedelta


def get_yesterday_date():
    # 获取今天的日期
    today = datetime.now()

    # 计算昨天的日期
    yesterday = today - timedelta(days=1)

    # 将昨天的日期格式化为 "2024-01-01" 格式
    return yesterday.strftime('%Y-%m-%d')


class Time:
    def __init__(self, date):
        self.class_name = "util.Time"
        self.date = date

    def get_start_time(self):
        return self.date + " 00:00:00"

    def get_end_time(self):
        return self.date + " 23:59:59"
