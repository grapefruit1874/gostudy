package main

import (
	"fmt"
)

// 显卡，cpu，内存
type IntelCPU struct {
	CPU
}
type IntelMemory struct {
	Memory
}
type IntelCard struct {
	Card
}

func (this *IntelCPU) Cal() {
	fmt.Println("IntelCPU")
}
func (this *IntelMemory) Sto() {
	fmt.Println("IntelMemory")
}
func (this *IntelCard) Dis() {
	fmt.Println("IntelCard")
}
