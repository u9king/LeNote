package main

import (
	"fmt"
)

func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1

	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}

	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	for ; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		result = result*10 + int(s[i]-'0')
		if sign*result < MinInt32 {
			return MinInt32
		}
		if sign*result > MaxInt32 {
			return MaxInt32
		}
	}
	return sign * result
}

func main() {
	//输入数据
	s := "00000-42a1234"

	//输出内容
	fmt.Println(myAtoi(s))
}
