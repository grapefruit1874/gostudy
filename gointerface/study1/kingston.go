package main

import (
	"fmt"
)

type KingStonMemory struct {
	Memory
}

func (this *KingStonMemory) Sto() {
	fmt.Println("KingStonMemory")
}
