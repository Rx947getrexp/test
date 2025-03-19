# -*- coding: utf-8 -*-
from decimal import Decimal
import sys
import time
from datetime import timedelta, datetime
import logging
import db_util
import util

class ReportUser:
    def __init__(self, date, start_time, end_time, month_start_time):
        logging.info("---> (%s) (%s) (%s) (%s)" % (date, start_time, end_time, month_start_time))
        self.date = date
        self.start_time = start_time
        self.end_time = end_time
        self.month_start_time = month_start_time
        self.db_speed_conn = db_util.Speed()
        self.db_report_conn = db_util.SpeedReport()
        self.db_collector_conn = db_util.SpeedCollector()

    def run(self):
        if self.check_collector():
            self.report_daily_user()
            self.report_daily_user_recharge()
            self.report_daily_user_recharge_times()
            self.report_daily_channel_user_recharge_times()
            self.report_daily_channel_user()
            self.report_daily_node()
            self.report_online_node_user()
            self.device_recharge_behavior()
            self.device_recharges()
            self.report_daily_channel_recharge_by_month()
            # self.report_device_retaind() #旧的统计7日，15日留存
            self.report_online_user()
            self.report_daily_device_retaind()  # 新的按设备统计次日，7日，15日留存
            self.report_monthly_device_retaind()  # 新的按设备统计次月留存
            self.report_ad_daily() #按日统计广告曝光量和点击量
            self.report_pay_daily() #按日统计支付金额
            self.report_channel_registration_pay_daily() #按日期获取注册用户数和付费用户数
            self.report_channel_retaind_daily() #按日期获取推广渠道的留存用户数
            self.report_channel_recharge_by_month() #按月统计推广渠道的充值用户数和充值金额
        self.db_speed_conn.close_connection()
        self.db_report_conn.close_connection()
        self.db_collector_conn.close_connection()

    def check_collector(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        flag = self.db_collector_conn.check_task(self.date)
        while not flag:
            self.db_collector_conn.close_connection()
            self.db_collector_conn = db_util.SpeedCollector()
            flag = self.db_collector_conn.check_task(self.date)
            time.sleep(600)
        return True

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
            data[goods_id]["new_cnt"] = self.db_speed_conn.count_user_recharge_by_create_time(goods_id, self.start_time, self.end_time)

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
            data[goods_id]["new_cnt"] = self.db_speed_conn.count_user_recharge_times_by_create_time(goods_id, self.start_time, self.end_time)

        self.db_report_conn.insert_daily_user_recharge_times(self.date, data)

    def report_daily_channel_user_recharge_times(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        """获取商品id列表"""
        goods_list = self.db_speed_conn.query_goods_id_list()
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
                    total_cnt = self.db_speed_conn.count_recharge_total_channel_user(channel, goods_id, currency, self.end_time)
                    """按渠道套餐新增次数"""
                    new_cnt = self.db_speed_conn.count_recharge_channel_user_by_create_time(channel, goods_id, currency, self.start_time, self.end_time)
                    data[channel][goods_id][currency] = {"total_cnt": total_cnt, "new_cnt": new_cnt}

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
            data[ip]["retained_cnt"] = self.db_speed_conn.count_node_online(ip, self.date, self.end_time)

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
            data[channel_id]["new_cnt"] = self.db_speed_conn.count_user_by_create_time(channel_id, self.start_time, self.end_time)
            """ 留存用户数量 """
            data[channel_id]["retained_cnt"] = self.db_speed_conn.count_user_online(channel_id, self.date, self.start_time, self.end_time)
            """月留存用户数量 """
            data[channel_id]["month_retained_cnt"] = self.db_speed_conn.count_user_month_online(channel_id, self.month_start_time, self.end_time)
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

            """新增推广渠道用户数量"""
            data[channel]["new_cnt"] = self.db_speed_conn.count_channel_user_by_create_time(channel, self.start_time, self.end_time)

            """推广渠道用户留存"""
            data[channel]["retained_cnt"] = self.db_speed_conn.count_channel_user_online(channel, self.date, self.start_time, self.end_time)

            """ 充值总人数 """
            data[channel]["total_recharge"] = self.db_speed_conn.count_total_number_of_recharges(channel, self.end_time)

            """ 充值总金额 """
            total_recharge_amount = self.db_speed_conn.count_total_recharge_amount(channel, self.end_time)
            data[channel]["total_recharge_amount"] = total_recharge_amount

            """ 新增充值金额 """
            new_recharge_amount = self.db_speed_conn.count_total_recharge_amount_by_create_time(channel, self.start_time, self.end_time)
            data[channel]["new_recharge_amount"] = new_recharge_amount

        self.db_report_conn.insert_daily_channel_user(self.date, data)

    def report_online_user(self):
        logging.info("*" * 20 + sys._getframe().f_code.co_name + "*" * 20)
        user_info = self.db_speed_conn.get_users()
        # user_ip_info = self.db_speed_conn.get_user_ip_list()
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
            last_login_ip = user_info[email]["last_login_ip"]
            if last_login_ip and last_login_ip != "":
                # get_country = f"https://ipinfo.io/{last_login_ip}/country"
                # try:
                #     country = requests.get(get_country).text.strip()
                # except:
                #     country="None"
                # if country == "None" or len(country) > 5:
                reader = util.IpSearch()
                country = reader.get_location(last_login_ip)
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
            data[device]["yesterday_day_clicks"] = self.db_report_conn.get_yesterday_day_clicks(device, self.start_time, self.end_time)
            data[device]["weekly_clicks"] = self.db_report_conn.get_weekly_clicks(device, week_ago_time, self.end_time)
            data[device]["total_users_clicked"] = self.db_report_conn.get_total_users_clicked(device, self.end_time)
            data[device]["yesterday_day_users_clicked"] = self.db_report_conn.get_yesterday_day_users_clicked(device, self.start_time, self.end_time)
            data[device]["weekly_users_clicked"] = self.db_report_conn.get_weekly_users_clicked(device, week_ago_time, self.end_time)
        self.db_report_conn.insert_daily_device_action(self.date, data)

    def device_recharges(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        data = {}
        rows = self.db_speed_conn.query_device_list(self.end_time)
        # 获取设备汇总数据: 总充值次数、总充值金额、新充值金额
        total_recharge = self.db_speed_conn.count_total_number_of_device_recharges(self.end_time)
        total_recharge_amount = self.db_speed_conn.count_total_device_recharge_amount(self.end_time)
        new_recharge_amount = self.db_speed_conn.count_total_device_recharge_amount_by_create_time(self.start_time, self.end_time)
        for device in set(total_recharge) | set(total_recharge_amount) | set(new_recharge_amount):
            data[device] = {
                "total_recharge": total_recharge.get(device, 0),
                "total_recharge_amount": total_recharge_amount.get(device, 0),
                "new_recharge_amount": new_recharge_amount.get(device, 0),
            }
        equipped_total_cnt = 0
        equipped_new_cnt = 0
        equipped_retained_cnt = 0
        for row in rows:
            device = row["device"]
            total_cnt = self.db_speed_conn.count_total_device_user(device, self.end_time)
            new_cnt = self.db_speed_conn.count_device_user_by_create_time(device, self.start_time, self.end_time)
            retained_cnt = self.db_speed_conn.count_device_user_online(device, self.date, self.start_time, self.end_time)
            equipped_total_cnt += total_cnt
            equipped_new_cnt += new_cnt
            equipped_retained_cnt += retained_cnt
            if device not in data:
                data[device] = {}
            data[device].update({
                "total_cnt": total_cnt,
                "new_cnt": new_cnt,
                "retained_cnt": retained_cnt
            })
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
            data[channel]["month_retained_cnt"] = self.db_speed_conn.count_channel_user_by_create_time(channel, self.month_start_time, self.end_time)
            """月用户新增"""
            data[channel]["month_new_cnt"] = self.db_speed_conn.count_channel_user_by_month(channel, self.month_start_time, self.end_time)
            """总充值人数"""
            data[channel]["total_recharge_cnt"] = self.db_speed_conn.count_total_number_of_recharges(channel, self.end_time)
            """总充值金额"""
            data[channel]["total_recharge_money_cnt"] = self.db_speed_conn.count_total_recharge_amount(channel, self.end_time)
            """月充值人数"""
            data[channel]["month_recharge_cnt"] = self.db_speed_conn.count_total_number_of_recharges_by_create_time(channel, self.month_start_time, self.end_time)
            """月充值金额"""
            data[channel]["month_recharge_money_cnt"] = self.db_speed_conn.count_total_recharge_amount_by_create_time(channel, self.month_start_time, self.end_time)
        self.db_report_conn.insert_daily_channel_recharge_by_month(self.date, data)

    def report_device_retaind(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        data = {}
        rows = self.db_speed_conn.query_device_list(self.end_time)
        end_date = datetime.strptime(self.date, '%Y-%m-%d')
        day7_date = end_date - timedelta(days=7)
        day15_date = end_date - timedelta(days=15)
        day7_start_time = end_date - timedelta(days=15)
        day15_start_time = end_date - timedelta(days=15)
        for row in rows:
            device = row["device"]
            # 获取新增用户数量
            new_cnt = self.db_speed_conn.count_device_user_by_create_time(device, self.start_time, self.end_time)
            # 获取次日留存
            retained_cnt = self.db_speed_conn.count_device_user_online(device, self.date, self.start_time, self.end_time)
            # 获取7日留存
            day7_retained = self.db_speed_conn.count_device_retained(device, day7_date.strftime('%Y-%m-%d'), day7_start_time, self.end_time)
            # 获取15日留存
            day15_retained = self.db_speed_conn.count_device_retained(device, day15_date.strftime('%Y-%m-%d'), day15_start_time, self.end_time)
            if device not in data:
                data[device] = {}
            data[device].update({
                "new_cnt": new_cnt,
                "retained_cnt": retained_cnt,
                "day7_retained": day7_retained,
                "day15_retained": day15_retained
            })
        self.db_report_conn.insert_daily_device_retention(self.date, data)

    def report_daily_device_retaind(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        # 获取前15天的数据并插入到t_user_device_retention中
        start_day = util.get_previous_days(15)  # 获取当前日期的前15天
        current_day = datetime.strptime(start_day, '%Y-%m-%d')
        end_day = datetime.now().strftime('%Y-%m-%d')

        while current_day.strftime('%Y-%m-%d') <= end_day:
            day_str = current_day.strftime('%Y-%m-%d')
            # 获取新用户统计信息
            new_users_results = self.db_speed_conn.query_registered_users_by_day(day_str)

            # 获取注册用户的设备类型
            registered_emails = self.db_speed_conn.query_registered_emails_by_day(day_str)
            registered_device_emails = self.db_speed_conn.query_registered_device_emails(registered_emails)
           
            # 新用户按设备类型分组
            new_users_by_os = {os_type: 0 for os_type in util.os_types}

            for row in new_users_results:
                _, email, _ = row
                original_os = registered_device_emails.get(email, 'Others')
                categorized_os = util.categorize_os(original_os)
                new_users_by_os[categorized_os] += 1

            # 计算次日留存用户数
            next_day_retention_by_os = self.db_speed_conn.query_retention_of_next_days(day_str, 1)
            
            # 计算7日留存用户数
            seven_days_retention_by_os = self.db_speed_conn.query_retention_of_next_days(day_str, 7)
            
            # 计算15日留存用户数
            fifteen_days_retention_by_os = self.db_speed_conn.query_retention_of_next_days(day_str, 15)

            # 处理results并插入到t_user_device_retention数据表中
            if new_users_results:
                stat_day = new_users_results[0][0]

                #user_count = sum([row[2] for row in registered_results])
            else:
                stat_day = day_str
                #user_count = 0

            for os in util.os_types:
                total_new_users = new_users_by_os.get(os, 0)
                next_day_retention_count = len(next_day_retention_by_os.get(os, set()))
                seven_days_retention_count = len(seven_days_retention_by_os.get(os, set()))
                fifteen_days_retention_count = len(fifteen_days_retention_by_os.get(os, set()))
                self.db_report_conn.insert_or_update_daily_device_report(stat_day, os, total_new_users, next_day_retention_count, seven_days_retention_count, fifteen_days_retention_count,)

            # 移动到下一天
            current_day += timedelta(days=1)

    def report_monthly_device_retaind(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        start_month = util.get_previous_months(2) #获取当前月份的前几个月

        current_month = datetime.strptime(start_month, '%Y-%m')
        end_month = datetime.now().strftime('%Y-%m')  # 终止于当前月份

        self.db_report_conn.clear_t_user_report_monthly()

        while current_month.strftime('%Y-%m') <= end_month:
            month_str = current_month.strftime('%Y-%m')
            # 获取注册用户的统计信息
            registered_results = self.db_speed_conn.get_registered_users_by_month(month_str)
            # 获取设备类型和计数
            device_results = self.db_speed_conn.get_device_types_and_counts(month_str)
            # 获取新用户统计信息
            new_users_results = self.db_speed_conn.get_new_users_in_month(month_str)
            # 新用户按设备类型分组
            new_users_by_os = {os_type: 0 for os_type in util.os_types}

            for row in new_users_results:
                original_os, email, new_users = row
                categorized_os = util.categorize_os(original_os)
                new_users_by_os[categorized_os] += new_users
            
            # 计算次月留存用户数
            retained_users_by_os = self.db_speed_conn.calculate_retention_of_next_month(month_str)

             # 处理results并插入到t_user_report_monthly中

            if registered_results:
                stat_month = registered_results[0][0]  # 获取正确的stat_month
                user_count = sum([row[2] for row in registered_results])  # 计算总用户数
            else:
                stat_month = month_str
                user_count = 0
            
            # 设备类型和计数
            #device_os_counts = {row[0]: row[1] for row in device_results}

            for os in util.os_types:
                total_new_users = new_users_by_os.get(os, 0)
                # total_device_count = device_os_counts.get(os, 0)
                retained_users_count = len(retained_users_by_os.get(os, set()))
                self.db_report_conn.insert_into_report_monthly(stat_month, os, user_count, total_new_users, retained_users_count)

            # 移动到下一个月
            current_month += timedelta(days=32)
            current_month = current_month.replace(day=1)

    def report_ad_daily(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        # print(self.date) #昨天的日期
        # 获取广告列表
        ad_list_data = self.db_speed_conn.get_ad_list()
        # 创建广告ID到广告名称的映射
        ad_map = {ad['id']: ad['name'] for ad in ad_list_data}

        # 获取曝光统计数据
        exposure_stats = self.db_speed_conn.get_ad_exposure_statistics(self.date, 'view')
        exposure_dict = {stat['ad_name']: stat['ad_count'] for stat in exposure_stats}
        
        # 获取点击统计数据
        click_stats = self.db_speed_conn.get_ad_exposure_statistics(self.date, 'click')
        click_dict = {stat['ad_name']: stat['ad_count'] for stat in click_stats}
        
        # 获取完播数
        ad_gift_user = self.db_speed_conn.get_ad_gift_user(self.date)
        gift_dict = {stat['ad_name']: stat['total_count'] for stat in ad_gift_user}

        # print(ad_gift_user)
        # 前一天日期，统计前一天数据
        date_int = int(self.date.replace('-', ''))
        
        # # 合并统计数据
        for ad_id, ad_name in ad_map.items():
            exposure = exposure_dict.get(ad_name, 0)
            clicks = click_dict.get(ad_name, 0) 
            rewards = gift_dict.get(ad_name, 0) 
            # print(ad_id, ad_name, date_int, exposure, clicks, rewards)
            self.db_report_conn.insert_into_report_ad_daily(ad_id, ad_name, date_int, exposure, clicks, rewards)

    def report_pay_daily(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        # print(self.date) #昨天的日期
        payment_channels = self.db_speed_conn.get_payment_channel()
        # print(payment_channel)

        # 获取每日付费数据
        pay_orders = self.db_speed_conn.get_pay_order(self.date)
        # print(pay_order)

        # 初始化一个字典用于存储每个支付渠道的总金额
        total_by_channel = {channel: Decimal('0.00') for channel in payment_channels}

        for pay_order in pay_orders:
            channel = pay_order['payment_channel_id']
            # 确保支付渠道在字典中
            if channel in total_by_channel:
                # 将字符串形式的实际支付金额转换为Decimal类型并累加
                total_by_channel[channel] += Decimal(pay_order['order_reality_amount'])
        
        # 格式化结果
        formatted_output = {channel: str(amount) for channel, amount in total_by_channel.items()}
        # print(formatted_output)

        date_int = int(self.date.replace('-', ''))
        current_time = datetime.now().strftime('%Y-%m-%d %H:%M:%S')

        for channel, amount in formatted_output.items():
            # print(date_int, channel, amount)
            self.db_report_conn.insert_into_report_pay_daily(date_int, channel, amount, current_time)

    #根据日期获取注册用户数和付费用户数
    def report_channel_registration_pay_daily(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        date_init = self.date.replace('-', '')
        # print(date_init,self.month_start_time)
        """ 获取channel列表 """
        rows = self.db_speed_conn.query_channel_list(self.end_time)
        for row in rows:
            channel = row["channel"]
            """当前推广渠道新增用户数量"""
            new_users = self.db_speed_conn.count_channel_user_by_create_time(channel, self.start_time, self.end_time)
            """当前推广渠道日活用户数量"""
            daily_active_users = self.db_speed_conn.count_channel_active_users(channel, self.start_time, self.end_time)
            """当前推广渠道月活用户数量"""
            monthly_active_users = self.db_speed_conn.count_channel_active_users(channel, self.month_start_time, self.end_time)
            """当前推广渠道付费用户数量"""
            total_recharge_users = self.db_speed_conn.count_total_number_of_recharges_by_create_time(channel, self.start_time, self.end_time)
            """当前推广渠道付费金额数量"""
            total_recharge_amount = self.db_speed_conn.count_total_recharge_amount_by_create_time(channel, self.start_time, self.end_time)
            self.db_report_conn.insert_into_report_channel_registration_pay_daily(date_init, channel, new_users, daily_active_users, monthly_active_users, total_recharge_users, total_recharge_amount)
        # print(data)
    
    # 根据日期获取推广渠道的留存用户数
    ###############################
    #因为是需要统计次日留存，七日留存，十五日留存，三十日留存，
    #所以只能从30日之前统计，所以需要获取30天前的数据
    ###############################
    def report_channel_retaind_daily(self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        date_init = self.date.replace('-','')
        #30天前的日期
        thirty_days_ago = util.get_previous_days(30) #获取30天前的日期 比如2025-01-13
        t = util.Time(thirty_days_ago) #获取30天前的日期
        date = t.date.replace('-','') #当天的日期比如2025-2-12
        st = t.get_start_time() #当天开始时间 比如2025-2-12 17:40:51
        et = t.get_end_time() #当天结束时间 比如2025-2-12 17:40:54
        # print(thirty_days_ago)
        # print(date, st, et)
        #30天前的次日日期
        next_day_date = (datetime.strptime(thirty_days_ago, '%Y-%m-%d') + timedelta(days=1)).strftime('%Y-%m-%d')
        t2 = util.Time(next_day_date) #获取30天前次日的日期
        date2 = t2.date.replace('-','') #当天的日期比如2025-2-12
        st2 = t2.get_start_time() #当天开始时间 比如2025-2-12 17:40:51
        et2 = t2.get_end_time() #当天结束时间 比如2025-2-12 17:40:54
        # print(next_day_date)
        # print(date2, st2, et2)
        seven_day_date = (datetime.strptime(thirty_days_ago, '%Y-%m-%d') + timedelta(days=7)).strftime('%Y-%m-%d')
        t7 = util.Time(seven_day_date) #获取30天前7日的日期
        date7 = t7.date.replace('-','') #当天的日期比如2025-2-12
        # st7 = t7.get_start_time() #当天开始时间 比如2025-2-12 17:40:51
        et7 = t7.get_end_time() #当天结束时间 比如2025-2-12 17:40:54
        fifteen_day_date = (datetime.strptime(thirty_days_ago, '%Y-%m-%d') + timedelta(days=15)).strftime('%Y-%m-%d')
        t15 = util.Time(fifteen_day_date) #获取30天前15日的日期
        date15 = t15.date.replace('-','') #当天的日期比如2025-2-12
        # st15 = t15.get_start_time() #当天开始时间 比如2025-2-12 17:40:51
        et15 = t15.get_end_time() #当天结束时间 比如2025-2-12 17:40:54
        # print(thirty_days_ago,seven_day_date,fifteen_day_date)
        """ 获取channel列表 """
        rows = self.db_speed_conn.query_channel_list(et)
        for row in rows:
            channel = row["channel"]
            # print("channel=>",channel)
            """当前推广渠道新增用户数量"""
            new_users = self.db_speed_conn.count_channel_user_by_create_time(channel, st, et)
            """当前推广渠道新增用户次日留存数量"""
            day_2_retained = self.db_speed_conn.count_channel_active_users_by_at(channel, st, et, date2, date2, st2, et2)
            """当前推广渠道新增用户7日留存数量"""
            day_7_retained = self.db_speed_conn.count_channel_active_users_by_at(channel, st, et, date2, date7, st2, et7)
            """当前推广渠道新增用户15日留存数量"""
            day_15_retained = self.db_speed_conn.count_channel_active_users_by_at(channel, st, et, date2, date15, st2, et15)
            """当前推广渠道新增用户30日留存数量"""
            day_30_retained = self.db_speed_conn.count_channel_active_users_by_at(channel, st, et, date2, date_init, st2, self.end_time)
            self.db_report_conn.insert_channel_retaind_daily(date,channel,new_users,day_2_retained,day_7_retained,day_15_retained,day_30_retained)

    def report_channel_recharge_by_month (self):
        logging.info(f"{'*' * 20}{sys._getframe().f_code.co_name}{'*' * 20}")
        last_month = util.get_previous_months(1) #获取当前月份的前1个月
        last_first_day, last_last_day = util.get_month_start_end(last_month)
        # print(last_first_day, last_last_day)
        mst = util.Time(last_first_day).get_start_time()
        met = util.Time(last_last_day).get_end_time()
        # print(mst, met)
        # 本月第一天
        first_day = util.get_month_date()
        last_day = util.get_yesterday_date()
        nmst = util.Time(first_day).get_start_time()
        nmet = util.Time(last_day).get_end_time()
        # print(nmst, nmet)
        """ 获取channel列表 """
        rows = self.db_speed_conn.query_channel_list(self.end_time)
        for row in rows:
            """当前推广渠道"""
            channel = row["channel"]
            """当前推广渠道充值用户数量"""
            total_recharge_users = self.db_speed_conn.count_channel_recharge_times(channel, mst, met)
            """当前推广渠道充值金额总数"""
            total_recharge_amount = self.db_speed_conn.count_total_recharge_amount_by_create_time(channel, mst, met)
            """当前推广渠道充值用户次月留存数量"""
            total_recharge_retained = self.db_speed_conn.count_channel_recharge_users_retained(channel, mst, met, first_day.replace('-',''), last_day.replace('-',''), nmst, nmet)
            """当前推广渠道充值用户次月续费人数数量"""
            total_renewals_users = self.db_speed_conn.count_channel_renewal_users(channel, mst, met, nmst, nmet)
            """当前推广渠道充值用户次月续费金额总数"""
            total_renewals_users_amount = self.db_speed_conn.count_channel_renewal_amount(channel, mst, met, nmst, nmet)
            self.db_report_conn.insert_channel_recharge_monthly(last_month.replace('-',''), channel, total_recharge_users, total_recharge_amount, total_recharge_retained, total_renewals_users, total_renewals_users_amount)