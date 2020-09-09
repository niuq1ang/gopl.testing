package issuetype

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

const (
	IssueTypeStatusNormal = 1
	IssueTypeColumns      = "uuid, team_uuid, name, name_pinyin, icon, built_in, default_selected, create_time, status, default_configs, type, detail_type"
)

// 任务类型
type IssueType struct {
	UUID            string `db:"uuid"`
	TeamUUID        string `db:"team_uuid"`
	Name            string `db:"name"`
	NamePinyin      string `db:"name_pinyin"`
	Icon            int    `db:"icon"`
	BuiltIn         bool   `db:"built_in"`
	DefaultSelected bool   `db:"default_selected"`
	CreateTime      int64  `db:"create_time"`
	Status          int    `db:"status"`
	DefaultConfigs  string `db:"default_configs"`
	Type            int    `db:"type"`
	DetailType      int    `db:"detail_type"`
	Index           int
}

func GetIssueType(src gorp.SqlExecutor, teamUUID string, typeUUID string) (*IssueType, error) {
	sql := fmt.Sprintf("SELECT %s FROM issue_type WHERE team_uuid=? AND uuid=? AND status=?;", IssueTypeColumns)
	types := make([]*IssueType, 0)
	_, err := src.Select(&types, sql, teamUUID, typeUUID, IssueTypeStatusNormal)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(types) > 0 {
		return types[0], nil
	} else {
		return nil, nil
	}
}

func ListTeamIssueType(src gorp.SqlExecutor, teamUUID string) ([]*IssueType, error) {
	sql := fmt.Sprintf("SELECT uuid, name, IFNULL(status,0) as status FROM issue_type WHERE team_uuid=? ORDER By create_time ASC;")
	types := make([]*IssueType, 0)
	_, err := src.Select(&types, sql, teamUUID)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return types, nil
}
