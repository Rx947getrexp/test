# -*- coding: utf-8 -*-
import sys

import logging
import db_util


class ReportUser:
    def __init__(self, date, start_time, end_time):
        self.date = date
        self.start_time = start_time
        self.end_time = end_time

    def run(self):
        self.report_daily_user()

    def report_daily_user(self):
        """
        用户报表
        """
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)

        db_speed = db_util.Speed()

        """ 获取channel_id列表 """
        rows = db_speed.query_channel_id_list(self.end_time)
        data = {}
        for row in rows:
            channel = row["channel"]
            channel_id = 0
            if channel.isdigit():
                channel_id = int(channel)

            data[channel_id] = {}
            """ 用户总量 """
            data[channel_id]["total_cnt"] = db_speed.count_total_user(channel, self.end_time)

            """ 新增用户数量 """
            data[channel_id]["new_cnt"] = db_speed.count_user_by_create_time(channel, self.start_time, self.end_time)

            """ 留存用户数量 """
            data[channel_id]["retained_cnt"] = db_speed.count_user_online(channel, self.date, self.end_time)

        db_util.SpeedReport().insert_daily_user(self.date, data)
