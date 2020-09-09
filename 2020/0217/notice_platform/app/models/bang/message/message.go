package message

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	gorp "gopkg.in/gorp.v1"
)

const (
	messageColumns = "uuid, team_uuid, reference_type, reference_id, from_uuid, to_uuid, send_time, message, type, resource_uuid, subject_type, subject_id, action, object_type, object_name, object_attr, old_value, new_value, ext"
)

func GetMessageByUUID(src gorp.SqlExecutor, uuid string) (*Message, error) {
	message := new(Message)
	sql := fmt.Sprintf("SELECT %s FROM message WHERE uuid=?;", messageColumns)
	rows, err := src.Select(&message, sql, uuid)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		return rows[0].(*Message), nil
	}
	return nil, nil
}

func ListMessageWithLimit(src gorp.SqlExecutor, since int64, limit int) ([]*Message, error) {
	var messages []*Message
	sql := fmt.Sprintf("SELECT %s FROM message WHERE send_time>? ORDER BY send_time ASC limit ?;", messageColumns)
	_, err := src.Select(&messages, sql, since, limit)
	if err != nil {
		return nil, errors.Sql(err)
	}
	return messages, nil
}
