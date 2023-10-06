package main

//左指针法
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
