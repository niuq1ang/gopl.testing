package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r io.Reader
	n int
}

func (this *limitReader) Read(p []byte) (int, error) {
	if this.n < 0 {
		return 0, io.EOF
	}
	size := len(p)
	if size > this.n {
		size = this.n
	}
	n, err := this.r.Read(p[:size])
	if err != nil {
		return 0, err
	}
	this.n -= n
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		r: r,
		n: int(n),
	}
}

func main() {
	r := strings.NewReader("i am a test string !")
	buf := make([]byte, 100)
	n, err := LimitReader(r, 10).Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf), n)
}
