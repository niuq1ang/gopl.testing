package sequence

import (
	"strconv"
	"time"
)

const SelNoLen = 8

var (
	currentDay string
	currentSeq int64
)

func GetSeqID(teamUUID string, seqNo int64) string {
	return teamUUID + strconv.FormatInt(seqNo, 10)
}

func GetTeamUUIDFromSeqID(seqId string) string {
	return seqId[:8]
}

func GetSeqNoFromSeqID(seqId string) int64 {
	seqNo, _ := strconv.ParseInt(seqId[8:], 10, 64)
	return seqNo
}

func GenSeqNo() int64 {
	currentday, flag := getCurrentDay()
	if flag {
		currentSeq++
	} else {
		currentSeq = 0
	}
	result, _ := strconv.ParseInt(appendZero(currentday, currentSeq), 10, 64)
	return result
}

func getCurrentDay() (string, bool) {
	day := time.Now().Format("20060102")
	if currentDay == day {
		return currentDay, true
	}
	currentDay = day
	return currentDay, false
}

func appendZero(d string, s int64) string {
	tmp := strconv.FormatInt(s, 10)
	for len(tmp) < SelNoLen {
		tmp = "0" + tmp
	}
	return d + tmp
}
