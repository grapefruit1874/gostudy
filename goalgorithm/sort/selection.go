// 选择排序 从第一个值开始比较，找出最小值，两两交换
// 优化练习：找最小值时，顺便找最大值，会少1/2的遍历循环
// 大o分析：时间复杂度 o(n的2次方) 空间复杂度o(1) 不稳（相等的数的相对位置会变化）
// 怎样验证算法的对错【对数器】，大量随机样本，与程序本身的算法进行对比，看看结果是否相同
// 写一个程序证明不稳定
package main

func main() {
	arr := []int{5, 3, 6, 8, 1, 100, 9, 4, 2}
	ssort(arr)
}
func ssort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		mini := findmini(arr, i)
		swap(arr, i, mini)
	}
	print(arr)
}
func findmini(arr []int, i int) int {
	mini := i
	for j := i + 1; j < len(arr); j++ {
		if arr[j] < arr[mini] {
			mini = j
		}
	}
	return mini
}
