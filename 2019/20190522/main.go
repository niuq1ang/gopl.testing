package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	parentForder   = "/Users/niuqiang/Downloads/ones_master/nginx"
	includeParts   = "project.web.ones.ai_access.log"
	outputFileName = "/Users/niuqiang/go_home/src/github.com/bangwork/ones-monitor/weblogkit/project_web_access.log"
)

func main() {
	file, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	w := bufio.NewWriter(file)
	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, fileName := range getFilesPath() {
		for _, content := range readFileContent(fileName) {
			w.WriteString(content + "\n")
		}
	}
	w.Flush()
	defer file.Close()
}

func getFilesPath() []string {
	fileNames := make([]string, 0)
	files, err := ioutil.ReadDir(parentForder)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			if strings.Contains(file.Name(), includeParts) {
				fileNames = append(fileNames, path.Join(parentForder, file.Name()))
			}
		}
	}
	fmt.Println(len(fileNames))
	return fileNames
}

func readFileContent(fileName string) []string {
	var result []string
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 4<<20)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
