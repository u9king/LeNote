package main

import "fmt"

//递增数组原地去重
//双指针法
func removeDuplicates(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	slow := 1
	for fast := 1;fast < n ; fast++{
		if nums[fast] != nums[fast-1]{
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}

func main(){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	ans := removeDuplicates(nums)
	fmt.Println(ans)
}