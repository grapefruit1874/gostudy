// 为sync.waitgroup wait实现waittimeout功能
// wait time.timer 都是阻塞的 所以都需要启动协程，难点在于怎么知道哪个阻塞先完成
// 声明一个没有缓冲的chan，谁先完成谁优先向管道写入数据
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}
	if waittimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("time out exit")
	}
	time.Sleep(time.Second * 10)
}

// 超时返回true 否则返回false
func waittimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan bool, 1)
	go time.AfterFunc(timeout, func() {
		ch <- true
	})
	go func() {
		wg.Wait()
		ch <- false
	}()
	return <-ch
}
