# wiki 数据字典

## article - 用户保存的编辑内容

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| draft_uuid | varchar(8) | null | NO | page uuid |
| page_uuid | varchar(8) | null | NO | page uuid |
| title | varchar(64) | null | NO | 标题 |
| content | longtext | null | NO | 内容 |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 最近更新时间 |

## attachment

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| space_uuid | varchar(8) | null | NO |  |
| page_uuid | varchar(8) | null | NO |  |
| draft_uuid | varchar(8) | null | NO |  |
| name | varchar(255) | null | NO |  |
| status | tinyint(4) | null | NO | 1: 正常, 2: 删除 |
| server_update_stamp | bigint(20) | null | NO |  |

## card

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

## card_attachment

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| card_uuid | varchar(8) | null | NO |  |
| attachment_uuid | varchar(8) | null | NO |  |
| status | tinyint(4) | null | NO | 1: 正常, 2: 删除 |
| server_update_stamp | bigint(20) | null | NO |  |

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

## page_comment - page 评论

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| page_uuid | varchar(8) | null | NO | page uuid |
| owner_uuid | varchar(8) | null | NO | 拥有者 |
| subject | varchar(1024) | null | NO | 主题 |
| status | tinyint(4) | null | NO | 状态 (1: 正常状态 2: 删除) |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 最近更新时间 |

## page_detail - page 对应版本信息及详情信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| page_uuid | varchar(8) | null | NO | uuid |
| version | int(11) | null | NO | page 版本, 从0开始递增 |
| draft_uuid | varchar(8) | null | NO | uuid |
| is_revert | tinyint(1) | null | NO |  |

## page_draft - 所有草稿及发布的 page

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| team_uuid | varchar(8) | null | NO | 所属团队 |
| space_uuid | varchar(8) | null | NO | 所属 space |
| page_uuid | varchar(8) | null | NO | page uuid |
| owner_uuid | varchar(8) | null | NO | 拥有者 |
| version | int(11) | null | YES | -1: 编辑状态, >0: 已发布, NULL: 草稿/已删除 |
| from_version | int(11) | null | NO | 草稿来源那个发布版本 |
| status | tinyint(4) | null | NO | 1: 草稿状态, 2: 编辑状态 3: 发布状态 4: 删除 |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 最近更新时间 |
| old_status | tinyint(4) | null | NO |  |

## page_share_config - 页面分享配置

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | 对应页面的共享uuid，随机生成的8位uuid |
| team_uuid | varchar(8) | null | NO | 所属团队 |
| space_uuid | varchar(8) | null | NO | 页面组uuid |
| page_uuid | varchar(8) | null | NO | 页面uuid |
| user_uuid | varchar(8) | null | NO | 分享人uuid |
| share_permission_type | int(1) | null | NO | 共享页面类型：0:不共享；1:共享给所有人查看；2:共享给所有人编辑；3:指定成员可以查看或编辑 |
| is_share_sub_page | int(1) | 0 | NO | 是否共享子页面：默认0：不共享子页面；1：共享子页面 |
| create_time | bigint(20) | null | NO | 创建时间 |
| update_time | bigint(20) | null | NO | 更新时间 |

## page_tree - page

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| team_uuid | varchar(8) | null | NO | 所属团队 |
| space_uuid | varchar(8) | null | NO | 所属 space |
| owner_uuid | varchar(8) | null | NO | 拥有者 |
| title | varchar(64) | null | NO | 最新标题 |
| parent_uuid | varchar(8) | null | NO | 父节点 uuid |
| status | tinyint(4) | null | NO | 状态 (1: 正常, 2: 删除) |
| position | bigint(20) | null | NO |  |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 最近更新时间 |
| old_parent_uuid | varchar(8) |  | NO | 被删除之前的父节点 uuid |
| old_previous_uuid | varchar(8) |  | NO | 被删除之前的前一个节点 uuid |
| encrypt_status | tinyint(4) | 1 | NO | 页面加密状态 (1: 不加密, 2: 仅查看，3：加密) |

## page_watch - 关注 page 的用户表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| page_uuid | varchar(8) | null | NO | page uuid |
| user_uuid | varchar(8) | null | NO | user uuid |

## space - space

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| team_uuid | varchar(8) | null | NO | 所属团队 |
| owner_uuid | varchar(8) | null | NO | 拥有者 |
| title | varchar(64) | null | NO | 标题 |
| title_pinyin | varchar(512) | null | NO |  |
| description | varchar(1024) | null | NO | 描述 |
| status | tinyint(4) | null | NO | 状态 (1: 正常, 2: 删除) |
| is_archive | tinyint(1) | null | NO |  |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 最近更新时间 |
| is_open_share_page | int(1) | 0 | NO | 是否开启共享页面的开关：默认0：不开启共享页面；1：开启共享页面 |

## space_browse - 用户浏览space记录表

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| user_uuid | varchar(8) | null | NO | 用户uuid |
| space_uuid | varchar(16) | null | NO | space_uuid |
| browse_time | bigint(20) | 0 | NO | 浏览时间 |

## space_pin - space置顶

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO |  |
| user_uuid | varchar(8) | null | NO |  |
| space_uuid | varchar(8) | null | NO | space_uuid |
| updated_time | bigint(20) | null | NO | 更新时间 |

## space_update_count - 记录space修改次数

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| team_uuid | varchar(8) | null | NO | 团队uuid |
| space_uuid | varchar(16) | null | NO | space_uuid |
| day | tinyint(4) | null | NO | 当月的第几天 |
| update_time | bigint(20) | 0 | NO | 统计时间 |
| count | bigint(20) | 0 | NO | 修改次数 |

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

## template - template

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

## template_article - 用户保存的编辑内容

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| template_draft_uuid | varchar(8) |  | NO | templata draft uuid |
| template_uuid | varchar(8) | null | NO | template uuid |
| title | varchar(64) | null | NO | 标题 |
| content | longtext | null | NO | 内容 |
| create_time | bigint(20) | null | NO | 创建时间 |
| update_time | bigint(20) | null | NO | 最近更新时间 |

## template_attachment

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| space_uuid | varchar(8) | null | NO |  |
| template_uuid | varchar(8) | null | NO |  |
| template_draft_uuid | varchar(8) |  | NO |  |
| name | varchar(255) | null | NO |  |
| status | tinyint(4) | null | NO | 1: 正常, 2: 删除 |
| server_update_stamp | bigint(20) | null | NO |  |

## template_card

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO |  |
| space_uuid | varchar(8) | null | NO |  |
| template_uuid | varchar(8) | null | NO |  |
| template_draft_uuid | varchar(8) | null | NO |  |
| type | int(11) | null | NO |  |
| status | tinyint(4) | null | NO | 1: 正常, 2: 删除 |
| create_time | bigint(20) | null | NO |  |
| config | text | null | NO | 配置 |

## template_card_attachment

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| template_card_uuid | varchar(8) | null | NO |  |
| template_attachment_uuid | varchar(8) | null | NO |  |
| status | tinyint(4) | null | NO | 1: 正常, 2: 删除 |
| server_update_stamp | bigint(20) | null | NO |  |

## template_detail - template 对应版本信息及详情信息

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| template_uuid | varchar(8) | null | NO | teamplate_uuid |
| version | int(11) | null | NO | template 版本, 从0开始递增 |
| template_draft_uuid | varchar(8) |  | NO | template draft uuid |
| is_revert | tinyint(1) | null | NO | 标记是否为回滚某个版本的 template |

## template_draft - 所有草稿及发布的 template

| 字段名 | 数据类型 | 默认值 | 允许为空 | 备注 |
| :--- | :--- | :--- | :--- | :--- |
| uuid | varchar(8) | null | NO | uuid |
| team_uuid | varchar(8) | null | NO | 所属团队 |
| space_uuid | varchar(8) | null | NO | 所属 space |
| template_uuid | varchar(8) | null | NO | template uuid |
| owner_uuid | varchar(8) | null | NO | 拥有者 |
| version | int(11) | null | YES | -1: 编辑状态, >0: 已发布, NULL: 草稿/已删除 |
| from_version | int(11) | null | NO | 草稿来源那个发布版本 |
| status | tinyint(4) | null | NO | 1: 草稿状态, 2: 编辑状态 3: 发布状态 4: 删除 |
| create_time | bigint(20) | null | NO | 创建时间 |
| updated_time | bigint(20) | null | NO | 最近更新时间 |

