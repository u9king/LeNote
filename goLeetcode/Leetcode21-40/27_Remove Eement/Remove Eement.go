package main

import (
	"fmt"
)

//单指针法（双指针法变种
func removeElement(nums []int, val int) int {
	lk := 0
	for _,v := range nums{
		if v != val{		//左指针移动条件
			nums[lk] = v	//处理题目逻辑
			lk++			//移动指针
		}
	}
	return lk
}


func main() {
	//输入数据
	arr := []int{0,1,2,2,3,0,4,2}

	//输出内容
	k := removeElement(arr,2)
	fmt.Println(k,arr)
}
