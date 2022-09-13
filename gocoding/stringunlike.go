// 判断字符串中字符是否全都不同
// 请实现一个算法，确定一个字符串的所有字符是否全部不同，要求不允许使用额外的存储结构
// 保证字符串中的字符为ascii字符，字符串长度<=3000
// strings.Count判断某字符的个数；strings.Index/LastIndx判断字符串在另外一个字符的索引位置
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123时的"
	isunlike := isunlike(str)
	fmt.Print(isunlike)
}
func isunlike(str string) bool {
	strl := len([]rune(str))
	if strl > 3000 {
		return false
	}
	for index, val := range str {
		/*	if strings.Count(str, string(val)) > 1 {
			return false
		}*/
		if strings.Index(str, string(val)) != index {
			return false
		}
	}
	return true
}
