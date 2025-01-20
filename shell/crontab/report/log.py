# -*- coding: utf-8 -*-
# import os
# import subprocess
# import sys
# import time
# import fcntl

import logging
from logging.handlers import RotatingFileHandler


def init_logging(file_name):
    log = logging.getLogger()
    log.setLevel(logging.DEBUG)
    formatter = logging.Formatter('%(asctime)s - [line]:%(lineno)d - %(levelname)s - %(message)s',
                                  '%Y-%m-%d %H:%M:%S')

    ch = logging.StreamHandler()  # 输出到控制台的handler
    ch.setFormatter(formatter)
    ch.setLevel(logging.DEBUG)  # 也可以不设置，不设置就默认用logger的level

    # log_file_handler = TimedRotatingFileHandler(filename=file_name, when="D", interval=1, backupCount=14)
    # log_file_handler.setFormatter(formatter)
    # log.addHandler(log_file_handler)

    handler = RotatingFileHandler(filename=file_name, mode='a', maxBytes=1024 * 1024 * 200, backupCount=2,encoding='utf-8')
    handler.setFormatter(formatter)
    log.addHandler(handler)
    logging.info("init_logging success")
