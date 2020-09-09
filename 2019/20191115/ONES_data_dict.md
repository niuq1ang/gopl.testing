# ONES 数据字典

## `batch_task` - 批量任务

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `owner` | varchar(8) | 空串 | no |  |
| `job_id` | varchar(16) | 空串 | no | 任务id |
| `job_type` | varchar(128) | 0 | no | 批量任务类型(1.批量删除 2. 批量移动) |
| `job_status` | tinyint(4) | 1 | no | 批量任务状态(1.等待中 2.进行中 3.中止 4.<br>完成) |
| `batch_type` | varchar(64) | batch | no | 批量任务的默认类型（批量，单个操作） |
| `status` | tinyint(4) | 1 | no | 任务状态(1.显示 2.用户关闭) |
| `start_time` | int(11) | 0 | no | 开始时间 |
| `end_time` | int(11) | 0 | no | 结束时间 |
| `payload` | mediumtext | 无 | no | 请求的任务body |
| `extra` | mediumtext | 无 | yes | 队列的额外信息 |
| `successful` | mediumtext | 无 | yes | 成功的数据 |
| `unsuccessful` | mediumtext | 无 | yes | 失败的数据 |
| `unprocessed` | mediumtext | 无 | yes | 待处理的数据(通常是用户取消了) |
| `successful_count` | int(11) | 0 | no | 成功的数量 |
| `unsuccessful_count` | int(11) | 0 | no | 失败的数量 |
| `unprocessed_count` | int(11) | 0 | no | 未处理的数量 |

## `batch_task_record` - 批量操作记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `batch_task_uuid` | varchar(8) | 空串 | no | batch task uuid |
| `team_uuid` | varchar(8) | 空串 | no | 团队id |
| `number` | int(20) | 无 | no | 序号 |
| `context_type` | int(11) | 0 | no | 上下文类型, 1. import_task(project_task)<br> |
| `context_param_1` | varchar(16) | 空串 | no | 上下文参数1 |
| `context_param_2` | varchar(16) | 空串 | no | 上下文参数2 |
| `create_time` | bigint(20) | 0 | no | 创建时间(秒) |
| `extra` | mediumtext | 无 | yes | 额外信息 |
| `status` | tinyint(4) | 0 | no | 状态(1:正常 2:删除) |

## `batch_task_row`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `batch_task_uuid` | varchar(8) | 空串 | no | batch task uuid |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `number` | int(11) | 无 | no |  |
| `context_type` | int(11) | 无 | no | 上下文类型 |
| `context_param_1` | varchar(16) | 空串 | no | 上下文参数1 |
| `context_param_2` | varchar(16) | 空串 | no | 上下文参数2 |
| `create_time` | int(11) | 无 | no | 创建时间（秒） |
| `extra` | mediumtext | 无 | yes | 额外信息 |

## `blog_subscribe`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `id` | int(11) | 无 | no |  |
| `email` | varchar(128) | 无 | no |  |
| `create_time` | datetime | CURRENT_TIMESTAMP | no |  |
| `ip` | varchar(15) | 无 | yes |  |

## `blog_subscribe_count` - 博客统计

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `count_value` | int(11) | 0 | no |  |
| `id` | int(11) | 无 | no |  |

## `board` - board

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `name` | varchar(32) | 无 | no | 名字 |
| `project_uuid` | varchar(16) | 无 | no | 所属项目uuid |
| `position` | bigint(20) | 无 | no | 顺序(从小到大) |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 更新时间 |

## `build_result`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `hash` | varchar(48) | 无 | no |  |
| `build_id` | int(11) | 无 | no |  |
| `display_url` | varchar(128) | 无 | no |  |
| `result` | varchar(16) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |

## `card` - 仪表盘卡片

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `dashboard_uuid` | varchar(8) | 空串 | no | 仪表盘UUID |
| `name` | varchar(32) | 空串 | no | 卡片名称 |
| `type` | int(11) | 0 | no | 卡片类型 |
| `layout_x` | int(11) | 0 | no | x位置 |
| `layout_y` | int(11) | 0 | no | y位置 |
| `layout_w` | int(11) | 0 | no | 宽度 |
| `layout_h` | int(11) | 0 | no | 高度 |
| `status` | tinyint(4) | 0 | no | 1:正常;2:删除 |
| `config` | text | 无 | no | 卡片配置 |
| `uuid` | varchar(8) | 无 | no |  |
| `space_uuid` | varchar(8) | 无 | no |  |
| `page_uuid` | varchar(8) | 无 | no |  |
| `draft_uuid` | varchar(8) | 无 | no |  |
| `type` | int(11) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `create_time` | bigint(20) | 无 | no |  |
| `config` | mediumtext | 无 | no | 配置 |

## `cas_config` - CAS 配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no | 组织UUID |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `cas_server_url` | varchar(255) | 空串 | no | CAS服务器地址 |
| `user_name` | varchar(255) | 空串 | no | CAS配置里代表用户姓名的属性 |
| `email` | varchar(255) | 空串 | no | CAS配置里代表用户邮箱的属性 |
| `user_title` | varchar(255) | 空串 | no | CAS配置里代表用户职位的属性 |
| `create_time` | bigint(20) | 0 | no | 创建时间，单位是微秒 |
| `update_time` | bigint(20) | 0 | no | 更新时间，单位是微秒 |
| `status` | tinyint(4) | 无 | no | 1:正常 2:删除 |

## `cas_config_status` - 组织的CAS配置状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no | 组织UUID |
| `is_open` | tinyint(1) | 无 | no | cas配置是否开启 0:关闭 1:开启 |

## `commit`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `hash` | varchar(48) | 无 | no |  |
| `revision` | bigint(20) | -1 | no | svn revision |
| `owner` | varchar(128) | 无 | no |  |
| `domin` | varchar(32) | 无 | no |  |
| `repository_name` | varchar(128) | 无 | no |  |
| `branch` | varchar(128) | 无 | no |  |
| `url` | varchar(1024) | 无 | no |  |
| `message` | text | 无 | no |  |
| `timestamp` | bigint(20) | 无 | no |  |
| `scm` | varchar(3) | 无 | no | 代码管理工具 |

## `component`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `template_uuid` | varchar(8) | 空串 | no | 模板UUID |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `project_uuid` | varchar(16) | 空串 | no |  |
| `parent_uuid` | varchar(8) | 无 | yes | 组件父级 |
| `name` | varchar(64) | 空串 | no | 组件名称 |
| `name_pinyin` | varchar(256) | 空串 | no | 组件名称拼音 |
| `desc` | varchar(256) | 空串 | no | 组件描述 |
| `type` | int(11) | 无 | no | 组件类型(1. 文件夹 2.内置组件 3.场景特化<br>组件 4. 自定义任务类型组件) |
| `create_time` | int(11) | 无 | no | 创建时间 |
| `position` | int(11) | 0 | yes | 排序 |
| `objects` | varchar(512) | [] | yes | 组件源类型/组件源uuid |
| `status` | tinyint(1) | 1 | yes | 1正常 2已删除 |
| `settings` | text | 无 | yes | 功能设置 |
| `kanban_setting` | text | 无 | yes | 敏捷看板设置 |
| `url_setting` | text | 无 | yes | 自定义链接设置 |

## `component_template`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `name` | varchar(64) | 空串 | no | 组件名称 |
| `name_pinyin` | varchar(256) | 空串 | no | 组件名称拼音 |
| `desc` | varchar(256) | 空串 | no | 组件描述 |
| `type` | int(11) | 无 | no | 组件类型(1. 文件夹 2.内置组件 3.场景特化<br>组件 4. 自定义任务类型组件) |
| `category` | int(11) | 1 | yes | 组件分组 |
| `status` | int(11) | 1 | yes | 状态(1.正常 2.删除) |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `position` | int(11) | 0 | yes | 位置 |
| `objects` | varchar(512) | [] | yes | 组件源类型/组件源uuid |

## `component_view`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 空串 | no |  |
| `component_uuid` | varchar(8) | 空串 | no | 组件uuid |
| `name` | varchar(64) | 空串 | no | 视图名称 |
| `layout` | tinyint(1) | 3 | no | 布局方式(1.表格 2.宽详情 3.窄详情) |
| `query` | text | 无 | yes | 筛选方式 |
| `condition` | text | 无 | no | 筛选条件 |
| `sort` | varchar(256) | 空串 | yes | 排序方式 |
| `group_by` | varchar(256) | 空串 | yes | 分组方式 |
| `create_time` | bigint(20) | 无 | no |  |
| `built_in` | tinyint(1) | 1 | no | 是否默认视图 |
| `include_subtasks` | tinyint(1) | 0 | no | 是否包含子任务 |
| `table_field_settings` | text | 无 | yes | 表格模式的表头 |
| `board_settings` | text | 无 | yes | 看板视图设置 |
| `is_fold_all_groups` | tinyint(1) | 0 | no | 是否折叠分组 |
| `status` | tinyint(1) | 1 | yes | 1正常 2已删除 |
| `display_type` | tinyint(4) | 0 | yes | 1:平铺展示, 2:子树展示，3:折叠展示 |
| `is_show_derive` | tinyint(1) | 0 | no | 是否显示派生 |

## `component_view_user`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `owner` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 空串 | no |  |
| `component_uuid` | varchar(8) | 空串 | no | 组件uuid |
| `name` | varchar(64) | 空串 | no | 视图名称 |
| `layout` | tinyint(1) | 3 | no | 布局方式(1.表格 2.宽详情 3.窄详情) |
| `query` | text | 无 | yes | 筛选方式 |
| `condition` | text | 无 | no | 筛选条件 |
| `sort` | varchar(256) | 空串 | yes | 排序方式 |
| `group_by` | varchar(256) | 空串 | yes | 分组方式 |
| `create_time` | bigint(20) | 无 | no |  |
| `built_in` | tinyint(1) | 0 | no | 是否默认视图 |
| `include_subtasks` | tinyint(1) | 0 | no | 是否包含子任务 |
| `table_field_settings` | text | 无 | yes | 表格模式的表头 |
| `board_settings` | text | 无 | yes | 看板视图设置 |
| `is_fold_all_groups` | tinyint(1) | 0 | no | 是否折叠分组 |
| `status` | tinyint(1) | 1 | yes | 1正常 2已删除 |
| `project_filter_uuid` | varchar(8) | 空串 | yes | 项目下用户的筛选器uuid |
| `display_type` | tinyint(4) | 无 | yes | 1:平铺展示, 2:子树展示，3:折叠展示 |
| `is_show_derive` | tinyint(1) | 0 | no | 是否显示派生 |

## `component_view_user_config`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `component_uuid` | varchar(8) | 无 | no | 组件uuid |
| `owner` | varchar(8) | 无 | no |  |
| `view_uuid` | varchar(8) | 空串 | no | 视图uuid |
| `type` | tinyint(1) | 无 | no | 1.组件视图 2.用户视图 |
| `position` | int(11) | 无 | no | 视图位置 |
| `is_show` | tinyint(1) | 1 | no | 1.显示 0.隐藏 |

## `converter_config_status` - 组织的converter配置状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no | 组织UUID |
| `is_open` | tinyint(1) | 无 | no | 是否开启 0:关闭 1:开启 |

## `corp_sync_department` - 某些第三方如LDAP没有提供部门sync_id的功能, 这里提供一种自管理sync_id的方式

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `department_key` | varchar(255) | 无 | no | 第三方登录企业id下标识唯一部门的关键词 |
| `sync_id` | bigint(20) | 无 | no | 同步部门的id |

## `create_team_code` - 创建团队邮箱验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(32) | 空串 | no | 验证码 |
| `email` | varchar(255) | 空串 | no | 邮箱地址 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `ip` | varchar(15) | 空串 | no | 请求IP地址 |
| `status` | tinyint(4) | 无 | no | 1:正常;2:无效 |
| `referrer` | text | 无 | no | 渠道信息 |
| `keyword` | text | 无 | no | 搜索关键词 |

## `crm_permission`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `email` | varchar(32) | 无 | no |  |

## `csm`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `email` | varchar(128) | 无 | no |  |
| `name` | varchar(24) | 无 | no |  |
| `title` | varchar(24) | 无 | no |  |
| `phone` | varchar(32) | 无 | no |  |
| `status` | tinyint(4) | -1 | no |  |

## `dashboard` - 仪表盘

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `owner` | varchar(8) | 空串 | no | 创建者UUID |
| `name` | varchar(32) | 空串 | no | 仪表盘名称 |
| `name_pinyin` | varchar(256) | 空串 | no | 仪表盘名称拼音 |
| `shared` | tinyint(4) | 0 | no | 是否共享 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `status` | tinyint(4) | 0 | no | 1:正常;2:删除 |

## `dashboard_member` - 仪表盘成员（被分享者）

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `dashboard_uuid` | varchar(8) | 空串 | no | 仪表盘UUID |
| `user_domain_param` | varchar(8) | 空串 | no | 成员域参数 |
| `user_domain_type` | tinyint(4) | 0 | no | 成员域类型 |
| `position` | int(11) | 0 | no | 位置 |

## `dashboard_preference` - 仪表盘用户配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `user_uuid` | varchar(8) | 无 | no | 用户UUID |
| `dashboard_uuid` | varchar(8) | 空串 | no | 仪表盘UUID |
| `pinned` | tinyint(4) | 0 | no | 是否常用仪表盘 |
| `is_default` | tinyint(4) | 0 | no | 是否默认仪表盘 |

## `department`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `parent_uuid` | varchar(8) | 空串 | yes | 父节点id |
| `name` | varchar(32) | 无 | no | 部门名称 |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `next_uuid` | varchar(8) | 空串 | yes | 兄弟节点id |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `sync_id` | bigint(20) | 无 | yes | 同步部门的id |
| `corp_uuid` | varchar(64) | 无 | yes | 第三方登录企业id |

## `department_member`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `department_uuid` | varchar(8) | 无 | yes |  |
| `user_uuid` | varchar(8) | 无 | yes |  |

## `dingding_sync_relation`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `corp_uuid` | varchar(64) | 空串 | no | 同步企业的id |
| `unionid` | varchar(128) | 空串 | no |  |
| `corp_name` | varchar(255) | 无 | no | 第三方企业名称 |
| `sync_id` | varchar(255) | 空串 | no | 同步账号的id |
| `status` | tinyint(4) | -1 | no | 状态:1:正常 2:删除 |

## `domain_info`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `id` | int(11) | 无 | no |  |
| `imap` | varchar(512) | 无 | yes |  |
| `smtp` | varchar(512) | 无 | yes |  |
| `pop` | varchar(512) | 无 | yes |  |
| `mailboxes` | varchar(512) | 无 | yes |  |
| `domainMatch` | varchar(2048) | 无 | yes |  |
| `mx-match` | varchar(64) | 无 | yes |  |

## `email_code` - 用户邮箱验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(32) | 空串 | no | 验证码 |
| `user_uuid` | varchar(8) | 无 | no | 用户ID |
| `email` | varchar(255) | 无 | no | 邮箱地址 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `ip` | varchar(15) | 无 | no | 请求IP地址 |
| `status` | tinyint(4) | 无 | no | 1:正常;2:无效 |

## `exchange_code` - 兑换码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(8) | 无 | no |  |
| `email` | varchar(128) | 无 | no |  |
| `phone` | varchar(128) | 无 | no |  |
| `team_name` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `user_name` | varchar(8) | 无 | no |  |
| `is_pay` | tinyint(1) | 无 | no | 是否付费 |
| `source` | varchar(128) | 无 | no | 来源 |
| `exchange_time` | int(11) | 无 | no | 团队有效期(天) |
| `create_time` | bigint(20) | 无 | no | 创建时间(秒) |
| `status` | tinyint(4) | 无 | no |  |
| `expired_time` | bigint(11) | 0 | no | 兑换码有效期(具体时间的秒数) |
| `scale` | int(6) | 无 | no | 团队成员规模 |
| `csm` | varchar(128) | 空串 | yes | 负责的CSM |

## `field` - 全局属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `type` | int(11) | 0 | no | 类型 |
| `name` | varchar(32) | 空串 | no | 属性名 |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `description` | text | 无 | no | 属性描述 |
| `default_value` | text | 无 | yes | 属性默认值 |
| `renderer` | int(11) | 0 | no | 属性显示方式 |
| `filter_option` | int(11) | 0 | no | 属性过滤器选项(0:不可过滤,1:可过滤) |
| `search_option` | int(11) | 0 | no | 属性搜索选项(0:不可搜索,1:可搜索) |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `built_in` | tinyint(1) | 0 | no | 系统内置属性 |
| `step_settings` | varchar(256) | 无 | yes | 属性间隔时间设置 |
| `stay_settings` | varchar(256) | 无 | yes | 属性停留时间与次数设置 |
| `related_type` | int(11) | 0 | yes | 关联类型 |
| `related_uuid` | varchar(8) | 无 | yes | 关联的uuid |
| `status` | tinyint(4) | 无 | yes | 1:正常;null:删除 |

## `field_access_rule` - 属性权限规则

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `issue_type_uuid` | varchar(8) | 空串 | no | 任务类型UUID |
| `field_uuid` | varchar(8) | 无 | no | 属性UUID |
| `status_uuid` | varchar(8) | 空串 | no | 任务状态UUID |
| `user_domain_type` | int(11) | 0 | no | 用户域类型 |
| `user_domain_param` | varchar(16) | 空串 | no | 用户域参数 |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `access` | int(11) | 0 | no | 访问权限 |

## `field_config` - 属性配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `issue_type_uuid` | varchar(8) | 空串 | no | 任务类型UUID |
| `field_uuid` | varchar(8) | 无 | no | 属性UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `default_value` | text | 无 | yes | 属性默认值 |
| `renderer` | int(11) | 0 | no | 属性显示方式 |
| `filter_option` | int(11) | 0 | no | 属性过滤器选项(0:不可过滤,1:可过滤) |
| `search_option` | int(11) | 0 | no | 属性搜索选项(0:不可搜索,1:可搜索) |
| `required` | tinyint(4) | 0 | no | 属性是否必填(0:否,1:是) |
| `can_modify_required` | tinyint(1) | 无 | no | 是否可以修改 选填 |
| `can_delete` | tinyint(1) | 无 | no | 是否可以删除 |
| `position` | int(11) | 0 | no | 属性位置 |

## `field_multi_option_value` - 任务多选属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 空串 | no |  |
| `field_uuid` | varchar(8) | 空串 | no |  |
| `type` | int(11) | 无 | no | 选项类型 |
| `value` | text | 无 | yes | 选项值 |
| `position` | int(11) | 0 | yes | 选项排序 |
| `status` | tinyint(4) | 1 | yes | 1: 正常 2: 删除 3: 储存 |

## `field_option` - 属性选项，用于单选或者多选类型的属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `field_uuid` | varchar(8) | 空串 | no | 属性UUID |
| `value` | varchar(64) | 空串 | no | 属性值 |
| `default_selected` | tinyint(4) | 0 | no | 是否默认选中 |
| `background_color` | varchar(9) | 无 | no |  |
| `color` | varchar(9) | 无 | no |  |
| `position` | int(11) | 0 | no | 排序序号 |
| `status` | tinyint(4) | 无 | yes | 1:正常;2:删除 |
| `desc` | varchar(128) | 空串 | yes | 属性值描述 |

## `field_option_config` - 属性配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `issue_type_uuid` | varchar(8) | 空串 | no | 属性类型UUID |
| `field_uuid` | varchar(8) | 无 | no | 属性UUID |
| `option_uuid` | varchar(8) | 空串 | no | 属性选项UUID |
| `position` | int(11) | 0 | no | 属性选项排序序号 |
| `default_selected` | tinyint(4) | 0 | no | 属性选项是否默认选中 |

## `field_value` - 任务属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 空串 | no | 任务UUID |
| `field_uuid` | varchar(8) | 空串 | no | 属性UUID |
| `type` | int(11) | 0 | no | 值类型 |
| `value` | text | 无 | yes | 属性值 |
| `number_value` | bigint(20) | 无 | yes | 数字属性值 |
| `status` | int(11) | 0 | no | 状态(1:正常,2:删除) |

## `field_value_history`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `task_uuid` | varchar(16) | 空串 | no |  |
| `field_uuid` | varchar(8) | 空串 | no |  |
| `type` | int(11) | 0 | no | field type |
| `value` | varchar(16) | 空串 | no | field的值，超过16位保存hash值 |
| `valid_from` | int(11) | 无 | no | 更新时间 |
| `valid_to` | int(11) | 无 | yes | 失效时间 |

## `file`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `hash` | varchar(28) | 无 | no |  |
| `type` | int(11) | 无 | no |  |
| `mime` | varchar(255) | 无 | yes |  |
| `size` | bigint(20) | -1 | no |  |
| `image_width` | int(11) | -1 | no |  |
| `image_height` | int(11) | -1 | no |  |
| `exif` | varchar(64) | 空串 | no | 文件 exif 属性 |

## `file_resource_view` - VIEW

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `reference_type` | tinyint(4) | 无 | no |  |
| `reference_id` | varchar(16) | 无 | yes |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `project_uuid` | varchar(16) | 无 | no |  |
| `owner_uuid` | varchar(8) | 无 | yes |  |
| `type` | tinyint(4) | -1 | no |  |
| `source` | tinyint(4) | 0 | no |  |
| `name` | varchar(255) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `modify_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | -1 | no |  |
| `description` | varchar(64) | 无 | no |  |
| `hash` | varchar(28) | 无 | no |  |
| `mime` | varchar(255) | 无 | yes |  |
| `size` | bigint(20) | -1 | no |  |
| `image_width` | int(11) | -1 | no |  |
| `image_height` | int(11) | -1 | no |  |
| `exif` | varchar(64) | 空串 | no | 文件 exif 属性 |

## `filter`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `owner` | varchar(8) | 无 | yes |  |
| `name` | varchar(128) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `query` | text | 无 | no |  |
| `sort` | text | 无 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `include_subtasks` | tinyint(4) | 0 | no | 是否包含子任务 |

## `gantt_chart`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | gantt chart uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `name` | varchar(32) | 无 | no | gantt chart名称 |
| `name_pinyin` | varchar(256) | 空串 | no | gantt chart名称拼音 |
| `owner` | varchar(8) | 无 | no | 创建者 uuid |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `shared` | tinyint(4) | 0 | no | 是否共享 |
| `import_config` | text | 无 | yes | 导入配置 |
| `status` | tinyint(4) | 无 | yes | 1:正常 2:删除 |
| `sync_from_project` | tinyint(4) | 0 | no | 开关：关联甘特数据同步project的变动 |
| `sync_to_project` | tinyint(4) | 0 | no | 开关：关联甘特数据变动同步到project |

## `gantt_chart_field_value` - 甘特图属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `gantt_chart_uuid` | varchar(16) | 无 | no | 甘特图UUID |
| `alias` | varchar(32) | 无 | no | 别名 |
| `type` | int(11) | 无 | no | 值类型 |
| `value` | text | 无 | yes | 属性值 |

## `gantt_chart_personal_config`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `gantt_chart_uuid` | varchar(8) | 无 | no | gantt chart uuid |
| `user_uuid` | varchar(8) | 无 | no | 用户 uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `expand` | tinyint(4) | 1 | no | 0:收起  1: 展开 |
| `zooming` | tinyint(4) | 0 | no | 缩放 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `status` | tinyint(4) | 1 | no | 1: 正常 2: 删除 |

## `gantt_data`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | 甘特图数据 uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `gantt_chart_uuid` | varchar(8) | 无 | no | gantt chart uuid |
| `type` | tinyint(4) | 0 | yes | 1:甘特图数据类型, 1:任务 2:分组 3:里程碑<br> |
| `name` | varchar(128) | 无 | no | 名称 |
| `name_pinyin` | varchar(512) | 空串 | no | 名称拼音 |
| `owner` | varchar(8) | 无 | no | 创建者 uuid |
| `start_time` | bigint(20) | 无 | no | 开始时间 |
| `end_time` | bigint(20) | 无 | no | 结束时间 |
| `progress` | int(11) | 0 | no | 进度 |
| `parent` | varchar(8) | 无 | no | 父节点 uuid |
| `path` | varchar(169) | 无 | no | 层级关系 |
| `position` | bigint(20) | 0 | no | 显示位置 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `status` | tinyint(4) | 无 | yes | 1:正常 2:删除 |
| `assign` | varchar(8) | 无 | yes | 负责人uuid |
| `related_type` | tinyint(4) | 0 | no | 关联数据类型 |
| `related_project_uuid` | varchar(16) | 无 | no | 关联的项目uuid |
| `related_sprint_uuid` | varchar(8) | 无 | no | 关联的迭代uuid |

## `important_field` - 关键属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `issue_type_uuid` | varchar(8) | 空串 | no | 任务类型UUID |
| `field_uuid` | varchar(8) | 无 | no | 属性UUID |
| `position` | int(11) | 0 | no | 属性位置 |

## `invitation`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(32) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `inviter_uuid` | varchar(8) | 无 | no |  |
| `email` | varchar(255) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | 无 | no |  |

## `issue_type` - 任务类型

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `name` | varchar(32) | 空串 | no | 任务类型名 |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `icon` | int(11) | 0 | no | 任务类型图标编号 |
| `built_in` | tinyint(4) | 0 | no | 是否系统内置任务类型 |
| `default_selected` | tinyint(4) | 0 | no | 是否默认选中 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `status` | tinyint(4) | 无 | yes | 1:正常;null:删除 |
| `default_configs` | longtext | 无 | no | 默认配置 |
| `type` | tinyint(4) | 0 | yes | 0:标准任务类型, 1:子任务类型 |
| `detail_type` | tinyint(4) | 0 | yes | issue_type的具体类型: 0:自定义类型, 1:需<br>求, 2:任务, 3:缺陷, 4:子任务, 5:子需求 |

## `kanban_status_sort`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 无 | no |  |
| `issue_type_uuid` | varchar(8) | 无 | no |  |
| `status_uuid` | varchar(8) | 无 | no |  |
| `position` | bigint(20) | 无 | no |  |

## `kanban_task_sort`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 无 | no |  |
| `position` | bigint(20) | 无 | no |  |

## `ldap_config` - LDAP 配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no | 组织UUID |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `host` | varchar(255) | 空串 | no | LDAP服务器 |
| `port` | int(11) | -1 | no | LDAP服务器端口 |
| `encryption` | varchar(10) | 空串 | no | 加密方式(plain,ssl,starttls) |
| `domain_base` | varchar(255) | 空串 | no | 需要同步的LDAP根节点 |
| `user_id` | varchar(255) | 空串 | no | LDAP配置里代表用户唯一ID的属性 |
| `user_name` | varchar(255) | 空串 | no | LDAP配置里代表用户姓名的属性 |
| `email` | varchar(255) | 空串 | no | LDAP配置里代表用户邮箱的属性 |
| `domain_search_user` | varchar(255) | 空串 | no | 执行LDAP搜索时使用的用户账户 |
| `domain_search_password` | varchar(255) | 空串 | no | 执行LDAP搜索时使用的用户密码 |
| `user_filter` | varchar(1024) | 空串 | no | 用户过滤规则 |
| `create_time` | bigint(20) | 0 | no | 创建时间，单位是微秒 |
| `update_time` | bigint(20) | 0 | no | 更新时间，单位是微秒 |
| `status` | tinyint(4) | 无 | no | 1:正常 2:删除 |
| `domain` | varchar(255) | 空串 | no | AD 需要域的概念 |
| `type` | int(11) | 1 | no | LDAP模式 1.LDAP 2.AD |
| `way_to_get_mail` | int(11) | 1 | no | 获取邮箱方式 1.按属性值 2.按用户登录名@域<br>名方式 |
| `enable_sync` | int(1) | 0 | no | 是否开启同步 |
| `department_filter` | varchar(1024) | 空串 | no | 部门过滤规则 |
| `avatar` | varchar(255) | 空串 | no | LDAP配置里代表用户头像的属性 |
| `user_title` | varchar(255) | 空串 | no | LDAP配置里代表用户职位的属性 |

## `ldap_config_status` - 组织的LDAP配置状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no | 组织UUID |
| `is_open` | tinyint(1) | 无 | no | ldap配置是否开启 0:关闭 1:开启 |

## `lint_issue` - lint问题记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(22) | 空串 | no | uuid |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `hash` | varchar(28) | 空串 | no | 用于跟踪相同issue的哈希 |
| `create_time` | bigint(20) | 无 | no | 首次出现时间（秒） |
| `linter` | varchar(20) | 空串 | no | 使用的linter |
| `language` | varchar(20) | 空串 | no | 使用的语言 |
| `file` | varchar(4096) | 空串 | no | 文件相对路径 |
| `line` | int(11) | -1 | no | 文件行号 |
| `column` | int(11) | -1 | no | 文件列号，没有列号则为-1 |
| `severity` | tinyint(4) | -1 | no | 严重程度(1:ignore 2:info 3:warning 4:err<br>or 5:fatal) |
| `category` | varchar(255) | 空串 | no | 所属类别，可能为空 |
| `rule` | varchar(255) | 空串 | no | 规则名称，可能为空 |
| `message` | varchar(4096) | 空串 | no | lint提示信息/描述 |
| `author` | varchar(40) | 空串 | no | 问题代码提交者名称 |

## `manhour`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `task_uuid` | varchar(16) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `start_time` | bigint(20) | 0 | no |  |
| `hours` | bigint(20) | 0 | no |  |
| `remark` | varchar(1024) | 无 | no |  |
| `status` | tinyint(4) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |

## `manhour_report`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `project_uuid` | varchar(16) | 空串 | no | 项目uuid |
| `name` | varchar(128) | 空串 | no | 报表名称 |
| `owner` | varchar(8) | 无 | no | 创建者uuid |
| `config` | text | 无 | no | 项目报表配置 |
| `create_time` | bigint(20) | 0 | no | 创建时间，秒 |
| `update_time` | bigint(20) | 0 | no | 更新时间，秒 |
| `fold` | tinyint(1) | 0 | no | 0:展开, 1:折叠 |
| `status` | tinyint(4) | 0 | no | 状态(1:正常, 2:删除) |

## `message`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(32) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `reference_type` | tinyint(4) | -1 | no |  |
| `reference_id` | varchar(16) | 无 | yes |  |
| `from_uuid` | varchar(16) | 无 | yes |  |
| `to_uuid` | varchar(16) | 无 | yes |  |
| `send_time` | bigint(20) | 0 | no |  |
| `message` | text | 无 | no |  |
| `type` | tinyint(4) | 无 | no |  |
| `resource_uuid` | varchar(8) | 无 | yes |  |
| `subject_type` | tinyint(4) | 无 | no |  |
| `subject_id` | varchar(16) | 无 | yes |  |
| `action` | tinyint(4) | 无 | no |  |
| `object_type` | tinyint(4) | 无 | no |  |
| `object_id` | varchar(16) | 无 | yes |  |
| `object_name` | text | 无 | no |  |
| `object_attr` | tinyint(4) | 无 | no |  |
| `old_value` | text | 无 | no |  |
| `new_value` | text | 无 | no |  |
| `ext` | text | 无 | no |  |
| `uuid` | varchar(8) | 无 | no | uuid |
| `user_uuid` | varchar(8) | 无 | no | 发送者 |
| `space_uuid` | varchar(8) | 无 | no | space uuid |
| `page_uuid` | varchar(8) | 无 | no | page uuid |
| `send_time` | bigint(20) | 无 | no | 创建时间 |
| `action` | tinyint(4) | 无 | no | 操作动作 |
| `type` | tinyint(4) | 无 | no | 类型(评论消息,系统消息等) |
| `title` | varchar(64) | 无 | no | 标题 |
| `message` | varchar(512) | 无 | no |  |
| `object_attr` | tinyint(4) | 无 | no |  |
| `ext` | varchar(512) | 无 | no |  |

## `mq_failed_message`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `body` | text | 无 | no |  |
| `error` | text | 无 | no |  |
| `create_time` | bigint(20) | 无 | no |  |

## `notice`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `task_uuid` | varchar(16) | 无 | no |  |
| `message_uuid` | varchar(8) | 无 | no |  |
| `status` | tinyint(4) | 无 | no |  |
| `server_update_stamp` | bigint(20) | 0 | no |  |

## `notice_config` - 任务通知配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | 任务通知配置的uuid |
| `status` | tinyint(4) | 无 | no | 1:正常 2:无效 |
| `create_time` | bigint(20) | 无 | no | 创建时间, 秒 |
| `update_time` | bigint(20) | 无 | no | 更新时间，秒 |

## `notice_config_sending_method` - 任务通知配置-发送方式

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `notice_config_uuid` | varchar(8) | 无 | no | 任务通知配置的uuid |
| `method` | enum('notice_center','email','dingding',<br>'wechat') | 无 | no | 通知发送的方式（消息中心、邮件、钉钉、企业微<br>信） |
| `status` | tinyint(4) | 无 | no | 1:正常 2:无效 |

## `notice_config_user_domain` - 接收通知的用户域

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `notice_config_uuid` | varchar(8) | 无 | no | 任务通知配置的uuid |
| `user_domain_type` | int(11) | 0 | no | 用户域类型 |
| `user_domain_param` | varchar(16) | 无 | no | 用户域参数 |
| `status` | tinyint(4) | 无 | no | 1:正常 2:无效 |

## `notification_rule`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `scheme_uuid` | varchar(8) | 空串 | no | 所属通知配置方案uuid |
| `create_time` | bigint(20) | 0 | no | 创建时间(秒) |
| `user_domain_type` | int(11) | 0 | no | 用户域类型 |
| `user_domain_param` | varchar(16) | 空串 | no | 用户域参数 |
| `event` | int(11) | 0 | no | 事件编号 |
| `email` | tinyint(4) | 0 | no | 邮件选项(1:开启 2:关闭) |
| `badge` | tinyint(4) | 0 | no | 红点选项(1:开启 2:关闭) |
| `alert` | tinyint(4) | 0 | no | 推送通知选项(1:开启 2:关闭) |
| `status` | tinyint(4) | 0 | no | 状态(1:正常 2:删除) |

## `notification_scheme`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `name` | varchar(32) | 空串 | no | 通知配置方案名称 |
| `description` | text | 无 | no | 通知配置方案描述 |
| `status` | tinyint(4) | 0 | no | 状态(1:正常 2:删除) |

## `ones_app`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(64) | 无 | no | 每一个私有部署客户对应一个ones_appid(one<br>s_app_uuid) |
| `name` | varchar(256) | 无 | no | 每一个私有部署客户的名称（组织名、团队名） |

## `ones_app_base_url`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `ones_app_uuid` | varchar(64) | 无 | no | 每一个私有部署客户对应一个ones_appid(one<br>s_app_uuid) |
| `service_name` | varchar(128) | 无 | no | 服务名称（例子： project-api wiki-api） |
| `create_time` | int(11) | 0 | no | 创建时间（秒） |
| `update_time` | int(11) | 0 | no | 更新时间（秒） |
| `base_url` | text | 无 | no |  |

## `organization`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | -1 | no |  |
| `name` | varchar(16) | 无 | no |  |
| `owner` | varchar(8) | 无 | yes |  |
| `logo` | varchar(255) | 无 | no |  |
| `expire_time` | bigint(20) | 无 | no |  |
| `scale` | int(6) | 无 | no |  |
| `csm` | varchar(128) | 空串 | yes |  |
| `type` | tinyint(4) | -1 | no |  |
| `visibility` | tinyint(1) | 无 | no |  |

## `organization_update_stamp` - 组织相关数据更新记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no | 组织id |
| `type` | tinyint(4) | 0 | no | 更新类型 |
| `server_update_stamp` | bigint(20) | 0 | no | 更新时间戳 |

## `permission_rule`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | 权限规则id |
| `team_uuid` | varchar(8) | 空串 | no | 团队id |
| `org_uuid` | varchar(8) | 空串 | no | 组织uuid |
| `context_type` | int(11) | 0 | no | 权限上下文类型 |
| `context_param_1` | varchar(16) | 空串 | no | 权限上下文参数1 |
| `context_param_2` | varchar(16) | 空串 | no | 权限上下文参数2 |
| `user_domain_type` | int(11) | 0 | no | 用户域类型 |
| `user_domain_param` | varchar(16) | 无 | no | 用户域参数 |
| `permission` | int(11) | 0 | no | 权限编号 |
| `create_time` | bigint(20) | 0 | no | 创建时间(秒) |
| `status` | tinyint(4) | 0 | no | 状态(1:正常 2:删除) |
| `read_only` | tinyint(4) | 0 | no | 是否只读 |
| `position` | int(11) | 0 | no | 权限显示位置 |

## `phone_code` - 手机号验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `phone` | varchar(32) | 空串 | no | 手机号 |
| `code` | varchar(6) | 空串 | no | 验证码 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `ip` | varchar(15) | 空串 | no | 请求IP地址 |
| `status` | tinyint(4) | 无 | no | 1:正常;2:无效 |

## `pipeline` - pipeline

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `token` | varchar(22) | 无 | no | 回调接口使用的token |
| `name` | varchar(128) | 空串 | no | 名字 |
| `name_pinyin` | varchar(1024) | 空串 | no | 名字拼音 |
| `owner` | varchar(8) | 无 | no | 创建者uuid |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `status` | tinyint(4) | -1 | no | 1:正常 2:删除 |
| `scm_status` | tinyint(4) | -1 | no | scm配置状态(1:未知 2:成功) |
| `scm_secret_key` | varchar(22) | 空串 | no | scm github/gitlab 密钥 |
| `scm_last_success_time` | bigint(20) | 无 | no | scm上次成功时间 |
| `ci_status` | tinyint(4) | -1 | no | ci配置状态(1:未知 2:成功) |
| `ci_last_success_time` | bigint(20) | 无 | no | ci上次成功时间 |
| `lint_status` | tinyint(4) | -1 | no | lint配置状态(1:未知 2:成功) |
| `lint_last_success_time` | bigint(20) | 无 | no | lint上次成功时间 |
| `test_status` | tinyint(4) | -1 | no | test配置状态(1:未知 2:成功) |
| `test_last_success_time` | bigint(20) | 无 | no | test上次成功时间 |
| `artifact_status` | tinyint(4) | -1 | no | artifact配置状态(1:未知 2:成功) |
| `artifact_last_success_time` | bigint(20) | 无 | no | artifact上次成功时间 |
| `sprint_binding_rule` | varchar(1024) | 空串 | no | 迭代绑定脚本 |

## `pipeline_run` - pipeline 单次执行

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `pipeline_uuid` | varchar(8) | 空串 | no | pipeline uuid |
| `number` | int(11) | 无 | no | 执行编号 |
| `repository` | varchar(255) | 空串 | no | 代码仓库名称 |
| `branch` | varchar(255) | 空串 | no | 分支名称 |
| `start_time` | bigint(20) | 无 | no | 开始时间 |
| `finish_time` | bigint(20) | 无 | no | 结束时间 |
| `build_id` | varchar(255) | 空串 | no | ci构建id |
| `build_url` | varchar(2048) | 空串 | no | ci构建url |
| `hint` | varchar(2048) | 空串 | no | 构建携带的额外信息 |
| `status` | int(11) | 无 | no | pipeline 执行状态 |
| `log` | mediumtext | 无 | no | 构建log |

## `pipeline_run_artifact` - pipeline 单次执行绑定的工件

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `run_number` | int(11) | 无 | no | 流水线执行编号 |
| `env` | varchar(32) | 空串 | no | 环境 |
| `name` | varchar(255) | 空串 | no | 工件名称 |
| `url` | varchar(2048) | 空串 | no | 工件url |
| `create_time` | bigint(20) | -1 | no | 创建时间 |

## `pipeline_run_commit` - pipeline单次执行和commit绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `run_number` | int(11) | 无 | no | 流水线执行编号 |
| `commit_uuid` | varchar(8) | 空串 | no | 代码提交uuid |

## `pipeline_run_lint_issue` - pipeline单次执行和lint issue绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `run_number` | int(11) | 无 | no | 流水线执行编号 |
| `issue_uuid` | varchar(22) | 无 | no | 问题uuid |
| `status` | tinyint(4) | 无 | no | 状态(1:开启 2:关闭) |
| `create_time` | bigint(20) | 无 | no | 创建时间（秒） |

## `pipeline_run_test_case` - pipeline单次执行和自动化测试用例绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `run_number` | int(11) | 无 | no | 流水线执行编号 |
| `case_uuid` | varchar(22) | 空串 | no | 用例uuid |
| `result` | tinyint(4) | 无 | no | 用例执行结果 |
| `start_time` | bigint(20) | 无 | no | 开始时间（毫秒） |
| `finish_time` | bigint(20) | 无 | no | 结束时间（毫秒） |
| `message` | text | 无 | no | 用例执行信息 |

## `pipeline_stage` - pipeline 单次执行阶段

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `pipeline_uuid` | varchar(8) | 空串 | no | pipeline uuid |
| `run_number` | int(11) | 无 | no | 执行编号 |
| `type` | int(11) | 无 | no | 阶段类型(1:start 2:finish 3:build 4:lint<br> 5:test 6:deploy) |
| `name` | varchar(255) | 空串 | no | 阶段名称 |
| `start_time` | bigint(20) | 无 | no | 开始时间 |
| `finish_time` | bigint(20) | 无 | no | 结束时间 |
| `status` | int(11) | 无 | no | 运行状态(1:未知 2:成功 3:失败) |

## `pipeline_test_case` - pipeline自动化测试用例记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(22) | 空串 | no | uuid |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `hash` | varchar(28) | 空串 | no | 用于跟踪相同测试用例的哈希 |
| `create_time` | bigint(20) | 无 | no | 首次出现时间（秒） |
| `language` | varchar(20) | 无 | no | 使用的语言 |
| `framework` | varchar(20) | 空串 | no | 使用的测试框架 |
| `id` | varchar(255) | 空串 | no | 用例 id |
| `name` | varchar(255) | 无 | no | 用例名称 |
| `description` | varchar(4096) | 空串 | no | 用例描述 |

## `plugin`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `name` | varchar(64) | 空串 | no |  |
| `dl` | varchar(128) | 空串 | no |  |
| `uuid` | varchar(8) | 无 | no | uuid |
| `name` | varchar(64) | 无 | no |  |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `avatar_url` | varchar(128) | 无 | no |  |
| `summary` | varchar(128) | 无 | no |  |
| `status` | int(11) | 无 | no | 状态 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |

## `product_authorization`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `org_uuid` | varchar(8) | 空串 | no |  |
| `type` | tinyint(4) | 无 | no |  |
| `status` | tinyint(4) | 1 | no |  |

## `product_trial_info` - 申请使用产品

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `email` | varchar(64) | 空串 | no | 邮箱 |
| `phone` | varchar(64) | 空串 | no | 手机号 |
| `nickname` | varchar(128) | 空串 | no | 用户昵称 |
| `team_name` | varchar(128) | 空串 | no | 团队名 |
| `team_scale` | varchar(128) | 空串 | no | 团队规模 |
| `industry` | varchar(128) | 空串 | no | 所处行业 |
| `status` | int(11) | 0 | no | 状态 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |

## `program`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | 甘特图数据 uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `owner` | varchar(8) | 无 | no | 创建者 uuid |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `name` | varchar(64) | 无 | no | 名称 |
| `name_pinyin` | varchar(512) | 空串 | no | 名称拼音 |
| `plan_start_time` | bigint(20) | 无 | yes | 开始时间 |
| `plan_end_time` | bigint(20) | 无 | yes | 结束时间 |
| `parent` | varchar(8) | 无 | no | 父节点 uuid |
| `path` | varchar(169) | 无 | no | 层级关系 |
| `position` | bigint(20) | 0 | no | 显示位置 |
| `status` | tinyint(4) | 无 | yes | 1:正常 2:删除 |
| `assign` | varchar(8) | 无 | yes | 负责人uuid |
| `related_type` | tinyint(4) | 0 | no | 关联数据类型 |
| `related_project_uuid` | varchar(16) | 无 | no | 关联的项目uuid |

## `program_role` - 角色

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | UUID |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `context` | varchar(32) | 空串 | no | 上下文 |
| `type` | int(11) | 0 | no | 类型 |
| `name` | varchar(32) | 无 | no | 角色名 |
| `name_pinyin` | varchar(256) | 无 | no | 角色名拼音 |
| `create_time` | bigint(20) | 无 | no | 创建时间（秒） |
| `built_in` | tinyint(4) | 0 | no | 是否系统内置角色（1:是 2:否） |
| `status` | tinyint(4) | 无 | yes | 状态（1:正常 NULL:删除） |

## `program_role_member` - 角色成员

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `role_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |

## `project`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(16) | 无 | no |  |
| `name` | varchar(128) | 无 | no |  |
| `name_pinyin` | varchar(1024) | 无 | no |  |
| `owner` | varchar(8) | 无 | yes |  |
| `assign` | varchar(8) | 无 | no | 项目负责人 |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `access_mode` | tinyint(4) | 无 | no |  |
| `status_uuid` | varchar(16) | 无 | no | 状态alias |
| `status` | tinyint(4) | 无 | no |  |
| `announcement` | text | 无 | no |  |
| `announcement_text` | text | 无 | no |  |
| `plan_start_time` | bigint(20) | 无 | yes | 计划开始时间 |
| `plan_end_time` | bigint(20) | 无 | yes | 计划结束时间 |
| `deadline` | bigint(20) | 0 | no |  |
| `template_uuid` | varchar(8) | 空串 | no | 模板UUID |
| `is_open_email_notify` | tinyint(1) | 无 | no | 邮件通知开关 |

## `project_browse` - 用户浏览项目记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `user_uuid` | varchar(8) | 无 | no | 用户uuid |
| `project_uuid` | varchar(16) | 无 | no | 项目uuid |
| `browse_time` | bigint(20) | 0 | no | 浏览时间 |

## `project_daily_stats` - 项目按天统计记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `project_uuid` | varchar(16) | 无 | no | 项目uuid |
| `type` | tinyint(4) | 0 | no | 统计类型 |
| `create_time` | bigint(20) | 0 | no | 统计时间 |
| `day` | tinyint(4) | 0 | no | 统计时间天(几号) |
| `count` | text | 无 | no | 统计结果 |

## `project_field` - 项目属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | UUID |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `context` | varchar(32) | 空串 | no | 上下文 |
| `alias` | varchar(16) | 空串 | no | 别名 |
| `status` | tinyint(4) | 无 | yes | 状态（1:正常 NULL:删除） |
| `type` | int(11) | 无 | no | 类型 |
| `name` | varchar(32) | 无 | no | 属性名 |
| `name_pinyin` | varchar(256) | 无 | no | 属性名拼音 |
| `required` | tinyint(4) | 无 | no | 是否必填 |
| `built_in` | tinyint(1) | 无 | no | 系统内置属性 |
| `create_time` | bigint(20) | 无 | no | 创建时间（秒） |
| `max_length` | int(11) | 无 | yes | 最大长度 |
| `options` | text | 无 | yes | 选项配置 |
| `statuses` | text | 无 | yes | 属性配置 |

## `project_field_value` - 项目属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `project_uuid` | varchar(16) | 无 | no | 项目UUID |
| `alias` | varchar(32) | 无 | no | 别名 |
| `type` | int(11) | 无 | no | 值类型 |
| `value` | text | 无 | yes | 属性值 |
| `number_value` | bigint(20) | 无 | yes | 数字属性值 |

## `project_filter`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `project_uuid` | varchar(16) | 无 | yes |  |
| `owner` | varchar(8) | 无 | yes |  |
| `name` | varchar(128) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `query` | text | 无 | no |  |
| `sort` | text | 无 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `include_subtasks` | tinyint(4) | 0 | no | 是否包含子任务 |

## `project_issue_type`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 无 | no |  |
| `issue_type_uuid` | varchar(8) | 空串 | no |  |
| `position` | int(11) | 0 | no | 排序 |

## `project_member`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 无 | no |  |
| `user_domain_param` | varchar(8) | 空串 | no | 成员域id |
| `user_domain_type` | tinyint(4) | 0 | no | 成员域类型 |

## `project_notice_config` - project的任务通知配置关联表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 无 | no | project(项目) uuid |
| `issue_type_uuid` | varchar(8) | 无 | no | 任务类型 uuid |
| `notice_config_uuid` | varchar(8) | 无 | no | 任务通知配置的uuid |
| `config_type` | enum('create_task','update_task_assign',<br>'update_task_status','update_task_priori<br>ty','update_task_title','update_task_des<br>cription','set_task_deadline','update_ta<br>sk_sprint','update_task_message','update<br>_task_watcher','update_task_related_task<br>','upload_task_file','update_task_access<br>_manhour','update_task_other_property','<br>delete_task','related_test_case','relate<br>d_plan_case','task_related_testcase_plan<br>','update_issue_type','update_std_to_sub<br>_issue_type','update_sub_to_sub_issue_ty<br>pe','update_sub_issue_type_parent','upda<br>te_sub_to_std_issue_type','update_task_r<br>emaining_manhour','update_task_record_ma<br>nhour') | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1:正常 2:无效 |

## `project_panel`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `title` | varchar(32) | 无 | no |  |
| `url` | varchar(128) | 无 | no |  |

## `project_pin` - 项目Pin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 无 | no | 所属项目 uuid |
| `updated_time` | bigint(20) | 无 | no | 更新时间 |

## `project_pipeline` - 项目和pipeline绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 无 | no |  |
| `pipeline_uuid` | varchar(8) | 空串 | no |  |
| `create_time` | bigint(20) | 无 | no | 绑定时间（秒） |

## `project_plugin` - 项目插件状态表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 无 | no | 项目id |
| `plugin_uuid` | varchar(8) | 无 | no | 插件id |
| `config` | text | 无 | no |  |
| `status` | tinyint(1) | 0 | no | 是否开启工时 1：开启，2：关闭 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |

## `project_report`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `project_uuid` | varchar(16) | 空串 | no | 项目uuid |
| `name` | varchar(128) | 空串 | no | 常用项目报表名称 |
| `owner` | varchar(8) | 无 | no | 创建者uuid |
| `config` | text | 无 | no | 项目报表配置 |
| `create_time` | bigint(20) | 0 | no | 创建时间，秒 |
| `update_time` | bigint(20) | 0 | no | 更新时间，秒 |
| `status` | tinyint(4) | 0 | no | 状态(1:正常, 2:删除) |

## `project_sprint_field` - 项目下的迭代属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `type` | int(11) | 0 | no | 类型 |
| `name` | varchar(32) | 空串 | no | 属性名 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `status` | tinyint(4) | 0 | no | 1:正常;2:删除 |
| `position` | int(11) | 0 | no | 排序序号 |

## `project_sprint_field_option` - 迭代属性选项，用于单选或者多选类型的属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `field_uuid` | varchar(8) | 空串 | no | 属性UUID |
| `value` | varchar(64) | 空串 | no | 属性值 |
| `default_selected` | tinyint(4) | 0 | no | 是否默认选中 |
| `position` | int(11) | 0 | no | 排序序号 |

## `project_sprint_status` - 项目下的迭代阶段

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `name` | varchar(32) | 空串 | no | 阶段名 |
| `category` | int(11) | 0 | no | 状态分类 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `position` | int(11) | 0 | no | 排序序号 |
| `built_in` | tinyint(4) | 0 | no | 是否内置状态 |
| `status` | tinyint(4) | 0 | no | 1:正常;2:删除 |

## `project_task_sort` - 项目任务排序表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `attribute_uuid` | varchar(16) | 空串 | no | 任务UUID/分段UUID |
| `attribute_type` | int(11) | 0 | no | 类型 |
| `position` | bigint(20) | 0 | no | 任务排序权重 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 最近修改时间 |

## `report_project_attribute` - 项目属性分析

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `day_time` | int(11) | 无 | no | 某天时间戳 |
| `project_uuid` | varchar(16) | 无 | no | 项目uuid |
| `attribute_uuid` | varchar(8) | 无 | no | 属性名 |
| `attribute_value` | varchar(32) | 无 | no | 属性值 |
| `count` | int(11) | 0 | no | 个数(累计) |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |

## `report_project_progress` - 项目总体进度表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `day_time` | int(11) | 无 | no | 某天时间戳 |
| `project_uuid` | varchar(16) | 无 | no | 项目uuid |
| `incomplete` | int(11) | 0 | no | 未完成任务数(累计) |
| `complete` | int(11) | 0 | no | 已完成任务数(累计) |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |

## `report_project_target` - 项目目标追踪

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `day_time` | int(11) | 无 | no | 某天时间戳 |
| `project_uuid` | varchar(16) | 无 | no | 项目uuid |
| `scores` | bigint(20) | 0 | no | 完成目标值 |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |

## `request_code`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(64) | 无 | no |  |
| `email` | varchar(64) | 无 | no |  |
| `phone` | varchar(32) | 无 | no |  |
| `password` | varchar(60) | 无 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `ip` | varchar(15) | 无 | yes |  |
| `status` | tinyint(4) | -1 | no |  |

## `reset_code`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(32) | 无 | no |  |
| `email` | varchar(64) | 无 | no |  |
| `phone` | varchar(64) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `ip` | varchar(15) | 无 | yes |  |
| `status` | tinyint(4) | 无 | no |  |

## `reset_password_code` - 重置团队成员密码邮箱验证码

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `code` | varchar(32) | 无 | no | 验证码 |
| `user_uuid` | varchar(8) | 无 | no | 用户ID |
| `member_uuid` | varchar(8) | 无 | no | 团队成员的用户ID |
| `email` | varchar(255) | 无 | no | 邮箱地址 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `ip` | varchar(15) | 无 | no | 请求IP地址 |
| `status` | tinyint(4) | 无 | no | 1:正常,2:无效 |

## `resource`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `reference_type` | tinyint(4) | 无 | no |  |
| `reference_id` | varchar(16) | 无 | yes |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `project_uuid` | varchar(16) | 无 | no |  |
| `owner_uuid` | varchar(8) | 无 | yes |  |
| `type` | tinyint(4) | -1 | no |  |
| `source` | tinyint(4) | 0 | no |  |
| `ext_id` | varchar(28) | 无 | yes |  |
| `name` | varchar(255) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `modify_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | -1 | no |  |
| `description` | varchar(64) | 无 | no |  |

## `role`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `name` | varchar(32) | 空串 | no | 角色名称 |
| `name_pinyin` | varchar(256) | 空串 | no | 角色名称拼音 |
| `built_in` | tinyint(4) | 0 | no | 是否系统内置角色 |
| `is_project_member` | tinyint(4) | 0 | no | 是否项目成员角色 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `status` | tinyint(4) | 无 | yes | 1:正常;null:删除 |

## `role_config` - 角色配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 无 | no |  |
| `role_uuid` | varchar(8) | 空串 | no |  |
| `create_time` | bigint(20) | 0 | no |  |

## `role_member` - 角色成员

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no |  |
| `role_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 空串 | no |  |

## `section` - 分段表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | 项目UUID |
| `source` | varchar(8) | 空串 | no | 来源 |
| `source_uuid` | varchar(16) | 空串 | no | 来源对应UUID |
| `mark` | varchar(32) | 空串 | no | 标记 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 最近修改时间 |

## `session_token`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `user_id` | varchar(8) | 无 | yes |  |
| `session_token` | varchar(64) | 无 | yes |  |
| `session_token_status` | tinyint(4) | -1 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `ip` | varchar(15) | 无 | yes |  |
| `type` | int(11) | 0 | no |  |

## `setting`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `name` | varchar(255) | 无 | no |  |
| `value` | varchar(255) | 无 | yes |  |

## `sprint`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 无 | yes |  |
| `title` | varchar(32) | 无 | yes |  |
| `description` | text | 无 | yes |  |
| `start_time` | bigint(20) | 0 | yes |  |
| `end_time` | bigint(20) | 0 | yes |  |
| `status` | tinyint(4) | -1 | no |  |
| `status_uuid` | varchar(8) | 无 | no | 状态UUID |
| `create_time` | bigint(20) | 0 | yes |  |
| `title_pinyin` | varchar(256) | 无 | no |  |
| `assign` | varchar(8) | 空串 | no | 迭代负责人 uuid |
| `goal` | text | 无 | no | 迭代目标 |

## `sprint_field_value` - 迭代属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `sprint_uuid` | varchar(8) | 空串 | no | 迭代UUID |
| `field_uuid` | varchar(8) | 无 | no | 属性UUID |
| `type` | int(11) | 0 | no | 值类型 |
| `value` | text | 无 | yes | 属性值 |
| `number_value` | bigint(20) | 无 | yes | 数字属性值 |

## `sprint_pin`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `sprint_uuid` | varchar(8) | 无 | no |  |
| `updated_time` | bigint(20) | 无 | no | 更新时间 |

## `sprint_pipeline` - 迭代和pipeline分支绑定关系

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `sprint_uuid` | varchar(8) | 空串 | no | 迭代uuid |
| `pipeline_uuid` | varchar(8) | 空串 | no | 流水线uuid |
| `repository` | varchar(255) | 空串 | no | 代码仓库名称 |
| `branch` | varchar(255) | 空串 | no | 分支名称 |
| `create_time` | bigint(20) | 无 | no | 绑定时间（秒） |

## `sprint_status_value` - 迭代的阶段信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `sprint_uuid` | varchar(8) | 空串 | no | 迭代UUID |
| `status_uuid` | varchar(8) | 空串 | no | 阶段UUID |
| `plan_start_time` | bigint(20) | 无 | yes | 计划开始时间 |
| `plan_end_time` | bigint(20) | 无 | yes | 计划结束时间 |
| `actual_start_time` | bigint(20) | 无 | yes | 实际开始时间 |
| `actual_end_time` | bigint(20) | 无 | yes | 实际结束时间 |
| `desc_rich` | text | 无 | no | 阶段描述(富文本) |

## `sprint_timeline`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `sprint_uuid` | varchar(8) | 无 | no |  |
| `type` | varchar(16) | 无 | no |  |
| `time` | bigint(20) | 无 | yes |  |

## `sso_verify_info` - sso验证

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `auth_code` | varchar(512) | 空串 | no |  |
| `sync_id` | varchar(255) | 空串 | no | 同步账号的id |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `create_time` | bigint(20) | 0 | no |  |

## `support_request` - 申请支持（升级到专业版）请求

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `email` | varchar(255) | 空串 | no | 邮箱 |
| `phone` | varchar(32) | 空串 | no | 手机号 |
| `source` | int(11) | 0 | no | 来源(1:创建团队 2:升级到专业版) |
| `urgent` | tinyint(4) | 0 | no | 是否紧急 |
| `team_uuid` | varchar(8) | 无 | no | 请求者团队uuid |
| `user_uuid` | varchar(8) | 无 | no | 请求者uuid |
| `name` | varchar(128) | 空串 | no | 用户名称 |
| `team_name` | varchar(128) | 空串 | no | 团队名 |
| `team_scale` | varchar(128) | 空串 | no | 团队规模 |
| `industry` | varchar(128) | 空串 | no | 所处行业 |
| `create_time` | datetime | CURRENT_TIMESTAMP | no | 创建时间 |
| `status` | int(11) | 0 | no | 1:正常 2:删除 |
| `referrer` | text | 无 | no | 渠道信息 |
| `keyword` | text | 无 | no | 搜索关键词 |
| `ip_address` | varchar(15) | 空串 | no | 用户IP |
| `ip_location` | varchar(255) | 空串 | no | IP地理位置 |
| `channel` | varchar(255) | 空串 | no | 渠道来源 |
| `issue` | varchar(48) | 空串 | no | 希望解决的问题 |
| `team_role` | varchar(24) | 空串 | no | 团队角色 |
| `demo_name` | varchar(64) | 空串 | no | demo?? |
| `email` | varchar(255) | 空串 | no | 邮箱 |
| `phone` | varchar(32) | 空串 | no | 手机号 |
| `source` | int(11) | 0 | no | 来源(1:创建团队 2:升级到专业版) |
| `urgent` | tinyint(4) | 0 | no | 是否紧急 |
| `team_uuid` | varchar(8) | 无 | no | 请求者团队uuid |
| `user_uuid` | varchar(8) | 无 | no | 请求者uuid |
| `name` | varchar(128) | 空串 | no | 用户名称 |
| `team_name` | varchar(128) | 空串 | no | 团队名 |
| `team_scale` | varchar(128) | 空串 | no | 团队规模 |
| `industry` | varchar(128) | 空串 | no | 所处行业 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `status` | int(11) | 0 | no | 1:正常 2:删除 |
| `referrer` | text | 无 | no | 渠道信息 |
| `keyword` | text | 无 | no | 搜索关键词 |
| `ip_address` | varchar(15) | 空串 | no | 用户IP |
| `ip_location` | varchar(255) | 空串 | no | IP地理位置 |
| `channel` | varchar(255) | 空串 | no | 渠道来源 |
| `demo_name` | varchar(64) | 空串 | no | demo名称 |
| `issue` | varchar(48) | 空串 | no | 希望解决的问题 |
| `team_role` | varchar(24) | 空串 | no | 团队角色 |

## `sync_department_config` - 选中第三方部门配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `org_uuid` | varchar(8) | 无 | no | 组织uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `sync_id` | bigint(20) | 无 | yes | 同步部门的id |

## `sync_user` - 同步用户表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `user_uuid` | varchar(8) | 空串 | no | user id |
| `sync_id` | varchar(255) | 空串 | no | 同步账号的id |
| `sync_type` | tinyint(4) | 0 | no | 同步账号类型 |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `status` | tinyint(4) | -1 | no | 状态:1:正常;2:删除 |

## `sync_user_config` - 选中第三方用户配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `org_uuid` | varchar(8) | 无 | no | 组织uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `sync_id` | varchar(255) | 空串 | no | 同步账号的id |

## `tag` - 标签库

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `name` | varchar(32) | 空串 | no | 标签名 |

## `task`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(16) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `priority` | varchar(8) | 无 | yes |  |
| `owner` | varchar(8) | 无 | yes |  |
| `assign` | varchar(8) | 无 | yes |  |
| `tags` | text | 无 | yes |  |
| `sprint_uuid` | varchar(8) | 无 | yes |  |
| `project_uuid` | varchar(16) | 无 | yes |  |
| `issue_type_uuid` | varchar(8) | 空串 | no | 任务类型UUID |
| `status_uuid` | varchar(8) | 无 | no | 任务状态UUID |
| `deadline` | bigint(20) | 无 | yes |  |
| `status` | tinyint(4) | 0 | no |  |
| `summary` | text | 无 | no |  |
| `desc` | longtext | 无 | yes | 任务描述 |
| `desc_rich` | longtext | 无 | yes | 任务描述（富文本） |
| `server_update_stamp` | bigint(20) | 0 | no |  |
| `complete_time` | bigint(20) | 0 | no |  |
| `open_time` | bigint(20) | 0 | no |  |
| `score` | bigint(20) | 无 | yes |  |
| `parent_uuid` | varchar(16) | 无 | no |  |
| `position` | bigint(20) | 0 | no | 子任务排序 |
| `number` | int(11) | 无 | no | 编号 |
| `assess_manhour` | bigint(20) | 无 | yes | 预估工时 |
| `path` | varchar(169) | 无 | no | 支持10层子任务（n*17-1） |
| `sub_issue_type_uuid` | varchar(8) | 空串 | no | sub_issue_type_uuid为空则不是子任务； 不为<br>空是子任务，此时issue_type_uuid是父任务的<br>标准任务类型 |
| `related_count` | int(11) | 0 | no | 关联任务的数量 |
| `remaining_manhour` | bigint(20) | 无 | yes | 剩余工时 |

## `task_attendee`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 无 | yes |  |
| `user_uuid` | varchar(8) | 无 | yes |  |

## `task_attribute` - 任务属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 空串 | no | 任务UUID |
| `template_uuid` | varchar(8) | 空串 | no | 模板UUID |
| `template_attribute_uuid` | varchar(8) | 空串 | no | 属性uuid |
| `value` | text | 无 | yes | 属性值 |
| `type` | int(11) | 0 | no | 值类型 |
| `status` | int(11) | 0 | no | 状态(1:正常,2:删除) |

## `task_case`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 无 | no |  |
| `case_uuid` | varchar(8) | 无 | no |  |

## `task_commit`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 无 | no |  |
| `commit_uuid` | varchar(8) | 无 | no |  |

## `task_gantt_chart`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | gantt chart uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `component_uuid` | varchar(8) | 无 | no | 组件 uuid |
| `name` | varchar(32) | 无 | no | gantt chart名称 |
| `name_pinyin` | varchar(256) | 空串 | no | gantt chart名称拼音 |
| `owner` | varchar(8) | 无 | no | 创建者 uuid |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `status` | tinyint(4) | 无 | yes | 1:正常 2:删除 |

## `task_gantt_chart_personal_config`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `gantt_chart_uuid` | varchar(8) | 无 | no | gantt chart uuid |
| `user_uuid` | varchar(8) | 无 | no | 用户 uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `expand` | tinyint(4) | 1 | no | 0:收起  1: 展开 |
| `zooming` | tinyint(4) | 0 | no | 缩放 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `status` | tinyint(4) | 1 | no | 1: 正常 2: 删除 |

## `task_gantt_data`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | 甘特图数据 uuid |
| `team_uuid` | varchar(8) | 无 | no | 团队 uuid |
| `gantt_chart_uuid` | varchar(8) | 无 | no | gantt chart uuid |
| `type` | tinyint(4) | 0 | yes | 1:甘特图数据类型, 1:任务 2:分组 3:里程碑<br> |
| `name` | varchar(128) | 无 | no | 名称 |
| `name_pinyin` | varchar(512) | 空串 | no | 名称拼音 |
| `owner` | varchar(8) | 无 | no | 创建者 uuid |
| `start_time` | bigint(20) | 无 | no | 开始时间 |
| `end_time` | bigint(20) | 无 | no | 结束时间 |
| `progress` | int(11) | 0 | no | 进度 |
| `parent` | varchar(8) | 无 | no | 父节点 uuid |
| `path` | varchar(169) | 无 | no | 层级关系 |
| `position` | bigint(20) | 0 | no | 显示位置 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `status` | tinyint(4) | 无 | yes | 1:正常 2:删除 |
| `assign` | varchar(8) | 无 | yes | 负责人uuid |
| `related_type` | tinyint(4) | 0 | no | 关联数据类型 |
| `related_project_uuid` | varchar(16) | 无 | no | 关联的项目uuid |
| `related_task_uuid` | varchar(16) | 无 | no | 关联的任务uuid |

## `task_link_type`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `name` | varchar(32) | 空串 | no | 任务关联类型名称 |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `link_out_desc` | varchar(32) | 空串 | no | 链出描述 |
| `link_out_desc_pinyin` | varchar(256) | 空串 | no | 链出描述拼音 |
| `link_in_desc` | varchar(32) | 空串 | no | 链入描述 |
| `link_in_desc_pinyin` | varchar(256) | 空串 | no | 链入描述拼音 |
| `built_in` | tinyint(1) | 0 | no | 系统内置默认任务关联类型,1:是,0否 |
| `create_time` | bigint(20) | 0 | yes |  |
| `status` | tinyint(4) | 1 | no | 1:正常 2:删除 |

## `task_plan_case`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 无 | no |  |
| `plan_uuid` | varchar(8) | 无 | no |  |
| `case_uuid` | varchar(8) | 无 | no |  |

## `task_preference`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `user_uuid` | varchar(8) | 无 | no |  |
| `task_uuid` | varchar(16) | 无 | no |  |
| `order` | int(11) | 0 | no |  |
| `alarm_time` | bigint(20) | 0 | no |  |
| `server_update_stamp` | bigint(20) | 0 | no |  |

## `task_related`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no |  |
| `source_task_uuid` | varchar(16) | 无 | no |  |
| `target_task_uuid` | varchar(16) | 无 | no |  |
| `task_link_type_uuid` | varchar(8) | 空串 | no | 任务关联类型UUID |
| `create_time` | bigint(20) | 0 | yes |  |

## `task_richtext_message` - 任务富文本属性动态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `message_uuid` | varchar(8) | 无 | no | message_uuid |
| `task_uuid` | varchar(16) | 无 | no | 任务UUID |
| `field_uuid` | varchar(8) | 无 | no | field_uuid |
| `old_value_time` | bigint(20) | 无 | no | 产生上一版本任务富文本属性的时间 |
| `new_value_time` | bigint(20) | 无 | no | 产生当前版本任务富文本属性的时间 |
| `old_value` | longtext | 无 | yes | 更改前的任务富文本属性值 |
| `new_value` | longtext | 无 | yes | 更改后的任务富文本属性值 |

## `task_sort` - task_sort

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 无 | no | uuid |
| `project_uuid` | varchar(16) | 无 | no | 所属项目 uuid |
| `board_uuid` | varchar(8) | 无 | no | 所属board uuid |
| `position` | bigint(20) | 无 | no | 顺序(从小到大) |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 更新时间 |

## `task_status` - 任务状态

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `name` | varchar(32) | 空串 | no | 任务状态名 |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `category` | int(11) | 0 | no | 任务状态分类 |
| `built_in` | tinyint(4) | 0 | no | 是否系统内置状态 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `status` | tinyint(4) | 无 | yes | 1:正常;null:删除 |

## `task_status_config` - 任务状态在某项目的某任务类型中的配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `project_uuid` | varchar(16) | 空串 | no | 项目uuid |
| `issue_type_uuid` | varchar(8) | 空串 | no | 任务类型uuid |
| `status_uuid` | varchar(8) | 空串 | no | 任务状态uuid |
| `is_default` | tinyint(4) | 0 | no | 是否默认状态 |
| `position` | int(11) | 0 | no | 状态位置 |

## `task_testcase_plan`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `task_uuid` | varchar(16) | 无 | no |  |
| `plan_uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |

## `task_watcher`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `task_uuid` | varchar(16) | 空串 | no |  |
| `user_uuid` | varchar(8) | 空串 | no |  |
| `create_time` | bigint(20) | 0 | no |  |

## `team`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | -1 | no |  |
| `name` | varchar(16) | 无 | no |  |
| `owner` | varchar(8) | 无 | yes |  |
| `logo` | varchar(255) | 无 | no |  |
| `domain` | varchar(255) | 无 | yes |  |
| `cover_url` | varchar(128) | 空串 | no |  |
| `expire_time` | bigint(20) | 无 | no |  |
| `scale` | int(6) | 无 | no |  |
| `csm` | varchar(128) | 空串 | yes |  |
| `type` | tinyint(4) | -1 | no |  |
| `sprint_alias` | varchar(32) | 空串 | no |  |
| `org_uuid` | varchar(8) | 无 | no |  |
| `workdays` | varchar(8) | 无 | no |  |
| `workhours` | bigint(20) | 0 | yes |  |
| `uuid` | varchar(8) | 无 | no |  |
| `is_create_default_space` | tinyint(4) | 无 | no |  |

## `team_member`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `modify_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | -1 | no |  |
| `user_avatar` | varchar(255) | 空串 | no | 头像 |
| `user_title` | varchar(16) | 空串 | no | 职位 |
| `user_name` | varchar(16) | 空串 | no | 姓名 |
| `user_name_pinyin` | varchar(128) | 无 | no |  |
| `inviter_uuid` | varchar(8) | 无 | no | 邀请人uuid |
| `user_team_constraint` | varchar(32) | 无 | yes | 用户只有一个有效团队 |
| `org_uuid` | varchar(8) | 无 | no |  |

## `team_notice_config` - team的全局任务通知配置关联表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | team（团队） uuid |
| `issue_type_uuid` | varchar(8) | 无 | no | 任务类型 uuid |
| `notice_config_uuid` | varchar(8) | 无 | no | 任务通知配置的uuid |
| `config_type` | enum('create_task','update_task_assign',<br>'update_task_status','update_task_priori<br>ty','update_task_title','update_task_des<br>cription','set_task_deadline','update_ta<br>sk_sprint','update_task_message','update<br>_task_watcher','update_task_related_task<br>','upload_task_file','update_task_access<br>_manhour','update_task_other_property','<br>delete_task','related_test_case','relate<br>d_plan_case','task_related_testcase_plan<br>','update_issue_type','update_std_to_sub<br>_issue_type','update_sub_to_sub_issue_ty<br>pe','update_sub_issue_type_parent','upda<br>te_sub_to_std_issue_type','update_task_r<br>emaining_manhour','update_task_record_ma<br>nhour') | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1:正常 2:无效 |

## `team_plugin`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | uuid |
| `plugin_uuid` | varchar(8) | 无 | no | uuid |
| `is_open` | tinyint(1) | 无 | no |  |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `config` | text | 无 | no |  |

## `team_report`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `name` | varchar(128) | 无 | no | 常用团队报表名称 |
| `owner` | varchar(8) | 无 | no | 创建者uuid |
| `config` | text | 无 | no | 项目报表配置 |
| `create_time` | bigint(20) | 0 | no | 创建时间，秒 |
| `update_time` | bigint(20) | 0 | no | 更新时间，秒 |
| `status` | tinyint(4) | 0 | no | 状态(1:正常, 2:删除) |

## `team_scm`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队uuid |
| `status` | tinyint(4) | 1 | no | 1：正常 2：删除 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `scm_secret_key` | varchar(22) | 无 | no | scm github/gitlab 密钥 |
| `scm_last_success_time` | bigint(20) | 无 | no | 最后的更新时间 |

## `template` - 模板基础表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `owner` | varchar(16) | 空串 | no | 拥有者 |
| `category_uuid` | varchar(8) | 空串 | no | 分类uuid |
| `name` | varchar(32) | 空串 | no | 模板名 |
| `type` | tinyint(4) | 0 | no | 类型 |
| `descs` | text | 无 | yes | 描述 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |
| `is_open_manhour` | tinyint(1) | 0 | yes | 工时开关 |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `space_uuid` | varchar(8) | 无 | no | 所属 space |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `title` | varchar(64) | 无 | no | 模版标题 |
| `title_pinyin` | varchar(256) | 空串 | no | 模版标题-拼音 |
| `status` | tinyint(4) | 无 | no | 状态 (1: 正常, 2: 删除) |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `update_time` | bigint(20) | 无 | no | 最近更新时间 |

## `template_attribute` - 模板任务属性表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `template_uuid` | varchar(8) | 空串 | no | 模板UUID |
| `name` | varchar(32) | 空串 | no | 属性名 |
| `type` | int(11) | 0 | no | 类型 |
| `text` | text | 无 | no | 文本属性值 |
| `number` | bigint(20) | 无 | yes | 数字属性值 |
| `is_default_show` | tinyint(1) | 1 | no | 任务列表是否显示 |
| `position` | int(11) | 0 | no | 排序序号 |

## `template_attribute_value` - 属性值表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `attribute_uuid` | varchar(8) | 空串 | no | 属性UUID |
| `value` | varchar(64) | 空串 | no | 属性值 |
| `font_color` | varchar(8) | 空串 | no | 字体颜色 |
| `bg_color` | varchar(8) | 空串 | no | 字体颜色 |
| `is_selected` | tinyint(1) | 0 | no | 单选框中值是否被选择 |
| `position` | int(11) | 0 | no | 排序序号 |

## `template_category` - 模板分类信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `name` | varchar(128) | 空串 | no | 分类名 |
| `pic_url` | varchar(1024) | 空串 | no | 分类图标地址 |

## `template_tag_relational` - 模板标签映射

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `template_uuid` | varchar(8) | 空串 | no | UUID |
| `tag_uuid` | varchar(16) | 空串 | no | 拥有者 |

## `template_target` - 模板任务目标表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `template_uuid` | varchar(8) | 空串 | no | 模板UUID |
| `name` | varchar(32) | 空串 | no | 量化指标 |
| `score` | bigint(20) | 无 | yes | 目标值 |
| `unit` | varchar(32) | 空串 | no | 单位 |

## `testcase_case`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `library_uuid` | varchar(8) | 无 | no |  |
| `path` | varchar(255) | 无 | no |  |
| `module_uuid` | varchar(8) | 无 | no |  |
| `name` | varchar(1024) | 无 | no |  |
| `priority` | varchar(8) | 无 | no |  |
| `type` | varchar(16) | 无 | no |  |
| `assign` | varchar(8) | 无 | no |  |
| `desc` | text | 无 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `number` | int(11) | 无 | no | 编号 |
| `status` | tinyint(4) | 0 | no |  |
| `condition` | text | 无 | yes |  |
| `name_pinyin` | varchar(10000) | 无 | no | 用例名字拼音 |

## `testcase_case_step`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `library_uuid` | varchar(8) | 无 | no |  |
| `case_uuid` | varchar(8) | 无 | no |  |
| `desc` | text | 无 | yes |  |
| `result` | text | 无 | yes |  |
| `index` | int(5) | 无 | no |  |

## `testcase_default_member`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 无 | no |  |
| `user_domain_type` | tinyint(4) | 0 | no | 成员域类型 |
| `user_domain_param` | varchar(8) | 空串 | no | 成员域参数 |

## `testcase_field` - TestCase属性

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | UUID |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `context` | varchar(32) | 空串 | no | 上下文 |
| `alias` | varchar(16) | 空串 | no | 别名 |
| `status` | tinyint(4) | 无 | yes | 状态（1:正常 NULL:删除） |
| `type` | int(11) | 无 | no | 类型 |
| `name` | varchar(32) | 无 | no | 属性名 |
| `name_pinyin` | varchar(256) | 无 | no | 属性名拼音 |
| `required` | tinyint(4) | 无 | no | 是否必填 |
| `built_in` | tinyint(1) | 无 | no | 系统内置属性 |
| `create_time` | bigint(20) | 无 | no | 创建时间（秒） |
| `max_length` | int(11) | 无 | yes | 最大长度 |
| `options` | text | 无 | yes | 选项配置 |
| `statuses` | text | 无 | yes | 属性配置 |
| `default_value` | text | 无 | yes | 属性默认值 |

## `testcase_field_config` - TestCase属性配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `name` | varchar(32) | 无 | no | 属性配置模版 |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `create_time` | bigint(20) | 无 | no | 创建时间（秒） |
| `is_default` | tinyint(1) | 无 | no | 是否为默认属性配置 |

## `testcase_field_value` - TestCase属性值

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队UUID |
| `object_uuid` | varchar(8) | 无 | no | testcase的所有实体，包括用例，用例库，测试计<br>划等等 |
| `item_type` | varchar(32) | 无 | no | itemType |
| `alias` | varchar(32) | 无 | no | 别名 |
| `type` | int(11) | 无 | no | 值类型 |
| `value` | text | 无 | yes | 属性值 |
| `number_value` | bigint(20) | 无 | yes | 数字属性值 |

## `testcase_library`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `name` | varchar(32) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `name_pinyin` | varchar(256) | 无 | no | 用例库名字拼音 |
| `field_config_uuid` | varchar(8) | 无 | no | 属性配置 |

## `testcase_library_pin` - 用例库Pin

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `library_uuid` | varchar(16) | 无 | no | 所属library uuid |
| `updated_time` | bigint(20) | 无 | no | 更新时间 |

## `testcase_module`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `parent_uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `library_uuid` | varchar(8) | 无 | no |  |
| `path` | varchar(255) | 无 | no |  |
| `name` | varchar(32) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `is_default` | tinyint(1) | 无 | no | 是否为无所属模块 |
| `position` | bigint(20) | 无 | no |  |
| `name_pinyin` | varchar(256) | 无 | no | 模块名字拼音 |

## `testcase_plan`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `owner` | varchar(8) | 无 | no |  |
| `name` | varchar(32) | 无 | no |  |
| `stage` | varchar(16) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `plan_status` | tinyint(4) | 0 | no |  |
| `related_project` | varchar(16) | 无 | yes |  |
| `related_sprint` | varchar(8) | 无 | yes |  |
| `related_issue_type` | varchar(8) | 无 | yes |  |
| `name_pinyin` | varchar(256) | 无 | no | 测试计划名字拼音 |

## `testcase_plan_assigns`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `plan_uuid` | varchar(8) | 空串 | no |  |
| `user_uuid` | varchar(8) | 空串 | no |  |
| `create_time` | bigint(20) | 无 | no | 创建时间 |

## `testcase_plan_case`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `plan_uuid` | varchar(8) | 无 | no |  |
| `case_uuid` | varchar(8) | 无 | no |  |
| `executor` | varchar(8) | 无 | no |  |
| `result` | varchar(8) | 无 | no |  |
| `note` | text | 无 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `warn_step` | varchar(8) | 无 | yes |  |
| `team_uuid` | varchar(8) | 无 | no |  |

## `testcase_plan_case_step`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `plan_uuid` | varchar(8) | 无 | no |  |
| `case_uuid` | varchar(8) | 无 | no |  |
| `step_uuid` | varchar(8) | 无 | no |  |
| `step_result` | varchar(8) | 无 | no |  |
| `actual_result` | text | 无 | yes |  |
| `team_uuid` | varchar(8) | 无 | no |  |

## `testcase_project_config`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `project_uuid` | varchar(16) | 无 | no |  |
| `issue_type_uuid` | varchar(8) | 无 | no |  |

## `testcase_report`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `plan_uuid` | varchar(8) | 无 | no |  |
| `start_time` | bigint(20) | 0 | no |  |
| `end_time` | bigint(20) | 0 | no |  |
| `summary` | text | 无 | yes |  |
| `title` | varchar(32) | 空串 | yes |  |
| `team_uuid` | varchar(8) | 无 | no |  |

## `testcase_report_template`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `content` | text | 无 | yes |  |
| `name` | varchar(32) | 无 | no |  |
| `status` | tinyint(4) | 0 | no |  |

## `third_party_config`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `wechat_suite_ticket` | varchar(512) | 空串 | yes | 企业微信suite_ticket |
| `dingding_suite_ticket` | varchar(512) | 空串 | yes | 钉钉suite_ticket |

## `third_party_organization`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `corp_uuid` | varchar(64) | 无 | no | 第三方登录企业id |
| `org_uuid` | varchar(8) | 无 | no | 对应组织uuid |
| `corp_name` | varchar(255) | 无 | no | 第三方企业名称 |
| `org_type` | tinyint(4) | 1 | no | 来源平台 |
| `auth_status` | tinyint(4) | 1 | no | 授权状态 |
| `permanent_code` | varchar(512) | 无 | no | 第三方企业微信永久授权码 |
| `visible_scale` | int(6) | 无 | no | 可见范围内人数 |
| `agent_id` | int(20) | 无 | no | 授权应用agentid |

## `transition` - 工作流步骤

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `project_uuid` | varchar(16) | 空串 | no | 项目UUID |
| `issue_type_uuid` | varchar(8) | 空串 | no | 任务类型UUID |
| `start_status` | varchar(8) | 无 | no | 起始状态UUID |
| `end_status` | varchar(8) | 无 | no | 结束状态UUID |
| `name` | varchar(64) | 空串 | no | 名称 |
| `trigger` | text | 无 | no | 触发器 |
| `condition` | text | 无 | no | 验证条件 |
| `post_function` | text | 无 | no | 后置动作 |
| `fields` | text | 无 | no | 步骤属性 |
| `position` | int(11) | 0 | no | 步骤位置 |

## `update_stamp` - 获取更新状态记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队id |
| `type` | tinyint(4) | 0 | no | 更新类型 |
| `user_uuid` | varchar(8) | 空串 | no | 用户id |
| `server_update_stamp` | bigint(20) | 0 | no | 更新标识 |

## `user`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `Host` | char(60) | 空串 | no |  |
| `User` | char(32) | 空串 | no |  |
| `Select_priv` | enum('N','Y') | N | no |  |
| `Insert_priv` | enum('N','Y') | N | no |  |
| `Update_priv` | enum('N','Y') | N | no |  |
| `Delete_priv` | enum('N','Y') | N | no |  |
| `Create_priv` | enum('N','Y') | N | no |  |
| `Drop_priv` | enum('N','Y') | N | no |  |
| `Reload_priv` | enum('N','Y') | N | no |  |
| `Shutdown_priv` | enum('N','Y') | N | no |  |
| `Process_priv` | enum('N','Y') | N | no |  |
| `File_priv` | enum('N','Y') | N | no |  |
| `Grant_priv` | enum('N','Y') | N | no |  |
| `References_priv` | enum('N','Y') | N | no |  |
| `Index_priv` | enum('N','Y') | N | no |  |
| `Alter_priv` | enum('N','Y') | N | no |  |
| `Show_db_priv` | enum('N','Y') | N | no |  |
| `Super_priv` | enum('N','Y') | N | no |  |
| `Create_tmp_table_priv` | enum('N','Y') | N | no |  |
| `Lock_tables_priv` | enum('N','Y') | N | no |  |
| `Execute_priv` | enum('N','Y') | N | no |  |
| `Repl_slave_priv` | enum('N','Y') | N | no |  |
| `Repl_client_priv` | enum('N','Y') | N | no |  |
| `Create_view_priv` | enum('N','Y') | N | no |  |
| `Show_view_priv` | enum('N','Y') | N | no |  |
| `Create_routine_priv` | enum('N','Y') | N | no |  |
| `Alter_routine_priv` | enum('N','Y') | N | no |  |
| `Create_user_priv` | enum('N','Y') | N | no |  |
| `Event_priv` | enum('N','Y') | N | no |  |
| `Trigger_priv` | enum('N','Y') | N | no |  |
| `Create_tablespace_priv` | enum('N','Y') | N | no |  |
| `ssl_type` | enum('','ANY','X509','SPECIFIED') | 空串 | no |  |
| `ssl_cipher` | blob | 无 | no |  |
| `x509_issuer` | blob | 无 | no |  |
| `x509_subject` | blob | 无 | no |  |
| `max_questions` | int(11) unsigned | 0 | no |  |
| `max_updates` | int(11) unsigned | 0 | no |  |
| `max_connections` | int(11) unsigned | 0 | no |  |
| `max_user_connections` | int(11) unsigned | 0 | no |  |
| `plugin` | char(64) | mysql_native_password | no |  |
| `authentication_string` | text | 无 | yes |  |
| `password_expired` | enum('N','Y') | N | no |  |
| `password_last_changed` | timestamp | 无 | yes |  |
| `password_lifetime` | smallint(5) unsigned | 无 | yes |  |
| `account_locked` | enum('N','Y') | N | no |  |
| `uuid` | varchar(8) | 无 | no |  |
| `email` | varchar(128) | 无 | yes |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | 0 | no |  |
| `verify_status` | tinyint(4) | -1 | no | 邮箱验证状态 |
| `name` | varchar(16) | 无 | no |  |
| `avatar` | varchar(255) | 无 | yes |  |
| `phone` | varchar(32) | 无 | yes | 手机号 |
| `channel` | varchar(32) | 无 | yes |  |
| `password` | varchar(60) | 无 | yes |  |
| `access_time` | bigint(20) | 0 | yes | 最近访问时间 |
| `sync_id` | varchar(255) | 无 | yes | 同步账号的id |

## `user_access` - 用户访问记录

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队id |
| `user_uuid` | varchar(8) | 无 | no | 用户id |
| `type` | tinyint(4) | 无 | no | 产品类型 |
| `access_time` | bigint(20) | 0 | no | 登录时间 |
| `weekday` | tinyint(4) | 无 | no | 星期几 |

## `user_filter_view` - filter表迁移到本表,filter表将被弃用，后续filter表仅用于兼容旧客户端

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `filter_uuid` | varchar(8) | 无 | no | filter_uuid,是系统内置视图(比如ft-t-001)<br>时与uuid不相等，其他情况与uuid相等 |
| `owner` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `name` | varchar(64) | 空串 | no | 视图名称 |
| `layout` | tinyint(1) | 3 | no | 布局方式(1.表格 2.宽详情 3.窄详情) |
| `query` | text | 无 | yes | 筛选方式 |
| `condition` | text | 无 | no | 筛选条件 |
| `sort` | varchar(256) | 空串 | yes | 排序方式 |
| `group_by` | varchar(256) | 空串 | yes | 分组方式 |
| `create_time` | bigint(20) | 无 | no |  |
| `built_in` | tinyint(1) | 0 | no | 是否默认视图 |
| `include_subtasks` | tinyint(1) | 0 | no | 是否包含子任务 |
| `table_field_settings` | text | 无 | yes | 表格模式的表头 |
| `board_settings` | text | 无 | yes | 看板视图设置 |
| `is_fold_all_groups` | tinyint(1) | 0 | no | 是否折叠分组 |
| `status` | tinyint(1) | 1 | yes | 1正常 2已删除 |
| `display_type` | tinyint(4) | 无 | yes | 1:平铺展示, 2:子树展示，3:折叠展示 |
| `is_show_derive` | tinyint(1) | 0 | no | 是否显示派生 |
| `shared` | tinyint(4) | 0 | no | 是否共享,0: 不共享，1: 共享 |

## `user_filter_view_config`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `filter_uuid` | varchar(8) | 无 | no | 对应user_filter_view的filter_uuid |
| `owner` | varchar(8) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `position` | int(11) | 无 | no | 视图位置 |
| `is_show` | tinyint(1) | 1 | no | 1.显示 0.隐藏 |

## `user_filter_view_member` - 筛选器成员（被分享者）

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `filter_uuid` | varchar(8) | 空串 | no | 筛选器UUID |
| `user_domain_param` | varchar(8) | 空串 | no | 成员域参数 |
| `user_domain_type` | tinyint(4) | 0 | no | 成员域类型 |
| `position` | int(11) | 0 | no | 位置 |

## `user_group`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `name` | varchar(32) | 空串 | no |  |
| `name_pinyin` | varchar(256) | 无 | no |  |
| `desc` | text | 无 | no |  |
| `status` | tinyint(4) | 0 | no |  |

## `user_group_member`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `group_uuid` | varchar(8) | 空串 | no |  |
| `user_uuid` | varchar(8) | 空串 | no |  |

## `user_role` - 用户角色映射更新表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `user_uuid` | varchar(8) | 无 | no | 用户uuid |
| `role_uuid` | varchar(8) | 无 | no | 角色uuid |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | no |  |

## `user_task_sort` - 我的任务排序表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `user_uuid` | varchar(8) | 空串 | no | 用户UUID |
| `attribute_uuid` | varchar(16) | 空串 | no | 任务UUID/分段UUID |
| `attribute_type` | int(11) | 0 | no | 类型 |
| `position` | bigint(20) | 0 | no | 任务排序权重 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 最近修改时间 |

## `user_template_relational` - 用户模板映射

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `user_uuid` | varchar(8) | 空串 | no | 用户UUID |
| `template_uuid` | varchar(16) | 空串 | no | 模板UUID |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |

## `version` - 版本管理

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no |  |
| `owner` | varchar(8) | 空串 | yes | 创建者 |
| `team_uuid` | varchar(8) | 空串 | yes | 所属团队 |
| `title` | varchar(64) | 空串 | no | 版本号 |
| `desc` | text | 无 | yes | 版本描述（富文本） |
| `assign` | varchar(8) | 无 | yes | 负责人 |
| `release_time` | bigint(20) | 无 | yes | 版本发布时间 |
| `create_time` | bigint(20) | 0 | no |  |
| `category` | tinyint(1) | 1 | no | 分类 未开始：1，进行中：2，已完成：3 |
| `status` | tinyint(1) | 1 | no | 状态 正常：1，删除：2 |
| `sys_version` | varchar(5) | 空串 | no |  |
| `mysql_version` | varchar(6) | 空串 | no |  |

## `version_sprint`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `version_uuid` | varchar(8) | 空串 | no |  |
| `sprint_uuid` | varchar(8) | 空串 | no |  |
| `project_uuid` | varchar(16) | 空串 | no |  |
| `create_time` | bigint(20) | 0 | no |  |

## `article` - 用户保存的编辑内容

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `draft_uuid` | varchar(8) | 无 | no | page uuid |
| `page_uuid` | varchar(8) | 无 | no | page uuid |
| `title` | varchar(64) | 无 | no | 标题 |
| `content` | longtext | 无 | no | 内容 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 最近更新时间 |

## `attachment`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `space_uuid` | varchar(8) | 无 | no |  |
| `page_uuid` | varchar(8) | 无 | no |  |
| `draft_uuid` | varchar(8) | 无 | no |  |
| `name` | varchar(255) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `server_update_stamp` | bigint(20) | 无 | no |  |

## `card`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `team_uuid` | varchar(8) | 空串 | no | 团队UUID |
| `dashboard_uuid` | varchar(8) | 空串 | no | 仪表盘UUID |
| `name` | varchar(32) | 空串 | no | 卡片名称 |
| `type` | int(11) | 0 | no | 卡片类型 |
| `layout_x` | int(11) | 0 | no | x位置 |
| `layout_y` | int(11) | 0 | no | y位置 |
| `layout_w` | int(11) | 0 | no | 宽度 |
| `layout_h` | int(11) | 0 | no | 高度 |
| `status` | tinyint(4) | 0 | no | 1:正常;2:删除 |
| `config` | text | 无 | no | 卡片配置 |
| `uuid` | varchar(8) | 无 | no |  |
| `space_uuid` | varchar(8) | 无 | no |  |
| `page_uuid` | varchar(8) | 无 | no |  |
| `draft_uuid` | varchar(8) | 无 | no |  |
| `type` | int(11) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `create_time` | bigint(20) | 无 | no |  |
| `config` | mediumtext | 无 | no | 配置 |

## `card_attachment`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `card_uuid` | varchar(8) | 无 | no |  |
| `attachment_uuid` | varchar(8) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `server_update_stamp` | bigint(20) | 无 | no |  |

## `message`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(32) | 无 | no |  |
| `team_uuid` | varchar(8) | 无 | yes |  |
| `reference_type` | tinyint(4) | -1 | no |  |
| `reference_id` | varchar(16) | 无 | yes |  |
| `from_uuid` | varchar(16) | 无 | yes |  |
| `to_uuid` | varchar(16) | 无 | yes |  |
| `send_time` | bigint(20) | 0 | no |  |
| `message` | text | 无 | no |  |
| `type` | tinyint(4) | 无 | no |  |
| `resource_uuid` | varchar(8) | 无 | yes |  |
| `subject_type` | tinyint(4) | 无 | no |  |
| `subject_id` | varchar(16) | 无 | yes |  |
| `action` | tinyint(4) | 无 | no |  |
| `object_type` | tinyint(4) | 无 | no |  |
| `object_id` | varchar(16) | 无 | yes |  |
| `object_name` | text | 无 | no |  |
| `object_attr` | tinyint(4) | 无 | no |  |
| `old_value` | text | 无 | no |  |
| `new_value` | text | 无 | no |  |
| `ext` | text | 无 | no |  |
| `uuid` | varchar(8) | 无 | no | uuid |
| `user_uuid` | varchar(8) | 无 | no | 发送者 |
| `space_uuid` | varchar(8) | 无 | no | space uuid |
| `page_uuid` | varchar(8) | 无 | no | page uuid |
| `send_time` | bigint(20) | 无 | no | 创建时间 |
| `action` | tinyint(4) | 无 | no | 操作动作 |
| `type` | tinyint(4) | 无 | no | 类型(评论消息,系统消息等) |
| `title` | varchar(64) | 无 | no | 标题 |
| `message` | varchar(512) | 无 | no |  |
| `object_attr` | tinyint(4) | 无 | no |  |
| `ext` | varchar(512) | 无 | no |  |

## `page_comment` - page 评论

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `page_uuid` | varchar(8) | 无 | no | page uuid |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `subject` | varchar(1024) | 无 | no | 主题 |
| `status` | tinyint(4) | 无 | no | 状态 (1: 正常状态 2: 删除) |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 最近更新时间 |

## `page_detail` - page 对应版本信息及详情信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `page_uuid` | varchar(8) | 无 | no | uuid |
| `version` | int(11) | 无 | no | page 版本, 从0开始递增 |
| `draft_uuid` | varchar(8) | 无 | no | uuid |
| `is_revert` | tinyint(1) | 无 | no |  |

## `page_draft` - 所有草稿及发布的 page

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `space_uuid` | varchar(8) | 无 | no | 所属 space |
| `page_uuid` | varchar(8) | 无 | no | page uuid |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `version` | int(11) | 无 | yes | -1: 编辑状态, >0: 已发布, NULL: 草稿/已删<br>除 |
| `from_version` | int(11) | 无 | no | 草稿来源那个发布版本 |
| `status` | tinyint(4) | 无 | no | 1: 草稿状态, 2: 编辑状态 3: 发布状态 4: <br>删除 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 最近更新时间 |
| `old_status` | tinyint(4) | 无 | no |  |

## `page_share_config` - 页面分享配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | 对应页面的共享uuid，随机生成的8位uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `space_uuid` | varchar(8) | 无 | no | 页面组uuid |
| `page_uuid` | varchar(8) | 无 | no | 页面uuid |
| `user_uuid` | varchar(8) | 无 | no | 分享人uuid |
| `share_permission_type` | int(1) | 无 | no | 共享页面类型：0:不共享；1:共享给所有人查看<br>；2:共享给所有人编辑；3:指定成员可以查看或<br>编辑 |
| `is_share_sub_page` | int(1) | 0 | no | 是否共享子页面：默认0：不共享子页面；1：共享<br>子页面 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `update_time` | bigint(20) | 无 | no | 更新时间 |

## `page_tree` - page

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `space_uuid` | varchar(8) | 无 | no | 所属 space |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `title` | varchar(64) | 无 | no | 最新标题 |
| `parent_uuid` | varchar(8) | 无 | no | 父节点 uuid |
| `status` | tinyint(4) | 无 | no | 状态 (1: 正常, 2: 删除) |
| `position` | bigint(20) | 无 | no |  |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 最近更新时间 |
| `old_parent_uuid` | varchar(8) | 空串 | no | 被删除之前的父节点 uuid |
| `old_previous_uuid` | varchar(8) | 空串 | no | 被删除之前的前一个节点 uuid |
| `encrypt_status` | tinyint(4) | 1 | no | 页面加密状态 (1: 不加密, 2: 仅查看，3：加密<br>) |

## `page_watch` - 关注 page 的用户表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `page_uuid` | varchar(8) | 无 | no | page uuid |
| `user_uuid` | varchar(8) | 无 | no | user uuid |

## `space` - space

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `title` | varchar(64) | 无 | no | 标题 |
| `title_pinyin` | varchar(512) | 无 | no |  |
| `description` | varchar(1024) | 无 | no | 描述 |
| `status` | tinyint(4) | 无 | no | 状态 (1: 正常, 2: 删除) |
| `is_archive` | tinyint(1) | 无 | no |  |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 最近更新时间 |
| `is_open_share_page` | int(1) | 0 | no | 是否开启共享页面的开关：默认0：不开启共享页<br>面；1：开启共享页面 |

## `space_browse` - 用户浏览space记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `user_uuid` | varchar(8) | 无 | no | 用户uuid |
| `space_uuid` | varchar(16) | 无 | no | space_uuid |
| `browse_time` | bigint(20) | 0 | no | 浏览时间 |

## `space_pin` - space置顶

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no |  |
| `user_uuid` | varchar(8) | 无 | no |  |
| `space_uuid` | varchar(8) | 无 | no | space_uuid |
| `updated_time` | bigint(20) | 无 | no | 更新时间 |

## `space_update_count` - 记录space修改次数

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `team_uuid` | varchar(8) | 无 | no | 团队uuid |
| `space_uuid` | varchar(16) | 无 | no | space_uuid |
| `day` | tinyint(4) | 无 | no | 当月的第几天 |
| `update_time` | bigint(20) | 0 | no | 统计时间 |
| `count` | bigint(20) | 0 | no | 修改次数 |

## `team`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `create_time` | bigint(20) | 0 | no |  |
| `status` | tinyint(4) | -1 | no |  |
| `name` | varchar(16) | 无 | no |  |
| `owner` | varchar(8) | 无 | yes |  |
| `logo` | varchar(255) | 无 | no |  |
| `domain` | varchar(255) | 无 | yes |  |
| `cover_url` | varchar(128) | 空串 | no |  |
| `expire_time` | bigint(20) | 无 | no |  |
| `scale` | int(6) | 无 | no |  |
| `csm` | varchar(128) | 空串 | yes |  |
| `type` | tinyint(4) | -1 | no |  |
| `sprint_alias` | varchar(32) | 空串 | no |  |
| `org_uuid` | varchar(8) | 无 | no |  |
| `workdays` | varchar(8) | 无 | no |  |
| `workhours` | bigint(20) | 0 | yes |  |
| `uuid` | varchar(8) | 无 | no |  |
| `is_create_default_space` | tinyint(4) | 无 | no |  |

## `template` - template

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 空串 | no | UUID |
| `owner` | varchar(16) | 空串 | no | 拥有者 |
| `category_uuid` | varchar(8) | 空串 | no | 分类uuid |
| `name` | varchar(32) | 空串 | no | 模板名 |
| `type` | tinyint(4) | 0 | no | 类型 |
| `descs` | text | 无 | yes | 描述 |
| `create_time` | bigint(20) | 0 | no | 创建时间 |
| `updated_time` | bigint(20) | 0 | no | 修改时间 |
| `is_open_manhour` | tinyint(1) | 0 | yes | 工时开关 |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `space_uuid` | varchar(8) | 无 | no | 所属 space |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `title` | varchar(64) | 无 | no | 模版标题 |
| `title_pinyin` | varchar(256) | 空串 | no | 模版标题-拼音 |
| `status` | tinyint(4) | 无 | no | 状态 (1: 正常, 2: 删除) |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `update_time` | bigint(20) | 无 | no | 最近更新时间 |

## `template_article` - 用户保存的编辑内容

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `template_draft_uuid` | varchar(8) | 空串 | no | templata draft uuid |
| `template_uuid` | varchar(8) | 无 | no | template uuid |
| `title` | varchar(64) | 无 | no | 标题 |
| `content` | longtext | 无 | no | 内容 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `update_time` | bigint(20) | 无 | no | 最近更新时间 |

## `template_attachment`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `space_uuid` | varchar(8) | 无 | no |  |
| `template_uuid` | varchar(8) | 无 | no |  |
| `template_draft_uuid` | varchar(8) | 空串 | no |  |
| `name` | varchar(255) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `server_update_stamp` | bigint(20) | 无 | no |  |

## `template_card`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no |  |
| `space_uuid` | varchar(8) | 无 | no |  |
| `template_uuid` | varchar(8) | 无 | no |  |
| `template_draft_uuid` | varchar(8) | 无 | no |  |
| `type` | int(11) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `create_time` | bigint(20) | 无 | no |  |
| `config` | text | 无 | no | 配置 |

## `template_card_attachment`

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `template_card_uuid` | varchar(8) | 无 | no |  |
| `template_attachment_uuid` | varchar(8) | 无 | no |  |
| `status` | tinyint(4) | 无 | no | 1: 正常, 2: 删除 |
| `server_update_stamp` | bigint(20) | 无 | no |  |

## `template_detail` - template 对应版本信息及详情信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `template_uuid` | varchar(8) | 无 | no | teamplate_uuid |
| `version` | int(11) | 无 | no | template 版本, 从0开始递增 |
| `template_draft_uuid` | varchar(8) | 空串 | no | template draft uuid |
| `is_revert` | tinyint(1) | 无 | no | 标记是否为回滚某个版本的 template |

## `template_draft` - 所有草稿及发布的 template

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| `uuid` | varchar(8) | 无 | no | uuid |
| `team_uuid` | varchar(8) | 无 | no | 所属团队 |
| `space_uuid` | varchar(8) | 无 | no | 所属 space |
| `template_uuid` | varchar(8) | 无 | no | template uuid |
| `owner_uuid` | varchar(8) | 无 | no | 拥有者 |
| `version` | int(11) | 无 | yes | -1: 编辑状态, >0: 已发布, NULL: 草稿/已删<br>除 |
| `from_version` | int(11) | 无 | no | 草稿来源那个发布版本 |
| `status` | tinyint(4) | 无 | no | 1: 草稿状态, 2: 编辑状态 3: 发布状态 4: <br>删除 |
| `create_time` | bigint(20) | 无 | no | 创建时间 |
| `updated_time` | bigint(20) | 无 | no | 最近更新时间 |

