package main

import (
	"bufio"
	// "bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func main() {
	var filePath = "./wiki.sql"
	err := replaceFileLine(filePath, replaceDBName)
	if err != nil {
		panic(err)
	}
	// bytes.NewBuffer()
	// writeLine("./wiki_test.sql", lines, 0644)
}

func replaceDBName(line string) (string, bool) {
	if line == "USE `wiki_test`;" {
		return "USE `wiki`;", true
	}
	return "", false
}
func replaceFileLine(filePath string, replaceLine func(line string) (string, bool)) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	defer file.Close()

	const (
		bufferLineSize  = 1 << 20
		bufferTotalSize = 64 << 20
	)

	reader := bufio.NewReaderSize(file, bufferLineSize)
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	tmpFilePath := path.Join(filepath.Dir(abs), "."+path.Base(filePath)+"_tmp")
	bufferLine, bufferTotal := make([]byte, 0), make([]byte, 0, bufferTotalSize)

	copyOrWrite := func(dst, src []byte) ([]byte, error) {
		if len(src) >= bufferTotalSize {
			if err := writeLine(tmpFilePath, src, fileInfo.Mode()); err != nil {
				return nil, err
			}
			return dst, nil
		}

		if len(dst)+len(src) >= bufferTotalSize {
			if err := writeLine(tmpFilePath, dst, fileInfo.Mode()); err != nil {
				return nil, err
			}
			dst = make([]byte, len(src), bufferTotalSize)
			copy(dst, src)
			return dst, nil
		}

		l := len(dst)
		dst = dst[:l+len(src)]
		copy(dst[l:], src)

		return dst, nil
	}

	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		bufferLine = append(bufferLine, line...)
		if isPrefix {
			continue
		}

		result, ok := replaceLine(string(bufferLine))
		if ok {
			bufferLine = []byte(result)
		}

		bufferLine = append(bufferLine, '\n')
		bufferTotal, err = copyOrWrite(bufferTotal, bufferLine)
		if err != nil {
			return err
		}
		bufferLine = bufferLine[0:0]
	}

	if len(bufferTotal) > 0 {
		if err := writeLine(tmpFilePath, bufferTotal, fileInfo.Mode()); err != nil {
			return err
		}
	}
	if err := os.Remove(filePath); err != nil {
		return err
	}
	if err := os.Rename(tmpFilePath, filePath); err != nil {
		return err
	}
	return nil
}

func useDBName(dbName string) []byte {
	return []byte(fmt.Sprintf("USE `%s`;\n", dbName))
}

func writeLine(filePath string, outPut []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, perm)
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
