package main

import "fmt"

//最优解
//考察：移动视窗
func lengthOfLongestSubstring(s string) int {
	hashMap := map[byte]int{}
	rk := -1
	maxLength := 0
	for i := range s{
		if i != 0{
			delete(hashMap,s[i-1])
		}
		for rk+1<len(s) && hashMap[s[rk+1]] == 0{
			hashMap[s[rk+1]]++
			rk++
		}
		if maxLength < rk - i + 1{
			maxLength = rk -i + 1
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
