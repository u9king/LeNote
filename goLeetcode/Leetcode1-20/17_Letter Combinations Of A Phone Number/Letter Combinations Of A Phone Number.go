package main

import (
	"fmt"
)

var results []string

var letterMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	results = []string{} //不加这句话本地对，力扣错；加上就都对了，为什么？
	backTracking(digits, 0, "")
	return results
}

func backTracking(digits string, index int, combination string) {
	//combination作为局部变量避免全局变量的多次修改导致错误
	if index == len(digits) {
		results = append(results, combination)
		return
	}
	digit := digits[index]
	letters := letterMap[int(digit-'0')] //ASCII转int
	for i := 0; i < len(letters); i++ {
		backTracking(digits, index+1, combination+string(letters[i])) //回溯
	}
}

func main() {
	//输入数据
	digits := "2"
	//输出内容
	fmt.Println(letterCombinations(digits))
}
