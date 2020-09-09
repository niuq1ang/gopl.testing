package main

import (
	"fmt"
	"strings"
)

func main() {
	taskName := "四川航空（华云）私有部署升级申请"
	nameSlice := strings.Split(taskName, "私有部署升级申请")
	fmt.Printf("%#v", nameSlice)
}
