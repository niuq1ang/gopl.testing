package main

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(new(User) == nil)
	fmt.Println(utf8.RuneCountInString("大江东去浪淘尽"))
}
