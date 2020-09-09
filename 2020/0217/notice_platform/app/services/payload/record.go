package payload

import (
	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	noticeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/record"
	webhookModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/webhook"
	"github.com/ngaut/log"
)

func GenerateAndSaveRocords(teamUUID string, seqNo int64, notice *noticeModel.Notice) {
	var recordList []*record.Record
	for uuid, v := range webhook.TunnelMap {
		if v.WebHook.TeamUUID == teamUUID && v.Status != webhookModel.WebHookStatusInvalid &&
			v.Status != webhookModel.WebHookStatusDelete {
			rec := &record.Record{
				SeqNo:      seqNo,
				TeamUUID:   teamUUID,
				TunnelUUID: uuid,
				Status:     record.RecordStatusReady,
			}
			recordList = append(recordList, rec)
		}
	}
	if len(recordList) > 0 {
		if err := record.InsertIntoRecord(db.NoticeDBMap, recordList); err != nil {
			log.Error(err)
			return
		}
		RpushRecord(db.RedisClient, recordList...)
	}
}
