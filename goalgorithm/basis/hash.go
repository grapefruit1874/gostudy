package main

import "fmt"

func main() {
	//哈希表
	var dic = map[string]int{"a": 12, "b": 23}
	dic["c"] = 34 //添加
	if value, ok := dic["a"]; ok == true {
		fmt.Println(value)
	}
	delete(dic, "a") //删除

	//集合
	hashSet := make(map[string]struct{}) //初始化
	data := []string{"Hello", "World", "213", "3213", "213", "World"}
	for _, v := range data {
		hashSet[v] = struct{}{} //添加
	}
	for k, _ := range hashSet {
		fmt.Println(k) //打印
	}
	delete(hashSet, "Hello") //删除

}
