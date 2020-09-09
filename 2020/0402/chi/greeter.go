package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
}

// this is exported
var Greeter greeting
