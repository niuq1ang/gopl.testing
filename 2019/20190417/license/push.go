package license

import (
	"encoding/json"
	"github.com/bangwork/bang-api/app/utils/log"
	"net/http"
	"net/url"
)

var POSTURL = "'"

func PushLicenseLogCron() {
	datas := ReadLogFile()
	pushData := Marshal(datas)
	if pushData == "" {
		msg := reporter.GetDefaultMsg("---")
		jmsg, _ := json.Marshal(msg)
		pushData = Marshal([]string{string(jmsg)})
	}
	err := UploadLog(POSTURL, pushData)
	if err == nil {
		DeleteLogFile()
		return
	}
	log.Error("license log push err; err=%v", err)
}

func UploadLog(urlStr, teamDataStr string) error {
	params := url.Values{}
	params.Add("license", teamDataStr)
	_, err := http.PostForm(urlStr, params)
	if err != nil {
		SacnFiles()
	}
	return err
}
