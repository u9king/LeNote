package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			lk, rk := j+1, len(nums)-1
			for lk < rk {
				if nums[i]+nums[j]+nums[lk]+nums[rk] > target {
					rk--
				} else if nums[i]+nums[j]+nums[lk]+nums[rk] < target {
					lk++
				} else if nums[i]+nums[j]+nums[lk]+nums[rk] == target {
					result = append(result, []int{nums[i], nums[j], nums[lk], nums[rk]})
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
	}
	return result
}

func main() {
	//输入数据
	nums := []int{2, 2, 2, 2}
	target := 8
	//输出内容
	fmt.Println(fourSum(nums, target))
}
