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
	backTracking(digits, 0)
	return results
}

var combination string

func backTracking(digits string, index int) {
	if index == len(digits) {
		results = append(results, combination)
		return
	}
	digit := digits[index]
	letters := letterMap[int(digit-'0')] //ASCII转int
	for i := 0; i < len(letters); i++ {
		combination += string(letters[i])
		backTracking(digits, index+1)                  //回溯
		combination = combination[:len(combination)-1] //撤销
	}
}

//版本2将combination融合在backTracking函数中
var results_v2 []string

func letterCombinations_v2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	backTracking_v2(digits, 0, "")
	return results
}

func backTracking_v2(digits string, index int, combination string) {
	if index == len(digits) {
		results_v2 = append(results_v2, combination)
		return
	}
	digit := digits[index]
	letters := letterMap[int(digit-'0')] //ASCII转int
	for i := 0; i < len(letters); i++ {
		backTracking_v2(digits, index+1, combination) //回溯
	}
}

func main() {
	//输入数据
	digits := "23"
	//输出内容
	//["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations(digits))
}
