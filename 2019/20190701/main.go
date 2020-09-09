package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	bytes, err := base64.StdEncoding.DecodeString("yizhisec")
	fmt.Println(string(bytes), err)
}
