package main

import (
	"fmt"
	"time"
	// "github.com/bangwork/mygo/20190417/log"
)

func main() {
	fmt.Println(time.Now().AddDate(0, 0, 1).Unix())
	tm := time.Unix(1270915200, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
	// logReporter := log.InitLogReporter("ones.ai", "xxxx", "复临科技")
	// encodeData, err1 := log.EncodeData("")
	// fmt.Println("EncodeData: ", encodeData, err1)
	// fmt.Println("marshal: ", log.Marshal([]string{}))
	// err := logReporter.RecordMsg("licenseDatas.OrgIdentity, org.UUID, org.Name")
	// err = logReporter.RecordMsg("organization scale exceed the limit")
	// fmt.Println(err)
	// data := log.PushLicenseLogCron()
	// array := log.DecodeData(data)
	// for _, a := range array {
	// 	fmt.Println("decrtpt", log.AESCBCDecrypter(a))
	// }
	// log.SacnFiles()
	// logReporter.LogBatch.Close()
}
