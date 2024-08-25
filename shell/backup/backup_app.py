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
    os.system("cp /wwwroot/go/go-admin/*.yaml /hs-app-backup/wwwroot/go/go-admin/")
    os.system("cp /wwwroot/go/go-admin/go-admin /hs-app-backup/wwwroot/go/go-admin/")
    os.system("cp /wwwroot/go/go-admin/*.sh /hs-app-backup/wwwroot/go/go-admin/")

    # api
    os.system("cp -rf /wwwroot/go/go-api/config/* /hs-app-backup/wwwroot/go/go-api/config/")
    os.system("cp /wwwroot/go/go-api/*.yaml /hs-app-backup/wwwroot/go/go-api/")
    os.system("cp -rf /wwwroot/go/go-api/geo/* /hs-app-backup/wwwroot/go/go-api/geo/")
    os.system("cp /wwwroot/go/go-api/go-api /hs-app-backup/wwwroot/go/go-api/")
    os.system("cp /wwwroot/go/go-api/speedctl /hs-app-backup/wwwroot/go/go-api/")
    os.system("cp /wwwroot/go/go-api/*.sh /hs-app-backup/wwwroot/go/go-api/")

    # fly
    os.system("cp -rf /wwwroot/go/go-fly/config/* /hs-app-backup/wwwroot/go/go-fly/config/")
    os.system("cp /wwwroot/go/go-fly/go-fly /hs-app-backup/wwwroot/go/go-fly/")
    os.system("cp -rf /wwwroot/go/go-fly/static/* /hs-app-backup/wwwroot/go/go-fly/static/")
    os.system("cp /wwwroot/go/go-fly/*.sh /hs-app-backup/wwwroot/go/go-fly/")

    # job
    os.system("cp /wwwroot/go/go-job/*.yaml /hs-app-backup/wwwroot/go/go-job/")
    os.system("cp /wwwroot/go/go-job/go-job /hs-app-backup/wwwroot/go/go-job/")
    os.system("cp /wwwroot/go/go-job/*.sh /hs-app-backup/wwwroot/go/go-job/")

    # upload
    os.system("cp /wwwroot/go/go-upload/*.yaml /hs-app-backup/wwwroot/go/go-upload/")
    os.system("cp /wwwroot/go/go-upload/go-upload /hs-app-backup/wwwroot/go/go-upload/")
    os.system("cp -rf /wwwroot/go/go-upload/public/* /hs-app-backup/wwwroot/go/go-upload/public/")
    os.system("cp /wwwroot/go/go-upload/*.sh /hs-app-backup/wwwroot/go/go-upload/")

    # h5
    os.system("cp -rf /wwwroot/h5/* /hs-app-backup/wwwroot/h5/")

    # shell
    os.system("cp /shell/*.sh /hs-app-backup/shell/")
    os.system("cp /shell/*.py /hs-app-backup/shell/")
    os.system("cp -rf /shell/node_check/* /hs-app-backup/shell/node_check/")
    os.system("cp -rf /shell/report/* /hs-app-backup/shell/report/")

def tar_app_dir():
    timestamp = subprocess.run(["date", "+%Y%m%d%H%M%S"], capture_output=True, text=True).stdout.strip()
    tar_file_name = f"./hs-app-backup-{timestamp}.tar.gz"
    os.system("tar -czf %s /hs-app-backup/" % tar_file_name)
    return tar_file_name

def execute_cmd(command):
    result = subprocess.run(command, shell=True, check=True)
    print(result)

def run():
    create_app_dir()
    cp_app_files()
    tar_file_name = tar_app_dir()
    execute_cmd("scp %s root@185.22.152.47:/root/workdir/" % tar_file_name)


if __name__ == '__main__':
    run()



# scp hs-app-backup-20240821195914.tar.gz root@45.251.243.140:/root/backup
# scp speed-2024-08-21_20-12-05.sql root@45.251.243.140:/root/backup
# mysqldump -u root -p --no-data --databases speed go_fly2 speed_report > all_databases_structure.sql