package notice

import (
	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"gopkg.in/gorp.v1"
)

const (
	NoticeStatusUnread = 1
	NoticeStatusRead   = 2
)

type Notice struct {
	UUID              string `db:"uuid"`
	TeamUUID          string `db:"team_uuid"`
	UserUUID          string `db:"user_uuid"`
	TaskUUID          string `db:"task_uuid"`
	MessageUUID       string `db:"message_uuid"`
	Status            int    `db:"status"`
	ServerUpdateStamp int64  `db:"server_update_stamp"`
}

func ListTeamNotices(src gorp.SqlExecutor, since int64, limit int) ([]*Notice, error) {
	var notices []*Notice
	sql := "SELECT team_uuid, user_uuid, task_uuid, message_uuid, status, server_update_stamp FROM notice "
	sql += "WHERE server_update_stamp>? ORDER BY server_update_stamp AESC limit ?;"
	_, err := src.Select(&notices, sql, since, limit)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return notices, nil
}

func ListNoticesByMessageUUID(src gorp.SqlExecutor, messageUUID string) ([]*Notice, error) {
	var notices []*Notice
	sql := "SELECT team_uuid, user_uuid, task_uuid, message_uuid, status, server_update_stamp FROM notice WHERE message_uuid=?;"
	_, err := src.Select(&notices, sql, messageUUID)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return notices, nil
}
