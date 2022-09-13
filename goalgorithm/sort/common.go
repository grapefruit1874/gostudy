package main

import "fmt"

// i与j进行交换
func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func print(arr []int) {
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}
