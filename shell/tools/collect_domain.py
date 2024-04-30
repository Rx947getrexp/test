# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import time
import fcntl
import traceback

import re

# 定义一个函数，用于检查字符串是否为 IP 地址
def is_ip_address(s):
    # 使用正则表达式匹配 IPv4 地址
    ipv4_pattern = r"^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
    ipv4_match = re.match(ipv4_pattern, s)
    if ipv4_match:
        return True

    # 使用正则表达式匹配 IPv6 地址
    ipv6_pattern = r"^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|(:(:[0-9a-fA-F]{1,4}){1,4}){1,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$"
    ipv6_match = re.match(ipv6_pattern, s)

    if ipv6_match:
        return True

    return False
    # return ipv4_match is not None or ipv6_match is not None


def run():
    cmd = "cat /var/log/v2ray/access.log | awk '{print $5}' | grep '\.cn\|\.weixin\.\|\.qq\.'"
    cmd = "cat /var/log/v2ray/access.log | awk '{print $5}' | grep '\.weixin\.\|\.qq\.\|wx\.'"
    cmd = "cat /var/log/v2ray/access.log | awk '{print $5}'"
    process = subprocess.Popen(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
    output, error = process.communicate()

    # 如果有错误信息，打印错误信息
    if error:
        print(f"Error: {error.decode('utf-8')}")
        return

    # print(output.decode('utf-8'))
    domain_dic = {}
    for line in output.decode('utf-8').split('\n'):
        stripped_line = line.strip()
        stripped_line = stripped_line.replace("[", "")
        stripped_line = stripped_line.replace("]", "")
        if not stripped_line:
            continue

        split_value = stripped_line.split(':')
        length = len(split_value)
        if length in [1,2]:
            # print(split_value)
            continue

        if length != 3:
            continue
        # if length < 3 or length > 10:
        #     print(split_value, "长度检查失败")
        #     sys.exit()

        # domain = ":".join(split_value[1:-1])
        domain = split_value[1]
        # if length == 3:
        #     domain = split_value[1]
        # else:
        #     temp = [split_value[1],split_value[2],split_value[3],split_value[4],split_value[5],split_value[6]]
        #     domain = ":".join(temp)
        if domain == "one.one.one.one" or is_ip_address(domain):
            continue

        if domain in domain_dic:
            domain_dic[domain] = domain_dic[domain]+1
        else:
            domain_dic[domain] = 1

    sorted_dict = sorted(domain_dic.items(), key=lambda x: x[1], reverse=True)
    for key, value in sorted_dict:
        # print(f"{key}\t{value}")
        print(f"{key}")


def ping(domain):
    command = ['ping', '-c', '1', domain]
    return subprocess.call(command, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL) == 0

def check_domain():
    with open('domains.txt', 'r') as file:
        for line in file:
            domain = line.strip()
            if ping(domain):
                # print(f'{domain} is reachable')
                print(f'{domain}')
            # else:
            #     print(f'{domain} is not reachable')

if __name__ == '__main__':
    # check_domain()
    run()
    # v = "tcp:tns-counter.ru:443"
    # v = "udp:[2606:4700:3032::6815:fe5]:443"
    # v = "udp:2001:b28:f23d:f101::111:208:1400"
    # print(len(v.split(":")))
    # print(v.split(":"))
