package main

import "fmt"

//双指针 快慢指针法
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main(){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	ans := removeDuplicates(nums)
	fmt.Println(ans)
}
