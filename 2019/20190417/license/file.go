package license

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bangwork/bang-api/app/utils/log"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const PATH = "/tmp/c4b39a80-a91f-4911-94f1-320114e73bd2"

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//MakeDirIfNotExist 判断文件夹是否存在 不存在则创建
func MakeDirIfNotExist(filePath string) {
	exist := PathExist(filePath)
	if exist {
		return
	}
	os.Mkdir(filePath, os.ModePerm)
	return
}

func GenerateFileName(path string) string {
	MakeDirIfNotExist(path)
	return getFileName(fmt.Sprintf("%s.log", time.Now().Format("20060102")))

}

func WriteLog2File(logs []LogReporter) error {
	fileName := GenerateFileName(PATH)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, item := range logs {
		jdata, err := json.Marshal(item)
		if err != nil {
			return err
		}
		_, err = w.WriteString(AESCBCEncrypter(string(jdata)) + "\n")
		if err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}

func ReadLogFile() []string {
	fileNames := getFileNames()
	fialedItems := make([]string, 0)
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			log.Warn("read file err %v", err)
			continue
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			str := scanner.Text()
			fialedItems = append(fialedItems, str)
		}
	}
	return fialedItems
}

func getFileName(fimeName string) string {
	return fmt.Sprintf("%s/%s", PATH, fimeName)
}

func getFileNames() []string {
	currentLogFileName := GenerateFileName(PATH)
	fileNames := make([]string, 0)
	files, _ := ioutil.ReadDir(PATH)
	for _, file := range files {
		if !file.IsDir() {
			if currentLogFileName == getFileName(file.Name()) {
				continue
			}
			fileNames = append(fileNames, getFileName(file.Name()))
		}
	}
	return fileNames
}

func DeleteLogFile() {
	fileNames := getFileNames()
	for _, fileName := range fileNames {
		os.Remove(fileName)
	}
}
func SacnFiles() {
	fileNames := getFileNames()
	if len(fileNames) <= 10 {
		return
	}
	removeList := make(map[int64]string)
	for _, fileName := range fileNames {
		pathList := strings.Split(fileName, "/")
		if len(pathList) > 3 {
			spList := strings.Split(pathList[len(pathList)-1], ".")
			if len(spList) == 2 {
				i, err := strconv.ParseInt(spList[0], 10, 64)
				if err != nil {
					continue
				}
				removeList[i] = fileName
			}
		}

	}
	if len(removeList) > 10 {
		var keys []int
		for k := range removeList {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)
		for i, m := 0, len(keys); i < m-10; i++ {
			os.Remove(removeList[int64(keys[i])])
		}
	}
}
