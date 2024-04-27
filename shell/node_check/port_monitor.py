import concurrent.futures
import json
import telnetlib
import time
import logging

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

with open("health_scores.json", "r") as f:
    health_scores = json.load(f)

# 健康阈值
HEALTH_THRESHOLD = 90
# 恢复阈值
RECOVERY_THRESHOLD = 100


def check_port(ip, port):
    try:
        with telnetlib.Telnet(ip, port, timeout=3) as telnet:
            logging.info(f"成功连接到 {ip}:{port}")
            return True
    except Exception as e:
        logging.warning(f"连接到 {ip}:{port} 失败")
        return False


def update_health_score(node, ip, port, success):
    # 根据探测结果更新健康值
    if success:
        health_scores[node][ip]["health_score"] += 1
    else:
        health_scores[node][ip]["health_score"] -= 1

    # 确保健康值在 0 到 100 之间
    health_scores[node][ip]["health_score"] = max(0, min(100, health_scores[node][ip]["health_score"]))

    with open("health_scores.json", "w") as f:
        json.dump(health_scores, f)


def main():
    nodes = {
        "Hong Kong": ["103.84.110.102"],
        "Moscow": ["185.143.220.131"],
        "Latvia": ["193.124.22.221"],
        "Inner Mongolia": ["211.101.233.165"],
    }
    ports = [80, 443, 10085, 15003, 13001, 13002, 13003, 13004, 13005]
    with concurrent.futures.ThreadPoolExecutor() as executor:
        tasks = []
        for node, ips in nodes.items():
            for ip in ips:
                for port in ports:
                    task = executor.submit(check_port, ip, port)
                    tasks.append((task, node, ip, port))
        concurrent.futures.wait([task[0] for task in tasks])

    # 更新健康值
    for task in tasks:
        success = task[0].result()
        node = task[1]
        ip = task[2]
        port = task[3]
        if ip not in health_scores.get(node, {}):
            health_scores[node][ip] = {"health_score": 0}
        update_health_score(node, ip, port, success)

    # 检查所有端口是否成功
    for node, ips in nodes.items():
        for ip in ips:
            for port in ports:
                if is_unhealthy(node, ip, port):
                    # 检查机器是否处于上架状态
                    # 处于上架状态进行下架操作
                    # 处于下架状态不作操作
                    # 进行切换逻辑
                    message="""
                    节点异常需要进行切换
                    检查机器是否处于上架状态
                    处于上架状态进行下架操作
                    处于下架状态不作操作
                    进行切换逻辑
                    """
                    logging.info(message)
                    pass
                else:
                    # 检查机器是否处于上架状态
                    # 处于上架状态不作操作
                    # 处于下架状态恢复上架操作
                    message="""
                    节点正常需要进行恢复
                    检查机器是否处于上架状态
                    处于上架状态进行下架操作
                    处于下架状态不作操作
                    进行切换逻辑
                    """
                    logging.info(message)
                    pass


def is_unhealthy(node, ip, port):
    health_score = health_scores[node][ip]["health_score"]

    # 检查健康得分是否低于健康阈值
    return health_score < HEALTH_THRESHOLD


if __name__ == "__main__":
    start_time = int(time.time())
    main()
    end_time = int(time.time())
    logging.info(f"运行时间：{end_time - start_time}s")
    logging.info(f"各节点健康状态：{health_scores}s")
