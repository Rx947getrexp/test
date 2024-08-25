# -*- coding: utf-8 -*-
import os
import subprocess
import sys
import time
import fcntl
from datetime import datetime


import os
import glob

def execute_cmd(command):
    # 记录开始时间
    start_time = time.time()
    print(subprocess.run(command, shell=True, check=True))
    # 计算耗时并打印
    elapsed_time = time.time() - start_time
    print(f"调用函数 f1 的耗时为：{elapsed_time:.6f} 秒\n\n")


if __name__ == '__main__':
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://eigrrht.xyz/app-api/get_official_docs""")
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://siaax.xyz/app-api/get_official_docs""")
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://beiyo.xyz/app-api/get_official_docs""")
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://thertee.xyz/app-api/get_official_docs""")
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://weechat.xyz/app-api/get_official_docs""")
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://2yiny.xyz/app-api/get_official_docs""")
    execute_cmd("""curl -X POST -H "Content-Type: application/json" -H "Lang: cn" https://yinyong.xyz/app-api/get_official_docs""")
