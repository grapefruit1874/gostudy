package main

import (
	"fmt"
)

type NviDiaCard struct {
	Card
}

func (this *NviDiaCard) Dis() {
	fmt.Println("NviDiaCard")
}
