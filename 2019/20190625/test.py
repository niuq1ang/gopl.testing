#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
#import urllib2,urllib
import json
import requests
import MySQLdb
import time

#nowTime = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime())
#nowhourTime = time.strftime("%H:%M:%S", time.localtime())
t1 = time.localtime(time.time()-125)
nowTime = time.strftime("%Y-%m-%d %H:%M:%S", t1)
a = nowTime.split()
nowtimeYear = a[0]
nowtimeHour = a[1]

t5 = time.localtime(time.time()-425)
oldTime = time.strftime("%Y-%m-%d %H:%M:%S", t5)
b = oldTime.split()
oldtimeYear = b[0]
oldtimeHour = b[1]

print(nowTime)
print(oldTime)


#接口url,拼接url（'http://127.0.0.1:10080/log/v1/wiki/stat/all?startTime='+bianliangtimeyear+'%20'+biantimehour+'&endTime='+endtimeyear+'%20'+endtimehour+'&isSuccess=-&timeSpanFlag=true&curr=1&nums=10&field=median&order=desc' ）
#Url = 'http://127.0.0.1:10080/log/v1/wiki/stat/all?startTime=2019-01-29%2015:24:08&endTime=2019-01-29%2015:39:08&isSuccess=-&timeSpanFlag=true&curr=1&nums=10&field=median&order=desc'

Url = 'http://127.0.0.1:10080/log/v1/project/stat/all?startTime='+oldtimeYear+'%20'+oldtimeHour+'&endTime='+nowtimeYear+'%20'+nowtimeHour+'&isSuccess=-&timeSpanFlag=true&curr=1&nums=10&field=median&order=desc'

#连接数据库
conn = MySQLdb.connect(host='127.0.0.1', port=10063, user='root', passwd='ones10086', db='ones')
cur = conn.cursor()

req = requests.get(Url)

content = json.loads(req.content)
conData = content['data']
con = json.dumps(conData, indent=4)

#print(con)
print type(conData) 
#print conData[0]['max']
i = 0
for i in range(10):
    x = conData[i]['min'], conData[i]['max'], conData[i]['avg'], conData[i]['median'], conData[i]['uri'], conData[i]['times'], conData[i]['percentile']
    print(x)
    #使用curs()方法操作浮标
    cursor = conn.cursor()
    sql = "INSERT INTO project_log_grafana(min,max,avg,med,uri,times,percentile) VALUES(%s,%s,%s,%s,'%s',%s,'%s')"%(x[0],x[1],x[2],x[3],x[4],x[5],x[6])
    print(sql)
    try:
        cursor.execute(sql)
        conn.commit()
    except:
        conn.rollback()
        conn.close()
    i += 1
    print("see see")