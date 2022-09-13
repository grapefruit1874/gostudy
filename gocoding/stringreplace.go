// 替换字符串
// 把空格全部变成%20
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := " v vvv vv "
	check, stred := replace(str)
	fmt.Print(check, stred)
}

func replace(str string) (bool, string) {
	strl := len([]rune(str))
	if strl > 1000 {
		return false, str
	}
	for _, val := range str {
		if string(val) != " " && unicode.IsLetter(val) == false {
			return false, str
		}
	}
	str = strings.Replace(str, " ", "%20", -1)
	return true, str
}
