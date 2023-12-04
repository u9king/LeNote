package main

import "math"

func reverse(x int)  (rev int) {
	for x != 0{
		digit := x % 10
		x /= 10
		rev = rev * 10 + digit
		if rev > 1<<31 - 1 || rev < -1<<31{  //1<<31-1就是2的31次方-1，位运算
			return 0
		}
	}
	return rev
}

func reverse_v1(x int) (rev int) {
	for x != 0 {
		if rev > math.MaxInt32/10 || rev < math.MinInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}
