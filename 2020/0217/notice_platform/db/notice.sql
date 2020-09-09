create database ones_notice;

 create table `notice`(
  `seq_no` bigint(20) NOT NULL DEFAULT '0' COMMENT '唯一序列号16位，取系统更新时间,微妙',
  `team_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
  `task_uuid` varchar(16) COLLATE latin1_bin NOT NULL,
  `from_user_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
  `message_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
  `message` text CHARACTER SET utf8mb4 NOT NULL,
   PRIMARY KEY (`seq_no`),
   KEY `idx_team_uuid` (`team_uuid`),
   KEY `idx_message_uuid` (`message_uuid`)
 )ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;
 
 create table `send_record`(
   `seq_no` bigint(20) NOT NULL DEFAULT '0' COMMENT '唯一序列号16位，例：2020011500000000',
   `team_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
   `tunnel_uuid` varchar(8) COLLATE latin1_bin NOT NULl,
   `status` tinyint(4) NOT NULL COMMENT '发送状态：1：ready, 2 success, 3 failure, 4 resend, 5 filterout',
   PRIMARY KEY (`seq_no`),
   KEY `idx_team_uuid` (`team_uuid`),
   KEY `idx_tunnel_uuid` (`tunnel_uuid`)
  )ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

 create table `tunnel`(
  `uuid` varchar(8) COLLATE latin1_bin NOT NULl,
  `team_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
  `link_type` tinyint(4) NOT NULL COMMENT '连接类型: webhook:1, wechat:2, dingding:3 ...',
  `link_uuid` varchar(8) COLLATE latin1_bin NOT NULl COMMENT '连接配置信息',
  `filter_type` tinyint(4) NOT NULL COMMENT '过滤器类型，task_filter:1,',
  `filter_uuid` varchar(8) COLLATE latin1_bin NOT NULl COMMENT '过滤器配置信息',
   PRIMARY KEY (`uuid`)
 )ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

 create table `webhook`(
  `uuid` varchar(8) COLLATE latin1_bin NOT NULl,
  `team_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
  `name` varchar(32) COLLATE latin1_bin NOT NULL,
  `url` varchar(1024) COLLATE latin1_bin NOT NULl,
  `is_batch` tinyint(4) NOT NULL COMMENT '是否批量发送，1 是，2否',
  `batch_count` int(8) NOT NUll COMMENT '批量发送的条数',
  `status` tinyint(4) NOT NULL COMMENT '状态: 1 有效, 2 无效, 3 (不可用)unavailble, 4 已删除',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间(秒)',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间(秒)',
   PRIMARY KEY (`uuid`),
   KEY `idx_team_uuid` (`team_uuid`)
 )ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

 create table `task_filter`(
  `uuid` varchar(8) COLLATE latin1_bin NOT NULl,
  `team_uuid` varchar(8) COLLATE latin1_bin NOT NULL,
  `field_filter` bigint(20)  NOT NUll COMMENT '工作项属性过滤器',
  `issue_type_filter` bigint(20)  NOT NUll COMMENT '工作项类型过滤器',
   PRIMARY KEY (`uuid`),
   KEY `idx_team_uuid` (`team_uuid`)
 )ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;