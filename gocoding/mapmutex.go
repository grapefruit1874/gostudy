//高并发下的锁与map的读写
//在一个高并发的web服务器中，要限制ip的频繁访问，现模拟100个ip同时并发访问服务器，每个ip要重复访问1000次
//每个ip三分钟之内只能访问一次，修改以下代码完成该过程，要求能成功输出success:100

// 1，首先要保证启动的go得到的参数是正确的
// 2，保证map的并发读写
// 3，保证3分钟只能访问一次
// 4，多cpu下修改int极端情况下会存在不同步情况，因此需要原子性的修改int值
// 启动一个协程每分钟检查map中的过期ip for启动协程时传参
package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}
func (o *Ban) visit(ip string) bool {
	//竞态
	o.lock.Lock()
	defer o.lock.Unlock()

	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ban := NewBan(ctx)

	//并发安全
	wg := &sync.WaitGroup{}
	wg.Add(1000 * 100)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) { //参数传值
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1) //原子性
				}
			}(j)
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}
