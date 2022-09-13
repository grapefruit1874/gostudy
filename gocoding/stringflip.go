// 翻转字符串
// 思想：以字符串中心为轴，前后翻转
package main

func main() {
	str := "你是一头猪hhhhhq"
	_, stred := flip(str)
	print(stred)
}

func flip(str string) (bool, string) {
	strarr := []rune(str)
	len := len(strarr)
	if len > 5000 {
		return false, str
	}
	for i := 0; i < len/2; i++ {
		strarr[i], strarr[len-1-i] = strarr[len-1-i], strarr[i]
	}
	return true, string(strarr)
}
