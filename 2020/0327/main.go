package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	reg1 := regexp.MustCompile(`([A-Z][0-9]{1,6})|[*]{1,}`)
	fmt.Println(reg1.FindString("初步沟通"))
	fmt.Println(msToTime(1585285167704))
}

func msToTime(ms int64) (string, error) {
	tm := time.Unix(ms/1000, 0)
	return tm.Format("2006-01-02 15:04:05.000"), nil
}
