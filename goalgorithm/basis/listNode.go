package main

import "fmt"

func main() {
	test := List{&Node{
		Data: 0,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: nil,
			},
		},
	}}
	test.Insert(2, 3)
	fmt.Printf("test:%+v", test)
}

type Node struct {
	Data int
	Next *Node
}
type List struct {
	headNode *Node
}

// IsEmpty 判断链表是否为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

// Length 获得链表长度
func (this *List) Length() int {
	cur := this.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// Add 在链表头部添加元素
func (this *List) Add(data int) {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
	return
}

// Append 在链表尾部添加元素
func (this *List) Append(data int) {
	node := &Node{Data: data}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// Insert 在某个位置插入
func (this *List) Insert(index int, data int) {
	if index < 0 {
		this.Add(data)
	} else if index > this.Length() {
		this.Append(data)
	} else {
		count := 0
		cur := this.headNode
		for count < (index - 1) {
			cur = cur.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = cur.Next
		cur.Next = node
	}
}
