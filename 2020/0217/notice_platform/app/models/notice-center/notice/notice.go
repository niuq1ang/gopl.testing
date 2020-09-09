package notice

import (
	"fmt"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/message"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/bang/task"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/ngaut/log"
	"gopkg.in/gorp.v1"
)

const (
	RecordStatusReady    = 1
	RecordStatusSuccess  = 2
	RecordStatusFailures = 3

	maxInsertCount = 200
	bufferSize     = 1000
	noticeColumns  = "seq_no, team_uuid, task_uuid, from_user_uuid, message_uuid, message"
)

type Notice struct {
	SeqNo                 int64  `db:"seq_no"`
	TeamUUID              string `db:"team_uuid"`
	TaskUUID              string `db:"task_uuid"`
	FromUserUUID          string `db:"from_user_uuid"`
	MessageUUID           string `db:"message_uuid"`
	MessagePayLoadContent string `db:"message"`
	Task                  *task.Task
	Message               *message.Message
	MessagePayLoad        *MessagePayLoad
}

type User struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
type MessagePayLoad struct {
	ToUsers     []*User                 `json:"to_users"`
	MessageDesc string                  `json:"message_text"`
	Title       string                  `json:"title"`
	Message     *message.MessagePayload `json:"message"`
	URL         string                  `json:"url"`
}

var noticeChan chan *Notice

func InitNoticehan() {
	noticeChan = make(chan *Notice)
}

func SaveNoticeChan(notice *Notice) {
	size := len(noticeChan)
	if size >= bufferSize {
		log.Warn("notice chan full")
		return
	}
	select {
	case noticeChan <- notice: // SUCCESS
	default:
		log.Warn("no notice send")
	}
	return
}

func SaveNoticeRoutine() {
	if len(noticeChan) == 0 {
		return
	}
	var noticeList []*Notice
	for notice := range noticeChan {
		noticeList = append(noticeList, notice)
	}
	SaveBatch(db.NoticeDBMap, noticeList)
}

func SaveBatch(src gorp.SqlExecutor, items []*Notice) error {
	data := items
	for {
		if len(data) == 0 {
			return nil
		}

		size := maxInsertCount
		if len(data) < size {
			size = len(data)
		}
		group := data[:size]
		data = data[size:]
		err := InsertIntoNotice(src, group)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
}

func InsertIntoNotice(src gorp.SqlExecutor, notices []*Notice) error {
	length := len(notices)
	if length == 0 {
		return nil
	}
	fieldsNumber := 6
	sql := fmt.Sprintf("INSERT INTO notice(%s) VALUES %s;", noticeColumns, utils.SqlInMultiInsertValues(fieldsNumber, length))
	args := make([]interface{}, fieldsNumber*length)
	for i, notice := range notices {
		args[i*fieldsNumber+0] = notice.SeqNo
		args[i*fieldsNumber+1] = notice.TeamUUID
		args[i*fieldsNumber+2] = notice.TaskUUID
		args[i*fieldsNumber+3] = notice.FromUserUUID
		args[i*fieldsNumber+4] = notice.MessageUUID
		args[i*fieldsNumber+5] = notice.MessagePayLoadContent
	}
	_, err := src.Exec(sql, args...)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func SaveNotice(src gorp.SqlExecutor, notice *Notice) error {
	sql := fmt.Sprintf("INSERT INTO notice(%s) VALUES (?,?,?,?,?,?);", noticeColumns)
	_, err := src.Exec(sql, notice.SeqNo, notice.TeamUUID, notice.TaskUUID, notice.FromUserUUID, notice.MessageUUID, notice.MessagePayLoadContent)
	if err != nil {
		return errors.Sql(err)
	}
	return nil
}

func ListNoticeBySeqNo(src gorp.SqlExecutor, seqNo int64) (*Notice, error) {
	var notice *Notice
	sql := fmt.Sprintf("SELECT %s FROM notice WHERE seq_no=?;", noticeColumns)
	rows, err := src.Select(&notice, sql, seqNo)
	if err != nil {
		return nil, errors.Sql(err)
	}
	if len(rows) > 0 {
		return rows[0].(*Notice), nil
	}
	return nil, nil
}
