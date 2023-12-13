package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ansArr [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ { //i是从左到右的驱动器,lk从i+1开始,rk最右边
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		lk := i + 1
		rk := len(nums) - 1
		for lk < rk {
			if nums[i]+nums[lk]+nums[rk] > 0 {
				rk--
			} else if nums[i]+nums[lk]+nums[rk] < 0 {
				lk++
			} else if nums[i]+nums[lk]+nums[rk] == 0 {
				ansArr = append(ansArr, []int{nums[i], nums[lk], nums[rk]})
				lk++
				rk--
				for lk < rk && nums[lk] == nums[lk-1] {
					lk++
				}
				for lk < rk && nums[rk] == nums[rk+1] {
					rk--
				}
			}
		}
	}
	return ansArr
}

func main() {
	//输入数据
	strs := []int{-1, 0, 1, 2, -1, -4}
	arr := []int{1,2,3}
	//输出内容
	fmt.Println(arr[1:3])
	fmt.Println(threeSum(strs))
}
