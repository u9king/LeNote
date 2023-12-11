package main

import (
	"fmt"
)

type romanStruct struct {
	value  int
	symbol string
}

func intToRoman(num int) string {
	romanMap := []romanStruct{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}
	result := ""
	for _, v := range romanMap {
		for num >= v.value {
			result += v.symbol
			num -= v.value
		}
		if num == 0 {
			break
		}
	}
	return result
}

func romanToInt(s string) int {
	romanMap := map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400,
		"C": 100, "XC": 90, "L": 50, "XL": 40,
		"X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
	result := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if i < n-1 {
			double := s[i : i+2]
			if cur, ok := romanMap[double]; ok {
				result += cur
				i++
				continue
			}
		}
		single := s[i : i+1]
		result += romanMap[single]
	}
	return result
}

func main() {
	//输入数据
	s := "MCMXCIV"
	//输出内容
	fmt.Println(romanToInt(s))
}
