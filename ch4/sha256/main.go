package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

func main() {
	fmt.Println('A', int('A'))
	niu := sha256.Sum256([]byte("niu"))
	qiang := sha256.Sum256([]byte("qiang"))
	var niuInt []int
	for _, v := range niu {
		niuInt = append(niuInt, int(v))
	}
	var qiangInt []int
	for _, v := range qiang {
		qiangInt = append(qiangInt, int(v))
	}

	fmt.Printf("%x\nlen=%d\n", niu, len(niu))
	fmt.Printf("%x\nlen=%d\n", niuInt, len(niuInt))
	sort.Ints(niuInt)

	fmt.Printf("%x\nlen=%d\n", qiang, len(qiang))
	fmt.Printf("%x\nlen=%d\n", qiangInt, len(qiangInt))
	sort.Ints(qiangInt)

	fmt.Printf("%x\nlen=%d\n", niuInt, len(niuInt))
	fmt.Printf("%x\nlen=%d\n", qiangInt, len(qiangInt))

	fmt.Printf("%d", diff(niu[:], qiang[:]))
}

func diff(b1, b2 []byte) int {
	m := make(map[byte]int)
	n := make(map[byte]int)

	for _, v := range b1 {
		m[v]++
	}

	for _, v := range b2 {
		n[v]++
	}

	relult := max(len(b1), len(b2))
	equalCount := 0
	for k, v := range m {
		v2 := n[k]
		equalCount += min(v, v2)
	}

	return relult - equalCount
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
