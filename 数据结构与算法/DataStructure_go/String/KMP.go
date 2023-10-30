package main

import (
	"fmt"
)

//buildPMT 构建部分匹配表（Partial Match Table）
func buildPMT(patternStr SString) []int {
	pmt := make([]int, patternStr.length)
	pmt[0] = 0
	k := 0

	for i := 1; i < patternStr.length; i++ {
		for k > 0 && patternStr.data[i] != patternStr.data[k] {
			k = pmt[k-1]
		}
		if patternStr.data[i] == patternStr.data[k] {
			k++
		}
		pmt[i] = k
	}

	return pmt
}

//KMPMatch KMP匹配
func KMPMatch(mainStr SString, patternStr SString) int {
	if mainStr.length == 0 || patternStr.length == 0 || mainStr.length < patternStr.length {
		return -1 // 无法匹配
	}

	pmt := buildPMT(patternStr)

	j := 0
	for i := 0; i < mainStr.length; i++ {
		for j > 0 && mainStr.data[i] != patternStr.data[j] {
			j = pmt[j-1]
		}
		if mainStr.data[i] == patternStr.data[j] {
			j++
		}
		if j == patternStr.length {
			return i - j + 1 // 找到匹配
		}
	}

	return -1 // 未找到匹配
}

func main() {
	mainStr := SString{data: "Hello, World!", length: 13}
	patternStr := SString{data: "World", length: 5}

	index := KMPMatch(mainStr, patternStr)
	if index != -1 {
		fmt.Printf("Pattern found at index: %d\n", index)
	} else {
		fmt.Println("Pattern not found")
	}
}
