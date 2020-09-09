package filter

const (
	// 工作项提醒
	NoticeConst = 1 << 0
	// 创建工作项
	CreateTaskConst = 1 << 1
	// 更改负责人
	UpdateTaskAssignConst = 1 << 2
	// 更改工作项状态
	UpdateTaskStatusConst = 1 << 3
	// 更改工作项优先级
	UpdateTaskPriorityConst = 1 << 4
	// 更改工作项标题
	UpdateTaskTitleConst = 1 << 5
	// 更改工作项描述
	UpdateTaskDescriptionConst = 1 << 6
	// 设置工作项截止时间
	SetTaskDeadlineConst = 1 << 7
	// 更改工作项所属迭代
	UpdateTaskSprintConst = 1 << 8
	// 修改工作项评论
	UpdateTaskMessageConst = 1 << 9
	// 更改工作项关注者
	UpdateTaskWatcherconst = 1 << 10
	// 设置关联工作项
	UpdateTaskRelatedTaskConst = 1 << 11
	// 设置关联执行结果
	UpdateTaskRelatedPlanCaseConst = 1 << 12
	// 设置关联测试用例
	UpdateTaskRelatedTestCaseConst = 1 << 13
	// 设置关联测试计划
	UpdateTaskRelatedTestcasePlanConst = 1 << 14
	// 设置关联wiki页面
	UpdateTaskRelatedWikiPageConst = 1 << 15
	// 上传文件
	UploadTaskFileConst = 1 << 16
	// 更新预估工时
	UpdateTaskAccessManhourConst = 1 << 17
	// 更新剩余工时
	UpdateTaskRemainingManhourConst = 1 << 18
	// 更新登记工时
	UpdateTaskRecordManhourConst = 1 << 19
	// 更新自定义属性
	UpdateTaskOtherPropertyConst = 1 << 20
	// 删除任务
	DeleteTaskConst = 1 << 21
	// 更改所属产品
	UpdateTaskProductConst = 1 << 22
	// 变更工作项类型
	UpdateIssueTypeConst = 1 << 23
	// 转为子工作项
	UpdateStdToSubIssueTypeConst = 1 << 24
	// 变更子工作项类型
	UpdateSubToSubIssueTypeConst = 1 << 25
	// 子工作项转父工作项
	UpdateSubIssueTypeParentConst = 1 << 26
	// 子工作类型转标准工作项类型
	UpdateSubToStdIssueTypeConst = 1 << 27
)

const (
	// 工作项提醒
	TaskNotice = "task_notice"
	//  通知类型枚举
	NoticeTypeCreateTask                 = "create_task"
	NoticeTypeUpdateTaskAssign           = "update_task_assign"
	NoticeTypeUpdateTaskStatus           = "update_task_status"
	NoticeTypeUpdateTaskPriority         = "update_task_priority"
	NoticeTypeUpdateTaskTitle            = "update_task_title"
	NoticeTypeUpdateTaskDescription      = "update_task_description"
	NoticeTypeUpdateTaskDeadline         = "set_task_deadline"
	NoticeTypeUpdateTaskSprint           = "update_task_sprint"
	NoticeTypeUpdateTaskMessage          = "update_task_message"
	NoticeTypeUpdateTaskWatcher          = "update_task_watcher"
	NoticeTypeUpdateTaskRelatedTask      = "update_task_related_task"
	NoticeTypeRelatedPlanCase            = "related_plan_case"
	NoticeTypeRelatedTestCase            = "related_test_case"
	NoticeTypeTaskRelatedTestcasePlan    = "task_related_testcase_plan"
	NoticeTypeTaskRelatedWikiPage        = "task_related_wiki_page"
	NoticeTypeUploadTaskFile             = "upload_task_file"
	NoticeTypeUpdateTaskAssessManhour    = "update_task_access_manhour"
	NoticeTypeUpdateTaskOtherProperty    = "update_task_other_property"
	NoticeTypeDeleteTask                 = "delete_task"
	NoticeTypeUpdateTaskRemainingManhour = "update_task_remaining_manhour"
	NoticeTypeUpdateTaskRecordManhour    = "update_task_record_manhour"
	NoticeTypeUpdateTaskProduct          = "update_task_product"

	// 标准 工作项类型
	NoticeTypeUpdateIssueType         = "update_issue_type"
	NoticeTypeUpdateStdToSubIssueType = "update_std_to_sub_issue_type"

	// 子 工作项类型
	NoticeTypeUpdateSubToSubIssueType  = "update_sub_to_sub_issue_type"
	NoticeTypeUpdateSubIssueTypeParent = "update_sub_issue_type_parent"
	NoticeTypeUpdateSubToStdIssueType  = "update_sub_to_std_issue_type"
)

var IssueFieldIntergeMap = map[string]int64{
	TaskNotice:                           NoticeConst,
	NoticeTypeCreateTask:                 CreateTaskConst,
	NoticeTypeUpdateTaskAssign:           UpdateTaskAssignConst,
	NoticeTypeUpdateTaskPriority:         UpdateTaskPriorityConst,
	NoticeTypeUpdateTaskTitle:            UpdateTaskTitleConst,
	NoticeTypeUpdateTaskDescription:      UpdateTaskDescriptionConst,
	NoticeTypeUpdateTaskDeadline:         SetTaskDeadlineConst,
	NoticeTypeUpdateTaskSprint:           UpdateTaskSprintConst,
	NoticeTypeUpdateTaskMessage:          UpdateTaskMessageConst,
	NoticeTypeUpdateTaskWatcher:          UpdateTaskWatcherconst,
	NoticeTypeUpdateTaskRelatedTask:      UpdateTaskRelatedTaskConst,
	NoticeTypeRelatedPlanCase:            UpdateTaskRelatedPlanCaseConst,
	NoticeTypeRelatedTestCase:            UpdateTaskRelatedTestCaseConst,
	NoticeTypeTaskRelatedTestcasePlan:    UpdateTaskRelatedTestcasePlanConst,
	NoticeTypeTaskRelatedWikiPage:        UpdateTaskRelatedWikiPageConst,
	NoticeTypeUploadTaskFile:             UploadTaskFileConst,
	NoticeTypeUpdateTaskAssessManhour:    UpdateTaskAccessManhourConst,
	NoticeTypeUpdateTaskOtherProperty:    UpdateTaskOtherPropertyConst,
	NoticeTypeDeleteTask:                 DeleteTaskConst,
	NoticeTypeUpdateTaskRemainingManhour: UpdateTaskRemainingManhourConst,
	NoticeTypeUpdateTaskRecordManhour:    UpdateTaskRecordManhourConst,
	NoticeTypeUpdateTaskProduct:          UpdateTaskProductConst,
	NoticeTypeUpdateIssueType:            UpdateIssueTypeConst,
	NoticeTypeUpdateStdToSubIssueType:    UpdateStdToSubIssueTypeConst,
	NoticeTypeUpdateSubToSubIssueType:    UpdateSubToSubIssueTypeConst,
	NoticeTypeUpdateSubIssueTypeParent:   UpdateSubIssueTypeParentConst,
	NoticeTypeUpdateSubToStdIssueType:    UpdateSubToStdIssueTypeConst,
}

var NoticeTypeNames = map[string]string{
	TaskNotice:                           "工作项提醒",
	NoticeTypeCreateTask:                 "创建任务",
	NoticeTypeUpdateTaskAssign:           "更改负责人",
	NoticeTypeUpdateTaskStatus:           "更新任务状态",
	NoticeTypeUpdateTaskPriority:         "更改优先级",
	NoticeTypeUpdateTaskTitle:            "更改任务标题",
	NoticeTypeUpdateTaskDescription:      "更改任务描述",
	NoticeTypeUpdateTaskDeadline:         "设置截止时间",
	NoticeTypeUpdateTaskSprint:           "设置所属迭代",
	NoticeTypeUpdateTaskMessage:          "新增/更改评论",
	NoticeTypeUpdateTaskWatcher:          "新增/取消关注任务",
	NoticeTypeUpdateTaskRelatedTask:      "设置关联任务",
	NoticeTypeRelatedPlanCase:            "设置关联执行结果",
	NoticeTypeRelatedTestCase:            "设置关联用例库用例",
	NoticeTypeTaskRelatedTestcasePlan:    "设置任务关联测试计划",
	NoticeTypeTaskRelatedWikiPage:        "设置任务关联 Wiki 页面",
	NoticeTypeUploadTaskFile:             "上传文件",
	NoticeTypeUpdateTaskAssessManhour:    "设置预估工时",
	NoticeTypeUpdateTaskOtherProperty:    "更改自定义属性",
	NoticeTypeDeleteTask:                 "删除任务",
	NoticeTypeUpdateTaskRemainingManhour: "更新剩余工时",
	NoticeTypeUpdateTaskRecordManhour:    "更新登记工时",
	NoticeTypeUpdateTaskProduct:          "更改所属产品",
	NoticeTypeUpdateIssueType:            "变更工作项类型",
	NoticeTypeUpdateStdToSubIssueType:    "转为子工作项",
	NoticeTypeUpdateSubToSubIssueType:    "变更子工作项类型",
	NoticeTypeUpdateSubIssueTypeParent:   "变更父工作项",
	NoticeTypeUpdateSubToStdIssueType:    "转为标准工作项",
}
