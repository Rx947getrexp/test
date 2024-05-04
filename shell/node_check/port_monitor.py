import telnetlib
import time
import logging
import json
import hashlib
import hmac
import requests

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

with open("health_scores.json", "r") as f:
    health_scores = json.load(f)

# 连续失败的阈值
FAILURE_THRESHOLD = 3


def Switching_machine_requests(siteName, status):
    url = "http://127.0.0.1:13002/machine_states_witching"
    secret_key = "3f5202f0-4ed3-4456-80dd-13638c975bda"
    params = {'Ip': siteName, 'status': status}
    signature = hmac.new(secret_key.encode(), secret_key.encode(), hashlib.sha256).hexdigest()
    headers = {"X-Signature": signature}
    response = requests.post(url, headers=headers, params=params)
    return response


# 检查端口是否开放
def check_port(ip, port):
    try:
        with telnetlib.Telnet(ip, port, timeout=10) as telnet:
            logging.info(f"成功连接到 {ip}:{port}")
            return True
    except Exception as e:
        logging.warning(f"连接到 {ip}:{port} 失败: {e}")
        return False
    finally:
        if 'telnet' in locals():
            telnet.close()


# 更新失败次数
def update_failure_count(siteName, ip, port, healthy):
    # 初始化键
    if str(port) not in health_scores[siteName][ip]:
        health_scores[siteName][ip][str(port)] = {"failure_count": 0}
    # 根据成功与否更新失败次数
    if healthy:
        # 连续失败超过三次了,后面成功了
        if health_scores[siteName][ip][str(port)]["failure_count"] >= FAILURE_THRESHOLD:
            recommission(siteName, ip)
        else:
            health_scores[siteName][ip][str(port)]["failure_count"] = 0
    else:
        # 如果失败，增加失败次数
        health_scores[siteName][ip][str(port)]["failure_count"] += 1
    # 如果连续失败超过三次,那么可以下架了
    if health_scores[siteName][ip][str(port)]["failure_count"] >= FAILURE_THRESHOLD:
        print(health_scores[siteName][ip])
        decommission(siteName, ip)


def recommission(siteName, siteNameIp):
    """
    :param siteName: 服务器名称
    :param siteNameIp: 服务器IP地址
    :return:
    """
    Switching_machine_requests(siteNameIp, "1")
    logging.warning(f"对 {siteName}进行上架架操作")


def decommission(siteName, siteNameIp):
    Switching_machine_requests(siteNameIp, "2")
    logging.warning(f"对 {siteName}进行下架操作")


# 主函数
def main():
    siteNames = {
        "Hong Kong": ["103.84.110.102"],
        "Moscow": ["185.143.220.131"],
        "Inner Mongolia": ["211.101.233.165"],
        "Germany": ["91.149.222.79"],
    }
    batch_ports = [443, 13001, 13002, 13003, 13004, 13005]
    for siteName, ips in siteNames.items():
        for ip in ips:
            batch_success = False
            for port in batch_ports:
                healthy = check_port(ip, port)
                batch_success |= healthy
                # 调用位运算符判断443~13005的结果
                if healthy:
                    break
            update_failure_count(siteName, ip, "443~13005", batch_success)
    with open("health_scores.json", "w") as f:
        json.dump(health_scores, f)


if __name__ == "__main__":
    start_time = time.time()
    main()
    end_time = time.time()
    logging.info(f"运行时间：{end_time - start_time}s")
    logging.info(f"各节点健康状态：{json.dumps(health_scores)}")
