<!-- TOC depthFrom:1 depthTo:2 withLinks:1 updateOnSave:1 orderedList:0 -->

# 目录

- [通用说明](#通用说明)
  - [筛选器 (filter)](#筛选器-filter)
  - [筛选器结构示例](#筛选器结构示例)
  - [query 运算符](#query-运算符)
  - [排序](#排序)
  - [可用于排序的属性](#可用于排序的属性)
  - [可用于分组的属性](#可用于分组的属性)
  - [任务列表项 (task_view_entry)](#任务列表项-taskviewentry)
- [API 说明](#api-说明)
  - [1. 获取全局筛选器列表](#1-获取全局筛选器列表)
  - [2. 获取某个过滤条件下的任务列表](#2-获取某个过滤条件下的任务列表)
  - [3. 任务列表任务导出](#3-任务列表任务导出)
  - [4. 获取项目筛选器列表](#4-获取项目筛选器列表)


<!-- /TOC -->

# 通用说明

## 筛选器 (filter)

| 参数名           | 值类型 | 取值范围 | 取值例子 | 说明                         |
|------------------|--------|----------|----------|------------------------------|
| uuid             | string | len=8    |          | uuid                         |
| name             | string |          |          | 创建者名称                   |
| owner            | string | len=8    |          | 创建者的用户id               |
| fixed            | bool   |          |          | 是否固有筛选器               |
| create_time      | int64  |          |          | 创建时间                     |
| query            | object |          |          | 查询条件，参考下方说明       |
| sort             | array  |          |          | 排序，参考下方说明           |
| with_boards      | bool  |          |          | 是否划分看板，如果划分看板，结果会变成 task_view_group_for_board          |
| boards           | array  |          |          | 看板配置          |
| include_subtasks | bool   |          |          | 是否包含子任务, 默认为 false |
| include_status_uuid | bool   |          |          | 是否包含任务状态, 默认为 false |
| include_issue_type | bool   |          |          | 是否包含任务类型, 默认为 false |
| include_project_uuid | bool   |          |          | 是否包含项目uuid, 默认为 false |
| is_show_derive    | bool   |          |          | 是否显示派生, 默认为 false |

## 筛选器结构示例

```json
{
    "uuid":"2FWsr6Xz",
    "name":"项目A、B内，创建者或者负责人是用户C，创建于11月的，bug状态不是'已回归', 包含附件的未完成任务，按优先级和创建时间排序",
    "owner":"6yQWMeg5",
    "create_time":1480550400000,
    "query":{
        "must":[
            {
                "in":{
                    "project_uuid":[
                        "SHArS5ZK",
                        "KOarS5LC"
                    ]
                }
            },
            {
                "match":{
                    "field_values.field001":"你好"
                }
            },
            {
                "should":[
                    {
                        "equal":{
                            "owner":"6yQWMeg5"
                        }
                    },
                    {
                        "equal":{
                            "assign":"6yQWMeg5"
                        }
                    }
                ]
            },
            {
                "range":{
                    "create_time":{
                        "gte":1477958400000,
                        "lt":1480550400000
                    }
                }
            },
            {
                "must_not":[
                    {
                        "equal":{
                            "field_values.3cRiVo4q":"5UiMxLCR"
                        }
                    }
                ]
            },
            {
                "equal":{
                    "status":1
                }
            },
            {
                "exists":"attachments"
            }
        ]
    },
    "sort":[
        {
            "priority":{
                "order":"asc"
            }
        },
        {
            "create_time":{
                "order":"desc"
            }
        }
    ],
    "with_boards":true,
    "boards":[
        {
            "uuid":"12345678",
            "name":"123311312",
            "status_uuids":[
                "123123123"
            ]
        }
    ],
    "include_subtasks":true,
    "is_show_derive":true
}
```

## query 运算符

运算符分为两种，一种是 `must`, `should`, `must_not` 这样的布尔运算，里面可以嵌套其它运算符，
另一种则是 `equal`, `in`, `range` 这样的条件运算，里面不能再嵌套其它运算符

另外，query 内只能有一个根运算符，嵌套的表达式不能超过 5 层

| 运算符     | 能否嵌套 | 说明                                                                                                                                                |
|:-----------|:---------|:----------------------------------------------------------------------------------------------------------------------------------------------------|
| must       | T        | and 运算，当数组内的所有条件都满足时为 true                                                                                                         |
| should     | T        | or 运算，当数组内任意一个条件满足时为 true                                                                                                          |
| must_not   | T        | and-not 运算，当数组内任意一个条件满足时为 false                                                                                                    |
| equal      | F        | 当存在字段，且字段的值等于指定的值时为 true                                                                                                         |
| not_equal  | F        | 当存在字段，且字段的值不等于指定的值时为 true                                                                                                       |
| in         | F        | 当存在字段，且字段的值为数组内的其中一个时为 true                                                                                                   |
| not_in     | F        | 当存在字段，且字段的值不是数组内的任何一个时为 true                                                                                                 |
| range      | F        | 当存在字段，且字段的值在范围内时为 true，可用的参数包括 'gt', 'gte', 'lt' 和 'lte'，分别表示 '>', '>=', '<' 和 '<='，仅能用在数字或者日期类型的属性 |
| date_range | F        | 当存在字段，且字段的值符合指定条件时为 true，仅能用在日期类型的属性                                                                                 |
| exists     | F        | 当存在字段，且字段的值不等于 null 或者空数组时为 true                                                                                               |

### range 参数列表

| 参数名    | 类型                 | 说明                                                  |
|:----------|:---------------------|:------------------------------------------------------|
| gt        | number,relative_date | 大于                                                  |
| gte       | number,relative_date | 大于等于                                              |
| lt        | number,relative_date | 小于                                                  |
| lte       | number,relative_date | 小于等于                                              |
| time_zone | string               | 进行日期比较时所用的时区，使用标准格式，比如 '+08:00' |

`relative_date` 类型必须满足以下正则表达式

```regex
^now(/d)?(\s*([+-])\s*(([1-9][0-9]*)d)?(([1-9][0-9]*)h)?(([1-9][0-9]*)m)?(([1-9][0-9]*)s)?)?$
```

使用方法可以部分参考[这里](https://www.elastic.co/guide/en/elasticsearch/reference/2.4/common-options.html#date-math)

`time_zone` 必须满足以下正则表达式

```regex
^([+-])([01][0-9]|2[0-3]):([0-5][0-9])$
```

### date_range 参数列表

| 参数名    | 类型   | 说明                                                     |
|:----------|:-------|:---------------------------------------------------------|
| equal     | date,relative_date   | 等于，即大于等于某个日期的 0 点，并且小于下一天的 0 点   |
| not_equal | date,relative_date   | 不等于，即小于某个日期的 0 点，或者大于等于下一天的 0 点 |
| gte       | date,relative_date   | 大于等于，即大于等于某个日期的 0 点                      |
| lte       | date,relative_date   | 小于等于，即小于下一天的 0 点             |
| lt        | date,relative_date   | 小于                                  |
| lte       | date,relative_date   | 小于等于                             |
| time_zone | string | 进行日期比较时所用的时区，使用标准格式，比如 '+08:00'    |

`date` 类型必须满足 ISO 8601 格式中的日期部分格式，即 `yyyy-MM-dd`

`relative_date` 类型必须满足以下正则表达式

```regex
^today\\s*(([+-])\\s*(([1-9][0-9]*)d)?)?$
```

例如

```regex
today+1d  today-1d
```

## 可作为条件的属性

| 属性名                    | 类型         | 说明                   |
|:--------------------------|:-------------|:-----------------------|
| create_time               | int64        | 创建时间，微秒         |
| server_update_stamp       | int64        | 修改时间，微秒         |
| priority                  | string       | 优先级 uuid            |
| owner                     | string       | 任务创建者 uuid        |
| assign                    | string       | 任务负责人 uuid        |
| project_uuid              | string       | 项目 uuid              |
| issue_type_uuid           | string       | 任务类型 uuid          |
| sprint_uuid               | string       | 迭代 uuid              |
| status_uuid               | string       | 任务状态 uuid          |
| deadline                  | int64        | 截止日期，秒           |
| field_values.{field_uuid} | string,int64 | 自定义属性或者固有属性 |

## 排序

和 SQL 语句类似，按照 `sort` 数组里的顺序依次进行排序，`asc` 表示升序，`desc` 表示
降序，默认使用的顺序如下

```json
{
    "sort": [
        {
            "create_time": {
                "order": "desc"
            }
        }
    ]
}
```

## 可用于排序的属性

| 属性名                    | 说明                                                           |
|:--------------------------|:---------------------------------------------------------------|
| create_time               | 创建时间                                                       |
| server_update_stamp       | 更新时间                                                       |
| deadline                  | 截止日期                                                       |
| field_values.{field_uuid} | 日期或者时间类型的自定义属性，以上固有属性也可以用这个方式表示 |
| kanban_position           | 看板中的自定义排序，如果未更改过，则按照创建时间排序                   | 

## 可用于分组的属性

| 属性名                    | 说明                                                                 |
|:--------------------------|:---------------------------------------------------------------------|
| assign                    | 任务创建者                                                           |
| owner                     | 任务负责人                                                           |
| status_uuid               | 任务状态                                                             |
| status_category           | 任务状态分类（暂未实现）                                             |
| issue_type_uuid           | 任务类型                                                             |
| priority                  | 任务优先级                                                           |
| sprint_uuid               | 迭代                                                                 |
| project_uuid              | 项目                                                                 |
| field_values.{field_uuid} | 单选，迭代或者文本类型的自定义属性，以上固有属性也可以用这个方式表示 |
| kanban                    | 其实就是根据 status_uuid 来分组。不同的是，看板模式下，分组间有自定义排序      |

## 任务列表分组（task_view_group）

| 参数名  | 值类型 | 取值范围 | 说明                                             |
|---------|--------|----------|--------------------------------------------------|
| key     | string |          | 分组唯一标识                                     |
| name    | string |          | 分组名称。例如，按单选项分组，这个值就是选项名称 |
| total   | int64  |          | 分组下的任务总数                                 |
| count   | int64  |          | 分组下本次取出的任务数量                         |
| entries | array  |          | task_view_entry 列表                             |

其中 key 有以下几种模式

| 模式       | 说明                                         |
|------------|----------------------------------------------|
| n-{value}  | 一般分组，分组下任务该属性的值等于 value     |
| s-null     | 分组下任务该属性的值等于未设置               |
| s-no-field | 分组下任务没有该属性，在全局筛选器中比较常见 |
| s-no-group | 没有填写 group_by，当前分组是唯一分组        |

## 任务列表看板分组（task_view_group_for_board）

| 参数名  | 值类型 | 取值范围 | 说明                                             |
|---------|--------|----------|--------------------------------------------------|
| key     | string |          | 分组唯一标识                                     |
| name    | string |          | 分组名称。例如，按单选项分组，这个值就是选项名称 |
| total   | int64  |          | 分组下的任务总数                                 |
| count   | int64  |          | 分组下本次取出的任务数量                         |
| boards  | array  |          | task_view_board 列表                             |

其中 key 的规则 与 task_view_group 相同

## 任务列表看板（task_view_board）

| 参数名  | 值类型 | 取值范围 | 说明                                             |
|---------|--------|----------|--------------------------------------------------|
| key     | string |          | 看板uuid                                    |
| name    | string |          | 看板名称。 |
| total   | int64  |          | 看板下的任务总数                                 |
| count   | int64  |          | 看板下的任务总数                         |
| entries  | array  |          | task_view_entry 列表                             |

## 任务列表项 (task_view_entry)

| 参数名              | 值类型 | 取值范围 | 取值例子 | 说明       |
|---------------------|--------|----------|----------|------------|
| uuid                | string | len=16   |          | 任务 uuid  |
| server_update_stamp | int64  |          |          | 任务更新戳 |

# API 说明

## 1. 获取全局筛选器列表

列出用户在当前团队下能使用的所有筛选器

### URL

https://your-host-name/project/api/project/team/:teamUUID/filters

### HTTP Method

GET

### 传值方式

URL

### 参数列表

无

### 返回JSON

| JSON键名 | 值类型 | 取值范围 | 取值例子 | 说明                      |
|:---------|:-------|:---------|:---------|:--------------------------|
| filters  | Array  |          | []       | Filter 列表, 参考通用说明 |

### 请求体示例

```curl
curl -X GET \
  https://your-host-name/project/api/project/team/3pDzCwAe/filters \
  -H 'Content-Type: application/json' \
  -H 'Ones-Auth-Token: WuX0XphkP7370DjtlTq3TrXQ0VdGWmQmDFrGEpj3HNvevWbLgH5TNPycbbzvSYhT' \
  -H 'Ones-User-Id: DU6krHBN' \
  -H 'cache-control: no-cache'
```

### 返回体示例

```json
{
    "filters": [
        {
            "uuid": "ft-t-001",
            "team_uuid": "3pDzCwAe",
            "name": "我负责的任务",
            "owner": "DU6krHBN",
            "create_time": 1472688000000000,
            "fixed": true,
            "query": {
                "must": [
                    {
                        "in": {
                            "field_values.field004": [
                                "DU6krHBN"
                            ]
                        }
                    }
                ]
            },
            "sort": null,
            "include_subtasks": true
        },
        {
            "uuid": "ft-t-002",
            "team_uuid": "3pDzCwAe",
            "name": "我关注的任务",
            "owner": "DU6krHBN",
            "create_time": 1472688000000000,
            "fixed": true,
            "query": {
                "must": [
                    {
                        "in": {
                            "field_values.field008": [
                                "DU6krHBN"
                            ]
                        }
                    }
                ]
            },
            "sort": null,
            "include_subtasks": true
        }
    ]
}
```

## 2. 获取某个过滤条件下的任务列表

### URL

https://your-host-name/project/api/project/team/:teamUUID/filters/peek

### HTTP Method

GET, POST

### 是否需要登录

是

### 传值方式

JSON

### 参数列表

| 参数名           | 值类型 | 取值范围 | 是否必填 | 说明                             |
|------------------|--------|----------|----------|----------------------------------|
| query            | object |          | T        | 使用的过滤条件，参考通用说明     |
| sort             | array  |          | F        | 使用的排序方式，参考通用说明     |
| group_by         | string |          | F        | 是否按某个字段分组，参考通用说明 |
| include_subtasks | bool   |          | F        | 是否包含子任务                   |

### 返回JSON

| 参数名 | 值类型 | 取值范围 | 说明                                   |
|--------|--------|----------|----------------------------------------|
| groups | array  |          | task_view_group 或 task_view_group_for_board 列表，参考顶端相关说明 |
| total  | int    |          | 符合过滤条件的所有任务数量             |
| count  | int    |          | 本次取出的任务数量                     |

### 请求体示例

```curl
curl -X POST \
  https://your-host-name/project/api/project/team/3pDzCwAe/filters/peek \
  -H 'Content-Type: application/json' \
  -H 'Ones-Auth-Token: WuX0XphkP7370DjtlTq3TrXQ0VdGWmQmDFrGEpj3HNvevWbLgH5TNPycbbzvSYhT' \
  -H 'Ones-User-Id: DU6krHBN' \
  -H 'cache-control: no-cache' \
  -d '{
    "with_boards": false,
    "boards": null,
    "query": {
        "must": [
            {
                "must": [
                    {
                        "in": {
                            "field_values.field006": [
                                "DU6krHBNJEEeoG8G"
                            ]
                        }
                    },
                    {
                        "in": {
                            "field_values.field007": [
                                "GLLfcQxq"
                            ]
                        }
                    }
                ]
            },
            {
                "should": [
                    {
                        "must": [
                            {
                                "in": {
                                    "field_values.field004": [
                                        "RGzJnspW",
                                        "DU6krHBN"
                                    ]
                                }
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "group_by": "",
    "sort": [
        {
            "field_values.field009": {
                "order": "desc"
            }
        }
    ],
    "include_subtasks": true,
    "include_status_uuid": false,
    "include_issue_type": false,
    "include_project_uuid": false,
    "is_show_derive": false
}'
```

### 返回体示例

```json
{
    "groups": [
        {
            "key": "s-no-group",
            "name": "",
            "total": 1,
            "count": 1,
            "entries": [
                {
                    "uuid": "DU6krHBNvC6QJFXI",
                    "path": "DU6krHBNvC6QJFXI",
                    "server_update_stamp": 1566203217853808,
                    "is_derive": false
                }
            ]
        }
    ],
    "total": 1,
    "count": 1
}
```

## 3. 任务列表任务导出

### URL

https://your-host-name/project/api/project/team/:teamUUID/filters/export

### HTTP Method

POST

### 是否需要登录

是

### 传值方式

JSON

### 参数列表

| 参数名           | 值类型 | 允许空值 | 取值范围 | 默认值 | 取值例子 | 说明                                |
|:-----------------|:-------|:---------|:---------|:-------|:---------|:------------------------------------|
| name             | string | F        |          |        |          | 列表名                              |
| query            | object | F        |          |        |          | filter 使用的过滤条件，参考通用说明 |
| sort             | object | T        |          |        |          | filter 使用的排序方式，参考通用说明 |
| group_by         | string | T        |          |        |          | 是否按某个字段分组，参考通用说明    |
| include_subtasks | string | T        |          |        |          | 是否包含子任务                      |

### 返回的HTTP status code

状态码|说明
:----|:----
200|

### 请求体参考

```curl
curl -X POST \
  https://your-host-name/project/api/project/team/3pDzCwAe/filters/export \
  -H 'Content-Type: application/json' \
  -H 'Ones-Auth-Token: WuX0XphkP7370DjtlTq3TrXQ0VdGWmQmDFrGEpj3HNvevWbLgH5TNPycbbzvSYhT' \
  -H 'Ones-User-Id: DU6krHBN' \
  -H 'cache-control: no-cache' \
  -d '{
    "name": "test-任务-全部任务.csv",
    "query": {
        "must": [
            {
                "must": [
                    {
                        "in": {
                            "field_values.field006": [
                                "DU6krHBNJEEeoG8G"
                            ]
                        }
                    },
                    {
                        "in": {
                            "field_values.field007": [
                                "GLLfcQxq"
                            ]
                        }
                    }
                ]
            },
            {
                "should": [
                    {
                        "must": [
                            {
                                "in": {
                                    "field_values.field004": [
                                        "DU6krHBN"
                                    ]
                                }
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "include_subtasks": true,
    "group_by": "",
    "sort": [
        {
            "field_values.field009": {
                "order": "desc"
            }
        }
    ]
}'
```

### 返回体示例

```text
任务ID,任务标题,任务类型,任务所属项目,任务负责人,任务创建者,任务状态,任务描述,任务关注者,所属迭代,创建时间,最后更新时间,优先级,截止日期,是否子任务,父任务,预估工时,已登记工时,剩余工时
#47,test_task,任务,test,会飞的犀牛长,会飞的犀牛长,未开始,test task,会飞的犀牛长,,2019-08-19 16:26:57,2019-08-19 16:26:57,普通,,否,,,0,
```

## 4. 获取项目筛选器列表

列出项目下能使用的所有 filter

### URL

https://your-host-name/project/api/project/team/:teamUUID/project/:projectUUID/filters

### HTTP Method

GET

### 传值方式

URL

### 参数列表

无

### 返回的HTTP status code

状态码|说明
:----|:----
200|
401|
500|

### 返回JSON

| JSON键名 | 值类型 | 取值范围 | 取值例子 | 说明                      |
|:---------|:-------|:---------|:---------|:--------------------------|
| filters  | Array  |          | []       | Filter 列表, 参考通用说明 |

### 请求体示例

```curl
curl -X GET \
  https://your-host-name/project/api/project/team/3pDzCwAe/project/DU6krHBNJEEeoG8G/filters \
  -H 'Content-Type: application/json' \
  -H 'Ones-Auth-Token: WuX0XphkP7370DjtlTq3TrXQ0VdGWmQmDFrGEpj3HNvevWbLgH5TNPycbbzvSYhT' \
  -H 'Ones-User-Id: DU6krHBN' \
  -H 'cache-control: no-cache'
```

### 返回值示例

```json
{
    "filters": [
        {
            "uuid": "ft-p-001DU6krHBNJEEeoG8G",
            "team_uuid": "3pDzCwAe",
            "project_uuid": "DU6krHBNJEEeoG8G",
            "name": "我负责的任务",
            "owner": "DU6krHBN",
            "fixed": true,
            "create_time": 1472688000000000,
            "query": {
                "must": [
                    {
                        "in": {
                            "field_values.field006": [
                                "DU6krHBNJEEeoG8G"
                            ]
                        }
                    },
                    {
                        "in": {
                            "field_values.field004": [
                                "DU6krHBN"
                            ]
                        }
                    }
                ]
            },
            "sort": null,
            "include_subtasks": true
        },
        {
            "uuid": "ft-p-002DU6krHBNJEEeoG8G",
            "team_uuid": "3pDzCwAe",
            "project_uuid": "DU6krHBNJEEeoG8G",
            "name": "我关注的任务",
            "owner": "DU6krHBN",
            "fixed": true,
            "create_time": 1472688000000000,
            "query": {
                "must": [
                    {
                        "in": {
                            "field_values.field006": [
                                "DU6krHBNJEEeoG8G"
                            ]
                        }
                    },
                    {
                        "in": {
                            "field_values.field008": [
                                "DU6krHBN"
                            ]
                        }
                    }
                ]
            },
            "sort": null,
            "include_subtasks": true
        }
    ]
}
```
