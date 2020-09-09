package main

import (
	"bufio"
	"fmt"
	"path"
	"path/filepath"
	"strings"
	// "golang.org/x/text/encoding"
	// "golang.org/x/text/transform"
	"io"
	"os"
	// "unicode"
	"unicode/utf8"
)

func main() {
	var filePath = "./wiki.sql"
	fmt.Println(filepath.Abs(filePath))
	fmt.Println(filepath.Dir(filePath))
	fmt.Println(path.Join(filepath.Dir(filePath), ".", "."+path.Base(filePath)))

	fmt.Println(utf8.ValidString("也让玩家体验非常好"))
	// contents, err := readFile("wiki.sql")
	// if err != nil {
	// 	panic(err)
	// }
	// // err = writeFile("txt1", contents)
	// if err != nil {
	// 	panic(err)
	// }
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileInfo.Mode()
	// tmpFileName := ""
	output := make([]byte, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}

		// if fileType == "project" {
		// 	if "USE `project`;" == strings.TrimSpace(string(line)) {
		// 		output = append(output, []byte(fmt.Sprintf("USE `%s`;\n", options.ProjectDBName))...)
		// 		continue
		// 	}
		// }

		if "wiki" == "wiki" {
			if "USE `wiki`;" == strings.TrimSpace(string(line)) {
				output = append(output, useDBName("wiki")...)
				continue
			}
		}

		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
}

func useDBName(dbName string) []byte {
	return []byte(fmt.Sprintf("USE `%s`;\n", dbName))
}

func writeLine(filePath string, outPut []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
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

// func readFile(filePath string) ([]byte, error) {
// 	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer f.Close()

// 	reader := bufio.NewReader(transform.NewReader(f, encoding.UTF8Validator))
// 	output := make([]byte, 0)
// 	i := 0
// 	for {
// 		i++
// 		line, _, err := reader.ReadLine()
// 		if err != nil {
// 			if err == io.EOF {
// 				return output, nil
// 			}
// 			return nil, err
// 		}
// 		// var result []byte
// 		// _, _, err = encoding.UTF8Validator.Transform(result, line, false)
// 		// if err != nil {
// 		// 	panic(err)
// 		// }
// 		// fmt.Println(string(line))

// 		// fmt.Printf("i=%d, %s\n", i, string(line))
// 		output = append(output, line...)
// 		output = append(output, []byte("\n")...)
// 	}
// }

// func writeToFile(filePath string, outPut []byte) error {
// 	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
// 	defer f.Close()
// 	if err != nil {
// 		return err
// 	}
// 	writer := bufio.NewWriter(transform.NewWriter(f, encoding.UTF8Validator))

// 	for i, r := range string(outPut) {

// 		if i > 35498 && i < 35511 {
// 			fmt.Println(i, unicode.Is(unicode.Han, r), string(r))
// 		}
// 	}
// 	fmt.Println(outPut[35502:35506])
// 	index, err := writer.Write(outPut)
// 	if err != nil {
// 		fmt.Println(index)
// 		return err
// 	}
// 	writer.Flush()
// 	return nil
// }
