package log

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "github.com/bangwork/bang-api/app/utils/log"
	"net/http"
	"net/url"
)

var POSTURL = "'"

func PushLicenseLogCron() string {
	datas := ReadLogFile()
	fmt.Println("datas", datas)
	pushData := Marshal(datas)

	fmt.Println("pushData", pushData)
	// err := uploadLog(POSTURL, pushData)
	// if err == nil {
	// 	DeleteLogFile()
	// 	return
	// }
	// log.Error("license log push err; err=%v", err)
	return pushData

}
func Marshal(items []string) string {
	if len(items) == 0 {
		return ""
	}
	jItems, _ := json.Marshal(items)
	result, _ := EncodeData(string(jItems))
	return result
}

func uploadLog(urlStr, teamDataStr string) error {
	params := url.Values{}
	params.Add("license", teamDataStr)
	_, err := http.PostForm(urlStr, params)
	if err != nil {
		SacnFiles()
	}
	return err
}

func GzipData(data string) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write([]byte(data))
	if err != nil {
		zw.Close()
		return nil, err
	}
	zw.Close()

	return buf.Bytes(), nil
}

func EncodeData(data string) (string, error) {
	if data == "" {
		return "", nil
	}
	gdata, err := GzipData(data)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(gdata)
	return encoded, nil
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

func DecodeData(data string) []string {
	decode, _ := base64.StdEncoding.DecodeString(data)

	gdata, _ := GzipDecode(decode)
	var array []string
	json.Unmarshal(gdata, &array)
	return array
}
