package main

import (
	"fmt"
	"math"
)

//倍增除数法
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 { //-2^31没有成对的相反数，单独处理这个数就可以
		return math.MaxInt32
	}
	result := 0  //记录被除次数，也就是无符号的商
	var sign int //记录符号
	if dividend*divisor > 0 {
		sign = 1
	} else {
		sign = -1
	}
	dvd, div := abs(dividend), abs(divisor)
	for dvd >= div {
		tmp := div          //初始化除数
		power := 1          //初始化权重
		for tmp<<1 <= dvd { //找寻最大权重的除数，也就是div左移到最大
			tmp <<= 1   //这里不用<<位计算用2*tmp效果是一样的
			power <<= 1 //提升权重
		}
		dvd -= tmp      //减掉计算出的最大数
		result += power //累计权重
	}
	return sign * result
}

//abs 求绝对值函数
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//输入数据
	dividend := 10
	divisor := 3       //0....0011
	//divisor2 := 3 << 1 //0....00110
	//fmt.Println(divisor2)

	//输出内容
	ans := divide(dividend, divisor)
	fmt.Println(ans)

}
