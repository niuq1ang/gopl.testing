package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	FirstName, SecondNames string
)

func main() {
	file, err := os.Open("./xx.txt")
	fmt.Println(file.Name(), err)
	fmt.Println(filepath.Abs("./xx.txt"))
	fmt.Printf("Please enter your full name: ")
	fmt.Scanln(&FirstName, &SecondNames) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	fmt.Printf("Hi %s %s!\n", FirstName, SecondNames)
}
