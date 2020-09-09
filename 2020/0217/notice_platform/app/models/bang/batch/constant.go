package batch

const (
	JobIDMaxLen = 16
)

const (
	JobStatusWaiting     = 1
	JobStatusInProgress  = 2
	JobStatusInterrupted = 3
	JobStatusDone        = 4
)

const (
	StatusShow   = 1
	StatusClosed = 2
)

const (
	RecordStatusNormal = 1
	RecordStatusDelete = 2
)

const (
	RecordContextTypeImportTask            = 1
	RecordContextTypeImportLibraryTestcase = 2
)
const (
	RowContextTypeUpdateIssueType = 1
	RowContextTypeUpdateParent    = 2
)

var JobStatusLabelMap = map[string]int{
	"waiting":     JobStatusWaiting,
	"in_progress": JobStatusInProgress,
	"interrupted": JobStatusInterrupted,
	"done":        JobStatusDone,
}

var JobStatusMap = map[int]string{
	JobStatusWaiting:     "waiting",
	JobStatusInProgress:  "in_progress",
	JobStatusInterrupted: "interrupted",
	JobStatusDone:        "done",
}

var JobStatusFinishedMap = map[int]bool{
	JobStatusInterrupted: true,
	JobStatusDone:        true,
}

func JobStatus(id int) string {
	return JobStatusMap[id]
}

var StatusLabelMap = map[string]int{
	"show":   StatusShow,
	"closed": StatusClosed,
}

var StatusMap = map[int]string{
	StatusShow:   "show",
	StatusClosed: "closed",
}
