package main

import (
	"crypto/rand"
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	var limit int64 = 1024 * 500
	// we will copy 200 Mb from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	bar := pb.New(1024 * 500).SetUnits(pb.U_BYTES_DEC).SetRefreshRate(time.Millisecond * 10)
	// 显示下载速度
	bar.ShowSpeed = true
	// 显示剩余时间
	bar.ShowTimeLeft = true
	// 显示完成时间
	bar.ShowFinalTime = true
	bar.SetWidth(80)
	bar.Start()
	// bar.NewProxyReader(reader)
	barReader := bar.NewProxyReader(reader)
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	var buf []byte
	buf, err = ioutil.ReadAll(barReader)
	fmt.Println("\n xxxx", err)
	file.Write(buf)
	file.Close()
	bar.Finish()
}
