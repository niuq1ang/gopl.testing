package task

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

const (
	taskColumns = "uuid, team_uuid, create_time, priority, owner, assign, tags, sprint_uuid, " +
		"project_uuid, issue_type_uuid, sub_issue_type_uuid, status_uuid, deadline, status, summary, `desc`, desc_rich, " +
		"server_update_stamp, complete_time, open_time, score, parent_uuid, path, position, number, assess_manhour, related_count, remaining_manhour"
)

// task
const (
	TaskStatusIncomplete = 1
	TaskStatusCompleted  = 2
	TaskStatusDeleted    = 3

	// FreeTeamTaskMaxCount = 600 已弃用,改为按组织做限量
)

type Task struct {
	UUID                  string  `db:"uuid" json:"uuid" dbmode:"r"`
	TeamUUID              string  `db:"team_uuid" json:"-"`
	CreateTime            int64   `db:"create_time" json:"create_time" dbmode:"r"`
	Priority              string  `db:"priority" json:"priority"`
	Owner                 string  `db:"owner" json:"owner"`
	Assign                string  `db:"assign" json:"assign"`
	Tags                  string  `db:"tags" json:"tags"`
	SprintUUID            *string `db:"sprint_uuid" json:"sprint_uuid"`
	ProjectUUID           string  `db:"project_uuid" json:"project_uuid"`
	StandardIssueTypeUUID string  `db:"issue_type_uuid" json:"issue_type_uuid"`
	SubIssueTypeUUID      string  `db:"sub_issue_type_uuid" json:"sub_issue_type_uuid"`
	StatusUUID            string  `db:"status_uuid" json:"status_uuid"`
	Deadline              *int64  `db:"deadline" json:"deadline"`
	Status                int     `db:"status" json:"status"`
	Summary               string  `db:"summary" json:"summary"`
	Desc                  string  `db:"desc" json:"desc"`
	DescRich              string  `db:"desc_rich" json:"desc_rich"`
	Score                 *int64  `db:"score" json:"score"`
	ParentUUID            string  `db:"parent_uuid" json:"parent_uuid"`
	Path                  string  `db:"path" json:"path"`
	Position              int64   `db:"position" json:"position"`
	ServerUpdateStamp     int64   `db:"server_update_stamp" json:"server_update_stamp"`
	CompleteTime          int64   `db:"complete_time" json:"complete_time"`
	OpenTime              int64   `db:"open_time" json:"open_time"`
	Number                int     `db:"number" json:"number"`
	AssessManhour         *int64  `db:"assess_manhour" json:"assess_manhour"`
	RelatedCount          int     `db:"related_count" json:"related_count"`
	RemainingManhour      *int64  `db:"remaining_manhour" json:"remaining_manhour"`
}

// 获取逻辑上的任务类型， 如果是子任务返回子任务类型， 如果是根节点父任务，返回任务类型
func (t *Task) GetIssueTypeUUID() string {
	if t.ParentUUID != "" {
		return t.SubIssueTypeUUID
	}
	return t.StandardIssueTypeUUID
}

func GetTaskIfExists(src gorp.SqlExecutor, uuid string) (*Task, error) {
	sql := "SELECT %s FROM task WHERE uuid=?;"
	sql = fmt.Sprintf(sql, taskColumns)
	task := new(Task)
	rows, err := src.Select(task, sql, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		task = rows[0].(*Task)
		return task, nil
	} else {
		return nil, nil
	}
}

func GetTask(src gorp.SqlExecutor, uuid string) (*Task, error) {
	sql := "SELECT %s FROM task WHERE uuid=? AND status!=?;"
	sql = fmt.Sprintf(sql, taskColumns)
	task := new(Task)
	rows, err := src.Select(task, sql, uuid, TaskStatusDeleted)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		task = rows[0].(*Task)
		return task, nil
	} else {
		return nil, nil
	}
}
