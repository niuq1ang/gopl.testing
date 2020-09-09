package license

import (
	"time"
)

type LogReporter struct {
	ID         string    `json:"id"`
	OrgUUID    string    `json:"org_uuid"`
	OrgName    string    `json:"org_name"`
	RecordTime string    `json:"record_time"`
	Msg        string    `json:"msg"`
	LogBatch   *LogBatch `json:"-"`
}

func InitLogReporter(id, orgUUID, orgName string) *LogReporter {
	return &LogReporter{
		ID:       id,
		OrgUUID:  orgUUID,
		OrgName:  orgName,
		LogBatch: InitLogBatch(),
	}
}

func (log *LogReporter) RecordMsg(msg string) error {
	log.RecordTime = time.Now().Format("2006-01-02 15:04:05")
	log.Msg = msg
	return log.LogBatch.AddLog2LogBatch(*log)
}

func (log *LogReporter) GetDefaultMsg(msg string) LogReporter {
	log.RecordTime = time.Now().Format("2006-01-02 15:04:05")
	log.Msg = msg
	return *log
}
