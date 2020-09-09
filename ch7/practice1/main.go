package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type byteCounter int
type wordCounter int
type lineCounter int

func main() {
	var c byteCounter
	c.Write([]byte("hello this is a line"))
	fmt.Println(c)

	var w wordCounter
	w.Write([]byte("hello this is a line"))
	fmt.Println(w)

	var l lineCounter
	l.Write([]byte("hello \nthis \n is \n a line \n.\n.\n"))
	fmt.Println(l)
}

func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}

func (w *wordCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanWords)
	*w += wordCounter(count)
	return count, nil
}

func (l *lineCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanLines)
	*l += lineCounter(count)
	return count, nil
}

func retCount(p []byte, fn bufio.SplitFunc) (count int) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)

	count = 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
	return
}
