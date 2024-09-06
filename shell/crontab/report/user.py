# -*- coding: utf-8 -*-
import sys
from datetime import timedelta, datetime
import logging
import requests
import db_util
import util


class ReportUser:
    def __init__(self, date, start_time, end_time,month_start_time):
        self.date = date
        self.start_time = start_time
        self.end_time = end_time
        self.month_start_time = month_start_time
        self.db_speed_conn = db_util.Speed()
        self.db_report_conn = db_util.SpeedReport()

    def run(self):
        self.report_daily_user()
        #self.report_online_user()
        #self.report_daily_user_recharge()
        #self.report_daily_user_recharge_times()
        #self.report_daily_channel_user_recharge_times()
        #self.report_daily_channel_user()
        #self.report_daily_node()
        #self.report_online_node_user()
        #self.device_recharge_behavior()
        #self.device_recharges()
        #self.report_daily_channel_recharge_by_month()
        self.db_speed_conn.close_connection()
        self.db_report_conn.close_connection()

    def report_daily_user_recharge(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        """获取商品id列表"""
        rows = self.db_speed_conn.query_goods_id_list()
        data = {}
        for row in rows:
            goods_id = row["id"]
            data[goods_id] = {}
            """ 套餐充值人数总量 """
            data[goods_id]["total_cnt"] = self.db_speed_conn.count_total_user_recharge(goods_id, self.end_time)
            """ 套餐新增用户数量 """
            data[goods_id]["new_cnt"] = self.db_speed_conn.count_user_recharge_by_create_time(goods_id, self.start_time,
                                                                                              self.end_time)

        self.db_report_conn.insert_daily_user_recharge(self.date, data)

    def report_daily_user_recharge_times(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        """获取商品id列表"""
        rows = self.db_speed_conn.query_goods_id_list()
        data = {}
        for row in rows:
            goods_id = row["id"]
            data[goods_id] = {}
            data[goods_id]["total_cnt"] = self.db_speed_conn.count_total_user_recharge_times(goods_id, self.end_time)
            data[goods_id]["new_cnt"] = self.db_speed_conn.count_user_recharge_times_by_create_time(goods_id,self.start_time,self.end_time)

        self.db_report_conn.insert_daily_user_recharge_times(self.date, data)

    def report_daily_channel_user_recharge_times(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        """获取商品id列表"""
        goods_list = self.db_speed_conn.query_goods_id_list()
    #today = "2024-07-21 00:00:00"
        """获取不同币种的列表"""
        currency_list = self.db_speed_conn.query_currency_list()
        rows = self.db_speed_conn.query_recharge_channel_list(self.end_time)
        data = {}
        for row in rows:
            channel = row["channel"]
            data[channel] = {}
            for goods in goods_list:
                goods_id = goods["id"]
                data[channel][goods_id] = {}
                for currencies in currency_list:
                    currency = currencies["currency"]
                    """按渠道套餐充值总次数"""
                    total_cnt = self.db_speed_conn.count_recharge_total_channel_user(channel,
                                                                                     goods_id,
                                                                                     currency,
                                                                                     self.end_time)
                    """按渠道套餐新增次数"""
                    new_cnt = self.db_speed_conn.count_recharge_channel_user_by_create_time(channel, goods_id, currency,
                                                                                            self.start_time,
                                                                                            self.end_time)
                    data[channel][goods_id][currency] = {
                        "total_cnt": total_cnt,
                        "new_cnt": new_cnt
                    }

        self.db_report_conn.insert_daily_channel_user_recharge(self.date, data)

    def report_daily_node(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        """ 获取ip列表 """
        rows = self.db_speed_conn.query_node_traffic_list(self.end_time)
        data = {}
        for row in rows:
            ip = row["ip"]
            data[ip] = {}
            """ 节点使用总量 """
            data[ip]["total_cnt"] = self.db_speed_conn.count_total_node(ip, self.end_time)
            """ 新增用户数量 """
            data[ip]["new_cnt"] = self.db_speed_conn.count_node_by_create_time(ip, self.start_time, self.end_time)
            """ 节点使用留存 """
            data[ip]["retained_cnt"] = self.db_speed_conn.count_node_online(ip, self.date,
                                                                            self.end_time)

        self.db_report_conn.insert_daily_node(self.date, data)

    def report_daily_user(self):
        """
        用户报表
        """
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)

        """ 获取channel_id列表 """
        rows = self.db_speed_conn.query_channel_id_list(self.end_time)
        data = {}
        for row in rows:
            channel_id = row["channel_id"]
            data[channel_id] = {}
            """ 用户总量 """
            data[channel_id]["total_cnt"] = self.db_speed_conn.count_total_user(channel_id, self.end_time)

            """ 新增用户数量 """
            data[channel_id]["new_cnt"] = self.db_speed_conn.count_user_by_create_time(channel_id, self.start_time,
                                                                                       self.end_time)

            """ 留存用户数量 """
            data[channel_id]["retained_cnt"] = self.db_speed_conn.count_user_online(channel_id, self.date,self.end_time)
            """ 月留存用户数量 """
            data[channel_id]["month_retained_cnt"] = self.db_speed_conn.count_user_month_online(channel_id,self.month_start_time,self.end_time)

        self.db_report_conn.insert_daily_user(self.date, data)

    def report_daily_channel_user(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)

        """ 获取channel列表 """
        rows = self.db_speed_conn.query_channel_list(self.end_time)
        data = {}
        for row in rows:
            channel = row["channel"]
            data[channel] = {}
            """ 推广渠道用户总量 """
            data[channel]["total_cnt"] = self.db_speed_conn.count_total_channel_user(channel, self.end_time)

            """ 新增推广渠道用户数量 """
            data[channel]["new_cnt"] = self.db_speed_conn.count_channel_user_by_create_time(channel, self.start_time,
                                                                                            self.end_time)

            """ 新增推广渠道用户数量 """
            data[channel]["retained_cnt"] = self.db_speed_conn.count_channel_user_online(channel, self.date,
                                                                                         self.end_time)

            """ 充值总人数 """
            data[channel]["total_recharge"] = self.db_speed_conn.count_total_number_of_recharges(channel, self.end_time)

            """ 充值总金额 """
            total_recharge_amount = self.db_speed_conn.count_total_recharge_amount(channel, self.end_time)
            data[channel]["total_recharge_amount"] = total_recharge_amount

            """ 新增充值金额 """
            new_recharge_amount = self.db_speed_conn.count_total_recharge_amount_by_create_time(channel,
                                                                                                self.start_time,
                                                                                                self.end_time)
            data[channel]["new_recharge_amount"] = new_recharge_amount

        self.db_report_conn.insert_daily_channel_user(self.date, data)

    def report_online_user(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        user_info = self.db_speed_conn.get_users()
        user_ip_info = self.db_speed_conn.get_user_ip_list()
        rows = self.db_speed_conn.query_user_traffic_list(self.date)
        user_online_data = []
        for row in rows:
            email = row["email"]
            logs = self.db_speed_conn.query_user_traffic_log_list(email, self.start_time, self.end_time)
            if len(logs) == 0:
                continue

            uplink, downlink = 0, 0
            for log in logs:
                uplink = uplink + log["uplink"]
                downlink = downlink + log["downlink"]

            total_duration = timedelta()
            gap_duration = timedelta(minutes=5)

            for i in range(len(logs) - 1):
                t1 = util.time_format(logs[i]["date_time"])
                t2 = util.time_format(logs[i + 1]["date_time"])
                time_diff = t2 - t1
                if time_diff <= timedelta(minutes=15):
                    total_duration += time_diff
                else:
                    total_duration += time_diff + gap_duration
            logging.info(total_duration)
            online_time = int(total_duration.total_seconds())
            logging.info(online_time)
            if online_time == 0:
                online_time = 5 * 60
            elif online_time > 60 * 60 * 24:
                online_time = 60 * 60 * 24
            user_email = user_info.get(email, None)
            if user_email == None:
                continue
            user_online_info = {
                "email": email,
                "online_duration": online_time,
                "uplink": uplink,
                "downlink": downlink,
                "country": 'unknown',
                # "register_date": user_info.get("register_date"),
                "channel": user_email.get("channel", None)
            }
            if email in user_ip_info.keys():
                user_ip = user_ip_info[email]["ip"]
                get_country = f"https://ipinfo.io/{user_ip}/country"
                try:
                    country = requests.get(get_country).text.strip()
                except:
                    country="None"
                if country == "None" or len(country) > 5:
                    reader = util.IpSearch()
                    country = reader.get_location(user_ip)
                user_online_info["country"] = country
            user_online_data.append(user_online_info)
        logging.info(user_online_data)
        self.db_report_conn.insert_online_user_day(self.date, user_online_data)

    def report_online_node_user(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)

        user_info = self.db_speed_conn.get_users()
        rows = self.db_speed_conn.query_node_user_traffic_list(self.date)
        user_online_data = []
        for row in rows:
            email = row["email"]
            ip = row["ip"]
            uplink = row["uplink"]
            downlink = row["downlink"]
            logs = self.db_speed_conn.query_user_traffic_log_list(email, self.start_time, self.end_time)
            if len(logs) == 0:
                continue
            filtered_logs = [log for log in logs if log["ip"] == ip]
            if len(filtered_logs) == 0:
                continue
            total_duration = timedelta()
            gap_duration = timedelta(minutes=5)
            for i in range(len(filtered_logs) - 1):
                t1 = util.time_format(filtered_logs[i]["date_time"])
                t2 = util.time_format(filtered_logs[i + 1]["date_time"])
                time_diff = t2 - t1
                if time_diff <= timedelta(minutes=15):
                    total_duration += time_diff
                else:
                    total_duration += time_diff + gap_duration
            logging.info(total_duration)
            online_time = int(total_duration.total_seconds())
            logging.info(online_time)
            if online_time == 0:
                online_time = 5 * 60
            elif online_time > 60 * 60 * 24:
                online_time = 60 * 60 * 24
            register_date = user_info.get(email, None)
            if register_date == None:
                continue
            user_online_info = {
                "email": email,
                "online_duration": online_time,
                "uplink": uplink,
                "downlink": downlink,
                "node": ip,
                "register_date": register_date.get("register_date", None),
                "channel": user_info.get(email, None).get("channel", None)
            }
            user_online_data.append(user_online_info)
        logging.info(user_online_data)
        self.db_report_conn.insert_online_user_node_day(self.date, user_online_data)

    def device_recharge_behavior(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)

        """ 获取用户device列表 """
        rows = self.db_report_conn.query_devices_list(self.end_time)
        data = {}
        end_time = datetime.strptime(self.end_time, "%Y-%m-%d %H:%M:%S")
        week_ago_time = end_time - timedelta(days=7)
        for row in rows:
            device = row["device"]
            data[device] = {}
            data[device]["total_clicks"] = self.db_report_conn.get_total_clicks(device, self.end_time)
            data[device]["yesterday_day_clicks"] = self.db_report_conn.get_yesterday_day_clicks(device, self.start_time,
                                                                                                self.end_time)
            data[device]["weekly_clicks"] = self.db_report_conn.get_weekly_clicks(device, week_ago_time, self.end_time)
            data[device]["total_users_clicked"] = self.db_report_conn.get_total_users_clicked(device, self.end_time)
            data[device]["yesterday_day_users_clicked"] = self.db_report_conn.get_yesterday_day_users_clicked(device,
                                                                                                              self.start_time,
                                                                                                              self.end_time)
            data[device]["weekly_users_clicked"] = self.db_report_conn.get_weekly_users_clicked(device, week_ago_time,
                                                                                                self.end_time)
        self.db_report_conn.insert_daily_device_action(self.date, data)

    def device_recharges(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)

        """ 获取用户device列表 """
        rows = self.db_speed_conn.query_device_list(self.end_time)
        data = {}
        data['无设备'] = {}
        for row in rows:
            device = row["device"]
            data[device] = {}
            """ 设备类型用户总量 """
            data[device]["total_cnt"] = self.db_speed_conn.count_total_device_user(device, self.end_time)

            """ 设备类型用户新增数量 """
            data[device]["new_cnt"] = self.db_speed_conn.count_device_user_by_create_time(device, self.start_time,
                                                                                          self.end_time)

            """ 设备类型用户留存"""
            data[device]["retained_cnt"] = self.db_speed_conn.count_device_user_online(device, self.date,
                                                                                       self.end_time)

            """ 充值总人数 """
            data[device]["total_recharge"] = self.db_speed_conn.count_total_number_of_device_recharges(device,self.end_time)

            """ 充值总金额 """
            total_recharge_amount = self.db_speed_conn.count_total_device_recharge_amount(device, self.end_time)
            data[device]["total_recharge_amount"] = total_recharge_amount

            """ 新增充值金额 """
            new_recharge_amount = self.db_speed_conn.count_total_device_recharge_amount_by_create_time(device,
                                                                                                       self.start_time,
                                                                                                       self.end_time)
            data[device]["new_recharge_amount"] = new_recharge_amount
        data['无设备']["total_cnt"] = self.db_speed_conn.count_total_user(1, self.end_time)
        data['无设备']["new_cnt"] = self.db_speed_conn.count_user_by_create_time(1, self.start_time, self.end_time)
        data['无设备']["retained_cnt"] = self.db_speed_conn.count_user_online(1, self.date, self.end_time)
        data['无设备']['total_recharge'] = self.db_speed_conn.all_device_recharge(self.end_time)
        data['无设备']['total_recharge_amount'] = self.db_speed_conn.all_device_total_recharge_amount(self.end_time)
        data['无设备']['new_recharge_amount'] = self.db_speed_conn.all_device_total_recharge_amount_by_create_time(
            self.start_time, self.end_time)
        equipped_total_cnt = 0
        equipped_new_cnt = 0
        equipped_retained_cnt = 0
        equipped_total_recharge = 0
        equipped_total_recharge_amount = 0
        equipped_new_recharge_amount = 0
        for device in data:
            if device != '无设备':
                total_cnt = data[device].get('total_cnt')
                equipped_total_cnt += total_cnt
                new_cnt = data[device].get('new_cnt')
                equipped_new_cnt += new_cnt
                retained_cnt = data[device].get('retained_cnt')
                equipped_retained_cnt += retained_cnt
                total_recharge = data[device].get('total_recharge')
                equipped_total_recharge += total_recharge
                total_recharge_amount = data[device].get('total_recharge_amount')
                equipped_total_recharge_amount += total_recharge_amount
                new_recharge_amount = data[device].get('new_recharge_amount')
                equipped_new_recharge_amount += new_recharge_amount
        data['无设备']['total_cnt'] = data['无设备']['total_cnt'] - equipped_total_cnt
        data['无设备']['new_cnt'] = data['无设备']['new_cnt'] - equipped_new_cnt
        data['无设备']['retained_cnt'] = data['无设备']['retained_cnt'] - equipped_retained_cnt
        data['无设备']['total_recharge'] = data['无设备']['total_recharge'] - equipped_total_recharge
        data['无设备']['total_recharge_amount'] = data['无设备']['total_recharge_amount'] - equipped_total_recharge_amount
        data['无设备']['new_recharge_amount'] = data['无设备']['new_recharge_amount'] - equipped_new_recharge_amount
        self.db_report_conn.insert_daily_device_user(self.date, data)

    def report_daily_channel_recharge_by_month(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        """ 获取channel列表 """
        rows = self.db_speed_conn.query_channel_list(self.end_time)
        data = {}
        for row in rows:
            channel = row["channel"]
            data[channel] = {}
            """用户总量"""
            data[channel]["total_cnt"] = self.db_speed_conn.count_total_channel_user(channel, self.end_time)
            """月用户总量"""
            data[channel]["month_retained_cnt"] = self.db_speed_conn.count_channel_user_by_create_time(channel,self.month_start_time,self.end_time)
            """月用户新增"""
            data[channel]["month_new_cnt"] = self.db_speed_conn.count_channel_user_by_month(channel,self.month_start_time,self.end_time)
            """总充值人数"""
            data[channel]["total_recharge_cnt"] = self.db_speed_conn.count_total_number_of_recharges(channel, self.end_time)
            """总充值金额"""
            data[channel]["total_recharge_money_cnt"] = self.db_speed_conn.count_total_recharge_amount(channel, self.end_time)
            """月充值人数"""
            data[channel]["month_recharge_cnt"] = self.db_speed_conn.count_total_number_of_recharges_by_create_time(channel, self.month_start_time,self.end_time)
            """月充值金额"""
            data[channel]["month_recharge_money_cnt"] = self.db_speed_conn.count_total_recharge_amount_by_create_time(channel, self.month_start_time,self.end_time)
        self.db_report_conn.insert_daily_channel_recharge_by_month(self.date, data)
