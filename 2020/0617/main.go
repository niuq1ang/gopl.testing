package main

import "fmt"

func main() {
	needReplaceTaskMap := make(map[string][]string)
	needReplaceTaskMap["123"] = nil
	v, ok := needReplaceTaskMap["123"]
	fmt.Println(v, ok)
	fmt.Println(len(needReplaceTaskMap["123"]))
}
