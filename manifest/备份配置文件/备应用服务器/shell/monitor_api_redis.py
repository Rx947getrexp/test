# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import time
import fcntl

import logging
# from logging.handlers import TimedRotatingFileHandler
from logging.handlers import RotatingFileHandler

# * * * * * /usr/bin/python3 /shell/monitor_api_redis.py &
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

    filehandler = RotatingFileHandler(filename=file_name, mode='a', maxBytes=1024 * 1024 * 200, backupCount=2)
    filehandler.setFormatter(formatter)
    log.addHandler(filehandler)
    logging.info("init_logging success")

def is_process_alive(name):
    command = f"pgrep {name}"
    output = subprocess.getoutput(command)
    if output:
        logging.info("%s is running, output: %s" % (name, output))
        return True
    else:
        logging.info("%s is exited, output: %s" % (name, output))
        logging.info(name, "is not running", output)
        return False


def run_background_program(command):
    logging.info(command)
    logging.info(subprocess.Popen(command, shell=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL))


def print_cmd(cmd):
    result = os.popen(cmd)
    res = result.read()
    logging.info(res)


if __name__ == '__main__':
    lock_file = "/tmp/monitor_api_redis.lock"
    fp = open(lock_file, "w")
    try:
        fcntl.lockf(fp, fcntl.LOCK_EX | fcntl.LOCK_NB)
    except IOError:
        logging.info("已经有一个 monitor_api_redis 进程在运行，本进程将退出")
        sys.exit(1)

    init_logging("/shell/log/monitor_api_redis.log")
    logging.info("start monitor")
    while True:
        logging.info("-"*20)
        if not is_process_alive("redis"):
            run_background_program("/usr/local/bin/redis-server /etc/redis/redis.conf")
            time.sleep(1)

        if not is_process_alive("go-admin"):
            run_background_program("cd /wwwroot/go/go-admin && ./restart.sh")
            time.sleep(1)

        if not is_process_alive("go-api"):
            run_background_program("cd /wwwroot/go/go-api && ./restart.sh")
            time.sleep(1)

        #if not is_process_alive("go-upload"):
        #    run_background_program("cd /wwwroot/go/go-upload && ./restart.sh")
        #    time.sleep(1)

        #if not is_process_alive("go-job"):
        #    run_background_program("cd /wwwroot/go/go-job && ./restart.sh")
        #    time.sleep(1)

        if not is_process_alive("go-fly"):
            run_background_program("cd /wwwroot/go/go-fly && ./restart.sh")
            time.sleep(1)

        if not is_process_alive("node_exporter"):
            run_background_program(
                "nohup /data/hs-fly/node_exporter/node_exporter >> /var/log/node_exporter/output.log 2>&1 &")
            time.sleep(1)

        time.sleep(2)

