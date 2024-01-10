package main

import (
	"fmt"
)

//暴力破解
func strStr(haystack string, needle string) int {
	l1,l2 := len(haystack),len(needle)
	if l2 == 0 {
		return 0
	}
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}

func main() {
	//输入数据
	haystack := "mississippi"
	//needle := ""

	fmt.Println(haystack[1:1])

	//输出内容
	//ans := strStr(haystack, needle)
	//fmt.Println(ans)
}
