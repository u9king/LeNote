package main

import (
	"fmt"
)

//KmpSearch kmp算法实现
func KmpSearch(text string, patten string) int {
	next := computeNext(patten) //计算next数组
	j := 0
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != patten[j] {
			j = next[j-1] //回退
		}
		if text[i] == patten[j] {
			j++ //扩大文本串和模式串的公共匹配长度
		}
		if j == len(patten) { //匹配到底
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
	ans := KmpSearch(haystack, needle)
	fmt.Println(ans)
}
