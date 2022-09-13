// 冒泡排序：从左到右，两两比较，把小的替换到后面
// 优化？？目前是o(n的2次方)把最好时间复杂度变成o(n)
package main

func main() {
	arr := []int{5, 3, 6, 8, 1, 100, 9, 4, 2}
	bsort(arr)
}
func bsort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		findmax(arr, i)
	}
	print(arr)
}
func findmax(arr []int, i int) {
	for j := 0; j < i; j++ {
		if arr[j] > arr[j+1] {
			swap(arr, j, j+1)
		}
	}
}
