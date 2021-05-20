import pymysql
import json
import os
from IPython import embed

configFile = os.environ['HOME']+"/script/python_script/mysqlManage/config.json"
with open(configFile) as f:
    result = json.load(f)
accounts = result["accounts"]
chandler = [ account for account in accounts if account["nickName"] == "chandler" ][0]
embed()
pymysql.Connect(host=chandler['host'], user=chandler['name'], password=chandler['passWord'], charset="utf8")
