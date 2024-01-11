package main

import (
	"fmt"
)

//暴力破解
func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
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

//strStr_v1 kmp算法实现
func strStr_v1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := computeNext(needle) //计算next数组
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] //回退
		}
		if haystack[i] == needle[j] {
			j++ //扩大文本串和模式串的公共匹配长度
		}
		if j == len(needle) { //匹配到底
			return i - j + 1 //返回文本串符合条件的初始位置
		}
	}
	return -1
}

func computeNext(patten string) []int {
	next := make([]int, len(patten))   //创建next数组
	j := 0                             // j表示后缀末尾
	next[0] = j                        //初始化
	for i := 1; i < len(patten); i++ { //i表示前缀末尾
		for j > 0 && patten[i] != patten[j] {
			j = next[j-1] // 回退前一位
		}
		if patten[i] == patten[j] {
			j++ //扩大最大公共前后缀长度
		}
		next[i] = j // next[i]是i之前的最长相等前后缀长度
	}
	return next
}

func main() {
	//输入数据
	haystack := "leetcode"
	needle := "leeto"

	//输出内容
	ans := strStr(haystack, needle)
	fmt.Println(ans)
}
