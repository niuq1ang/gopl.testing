package main

import "fmt"

// slice 之间没有直接比较的方法， slice 只能和 nil 直接比较
func main() {
	var s = "Hello, 世界"
	fmt.Println(len([]rune(s)))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
