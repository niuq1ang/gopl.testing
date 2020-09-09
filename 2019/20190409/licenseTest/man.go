package main

import (
	"fmt"
	"github.com/bangwork/mygo/20190409/licenseTest/license"
)

func main() {
	err := license.InitLicense()
	if err != nil {
		fmt.Println(err)
	}
}
