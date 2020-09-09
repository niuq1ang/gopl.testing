#!/bin/bash
set -e

configPath="/data/conf/options.json"

port=$(cat $configPath | grep -w "MySQLPort" | awk -F ':' '{print $2}' |awk -F '"' '{print $2}')
host=$(cat $configPath | grep 'MySQLHost' | awk -F ':' '{print $2}' |awk -F '"' '{print $2}')
if [ "$host" == "" ];then
    host="127.0.0.1"
fi
userName=$(cat $configPath | grep 'MySQLUsername' | awk -F ':' '{print $2}' |awk -F '"' '{print $2}')
password=$(cat $configPath | grep 'MySQLPassword' | awk -F ':' '{print $2}' |awk -F '"' '{print $2}')
if [ $password == "{{.MySQLPassword}}" ];then
    password=$(/usr/local/ones-ai-config-proxy/bin/config-proxy -config=/usr/local/ones-ai-config-proxy/conf/config.json)
fi
dbName=$(cat $configPath | grep 'ProjectDBName' | awk -F ':' '{print $2}' | awk -F '"' '{print $2}')
if [ "$dbName" == "" ];then
    dbName="project"
fi


mysql -u${userName} -p${password} -h${host} -P${port} -D${dbName} -e "set names utf8mb4;select a.name as '产品名称', d.user_name as '负责人', if(b.count is not null,'是','否') as '是否新建模块', if(c.count is not null,'是','否') as '是否新建工作项' from product a left join (select product_uuid,count(1) as count from product_module where status=1 group by product_uuid) b on a.uuid=b.product_uuid left join (select product_uuid, count(1) as count from task_product group by product_uuid) c on a.uuid=c.product_uuid left join team_member d on a.assign=d.user_uuid where a.status=1;" | sed -e  "s/\t/,/g" -e "s/NULL/  /g" -e "s/\n/\r\n/g" > ./product_usage_stat.csv