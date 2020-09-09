package payload

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bangwork/ones-ai-api-common/utils/errors"
	noticeModel "github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/notice"
	"github.com/bangwork/ones-plugin/notice_platform/app/models/notice-center/record"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils"
	"github.com/ngaut/log"
	"gopkg.in/redis.v5"
)

const (
	NoticeExpireTime   = 30 * time.Minute
	AllTunnelRecordKey = "all.tunnel.record"
)

func SetNotice(client *redis.Client, notice *noticeModel.Notice) {
	data, _ := json.Marshal(notice)
	err := client.Set(strconv.FormatInt(notice.SeqNo, 10), string(data), NoticeExpireTime).Err()
	if err != nil {
		log.Error(errors.Redis(err))
	}
}

func GetNoticeData(client *redis.Client, seqNo int64) *noticeModel.Notice {
	data, err := client.Get(strconv.FormatInt(seqNo, 10)).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		log.Error(errors.Redis(err))
	}
	notice := new(noticeModel.Notice)
	err = json.Unmarshal([]byte(data), notice)
	if err != nil {
		log.Error(errors.Trace(err))
		return nil
	}
	return notice
}

func RpushRecord(client *redis.Client, records ...*record.Record) {
	if len(records) == 0 {
		return
	}
	for _, record := range records {
		bytes, _ := json.Marshal(record)
		err := client.RPush(AllTunnelRecordKey, string(bytes)).Err()
		if err != nil {
			log.Error(errors.Redis(err))
		}
	}
}

func LPopRecord(client *redis.Client) *record.Record {
	data, err := client.LPop(AllTunnelRecordKey).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		log.Error(errors.Redis(err))
		return nil
	}
	if utils.IsEmptyString(data) {
		return nil
	}
	var result record.Record
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Error(err)
		return nil
	}

	return &result
}
