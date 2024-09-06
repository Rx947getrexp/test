import os
import time
import subprocess
import logging
import fcntl

def setup_logging(log_file):
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        handlers=[
            logging.FileHandler(log_file),
            logging.StreamHandler()
        ]
    )

def is_nginx_running():
    try:
        status = subprocess.run(["systemctl", "is-active", "nginx"], capture_output=True, text=True)
        return status.stdout.strip() == "active"
    except Exception as e:
        logging.error(f"检查Nginx状态时出错: {e}")
        return False

def start_nginx():
    try:
        subprocess.run(["systemctl", "start", "nginx"], check=True)
        logging.info("Nginx服务已启动.")
    except Exception as e:
        logging.error(f"启动Nginx时出错: {e}")

def monitor_nginx():
    while True:
        if not is_nginx_running():
            logging.warning("检测到Nginx已停止，正在重新启动...")
            start_nginx()
        else:
            logging.info("Nginx正在运行.")
        time.sleep(3)

def prevent_multiple_instances(lock_file):
    try:
        lock_fd = os.open(lock_file, os.O_CREAT | os.O_RDWR)
        fcntl.lockf(lock_fd, fcntl.LOCK_EX | fcntl.LOCK_NB)
        return lock_fd
    except IOError:
        logging.error("另一个实例正在运行，退出...")
        exit(1)

def release_lock(lock_fd, lock_file):
    try:
        os.close(lock_fd)
        os.remove(lock_file)
    except Exception as e:
        logging.error(f"释放锁文件时出错: {e}")

if __name__ == "__main__":
    lock_file = "/tmp/nginx_monitor.lock"
    log_file = "/shell/log/nginx_monitor.log"
    setup_logging(log_file)
    lock_fd = prevent_multiple_instances(lock_file)

    try:
        monitor_nginx()
    finally:
        release_lock(lock_fd, lock_file)
