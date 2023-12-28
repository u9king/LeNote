package main

import "fmt"

//最优解
//考察：移动视窗
func lengthOfLongestSubstring(s string) int {
	hashMap := map[byte]bool{}
	lk := 0
	maxLength := 0
	for rk := range s {
		for hashMap[s[rk]] {	//判断右指针所指元素是否存在
			hashMap[s[lk]] = false	//存在就从左指针开始，删除直至右指针所指内容不存在
			lk++
		}
		hashMap[s[rk]] = true  //记录右指针所指元素
		if maxLength < rk-lk+1 {
			maxLength = rk - lk + 1		//更新最大长度
		}
	}
	return maxLength
}

//lengthOfLongestSubstring_v1 被优化不需要从左开始，需要从第一个重复字母开始
func lengthOfLongestSubstring_v1(s string) int {
	hashMap := map[byte]int{}
	rk := -1
	maxLength := 0
	for i := range s {
		if i != 0 {
			delete(hashMap, s[i-1])
		}
		for rk+1 < len(s) && hashMap[s[rk+1]] == 0 {
			hashMap[s[rk+1]]++
			rk++
		}
		if maxLength < rk-i+1 {
			maxLength = rk - i + 1
		}
	}
	return maxLength
}

func main() {
	//输入数据
	s := "pwwkew"
	//调用方法
	ans := lengthOfLongestSubstring(s)
	//输出结果
	fmt.Println(ans)
}
