# -*- coding: utf-8 -*-
import json
import logging
import sys

from common import *

import pymysql


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

    def count_user_by_create_time(self, channel, st, et):
        sql = """select count(*) as cnt from t_user where channel= '%s' and created_at >= '%s' and created_at <= '%s';""" % (channel, st, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_user_online(self, channel, date, et):
        sql = """select count(distinct email) as cnt from t_user_traffic where date = '%s' and email in (select email from t_user where channel='%s' and created_at <= '%s');""" % (
            date.replace("-", ""), channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def count_total_user(self, channel, et):
        sql = """select count(*) as cnt from t_user where channel = '%s' and created_at <= '%s';""" % (channel, et)
        rows = mysql_query_db(self.conn, sql)
        if len(rows) != 1:
            logging.error("sql: %s, rows: %d != 1" % (sql, len(rows)))
            sys.exit(1)
        return rows[0]["cnt"]

    def query_channel_id_list(self, et):
        sql = """select distinct channel as channel from t_user where created_at <= '%s';""" % et
        rows = mysql_query_db(self.conn, sql)
        return rows


class SpeedReport:
    def __init__(self):
        self.config = load_config("/shell/report/config.yaml")
        self.conn = mysql_connect(self.config["report-speed-db"])

    def insert_daily_user(self, date, data):
        logging.info(data)
        for channel_id in data.keys():
            user = data[channel_id]
            sql = """insert into t_user_report_day set date=%s, channel_id=%d, total=%d, new=%d, retained=%d, created_at=now();""" % (
                date.replace("-", ""), channel_id, user["total_cnt"], user["new_cnt"], user["retained_cnt"])
            mysql_execute(self.conn, sql)
