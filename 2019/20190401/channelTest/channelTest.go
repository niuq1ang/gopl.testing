package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var Lr LimitRate

//LimitRate 计数器实现限流
type LimitRate struct {
	rate  int           //计数周期内最多允许的请求数
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //计数周期内累计收到的请求数
	lock  sync.Mutex
}

func (l *LimitRate) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count >= l.rate-1 {
		for {
			now := time.Now()
			if now.Sub(l.begin) >= l.cycle {
				//速度允许范围内， 重置计数器
				l.Reset(now)
				return true
			}
			// wait
			time.Sleep(l.cycle - now.Sub(l.begin))
		}
	} else {
		//没有达到速率限制，计数加1
		l.count++
		return true
	}
}

//Set 周期内允许的请求数，cycle计数周期 单位time.Duration,1纳秒。限速为 r/cycle；
func (l *LimitRate) Set(r int, cycle time.Duration) {
	l.rate = r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}

func (l *LimitRate) Reset(t time.Time) {
	l.begin = t
	l.count = 0
}

func main() {
	t1 := time.Now()
	Lr.Set(1, 40*time.Millisecond)
	var lr LimitRate
	lr.Set(1, 400*time.Millisecond)
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup
	dataChan := make(chan int, 10)
	limitChan := make(chan bool, 100)
	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			if Lr.Allow() {
				limitChan <- true
				go func(i int) {
					randomInt := rand.Intn(100)
					time.Sleep(time.Duration(1000) * time.Millisecond)
					dataChan <- i
					fmt.Printf("producter data: %d, thread num: %d,randomInt %d \n", i, runtime.NumGoroutine(), randomInt)
					<-limitChan
				}(i)
			}
		}
	}()

	go func() {
		for data := range dataChan {
			fmt.Println("consumer data:", data, runtime.NumGoroutine(), len(dataChan))
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println(time.Since(t1))
}
