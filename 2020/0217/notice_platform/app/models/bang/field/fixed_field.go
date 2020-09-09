package field

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/common/i18n"
	"golang.org/x/text/language"
)

const (
	TaskUUID = "uuid"
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

var (
	FixedFieldCreateTime   = int64(1472688000000000)
	FixedFieldDesc         = ""
	FixedFieldRenderer     = 0
	FixedFieldFilterOption = 0
	FixedFieldSearchOption = 0
	FixedFieldBuiltIn      = true
	//固有属性配置第一项的position
	FixedFieldConfigFirstPosition = int64(100000)

	//获取属性列表需要返回的固有属性
	FixedFieldUUIDInFieldList = []string{
		SummaryFieldUUID,
		DescFieldUUID,
		OwnerFieldUUID,
		AssignFieldUUID,
		StatusFieldUUID,
		ProjectFieldUUID,
		StandardIssueTypeFieldUUID,
		WatcherFieldUUID,
		CreateTimeFieldUUID,
		UpdateTimeFieldUUID,
		SprintFieldUUID,
		PriorityFieldUUID,
		DeadlineFieldUUID,
		NumberFieldUUID,
		DescRichFieldUUID,
		AssessManhourFieldUUID,
		TotalManhourFieldUUID,
		RemainingManhourFieldUUID,
		SubIssueTypeFieldUUID,
		EstimateVarianceFieldUUID,
		TimeProgressFieldUUID,
		ProductsFieldUUID,
	}

	// 需要出现在「project配置中心」=>「工作项属性」中的 FixedField
	FixedFieldUUIDInProjectSetting = append(FixedFieldUUIDInFieldList,
		ProductModuleFieldUUID,
	)

	//创建任务中使用到的固有属性
	FixedFieldUUIDs = append(FixedFieldUUIDInFieldList,
		ParentFieldUUID,
		TaskUUID,
		RelatedCountFieldUUID,
		ProductModuleFieldUUID,
	)

	//报表用到的固有属性
	ReportFixedFieldUUIDs = append(FixedFieldUUIDInFieldList, ParentFieldUUID, StatusCategoryFieldUUID)

	FixedFieldMap = map[string]string{
		SummaryFieldUUID:           "fixed_field_title_summary",
		DescFieldUUID:              "fixed_field_title_desc",
		OwnerFieldUUID:             "fixed_field_title_owner",
		AssignFieldUUID:            "fixed_field_title_assign",
		StatusFieldUUID:            "fixed_field_title_status",
		ProjectFieldUUID:           "fixed_field_title_project",
		StandardIssueTypeFieldUUID: "fixed_field_title_type",
		WatcherFieldUUID:           "fixed_field_title_watcher",
		CreateTimeFieldUUID:        "fixed_field_title_create_time",
		UpdateTimeFieldUUID:        "fixed_field_title_upadate_time",
		SprintFieldUUID:            "fixed_field_title_sprint",
		PriorityFieldUUID:          "fixed_field_title_priority",
		DeadlineFieldUUID:          "fixed_field_title_deadline",
		ParentFieldUUID:            "fixed_field_title_parent",
		NumberFieldUUID:            "fixed_field_title_number",
		DescRichFieldUUID:          "fixed_field_title_desc_rich",
		AssessManhourFieldUUID:     "fixed_field_title_assess_manhour",
		TotalManhourFieldUUID:      "fixed_field_title_total_manhour",
		RemainingManhourFieldUUID:  "fixed_field_title_remaining_manhour",
		StatusCategoryFieldUUID:    "fixed_field_title_status_category",
		SubIssueTypeFieldUUID:      "fixed_field_title_sub_type",
		RelatedCountFieldUUID:      "fixed_field_title_related_count",
		EstimateVarianceFieldUUID:  "fixed_field_title_estimate_variance",
		TimeProgressFieldUUID:      "fixed_field_title_time_progress",
		ProductsFieldUUID:          "fixed_field_title_products",
		ProductModuleFieldUUID:     "fixed_field_title_product_modeule",
		PlanStartDateFieldUUID:     "fixed_field_title_plan_start_date",
		PlanEndDateFieldUUID:       "fixed_field_title_plan_end_date",
	}

	// 值由系统生成，用户不可修改的属性
	GeneratedFieldUUIDs = []string{
		OwnerFieldUUID,
		CreateTimeFieldUUID,
		UpdateTimeFieldUUID,
		NumberFieldUUID,
		StatusCategoryFieldUUID,
		TotalManhourFieldUUID,
		RelatedCountFieldUUID,
		EstimateVarianceFieldUUID,
		TimeProgressFieldUUID,
	}

	// 该hash表中的"系统字段"表示可以被修改
	ModifiedFixedFieldMap = map[string]bool{
		PriorityFieldUUID: true,
	}

	// 不允许设置默认值的固有字段
	DisableSetDefaultValueFixedFieldMap = func() map[string]bool {
		m := map[string]bool{}
		fieldUUIDs := []string{
			SummaryFieldUUID,
			OwnerFieldUUID,
			ProjectFieldUUID,
			StandardIssueTypeFieldUUID,
			CreateTimeFieldUUID,
			UpdateTimeFieldUUID,
			DeadlineFieldUUID,
			ParentFieldUUID,
			NumberFieldUUID,
			StatusCategoryFieldUUID,
			AssessManhourFieldUUID,
			TotalManhourFieldUUID,
			RemainingManhourFieldUUID,
			SubIssueTypeFieldUUID,
			EstimateVarianceFieldUUID,
			TimeProgressFieldUUID,
		}
		for _, fieldUUID := range fieldUUIDs {
			m[fieldUUID] = true
		}
		return m
	}()
)

// func GetFieldTypeByFixedFieldUUID(fixedFieldUUID string) (fieldType int, err error) {
// 	switch fixedFieldUUID {
// 	case SummaryFieldUUID, DescFieldUUID, DescRichFieldUUID:
// 		fieldType = FieldTypeText
// 	case OwnerFieldUUID, AssignFieldUUID:
// 		fieldType = FieldTypeUser
// 	case StatusFieldUUID:
// 		fieldType = FieldTypeTaskStatus
// 	case ProjectFieldUUID:
// 		fieldType = FieldTypeProject
// 	case StandardIssueTypeFieldUUID:
// 		fieldType = FieldTypeIssueType
// 	case SubIssueTypeFieldUUID:
// 		fieldType = FieldTypeIssueType
// 	case WatcherFieldUUID:
// 		fieldType = FieldTypeUserList
// 	case CreateTimeFieldUUID, UpdateTimeFieldUUID:
// 		fieldType = FieldTypeTime
// 	case SprintFieldUUID:
// 		fieldType = FieldTypeSprint
// 	case PriorityFieldUUID:
// 		fieldType = FieldTypeOption
// 	case DeadlineFieldUUID:
// 		fieldType = FieldTypeDate
// 	case ParentFieldUUID:
// 		fieldType = FieldTypeTask
// 	case NumberFieldUUID:
// 		fieldType = FieldTypeNumber
// 	case AssessManhourFieldUUID, TotalManhourFieldUUID, RemainingManhourFieldUUID, EstimateVarianceFieldUUID, TimeProgressFieldUUID:
// 		fieldType = FieldTypeFloat
// 	case StatusCategoryFieldUUID:
// 		fieldType = FieldTypeStatusCategory
// 	case TaskUUID:
// 		fieldType = FieldTypeTask
// 	case RelatedCountFieldUUID:
// 		fieldType = FieldTypeInteger
// 	case ProductsFieldUUID:
// 		fieldType = FieldTypeProduct
// 	case ProductModuleFieldUUID:
// 		fieldType = FieldTypeProductModule
// 	default:
// 		err = errors.UnexpectedArgumentsError(fixedFieldUUID)
// 	}
// 	return
// }

func GetRequiredByFixedFieldUUID(fixedFieldUUID string) (required bool, err error) {
	switch fixedFieldUUID {
	case SummaryFieldUUID,
		OwnerFieldUUID,
		AssignFieldUUID,
		StatusFieldUUID,
		ProjectFieldUUID,
		StandardIssueTypeFieldUUID,
		CreateTimeFieldUUID,
		UpdateTimeFieldUUID,
		PriorityFieldUUID:
		required = true
	case SprintFieldUUID,
		DeadlineFieldUUID,
		WatcherFieldUUID,
		NumberFieldUUID,
		DescFieldUUID,
		DescRichFieldUUID,
		ParentFieldUUID,
		AssessManhourFieldUUID,
		TotalManhourFieldUUID,
		RemainingManhourFieldUUID,
		SubIssueTypeFieldUUID,
		EstimateVarianceFieldUUID,
		TimeProgressFieldUUID,
		ProductsFieldUUID:
		required = false
	default:
		err = errors.UnexpectedArgumentsError(fixedFieldUUID)
	}
	return
}

func GetFixedFieldNameByKey(nameKey string) string {
	lang := []language.Tag{language.Chinese}
	return i18n.String(lang, nameKey)
}

func GetFixedFieldNameByUUID(fixedFieldUUID string) string {
	return GetFixedFieldNameByKey(FixedFieldMap[fixedFieldUUID])
}

func GetTaskPriorityValue(priUUID string) string {
	lang := []language.Tag{language.Chinese}
	return i18n.String(lang, priUUID)
}

func IsFixedField(id string) bool {
	for _, fieldUUID := range FixedFieldUUIDs {
		if id == fieldUUID {
			return true
		}
	}
	return false
}

// func BuildFixedField(fieldUUID string) (*Field, error) {
// 	index := -1
// 	for i, fuuid := range FixedFieldUUIDs {
// 		if fuuid == fieldUUID {
// 			index = i
// 			break
// 		}
// 	}
// 	if index < 0 {
// 		return nil, errors.UnexpectedArgumentsError(fieldUUID)
// 	}
// 	return buildFixedField(index)
// }

// func BuildReportFixedField(fieldUUID string) (*Field, error) {
// 	index := -1
// 	for i, fuuid := range ReportFixedFieldUUIDs {
// 		if fuuid == fieldUUID {
// 			index = i
// 			break
// 		}
// 	}
// 	if index < 0 {
// 		return nil, errors.UnexpectedArgumentsError(fieldUUID)
// 	}
// 	return buildReportFixedField(index)
// }

// func BuildAllFixedFields() ([]*Field, error) {
// 	l := len(FixedFieldUUIDs)
// 	fields := make([]*Field, l)
// 	for i := 0; i < l; i++ {
// 		f, err := buildFixedField(i)
// 		if err != nil {
// 			return nil, errors.Trace(err)
// 		}
// 		fields[i] = f
// 	}
// 	return fields, nil
// }

// func buildFixedField(index int) (*Field, error) {
// 	uuid := FixedFieldUUIDs[index]
// 	createTime := FixedFieldCreateTime + int64(index)
// 	return doBuildFixedField(uuid, createTime)
// }

// func buildReportFixedField(index int) (*Field, error) {
// 	uuid := ReportFixedFieldUUIDs[index]
// 	createTime := FixedFieldCreateTime + int64(index)
// 	return doBuildFixedField(uuid, createTime)
// }

// func doBuildFixedField(uuid string, createTime int64) (*Field, error) {
// 	nameKey := FixedFieldMap[uuid]
// 	name := GetFixedFieldNameByKey(nameKey)
// 	fieldType, err := GetFieldTypeByFixedFieldUUID(uuid)
// 	if err != nil {
// 		return nil, errors.Trace(err)
// 	}

// 	f := &Field{
// 		UUID:         uuid,
// 		Name:         name,
// 		Desc:         FixedFieldDesc,
// 		Type:         fieldType,
// 		DefaultValue: nil,
// 		Renderer:     FixedFieldRenderer,
// 		FilterOption: FixedFieldFilterOption,
// 		SearchOption: FixedFieldSearchOption,
// 		CreateTime:   createTime,
// 		BuiltIn:      FixedFieldBuiltIn,
// 	}
// 	return f, nil
// }

func CanModifyRequiredFixedField(fieldUUID string) bool {
	switch fieldUUID {
	case DeadlineFieldUUID,
		SprintFieldUUID,
		AssessManhourFieldUUID,
		DescRichFieldUUID,
		DescFieldUUID,
		ProductsFieldUUID,
		ProductModuleFieldUUID:
		return true
	}
	return false
}

// func CanModifyRequiredField(fd *Field, fieldUUID string) bool {
// 	if fd != nil {
// 		if ShouldCalFieldTypesSet[fd.Type] {
// 			return false
// 		}
// 	}
// 	return !IsFixedField(fieldUUID) || CanModifyRequiredFixedField(fieldUUID)
// }

func CanDeleteField(fieldUUID string) bool {
	if fieldUUID == ProductModuleFieldUUID {
		return true
	}
	return !IsFixedField(fieldUUID)
}

func CanAddField(fieldUUID string) bool { return CanDeleteField(fieldUUID) }

var generatedFieldSet = func() map[string]struct{} {
	set := make(map[string]struct{})
	for _, uuid := range GeneratedFieldUUIDs {
		set[uuid] = struct{}{}
	}
	return set
}()

func IsGeneratedField(fieldUUID string) bool {
	_, ok := generatedFieldSet[fieldUUID]
	return ok
}

func FieldPositionCorrection(from int64, fieldUUID string) int64 {
	// 截止日期特殊处理
	switch fieldUUID {
	case DeadlineFieldUUID, ProductsFieldUUID, ProductModuleFieldUUID:
		return from
	}
	// 除了截止日期以外的固有属性的处理
	for i, f := range FixedFieldUUIDs {
		if f == fieldUUID {
			return FixedFieldConfigFirstPosition + int64(i)
		}
	}
	return from
}

// 系统field是否可以被修改
func IsModifiedFixedField(fieldUUID string) bool {
	return ModifiedFixedFieldMap[fieldUUID]
}
