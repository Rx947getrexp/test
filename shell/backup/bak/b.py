# -*- coding: utf-8 -*-
import subprocess
import os


def create_app_dir():
    # 清理
    subprocess.run(["rm", "-rf", "/hs-app-backup/shell/"])
    subprocess.run(["rm", "-rf", "/hs-app-backup/wwwroot/"])

    #
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-admin/backup"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-admin/logs"])

    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-api/backup"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-api/config"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-api/geo"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-api/logs"])

    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-fly/backup"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-fly/config"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-fly/logs"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-fly/static"])

    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-job/backup"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-job/logs"])

    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-upload/backup"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-upload/logs"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/go/go-upload/public"])

    subprocess.run(["mkdir", "-p", "/hs-app-backup/wwwroot/h5"])

    subprocess.run(["mkdir", "-p", "/hs-app-backup/shell/database_backup"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/shell/log"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/shell/node_check"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/shell/report/log"])
    subprocess.run(["mkdir", "-p", "/hs-app-backup/shell/sql_backup"])

def cp_app_files():
    # admin
    # subprocess.run(["cp", "/wwwroot/go/go-admin/config.yaml", "/hs-app-backup/wwwroot/go/go-admin/"])
    subprocess.run(["cp", "/wwwroot/go/go-admin/*.yaml", "/hs-app-backup/wwwroot/go/go-admin/"])
    subprocess.run(["cp", "/wwwroot/go/go-admin/go-admin", "/hs-app-backup/wwwroot/go/go-admin/"])
    subprocess.run(["cp", "/wwwroot/go/go-admin/*.sh", "/hs-app-backup/wwwroot/go/go-admin/"])
    # subprocess.run(["cp", "/wwwroot/go/go-admin/restart.sh", "/hs-app-backup/wwwroot/go/go-admin/"])
    # subprocess.run(["cp", "/wwwroot/go/go-admin/update.sh", "/hs-app-backup/wwwroot/go/go-admin/"])

    # api
    os.system("cp -rf /wwwroot/go/go-api/config/* /hs-app-backup/wwwroot/go/go-api/config/")
    # subprocess.run(["cp", "/wwwroot/go/go-api/config.yaml", "/hs-app-backup/wwwroot/go/go-api/"])
    subprocess.run(["cp", "/wwwroot/go/go-api/*.yaml", "/hs-app-backup/wwwroot/go/go-api/"])
    os.system("cp -rf /wwwroot/go/go-api/geo/* /hs-app-backup/wwwroot/go/go-api/geo/")
    subprocess.run(["cp", "/wwwroot/go/go-api/go-api", "/hs-app-backup/wwwroot/go/go-api/"])
    subprocess.run(["cp", "/wwwroot/go/go-api/speedctl", "/hs-app-backup/wwwroot/go/go-api/"])
    subprocess.run(["cp", "/wwwroot/go/go-api/*.sh", "/hs-app-backup/wwwroot/go/go-api/"])
    # subprocess.run(["cp", "/wwwroot/go/go-api/restart.sh", "/hs-app-backup/wwwroot/go/go-api/"])
    # subprocess.run(["cp", "/wwwroot/go/go-api/update.sh", "/hs-app-backup/wwwroot/go/go-api/"])

    # fly
    os.system("cp -rf /wwwroot/go/go-fly/config/* /hs-app-backup/wwwroot/go/go-fly/config/")
    subprocess.run(["cp", "/wwwroot/go/go-fly/go-fly", "/hs-app-backup/wwwroot/go/go-fly/"])
    os.system("cp -rf /wwwroot/go/go-fly/static/* /hs-app-backup/wwwroot/go/go-fly/static/")
    subprocess.run(["cp", "/wwwroot/go/go-fly/*.sh", "/hs-app-backup/wwwroot/go/go-fly/"])
    # subprocess.run(["cp", "/wwwroot/go/go-fly/restart.sh", "/hs-app-backup/wwwroot/go/go-fly/"])
    # subprocess.run(["cp", "/wwwroot/go/go-fly/update.sh", "/hs-app-backup/wwwroot/go/go-fly/"])

    # job
    # subprocess.run(["cp", "/wwwroot/go/go-job/config.yaml", "/hs-app-backup/wwwroot/go/go-job/"])
    subprocess.run(["cp", "/wwwroot/go/go-job/*.yaml", "/hs-app-backup/wwwroot/go/go-job/"])
    subprocess.run(["cp", "/wwwroot/go/go-job/go-job", "/hs-app-backup/wwwroot/go/go-job/"])
    subprocess.run(["cp", "/wwwroot/go/go-job/*.sh", "/hs-app-backup/wwwroot/go/go-job/"])
    # subprocess.run(["cp", "/wwwroot/go/go-job/restart.sh", "/hs-app-backup/wwwroot/go/go-job/"])
    # subprocess.run(["cp", "/wwwroot/go/go-job/update.sh", "/hs-app-backup/wwwroot/go/go-job/"])

    # upload
    # subprocess.run(["cp", "/wwwroot/go/go-upload/config.yaml", "/hs-app-backup/wwwroot/go/go-upload/"])
    subprocess.run(["cp", "/wwwroot/go/go-upload/*.yaml", "/hs-app-backup/wwwroot/go/go-upload/"])
    subprocess.run(["cp", "/wwwroot/go/go-upload/go-upload", "/hs-app-backup/wwwroot/go/go-upload/"])
    os.system("cp -rf /wwwroot/go/go-upload/public/* /hs-app-backup/wwwroot/go/go-upload/public/")
    subprocess.run(["cp", "/wwwroot/go/go-upload/*.sh", "/hs-app-backup/wwwroot/go/go-upload/"])
    # subprocess.run(["cp", "/wwwroot/go/go-upload/restart.sh", "/hs-app-backup/wwwroot/go/go-upload/"])
    # subprocess.run(["cp", "/wwwroot/go/go-upload/update.sh", "/hs-app-backup/wwwroot/go/go-upload/"])

    # h5
    os.system("cp -rf /wwwroot/h5/* /hs-app-backup/wwwroot/h5/")

    # shell
    # subprocess.run(["cp", "/shell/auto_delete.sh", "/hs-app-backup/shell/auto_delete.sh"])
    # subprocess.run(["cp", "/shell/database_backup_speed_for_86.py", "/hs-app-backup/shell/"])
    # subprocess.run(["cp", "/shell/database_backup_speed.py", "/hs-app-backup/shell/"])
    # subprocess.run(["cp", "/shell/monitor_api_redis.py", "/hs-app-backup/shell/"])
    # subprocess.run(["cp", "/shell/port.py", "/hs-app-backup/shell/"])
    # subprocess.run(["cp", "/shell/port.sh", "/hs-app-backup/shell/"])
    # subprocess.run(["cp", "/shell/SqlBackUp.sh", "/hs-app-backup/shell/"])
    subprocess.run(["cp", "/shell/*.sh", "/hs-app-backup/shell/"])
    subprocess.run(["cp", "/shell/*.py", "/hs-app-backup/shell/"])
    os.system("cp -rf /shell/node_check/* /hs-app-backup/shell/node_check/")
    os.system("cp -rf /shell/report/* /hs-app-backup/shell/report/")

def tar_app_dir():
    timestamp = subprocess.run(["date", "+%Y%m%d%H%M%S"], capture_output=True, text=True).stdout.strip()
    tar_file_name = f"./hs-app-backup-{timestamp}.tar.gz"
    subprocess.run(["tar", "-czf", tar_file_name, "/hs-app-backup/wwwroot"])

def run():
    create_app_dir()
    cp_app_files()
    tar_app_dir()


if __name__ == '__main__':
    run()