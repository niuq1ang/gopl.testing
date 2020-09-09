package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

func main() {
	str := "http1://www.urlregex.com"
	validURL := govalidator.IsURL(str)
	fmt.Printf("%s is a valid URL : %v \n", str, validURL)
}
