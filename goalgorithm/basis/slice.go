package main

import "fmt"

func main() {
	//切片
	var a = []int{1, 2, 3, 4}
	a = append(a, 5)              // 在后面添加
	b := a[:len(a)-1]             // 删除最后一个,并把最后一个返回给b
	c := append(a[0:2], a[3:]...) // 删除中间元素
	fmt.Println(b, c)

	//栈
	var stack = []int{1, 2}
	// 入栈
	stack = append(stack, 3)
	// 出栈
	b1 := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(stack, b1)

	//队列
	//var queue = []int{1,2}
	// 入队
	//queue = append(0,queue) 没有这种写法
	// 出队
	//r := queue[len(queue)-1]
	//queue = queue[len(queue)-1]
	//fmt.Println(queue, r)

	////双端队列
	//	var queue = []int{1,2}
	//	// 左入队
	//	queue = append(0,queue)//没有这种写法
	//	// 右入队
	//	queue = append(queue,3)
	//	// 右出队
	//	r := queue[len(queue)-1]
	//	queue = queue[:len(queue)-1]
	//	// 左出队
	//	r := queue[0]
	//	queue = queue[1:]

}
