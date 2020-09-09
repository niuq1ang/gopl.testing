# project 数据字典

## batch_task - 批量任务

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| team_uuid | varchar(8) |  | NO |  |
| owner | varchar(8) |  | NO |  |
| job_id | varchar(16) |  | NO | 任务id |
| job_type | varchar(128) | 0 | NO | 批量任务类型(1.批量删除 2. 批量移动) |
| job_status | tinyint(4) | 1 | NO | 批量任务状态(1.等待中 2.进行中 3.中止 4.完成) |
| batch_type | varchar(64) | batch | NO | 批量任务的默认类型（批量，单个操作） |
| status | tinyint(4) | 1 | NO | 任务状态(1.显示 2.用户关闭) |
| start_time | int(11) | 0 | NO | 开始时间 |
| end_time | int(11) | 0 | NO | 结束时间 |
| payload | mediumtext | null | NO | 请求的任务body |
| extra | mediumtext | null | YES | 队列的额外信息 |
| successful | mediumtext | null | YES | 成功的数据 |
| unsuccessful | mediumtext | null | YES | 失败的数据 |
| unprocessed | mediumtext | null | YES | 待处理的数据(通常是用户取消了) |
| successful_count | int(11) | 0 | NO | 成功的数量 |
| unsuccessful_count | int(11) | 0 | NO | 失败的数量 |
| unprocessed_count | int(11) | 0 | NO | 未处理的数量 |

## batch_task_record - 批量操作记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| batch_task_uuid | varchar(8) |  | NO | batch task uuid |
| team_uuid | varchar(8) |  | NO | 团队id |
| number | int(20) | null | NO | 序号 |
| context_type | int(11) | 0 | NO | 上下文类型, 1. import_task(project_task) |
| context_param_1 | varchar(16) |  | NO | 上下文参数1 |
| context_param_2 | varchar(16) |  | NO | 上下文参数2 |
| create_time | bigint(20) | 0 | NO | 创建时间(秒) |
| extra | mediumtext | null | YES | 额外信息 |
| status | tinyint(4) | 0 | NO | 状态(1:正常 2:删除) |

## batch_task_row

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| batch_task_uuid | varchar(8) |  | NO | batch task uuid |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| number | int(11) | null | NO |  |
| context_type | int(11) | null | NO | 上下文类型 |
| context_param_1 | varchar(16) |  | NO | 上下文参数1 |
| context_param_2 | varchar(16) |  | NO | 上下文参数2 |
| create_time | int(11) | null | NO | 创建时间（秒） |
| extra | mediumtext | null | YES | 额外信息 |

## blog_subscribe

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| id | int(11) | null | NO |  |
| email | varchar(128) | null | NO |  |
| create_time | datetime | CURRENT_TIMESTAMP | NO |  |
| ip | varchar(15) | null | YES |  |

## blog_subscribe_count - 博客统计

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| count_value | int(11) | 0 | NO |  |
| id | int(11) | null | NO |  |

## board - board

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| name | varchar(32) | null | NO | 名字 |
| project_uuid | varchar(16) | null | NO | 所属项目uuid |
| position | bigint(20) | null | NO | 顺序(从小到大) |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 更新时间 |

## build_result

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| hash | varchar(48) | null | NO |  |
| build_id | int(11) | null | NO |  |
| display_url | varchar(128) | null | NO |  |
| result | varchar(16) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |

## card - 仪表盘卡片

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| dashboard_uuid | varchar(8) |  | NO | 仪表盘UUID |
| name | varchar(32) |  | NO | 卡片名称 |
| type | int(11) | 0 | NO | 卡片类型 |
| layout_x | int(11) | 0 | NO | x位置 |
| layout_y | int(11) | 0 | NO | y位置 |
| layout_w | int(11) | 0 | NO | 宽度 |
| layout_h | int(11) | 0 | NO | 高度 |
| status | tinyint(4) | 0 | NO | 1:正常;2:删除 |
| config | text | null | NO | 卡片配置 |
| uuid | varchar(8) | null | NO |  |
| space_uuid | varchar(8) | null | NO |  |
| page_uuid | varchar(8) | null | NO |  |
| draft_uuid | varchar(8) | null | NO |  |
| type | int(11) | null | NO |  |
| status | tinyint(4) | null | NO | 1: 正常, 2: 删除 |
| create_time | bigint(20) | null | NO |  |
| config | mediumtext | null | NO | 配置 |

## cas_config - CAS 配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO | 组织UUID |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| cas_server_url | varchar(255) |  | NO | CAS服务器地址 |
| user_name | varchar(255) |  | NO | CAS配置里代表用户姓名的属性 |
| email | varchar(255) |  | NO | CAS配置里代表用户邮箱的属性 |
| user_title | varchar(255) |  | NO | CAS配置里代表用户职位的属性 |
| create_time | bigint(20) | 0 | NO | 创建时间，单位是微秒 |
| update_time | bigint(20) | 0 | NO | 更新时间，单位是微秒 |
| status | tinyint(4) | null | NO | 1:正常 2:删除 |

## cas_config_status - 组织的CAS配置状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO | 组织UUID |
| is_open | tinyint(1) | null | NO | cas配置是否开启 0:关闭 1:开启 |

## commit

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| hash | varchar(48) | null | NO |  |
| revision | bigint(20) | -1 | NO | svn revision |
| owner | varchar(128) | null | NO |  |
| domin | varchar(32) | null | NO |  |
| repository_name | varchar(128) | null | NO |  |
| branch | varchar(128) | null | NO |  |
| url | varchar(1024) | null | NO |  |
| message | text | null | NO |  |
| timestamp | bigint(20) | null | NO |  |
| scm | varchar(3) | null | NO | 代码管理工具 |

## component

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| template_uuid | varchar(8) |  | NO | 模板UUID |
| team_uuid | varchar(8) |  | NO |  |
| project_uuid | varchar(16) |  | NO |  |
| parent_uuid | varchar(8) | null | YES | 组件父级 |
| name | varchar(64) |  | NO | 组件名称 |
| name_pinyin | varchar(256) |  | NO | 组件名称拼音 |
| desc | varchar(256) |  | NO | 组件描述 |
| type | int(11) | null | NO | 组件类型(1. 文件夹 2.内置组件 3.场景特化组件 4. 自定义任务类型组件) |
| create_time | int(11) | null | NO | 创建时间 |
| position | int(11) | 0 | YES | 排序 |
| objects | varchar(512) | [] | YES | 组件源类型/组件源uuid |
| status | tinyint(1) | 1 | YES | 1正常 2已删除 |
| settings | text | null | YES | 功能设置 |
| kanban_setting | text | null | YES | 敏捷看板设置 |
| url_setting | text | null | YES | 自定义链接设置 |

## component_template

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| team_uuid | varchar(8) |  | NO |  |
| name | varchar(64) |  | NO | 组件名称 |
| name_pinyin | varchar(256) |  | NO | 组件名称拼音 |
| desc | varchar(256) |  | NO | 组件描述 |
| type | int(11) | null | NO | 组件类型(1. 文件夹 2.内置组件 3.场景特化组件 4. 自定义任务类型组件) |
| category | int(11) | 1 | YES | 组件分组 |
| status | int(11) | 1 | YES | 状态(1.正常 2.删除) |
| create_time | bigint(20) | null | NO | 创建时间 |
| position | int(11) | 0 | YES | 位置 |
| objects | varchar(512) | [] | YES | 组件源类型/组件源uuid |

## component_view

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) |  | NO |  |
| component_uuid | varchar(8) |  | NO | 组件uuid |
| name | varchar(64) |  | NO | 视图名称 |
| layout | tinyint(1) | 3 | NO | 布局方式(1.表格 2.宽详情 3.窄详情) |
| query | text | null | YES | 筛选方式 |
| condition | text | null | NO | 筛选条件 |
| sort | varchar(256) |  | YES | 排序方式 |
| group_by | varchar(256) |  | YES | 分组方式 |
| create_time | bigint(20) | null | NO |  |
| built_in | tinyint(1) | 1 | NO | 是否默认视图 |
| include_subtasks | tinyint(1) | 0 | NO | 是否包含子任务 |
| table_field_settings | text | null | YES | 表格模式的表头 |
| board_settings | text | null | YES | 看板视图设置 |
| is_fold_all_groups | tinyint(1) | 0 | NO | 是否折叠分组 |
| status | tinyint(1) | 1 | YES | 1正常 2已删除 |
| display_type | tinyint(4) | 0 | YES | 1:平铺展示, 2:子树展示，3:折叠展示 |
| is_show_derive | tinyint(1) | 0 | NO | 是否显示派生 |

## component_view_user

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| owner | varchar(8) |  | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) |  | NO |  |
| component_uuid | varchar(8) |  | NO | 组件uuid |
| name | varchar(64) |  | NO | 视图名称 |
| layout | tinyint(1) | 3 | NO | 布局方式(1.表格 2.宽详情 3.窄详情) |
| query | text | null | YES | 筛选方式 |
| condition | text | null | NO | 筛选条件 |
| sort | varchar(256) |  | YES | 排序方式 |
| group_by | varchar(256) |  | YES | 分组方式 |
| create_time | bigint(20) | null | NO |  |
| built_in | tinyint(1) | 0 | NO | 是否默认视图 |
| include_subtasks | tinyint(1) | 0 | NO | 是否包含子任务 |
| table_field_settings | text | null | YES | 表格模式的表头 |
| board_settings | text | null | YES | 看板视图设置 |
| is_fold_all_groups | tinyint(1) | 0 | NO | 是否折叠分组 |
| status | tinyint(1) | 1 | YES | 1正常 2已删除 |
| project_filter_uuid | varchar(8) |  | YES | 项目下用户的筛选器uuid |
| display_type | tinyint(4) | null | YES | 1:平铺展示, 2:子树展示，3:折叠展示 |
| is_show_derive | tinyint(1) | 0 | NO | 是否显示派生 |

## component_view_user_config

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| component_uuid | varchar(8) | null | NO | 组件uuid |
| owner | varchar(8) | null | NO |  |
| view_uuid | varchar(8) |  | NO | 视图uuid |
| type | tinyint(1) | null | NO | 1.组件视图 2.用户视图 |
| position | int(11) | null | NO | 视图位置 |
| is_show | tinyint(1) | 1 | NO | 1.显示 0.隐藏 |

## converter_config_status - 组织的converter配置状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO | 组织UUID |
| is_open | tinyint(1) | null | NO | 是否开启 0:关闭 1:开启 |

## corp_sync_department - 某些第三方如LDAP没有提供部门sync_id的功能, 这里提供一种自管理sync_id的方式

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| department_key | varchar(255) | null | NO | 第三方登录企业id下标识唯一部门的关键词 |
| sync_id | bigint(20) | null | NO | 同步部门的id |

## create_team_code - 创建团队邮箱验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(32) |  | NO | 验证码 |
| email | varchar(255) |  | NO | 邮箱地址 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| ip | varchar(15) |  | NO | 请求IP地址 |
| status | tinyint(4) | null | NO | 1:正常;2:无效 |
| referrer | text | null | NO | 渠道信息 |
| keyword | text | null | NO | 搜索关键词 |

## crm_permission

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| email | varchar(32) | null | NO |  |

## csm

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| email | varchar(128) | null | NO |  |
| name | varchar(24) | null | NO |  |
| title | varchar(24) | null | NO |  |
| phone | varchar(32) | null | NO |  |
| status | tinyint(4) | -1 | NO |  |

## dashboard - 仪表盘

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| owner | varchar(8) |  | NO | 创建者UUID |
| name | varchar(32) |  | NO | 仪表盘名称 |
| name_pinyin | varchar(256) |  | NO | 仪表盘名称拼音 |
| shared | tinyint(4) | 0 | NO | 是否共享 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| status | tinyint(4) | 0 | NO | 1:正常;2:删除 |

## dashboard_member - 仪表盘成员（被分享者）

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| dashboard_uuid | varchar(8) |  | NO | 仪表盘UUID |
| user_domain_param | varchar(8) |  | NO | 成员域参数 |
| user_domain_type | tinyint(4) | 0 | NO | 成员域类型 |
| position | int(11) | 0 | NO | 位置 |

## dashboard_preference - 仪表盘用户配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| user_uuid | varchar(8) | null | NO | 用户UUID |
| dashboard_uuid | varchar(8) |  | NO | 仪表盘UUID |
| pinned | tinyint(4) | 0 | NO | 是否常用仪表盘 |
| is_default | tinyint(4) | 0 | NO | 是否默认仪表盘 |

## department

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| parent_uuid | varchar(8) |  | YES | 父节点id |
| name | varchar(32) | null | NO | 部门名称 |
| name_pinyin | varchar(256) | null | NO |  |
| next_uuid | varchar(8) |  | YES | 兄弟节点id |
| team_uuid | varchar(8) |  | NO |  |
| sync_id | bigint(20) | null | YES | 同步部门的id |
| corp_uuid | varchar(64) | null | YES | 第三方登录企业id |

## department_member

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO |  |
| department_uuid | varchar(8) | null | YES |  |
| user_uuid | varchar(8) | null | YES |  |

## dingding_sync_relation

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| corp_uuid | varchar(64) |  | NO | 同步企业的id |
| unionid | varchar(128) |  | NO |  |
| corp_name | varchar(255) | null | NO | 第三方企业名称 |
| sync_id | varchar(255) |  | NO | 同步账号的id |
| status | tinyint(4) | -1 | NO | 状态:1:正常 2:删除 |

## domain_info

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| id | int(11) | null | NO |  |
| imap | varchar(512) | null | YES |  |
| smtp | varchar(512) | null | YES |  |
| pop | varchar(512) | null | YES |  |
| mailboxes | varchar(512) | null | YES |  |
| domainMatch | varchar(2048) | null | YES |  |
| mx-match | varchar(64) | null | YES |  |

## email_code - 用户邮箱验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(32) |  | NO | 验证码 |
| user_uuid | varchar(8) | null | NO | 用户ID |
| email | varchar(255) | null | NO | 邮箱地址 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| ip | varchar(15) | null | NO | 请求IP地址 |
| status | tinyint(4) | null | NO | 1:正常;2:无效 |

## exchange_code - 兑换码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(8) | null | NO |  |
| email | varchar(128) | null | NO |  |
| phone | varchar(128) | null | NO |  |
| team_name | varchar(8) | null | NO |  |
| team_uuid | varchar(8) |  | NO |  |
| user_name | varchar(8) | null | NO |  |
| is_pay | tinyint(1) | null | NO | 是否付费 |
| source | varchar(128) | null | NO | 来源 |
| exchange_time | int(11) | null | NO | 团队有效期(天) |
| create_time | bigint(20) | null | NO | 创建时间(秒) |
| status | tinyint(4) | null | NO |  |
| expired_time | bigint(11) | 0 | NO | 兑换码有效期(具体时间的秒数) |
| scale | int(6) | null | NO | 团队成员规模 |
| csm | varchar(128) |  | YES | 负责的CSM |

## field - 全局属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| type | int(11) | 0 | NO | 类型 |
| name | varchar(32) |  | NO | 属性名 |
| name_pinyin | varchar(256) | null | NO |  |
| description | text | null | NO | 属性描述 |
| default_value | text | null | YES | 属性默认值 |
| renderer | int(11) | 0 | NO | 属性显示方式 |
| filter_option | int(11) | 0 | NO | 属性过滤器选项(0:不可过滤,1:可过滤) |
| search_option | int(11) | 0 | NO | 属性搜索选项(0:不可搜索,1:可搜索) |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| built_in | tinyint(1) | 0 | NO | 系统内置属性 |
| step_settings | varchar(256) | null | YES | 属性间隔时间设置 |
| stay_settings | varchar(256) | null | YES | 属性停留时间与次数设置 |
| related_type | int(11) | 0 | YES | 关联类型 |
| related_uuid | varchar(8) | null | YES | 关联的uuid |
| status | tinyint(4) | null | YES | 1:正常;null:删除 |

## field_access_rule - 属性权限规则

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| issue_type_uuid | varchar(8) |  | NO | 任务类型UUID |
| field_uuid | varchar(8) | null | NO | 属性UUID |
| status_uuid | varchar(8) |  | NO | 任务状态UUID |
| user_domain_type | int(11) | 0 | NO | 用户域类型 |
| user_domain_param | varchar(16) |  | NO | 用户域参数 |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| access | int(11) | 0 | NO | 访问权限 |

## field_config - 属性配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| issue_type_uuid | varchar(8) |  | NO | 任务类型UUID |
| field_uuid | varchar(8) | null | NO | 属性UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| default_value | text | null | YES | 属性默认值 |
| renderer | int(11) | 0 | NO | 属性显示方式 |
| filter_option | int(11) | 0 | NO | 属性过滤器选项(0:不可过滤,1:可过滤) |
| search_option | int(11) | 0 | NO | 属性搜索选项(0:不可搜索,1:可搜索) |
| required | tinyint(4) | 0 | NO | 属性是否必填(0:否,1:是) |
| can_modify_required | tinyint(1) | null | NO | 是否可以修改 选填 |
| can_delete | tinyint(1) | null | NO | 是否可以删除 |
| position | int(11) | 0 | NO | 属性位置 |

## field_multi_option_value - 任务多选属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) |  | NO |  |
| field_uuid | varchar(8) |  | NO |  |
| type | int(11) | null | NO | 选项类型 |
| value | text | null | YES | 选项值 |
| position | int(11) | 0 | YES | 选项排序 |
| status | tinyint(4) | 1 | YES | 1: 正常 2: 删除 3: 储存 |

## field_option - 属性选项，用于单选或者多选类型的属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| field_uuid | varchar(8) |  | NO | 属性UUID |
| value | varchar(64) |  | NO | 属性值 |
| default_selected | tinyint(4) | 0 | NO | 是否默认选中 |
| background_color | varchar(9) | null | NO |  |
| color | varchar(9) | null | NO |  |
| position | int(11) | 0 | NO | 排序序号 |
| status | tinyint(4) | null | YES | 1:正常;2:删除 |
| desc | varchar(128) |  | YES | 属性值描述 |

## field_option_config - 属性配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| issue_type_uuid | varchar(8) |  | NO | 属性类型UUID |
| field_uuid | varchar(8) | null | NO | 属性UUID |
| option_uuid | varchar(8) |  | NO | 属性选项UUID |
| position | int(11) | 0 | NO | 属性选项排序序号 |
| default_selected | tinyint(4) | 0 | NO | 属性选项是否默认选中 |

## field_value - 任务属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) |  | NO | 任务UUID |
| field_uuid | varchar(8) |  | NO | 属性UUID |
| type | int(11) | 0 | NO | 值类型 |
| value | text | null | YES | 属性值 |
| number_value | bigint(20) | null | YES | 数字属性值 |
| status | int(11) | 0 | NO | 状态(1:正常,2:删除) |

## field_value_history

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO |  |
| task_uuid | varchar(16) |  | NO |  |
| field_uuid | varchar(8) |  | NO |  |
| type | int(11) | 0 | NO | field type |
| value | varchar(16) |  | NO | field的值，超过16位保存hash值 |
| valid_from | int(11) | null | NO | 更新时间 |
| valid_to | int(11) | null | YES | 失效时间 |

## file

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| hash | varchar(28) | null | NO |  |
| type | int(11) | null | NO |  |
| mime | varchar(255) | null | YES |  |
| size | bigint(20) | -1 | NO |  |
| image_width | int(11) | -1 | NO |  |
| image_height | int(11) | -1 | NO |  |
| exif | varchar(64) |  | NO | 文件 exif 属性 |

## file_resource_view - VIEW

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| reference_type | tinyint(4) | null | NO |  |
| reference_id | varchar(16) | null | YES |  |
| team_uuid | varchar(8) | null | YES |  |
| project_uuid | varchar(16) | null | NO |  |
| owner_uuid | varchar(8) | null | YES |  |
| type | tinyint(4) | -1 | NO |  |
| source | tinyint(4) | 0 | NO |  |
| name | varchar(255) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| modify_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | -1 | NO |  |
| description | varchar(64) | null | NO |  |
| hash | varchar(28) | null | NO |  |
| mime | varchar(255) | null | YES |  |
| size | bigint(20) | -1 | NO |  |
| image_width | int(11) | -1 | NO |  |
| image_height | int(11) | -1 | NO |  |
| exif | varchar(64) |  | NO | 文件 exif 属性 |

## filter

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | YES |  |
| owner | varchar(8) | null | YES |  |
| name | varchar(128) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| query | text | null | NO |  |
| sort | text | null | NO |  |
| status | tinyint(4) | 0 | NO |  |
| include_subtasks | tinyint(4) | 0 | NO | 是否包含子任务 |

## gantt_chart

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | gantt chart uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| name | varchar(32) | null | NO | gantt chart名称 |
| name_pinyin | varchar(256) |  | NO | gantt chart名称拼音 |
| owner | varchar(8) | null | NO | 创建者 uuid |
| create_time | bigint(20) | null | NO | 创建时间 |
| shared | tinyint(4) | 0 | NO | 是否共享 |
| import_config | text | null | YES | 导入配置 |
| status | tinyint(4) | null | YES | 1:正常 2:删除 |
| sync_from_project | tinyint(4) | 0 | NO | 开关：关联甘特数据同步project的变动 |
| sync_to_project | tinyint(4) | 0 | NO | 开关：关联甘特数据变动同步到project |

## gantt_chart_field_value - 甘特图属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| gantt_chart_uuid | varchar(16) | null | NO | 甘特图UUID |
| alias | varchar(32) | null | NO | 别名 |
| type | int(11) | null | NO | 值类型 |
| value | text | null | YES | 属性值 |

## gantt_chart_personal_config

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| gantt_chart_uuid | varchar(8) | null | NO | gantt chart uuid |
| user_uuid | varchar(8) | null | NO | 用户 uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| expand | tinyint(4) | 1 | NO | 0:收起  1: 展开 |
| zooming | tinyint(4) | 0 | NO | 缩放 |
| create_time | bigint(20) | null | NO | 创建时间 |
| status | tinyint(4) | 1 | NO | 1: 正常 2: 删除 |

## gantt_data

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | 甘特图数据 uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| gantt_chart_uuid | varchar(8) | null | NO | gantt chart uuid |
| type | tinyint(4) | 0 | YES | 1:甘特图数据类型, 1:任务 2:分组 3:里程碑 |
| name | varchar(128) | null | NO | 名称 |
| name_pinyin | varchar(512) |  | NO | 名称拼音 |
| owner | varchar(8) | null | NO | 创建者 uuid |
| start_time | bigint(20) | null | NO | 开始时间 |
| end_time | bigint(20) | null | NO | 结束时间 |
| progress | int(11) | 0 | NO | 进度 |
| parent | varchar(8) | null | NO | 父节点 uuid |
| path | varchar(169) | null | NO | 层级关系 |
| position | bigint(20) | 0 | NO | 显示位置 |
| create_time | bigint(20) | null | NO | 创建时间 |
| status | tinyint(4) | null | YES | 1:正常 2:删除 |
| assign | varchar(8) | null | YES | 负责人uuid |
| related_type | tinyint(4) | 0 | NO | 关联数据类型 |
| related_project_uuid | varchar(16) | null | NO | 关联的项目uuid |
| related_sprint_uuid | varchar(8) | null | NO | 关联的迭代uuid |

## important_field - 关键属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| issue_type_uuid | varchar(8) |  | NO | 任务类型UUID |
| field_uuid | varchar(8) | null | NO | 属性UUID |
| position | int(11) | 0 | NO | 属性位置 |

## invitation

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(32) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| inviter_uuid | varchar(8) | null | NO |  |
| email | varchar(255) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | null | NO |  |

## issue_type - 任务类型

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| name | varchar(32) |  | NO | 任务类型名 |
| name_pinyin | varchar(256) | null | NO |  |
| icon | int(11) | 0 | NO | 任务类型图标编号 |
| built_in | tinyint(4) | 0 | NO | 是否系统内置任务类型 |
| default_selected | tinyint(4) | 0 | NO | 是否默认选中 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| status | tinyint(4) | null | YES | 1:正常;null:删除 |
| default_configs | longtext | null | NO | 默认配置 |
| type | tinyint(4) | 0 | YES | 0:标准任务类型, 1:子任务类型 |
| detail_type | tinyint(4) | 0 | YES | issue_type的具体类型: 0:自定义类型, 1:需求, 2:任务, 3:缺陷, 4:子任务, 5:子需求 |

## kanban_status_sort

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) | null | NO |  |
| issue_type_uuid | varchar(8) | null | NO |  |
| status_uuid | varchar(8) | null | NO |  |
| position | bigint(20) | null | NO |  |

## kanban_task_sort

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) | null | NO |  |
| position | bigint(20) | null | NO |  |

## ldap_config - LDAP 配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO | 组织UUID |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| host | varchar(255) |  | NO | LDAP服务器 |
| port | int(11) | -1 | NO | LDAP服务器端口 |
| encryption | varchar(10) |  | NO | 加密方式(plain,ssl,starttls) |
| domain_base | varchar(255) |  | NO | 需要同步的LDAP根节点 |
| user_id | varchar(255) |  | NO | LDAP配置里代表用户唯一ID的属性 |
| user_name | varchar(255) |  | NO | LDAP配置里代表用户姓名的属性 |
| email | varchar(255) |  | NO | LDAP配置里代表用户邮箱的属性 |
| domain_search_user | varchar(255) |  | NO | 执行LDAP搜索时使用的用户账户 |
| domain_search_password | varchar(255) |  | NO | 执行LDAP搜索时使用的用户密码 |
| user_filter | varchar(1024) |  | NO | 用户过滤规则 |
| create_time | bigint(20) | 0 | NO | 创建时间，单位是微秒 |
| update_time | bigint(20) | 0 | NO | 更新时间，单位是微秒 |
| status | tinyint(4) | null | NO | 1:正常 2:删除 |
| domain | varchar(255) |  | NO | AD 需要域的概念 |
| type | int(11) | 1 | NO | LDAP模式 1.LDAP 2.AD |
| way_to_get_mail | int(11) | 1 | NO | 获取邮箱方式 1.按属性值 2.按用户登录名@域名方式 |
| enable_sync | int(1) | 0 | NO | 是否开启同步 |
| department_filter | varchar(1024) |  | NO | 部门过滤规则 |
| avatar | varchar(255) |  | NO | LDAP配置里代表用户头像的属性 |
| user_title | varchar(255) |  | NO | LDAP配置里代表用户职位的属性 |

## ldap_config_status - 组织的LDAP配置状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO | 组织UUID |
| is_open | tinyint(1) | null | NO | ldap配置是否开启 0:关闭 1:开启 |

## lint_issue - lint问题记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(22) |  | NO | uuid |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| hash | varchar(28) |  | NO | 用于跟踪相同issue的哈希 |
| create_time | bigint(20) | null | NO | 首次出现时间（秒） |
| linter | varchar(20) |  | NO | 使用的linter |
| language | varchar(20) |  | NO | 使用的语言 |
| file | varchar(4096) |  | NO | 文件相对路径 |
| line | int(11) | -1 | NO | 文件行号 |
| column | int(11) | -1 | NO | 文件列号，没有列号则为-1 |
| severity | tinyint(4) | -1 | NO | 严重程度(1:ignore 2:info 3:warning 4:error 5:fatal) |
| category | varchar(255) |  | NO | 所属类别，可能为空 |
| rule | varchar(255) |  | NO | 规则名称，可能为空 |
| message | varchar(4096) |  | NO | lint提示信息/描述 |
| author | varchar(40) |  | NO | 问题代码提交者名称 |

## manhour

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | YES |  |
| task_uuid | varchar(16) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| start_time | bigint(20) | 0 | NO |  |
| hours | bigint(20) | 0 | NO |  |
| remark | varchar(1024) | null | NO |  |
| status | tinyint(4) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |

## manhour_report

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| project_uuid | varchar(16) |  | NO | 项目uuid |
| name | varchar(128) |  | NO | 报表名称 |
| owner | varchar(8) | null | NO | 创建者uuid |
| config | text | null | NO | 项目报表配置 |
| create_time | bigint(20) | 0 | NO | 创建时间，秒 |
| update_time | bigint(20) | 0 | NO | 更新时间，秒 |
| fold | tinyint(1) | 0 | NO | 0:展开, 1:折叠 |
| status | tinyint(4) | 0 | NO | 状态(1:正常, 2:删除) |

## message

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(32) | null | NO |  |
| team_uuid | varchar(8) | null | YES |  |
| reference_type | tinyint(4) | -1 | NO |  |
| reference_id | varchar(16) | null | YES |  |
| from_uuid | varchar(16) | null | YES |  |
| to_uuid | varchar(16) | null | YES |  |
| send_time | bigint(20) | 0 | NO |  |
| message | text | null | NO |  |
| type | tinyint(4) | null | NO |  |
| resource_uuid | varchar(8) | null | YES |  |
| subject_type | tinyint(4) | null | NO |  |
| subject_id | varchar(16) | null | YES |  |
| action | tinyint(4) | null | NO |  |
| object_type | tinyint(4) | null | NO |  |
| object_id | varchar(16) | null | YES |  |
| object_name | text | null | NO |  |
| object_attr | tinyint(4) | null | NO |  |
| old_value | text | null | NO |  |
| new_value | text | null | NO |  |
| ext | text | null | NO |  |
| uuid | varchar(8) | null | NO | uuid |
| user_uuid | varchar(8) | null | NO | 发送者 |
| space_uuid | varchar(8) | null | NO | space uuid |
| page_uuid | varchar(8) | null | NO | page uuid |
| send_time | bigint(20) | null | NO | 创建时间 |
| action | tinyint(4) | null | NO | 操作动作 |
| type | tinyint(4) | null | NO | 类型(评论消息,系统消息等) |
| title | varchar(64) | null | NO | 标题 |
| message | varchar(512) | null | NO |  |
| object_attr | tinyint(4) | null | NO |  |
| ext | varchar(512) | null | NO |  |

## mq_failed_message

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| body | text | null | NO |  |
| error | text | null | NO |  |
| create_time | bigint(20) | null | NO |  |

## notice

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| task_uuid | varchar(16) | null | NO |  |
| message_uuid | varchar(8) | null | NO |  |
| status | tinyint(4) | null | NO |  |
| server_update_stamp | bigint(20) | 0 | NO |  |

## notice_config - 任务通知配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | 任务通知配置的uuid |
| status | tinyint(4) | null | NO | 1:正常 2:无效 |
| create_time | bigint(20) | null | NO | 创建时间, 秒 |
| update_time | bigint(20) | null | NO | 更新时间，秒 |

## notice_config_sending_method - 任务通知配置-发送方式

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| notice_config_uuid | varchar(8) | null | NO | 任务通知配置的uuid |
| method | enum('notice_center','email','dingding','wechat') | null | NO | 通知发送的方式（消息中心、邮件、钉钉、企业微信） |
| status | tinyint(4) | null | NO | 1:正常 2:无效 |

## notice_config_user_domain - 接收通知的用户域

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| notice_config_uuid | varchar(8) | null | NO | 任务通知配置的uuid |
| user_domain_type | int(11) | 0 | NO | 用户域类型 |
| user_domain_param | varchar(16) | null | NO | 用户域参数 |
| status | tinyint(4) | null | NO | 1:正常 2:无效 |

## notification_rule

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| team_uuid | varchar(8) |  | NO |  |
| scheme_uuid | varchar(8) |  | NO | 所属通知配置方案uuid |
| create_time | bigint(20) | 0 | NO | 创建时间(秒) |
| user_domain_type | int(11) | 0 | NO | 用户域类型 |
| user_domain_param | varchar(16) |  | NO | 用户域参数 |
| event | int(11) | 0 | NO | 事件编号 |
| email | tinyint(4) | 0 | NO | 邮件选项(1:开启 2:关闭) |
| badge | tinyint(4) | 0 | NO | 红点选项(1:开启 2:关闭) |
| alert | tinyint(4) | 0 | NO | 推送通知选项(1:开启 2:关闭) |
| status | tinyint(4) | 0 | NO | 状态(1:正常 2:删除) |

## notification_scheme

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| team_uuid | varchar(8) |  | NO |  |
| name | varchar(32) |  | NO | 通知配置方案名称 |
| description | text | null | NO | 通知配置方案描述 |
| status | tinyint(4) | 0 | NO | 状态(1:正常 2:删除) |

## ones_app

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(64) | null | NO | 每一个私有部署客户对应一个ones_appid(ones_app_uuid) |
| name | varchar(256) | null | NO | 每一个私有部署客户的名称（组织名、团队名） |

## ones_app_base_url

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| ones_app_uuid | varchar(64) | null | NO | 每一个私有部署客户对应一个ones_appid(ones_app_uuid) |
| service_name | varchar(128) | null | NO | 服务名称（例子： project-api wiki-api） |
| create_time | int(11) | 0 | NO | 创建时间（秒） |
| update_time | int(11) | 0 | NO | 更新时间（秒） |
| base_url | text | null | NO |  |

## organization

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | -1 | NO |  |
| name | varchar(16) | null | NO |  |
| owner | varchar(8) | null | YES |  |
| logo | varchar(255) | null | NO |  |
| expire_time | bigint(20) | null | NO |  |
| scale | int(6) | null | NO |  |
| csm | varchar(128) |  | YES |  |
| type | tinyint(4) | -1 | NO |  |
| visibility | tinyint(1) | null | NO |  |

## organization_update_stamp - 组织相关数据更新记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO | 组织id |
| type | tinyint(4) | 0 | NO | 更新类型 |
| server_update_stamp | bigint(20) | 0 | NO | 更新时间戳 |

## permission_rule

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | 权限规则id |
| team_uuid | varchar(8) |  | NO | 团队id |
| org_uuid | varchar(8) |  | NO | 组织uuid |
| context_type | int(11) | 0 | NO | 权限上下文类型 |
| context_param_1 | varchar(16) |  | NO | 权限上下文参数1 |
| context_param_2 | varchar(16) |  | NO | 权限上下文参数2 |
| user_domain_type | int(11) | 0 | NO | 用户域类型 |
| user_domain_param | varchar(16) | null | NO | 用户域参数 |
| permission | int(11) | 0 | NO | 权限编号 |
| create_time | bigint(20) | 0 | NO | 创建时间(秒) |
| status | tinyint(4) | 0 | NO | 状态(1:正常 2:删除) |
| read_only | tinyint(4) | 0 | NO | 是否只读 |
| position | int(11) | 0 | NO | 权限显示位置 |

## phone_code - 手机号验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| phone | varchar(32) |  | NO | 手机号 |
| code | varchar(6) |  | NO | 验证码 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| ip | varchar(15) |  | NO | 请求IP地址 |
| status | tinyint(4) | null | NO | 1:正常;2:无效 |

## pipeline - pipeline

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| token | varchar(22) | null | NO | 回调接口使用的token |
| name | varchar(128) |  | NO | 名字 |
| name_pinyin | varchar(1024) |  | NO | 名字拼音 |
| owner | varchar(8) | null | NO | 创建者uuid |
| create_time | bigint(20) | null | NO | 创建时间 |
| status | tinyint(4) | -1 | NO | 1:正常 2:删除 |
| scm_status | tinyint(4) | -1 | NO | scm配置状态(1:未知 2:成功) |
| scm_secret_key | varchar(22) |  | NO | scm github/gitlab 密钥 |
| scm_last_success_time | bigint(20) | null | NO | scm上次成功时间 |
| ci_status | tinyint(4) | -1 | NO | ci配置状态(1:未知 2:成功) |
| ci_last_success_time | bigint(20) | null | NO | ci上次成功时间 |
| lint_status | tinyint(4) | -1 | NO | lint配置状态(1:未知 2:成功) |
| lint_last_success_time | bigint(20) | null | NO | lint上次成功时间 |
| test_status | tinyint(4) | -1 | NO | test配置状态(1:未知 2:成功) |
| test_last_success_time | bigint(20) | null | NO | test上次成功时间 |
| artifact_status | tinyint(4) | -1 | NO | artifact配置状态(1:未知 2:成功) |
| artifact_last_success_time | bigint(20) | null | NO | artifact上次成功时间 |
| sprint_binding_rule | varchar(1024) |  | NO | 迭代绑定脚本 |

## pipeline_run - pipeline 单次执行

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| pipeline_uuid | varchar(8) |  | NO | pipeline uuid |
| number | int(11) | null | NO | 执行编号 |
| repository | varchar(255) |  | NO | 代码仓库名称 |
| branch | varchar(255) |  | NO | 分支名称 |
| start_time | bigint(20) | null | NO | 开始时间 |
| finish_time | bigint(20) | null | NO | 结束时间 |
| build_id | varchar(255) |  | NO | ci构建id |
| build_url | varchar(2048) |  | NO | ci构建url |
| hint | varchar(2048) |  | NO | 构建携带的额外信息 |
| status | int(11) | null | NO | pipeline 执行状态 |
| log | mediumtext | null | NO | 构建log |

## pipeline_run_artifact - pipeline 单次执行绑定的工件

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| run_number | int(11) | null | NO | 流水线执行编号 |
| env | varchar(32) |  | NO | 环境 |
| name | varchar(255) |  | NO | 工件名称 |
| url | varchar(2048) |  | NO | 工件url |
| create_time | bigint(20) | -1 | NO | 创建时间 |

## pipeline_run_commit - pipeline单次执行和commit绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| run_number | int(11) | null | NO | 流水线执行编号 |
| commit_uuid | varchar(8) |  | NO | 代码提交uuid |

## pipeline_run_lint_issue - pipeline单次执行和lint issue绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| run_number | int(11) | null | NO | 流水线执行编号 |
| issue_uuid | varchar(22) | null | NO | 问题uuid |
| status | tinyint(4) | null | NO | 状态(1:开启 2:关闭) |
| create_time | bigint(20) | null | NO | 创建时间（秒） |

## pipeline_run_test_case - pipeline单次执行和自动化测试用例绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| run_number | int(11) | null | NO | 流水线执行编号 |
| case_uuid | varchar(22) |  | NO | 用例uuid |
| result | tinyint(4) | null | NO | 用例执行结果 |
| start_time | bigint(20) | null | NO | 开始时间（毫秒） |
| finish_time | bigint(20) | null | NO | 结束时间（毫秒） |
| message | text | null | NO | 用例执行信息 |

## pipeline_stage - pipeline 单次执行阶段

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| pipeline_uuid | varchar(8) |  | NO | pipeline uuid |
| run_number | int(11) | null | NO | 执行编号 |
| type | int(11) | null | NO | 阶段类型(1:start 2:finish 3:build 4:lint 5:test 6:deploy) |
| name | varchar(255) |  | NO | 阶段名称 |
| start_time | bigint(20) | null | NO | 开始时间 |
| finish_time | bigint(20) | null | NO | 结束时间 |
| status | int(11) | null | NO | 运行状态(1:未知 2:成功 3:失败) |

## pipeline_test_case - pipeline自动化测试用例记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(22) |  | NO | uuid |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| hash | varchar(28) |  | NO | 用于跟踪相同测试用例的哈希 |
| create_time | bigint(20) | null | NO | 首次出现时间（秒） |
| language | varchar(20) | null | NO | 使用的语言 |
| framework | varchar(20) |  | NO | 使用的测试框架 |
| id | varchar(255) |  | NO | 用例 id |
| name | varchar(255) | null | NO | 用例名称 |
| description | varchar(4096) |  | NO | 用例描述 |

## plugin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| name | varchar(64) |  | NO |  |
| dl | varchar(128) |  | NO |  |
| uuid | varchar(8) | null | NO | uuid |
| name | varchar(64) | null | NO |  |
| name_pinyin | varchar(256) | null | NO |  |
| avatar_url | varchar(128) | null | NO |  |
| summary | varchar(128) | null | NO |  |
| status | int(11) | null | NO | 状态 |
| create_time | bigint(20) | null | NO | 创建时间 |

## product_authorization

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| org_uuid | varchar(8) |  | NO |  |
| type | tinyint(4) | null | NO |  |
| status | tinyint(4) | 1 | NO |  |

## product_trial_info - 申请使用产品

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| email | varchar(64) |  | NO | 邮箱 |
| phone | varchar(64) |  | NO | 手机号 |
| nickname | varchar(128) |  | NO | 用户昵称 |
| team_name | varchar(128) |  | NO | 团队名 |
| team_scale | varchar(128) |  | NO | 团队规模 |
| industry | varchar(128) |  | NO | 所处行业 |
| status | int(11) | 0 | NO | 状态 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| updated_time | bigint(20) | 0 | NO | 修改时间 |

## program

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | 甘特图数据 uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| owner | varchar(8) | null | NO | 创建者 uuid |
| create_time | bigint(20) | null | NO | 创建时间 |
| name | varchar(64) | null | NO | 名称 |
| name_pinyin | varchar(512) |  | NO | 名称拼音 |
| plan_start_time | bigint(20) | null | YES | 开始时间 |
| plan_end_time | bigint(20) | null | YES | 结束时间 |
| parent | varchar(8) | null | NO | 父节点 uuid |
| path | varchar(169) | null | NO | 层级关系 |
| position | bigint(20) | 0 | NO | 显示位置 |
| status | tinyint(4) | null | YES | 1:正常 2:删除 |
| assign | varchar(8) | null | YES | 负责人uuid |
| related_type | tinyint(4) | 0 | NO | 关联数据类型 |
| related_project_uuid | varchar(16) | null | NO | 关联的项目uuid |

## program_role - 角色

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | UUID |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| context | varchar(32) |  | NO | 上下文 |
| type | int(11) | 0 | NO | 类型 |
| name | varchar(32) | null | NO | 角色名 |
| name_pinyin | varchar(256) | null | NO | 角色名拼音 |
| create_time | bigint(20) | null | NO | 创建时间（秒） |
| built_in | tinyint(4) | 0 | NO | 是否系统内置角色（1:是 2:否） |
| status | tinyint(4) | null | YES | 状态（1:正常 NULL:删除） |

## program_role_member - 角色成员

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| role_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |

## project

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(16) | null | NO |  |
| name | varchar(128) | null | NO |  |
| name_pinyin | varchar(1024) | null | NO |  |
| owner | varchar(8) | null | YES |  |
| assign | varchar(8) | null | NO | 项目负责人 |
| team_uuid | varchar(8) | null | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| access_mode | tinyint(4) | null | NO |  |
| status_uuid | varchar(16) | null | NO | 状态alias |
| status | tinyint(4) | null | NO |  |
| announcement | text | null | NO |  |
| announcement_text | text | null | NO |  |
| plan_start_time | bigint(20) | null | YES | 计划开始时间 |
| plan_end_time | bigint(20) | null | YES | 计划结束时间 |
| deadline | bigint(20) | 0 | NO |  |
| template_uuid | varchar(8) |  | NO | 模板UUID |
| is_open_email_notify | tinyint(1) | null | NO | 邮件通知开关 |

## project_browse - 用户浏览项目记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| user_uuid | varchar(8) | null | NO | 用户uuid |
| project_uuid | varchar(16) | null | NO | 项目uuid |
| browse_time | bigint(20) | 0 | NO | 浏览时间 |

## project_daily_stats - 项目按天统计记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| project_uuid | varchar(16) | null | NO | 项目uuid |
| type | tinyint(4) | 0 | NO | 统计类型 |
| create_time | bigint(20) | 0 | NO | 统计时间 |
| day | tinyint(4) | 0 | NO | 统计时间天(几号) |
| count | text | null | NO | 统计结果 |

## project_field - 项目属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | UUID |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| context | varchar(32) |  | NO | 上下文 |
| alias | varchar(16) |  | NO | 别名 |
| status | tinyint(4) | null | YES | 状态（1:正常 NULL:删除） |
| type | int(11) | null | NO | 类型 |
| name | varchar(32) | null | NO | 属性名 |
| name_pinyin | varchar(256) | null | NO | 属性名拼音 |
| required | tinyint(4) | null | NO | 是否必填 |
| built_in | tinyint(1) | null | NO | 系统内置属性 |
| create_time | bigint(20) | null | NO | 创建时间（秒） |
| max_length | int(11) | null | YES | 最大长度 |
| options | text | null | YES | 选项配置 |
| statuses | text | null | YES | 属性配置 |

## project_field_value - 项目属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| project_uuid | varchar(16) | null | NO | 项目UUID |
| alias | varchar(32) | null | NO | 别名 |
| type | int(11) | null | NO | 值类型 |
| value | text | null | YES | 属性值 |
| number_value | bigint(20) | null | YES | 数字属性值 |

## project_filter

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | YES |  |
| project_uuid | varchar(16) | null | YES |  |
| owner | varchar(8) | null | YES |  |
| name | varchar(128) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| query | text | null | NO |  |
| sort | text | null | NO |  |
| status | tinyint(4) | 0 | NO |  |
| include_subtasks | tinyint(4) | 0 | NO | 是否包含子任务 |

## project_issue_type

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) | null | NO |  |
| issue_type_uuid | varchar(8) |  | NO |  |
| position | int(11) | 0 | NO | 排序 |

## project_member

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) | null | NO |  |
| user_domain_param | varchar(8) |  | NO | 成员域id |
| user_domain_type | tinyint(4) | 0 | NO | 成员域类型 |

## project_notice_config - project的任务通知配置关联表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) | null | NO | project(项目) uuid |
| issue_type_uuid | varchar(8) | null | NO | 任务类型 uuid |
| notice_config_uuid | varchar(8) | null | NO | 任务通知配置的uuid |
| config_type | enum('create_task','update_task_assign','update_task_status','update_task_priority','update_task_title','update_task_description','set_task_deadline','update_task_sprint','update_task_message','update_task_watcher','update_task_related_task','upload_task_file','update_task_access_manhour','update_task_other_property','delete_task','related_test_case','related_plan_case','task_related_testcase_plan','update_issue_type','update_std_to_sub_issue_type','update_sub_to_sub_issue_type','update_sub_issue_type_parent','update_sub_to_std_issue_type','update_task_remaining_manhour','update_task_record_manhour') | null | NO |  |
| status | tinyint(4) | null | NO | 1:正常 2:无效 |

## project_panel

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| title | varchar(32) | null | NO |  |
| url | varchar(128) | null | NO |  |

## project_pin - 项目Pin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) | null | NO | 所属项目 uuid |
| updated_time | bigint(20) | null | NO | 更新时间 |

## project_pipeline - 项目和pipeline绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) | null | NO |  |
| pipeline_uuid | varchar(8) |  | NO |  |
| create_time | bigint(20) | null | NO | 绑定时间（秒） |

## project_plugin - 项目插件状态表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) | null | NO | 项目id |
| plugin_uuid | varchar(8) | null | NO | 插件id |
| config | text | null | NO |  |
| status | tinyint(1) | 0 | NO | 是否开启工时 1：开启，2：关闭 |
| create_time | bigint(20) | 0 | NO | 创建时间 |

## project_report

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| project_uuid | varchar(16) |  | NO | 项目uuid |
| name | varchar(128) |  | NO | 常用项目报表名称 |
| owner | varchar(8) | null | NO | 创建者uuid |
| config | text | null | NO | 项目报表配置 |
| create_time | bigint(20) | 0 | NO | 创建时间，秒 |
| update_time | bigint(20) | 0 | NO | 更新时间，秒 |
| status | tinyint(4) | 0 | NO | 状态(1:正常, 2:删除) |

## project_sprint_field - 项目下的迭代属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| type | int(11) | 0 | NO | 类型 |
| name | varchar(32) |  | NO | 属性名 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| status | tinyint(4) | 0 | NO | 1:正常;2:删除 |
| position | int(11) | 0 | NO | 排序序号 |

## project_sprint_field_option - 迭代属性选项，用于单选或者多选类型的属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| field_uuid | varchar(8) |  | NO | 属性UUID |
| value | varchar(64) |  | NO | 属性值 |
| default_selected | tinyint(4) | 0 | NO | 是否默认选中 |
| position | int(11) | 0 | NO | 排序序号 |

## project_sprint_status - 项目下的迭代阶段

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| name | varchar(32) |  | NO | 阶段名 |
| category | int(11) | 0 | NO | 状态分类 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| position | int(11) | 0 | NO | 排序序号 |
| built_in | tinyint(4) | 0 | NO | 是否内置状态 |
| status | tinyint(4) | 0 | NO | 1:正常;2:删除 |

## project_task_sort - 项目任务排序表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| attribute_uuid | varchar(16) |  | NO | 任务UUID/分段UUID |
| attribute_type | int(11) | 0 | NO | 类型 |
| position | bigint(20) | 0 | NO | 任务排序权重 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| updated_time | bigint(20) | 0 | NO | 最近修改时间 |

## report_project_attribute - 项目属性分析

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| day_time | int(11) | null | NO | 某天时间戳 |
| project_uuid | varchar(16) | null | NO | 项目uuid |
| attribute_uuid | varchar(8) | null | NO | 属性名 |
| attribute_value | varchar(32) | null | NO | 属性值 |
| count | int(11) | 0 | NO | 个数(累计) |
| updated_time | bigint(20) | 0 | NO | 修改时间 |

## report_project_progress - 项目总体进度表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| day_time | int(11) | null | NO | 某天时间戳 |
| project_uuid | varchar(16) | null | NO | 项目uuid |
| incomplete | int(11) | 0 | NO | 未完成任务数(累计) |
| complete | int(11) | 0 | NO | 已完成任务数(累计) |
| updated_time | bigint(20) | 0 | NO | 修改时间 |

## report_project_target - 项目目标追踪

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| day_time | int(11) | null | NO | 某天时间戳 |
| project_uuid | varchar(16) | null | NO | 项目uuid |
| scores | bigint(20) | 0 | NO | 完成目标值 |
| updated_time | bigint(20) | 0 | NO | 修改时间 |

## request_code

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(64) | null | NO |  |
| email | varchar(64) | null | NO |  |
| phone | varchar(32) | null | NO |  |
| password | varchar(60) | null | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| ip | varchar(15) | null | YES |  |
| status | tinyint(4) | -1 | NO |  |

## reset_code

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(32) | null | NO |  |
| email | varchar(64) | null | NO |  |
| phone | varchar(64) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| ip | varchar(15) | null | YES |  |
| status | tinyint(4) | null | NO |  |

## reset_password_code - 重置团队成员密码邮箱验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| code | varchar(32) | null | NO | 验证码 |
| user_uuid | varchar(8) | null | NO | 用户ID |
| member_uuid | varchar(8) | null | NO | 团队成员的用户ID |
| email | varchar(255) | null | NO | 邮箱地址 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| ip | varchar(15) | null | NO | 请求IP地址 |
| status | tinyint(4) | null | NO | 1:正常,2:无效 |

## resource

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| reference_type | tinyint(4) | null | NO |  |
| reference_id | varchar(16) | null | YES |  |
| team_uuid | varchar(8) | null | YES |  |
| project_uuid | varchar(16) | null | NO |  |
| owner_uuid | varchar(8) | null | YES |  |
| type | tinyint(4) | -1 | NO |  |
| source | tinyint(4) | 0 | NO |  |
| ext_id | varchar(28) | null | YES |  |
| name | varchar(255) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| modify_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | -1 | NO |  |
| description | varchar(64) | null | NO |  |

## role

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| name | varchar(32) |  | NO | 角色名称 |
| name_pinyin | varchar(256) |  | NO | 角色名称拼音 |
| built_in | tinyint(4) | 0 | NO | 是否系统内置角色 |
| is_project_member | tinyint(4) | 0 | NO | 是否项目成员角色 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| status | tinyint(4) | null | YES | 1:正常;null:删除 |

## role_config - 角色配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) | null | NO |  |
| role_uuid | varchar(8) |  | NO |  |
| create_time | bigint(20) | 0 | NO |  |

## role_member - 角色成员

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO |  |
| role_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) |  | NO |  |

## section - 分段表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | 项目UUID |
| source | varchar(8) |  | NO | 来源 |
| source_uuid | varchar(16) |  | NO | 来源对应UUID |
| mark | varchar(32) |  | NO | 标记 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| updated_time | bigint(20) | 0 | NO | 最近修改时间 |

## session_token

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| user_id | varchar(8) | null | YES |  |
| session_token | varchar(64) | null | YES |  |
| session_token_status | tinyint(4) | -1 | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| ip | varchar(15) | null | YES |  |
| type | int(11) | 0 | NO |  |

## setting

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| name | varchar(255) | null | NO |  |
| value | varchar(255) | null | YES |  |

## sprint

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) | null | YES |  |
| title | varchar(32) | null | YES |  |
| description | text | null | YES |  |
| start_time | bigint(20) | 0 | YES |  |
| end_time | bigint(20) | 0 | YES |  |
| status | tinyint(4) | -1 | NO |  |
| status_uuid | varchar(8) | null | NO | 状态UUID |
| create_time | bigint(20) | 0 | YES |  |
| title_pinyin | varchar(256) | null | NO |  |
| assign | varchar(8) |  | NO | 迭代负责人 uuid |
| goal | text | null | NO | 迭代目标 |

## sprint_field_value - 迭代属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| sprint_uuid | varchar(8) |  | NO | 迭代UUID |
| field_uuid | varchar(8) | null | NO | 属性UUID |
| type | int(11) | 0 | NO | 值类型 |
| value | text | null | YES | 属性值 |
| number_value | bigint(20) | null | YES | 数字属性值 |

## sprint_pin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| sprint_uuid | varchar(8) | null | NO |  |
| updated_time | bigint(20) | null | NO | 更新时间 |

## sprint_pipeline - 迭代和pipeline分支绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| sprint_uuid | varchar(8) |  | NO | 迭代uuid |
| pipeline_uuid | varchar(8) |  | NO | 流水线uuid |
| repository | varchar(255) |  | NO | 代码仓库名称 |
| branch | varchar(255) |  | NO | 分支名称 |
| create_time | bigint(20) | null | NO | 绑定时间（秒） |

## sprint_status_value - 迭代的阶段信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| sprint_uuid | varchar(8) |  | NO | 迭代UUID |
| status_uuid | varchar(8) |  | NO | 阶段UUID |
| plan_start_time | bigint(20) | null | YES | 计划开始时间 |
| plan_end_time | bigint(20) | null | YES | 计划结束时间 |
| actual_start_time | bigint(20) | null | YES | 实际开始时间 |
| actual_end_time | bigint(20) | null | YES | 实际结束时间 |
| desc_rich | text | null | NO | 阶段描述(富文本) |

## sprint_timeline

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| sprint_uuid | varchar(8) | null | NO |  |
| type | varchar(16) | null | NO |  |
| time | bigint(20) | null | YES |  |

## sso_verify_info - sso验证

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| auth_code | varchar(512) |  | NO |  |
| sync_id | varchar(255) |  | NO | 同步账号的id |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| create_time | bigint(20) | 0 | NO |  |

## support_request - 申请支持（升级到专业版）请求

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| email | varchar(255) |  | NO | 邮箱 |
| phone | varchar(32) |  | NO | 手机号 |
| source | int(11) | 0 | NO | 来源(1:创建团队 2:升级到专业版) |
| urgent | tinyint(4) | 0 | NO | 是否紧急 |
| team_uuid | varchar(8) | null | NO | 请求者团队uuid |
| user_uuid | varchar(8) | null | NO | 请求者uuid |
| name | varchar(128) |  | NO | 用户名称 |
| team_name | varchar(128) |  | NO | 团队名 |
| team_scale | varchar(128) |  | NO | 团队规模 |
| industry | varchar(128) |  | NO | 所处行业 |
| create_time | datetime | CURRENT_TIMESTAMP | NO | 创建时间 |
| status | int(11) | 0 | NO | 1:正常 2:删除 |
| referrer | text | null | NO | 渠道信息 |
| keyword | text | null | NO | 搜索关键词 |
| ip_address | varchar(15) |  | NO | 用户IP |
| ip_location | varchar(255) |  | NO | IP地理位置 |
| channel | varchar(255) |  | NO | 渠道来源 |
| issue | varchar(48) |  | NO | 希望解决的问题 |
| team_role | varchar(24) |  | NO | 团队角色 |
| demo_name | varchar(64) |  | NO | demo?? |
| email | varchar(255) |  | NO | 邮箱 |
| phone | varchar(32) |  | NO | 手机号 |
| source | int(11) | 0 | NO | 来源(1:创建团队 2:升级到专业版) |
| urgent | tinyint(4) | 0 | NO | 是否紧急 |
| team_uuid | varchar(8) | null | NO | 请求者团队uuid |
| user_uuid | varchar(8) | null | NO | 请求者uuid |
| name | varchar(128) |  | NO | 用户名称 |
| team_name | varchar(128) |  | NO | 团队名 |
| team_scale | varchar(128) |  | NO | 团队规模 |
| industry | varchar(128) |  | NO | 所处行业 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| status | int(11) | 0 | NO | 1:正常 2:删除 |
| referrer | text | null | NO | 渠道信息 |
| keyword | text | null | NO | 搜索关键词 |
| ip_address | varchar(15) |  | NO | 用户IP |
| ip_location | varchar(255) |  | NO | IP地理位置 |
| channel | varchar(255) |  | NO | 渠道来源 |
| demo_name | varchar(64) |  | NO | demo名称 |
| issue | varchar(48) |  | NO | 希望解决的问题 |
| team_role | varchar(24) |  | NO | 团队角色 |

## sync_department_config - 选中第三方部门配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| org_uuid | varchar(8) | null | NO | 组织uuid |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| sync_id | bigint(20) | null | YES | 同步部门的id |

## sync_user - 同步用户表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| user_uuid | varchar(8) |  | NO | user id |
| sync_id | varchar(255) |  | NO | 同步账号的id |
| sync_type | tinyint(4) | 0 | NO | 同步账号类型 |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| status | tinyint(4) | -1 | NO | 状态:1:正常;2:删除 |

## sync_user_config - 选中第三方用户配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| org_uuid | varchar(8) | null | NO | 组织uuid |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| sync_id | varchar(255) |  | NO | 同步账号的id |

## tag - 标签库

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| name | varchar(32) |  | NO | 标签名 |

## task

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(16) | null | NO |  |
| team_uuid | varchar(8) | null | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| priority | varchar(8) | null | YES |  |
| owner | varchar(8) | null | YES |  |
| assign | varchar(8) | null | YES |  |
| tags | text | null | YES |  |
| sprint_uuid | varchar(8) | null | YES |  |
| project_uuid | varchar(16) | null | YES |  |
| issue_type_uuid | varchar(8) |  | NO | 任务类型UUID |
| status_uuid | varchar(8) | null | NO | 任务状态UUID |
| deadline | bigint(20) | null | YES |  |
| status | tinyint(4) | 0 | NO |  |
| summary | text | null | NO |  |
| desc | longtext | null | YES | 任务描述 |
| desc_rich | longtext | null | YES | 任务描述（富文本） |
| server_update_stamp | bigint(20) | 0 | NO |  |
| complete_time | bigint(20) | 0 | NO |  |
| open_time | bigint(20) | 0 | NO |  |
| score | bigint(20) | null | YES |  |
| parent_uuid | varchar(16) | null | NO |  |
| position | bigint(20) | 0 | NO | 子任务排序 |
| number | int(11) | null | NO | 编号 |
| assess_manhour | bigint(20) | null | YES | 预估工时 |
| path | varchar(169) | null | NO | 支持10层子任务（n*17-1） |
| sub_issue_type_uuid | varchar(8) |  | NO | sub_issue_type_uuid为空则不是子任务； 不为空是子任务，此时issue_type_uuid是父任务的标准任务类型 |
| related_count | int(11) | 0 | NO | 关联任务的数量 |
| remaining_manhour | bigint(20) | null | YES | 剩余工时 |

## task_attendee

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) | null | YES |  |
| user_uuid | varchar(8) | null | YES |  |

## task_attribute - 任务属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) |  | NO | 任务UUID |
| template_uuid | varchar(8) |  | NO | 模板UUID |
| template_attribute_uuid | varchar(8) |  | NO | 属性uuid |
| value | text | null | YES | 属性值 |
| type | int(11) | 0 | NO | 值类型 |
| status | int(11) | 0 | NO | 状态(1:正常,2:删除) |

## task_case

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) | null | NO |  |
| case_uuid | varchar(8) | null | NO |  |

## task_commit

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) | null | NO |  |
| commit_uuid | varchar(8) | null | NO |  |

## task_gantt_chart

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | gantt chart uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| component_uuid | varchar(8) | null | NO | 组件 uuid |
| name | varchar(32) | null | NO | gantt chart名称 |
| name_pinyin | varchar(256) |  | NO | gantt chart名称拼音 |
| owner | varchar(8) | null | NO | 创建者 uuid |
| create_time | bigint(20) | null | NO | 创建时间 |
| status | tinyint(4) | null | YES | 1:正常 2:删除 |

## task_gantt_chart_personal_config

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| gantt_chart_uuid | varchar(8) | null | NO | gantt chart uuid |
| user_uuid | varchar(8) | null | NO | 用户 uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| expand | tinyint(4) | 1 | NO | 0:收起  1: 展开 |
| zooming | tinyint(4) | 0 | NO | 缩放 |
| create_time | bigint(20) | null | NO | 创建时间 |
| status | tinyint(4) | 1 | NO | 1: 正常 2: 删除 |

## task_gantt_data

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | 甘特图数据 uuid |
| team_uuid | varchar(8) | null | NO | 团队 uuid |
| gantt_chart_uuid | varchar(8) | null | NO | gantt chart uuid |
| type | tinyint(4) | 0 | YES | 1:甘特图数据类型, 1:任务 2:分组 3:里程碑 |
| name | varchar(128) | null | NO | 名称 |
| name_pinyin | varchar(512) |  | NO | 名称拼音 |
| owner | varchar(8) | null | NO | 创建者 uuid |
| start_time | bigint(20) | null | NO | 开始时间 |
| end_time | bigint(20) | null | NO | 结束时间 |
| progress | int(11) | 0 | NO | 进度 |
| parent | varchar(8) | null | NO | 父节点 uuid |
| path | varchar(169) | null | NO | 层级关系 |
| position | bigint(20) | 0 | NO | 显示位置 |
| create_time | bigint(20) | null | NO | 创建时间 |
| status | tinyint(4) | null | YES | 1:正常 2:删除 |
| assign | varchar(8) | null | YES | 负责人uuid |
| related_type | tinyint(4) | 0 | NO | 关联数据类型 |
| related_project_uuid | varchar(16) | null | NO | 关联的项目uuid |
| related_task_uuid | varchar(16) | null | NO | 关联的任务uuid |

## task_link_type

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| name | varchar(32) |  | NO | 任务关联类型名称 |
| name_pinyin | varchar(256) | null | NO |  |
| link_out_desc | varchar(32) |  | NO | 链出描述 |
| link_out_desc_pinyin | varchar(256) |  | NO | 链出描述拼音 |
| link_in_desc | varchar(32) |  | NO | 链入描述 |
| link_in_desc_pinyin | varchar(256) |  | NO | 链入描述拼音 |
| built_in | tinyint(1) | 0 | NO | 系统内置默认任务关联类型,1:是,0否 |
| create_time | bigint(20) | 0 | YES |  |
| status | tinyint(4) | 1 | NO | 1:正常 2:删除 |

## task_plan_case

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) | null | NO |  |
| plan_uuid | varchar(8) | null | NO |  |
| case_uuid | varchar(8) | null | NO |  |

## task_preference

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| user_uuid | varchar(8) | null | NO |  |
| task_uuid | varchar(16) | null | NO |  |
| order | int(11) | 0 | NO |  |
| alarm_time | bigint(20) | 0 | NO |  |
| server_update_stamp | bigint(20) | 0 | NO |  |

## task_related

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO |  |
| source_task_uuid | varchar(16) | null | NO |  |
| target_task_uuid | varchar(16) | null | NO |  |
| task_link_type_uuid | varchar(8) |  | NO | 任务关联类型UUID |
| create_time | bigint(20) | 0 | YES |  |

## task_richtext_message - 任务富文本属性动态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| message_uuid | varchar(8) | null | NO | message_uuid |
| task_uuid | varchar(16) | null | NO | 任务UUID |
| field_uuid | varchar(8) | null | NO | field_uuid |
| old_value_time | bigint(20) | null | NO | 产生上一版本任务富文本属性的时间 |
| new_value_time | bigint(20) | null | NO | 产生当前版本任务富文本属性的时间 |
| old_value | longtext | null | YES | 更改前的任务富文本属性值 |
| new_value | longtext | null | YES | 更改后的任务富文本属性值 |

## task_sort - task_sort

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) | null | NO | uuid |
| project_uuid | varchar(16) | null | NO | 所属项目 uuid |
| board_uuid | varchar(8) | null | NO | 所属board uuid |
| position | bigint(20) | null | NO | 顺序(从小到大) |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 更新时间 |

## task_status - 任务状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| name | varchar(32) |  | NO | 任务状态名 |
| name_pinyin | varchar(256) | null | NO |  |
| category | int(11) | 0 | NO | 任务状态分类 |
| built_in | tinyint(4) | 0 | NO | 是否系统内置状态 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| status | tinyint(4) | null | YES | 1:正常;null:删除 |

## task_status_config - 任务状态在某项目的某任务类型中的配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| project_uuid | varchar(16) |  | NO | 项目uuid |
| issue_type_uuid | varchar(8) |  | NO | 任务类型uuid |
| status_uuid | varchar(8) |  | NO | 任务状态uuid |
| is_default | tinyint(4) | 0 | NO | 是否默认状态 |
| position | int(11) | 0 | NO | 状态位置 |

## task_testcase_plan

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| task_uuid | varchar(16) | null | NO |  |
| plan_uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |

## task_watcher

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| task_uuid | varchar(16) |  | NO |  |
| user_uuid | varchar(8) |  | NO |  |
| create_time | bigint(20) | 0 | NO |  |

## team

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | -1 | NO |  |
| name | varchar(16) | null | NO |  |
| owner | varchar(8) | null | YES |  |
| logo | varchar(255) | null | NO |  |
| domain | varchar(255) | null | YES |  |
| cover_url | varchar(128) |  | NO |  |
| expire_time | bigint(20) | null | NO |  |
| scale | int(6) | null | NO |  |
| csm | varchar(128) |  | YES |  |
| type | tinyint(4) | -1 | NO |  |
| sprint_alias | varchar(32) |  | NO |  |
| org_uuid | varchar(8) | null | NO |  |
| workdays | varchar(8) | null | NO |  |
| workhours | bigint(20) | 0 | YES |  |
| uuid | varchar(8) | null | NO |  |
| is_create_default_space | tinyint(4) | null | NO |  |

## team_member

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| modify_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | -1 | NO |  |
| user_avatar | varchar(255) |  | NO | 头像 |
| user_title | varchar(16) |  | NO | 职位 |
| user_name | varchar(16) |  | NO | 姓名 |
| user_name_pinyin | varchar(128) | null | NO |  |
| inviter_uuid | varchar(8) | null | NO | 邀请人uuid |
| user_team_constraint | varchar(32) | null | YES | 用户只有一个有效团队 |
| org_uuid | varchar(8) | null | NO |  |

## team_notice_config - team的全局任务通知配置关联表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | team（团队） uuid |
| issue_type_uuid | varchar(8) | null | NO | 任务类型 uuid |
| notice_config_uuid | varchar(8) | null | NO | 任务通知配置的uuid |
| config_type | enum('create_task','update_task_assign','update_task_status','update_task_priority','update_task_title','update_task_description','set_task_deadline','update_task_sprint','update_task_message','update_task_watcher','update_task_related_task','upload_task_file','update_task_access_manhour','update_task_other_property','delete_task','related_test_case','related_plan_case','task_related_testcase_plan','update_issue_type','update_std_to_sub_issue_type','update_sub_to_sub_issue_type','update_sub_issue_type_parent','update_sub_to_std_issue_type','update_task_remaining_manhour','update_task_record_manhour') | null | NO |  |
| status | tinyint(4) | null | NO | 1:正常 2:无效 |

## team_plugin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | uuid |
| plugin_uuid | varchar(8) | null | NO | uuid |
| is_open | tinyint(1) | null | NO |  |
| create_time | bigint(20) | null | NO | 创建时间 |
| config | text | null | NO |  |

## team_report

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| name | varchar(128) | null | NO | 常用团队报表名称 |
| owner | varchar(8) | null | NO | 创建者uuid |
| config | text | null | NO | 项目报表配置 |
| create_time | bigint(20) | 0 | NO | 创建时间，秒 |
| update_time | bigint(20) | 0 | NO | 更新时间，秒 |
| status | tinyint(4) | 0 | NO | 状态(1:正常, 2:删除) |

## team_scm

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队uuid |
| status | tinyint(4) | 1 | NO | 1：正常 2：删除 |
| create_time | bigint(20) | null | NO | 创建时间 |
| scm_secret_key | varchar(22) | null | NO | scm github/gitlab 密钥 |
| scm_last_success_time | bigint(20) | null | NO | 最后的更新时间 |

## template - 模板基础表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| owner | varchar(16) |  | NO | 拥有者 |
| category_uuid | varchar(8) |  | NO | 分类uuid |
| name | varchar(32) |  | NO | 模板名 |
| type | tinyint(4) | 0 | NO | 类型 |
| descs | text | null | YES | 描述 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| updated_time | bigint(20) | 0 | NO | 修改时间 |
| is_open_manhour | tinyint(1) | 0 | YES | 工时开关 |
| uuid | varchar(8) | null | NO | uuid |
| team_uuid | varchar(8) | null | NO | 所属团队 |
| space_uuid | varchar(8) | null | NO | 所属 space |
| owner_uuid | varchar(8) | null | NO | 拥有者 |
| title | varchar(64) | null | NO | 模版标题 |
| title_pinyin | varchar(256) |  | NO | 模版标题-拼音 |
| status | tinyint(4) | null | NO | 状态 (1: 正常, 2: 删除) |
| create_time | bigint(20) | null | NO | 创建时间 |
| update_time | bigint(20) | null | NO | 最近更新时间 |

## template_attribute - 模板任务属性表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| template_uuid | varchar(8) |  | NO | 模板UUID |
| name | varchar(32) |  | NO | 属性名 |
| type | int(11) | 0 | NO | 类型 |
| text | text | null | NO | 文本属性值 |
| number | bigint(20) | null | YES | 数字属性值 |
| is_default_show | tinyint(1) | 1 | NO | 任务列表是否显示 |
| position | int(11) | 0 | NO | 排序序号 |

## template_attribute_value - 属性值表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| attribute_uuid | varchar(8) |  | NO | 属性UUID |
| value | varchar(64) |  | NO | 属性值 |
| font_color | varchar(8) |  | NO | 字体颜色 |
| bg_color | varchar(8) |  | NO | 字体颜色 |
| is_selected | tinyint(1) | 0 | NO | 单选框中值是否被选择 |
| position | int(11) | 0 | NO | 排序序号 |

## template_category - 模板分类信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| name | varchar(128) |  | NO | 分类名 |
| pic_url | varchar(1024) |  | NO | 分类图标地址 |

## template_tag_relational - 模板标签映射

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| template_uuid | varchar(8) |  | NO | UUID |
| tag_uuid | varchar(16) |  | NO | 拥有者 |

## template_target - 模板任务目标表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| template_uuid | varchar(8) |  | NO | 模板UUID |
| name | varchar(32) |  | NO | 量化指标 |
| score | bigint(20) | null | YES | 目标值 |
| unit | varchar(32) |  | NO | 单位 |

## testcase_case

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| library_uuid | varchar(8) | null | NO |  |
| path | varchar(255) | null | NO |  |
| module_uuid | varchar(8) | null | NO |  |
| name | varchar(1024) | null | NO |  |
| priority | varchar(8) | null | NO |  |
| type | varchar(16) | null | NO |  |
| assign | varchar(8) | null | NO |  |
| desc | text | null | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| number | int(11) | null | NO | 编号 |
| status | tinyint(4) | 0 | NO |  |
| condition | text | null | YES |  |
| name_pinyin | varchar(10000) | null | NO | 用例名字拼音 |

## testcase_case_step

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| library_uuid | varchar(8) | null | NO |  |
| case_uuid | varchar(8) | null | NO |  |
| desc | text | null | YES |  |
| result | text | null | YES |  |
| index | int(5) | null | NO |  |

## testcase_default_member

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) | null | NO |  |
| user_domain_type | tinyint(4) | 0 | NO | 成员域类型 |
| user_domain_param | varchar(8) |  | NO | 成员域参数 |

## testcase_field - TestCase属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | UUID |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| context | varchar(32) |  | NO | 上下文 |
| alias | varchar(16) |  | NO | 别名 |
| status | tinyint(4) | null | YES | 状态（1:正常 NULL:删除） |
| type | int(11) | null | NO | 类型 |
| name | varchar(32) | null | NO | 属性名 |
| name_pinyin | varchar(256) | null | NO | 属性名拼音 |
| required | tinyint(4) | null | NO | 是否必填 |
| built_in | tinyint(1) | null | NO | 系统内置属性 |
| create_time | bigint(20) | null | NO | 创建时间（秒） |
| max_length | int(11) | null | YES | 最大长度 |
| options | text | null | YES | 选项配置 |
| statuses | text | null | YES | 属性配置 |
| default_value | text | null | YES | 属性默认值 |

## testcase_field_config - TestCase属性配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| name | varchar(32) | null | NO | 属性配置模版 |
| name_pinyin | varchar(256) | null | NO |  |
| status | tinyint(4) | 0 | NO |  |
| create_time | bigint(20) | null | NO | 创建时间（秒） |
| is_default | tinyint(1) | null | NO | 是否为默认属性配置 |

## testcase_field_value - TestCase属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队UUID |
| object_uuid | varchar(8) | null | NO | testcase的所有实体，包括用例，用例库，测试计划等等 |
| item_type | varchar(32) | null | NO | itemType |
| alias | varchar(32) | null | NO | 别名 |
| type | int(11) | null | NO | 值类型 |
| value | text | null | YES | 属性值 |
| number_value | bigint(20) | null | YES | 数字属性值 |

## testcase_library

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| name | varchar(32) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | 0 | NO |  |
| name_pinyin | varchar(256) | null | NO | 用例库名字拼音 |
| field_config_uuid | varchar(8) | null | NO | 属性配置 |

## testcase_library_pin - 用例库Pin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| library_uuid | varchar(16) | null | NO | 所属library uuid |
| updated_time | bigint(20) | null | NO | 更新时间 |

## testcase_module

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| parent_uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| library_uuid | varchar(8) | null | NO |  |
| path | varchar(255) | null | NO |  |
| name | varchar(32) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | 0 | NO |  |
| is_default | tinyint(1) | null | NO | 是否为无所属模块 |
| position | bigint(20) | null | NO |  |
| name_pinyin | varchar(256) | null | NO | 模块名字拼音 |

## testcase_plan

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| owner | varchar(8) | null | NO |  |
| name | varchar(32) | null | NO |  |
| stage | varchar(16) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | 0 | NO |  |
| plan_status | tinyint(4) | 0 | NO |  |
| related_project | varchar(16) | null | YES |  |
| related_sprint | varchar(8) | null | YES |  |
| related_issue_type | varchar(8) | null | YES |  |
| name_pinyin | varchar(256) | null | NO | 测试计划名字拼音 |

## testcase_plan_assigns

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| plan_uuid | varchar(8) |  | NO |  |
| user_uuid | varchar(8) |  | NO |  |
| create_time | bigint(20) | null | NO | 创建时间 |

## testcase_plan_case

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| plan_uuid | varchar(8) | null | NO |  |
| case_uuid | varchar(8) | null | NO |  |
| executor | varchar(8) | null | NO |  |
| result | varchar(8) | null | NO |  |
| note | text | null | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| warn_step | varchar(8) | null | YES |  |
| team_uuid | varchar(8) | null | NO |  |

## testcase_plan_case_step

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| plan_uuid | varchar(8) | null | NO |  |
| case_uuid | varchar(8) | null | NO |  |
| step_uuid | varchar(8) | null | NO |  |
| step_result | varchar(8) | null | NO |  |
| actual_result | text | null | YES |  |
| team_uuid | varchar(8) | null | NO |  |

## testcase_project_config

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| project_uuid | varchar(16) | null | NO |  |
| issue_type_uuid | varchar(8) | null | NO |  |

## testcase_report

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| plan_uuid | varchar(8) | null | NO |  |
| start_time | bigint(20) | 0 | NO |  |
| end_time | bigint(20) | 0 | NO |  |
| summary | text | null | YES |  |
| title | varchar(32) |  | YES |  |
| team_uuid | varchar(8) | null | NO |  |

## testcase_report_template

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| content | text | null | YES |  |
| name | varchar(32) | null | NO |  |
| status | tinyint(4) | 0 | NO |  |

## third_party_config

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| wechat_suite_ticket | varchar(512) |  | YES | 企业微信suite_ticket |
| dingding_suite_ticket | varchar(512) |  | YES | 钉钉suite_ticket |

## third_party_organization

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| corp_uuid | varchar(64) | null | NO | 第三方登录企业id |
| org_uuid | varchar(8) | null | NO | 对应组织uuid |
| corp_name | varchar(255) | null | NO | 第三方企业名称 |
| org_type | tinyint(4) | 1 | NO | 来源平台 |
| auth_status | tinyint(4) | 1 | NO | 授权状态 |
| permanent_code | varchar(512) | null | NO | 第三方企业微信永久授权码 |
| visible_scale | int(6) | null | NO | 可见范围内人数 |
| agent_id | int(20) | null | NO | 授权应用agentid |

## transition - 工作流步骤

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO | UUID |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| project_uuid | varchar(16) |  | NO | 项目UUID |
| issue_type_uuid | varchar(8) |  | NO | 任务类型UUID |
| start_status | varchar(8) | null | NO | 起始状态UUID |
| end_status | varchar(8) | null | NO | 结束状态UUID |
| name | varchar(64) |  | NO | 名称 |
| trigger | text | null | NO | 触发器 |
| condition | text | null | NO | 验证条件 |
| post_function | text | null | NO | 后置动作 |
| fields | text | null | NO | 步骤属性 |
| position | int(11) | 0 | NO | 步骤位置 |

## update_stamp - 获取更新状态记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队id |
| type | tinyint(4) | 0 | NO | 更新类型 |
| user_uuid | varchar(8) |  | NO | 用户id |
| server_update_stamp | bigint(20) | 0 | NO | 更新标识 |

## user

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| Host | char(60) |  | NO |  |
| User | char(32) |  | NO |  |
| Select_priv | enum('N','Y') | N | NO |  |
| Insert_priv | enum('N','Y') | N | NO |  |
| Update_priv | enum('N','Y') | N | NO |  |
| Delete_priv | enum('N','Y') | N | NO |  |
| Create_priv | enum('N','Y') | N | NO |  |
| Drop_priv | enum('N','Y') | N | NO |  |
| Reload_priv | enum('N','Y') | N | NO |  |
| Shutdown_priv | enum('N','Y') | N | NO |  |
| Process_priv | enum('N','Y') | N | NO |  |
| File_priv | enum('N','Y') | N | NO |  |
| Grant_priv | enum('N','Y') | N | NO |  |
| References_priv | enum('N','Y') | N | NO |  |
| Index_priv | enum('N','Y') | N | NO |  |
| Alter_priv | enum('N','Y') | N | NO |  |
| Show_db_priv | enum('N','Y') | N | NO |  |
| Super_priv | enum('N','Y') | N | NO |  |
| Create_tmp_table_priv | enum('N','Y') | N | NO |  |
| Lock_tables_priv | enum('N','Y') | N | NO |  |
| Execute_priv | enum('N','Y') | N | NO |  |
| Repl_slave_priv | enum('N','Y') | N | NO |  |
| Repl_client_priv | enum('N','Y') | N | NO |  |
| Create_view_priv | enum('N','Y') | N | NO |  |
| Show_view_priv | enum('N','Y') | N | NO |  |
| Create_routine_priv | enum('N','Y') | N | NO |  |
| Alter_routine_priv | enum('N','Y') | N | NO |  |
| Create_user_priv | enum('N','Y') | N | NO |  |
| Event_priv | enum('N','Y') | N | NO |  |
| Trigger_priv | enum('N','Y') | N | NO |  |
| Create_tablespace_priv | enum('N','Y') | N | NO |  |
| ssl_type | enum('','ANY','X509','SPECIFIED') |  | NO |  |
| ssl_cipher | blob | null | NO |  |
| x509_issuer | blob | null | NO |  |
| x509_subject | blob | null | NO |  |
| max_questions | int(11) unsigned | 0 | NO |  |
| max_updates | int(11) unsigned | 0 | NO |  |
| max_connections | int(11) unsigned | 0 | NO |  |
| max_user_connections | int(11) unsigned | 0 | NO |  |
| plugin | char(64) | mysql_native_password | NO |  |
| authentication_string | text | null | YES |  |
| password_expired | enum('N','Y') | N | NO |  |
| password_last_changed | timestamp | null | YES |  |
| password_lifetime | smallint(5) unsigned | null | YES |  |
| account_locked | enum('N','Y') | N | NO |  |
| uuid | varchar(8) | null | NO |  |
| email | varchar(128) | null | YES |  |
| create_time | bigint(20) | 0 | NO |  |
| status | tinyint(4) | 0 | NO |  |
| verify_status | tinyint(4) | -1 | NO | 邮箱验证状态 |
| name | varchar(16) | null | NO |  |
| avatar | varchar(255) | null | YES |  |
| phone | varchar(32) | null | YES | 手机号 |
| channel | varchar(32) | null | YES |  |
| password | varchar(60) | null | YES |  |
| access_time | bigint(20) | 0 | YES | 最近访问时间 |
| sync_id | varchar(255) | null | YES | 同步账号的id |

## user_access - 用户访问记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队id |
| user_uuid | varchar(8) | null | NO | 用户id |
| type | tinyint(4) | null | NO | 产品类型 |
| access_time | bigint(20) | 0 | NO | 登录时间 |
| weekday | tinyint(4) | null | NO | 星期几 |

## user_filter_view - filter表迁移到本表,filter表将被弃用，后续filter表仅用于兼容旧客户端

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| filter_uuid | varchar(8) | null | NO | filter_uuid,是系统内置视图(比如ft-t-001)时与uuid不相等，其他情况与uuid相等 |
| owner | varchar(8) |  | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| name | varchar(64) |  | NO | 视图名称 |
| layout | tinyint(1) | 3 | NO | 布局方式(1.表格 2.宽详情 3.窄详情) |
| query | text | null | YES | 筛选方式 |
| condition | text | null | NO | 筛选条件 |
| sort | varchar(256) |  | YES | 排序方式 |
| group_by | varchar(256) |  | YES | 分组方式 |
| create_time | bigint(20) | null | NO |  |
| built_in | tinyint(1) | 0 | NO | 是否默认视图 |
| include_subtasks | tinyint(1) | 0 | NO | 是否包含子任务 |
| table_field_settings | text | null | YES | 表格模式的表头 |
| board_settings | text | null | YES | 看板视图设置 |
| is_fold_all_groups | tinyint(1) | 0 | NO | 是否折叠分组 |
| status | tinyint(1) | 1 | YES | 1正常 2已删除 |
| display_type | tinyint(4) | null | YES | 1:平铺展示, 2:子树展示，3:折叠展示 |
| is_show_derive | tinyint(1) | 0 | NO | 是否显示派生 |
| shared | tinyint(4) | 0 | NO | 是否共享,0: 不共享，1: 共享 |

## user_filter_view_config

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| filter_uuid | varchar(8) | null | NO | 对应user_filter_view的filter_uuid |
| owner | varchar(8) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| position | int(11) | null | NO | 视图位置 |
| is_show | tinyint(1) | 1 | NO | 1.显示 0.隐藏 |

## user_filter_view_member - 筛选器成员（被分享者）

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) |  | NO | 团队UUID |
| filter_uuid | varchar(8) |  | NO | 筛选器UUID |
| user_domain_param | varchar(8) |  | NO | 成员域参数 |
| user_domain_type | tinyint(4) | 0 | NO | 成员域类型 |
| position | int(11) | 0 | NO | 位置 |

## user_group

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| team_uuid | varchar(8) | null | NO |  |
| create_time | bigint(20) | 0 | NO |  |
| name | varchar(32) |  | NO |  |
| name_pinyin | varchar(256) | null | NO |  |
| desc | text | null | NO |  |
| status | tinyint(4) | 0 | NO |  |

## user_group_member

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| group_uuid | varchar(8) |  | NO |  |
| user_uuid | varchar(8) |  | NO |  |

## user_role - 用户角色映射更新表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| user_uuid | varchar(8) | null | NO | 用户uuid |
| role_uuid | varchar(8) | null | NO | 角色uuid |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO |  |
| team_uuid | varchar(8) | null | NO |  |

## user_task_sort - 我的任务排序表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| user_uuid | varchar(8) |  | NO | 用户UUID |
| attribute_uuid | varchar(16) |  | NO | 任务UUID/分段UUID |
| attribute_type | int(11) | 0 | NO | 类型 |
| position | bigint(20) | 0 | NO | 任务排序权重 |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| updated_time | bigint(20) | 0 | NO | 最近修改时间 |

## user_template_relational - 用户模板映射

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| user_uuid | varchar(8) |  | NO | 用户UUID |
| template_uuid | varchar(16) |  | NO | 模板UUID |
| create_time | bigint(20) | 0 | NO | 创建时间 |
| updated_time | bigint(20) | 0 | NO | 修改时间 |

## version - 版本管理

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) |  | NO |  |
| owner | varchar(8) |  | YES | 创建者 |
| team_uuid | varchar(8) |  | YES | 所属团队 |
| title | varchar(64) |  | NO | 版本号 |
| desc | text | null | YES | 版本描述（富文本） |
| assign | varchar(8) | null | YES | 负责人 |
| release_time | bigint(20) | null | YES | 版本发布时间 |
| create_time | bigint(20) | 0 | NO |  |
| category | tinyint(1) | 1 | NO | 分类 未开始：1，进行中：2，已完成：3 |
| status | tinyint(1) | 1 | NO | 状态 正常：1，删除：2 |
| sys_version | varchar(5) |  | NO |  |
| mysql_version | varchar(6) |  | NO |  |

## version_sprint

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| version_uuid | varchar(8) |  | NO |  |
| sprint_uuid | varchar(8) |  | NO |  |
| project_uuid | varchar(16) |  | NO |  |
| create_time | bigint(20) | 0 | NO |  |

