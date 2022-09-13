// 要求每秒调用一次proc并保证程序不退出
// 不退出就要使用recover()捕获panic错误
// 每秒定时执行
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				proc()
			}()
		}
	}
}

func proc() {
	panic("ok")
}
