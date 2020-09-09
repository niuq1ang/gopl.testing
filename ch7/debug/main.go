package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf1 bytes.Buffer
	fmt.Println(&buf1 == nil)
	var buf *bytes.Buffer
	fmt.Println("1", buf == nil) // true

	if debug {
		buf = new(bytes.Buffer)
		fmt.Println("2", buf == nil) // true
	}
	fmt.Println("3", buf == &buf1) // true

	fmt.Println(buf == nil) // true

	var w io.Writer
	fmt.Println(w == nil) // true
	w = buf
	fmt.Println(w == nil) // fasle

	f(buf)
	if debug {
		fmt.Println(buf.String())
	}
}

func f(out io.Writer) {
	fmt.Println(out == nil)
	if out != nil {
		out.Write([]byte("done!"))
	}
}
