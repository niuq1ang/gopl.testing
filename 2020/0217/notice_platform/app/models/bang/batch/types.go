package batch

// max length 64
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

// MapJobTypesIgnoreCountCompute 数量无需通过计算的JobType
// 不在这个map中的JobType, 数量通过 Successful 和 Unsuccessful 计算得出
var MapJobTypesIgnoreCountComputing = map[string]bool{
	JobTypeImportTasks: true,
}

type BatchTask struct {
	UUID              string  `db:"uuid" json:"uuid"`
	TeamUUID          string  `db:"team_uuid" json:"team_uuid"`
	Owner             string  `db:"owner" json:"owner"`
	JobID             string  `db:"job_id" json:"job_id"`
	JobType           string  `db:"job_type" json:"job_type"`
	JobStatus         int     `db:"job_status" json:"job_status"`
	BatchType         string  `db:"batch_type" json:"batch_type"`
	Status            int     `db:"status" json:"status"`
	StartTime         int64   `db:"start_time" json:"start_time"`
	EndTime           int64   `db:"end_time" json:"end_time"`
	Payload           string  `db:"payload" json:"payload"`
	Extra             *string `db:"extra" json:"extra"`
	Successful        *string `db:"successful" json:"successful"`
	Unsuccessful      *string `db:"unsuccessful" json:"unsuccessful"`
	Unprocessed       *string `db:"unprocessed" json:"unprocessed"`
	SuccessfulCount   int     `db:"successful_count" json:"successful_count"`
	UnsuccessfulCount int     `db:"unsuccessful_count" json:"unsuccessful_count"`
	UnprocessedCount  int     `db:"unprocessed_count" json:"unprocessed_count"`
}

func (b *BatchTask) IsFinished() bool {
	return JobStatusFinishedMap[b.JobStatus]
}

func (b *BatchTask) IsInterrupted() bool {
	return b.JobStatus == JobStatusInterrupted
}

func (b *BatchTask) IsInProgress() bool {
	return b.JobStatus == JobStatusInProgress
}

func (b *BatchTask) ProcessDone(total int) bool {
	processedCount := b.SuccessfulCount + b.UnsuccessfulCount + b.UnprocessedCount
	return processedCount == total
}

type BatchTaskRecord struct {
	BatchTaskUUID string `db:"batch_task_uuid"`
	TeamUUID      string `db:"team_uuid"`
	Number        int64  `db:"number"`
	ContextType   int    `db:"context_type"`
	ContextParam1 string `db:"context_param_1"`
	ContextParam2 string `db:"context_param_2"`
	CreateTime    int64  `db:"create_time"`
	Extra         string `db:"extra"`
	Status        int    `db:"status"`
}

type BatchTaskRow struct {
	BatchTaskUUID string `db:"batch_task_uuid"`
	TeamUUID      string `db:"team_uuid"`
	Number        int64  `db:"number"`
	ContextType   int    `db:"context_type"`
	ContextParam1 string `db:"context_param_1"`
	ContextParam2 string `db:"context_param_2"`
	CreateTime    int64  `db:"create_time"`
	Extra         string `db:"extra"`
}
