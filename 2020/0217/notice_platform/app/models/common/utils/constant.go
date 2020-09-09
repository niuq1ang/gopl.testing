package utils

import "github.com/bangwork/ones-plugin/notice_platform/app/models/bang/item"

var (
	FieldTypeOption         = item.MustFieldTypeEnum(item.FieldTypeOption)
	FieldTypeText           = item.MustFieldTypeEnum(item.FieldTypeText)
	FieldTypeInteger        = item.MustFieldTypeEnum(item.FieldTypeInteger)
	FieldTypeFloat          = item.MustFieldTypeEnum(item.FieldTypeFloat)
	FieldTypeDate           = item.MustFieldTypeEnum(item.FieldTypeDate)
	FieldTypeTime           = item.MustFieldTypeEnum(item.FieldTypeTime)
	FieldTypeSprint         = item.MustFieldTypeEnum(item.FieldTypeSprint)
	FieldTypeUser           = item.MustFieldTypeEnum(item.FieldTypeUser)
	FieldTypeProject        = item.MustFieldTypeEnum(item.FieldTypeProject)
	FieldTypeTask           = item.MustFieldTypeEnum(item.FieldTypeTask)
	FieldTypeIssueType      = item.MustFieldTypeEnum(item.FieldTypeIssueType)
	FieldTypeTaskStatus     = item.MustFieldTypeEnum(item.FieldTypeTaskStatus)
	FieldTypeUserList       = item.MustFieldTypeEnum(item.FieldTypeUserList)
	FieldTypeNumber         = item.MustFieldTypeEnum(item.FieldTypeNumber)
	FieldTypeMultiLineText  = item.MustFieldTypeEnum(item.FieldTypeMultiLineText)
	FieldTypeMultiOption    = item.MustFieldTypeEnum(item.FieldTypeMultiOption)
	FieldTypeStatusCategory = item.MustFieldTypeEnum(item.FieldTypeStatusCategory)
	FieldTypeStatus         = item.MustFieldTypeEnum(item.FieldTypeStatus)
	FieldTypeContext        = item.MustFieldTypeEnum(item.FieldTypeContext)
	FieldTypeRichText       = item.MustFieldTypeEnum(item.FieldTypeRichText)
	FieldTypeStepRange      = item.MustFieldTypeEnum(item.FieldTypeStepRange)
	FieldTypeStayCount      = item.MustFieldTypeEnum(item.FieldTypeStayCount)
	FieldTypeStayTime       = item.MustFieldTypeEnum(item.FieldTypeStayTime)
	FieldTypeProduct        = item.MustFieldTypeEnum(item.FieldTypeProduct)
	FieldTypeProductModule  = item.MustFieldTypeEnum(item.FieldTypeProductModule)
	FieldTypeAppearTime     = item.MustFieldTypeEnum(item.FieldTypeAppearTime)
)