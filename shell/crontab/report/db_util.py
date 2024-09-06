# -*- coding: utf-8 -*-
import logging
import sys

import pymysql

from common import load_config


def mysql_connect(db):
    """连接DB
    """
    return pymysql.connect(host=db["host"], port=db["port"], user=db["user"], passwd=db["pswd"], db=db["db"],
                           charset=db["charset"])


def mysql_query_db(_conn, sql):
    """读取DB数据
    """
    logging.info(sql)
    cursor = _conn.cursor(cursor=pymysql.cursors.DictCursor)
    cursor.execute(sql)
    rows = cursor.fetchall()
    cursor.close()
    return rows


def mysql_execute(_conn, sql):
    """执行sql
    """
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
        self.country = "193.233.48.70", "46.17.41.7", "45.147.201.21", "45.147.200.112", "46.17.44.132", "92.118.112.89", "38.55.136.95", "207.90.237.91", "91.149.218.194", "62.133.60.81", "82.118.21.147", "62.133.63.237", "213.159.68.106", "103.198.203.11", "110.42.42.229", "185.39.207.20", "92.118.112.133", "212.18.104.23", "147.45.178.51", "185.39.207.104"

    def close_connection(self):
        if self.conn:
            self.conn.close()

    def convert_to_rub(self, amount, currency):
        exchange_rates = {
            'USD': self.exchange_rate_usd,
            'WMZ': self.exchange_rate_wmz
        }
        return float(amount) * exchange_rates.get(currency, 1)

    def get_user_ip_list(self):
        sql = """SELECT t1.id, t1.email, t2.ip, t2.updated_at FROM t_user t1 JOIN (SELECT user_id, MAX(updated_at) AS max_updated_at FROM user_logs GROUP BY user_id
    ) AS latest_logs ON t1.id = latest_logs.user_id JOIN user_logs t2 ON t1.id = t2.user_id AND t2.updated_at = latest_logs.max_updated_at;"""
        rows = mysql_query_db(self.conn, sql)
        dic = {}
        for row in rows:
            dic[row["email"]] = {
                "id": row["id"],
                "ip": row["ip"],
            }
        return dic

    def get_users(self):
        sql = """select id, email, channel,last_login_country,created_at from t_user;"""
        rows = mysql_query_db(self.conn, sql)
        dic = {}
        for row in rows:
            dic[row["email"]] = {
                "id": row["id"],
                "channel": row["channel"],
                "country": row['last_login_country'],
                "register_date": row['created_at']
            }
        return dic

    def count_user_by_create_time(self, channel_id, st, et):
        sql = """select count(*) as cnt from t_user where channel_id= '%d' and created_at >= '%s' and created_at <= '%s';""" % (
            channel_id, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_node_by_create_time(self, ip, st, et):
        sql = """SELECT count(DISTINCT email) as cnt FROM t_user_traffic WHERE ip='%s' AND email IN (select email as cnt from t_user where created_at >= '%s' and created_at <= '%s') AND created_at >= '%s' and created_at <= '%s';""" % (
            ip, st, et, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_online(self, channel_id, date, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where date = '%s' and email in (select email from t_user where channel_id='%d' and created_at <= '%s');""" % (
            date.replace("-", ""), channel_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_month_online(self, channel_id, st, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where email in (select email from t_user where channel_id='%d' and created_at <= '%s') and created_at >= '%s' and created_at <= '%s';""" % (channel_id,et, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_node_online(self, ip, date, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where date = '%s' and ip = '%s' and email in (select email from t_user WHERE created_at <= '%s');""" % (
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
        sql = """select count(DISTINCT email) as cnt from t_user_traffic where ip = '%s' and created_at <= '%s';""" % (
            ip, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def query_channel_id_list(self, et):
        sql = """select distinct channel_id as channel_id from t_user where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_node_traffic_list(self, et):
        sql = """select distinct ip as ip from t_user_traffic where created_at <= '%s' and ip in %s;""" % (
            et, self.country)
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_user_traffic_list(self, date):
        sql = """select distinct email as email from t_user_traffic where date=%s;""" % date.replace("-", "")
        return mysql_query_db(self.conn, sql)

    def query_node_user_traffic_list(self, date):
        sql = """select email as email,ip,uplink,downlink from t_user_traffic where date=%s;""" % date.replace("-", "")
        return mysql_query_db(self.conn, sql)

    def query_user_traffic_log_list(self, email, st, et):
        sql = """select ip, date_time, uplink, downlink from t_user_traffic_log where email = '%s' and date_time>='%s' and date_time<='%s' order by date_time asc;""" % (
            email, st, et)
        return mysql_query_db(self.conn, sql)

    def query_channel_list(self, et):
        sql = """select distinct channel as channel from t_user where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_recharge_channel_list(self, et):
        sql = """SELECT distinct t2.channel as channel FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) AND t2.channel !=''and t1.created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def count_total_channel_user(self, channel, et):
        sql = """select count(*) as cnt from t_user where channel = '%s' and created_at <= '%s';""" % (channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_recharge_total_channel_user(self, channel, goods_id, currency, et):
        sql = """select count(*) as cnt FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and t1.goods_id = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.currency='%s' and t1.created_at <= '%s';""" % (
            channel, goods_id, currency, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_channel_user_by_create_time(self, channel, st, et):
        sql = """select count(*) as cnt from t_user where channel= '%s' and created_at >= '%s' and created_at <= '%s';""" % (
            channel, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_recharge_channel_user_by_create_time(self, channel, goods_id, currency, st, et):
        sql = """select count(*) as cnt FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and t1.goods_id = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.currency='%s' and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (
            channel, goods_id, currency, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_channel_user_online(self, channel, date, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where date = '%s' and email in (select email from t_user where channel='%s' and created_at <= '%s');""" % (
            date.replace("-", ""), channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_channel_user_by_month(self, channel, st, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where email in (select email from t_user where channel='%s' and created_at >= '%s' and created_at <= '%s') and created_at >= '%s' and created_at <= '%s';""" % (channel, st, et,st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]


    def query_goods_id_list(self):
        sql = """select id from t_goods;"""
        rows = mysql_query_db(self.conn, sql)
        return rows

    def query_currency_list(self):
        sql = """select DISTINCT currency from t_pay_order;"""
        rows = mysql_query_db(self.conn, sql)
        return rows

    def count_total_user_recharge(self, goods_id, et):
        sql = """select count(distinct user_id) as cnt from t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at <= '%s';""" % (
            goods_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_recharge_by_create_time(self, goods_id, st, et):
        sql = """select count(distinct user_id) as cnt from t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at >= '%s' and created_at <= '%s';""" % (
            goods_id, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_user_recharge_times(self, goods_id, et):
        sql = """select count(user_id) as cnt from t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at <= '%s';""" % (
            goods_id, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_recharge_times_by_create_time(self, goods_id, st, et):
        sql = """select count(user_id) as cnt from t_pay_order where ((currency = 'RUB' AND status = 'admin-confirm-passed') OR (currency IN ('USD', 'WMZ','RUB') AND status = 'paid')) and goods_id = %d and created_at >= '%s' and created_at <= '%s';""" % (
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
        sql = """select count(*) as cnt FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_number_of_recharges_by_create_time(self, channel, st, et):
        sql = """select count(*) as cnt FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (channel,st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_recharge_amount(self, channel, et):
        '''
        充值总金额
        '''
        sql = """select t1.currency, t1.order_reality_amount,t1.status FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (
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
        sql = """select t1.currency, t1.order_reality_amount,t1.status FROM t_pay_order t1 JOIN t_user t2 on t1.email = t2.email  WHERE t2.channel = '%s' and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (
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
        sql = """select distinct os as device from t_dev where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def count_total_device_user(self, device, et):
        sql = """SELECT count(*) as cnt FROM t_user u JOIN t_dev d ON u.client_id = d.client_id WHERE u.client_id !='' AND d.os= '%s' and u.created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_device_user_by_create_time(self, device, st, et):
        sql = """SELECT count(*) as cnt FROM t_user u JOIN t_dev d ON u.client_id = d.client_id WHERE u.client_id !='' AND d.os= '%s' AND u.created_at >= '%s' and u.created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_device_user_online(self, device, date, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where date = '%s' and email in (SELECT u.email FROM t_user u JOIN t_dev d ON u.client_id = d.client_id WHERE u.client_id !='' and d.os='%s' and u.created_at <= '%s');""" % (
            date.replace("-", ""), device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_number_of_device_recharges(self, device, et):
        sql = """select count(*) as cnt FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.email in (SELECT DISTINCT(u.email) FROM t_user u JOIN t_dev d ON u.client_id = d.client_id WHERE u.client_id !='' AND d.os='%s') and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_device_recharge_amount(self, device, et):
        sql = """select t1.currency, t1.order_reality_amount,t1.status FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE t2.email IN (SELECT DISTINCT u.email FROM t_user u JOIN t_dev d ON u.client_id = d.client_id WHERE u.client_id != '' AND d.os = '%s') and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        total_recharge = 0
        for row in rows:
            status = row['status']
            currency = row['currency']
            amount = row['order_reality_amount']
            if status == "paid":
                amount_in_rub = self.convert_to_rub(amount, currency)
            elif status == "admin-confirm-passed":
                amount_in_rub = float(amount)
            else:
                continue
            total_recharge += amount_in_rub
        return total_recharge

    def count_total_device_recharge_amount_by_create_time(self, device, st, et):
        sql = """select t1.currency, t1.order_reality_amount,t1.status  FROM t_pay_order t1 JOIN t_user t2 on t1.email = t2.email  WHERE t2.email in (SELECT u.email FROM t_user u JOIN t_dev d ON u.client_id = d.client_id WHERE u.client_id !='' AND d.os='%s') and ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at >= '%s' and t1.created_at <= '%s';""" % (
            device, st, et)
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

    def all_device_total_recharge_amount(self, et):
        sql = """select t2.email,t1.currency, t1.order_reality_amount,t1.status FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) and t1.created_at <= '%s';""" % (
            et)
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

    def all_device_total_recharge_amount_by_create_time(self, st, et):
        sql = """select t2.email,t1.currency, t1.order_reality_amount,t1.status FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) AND t1.created_at >= '%s' and t1.created_at <= '%s';""" % (
            st, et)
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

    def all_device_recharge(self, et):
        sql = """select count(*) as cnt FROM t_pay_order t1 JOIN t_user t2 on t1.email=t2.email  WHERE ((t1.currency = 'RUB' AND t1.status = 'admin-confirm-passed') OR (t1.currency IN ('USD', 'WMZ','RUB') AND t1.status = 'paid')) AND t1.created_at <= '%s';""" % (
            et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]


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
            "38.55.136.95": "美国2",
            "207.90.237.91": "美国3",
            "92.118.112.133": "美国4",
            "212.18.104.23": "美国5",
            "91.149.218.194": "法国1",
            "62.133.60.81": "德国1",
            "147.45.178.51": "德国2",
            "82.118.21.147": "波兰1",
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
        sql = """SELECT DISTINCT device_type as device FROM t_user_op_log  where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows

    def get_total_clicks(self, device, et):
        sql = """SELECT count(email) as cnt from t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_yesterday_day_clicks(self, device, st, et):
        sql = """SELECT count(email) as cnt from t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_weekly_clicks(self, device, st, et):
        sql = """SELECT count(email) as cnt from t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_total_users_clicked(self, device, et):
        sql = """SELECT count(distinct email) as cnt from t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at <= '%s';""" % (
            device, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_yesterday_day_users_clicked(self, device, st, et):
        sql = """SELECT count(distinct email) as cnt from t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def get_weekly_users_clicked(self, device, st, et):
        sql = """SELECT count(distinct email) as cnt from t_user_op_log WHERE content='进入充值页面' and device_type='%s' and created_at >= '%s' and created_at <= '%s';""" % (
            device, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            sys.exit(1)
        return rows[0]["cnt"]

    def insert_daily_user(self, date, data):
        logging.info(data)
        for channel_id in data.keys():
            user = data[channel_id]
            sql = """insert into t_user_report_day set date=%s, channel_id=%d, total=%d, new=%d, retained=%d,month_retained=%d, created_at=now();""" % (
                date.replace("-", ""), channel_id, user["total_cnt"], user["new_cnt"], user["retained_cnt"],
                user["month_retained_cnt"])
            mysql_execute(self.conn, sql)

    def insert_online_user_day(self, date, items):
        for item in items:
            sql = """insert into t_user_online_day set date='{}',email='{}',channel='{}',online_duration={},uplink={},downlink={},last_login_country='{}',created_at=now();""".format(
                date.replace("-", ""), item["email"], item["channel"], item["online_duration"], item["uplink"],
                item["downlink"], item["country"])
            mysql_execute(self.conn, sql)

    def insert_daily_channel_user(self, date, data):
        for channel in data.keys():
            user = data[channel]
            channel = channel if channel else '官网'
            sql = """insert into t_user_channel_day set date='{}', channel='{}', total={}, new={},retained={}, created_at=now(),total_recharge={},total_recharge_money={},new_recharge_money={};""".format(
                date.replace("-", ""), channel, user["total_cnt"], user["new_cnt"], user["retained_cnt"],
                user["total_recharge"], user["total_recharge_amount"], user["new_recharge_amount"])
            mysql_execute(self.conn, sql)

    def insert_daily_node(self, date, data):
        for ip in data.keys():
            user = data[ip]
            name = self.data_name.get(ip)
            if name:
                sql = """insert into t_user_node_day set date='{}', ip='{}', total={},new={},retained={}, created_at=now();""".format(
                    date.replace("-", ""), name, user["total_cnt"], user["new_cnt"], user["retained_cnt"])
                mysql_execute(self.conn, sql)

    def insert_online_user_node_day(self, date, items):
        for item in items:
            ip_name = self.data_name.get(item["node"])
            sql = """insert into t_user_node_online_day set date='{}',email='{}',channel='{}',online_duration={},uplink={},downlink={},node='{}',register_date='{}',created_at=now();""".format(
                date.replace("-", ""), item["email"], item["channel"], item["online_duration"], item["uplink"],
                item["downlink"], ip_name, item["register_date"])
            mysql_execute(self.conn, sql)

    def insert_daily_user_recharge(self, date, data):
        logging.info(data)
        for goods_id in data.keys():
            user = data[goods_id]
            sql = """insert into t_user_recharge_report_day set date=%s, goods_id=%d, total=%d, new=%d,created_at=now();""" % (
                date.replace("-", ""), goods_id, user["total_cnt"], user["new_cnt"])
            mysql_execute(self.conn, sql)

    def insert_daily_user_recharge_times(self, date, data):
        logging.info(data)
        for goods_id in data.keys():
            user = data[goods_id]
            sql = """insert into t_user_recharge_times_report_day set date=%s, goods_id=%d, total=%d, new=%d,created_at=now();""" % (
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
        for channel, user_data in data.items():
            for goods_id, recharge_data in user_data.items():
                goods_name = goods_name_dic.get(goods_id, "未知套餐")
                # 获取USD和RUB的支付总次数和新增次数
                usd_total = recharge_data.get("USD", {}).get("total_cnt", 0)
                usd_new = recharge_data.get("USD", {}).get("new_cnt", 0)
                rub_total = recharge_data.get("RUB", {}).get("total_cnt", 0)
                rub_new = recharge_data.get("RUB", {}).get("new_cnt", 0)
                sql = f"""INSERT INTO t_user_channel_recharge_day (date, channel, goods_name, usd_total, usd_new, rub_total, rub_new, created_at) 
                                         VALUES ({date.replace("-", "")}, "{channel}", "{goods_name}", {usd_total}, {usd_new}, {rub_total}, {rub_new}, NOW())"""
                mysql_execute(self.conn, sql)

    def insert_daily_device_action(self, date, data):
        for device_type in data.keys():
            device = data[device_type]
            sql = """insert into t_user_device_action_day set date='%s', device='%s', total_clicks=%d, yesterday_day_clicks=%d, weekly_clicks=%d,total_users_clicked=%d,yesterday_day_users_clicked=%d,weekly_users_clicked=%d,created_at=now();""" % (
                date.replace("-", ""), device_type, device["total_clicks"], device["yesterday_day_clicks"],
                device["weekly_clicks"], device["total_users_clicked"], device["yesterday_day_users_clicked"],
                device["weekly_users_clicked"])
            mysql_execute(self.conn, sql)

    def insert_daily_device_user(self, date, data):
        device_totals = {'android': {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0,
                                     'total_recharge_amount': 0, 'new_recharge_amount': 0},
                         'ios': {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0,
                                 'total_recharge_amount': 0, 'new_recharge_amount': 0},
                         'mac': {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0,
                                 'total_recharge_amount': 0, 'new_recharge_amount': 0},
                         'win': {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0,
                                 'total_recharge_amount': 0, 'new_recharge_amount': 0},
                         'null_device': {'total_cnt': 0, 'new_cnt': 0, 'retained_cnt': 0, 'total_recharge': 0,
                                         'total_recharge_amount': 0, 'new_recharge_amount': 0}
                         }
        for device, user in data.items():
            if 'android' in device.lower():
                device_totals['android']['total_cnt'] += user.get("total_cnt", 0)
                device_totals['android']['new_cnt'] += user.get("new_cnt", 0)
                device_totals['android']['retained_cnt'] += user.get("retained_cnt", 0)
                device_totals['android']['total_recharge'] += user.get("total_recharge", 0)
                device_totals['android']['total_recharge_amount'] += user.get("total_recharge_amount", 0)
                device_totals['android']['new_recharge_amount'] += user.get("new_recharge_amount", 0)
            elif 'iphone' in device.lower():
                device_totals['ios']['total_cnt'] += user.get("total_cnt", 0)
                device_totals['ios']['new_cnt'] += user.get("new_cnt", 0)
                device_totals['ios']['retained_cnt'] += user.get("retained_cnt", 0)
                device_totals['ios']['total_recharge'] += user.get("total_recharge", 0)
                device_totals['ios']['total_recharge_amount'] += user.get("total_recharge_amount", 0)
                device_totals['ios']['new_recharge_amount'] += user.get("new_recharge_amount", 0)
            elif 'mac' in device.lower():
                device_totals['mac']['total_cnt'] += user.get("total_cnt", 0)
                device_totals['mac']['new_cnt'] += user.get("new_cnt", 0)
                device_totals['mac']['retained_cnt'] += user.get("retained_cnt", 0)
                device_totals['mac']['total_recharge'] += user.get("total_recharge", 0)
                device_totals['mac']['total_recharge_amount'] += user.get("total_recharge_amount", 0)
                device_totals['mac']['new_recharge_amount'] += user.get("new_recharge_amount", 0)
            elif 'win' in device.lower():
                device_totals['win']['total_cnt'] += user.get("total_cnt", 0)
                device_totals['win']['new_cnt'] += user.get("new_cnt", 0)
                device_totals['win']['retained_cnt'] += user.get("retained_cnt", 0)
                device_totals['win']['total_recharge'] += user.get("total_recharge", 0)
                device_totals['win']['total_recharge_amount'] += user.get("total_recharge_amount", 0)
                device_totals['win']['new_recharge_amount'] += user.get("new_recharge_amount", 0)
            else:
                device_totals['null_device']['total_cnt'] += user.get("total_cnt", 0)
                device_totals['null_device']['new_cnt'] += user.get("new_cnt", 0)
                device_totals['null_device']['retained_cnt'] += user.get("retained_cnt", 0)
                device_totals['null_device']['total_recharge'] += user.get("total_recharge", 0)
                device_totals['null_device']['total_recharge_amount'] += user.get("total_recharge_amount", 0)
                device_totals['null_device']['new_recharge_amount'] += user.get("new_recharge_amount", 0)
        for system, values in device_totals.items():
            sql = """insert into t_user_device_day set date='{}', device='{}', total={}, new={}, retained={}, created_at=now(), total_recharge={}, total_recharge_money={}, new_recharge_money={};""".format(
                date.replace("-", ""), system, values['total_cnt'], values['new_cnt'], values['retained_cnt'],
                values['total_recharge'], values['total_recharge_amount'], values['new_recharge_amount'])
            mysql_execute(self.conn, sql)

    def insert_daily_channel_recharge_by_month(self, date, data):
        for channel in data.keys():
            user = data[channel]
            channel = channel if channel else '官网'
            sql = """insert into t_user_channel_month set date='{}', channel='{}', total={}, month_total={},month_new={},total_recharge={},total_recharge_money={},month_total_recharge={},month_recharge_money={},created_at=now();""".format(date.replace("-", ""), channel, user["total_cnt"], user["month_retained_cnt"], user["month_new_cnt"],user["total_recharge_cnt"], user["total_recharge_money_cnt"], user["month_recharge_cnt"], user["month_recharge_money_cnt"])
            mysql_execute(self.conn, sql)
