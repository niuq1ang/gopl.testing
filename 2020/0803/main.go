package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "/team/BDfDqJU7/permission_rules/add"
	results := strings.Split(str[1:], "/")
	fmt.Printf("%v", results)
}
