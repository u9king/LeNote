package main

import "sort"

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

func threeSum_v1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
