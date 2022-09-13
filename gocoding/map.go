// 实现阻塞(chen)读且并发安全(锁)的map
// go中map如何实现key不存在get操作时等待，直到key存在或超时，保证并发安全
package main

import (
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key/val,如果key读取go挂起，则唤醒，此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key,如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}
