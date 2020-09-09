package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02", "2019-08-01")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Unix())
}
