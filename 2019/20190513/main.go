package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)

	fmt.Println(time.Since(t))
	fmt.Println(time.Now())
}
