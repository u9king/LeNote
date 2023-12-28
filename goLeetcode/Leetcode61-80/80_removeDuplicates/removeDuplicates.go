package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return 0
	}
	slow,fast := 2,2
	for fast< n {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{0,0,1,1,1,1,2,3,3}
	ans := removeDuplicates(nums)
	fmt.Println(ans)
}
