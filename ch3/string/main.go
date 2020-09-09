package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a, b := "asdfasdf", "ffddssaa"
	fmt.Println(isReverse(a, b))

}

func test() {
	s := "我是人啊"
	fmt.Printf("%s\n", string(0x6211))
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r), len(string("我是人啊我是人啊我是人啊我是人啊")))

	fmt.Printf("%c\n", '\uFFFD')
	s = "Hello, 世界"

	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// comma inserts commas in a non-negative decimal integer string.
func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var result string
	for i := len(s); i > 3; i = i - 3 {
		result = "," + s[i-3:] + result
		s = s[:i-3]
	}

	return s + result
}

func isReverse(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]int)
	n := make(map[rune]int)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		n[v]++
	}

	if len(m) != len(n) {
		return false
	}

	for k, v := range m {
		if n[k] != v {
			return false
		}
	}
	return true
}
