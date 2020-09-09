package batch

type Action string
type Type string

func (a Action) String() string {
	return string(a)
}

func (t Type) String() string {
	return string(t)
}

const (
	BatchTypeDefault          Type = "none"
	BatchTypeUpdateTaskBatch  Type = "batch"
	BatchTypeUpdateTaskSingle Type = "single"
)

const (
	BatchActionModifyIssueType   Action = JobTypeModifyIssueType
	BatchActionStdToSubIssueType Action = JobTypeStdToSubIssueType
	BatchActionStdToStdIssueType Action = JobTypeStdToStdIssueType
	BatchActionSubToStdIssueType Action = JobTypeSubToStdIssueType
	BatchActionSubToSubIssueType Action = JobTypeSubToSubIssueType
	BatchActionSubModifyParent   Action = JobTypeSubModifyParent
)

var BatchUpdateTaskTypeSet = map[Type]bool{
	BatchTypeDefault:          true,
	BatchTypeUpdateTaskBatch:  true,
	BatchTypeUpdateTaskSingle: true,
}

var BatchActionSet = map[Action]bool{
	BatchActionModifyIssueType:   true,
	BatchActionStdToStdIssueType: true,
	BatchActionSubToSubIssueType: true,

	BatchActionStdToSubIssueType: true,
	BatchActionSubToStdIssueType: true,
	BatchActionSubModifyParent:   true,
}
