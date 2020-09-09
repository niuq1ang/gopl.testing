package field

// 全局属性
type Field struct {
	UUID               string  `db:"uuid"`
	TeamUUID           string  `db:"team_uuid"`
	Type               int     `db:"type"`
	Name               string  `db:"name"`
	NamePinyin         string  `db:"name_pinyin"`
	Desc               string  `db:"description"`
	CreateTime         int64   `db:"create_time"`
	BuiltIn            bool    `db:"built_in"`
	DefaultValue       *string `db:"default_value"`
	Renderer           int     `db:"renderer"`
	FilterOption       int     `db:"filter_option"`
	SearchOption       int     `db:"search_option"`
	Required           bool    `db:"-"`
	Status             int     `db:"status"`
	StepSettings       *string `db:"step_settings"`
	StaySettings       *string `db:"stay_settings"`
	RelatedType        int     `db:"related_type"`
	RelatedUUID        *string `db:"related_uuid"`
	AppearTimeSettings *string `db:"appear_time_settings"`

	Options []*FieldOption `db:"-"`
}

// 属性选项
type FieldOption struct {
	UUID            string `db:"uuid"`
	TeamUUID        string `db:"team_uuid"`
	FieldUUID       string `db:"field_uuid"`
	Value           string `db:"value"`
	DefaultSelected bool   `db:"default_selected"`
	BackgroundColor string `db:"background_color"`
	Color           string `db:"color"`
	Position        int    `db:"position"`
	Desc            string `db:"desc"`
}

// func (f *Field) Clone() *Field {
// 	n := *f
// 	if n.DefaultValue != nil {
// 		v := *n.DefaultValue
// 		n.DefaultValue = &v
// 	}
// 	if n.StepSettings != nil {
// 		v := *n.StepSettings
// 		n.StepSettings = &v
// 	}
// 	if n.StaySettings != nil {
// 		v := *n.StaySettings
// 		n.StaySettings = &v
// 	}
// 	for i, o := range n.Options {
// 		v := *o
// 		n.Options[i] = &v
// 	}
// 	return &n
// }

// func (f *Field) ApplyConfig(c *FieldConfig) error {
// 	f.DefaultValue = c.DefaultValue
// 	f.Renderer = c.Renderer
// 	f.FilterOption = c.FilterOption
// 	f.SearchOption = c.SearchOption
// 	f.Required = c.Required

// 	if len(f.Options) == 0 {
// 		return nil
// 	} else if len(c.OptionConfigs) == 0 {
// 		// TODO: 之后需要保证必须有 OptionConfigs
// 		return nil
// 		// return errors.New("field config must have at least 1 option")
// 	}

// 	allOps := make(map[string]*FieldOption, len(f.Options))
// 	for _, o := range f.Options {
// 		allOps[o.UUID] = o
// 	}
// 	newOps := make([]*FieldOption, len(c.OptionConfigs))
// 	for i, oc := range c.OptionConfigs {
// 		op := allOps[oc.OptionUUID]
// 		if op == nil {
// 			return errors.CorruptedDataError(
// 				fmt.Sprintf("invalid option '%s'", oc.OptionUUID),
// 				errors.FieldConfig)
// 		}
// 		newOp := *op
// 		newOp.Position = i
// 		newOp.DefaultSelected = oc.DefaultSelected
// 		newOps[i] = &newOp
// 	}
// 	return nil
// }

// func (f *Field) ExtractConfig(projectUUID string, issueTypeUUID string) *FieldConfig {
// 	c := new(FieldConfig)
// 	c.TeamUUID = f.TeamUUID
// 	c.ProjectUUID = projectUUID
// 	c.IssueTypeUUID = issueTypeUUID
// 	c.FieldUUID = f.UUID
// 	c.DefaultValue = f.DefaultValue
// 	c.Renderer = f.Renderer
// 	c.FilterOption = f.FilterOption
// 	c.SearchOption = f.SearchOption
// 	c.Required = f.Required

// 	for _, o := range f.Options {
// 		c.OptionConfigs = append(c.OptionConfigs, &FieldOptionConfig{
// 			FieldOptionConfigKey: FieldOptionConfigKey{
// 				ProjectUUID:   projectUUID,
// 				IssueTypeUUID: issueTypeUUID,
// 				FieldUUID:     f.UUID,
// 				OptionUUID:    o.UUID,
// 			},
// 			Position:        o.Position,
// 			DefaultSelected: o.DefaultSelected,
// 		})
// 	}
// 	return c
// }

// func (f *Field) OptionExists(optionUUID string) bool {
// 	for _, o := range f.Options {
// 		if o.UUID == optionUUID {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (f *Field) IsPriority() bool {
// 	return f.UUID == PriorityFieldUUID
// }

// func (f *Field) IsWatcher() bool {
// 	return f.UUID == WatcherFieldUUID
// }

// func (f *Field) IsSingleOrMultiOption() bool {
// 	switch f.Type {
// 	case FieldTypeMultiOption, FieldTypeOption:
// 		return true
// 	}
// 	return false
// }

// func (f *Field) IsMultiOption() bool {
// 	return IsMultiOptionFieldType(f.Type)
// }

// func (f *Field) IsShouldCalField() bool {
// 	return ShouldCalFieldTypesSet[f.Type]
// }

// func (f *FieldOption) IsPriority() bool {
// 	return f.FieldUUID == PriorityFieldUUID
// }

// // 属性值
// type FieldValue struct {
// 	FieldValueKey
// 	Type        int     `db:"type"`
// 	Value       *string `db:"value"`
// 	NumberValue *int64  `db:"number_value"`
// 	ArrayValue  []string
// }

// type FieldMultiValue struct {
// 	FieldValueKey
// 	Type     int    `db:"type"`
// 	Value    string `db:"value"`
// 	Position int    `db:"position"`
// 	Status   int    `db:"status"`
// }

// // func (v *FieldValue) IsEmpty() bool {
// // 	if IsMultiOptionFieldType(v.Type) {
// // 		return len(v.ArrayValue) <= 0
// // 	} else {
// // 		return v.Value == nil
// // 	}
// // }

// // func (v *FieldValue) IsMultiOption() bool {
// // 	return IsMultiOptionFieldType(v.Type)
// // }

// // func (v *FieldValue) ToMultiValue() []*FieldMultiValue {
// // 	if !v.IsMultiOption() {
// // 		return nil
// // 	}
// // 	if len(v.ArrayValue) == 0 {
// // 		return nil
// // 	}
// // 	multiValues := make([]*FieldMultiValue, 0, len(v.ArrayValue))
// // 	for i, av := range v.ArrayValue {
// // 		multiValues = append(multiValues, &FieldMultiValue{
// // 			FieldValueKey: v.FieldValueKey,
// // 			Type:          v.Type,
// // 			Value:         av,
// // 			Position:      i,
// // 			Status:        FieldMultiValueStatusNormal,
// // 		})
// // 	}
// // 	return multiValues
// // }

// func NewFieldValue(taskUUID string, fieldUUID string, fieldType int, value *string, numberValue *int64, arrayValue []string) *FieldValue {
// 	return &FieldValue{
// 		FieldValueKey: FieldValueKey{
// 			TaskUUID:  taskUUID,
// 			FieldUUID: fieldUUID,
// 		},
// 		Type:        fieldType,
// 		Value:       value,
// 		NumberValue: numberValue,
// 		ArrayValue:  arrayValue,
// 	}
// }

// type FieldValueKey struct {
// 	TaskUUID  string `db:"task_uuid" json:"-"`
// 	FieldUUID string `db:"field_uuid" json:"field_uuid"`
// }

// // 属性配置
// type FieldConfig struct {
// 	FieldConfigKey
// 	TeamUUID          string  `db:"team_uuid"`
// 	DefaultValue      *string `db:"default_value"`
// 	Renderer          int     `db:"renderer"`
// 	FilterOption      int     `db:"filter_option"`
// 	SearchOption      int     `db:"search_option"`
// 	Required          bool    `db:"required"`
// 	Position          int64   `db:"position"`
// 	CanModifyRequired bool    `db:"can_modify_required"`
// 	CanDelete         bool    `db:"can_delete"`

// 	OptionConfigs []*FieldOptionConfig `db:"-"`
// }

// func (conf *FieldConfig) HasDefaultValue() bool {
// 	if conf.DefaultValue == nil || *conf.DefaultValue == "" {
// 		return false
// 	}
// 	return true
// }

// type FieldConfigKey struct {
// 	ProjectUUID   string `db:"project_uuid"`
// 	IssueTypeUUID string `db:"issue_type_uuid"`
// 	FieldUUID     string `db:"field_uuid"`
// }

// // 属性选项配置
// type FieldOptionConfig struct {
// 	FieldOptionConfigKey
// 	Position        int  `db:"position"`
// 	DefaultSelected bool `db:"default_selected"`
// }

// type FieldOptionConfigKey struct {
// 	ProjectUUID   string `db:"project_uuid"`
// 	IssueTypeUUID string `db:"issue_type_uuid"`
// 	FieldUUID     string `db:"field_uuid"`
// 	OptionUUID    string `db:"option_uuid"`
// }

// // 状态属性配置
// type FieldAccessRule struct {
// 	FieldAccessRuleKey
// 	TeamUUID string `db:"team_uuid"`
// 	Access   int    `db:"access"`
// }

// type FieldAccessRuleKey struct {
// 	ProjectUUID     string `db:"project_uuid"`
// 	IssueTypeUUID   string `db:"issue_type_uuid"`
// 	FieldUUID       string `db:"field_uuid"`
// 	StatusUUID      string `db:"status_uuid"`
// 	UserDomainType  int    `db:"user_domain_type"`
// 	UserDomainParam string `db:"user_domain_param"`
// }

// type FieldValueHistory struct {
// 	TeamUUID  string `db:"team_uuid"`
// 	TaskUUID  string `db:"task_uuid"`
// 	FieldUUID string `db:"field_uuid"`
// 	Type      int    `db:"type"`
// 	Value     string `db:"value"`
// 	ValidFrom int64  `db:"valid_from"`
// 	ValidTo   *int64 `db:"valid_to"`
// }

// type HistoryRange struct {
// 	ValidFrom int64  `db:"valid_from"`
// 	ValidTo   *int64 `db:"valid_to"`
// }

// type HistoryRanges []HistoryRange

// func (s HistoryRanges) Len() int           { return len(s) }
// func (s HistoryRanges) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s HistoryRanges) Less(i, j int) bool { return s[i].ValidFrom < s[j].ValidFrom }

// func (s HistoryRanges) Total(nowUnix int64) int64 {
// 	var total int64 = 0
// 	for _, r := range s {
// 		var t int64
// 		if r.ValidTo == nil {
// 			t = nowUnix - r.ValidFrom
// 		} else {
// 			t = *r.ValidTo - r.ValidFrom
// 		}
// 		if t < 0 {
// 			t = 0
// 		}
// 		total += t
// 	}
// 	return total
// }

// type TaskHistoryRange struct {
// 	TaskUUID string `db:"task_uuid"`
// 	HistoryRange
// }

// type FieldStep struct {
// 	FieldType       int    `json:"field_type"`
// 	FieldUUID       string `json:"field_uuid"`
// 	FieldOptionUUID string `json:"field_option_uuid"`
// 	Method          string `json:"method"`
// }

// type FieldStepUnit struct {
// 	StepStart *FieldStep `json:"step_start"`
// 	StepEnd   *FieldStep `json:"step_end"`
// }

// func (fs *FieldStep) Validate() error {
// 	if StepAllowedFieldTypeSet[fs.FieldType] == false {
// 		return errors.InvalidParameterError(errors.FieldStep, errors.Type, errors.NotFound)
// 	}
// 	if len(fs.FieldUUID) != utilsModel.UUIDLen {
// 		return errors.InvalidParameterError(errors.FieldStep, errors.FieldUUID, errors.InvalidValue)
// 	}
// 	if fs.FieldType == FieldTypeTaskStatus {
// 		if len(fs.FieldOptionUUID) != utilsModel.UUIDLen {
// 			return errors.InvalidParameterError(errors.FieldStep, errors.Option, errors.InvalidValue)
// 		}
// 		if FieldStepRangeMethodMap[fs.Method] == false {
// 			return errors.InvalidParameterError(errors.FieldStep, errors.Method, errors.InvalidFieldStepMethod)
// 		}
// 	}
// 	return nil
// }

// type FieldStay struct {
// 	FieldType       int    `json:"field_type"`
// 	FieldUUID       string `json:"field_uuid"`
// 	FieldOptionUUID string `json:"field_option_uuid"`
// }

// func (fs *FieldStay) Validate() error {
// 	if StayAllowedFieldTypeSet[fs.FieldType] == false {
// 		return errors.InvalidParameterError(errors.FieldStay, errors.Type, errors.NotFound)
// 	}
// 	if len(fs.FieldUUID) != utilsModel.UUIDLen {
// 		return errors.InvalidParameterError(errors.FieldStay, errors.UUID, errors.InvalidValue)
// 	}
// 	if len(fs.FieldOptionUUID) != utilsModel.UUIDLen {
// 		return errors.InvalidParameterError(errors.FieldStay, errors.Option, errors.InvalidValue)
// 	}
// 	return nil
// }

// type FieldAppearTime struct {
// 	FieldUUID       string `json:"field_uuid"`
// 	FieldOptionUUID string `json:"field_option_uuid"`
// 	Method          string `json:"method"`
// }

// func (fat *FieldAppearTime) Validate() error {
// 	if len(fat.FieldUUID) != utilsModel.UUIDLen {
// 		return errors.InvalidParameterError(errors.FieldAppearTime, errors.FieldUUID, errors.InvalidValue)
// 	}
// 	if fat.FieldUUID != StatusFieldUUID {
// 		return errors.InvalidParameterError(errors.FieldAppearTime, errors.FieldUUID, errors.InvalidValue)
// 	}
// 	if len(fat.FieldOptionUUID) != utilsModel.UUIDLen {
// 		return errors.InvalidParameterError(errors.FieldAppearTime, errors.UUID, errors.InvalidValue)
// 	}
// 	if !FieldStepRangeMethodMap[fat.Method] {
// 		return errors.InvalidParameterError(errors.FieldAppearTime, errors.Method, errors.InvalidFieldAppearTimeMethod)
// 	}
// 	return nil
// }
