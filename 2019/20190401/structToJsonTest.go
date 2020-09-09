package main

import (
	"fmt"
	"sync"
)

type User struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"age,omitempty"`
}

func main() {

	mychannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		mychannel <- i
	}
	close(mychannel) //关闭管道
	fmt.Println("data lenght: ", len(mychannel))
	for v := range mychannel { //循环管道
		fmt.Println(v)
	}
	fmt.Printf("data lenght:  %d", len(mychannel))

	datachan := make(chan int, 2) //通讯数据管道
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go Productor(datachan, i, &wg) //生产数据
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datachan, &wg) //消费数据
		wg.Add(1)
	}
	wg.Wait()
}

func Productor(mychan chan int, data int, wait *sync.WaitGroup) {
	mychan <- data
	fmt.Println("product data：", data)
	wait.Done()
}
func Consumer(mychan chan int, wait *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer data：", a)
	wait.Done()
}
