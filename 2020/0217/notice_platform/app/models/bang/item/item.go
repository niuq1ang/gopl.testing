package item

import commonitem "github.com/bangwork/ones-ai-api-common/models/item"

type FieldTypeEnumConverter = commonitem.FieldTypeEnumConverter

const (
	FieldTypeAny                        = commonitem.FieldTypeAny
	FieldTypeOption                     = commonitem.FieldTypeOption
	FieldTypeText                       = commonitem.FieldTypeText
	FieldTypeInteger                    = commonitem.FieldTypeInteger
	FieldTypeFloat                      = commonitem.FieldTypeFloat
	FieldTypeDate                       = commonitem.FieldTypeDate
	FieldTypeTime                       = commonitem.FieldTypeTime
	FieldTypeUser                       = commonitem.FieldTypeUser
	FieldTypeStatus                     = commonitem.FieldTypeStatus
	FieldTypeUserList                   = commonitem.FieldTypeUserList
	FieldTypeNumber                     = commonitem.FieldTypeNumber
	FieldTypeMultiLineText              = commonitem.FieldTypeMultiLineText
	FieldTypeMultiOption                = commonitem.FieldTypeMultiOption
	FieldTypeStatusCategory             = commonitem.FieldTypeStatusCategory
	FieldTypeBoolean                    = commonitem.FieldTypeBoolean
	FieldTypeContext                    = commonitem.FieldTypeContext
	FieldTypeRichText                   = commonitem.FieldTypeRichText
	FieldTypeOptions                    = commonitem.FieldTypeOptions
	FieldTypeStatuses                   = commonitem.FieldTypeStatuses
	FieldTypeEnum                       = commonitem.FieldTypeEnum
	FieldTypeBucket                     = commonitem.FieldTypeBucket
	FieldTypeTimeSeries                 = commonitem.FieldTypeTimeSeries
	FieldTypeManhourReportConfig        = "manhour_report_config"
	FieldTypeManhour                    = "manhour"
	FieldTypeUserDomains                = commonitem.FieldTypeUserDomains
	FieldTypeGanttChartImportConfig     = "gantt_chart_import_config"
	FieldTypeActionScriptParam          = "action_script_param"
	FieldTypeTaskGanttChartImportConfig = "task_gantt_chart_import_config"
	FieldTypeTaskGanttChart             = "task_gantt_chart"

	FieldTypeSprint              = "sprint"
	FieldTypeProject             = "project"
	FieldTypeTask                = "task"
	FieldTypeIssueType           = "issue_type"
	FieldTypeTaskStatus          = "task_status"
	FieldTypeTestCaseLibrary     = "testcase_library"
	FieldTypeTestCaseModule      = "testcase_module"
	FieldTypeTestCaseCase        = "testcase_case"
	FieldTypeTestCaseCaseStep    = "testcase_case_step"
	FieldTypeTestCasePlan        = "testcase_plan"
	FieldTypeTestCaseFieldConfig = "testcase_field_config"
	FieldTypeProgram             = "program"
	FieldTypeProduct             = "product"
	FieldTypeProductModule       = "product_module"
	FieldTypeProductComponent    = "product_component"
	FieldTypeUserGroup           = "user_group"
	FieldTypeDepartment          = "department"

	FieldTypeStepRange  = "step_range"
	FieldTypeStayCount  = "stay_count"
	FieldTypeStayTime   = "stay_time"
	FieldTypeAppearTime = "appear_time"
)

var (
	NewDefaultFieldTypeEnumConverter = commonitem.NewDefaultFieldTypeEnumConverter
	NewFieldTypeEnumConverter        = commonitem.NewFieldTypeEnumConverter
)

var (
	DefaultFieldTypeEnumConverter *FieldTypeEnumConverter
)

func init() {
	c := NewDefaultFieldTypeEnumConverter()
	c.Set(FieldTypeSprint, 7)
	c.Set(FieldTypeProject, 9)
	c.Set(FieldTypeTask, 10)
	c.Set(FieldTypeIssueType, 11)
	c.Set(FieldTypeTaskStatus, 12)
	c.Set(FieldTypeManhourReportConfig, 27)
	c.Set(FieldTypeGanttChartImportConfig, 28)
	c.Set(FieldTypeTestCaseLibrary, 30)
	c.Set(FieldTypeTestCaseModule, 31)
	c.Set(FieldTypeTestCaseCase, 32)
	c.Set(FieldTypeTestCasePlan, 33)
	c.Set(FieldTypeTaskGanttChartImportConfig, 34)
	c.Set(FieldTypeTaskGanttChart, 35)
	c.Set(FieldTypeStepRange, 40)
	c.Set(FieldTypeStayCount, 41)
	c.Set(FieldTypeStayTime, 42)
	c.Set(FieldTypeProduct, 44)
	c.Set(FieldTypeAppearTime, 45)
	c.Set(FieldTypeProductModule, 46)
	DefaultFieldTypeEnumConverter = c
}

func MustFieldTypeEnum(label string) int {
	enum, _ := DefaultFieldTypeEnumConverter.Enum(label)
	return enum
}
