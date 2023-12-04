package main

import "fmt"

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

func fibonacci(n int)int{
	farr := make([]int, n)  //实现memo数组
	farr[0], farr[1] = 1, 1	//设定初始状态

	for i:= 2; i < n; i++ {
		farr[i] = farr[i-1] + farr[i-2]	//状态转移方程
	}
	return farr[n-1]	//返回最终解
}

func main() {
	//输入数据
	x := 123
	fmt.Println(fibonacci(5))
	fmt.Println(-2^31)

	//输出内容
	fmt.Println(reverse(x))
}
