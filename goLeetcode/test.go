package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	left := 0
	for _,v := range nums{
		if v != val{
			nums[left] = v
			left++
		}
	}
	return left
}


func main() {
	//输入数据
	arr := []int{0,1,2,2,3,0,4,2}

	//输出内容
	k := removeElement(arr,2)
	fmt.Println(k,arr)
}
