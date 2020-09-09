package distributor

import (
	"time"

	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/record"
	webhookModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/filter"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/payload"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/webhook"
	"github.com/ngaut/log"
)

func Sender() {
	for {
		r := payload.LPopRecord(db.RedisClient)
		if r == nil {
			time.Sleep(time.Second * 3)
			continue
		}
		n := payload.GetNoticeData(db.RedisClient, r.SeqNo)
		if n == nil {
			var err error
			n, err = payload.AssembleNotice(r.SeqNo)
			if err != nil {
				log.Error(err)
				continue
			}
		}

		t, ok := webhook.TunnelMap[r.TunnelUUID]
		if ok && t.Status == webhookModel.WebHookStatusValid {
			if filter.DoFilter(n.Message, n.Task, t.FilterType, t.FilterUUID) {
				webhook.DoSend(n, t.LinkType, t.LinkUUID)
			} else {
				_, err := record.UpdateRecordBySelNo(db.NoticeDBMap, record.RecordStatusFilterOut, r.SeqNo)
				if err != nil {
					log.Error(err)
					continue
				}
			}
		} else {
			_, err := record.UpdateRecordBySelNo(db.NoticeDBMap, record.RecordStatusFailure, r.SeqNo)
			if err != nil {
				log.Error(err)
				continue
			}
		}
	}
}
