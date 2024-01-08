package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	lk,rk := 0,0
	for rk < len(nums){		//双指针法的退出逻辑
		if nums[lk] != nums[rk]{	//左指针前进判断
			lk++
			nums[lk] = nums[rk]		//修改数组
		} else {
			rk++	//寻找不同值
		}
	}
	return lk+1
}

func main() {
	//输入数据
	arr := []int{0,0,1,1,1,2,2,3,3,4}

	//输出内容
	k := removeDuplicates(arr)
	fmt.Println(k,arr)
}
