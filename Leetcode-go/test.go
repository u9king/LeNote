package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	hashMap := map[byte]bool{}
	lk := 0
	maxLength := 0
	for rk := range s {
		for hashMap[s[rk]] {
			hashMap[s[lk]] = false
			lk++
		}
		hashMap[s[rk]] = true
		if maxLength < rk - lk + 1{
			maxLength = rk -lk + 1
		}
	}
	return maxLength
}

func main() {
	//s := "abcabcbb"
	s1 := "abca"

	//输出内容
	fmt.Println(lengthOfLongestSubstring(s1))
}
