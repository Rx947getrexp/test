#!/usr/bin/env python3
import xml.etree.ElementTree as ET
import subprocess
import datetime

# 定义要检测的端口和协议
PORTS = [22, 9100, 80, 443, 3306, 6379,13002,13005,8081]
PROTOCOL = "tcp"

# 定义日志文件路径
LOGFILE = "/shell/log/port_opening.log"

# 记录执行日期和时间
current_time = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
with open(LOGFILE, 'a') as log_file:
    log_file.write(f"执行日期: {current_time}\n")

def is_port_open_in_firewalld(port, protocol):
    # 解析防火墙规则文件
    firewalld_file = "/etc/firewalld/zones/public.xml"
    tree = ET.parse(firewalld_file)
    root = tree.getroot()

    # 查找指定端口和协议的规则
    for port_elem in root.findall("./port"):
        if int(port_elem.attrib.get("port")) == port and port_elem.attrib.get("protocol", "") == protocol:
            return True

    return False

def open_port(port):
    # 尝试开放端口
    subprocess.run(f"firewall-cmd --zone=public --add-port={port}/{PROTOCOL} --permanent", shell=True, check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    subprocess.run("firewall-cmd --reload", shell=True, check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    log(f"端口 {port} 已尝试开放。")

def log(message):
    with open(LOGFILE, 'a') as log_file:
        log_file.write(message + "\n")
    print(message)

# 主循环
for port in PORTS:
    if port != 22 and not is_port_open_in_firewalld(port, PROTOCOL):
        log(f"端口 {port} 未在防火墙规则中开放，尝试开放端口...")
        open_port(port)
    else:
        log(f"端口 {port} 已经在防火墙规则中开放，无需操作。")

