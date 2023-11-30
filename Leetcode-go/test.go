package main

import "fmt"

func longestPalindrome(s string) string {
	//滑动窗口
	hashMap := map[byte]bool{}
	lk := 0
	maxString := ""
	for rk := range s {
		for hashMap[s[rk]] {
			hashMap[s[lk]] = false
			lk++
		}
		hashMap[s[rk]] = true
		if rk-lk+1 > len(maxString) && rk-lk+1 == 1 {
			maxString = string(s[lk])
		} else if rk-lk+1 > len(maxString) && rk-lk+1 != 1 {
			maxString = s[lk : rk+1]
		}
	}
	return maxString
}
func longestPalindrome_v1(s string) string {
	lk, rk := 0, 0
	maxString := ""
	for lk < len(s) {
		for rk < len(s) - 1 && s[lk] == s[rk+1] {
			rk++
		}
		for lk-1 >= 0 && rk < len(s) - 1 && s[lk-1] == s[rk+1] {
			lk--
			rk++
		}
		if rk-lk > len(maxString) - 1 {
			maxString = s[lk : rk+1]
		}
		lk = (lk+rk)/2 + 1
		rk = lk
	}
	return maxString
}


func longestPalindrome_v2(s string) string {
	lk, rk := 0, 0
	maxString := ""
	for rk < len(s) {
		for rk < len(s) - 1 && s[lk] == s[rk+1] {
			rk++
		}
		for lk-1 >= 0 && rk < len(s) - 1 && s[lk-1] == s[rk+1] {
			lk--
			rk++
		}
		if rk-lk > len(maxString) - 1 {
			maxString = s[lk : rk+1]
		}
		rk = (lk+rk)/2 + 1
		lk = rk
	}
	return maxString
}

func main() {
	//输入数据
	s := "bacaca"

	//输出内容
	fmt.Println(longestPalindrome_v2(s))
}
