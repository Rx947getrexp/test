# -*- coding: utf-8 -*-
from datetime import datetime, timedelta
import logging
import sys
import pymysql
from common import load_config
import util

def mysql_connect(db):
    """连接DB"""
    return pymysql.connect(host=db["host"], port=db["port"], user=db["user"], passwd=db["pswd"], db=db["db"],charset=db["charset"])


def mysql_query_db(_conn, sql):
    """读取DB数据"""
    logging.info(sql)
    cursor = _conn.cursor(cursor=pymysql.cursors.DictCursor)
    cursor.execute(sql)
    rows = cursor.fetchall()
    cursor.close()
    return rows


def mysql_execute(_conn, sql):
    """执行sql"""
    logging.info(sql)
    cursor = _conn.cursor(cursor=pymysql.cursors.DictCursor)
    logging.info(cursor.execute(sql))
    # cursor.execute(sql)
    _conn.commit()
    cursor.close()

class Speed:
    def __init__(self):
        self.config = load_config("/shell/report/config.yaml")
        self.conn = mysql_connect(self.config["speed-db"])
        self.exchange_rate_usd = 90.23  # usdt汇率 1u=90.23rub
        self.exchange_rate_wmz = 65  # wmz汇率 1u=65rub
        self.country = "193.233.48.70", "46.17.41.7", "45.147.201.21", "45.147.200.112", "46.17.44.132", "92.118.112.89", "5.181.3.143", "207.90.237.91", "91.149.218.194", "62.133.60.81", "62.133.63.237", "213.159.68.106", "103.198.203.11", "110.42.42.229", "185.39.207.20", "92.118.112.133", "212.18.104.23", "147.45.178.51", "185.39.207.104", "193.124.41.88"

    def close_connection(self):
        if self.conn:
            self.conn.close()

    def convert_to_rub(self, amount, currency):
        exchange_rates = {
            'USD': self.exchange_rate_usd,
            'WMZ': self.exchange_rate_wmz
        }
        return float(amount) * exchange_rates.get(currency, 1)

    # def get_user_ip_list(self):
    #     sql = """SELECT t1.id, t1.email, t2.ip, t2.updated_at FROM speed.t_user t1 JOIN (SELECT user_id, MAX(updated_at) AS max_updated_at FROM user_logs GROUP BY user_id
    # ) AS latest_logs ON t1.id = latest_logs.user_id JOIN user_logs t2 ON t1.id = t2.user_id AND t2.updated_at = latest_logs.max_updated_at;"""
    #     rows = mysql_query_db(self.conn, sql)
    #     dic = {}
    #     for row in rows:
    #         dic[row["email"]] = {
    #             "id": row["id"],
    #             "ip": row["ip"],
    #         }
    #     return dic

    def get_users(self):
        sql = """select id, email, channel, last_login_ip, last_login_country, created_at from speed.t_user;"""
        rows = mysql_query_db(self.conn, sql)
        dic = {}
        for row in rows:
            dic[row["email"]] = {
                "id": row["id"],
                "channel": row["channel"],
                "last_login_ip": row["last_login_ip"],
                "country": row['last_login_country'],
                "register_date": row['created_at']
            }
        return dic

    def count_user_by_create_time(self, channel_id, st, et):
        sql = """select count(*) as cnt from speed.t_user where channel_id= %d and created_at >= '%s' and created_at <= '%s';""" % (
            channel_id, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_node_by_create_time(self, ip, st, et):
        sql = """SELECT count(DISTINCT email) as cnt FROM speed_collector.t_v2ray_user_traffic WHERE ip='%s' AND email IN (select email as cnt from speed.t_user where created_at >= '%s' and created_at <= '%s') AND created_at >= '%s' and created_at <= '%s';""" % (
            ip, st, et, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_online(self, channel_id, date,st, et):
        # sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where date = '%s' and email in (select email from speed.t_user where channel_id='%d' and created_at <= '%s');""" % (date.replace("-", ""), channel_id, et)
        sql = """SELECT COUNT(DISTINCT email) as cnt FROM (SELECT email FROM (SELECT email FROM speed_report.t_user_op_log WHERE (content LIKE '%点击连接%' OR content LIKE '%开始连接%') AND created_at >= '{}' AND result = 'success' AND created_at <= '{}' AND email IN (SELECT email FROM speed.t_user) UNION SELECT email FROM speed_collector.t_v2ray_user_traffic WHERE date = '{}' AND email IN (SELECT email FROM speed.t_user WHERE channel_id = {} AND created_at <= '{}')) AS combined_results) AS distinct_emails;""".format(st, et, date.replace("-", ""), channel_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]
    def count_user_month_online(self, channel_id, st, et):
        # sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where email in (select email from speed.t_user where channel_id='%d' and created_at <= '%s') and created_at >= '%s' and created_at <= '%s';""" % (channel_id,et, st, et)
        sql = """SELECT COUNT(DISTINCT email) as cnt FROM (SELECT email FROM (SELECT email FROM speed_report.t_user_op_log WHERE (content LIKE '%点击连接%' OR content LIKE '%开始连接%') AND created_at >= '{}' AND result = 'success' AND created_at <= '{}' AND email IN (SELECT email FROM speed.t_user) UNION SELECT email FROM speed_collector.t_v2ray_user_traffic WHERE email IN (SELECT email FROM speed.t_user WHERE channel_id = '{}' AND created_at <= '{}') AND created_at >= '{}' and created_at <= '{}') AS combined_results ) AS distinct_emails;""".format(st, et, channel_id, et,st,et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_node_online(self, ip, date, et):
        sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where date = '%s' and ip = '%s' and email in (select email from speed.t_user WHERE created_at <= '%s');""" % (
            date.replace("-", ""), ip, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_user(self, channel_id, et):
        sql = """select count(*) as cnt from t_user where channel_id = %d and created_at <= '%s';""" % (channel_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_node(self, ip, et):
        sql = """select count(DISTINCT email) as cnt from speed_collector.t_v2ray_user_traffic where ip = '%s' and created_at <= '%s';""" % (
            ip, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def query_channel_id_list(self, et):
        sql = """select distinct channel_id as channel_id from speed.t_user where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_channel_list(self, et):
        sql = """select distinct channel as channel from speed.t_user where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_node_traffic_list(self, et):
        sql = """select distinct ip as ip from speed_collector.t_v2ray_user_traffic where created_at <= '%s' and ip in %s;""" % (
            et, self.country)
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_user_traffic_list(self, date):
        sql = """select distinct email as email from speed_collector.t_v2ray_user_traffic where date=%s;""" % date.replace("-", "")
        return mysql_query_db(self.conn, sql)

    def query_node_user_traffic_list(self, date):
        sql = """select email as email,ip,uplink,downlink from speed_collector.t_v2ray_user_traffic where date=%s;""" % date.replace("-", "")
        return mysql_query_db(self.conn, sql)

    def query_user_traffic_log_list(self, email, st, et):
        sql = """select ip, date_time, uplink, downlink from speed_collector.t_v2ray_user_traffic_log where email = '%s' and date_time>='%s' and date_time<='%s' order by date_time asc;""" % (
            email, st, et)
        return mysql_query_db(self.conn, sql)

    def query_recharge_channel_list(self, et):
        sql = """SELECT distinct t2.channel as channel FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email=t2.email  WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) AND t2.channel !=''and t1.created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def count_total_channel_user(self, channel, et):
        sql = """select count(*) as cnt from speed.t_user where channel = '%s' and created_at <= '%s';""" % (channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_recharge_total_channel_user(self, channel, goods_id, currency, et):
        sql = """select count(*) as cnt FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and t1.goods_id = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.currency='%s' and t1.created_at <= '%s';""" % (
            channel, goods_id, currency, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_channel_user_by_create_time(self, channel, st, et):
        sql = """select count(*) as cnt from speed.t_user where channel= '%s' and created_at >= '%s' and created_at <= '%s';""" % (
            channel, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_recharge_channel_user_by_create_time(self, channel, goods_id, currency, st, et):
        sql = """select count(*) as cnt FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and t1.goods_id = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.currency='%s' and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (
            channel, goods_id, currency, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_channel_user_online(self, channel, date, st,et):
        # sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where date = '%s' and email in (select email from speed.t_user where channel='%s' and created_at <= '%s');""" % (date.replace("-", ""), channel, et)
        sql = "SELECT COUNT(DISTINCT email) as cnt FROM (SELECT email FROM (SELECT DISTINCT email FROM speed_report.t_user_op_log WHERE (content LIKE '%点击连接%' OR content LIKE '%开始连接%') AND created_at >= '{}' AND result = 'success' AND created_at <= '{}' AND email IN (SELECT email FROM speed.t_user WHERE channel='{}')) AS derived_table1 UNION SELECT email FROM (SELECT email FROM speed_collector.t_v2ray_user_traffic WHERE date = '{}' AND email IN (SELECT email FROM speed.t_user WHERE channel='{}' AND created_at <= '{}')) AS derived_table2) AS distinct_emails;".format(st, et, channel, date.replace("-", ""), channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_channel_user_by_month(self, channel, st, et):
        sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where email in (select email from speed.t_user where channel='%s' and created_at >= '%s' and created_at <= '%s') and created_at >= '%s' and created_at <= '%s';""" % (channel, st, et,st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def query_goods_id_list(self):
        sql = """select id from speed.t_goods;"""
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_currency_list(self):
        sql = """select DISTINCT currency from speed.t_pay_order;"""
        rows = mysql_query_db(self.conn, sql)
        return rows

    def count_total_user_recharge(self, goods_id, et):
        sql = """select count(distinct user_id) as cnt from speed.t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at <= '%s';""" % (
            goods_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_recharge_by_create_time(self, goods_id, st, et):
        sql = """select count(distinct user_id) as cnt from speed.t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at >= '%s' and created_at <= '%s';""" % (
            goods_id, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_user_recharge_times(self, goods_id, et):
        sql = """select count(user_id) as cnt from speed.t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at <= '%s';""" % (
            goods_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_recharge_times_by_create_time(self, goods_id, st, et):
        sql = """select count(user_id) as cnt from speed.t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at >= '%s' and created_at <= '%s';""" % (
            goods_id, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_number_of_recharges(self, channel, et):
        '''
        充值总人数
        '''
        sql = """select count(*) as cnt FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_number_of_recharges_by_create_time(self, channel, st, et):
        sql = """select count(*) as cnt FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (channel,st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_recharge_amount(self, channel, et):
        '''
        充值总金额
        '''
        sql = """select t1.currency, t1.order_reality_amount,t1.status FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (
            channel, et)
        rows = mysql_query_db(self.conn, sql)
        total_recharge = 0
        for row in rows:
            status = row['status']
            currency = row['currency']
            amount = row['order_reality_amount']
            if status == "paid":
                amount_in_rub = self.convert_to_rub(amount, currency)
            elif status == "admin-confirm-passed":
                amount_in_rub = amount
            else:
                continue
            total_recharge += float(amount_in_rub)
        return total_recharge

    def count_total_recharge_amount_by_create_time(self, channel, st, et):
        '''
        新增充值总金额
        '''
        sql = """select t1.currency, t1.order_reality_amount,t1.status FROM speed.t_pay_order t1 JOIN speed.t_user t2 on t1.email = t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (
            channel, st, et)
        rows = mysql_query_db(self.conn, sql)
        total_recharge = 0
        for row in rows:
            status = row['status']
            currency = row['currency']
            amount = row['order_reality_amount']
            if status == "paid":
                amount_in_rub = self.convert_to_rub(amount, currency)
            elif status == "admin-confirm-passed":
                amount_in_rub = amount
            else:
                continue
            total_recharge += float(amount_in_rub)
        return total_recharge

    def query_device_list(self, et):
        sql = """select distinct os as device from speed.t_dev where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def count_total_device_user(self, device, et):
        sql = """SELECT count(DISTINCT u.id) as cnt FROM speed.t_user u JOIN speed.t_user_device d ON u.id = d.user_id WHERE d.os= '%s' and u.created_at <= '%s';""" % (device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_device_user_by_create_time(self, device, st, et):
        sql = """SELECT count(DISTINCT u.id) as cnt FROM speed.t_user u JOIN speed.t_user_device d ON u.id = d.user_id WHERE d.os= '%s' AND u.created_at >= '%s' and u.created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_device_user_online(self, device, date, st,et):
        sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where date = '%s' and email in (SELECT u.email FROM speed.t_user u JOIN speed.t_user_device d ON u.id = d.user_id WHERE d.os= '%s' AND u.created_at >= '%s' and u.created_at <= '%s');""" % (
            date.replace("-", ""), device, st,et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]
    def count_device_retained(self, device, date, st,et):
        sql = """select count(distinct email) as cnt from speed_collector.t_v2ray_user_traffic where date >= '%s' and email in (SELECT u.email FROM speed.t_user u JOIN speed.t_user_device d ON u.id = d.user_id WHERE d.os= '%s' and u.created_at >= '%s' and u.created_at <= '%s');""" % (
            date.replace("-", ""), device,st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]
    def count_total_number_of_device_recharges(self, et):
        sql = """SELECT COALESCE(t1.device_type, 'unknown') AS device_type,COUNT(t1.user_id) AS cnt FROM speed.t_pay_order t1 JOIN speed.t_user t2 ON t1.user_id = t2.id WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ', 'RUB') AND t1.status = 'paid')) and t1.created_at <= '%s' GROUP BY t1.device_type;""" % (et)
        rows = mysql_query_db(self.conn, sql)
        return {row["device_type"]: row["cnt"] for row in rows}

    def count_total_device_recharge_amount(self, et):
        sql = """SELECT COALESCE(t1.device_type, 'unknown') AS device_type,t1.currency, t1.order_reality_amount,t1.status FROM speed.t_pay_order t1 JOIN speed.t_user t2 ON t1.user_id = t2.id WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ', 'RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (et)
        rows = mysql_query_db(self.conn, sql)
        device_recharge_totals = {}
        for row in rows:
            device_type = row['device_type']
            status = row['status']
            currency = row['currency']
            amount = row['order_reality_amount']
            if status == "paid":
                amount_in_rub = self.convert_to_rub(amount, currency)
            elif status == "admin-confirm-passed":
                amount_in_rub = float(amount)
            else:
                continue
            if device_type not in device_recharge_totals:
                device_recharge_totals[device_type] = 0
            device_recharge_totals[device_type] += amount_in_rub
        # 返回各个设备类型的总金额
        return device_recharge_totals

    def count_total_device_recharge_amount_by_create_time(self, st, et):
        sql = """SELECT COALESCE(t1.device_type, 'unknown') AS device_type,t1.currency, t1.order_reality_amount,t1.status FROM speed.t_pay_order t1 JOIN speed.t_user t2 ON t1.user_id = t2.id WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ', 'RUB') AND t1.status = 'paid')) and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (st, et)
        rows = mysql_query_db(self.conn, sql)
        device_recharge_totals = {}
        for row in rows:
            device_type = row['device_type']
            status = row['status']
            currency = row['currency']
            amount = row['order_reality_amount']
            if status == "paid":
                amount_in_rub = self.convert_to_rub(amount, currency)
            elif status == "admin-confirm-passed":
                amount_in_rub = float(amount)
            else:
                continue
            if device_type not in device_recharge_totals:
                device_recharge_totals[device_type] = 0
            device_recharge_totals[device_type] += amount_in_rub
        # 返回各个设备类型的总金额
        return device_recharge_totals
    
    ################################
    def query_registered_users_by_day(self, day_str):
        sql = """SELECT DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') AS stat_day, tu.email, COUNT(*) AS user_count FROM speed.t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') = '%s' GROUP BY tu.email, stat_day;""" % (day_str)
        rows = mysql_query_db(self.conn, sql)
        return [(row['stat_day'], row['email'], row['user_count']) for row in rows]
    
    def query_registered_emails_by_day(self, day_str):
        sql = """SELECT tu.email FROM speed.t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m-%%d') = '%s'""" % (day_str)
        results = mysql_query_db(self.conn, sql)
        return [row['email'] for row in results]
    
    def query_registered_device_emails(self, registered_emails):
        if not registered_emails:
            return {}
        # 将 registered_emails 转换为元组，并将其转换为逗号分隔的字符串
        email_list_str = ', '.join([f"'{email}'" for email in registered_emails])
        sql = f"""SELECT tud.os, tu.email FROM t_user_device tud JOIN t_user tu ON tud.user_id = tu.id WHERE tu.email IN ({email_list_str});"""
        results = mysql_query_db(self.conn, sql)
        return {item['email']: item['os'] for item in results}

    def query_active_emails_in_next_days(self, date, days):
        date_obj = datetime.strptime(date, '%Y-%m-%d')
        start_date = date_obj + timedelta(days=1)
        end_date = (date_obj + timedelta(days=days + 1)) - timedelta(seconds=1)
        sql = """SELECT tut.email FROM speed_collector.t_v2ray_user_traffic tut WHERE tut.date >= %s AND tut.date <= %s""" % (int(start_date.strftime('%Y%m%d')), int(end_date.strftime('%Y%m%d')))
        results = mysql_query_db(self.conn, sql)
        return [row['email'] for row in results]

    def query_retention_of_next_days(self, day_str, days):
        registered_emails = self.query_registered_emails_by_day(day_str)
        active_emails = set(self.query_active_emails_in_next_days(day_str, days))
        if not registered_emails:
            return {os_type: set() for os_type in util.os_types}
        registered_device_emails = self.query_registered_device_emails(registered_emails)
        retained_users_by_os = {os_type: set() for os_type in util.os_types}
        for email in registered_emails:
            if email in active_emails:
                original_os = registered_device_emails.get(email, 'Others')
                categorized_os = util.categorize_os(original_os)
                retained_users_by_os[categorized_os].add(email)
        return retained_users_by_os
    
    ################################
    def get_registered_users_by_month(self, month):
        sql = """SELECT DATE_FORMAT(tu.created_at, '%%Y-%%m') AS stat_month, tu.email, COUNT(*) AS user_count FROM speed.t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = '%s' GROUP BY tu.email, stat_month;""" % (month)
        rows = mysql_query_db(self.conn, sql)
        return [(row['stat_month'], row['email'], row['user_count']) for row in rows]

    def get_device_types_and_counts(self, month):
        query = """SELECT tud.os AS os, COUNT(*) AS device_count FROM speed.t_user_device tud JOIN t_user tu ON tud.user_id = tu.id WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = '%s' GROUP BY tud.os;""" % (month,)
        res = mysql_query_db(self.conn, query)
        results = [(row['os'], row['device_count']) for row in res]
        categorized_counts = {os_type: 0 for os_type in util.device_mapping.keys()}
        for row in results:
            original_os = row[0]
            categorized_os = util.categorize_os(original_os)
            categorized_counts[categorized_os] += row[1]
        return [(os_type, count) for os_type, count in categorized_counts.items()]
    
    def get_new_users_in_month(self, month):
        sql = """SELECT tud.os, tu.email, COUNT(*) AS new_users FROM speed.t_user tu JOIN speed.t_user_device tud ON tu.id = tud.user_id WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = '%s' GROUP BY tud.os, tu.email;""" % (month,)
        rows = mysql_query_db(self.conn, sql)
        return  [(row['os'], row['email'], row['new_users']) for row in rows]

    def get_registered_emails_by_month(self, month):
        sql = """SELECT tu.email FROM speed.t_user tu WHERE DATE_FORMAT(tu.created_at, '%%Y-%%m') = '%s'""" % (month,)
        results = mysql_query_db(self.conn, sql)
        return [row['email'] for row in results]
    
    def get_active_emails_in_next_month(self, month):
        month_date = datetime.strptime(month, '%Y-%m')
        next_month_date = month_date + timedelta(days=32)
        #next_month = next_month_date.strftime('%Y-%m')
        next_month_start = next_month_date.replace(day=1)
        next_month_end = (next_month_start + timedelta(days=32)).replace(day=1) - timedelta(days=1)
        sql = """SELECT tut.email FROM speed_collector.t_v2ray_user_traffic tut WHERE tut.date >= %s AND tut.date <= %s""" % (int(next_month_start.strftime('%Y%m%d')), int(next_month_end.strftime('%Y%m%d')))
        results = mysql_query_db(self.conn, sql)
        return [row['email'] for row in results]
    
    def calculate_retention_of_next_month(self, month):
        registered_emails = self.get_registered_emails_by_month(month)
        active_emails = set(self.get_active_emails_in_next_month(month))  # 转换成集合
        if not registered_emails:
            return {os_type: set() for os_type in util.os_types}  # 如果没有注册用户，直接返回空结果
        registered_device_emails = {}
        # 将 registered_emails 转换为元组，并将其转换为逗号分隔的字符串
        in_registered_emails = ', '.join([f"'{email}'" for email in registered_emails])
        sql = f"""SELECT tud.os, tu.email FROM speed.t_user_device tud JOIN speed.t_user tu ON tud.user_id = tu.id WHERE tu.email IN ({in_registered_emails});"""
        res = mysql_query_db(self.conn, sql)
        results = [(row['os'], row['email']) for row in res]
        registered_device_emails = {email: os for os, email in results}
        retained_users_by_os = {os_type: set() for os_type in util.os_types}
        for email in registered_emails:  # 直接遍历列表中的每个邮箱地址
            if email in active_emails:
                original_os = registered_device_emails.get(email, 'Others')
                categorized_os = util.categorize_os(original_os)
                retained_users_by_os[categorized_os].add(email)
        return retained_users_by_os

class SpeedReport:
    def __init__(self):
        self.config = load_config("/shell/report/config.yaml")
        self.conn = mysql_connect(self.config["report-speed-db"])
        self.data_name = {
            "193.233.48.70": "俄罗斯1",
            "46.17.41.7": "俄罗斯2",
            "45.147.201.21": "俄罗斯3",
            "45.147.200.112": "俄罗斯4",
            "46.17.44.132": "俄罗斯5",
            "92.118.112.89": "美国1",
            "5.181.3.143": "美国2",
            "207.90.237.91": "美国3",
            "92.118.112.133": "美国4",
            "212.18.104.23": "美国5",
            "91.149.218.194": "法国1",
            "62.133.60.81": "德国1",
            "147.45.178.51": "德国2",
            "193.124.41.88": "波兰1",
            "62.133.63.237": "土耳其1",
            "213.159.68.106": "芬兰1",
            "103.198.203.11": "香港1",
            "110.42.42.229": "中国1",
            "185.39.207.20": "希腊1",
            "185.39.207.104": "希腊2"
        }

    def close_connection(self):
        if self.conn:
            self.conn.close()

    def query_devices_list(self, et):
        sql = """SELECT DISTINCT device_type as device FROM speed_report.t_user_op_log  where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def get_total_clicks(self, device, et):
        sql = """SELECT count(email) as cnt from speed_report.t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_yesterday_day_clicks(self, device, st, et):
        sql = """SELECT count(email) as cnt from speed_report.t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_weekly_clicks(self, device, st, et):
        sql = """SELECT count(email) as cnt from speed_report.t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_total_users_clicked(self, device, et):
        sql = """SELECT count(distinct email) as cnt from speed_report.t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_yesterday_day_users_clicked(self, device, st, et):
        sql = """SELECT count(distinct email) as cnt from speed_report.t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_weekly_users_clicked(self, device, st, et):
        sql = """SELECT count(distinct email) as cnt from speed_report.t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def insert_daily_user(self, date, data):
        logging.info(data)
        mysql_execute(self.conn, "delete from speed_report.t_user_report_day where date=%s" % date.replace("-", ""))
        for channel_id in data.keys():
            user = data[channel_id]
            sql = """insert into t_user_report_day set date=%s, channel_id=%d, total=%d, new=%d, retained=%d,month_retained=%d, created_at=now();""" % (date.replace("-", ""), channel_id, user["total_cnt"], user["new_cnt"], user["retained_cnt"], user["month_retained_cnt"])
            mysql_execute(self.conn, sql)

    def insert_online_user_day(self, date, items):
        mysql_execute(self.conn, "delete from speed_report.t_user_online_day where date=%s" % date.replace("-", ""))
        for item in items:
            sql = """insert into speed_report.t_user_online_day set date='{}',email='{}',channel='{}',online_duration={},uplink={},downlink={},last_login_country='{}',created_at=now();""".format(
                date.replace("-", ""), item["email"], item["channel"], item["online_duration"], item["uplink"],
                item["downlink"], item["country"])
            mysql_execute(self.conn, sql)

    def insert_daily_channel_user(self, date, data):
        mysql_execute(self.conn, "delete from speed_report.t_user_channel_day where date=%s" % date.replace("-", ""))
        for channel in data.keys():
            user = data[channel]
            channel = channel if channel else '官网'
            sql = """insert into speed_report.t_user_channel_day set date='{}', channel='{}', total={}, new={},retained={}, created_at=now(),total_recharge={},total_recharge_money={},new_recharge_money={};""".format(
                date.replace("-", ""), channel, user["total_cnt"], user["new_cnt"], user["retained_cnt"],
                user["total_recharge"], user["total_recharge_amount"], user["new_recharge_amount"])
            mysql_execute(self.conn, sql)

    def insert_daily_node(self, date, data):
        mysql_execute(self.conn, "delete from speed_report.t_user_node_day where date=%s" % date.replace("-", ""))
        for ip in data.keys():
            user = data[ip]
            name = self.data_name.get(ip)
            if name:
                sql = """insert into speed_report.t_user_node_day set date='{}', ip='{}', total={},new={},retained={}, created_at=now();""".format(
                    date.replace("-", ""), name, user["total_cnt"], user["new_cnt"], user["retained_cnt"])
                mysql_execute(self.conn, sql)

    def insert_online_user_node_day(self, date, items):
        mysql_execute(self.conn, "delete from speed_report.t_user_node_online_day where date=%s" % date.replace("-", ""))
        for item in items:
            ip_name = self.data_name.get(item["node"])
            sql = """insert into speed_report.t_user_node_online_day set date='{}',email='{}',channel='{}',online_duration={},uplink={},downlink={},node='{}',register_date='{}',created_at=now();""".format(
                date.replace("-", ""), item["email"], item["channel"], item["online_duration"], item["uplink"],
                item["downlink"], ip_name, item["register_date"])
            mysql_execute(self.conn, sql)

    def insert_daily_user_recharge(self, date, data):
        logging.info(data)
        mysql_execute(self.conn, "delete from speed_report.t_user_recharge_report_day where date=%s" % date.replace("-", ""))
        for goods_id in data.keys():
            user = data[goods_id]
            sql = """insert into speed_report.t_user_recharge_report_day set date=%s, goods_id=%d, total=%d, new=%d,created_at=now();""" % (
                date.replace("-", ""), goods_id, user["total_cnt"], user["new_cnt"])
            mysql_execute(self.conn, sql)

    def insert_daily_user_recharge_times(self, date, data):
        logging.info(data)
        mysql_execute(self.conn, "delete from speed_report.t_user_recharge_times_report_day where date=%s" % date.replace("-", ""))
        for goods_id in data.keys():
            user = data[goods_id]
            sql = """insert into speed_report.t_user_recharge_times_report_day set date=%s, goods_id=%d, total=%d, new=%d,created_at=now();""" % (
                date.replace("-", ""), goods_id, user["total_cnt"], user["new_cnt"])
            mysql_execute(self.conn, sql)

    def insert_daily_channel_user_recharge(self, date, data):
        goods_name_dic = {
            1: "青铜会员30天",
            2: "青铜会员90天",
            3: "青铜会员180天",
            4: "青铜会员365天",
            5: "铂金会员30天",
            6: "铂金会员90天",
            7: "铂金会员180天",
            8: "铂金会员365天"
        }
        mysql_execute(self.conn, "delete from speed_report.t_user_channel_recharge_day where date=%s" % date.replace("-", ""))
        for channel, user_data in data.items():
            for goods_id, recharge_data in user_data.items():
                goods_name = goods_name_dic.get(goods_id, "未知套餐")
                # 获取USD和RUB的支付总次数和新增次数
                usd_total = recharge_data.get("USD", {}).get("total_cnt", 0)
                usd_new = recharge_data.get("USD", {}).get("new_cnt", 0)
                rub_total = recharge_data.get("RUB", {}).get("total_cnt", 0)
                rub_new = recharge_data.get("RUB", {}).get("new_cnt", 0)
                sql = f"""INSERT INTO speed_report.t_user_channel_recharge_day (date, channel, goods_name, usd_total, usd_new, rub_total, rub_new, created_at)
                                         VALUES ({date.replace("-", "")}, "{channel}", "{goods_name}", {usd_total}, {usd_new}, {rub_total}, {rub_new}, NOW())"""
                mysql_execute(self.conn, sql)

    def insert_daily_device_action(self, date, data):
        mysql_execute(self.conn, "delete from speed_report.t_user_device_action_day where date=%s" % date.replace("-", ""))
        for device_type in data.keys():
            device = data[device_type]
            sql = """insert into speed_report.t_user_device_action_day set date='%s', device='%s', total_clicks=%d, yesterday_day_clicks=%d, weekly_clicks=%d,total_users_clicked=%d,yesterday_day_users_clicked=%d,weekly_users_clicked=%d,created_at=now();""" % (
                date.replace("-", ""), device_type, device["total_clicks"], device["yesterday_day_clicks"],
                device["weekly_clicks"], device["total_users_clicked"], device["yesterday_day_users_clicked"],
                device["weekly_users_clicked"])
            mysql_execute(self.conn, sql)

    def insert_daily_device_user(self, date, data):
        # 设备类型映射表
        device_map = {'android': ['android'],'ios': ['iphone'],'mac': ['mac'],'win': ['win']}
        device_totals = {key: {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0, 'total_recharge_amount': 0,'new_recharge_amount': 0} for key in device_map.keys()}
        device_totals['other_device'] = {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0,'total_recharge_amount': 0, 'new_recharge_amount': 0}
        def update_device_totals(device_type, user_data):
            device_totals[device_type]['total_cnt'] += user_data.get("total_cnt", 0)
            device_totals[device_type]['new_cnt'] += user_data.get("new_cnt", 0)
            device_totals[device_type]['retained_cnt'] += user_data.get("retained_cnt", 0)
            device_totals[device_type]['total_recharge'] += user_data.get("total_recharge", 0)
            device_totals[device_type]['total_recharge_amount'] += user_data.get("total_recharge_amount", 0)
            device_totals[device_type]['new_recharge_amount'] += user_data.get("new_recharge_amount", 0)
        # 遍历数据，统计每种设备类型的数据
        for device, user in data.items():
            device_lower = device.lower()
            # 查找设备类型
            matched = False
            for device_type, patterns in device_map.items():
                if any(pattern in device_lower for pattern in patterns):
                    update_device_totals(device_type, user)
                    matched = True
                    break
            if not matched:
                update_device_totals('other_device', user)
        delete_sql = "DELETE FROM speed_report.t_user_device_day WHERE date=%s"
        mysql_execute(self.conn, delete_sql % date.replace("-", ""))
        insert_queries = []
        for system, values in device_totals.items():
            insert_sql = """INSERT INTO speed_report.t_user_device_day SET date='{date}', device='{device}', total={total}, new={new}, retained={retained}, created_at=NOW(), total_recharge={total_recharge}, total_recharge_money={total_recharge_money}, new_recharge_money={new_recharge_money};""".format(
                date=date.replace("-", ""),
                device=system,
                total=values['total_cnt'],
                new=values['new_cnt'],
                retained=values['retained_cnt'],
                total_recharge=values['total_recharge'],
                total_recharge_money=values['total_recharge_amount'],
                new_recharge_money=values['new_recharge_amount']
            )
            insert_queries.append(insert_sql)
        for query in insert_queries:
            mysql_execute(self.conn, query)

    def insert_daily_channel_recharge_by_month(self, date, data):
        mysql_execute(self.conn, "delete from speed_report.t_user_channel_month where date=%s" % date.replace("-", ""))
        for channel in data.keys():
            user = data[channel]
            channel = channel if channel else '官网'
            sql = """insert into speed_report.t_user_channel_month set date='{}', channel='{}', total={}, month_total={},month_new={},total_recharge={},total_recharge_money={},month_total_recharge={},month_recharge_money={},created_at=now();""".format(date.replace("-", ""), channel, user["total_cnt"], user["month_retained_cnt"], user["month_new_cnt"],user["total_recharge_cnt"], user["total_recharge_money_cnt"], user["month_recharge_cnt"], user["month_recharge_money_cnt"])
            mysql_execute(self.conn, sql)
    def insert_daily_device_retention(self, date, data):
        device_map = {'android': ['android'],'ios': ['iphone'],'mac': ['mac'],'win': ['win']}
        device_totals = {key: {'new_cnt': 0, 'retained_cnt': 0, 'day7_retained': 0, 'day15_retained': 0} for key in device_map.keys()}
        device_totals['other_device'] = {'new_cnt': 0, 'retained_cnt': 0, 'day7_retained': 0, 'day15_retained': 0}
        def update_device_totals(device_type, user_data):
            device_totals[device_type]['new_cnt'] += user_data.get("new_cnt", 0)
            device_totals[device_type]['retained_cnt'] += user_data.get("retained_cnt", 0)
            device_totals[device_type]['day7_retained'] += user_data.get("day7_retained", 0)
            device_totals[device_type]['day15_retained'] += user_data.get("day15_retained", 0)
        for device, user in data.items():
            device_lower = device.lower()
            matched = False
            for device_type, patterns in device_map.items():
                if any(pattern in device_lower for pattern in patterns):
                    update_device_totals(device_type, user)
                    matched = True
                    break
            if not matched:
                update_device_totals('other_device', user)
        delete_sql = "DELETE FROM speed_report.t_user_device_retention WHERE date=%s"
        mysql_execute(self.conn, delete_sql % date.replace("-", ""))
        insert_queries = []
        for system, values in device_totals.items():
            insert_sql = """INSERT INTO speed_report.t_user_device_retention SET date='{date}', device='{device}',new={new}, retained={retained}, created_at=NOW(), day7_retained={day7_retained}, day15_retained={day15_retained};""".format(
                date=date.replace("-", ""),
                device=system,
                new=values['new_cnt'],
                retained=values['retained_cnt'],
                day7_retained=values['day7_retained'],
                day15_retained=values['day15_retained'],
            )
            insert_queries.append(insert_sql)
        for query in insert_queries:
            mysql_execute(self.conn, query)

    def insert_or_update_daily_device_report(self, date, device, new, retained, day7_retained, day15_retained):
        date = int(date.replace('-', ''))
        current_time = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        sql = """
        INSERT INTO speed_report.t_user_device_retention (date, device, new, retained, day7_retained, day15_retained, created_at)
        VALUES (%s, '%s', %s, %s, %s, %s,'%s')
        ON DUPLICATE KEY UPDATE
        new = VALUES(new),
        retained = VALUES(retained),
        day7_retained = VALUES(day7_retained),
        day15_retained = VALUES(day15_retained);
        """ % (date, device, new, retained, day7_retained, day15_retained, current_time)
        mysql_execute(self.conn, sql)

    def clear_t_user_report_monthly(self):
        sql = """TRUNCATE TABLE t_user_report_monthly;"""
        mysql_execute(self.conn, sql)

    def insert_into_report_monthly(self, stat_month, os, user_count, new_users, retained_users):
        stat_month_date = int(stat_month.replace('-', ''))
        print(stat_month, os, user_count, new_users, retained_users)
        sql = """INSERT INTO t_user_report_monthly (stat_month, os, user_count, new_users, retained_users) VALUES (%s, '%s', %s, %s, %s);""" % (stat_month_date, os, user_count, new_users, retained_users)
        mysql_execute(self.conn, sql)

class SpeedCollector:
    def __init__(self):
        self.config = load_config("/shell/report/config.yaml")
        self.conn = mysql_connect(self.config["collector-speed-db"])
    def close_connection(self):
        if self.conn:
            self.conn.close()
    def check_task(self, date):
        sql = f"SELECT COUNT(*) as cnt FROM t_task WHERE ip='node-all' AND date='{date.replace('-','')}' AND type='node-user'"
        rows = mysql_query_db(self.conn, sql)
        if rows[0]["cnt"] > 0:
            return True  # 代表执行成功
        else:
            return False  # 代表没有执行成功
