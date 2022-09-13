// 最重要 稳定 对于基本有效的数组是最好用的
// 插入排序：从第二个数开始，与前面的数字进行比较，比前面的小就进行替换，一直到不能替换为止
// 优化：用临时变量记录标记项，去掉swap方法，可以使用位运算进行优化
package main

func main() {
	arr := []int{5, 3, 6, 8, 1, 100, 9, 4, 2}
	isort(arr)
}
func isort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insert(arr, i)
	}
	print(arr)
}
func insert(arr []int, i int) {
	for j := i; j > 0; j-- {
		if arr[j] < arr[j-1] {
			swap(arr, j, j-1)
		}
	}
}
