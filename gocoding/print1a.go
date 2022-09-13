// 交替打印数字和字母
// 使用两个goroutine,一个打印数字，一个打印字母，最终效果如下：12AB34CD
package main

import (
	"fmt"
	"sync"
)

func main() {
	numch, letterch := make(chan bool), make(chan bool)
	var wg = sync.WaitGroup{}
	go printnum(numch, letterch)
	numch <- true
	wg.Add(1)
	go printletter(numch, letterch, &wg)
	wg.Wait()
}
func printnum(numch, letterch chan bool) {
	i := 1
	for {
		select {
		case <-numch:
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letterch <- true
		}
	}
}

func printletter(numch, letterch chan bool, wg *sync.WaitGroup) {
	i := 'A'
	for {
		select {
		case <-letterch:
			if i >= 'Z' {
				wg.Done()
				return
			}
			fmt.Print(string(i))
			i++
			fmt.Print(string(i))
			i++
			numch <- true
		}
	}
}
