# db_util.py
import pymysql
import yaml

#加载配置文件
def load_db_config(config_path='config.yaml'):
    with open(config_path, 'r') as file:
        config = yaml.safe_load(file)
    return config['databases']

# 根据数据库名称连接到对应数据库
def connect_to_db(db_config, db_name):
    config = db_config[db_name]
    return pymysql.connect(
        host=config['host'],
        port=config['port'],
        user=config['user'],
        password=config['password'],
        database=config['database'],
        #charset=config['charset']
    )

# 获取数据库连接指针
def get_connection(db_name='speed', config_path='config.yaml'):
    db_config = load_db_config(config_path)
    return connect_to_db(db_config, db_name)