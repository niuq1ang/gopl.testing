package field

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	utilsModel "github.com/bangwork/ones-plugin/notice_platform/app/models/common/utils"
)

const (
	FieldUUIDLen         = 8
	ImportantFieldMaxLen = 5

	FieldStatusNormal = 1

	FieldValueStatusNormal  = 1
	FieldValueStatusDeleted = 2

	FieldMultiValueStatusNormal  = 1
	FieldMultiValueStatusDeleted = 2
	FieldMultiValueStatusStash   = 3

	FieldOptionStatusNormal  = 1
	FieldOptionStatusDeleted = 2

	AccessReadMask  = 0x1
	AccessWriteMask = 0x2

	FieldRelatedTypeTaskStatus = 1 // task_status
)

var (
	FieldTypeOption         = utilsModel.FieldTypeOption
	FieldTypeText           = utilsModel.FieldTypeText
	FieldTypeInteger        = utilsModel.FieldTypeInteger
	FieldTypeFloat          = utilsModel.FieldTypeFloat
	FieldTypeDate           = utilsModel.FieldTypeDate
	FieldTypeTime           = utilsModel.FieldTypeTime
	FieldTypeSprint         = utilsModel.FieldTypeSprint
	FieldTypeUser           = utilsModel.FieldTypeUser
	FieldTypeProject        = utilsModel.FieldTypeProject
	FieldTypeTask           = utilsModel.FieldTypeTask
	FieldTypeIssueType      = utilsModel.FieldTypeIssueType
	FieldTypeTaskStatus     = utilsModel.FieldTypeTaskStatus
	FieldTypeUserList       = utilsModel.FieldTypeUserList
	FieldTypeNumber         = utilsModel.FieldTypeNumber
	FieldTypeMultiLineText  = utilsModel.FieldTypeMultiLineText
	FieldTypeMultiOption    = utilsModel.FieldTypeMultiOption
	FieldTypeStatusCategory = utilsModel.FieldTypeStatusCategory
	FieldTypeRichText       = utilsModel.FieldTypeRichText
	FieldTypeStepRange      = utilsModel.FieldTypeStepRange
	FieldTypeStayCount      = utilsModel.FieldTypeStayCount
	FieldTypeStayTime       = utilsModel.FieldTypeStayTime
	FieldTypeProduct        = utilsModel.FieldTypeProduct
	FieldTypeProductModule  = utilsModel.FieldTypeProductModule
	FieldTypeAppearTime     = utilsModel.FieldTypeAppearTime
)

const (
	FieldStepRangeMethodFirstAt = "first_at"
	FieldStepRangeMethodLastAt  = "last_at"
)

var (
	FieldStepRangeMethodMap = map[string]bool{
		FieldStepRangeMethodFirstAt: true,
		FieldStepRangeMethodLastAt:  true,
	}
)

var (
	FieldTypes = []int{
		FieldTypeOption,
		FieldTypeText,
		FieldTypeInteger,
		FieldTypeFloat,
		FieldTypeDate,
		FieldTypeTime,
		FieldTypeSprint,
		FieldTypeUser,
		FieldTypeProject,
		FieldTypeTask,
		FieldTypeIssueType,
		FieldTypeTaskStatus,
		FieldTypeUserList,
		FieldTypeMultiLineText,
	}

	UsableFieldTypes = map[int]bool{
		FieldTypeOption:        true,
		FieldTypeText:          true,
		FieldTypeInteger:       true,
		FieldTypeFloat:         true,
		FieldTypeDate:          true,
		FieldTypeTime:          true,
		FieldTypeSprint:        true,
		FieldTypeUser:          true,
		FieldTypeMultiLineText: true,
		FieldTypeUserList:      true,
		FieldTypeMultiOption:   true,
		FieldTypeTaskStatus:    true,
		FieldTypeStepRange:     true,
		FieldTypeStayCount:     true,
		FieldTypeStayTime:      true,
		FieldTypeProduct:       true,
		FieldTypeAppearTime:    true,
		FieldTypeProductModule: true,
	}

	// 需要被计算的属性类型
	ShouldCalFieldTypesSet = map[int]bool{
		FieldTypeStepRange:  true,
		FieldTypeStayCount:  true,
		FieldTypeStayTime:   true,
		FieldTypeAppearTime: true,
	}

	// 「间隔时间」 属性支持的类型
	StepAllowedFieldTypeSet = map[int]bool{
		FieldTypeTime:       true,
		FieldTypeDate:       true,
		FieldTypeTaskStatus: true,
	}

	// [属性停留时间或次数] 属性支持的类型
	StayAllowedFieldTypeSet = map[int]bool{
		FieldTypeOption:     true,
		FieldTypeTaskStatus: true,
		FieldTypeSprint:     true,
		FieldTypeUser:       true,
		FieldTypeIssueType:  true,
	}

	// [出现时间] 属性支持的类型
	AppearTimeAllowedFieldTypeSet = map[int]bool{
		FieldTypeTaskStatus: true,
	}

	// 保存「数字类型」的属性
	NumberFieldTypeSet = map[int]bool{
		FieldTypeDate:       true,
		FieldTypeFloat:      true,
		FieldTypeInteger:    true,
		FieldTypeNumber:     true,
		FieldTypeTime:       true,
		FieldTypeStayCount:  true,
		FieldTypeStayTime:   true,
		FieldTypeAppearTime: true,
	}

	// 默认的优先级配置
	DefaultPriorityStyles = []FieldOption{
		{
			Value:           "最高",
			DefaultSelected: false,
			BackgroundColor: "#e63422",
			Desc:            "将会直接阻碍进程的问题。",
			Color:           "#FFFFFF",
			Position:        0,
		},
		{
			Value:           "较高",
			DefaultSelected: false,
			BackgroundColor: "#ffe9e2",
			Desc:            "可能会阻碍进程的严重问题。",
			Color:           "#ff6a39",
			Position:        1,
		},
		{
			Value:           "普通",
			DefaultSelected: true,
			BackgroundColor: "#e0ecfb",
			Desc:            "有可能影响进程的潜在问题。",
			Color:           "#307fe2",
			Position:        2,
		},
		{
			Value:           "较低",
			DefaultSelected: false,
			BackgroundColor: "#d9f4ed",
			Desc:            "容易解决的小问题。",
			Color:           "#00b388",
			Position:        3,
		},
		{
			Value:           "最低",
			DefaultSelected: false,
			BackgroundColor: "#EFEFEF",
			Desc:            "对进展影响甚微，微不足道的问题。",
			Color:           "#909090",
			Position:        4,
		},
	}
)

func IsFieldTypeValid(f int) bool {
	for _, t := range FieldTypes {
		if t == f {
			return true
		}
	}
	return false
}

func IsOptionFieldType(f int) bool {
	return f == FieldTypeOption
}

func IsUserFieldType(f int) bool {
	return f == FieldTypeUser
}

func IsMultiOptionFieldType(f int) bool {
	switch f {
	case FieldTypeUserList, FieldTypeMultiOption, FieldTypeProduct, FieldTypeProductModule:
		return true
	}
	return false
}

func IsStepRangeFieldType(f int) bool {
	return f == FieldTypeStepRange
}

func IsStayFieldType(f int) bool {
	return f == FieldTypeStayTime || f == FieldTypeStayCount
}

func IsAppearTimeFieldType(f int) bool {
	return f == FieldTypeAppearTime
}

// 类型是否能用在自定义属性上
func IsFieldTypeUsable(t int) bool {
	return UsableFieldTypes[t]
}

func FieldTypeName(t int) (name string, err error) {
	switch t {
	case FieldTypeOption:
		name = "option"
	case FieldTypeText:
		name = "text"
	case FieldTypeInteger:
		name = "integer"
	case FieldTypeFloat:
		name = "float"
	case FieldTypeDate:
		name = "date"
	case FieldTypeTime:
		name = "time"
	case FieldTypeSprint:
		name = "sprint"
	case FieldTypeUser:
		name = "user"
	case FieldTypeProject:
		name = "project"
	case FieldTypeTask:
		name = "task"
	case FieldTypeIssueType:
		name = "issue_type"
	case FieldTypeTaskStatus:
		name = "task_status"
	case FieldTypeUserList:
		name = "user_list"
	case FieldTypeMultiLineText:
		name = "multi_line_text"
	case FieldTypeMultiOption:
		name = "multi_option"
	case FieldTypeRichText:
		name = "rich_text"
	case FieldTypeStepRange:
		name = "step_range"
	case FieldTypeStayCount:
		name = "stay_count"
	case FieldTypeStayTime:
		name = "stay_time"
	case FieldTypeAppearTime:
		name = "appear_time"
	default:
		err = errors.InvalidEnumError(t, errors.Field, errors.Type)
	}
	return
}

func FieldTypeByName(name string) (t int, err error) {
	switch name {
	case "option":
		t = FieldTypeOption
	case "text":
		t = FieldTypeText
	case "integer":
		t = FieldTypeInteger
	case "float":
		t = FieldTypeFloat
	case "date":
		t = FieldTypeDate
	case "time":
		t = FieldTypeTime
	case "sprint":
		t = FieldTypeSprint
	case "user":
		t = FieldTypeUser
	case "project":
		t = FieldTypeProject
	case "task":
		t = FieldTypeTask
	case "issue_type":
		t = FieldTypeIssueType
	case "task_status":
		t = FieldTypeTaskStatus
	case "user_list":
		t = FieldTypeUserList
	case "multi_line_text":
		t = FieldTypeMultiLineText
	case "multi_option":
		t = FieldTypeMultiOption
	case "rich_text":
		t = FieldTypeRichText
	case "step_range":
		t = FieldTypeStepRange
	case "stay_count":
		t = FieldTypeStayCount
	case "stay_time":
		t = FieldTypeStayTime
	case "appear_time":
		t = FieldTypeAppearTime
	default:
		err = errors.InvalidEnumError(name, errors.Field, errors.Type)
	}
	return
}
