package main

import (
	"fmt"
	"log"
	"os"
)

const filePath = "bi-tools-status"

func main() {
	file, err := os.OpenFile("testXXX.csv", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Name())

	MakeDirIfNotExist(filePath)
}

//PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//MakeDirIfNotExist 判断文件夹是否存在 不存在则创建
func MakeDirIfNotExist(fileName string) {
	exist, err := PathExists(fileName)
	if err != nil {
		log.Panic(err)
	}
	if exist {
		return
	}
	err = os.Mkdir(fileName, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	return
}
