package license

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

func Marshal(data []string) string {
	if len(data) == 0 {
		return ""
	}
	jdata, _ := json.Marshal(data)
	encode, _ := Base64EncodeData(string(jdata))
	return encode
}

func GzipEncode(data string) ([]byte, error) {
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

func Base64EncodeData(data string) (string, error) {
	if data == "" {
		return "", nil
	}
	gdata, err := GzipEncode(data)
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

func Base64DecodeData(data string) []string {
	decode, _ := base64.StdEncoding.DecodeString(data)

	gdata, _ := GzipDecode(decode)
	var array []string
	json.Unmarshal(gdata, &array)
	return array
}
