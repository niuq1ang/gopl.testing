package filter

import (
	"encoding/json"

	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/batch"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/message"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/ngaut/log"
)

const (
	JobTypeMoveTasks             = "move_tasks"
	JobTypeDeleteTasks           = "delete_tasks"
	JobTypeImportTasks           = "import_tasks"
	JobTypeImportLibraryTestcase = "import_library_testcase"

	JobTypeModifyIssueType   = "modify_issue_type"
	JobTypeStdToSubIssueType = "std_to_sub_issue_type"
	JobTypeStdToStdIssueType = "std_to_std_issue_type"
	JobTypeSubToStdIssueType = "sub_to_std_issue_type"
	JobTypeSubToSubIssueType = "sub_to_sub_issue_type"
	JobTypeSubModifyParent   = "sub_modify_parent"
)

const (
	BatchActionModifyIssueType   string = JobTypeModifyIssueType
	BatchActionStdToSubIssueType string = JobTypeStdToSubIssueType
	BatchActionStdToStdIssueType string = JobTypeStdToStdIssueType
	BatchActionSubToStdIssueType string = JobTypeSubToStdIssueType
	BatchActionSubToSubIssueType string = JobTypeSubToSubIssueType
	BatchActionSubModifyParent   string = JobTypeSubModifyParent
)

const (
	SummaryFieldUUID           = "field001" // 任务标题
	DescFieldUUID              = "field002" // 任务描述（纯文本）
	OwnerFieldUUID             = "field003" // 任务创建者
	AssignFieldUUID            = "field004" // 任务负责人
	StatusFieldUUID            = "field005" // 任务状态
	ProjectFieldUUID           = "field006" // 任务所属项目
	StandardIssueTypeFieldUUID = "field007" // 任务类型
	WatcherFieldUUID           = "field008" // 任务关注者
	CreateTimeFieldUUID        = "field009" // 任务创建时间
	UpdateTimeFieldUUID        = "field010" // 任务更新时间
	SprintFieldUUID            = "field011" // 所属迭代
	PriorityFieldUUID          = "field012" // 优先级
	DeadlineFieldUUID          = "field013" // 截止日期
	ParentFieldUUID            = "field014" // 父任务
	NumberFieldUUID            = "field015" // 任务编号
	DescRichFieldUUID          = "field016" // 任务描述（富文本）
	StatusCategoryFieldUUID    = "field017" // 任务状态分类
	AssessManhourFieldUUID     = "field018" // 预估工时
	TotalManhourFieldUUID      = "field019" // 已登记工时合计
	RemainingManhourFieldUUID  = "field020" // 剩余工时
	SubIssueTypeFieldUUID      = "field021" // 子任务类型
	RelatedFieldUUID           = "field022" // 关联任务
	RelatedCountFieldUUID      = "field023" // 关联任务数量
	AltIssueTypeFieldUUID      = "field024" // 特殊的任务类型（包含「任务类型」「子任务类型」，仅作为筛选，无实际字段）
	EstimateVarianceFieldUUID  = "field025" // 预估偏差
	TimeProgressFieldUUID      = "field026" // 工时进度
	PlanStartDateFieldUUID     = "field027" // 计划开始日期
	PlanEndDateFieldUUID       = "field028" // 计划结束日期
	ProductsFieldUUID          = "field029" // 所属产品（多选项）
	ProductModuleFieldUUID     = "field030" // 所属产品模块（多选项）
)

var NoticeWithBatchActionSet = map[batch.Action]string{
	batch.BatchActionModifyIssueType:   NoticeTypeUpdateIssueType,
	batch.BatchActionStdToStdIssueType: NoticeTypeUpdateIssueType,
	batch.BatchActionStdToSubIssueType: NoticeTypeUpdateStdToSubIssueType,
	batch.BatchActionSubToSubIssueType: NoticeTypeUpdateSubToSubIssueType,
	batch.BatchActionSubModifyParent:   NoticeTypeUpdateSubIssueTypeParent,
	batch.BatchActionSubToStdIssueType: NoticeTypeUpdateSubToStdIssueType,
}

func GetMessageNoticeType(m *message.Message) string {
	if m.Type == message.MessageTypeDiscussion {
		if m.ResourceUUID == "" {
			return NoticeTypeUpdateTaskMessage
		} else {
			return NoticeTypeUploadTaskFile
		}
	}
	switch m.ObjectAttr {
	case 0:
		switch m.Action {
		case utils.EntityActionAdd:
			return NoticeTypeCreateTask
		}
	case utils.EntityAttrStatus:
		switch m.Action {
		case utils.EntityActionDelete:
			return NoticeTypeDeleteTask
		}
	case utils.EntityAttrField:
		fieldExt := &message.TaskFieldExt{}
		err := json.Unmarshal([]byte(m.Ext), fieldExt)
		if err != nil {
			log.Warn(err.Error())
			return NoticeTypeUpdateTaskOtherProperty
		}
		return fieldUUIDToTaskType(fieldExt.FieldUUID)
	case utils.EntityAttrRelatedTask:
		return NoticeTypeUpdateTaskRelatedTask
	case utils.EntityAttrSubtask:
	case utils.EntityAttrWatchers:
		return NoticeTypeUpdateTaskWatcher
	case utils.EntityAttrAssessManhour:
	case utils.EntityAttrAttachments:
		return NoticeTypeUploadTaskFile
	case utils.EntityAttrRelatedTestCase:
		return NoticeTypeRelatedTestCase
	case utils.EntityAttrRelatedPlanCase:
		return NoticeTypeRelatedPlanCase
	case utils.EntityAttrTaskRelatedTestCasePlan:
		return NoticeTypeTaskRelatedTestcasePlan
	case utils.EntityAttrIssueType, utils.EntityAttrParent:
		if len(m.Ext) > 0 {
			var ext message.TaskFieldExt
			err := json.Unmarshal([]byte(m.Ext), &ext)
			if err != nil {
				return NoticeTypeUpdateIssueType
			}
			noticeType := NoticeWithBatchActionSet[ext.BatchAction]
			return noticeType
		}
		return NoticeTypeUpdateIssueType
	case utils.EntityAttrManhours:
		return NoticeTypeUpdateTaskRecordManhour
	case utils.EntityAttrProduct:
		return NoticeTypeUpdateTaskProduct
	case utils.EntityAttrTaskRelatedWikiPage:
		return NoticeTypeTaskRelatedWikiPage
	default:
	}
	return NoticeTypeUpdateTaskOtherProperty
}

func fieldUUIDToTaskType(fieldUUID string) string {
	switch fieldUUID {
	case SummaryFieldUUID: // 任务标题
		return NoticeTypeUpdateTaskTitle
	case DescFieldUUID: // 任务描述（纯文本）
		return NoticeTypeUpdateTaskDescription
	case OwnerFieldUUID: // 任务创建者
	case AssignFieldUUID: // 任务负责人
		return NoticeTypeUpdateTaskAssign
	case StatusFieldUUID: // 任务状态
		return NoticeTypeUpdateTaskStatus
	case ProjectFieldUUID: // 任务所属项目
	case StandardIssueTypeFieldUUID: // 任务类型
	case SubIssueTypeFieldUUID: // 子任务类型
	case WatcherFieldUUID: // 任务关注者
		return NoticeTypeUpdateTaskWatcher
	case CreateTimeFieldUUID: // 任务创建时间
	case UpdateTimeFieldUUID: // 任务更新时间
	case SprintFieldUUID: // 所属迭代
		return NoticeTypeUpdateTaskSprint
	case PriorityFieldUUID: // 优先级
		return NoticeTypeUpdateTaskPriority
	case DeadlineFieldUUID: // 截止日期
		return NoticeTypeUpdateTaskDeadline
	case ParentFieldUUID: // 父任务
	case NumberFieldUUID: // 任务编号
	case DescRichFieldUUID: // 任务描述（富文本）
		return NoticeTypeUpdateTaskDescription
	case StatusCategoryFieldUUID: // 任务状态分类
	case AssessManhourFieldUUID: // 预估工时
		return NoticeTypeUpdateTaskAssessManhour
	case TotalManhourFieldUUID: // 已登记工时合计
	case RemainingManhourFieldUUID: // 剩余工时
		return NoticeTypeUpdateTaskRemainingManhour
	case EstimateVarianceFieldUUID: // 预估偏差
	case TimeProgressFieldUUID: //工时进度
	case ProductsFieldUUID: // 所属产品
		return NoticeTypeUpdateTaskProduct
	}
	return NoticeTypeUpdateTaskOtherProperty
}
