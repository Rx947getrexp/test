# -*- coding: utf-8 -*-
import calendar
from datetime import date, datetime, timedelta
from dateutil.relativedelta import relativedelta
import geoip2.database


def get_month_date():
    # 获取今天的日期
    today = datetime.now()
    # 判断今天是否是1号
    # 获取今天的日期字符串
    #today_str = "2024-09-12 00:00:00"
    # 将字符串转换为 datetime 对象
    #today = datetime.strptime(today_str, "%Y-%m-%d %H:%M:%S")
    if today.day == 1:
        # 获取上个月的第一天
        first_day_last_month = (today.replace(day=1) - timedelta(days=1)).replace(day=1)
        return first_day_last_month.strftime("%Y-%m-%d")
    else:
        # 如果不是1号，获取当月的第一天
        first_day_this_month = today.replace(day=1)
        return first_day_this_month.strftime("%Y-%m-%d")
def get_yesterday_date():
    # 获取今天的日期
    today = datetime.now()
    # 计算昨天的日期
    yesterday = today - timedelta(days=1)
    # 将昨天的日期格式化为 "2024-01-01" 格式
    return yesterday.strftime('%Y-%m-%d')
    #DATE=input("请输入：")
    #return DATE
    #return "2024-09-10"
def get_previous_months(months_back):
    """
    获取当前日期之前的指定月数的月份，格式为 "YYYY-MM"。
    :param months_back: int, 要回溯的月数
    :return: str, 指定月数之前的月份，格式为 "YYYY-MM"
    """
    # 计算过去的月份
    previous_month = datetime.now() - relativedelta(months=months_back)
    # 格式化输出为 "YYYY-MM"
    formatted_date = previous_month.strftime("%Y-%m")
    return formatted_date

def get_previous_days(days=15):
    # 获取今天的日期
    today = datetime.now()
    # 计算多少天前的日期
    previous_day = today - timedelta(days)
    # 将多少天前的日期格式化为 "2024-01-01" 格式
    return previous_day.strftime('%Y-%m-%d')

def time_format(s):
    return datetime.strptime(s, "%Y-%m-%d %H:%M:%S")

# 设备型号映射表
device_mapping = {
    'Android': ['Android'],
    'iPhone': ['iPhone', 'iOS'],
    'Mac': ['Mac', 'macOS', 'Mac OS'],
    'Windows': ['Windows'],
    'Others': []  # 任何未匹配的都归为 Others
}
os_types = list(device_mapping.keys())
def categorize_os(original_os):
    if 'Mac OS' in original_os:
        if 'iPhone' in original_os:
            return 'iPhone'
        else:
            return 'Mac'
    elif 'Android' in original_os:
        return 'Android'
    elif 'Windows' in original_os:
        return 'Windows'
    else:
        return 'Others'

def get_month_start_end(date_str):
    """
    根据输入的 YYYY-MM 格式字符串，返回当月第一天和最后一天的日期字符串（YYYY-MM-DD）。
    
    :param date_str: str, 格式为 'YYYY-MM'
    :return: tuple (first_day, last_day) -> ('YYYY-MM-DD', 'YYYY-MM-DD')
    """
    year, month = map(int, date_str.split('-'))  # 解析年份和月份
    
    # 获取当月第一天
    first_day = date(year, month, 1)
    
    # 获取当月最后一天
    last_day = date(year, month, calendar.monthrange(year, month)[1])
    
    return first_day.strftime('%Y-%m-%d'), last_day.strftime('%Y-%m-%d')

class Time:
    def __init__(self, date):
        self.class_name = "util.Time"
        self.date = date

    def get_start_time(self):
        return self.date + " 00:00:00"

    def get_end_time(self):
        return self.date + " 23:59:59"


class IpSearch:
    def __init__(self):
        self.reader = geoip2.database.Reader('/shell/report/GeoLite2-City.mmdb')

    def get_location(self, ip_address):
        response = self.reader.city(ip_address)
        Country_IsoCode = response.country.iso_code
        if (Country_IsoCode == None):
            Country_IsoCode = "None"
        return Country_IsoCode
