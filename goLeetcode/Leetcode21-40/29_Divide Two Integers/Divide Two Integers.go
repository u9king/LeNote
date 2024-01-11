package main

import (
	"fmt"
	"math"
)

//二分搜索
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	result := 0
	sign := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs(dividend), abs(divisor)
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		result += m
	}
	return sign * result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//输入数据
	dividend := 10
	divisor := 3

	//输出内容
	ans := divide(dividend, divisor)
	fmt.Println(ans)

}
