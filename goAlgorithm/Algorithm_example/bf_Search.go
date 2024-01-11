package main

import "fmt"

//bruteForceSearch 暴力搜索算法，基本思想是在主串上滑动子串，逐一比较字符是否匹配
func bruteForceSearch(text, pattern string) []int {
	var result []int
	n := len(text)
	m := len(pattern)
	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				match = false
				break
			}
		}
		if match {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"
	result := bruteForceSearch(text, pattern)
	fmt.Println(result)

}
