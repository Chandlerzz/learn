import pymysql
import json
import os
from IPython import embed


def connect(nickName):
    configFile = os.environ['HOME']+"/script/python_script/mysqlManage/config.json"
    with open(configFile) as f:
        result = json.load(f)
        accounts = result["accounts"]
    nickName = [ account for account in accounts if account["nickName"] == nickName ][0]
    host = nickName ['host']
    user = nickName ['name']
    password = nickName ['passWord']
    port = int(nickName ['port'])
    conn = pymysql.Connect(host=host, user=user, password = password,port = port, charset="utf8")
    return conn

