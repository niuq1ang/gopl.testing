package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var dbSpec = "root:123456@tcp(127.0.0.1:3306)/project?parseTime=true&loc=Asia%2FShanghai&charset=utf8mb4"

func main() {
	fmt.Println(getTmpDBSpec("tmp"))
	fmt.Println(resolveCurrentVersion("ones-data_2.93.40373-201911221650.tar"))
}

func getTmpDBSpec(db string) string {
	arrayPrifix := strings.Split(dbSpec, ")/")
	if len(arrayPrifix) == 2 {
		arraySubfix := strings.Split(arrayPrifix[1], "?parseTime")
		if len(arraySubfix) == 2 {
			return fmt.Sprintf("%s)/%s?parseTime%s", arrayPrifix[0], db, arraySubfix[1])
		}
	}
	panic(fmt.Errorf("db_spec error: %s", dbSpec))
}

func resolveCurrentVersion(fileName string) string {
	arrs := strings.Split(fileName, "_")
	if len(arrs) == 2 {
		arrs = strings.Split(arrs[1], "-")
		return arrs[0]
	}
	return ""
}

func test1() {
	filePath := "/Users/niuqiang/Downloads/ones-data/wiki.sql"
	output, err := readFile(filePath, "wiki")
	if err != nil {
		panic(err)
	}
	err = writeToFile(filePath, output)
	if err != nil {
		panic(err)
	}
}

func readFile(filePath, fileType string) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	output := make([]byte, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}

		if fileType == "project" {
			if "USE `project`;" == strings.TrimSpace(string(line)) {
				output = append(output, []byte(fmt.Sprintf("USE `%s`;\n", "test_db_name_project"))...)
				continue
			}
		}

		if fileType == "wiki" {
			if "USE `wiki`;" == strings.TrimSpace(string(line)) {
				output = append(output, []byte(fmt.Sprintf("USE `%s`;\n", "test_db_name_wiki"))...)
				continue
			}
		}

		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
}

func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
