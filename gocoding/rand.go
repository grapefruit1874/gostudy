// 实现两个goroutine,其中一个产生五个随机数，另一个读取随机数进行打印
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	out := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Print(i)
		}
	}()
	wg.Wait()
}
