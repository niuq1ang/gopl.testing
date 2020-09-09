package filter

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

const (
	taskFilterTableColumn = "uuid, team_uuid, field_filter, issue_type_filter"
)

type TaskFilter struct {
	UUID            string `db:"uuid"`
	TeamUUID        string `db:"team_uuid"`
	FieldFilter     int64  `db:"field_filter"`
	IssueTypeFilter int64  `db:"issue_type_filter"`
}

func InsertIntoFilter(src gorp.SqlExecutor, item *TaskFilter) error {
	if item == nil {
		return nil
	}
	sql := fmt.Sprintf("INSERT INTO task_filter(%s) VALUES (?,?,?,?);", taskFilterTableColumn)
	_, err := src.Exec(sql, item.UUID, item.TeamUUID, item.FieldFilter, item.IssueTypeFilter)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func DeleteTaskFilterByUUID(src gorp.SqlExecutor, uuid string) error {
	sql := fmt.Sprintf("DELETE FROM task_filter where uuid=?;")
	_, err := src.Exec(sql, uuid)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}
func UpdateFilter(src gorp.SqlExecutor, uuid string, feild, issueType int64) error {
	sql := fmt.Sprintf("UPDATE task_filter set field_filter=?, issue_type_filter=? WHERE uuid='%s';", uuid)
	_, err := src.Exec(sql, feild, issueType)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func ListTaskFilterByUUID(src gorp.SqlExecutor, uuid string) (*TaskFilter, error) {
	var result *TaskFilter
	sql := fmt.Sprintf("SELECT %s FROM task_filter WHERE uuid=?;", taskFilterTableColumn)
	rows, err := src.Select(&result, sql, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		return rows[0].(*TaskFilter), nil
	}
	return nil, nil
}

func ListAllTaskFilter(src gorp.SqlExecutor) ([]*TaskFilter, error) {
	var results []*TaskFilter
	sql := fmt.Sprintf("SELECT %s FROM task_filter;", taskFilterTableColumn)
	_, err := src.Select(&results, sql)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return results, nil
}
