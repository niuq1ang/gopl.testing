package main

import (
	"fmt"
	"io"
	"os"
)

type CounterWriter struct {
	cw io.Writer
	C  int64
}

func (c *CounterWriter) Write(p []byte) (int, error) {
	c.C = int64(len(p))
	n, err := c.cw.Write(p)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &CounterWriter{cw: w}
	return c, &c.C
}

func main() {
	rw, ilen := CountingWriter(os.Stdout)
	fmt.Println(*ilen)
	rw.Write([]byte("hello world!\n"))
	fmt.Println(*ilen)
}
