package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// const path string = `/Users/niuqiang/go_home/src/github.com/bangwork/mygo/doc1`

var _path = flag.String("path", "", "scan forder")

func main() {
	flag.Parse()
	filenames := getAllFileNames(*_path)

	for _, filename := range filenames {
		contents := readAndChange(filename)
		writeFile(filename, contents)
	}
}

func getAllFileNames(_path string) []string {
	fileNames := make([]string, 0)
	// 获取所有文件
	files, _ := ioutil.ReadDir(_path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			name := path.Join(_path, file.Name())
			fileNames = append(fileNames, name)
			fmt.Println(name)
		}
	}

	return fileNames
}

func readAndChange(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	const (
		wikiSaas       = "https://api.ones.ai/wiki/v1/"
		wikiPrivate    = "https://api.ones.ai/wiki/v1//"
		projectSaas    = "https://api.ones.ai/v1"
		projectPrivate = "https://your-host-name/project/api/project"
		addLline       = "    -H 'Referer: https://your-host-name' \\\n"
	)

	var contents []string
	//是否有下一行
	for scanner.Scan() {
		readLine := scanner.Text()

		if strings.Contains(readLine, projectSaas) {
			readLine = strings.Replace(readLine, projectSaas, projectPrivate, -1)
		}

		// if strings.Contains(readLine, "https://testapp.ones.team/wiki/api/wiki") {
		// 	readLine = strings.Replace(readLine, " https://testapp.ones.team/wiki/api/wiki", "https://your-host-name/wiki/api/wiki", -1)
		// }

		// if strings.Contains(readLine, "https://192.168.20.34:443/wiki/api/wiki") {
		// 	readLine = strings.Replace(readLine, "https://192.168.20.34:443/wiki/api/wiki", "https://your-host-name/wiki/api/wiki", -1)
		// }

		// if strings.Contains(readLine, wikiPrivate) {
		// 	readLine = strings.Replace(readLine, wikiPrivate, wikiSaas, -1)
		// }
		// if addLline == readLine {
		// 	continue
		// }
		// if "-H 'cache-control: no-cache'" == strings.TrimSpace(readLine) || "-H 'cache-control: no-cache' \\" == strings.TrimSpace(readLine) {
		// 	readLine = addLline + readLine
		// }

		contents = append(contents, readLine)
	}

	return contents
}

func writeFile(filename string, strs []string) {
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, str := range strs {
		_, err := w.WriteString(str + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	w.Flush()
}
