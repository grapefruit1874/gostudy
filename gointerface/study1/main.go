package main

// 开闭原则
// 依赖倒转原则
func main() {
	com1 := NewComputer(&IntelCard{}, &IntelMemory{}, &IntelCPU{})
	com1.Dowork()

	com2 := NewComputer(&NviDiaCard{}, &KingStonMemory{}, &IntelCPU{})
	com2.Dowork()
}
