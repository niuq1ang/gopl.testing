package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/prd_log", server) //
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	// fmt.Println("license=", Base64DecodeData(r.Form["license"]))
	for _, val := range r.Form["license"] {
		array := Base64DecodeData(val)
		for _, arr := range array {
			fmt.Println(AESCBCDecrypter(arr))
		}

	}
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

const AESKEY = "6368616e676520746869732070617373"

var key []byte

func GetAESKey() []byte {
	if len(key) > 0 {
		return key
	}
	key, _ = hex.DecodeString(AESKEY)
	return key
}
func AESCBCDecrypter(cipher_text string) string {
	key := GetAESKey()
	ciphertext, _ := hex.DecodeString(cipher_text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return cipher_text
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)
	return fmt.Sprintf("%s", ciphertext)
}

func AESCBCEncrypter(plain_text string) string {
	key := GetAESKey()
	plaintext := []byte(plain_text)
	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return plain_text
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return plain_text
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return fmt.Sprintf("%x", ciphertext)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
