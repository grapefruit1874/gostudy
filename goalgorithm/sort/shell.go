// 希尔排序，改进的插入排序
// 思想：按照间隔排序
package main

func main() {
	arr := []int{5, 3, 6, 8, 1, 100, 9, 4, 2}
	shellsort(arr)
}
func shellsort(arr []int) {

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
