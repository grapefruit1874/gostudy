// 判断两个给定的字符串排序后是否一致
// 思想：判断s1中的字符在s2中是否存在，又一个不存在就代表不可能相等
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1, s2 := "123", "321"
	check := islike(s1, s2)
	fmt.Print(check)
}

func islike(s1, s2 string) bool {
	sl1, sl2 := len([]rune(s1)), len([]rune(s2))
	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
