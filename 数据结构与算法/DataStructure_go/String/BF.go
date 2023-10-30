package main

import (
	"fmt"
)

func BFMatch(mainStr SString, patternStr SString) int {
	if mainStr.length == 0 || patternStr.length == 0 || mainStr.length < patternStr.length {
		return -1 // 无法匹配
	}

	i := 0
	j := 0

	for i < mainStr.length && j < patternStr.length {
		if mainStr.data[i] == patternStr.data[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j >= patternStr.length {
		return i - patternStr.length // 找到匹配
	}

	return -1 // 未找到匹配
}

func main() {
	mainStr := SString{
		data:   "Hello, World!",
		length: 13,
	}
	patternStr := SString{
		data:   "World",
		length: 5,
	}

	index := BFMatch(mainStr, patternStr)
	if index != -1 {
		fmt.Printf("Pattern found at index: %d\n", index)
	} else {
		fmt.Println("Pattern not found")
	}
}
