package record

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"gopkg.in/gorp.v1"
)

const (
	RecordStatusReady     = 1
	RecordStatusSuccess   = 2
	RecordStatusFailure   = 3
	RecordStatusResend    = 4
	RecordStatusFilterOut = 5

	recordColumns = "seq_no, team_uuid, tunnel_uuid, status"
)

type Record struct {
	SeqNo      int64  `db:"seq_no" json:"seq_no"`
	TeamUUID   string `db:"team_uuid" json:"team_uuid"`
	TunnelUUID string `db:"tunnel_uuid" json:"tunnel_uuid"`
	Status     int    `db:"status" json:"status"`
}

func InsertIntoRecord(src gorp.SqlExecutor, records []*Record) error {
	length := len(records)
	if length == 0 {
		return nil
	}
	fieldsNumber := 4
	sql := fmt.Sprintf("INSERT INTO send_record(%s) VALUES %s;", recordColumns, utils.SqlInMultiInsertValues(fieldsNumber, length))
	args := make([]interface{}, fieldsNumber*length)
	for i, record := range records {
		args[i*fieldsNumber+0] = record.SeqNo
		args[i*fieldsNumber+1] = record.TeamUUID
		args[i*fieldsNumber+2] = record.TunnelUUID
		args[i*fieldsNumber+3] = record.Status
	}
	_, err := src.Exec(sql, args...)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func UpdateRecordBySelNo(src gorp.SqlExecutor, status int, seqNo int64) (bool, error) {
	sql := fmt.Sprintf("UPDATE send_record SET status=? WHERE seq_no=?;")
	result, err := src.Exec(sql, status, seqNo)
	if err != nil {
		return false, errors.Sql(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return false, errors.Sql(err)
	}
	return count == 1, nil
}

func UpdateRecordsBySelNo(src gorp.SqlExecutor, status int, seqNo []int64) (bool, error) {
	if len(seqNo) == 0 {
		return true, nil
	}
	sql := fmt.Sprintf("UPDATE send_record SET status=? WHERE seq_no in (%s);", utils.SqlPlaceholds(len(seqNo)))
	result, err := src.Exec(sql, status, seqNo[0])
	if err != nil {
		return false, errors.Sql(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return false, errors.Sql(err)
	}
	return count == int64(len(seqNo)), nil
}
