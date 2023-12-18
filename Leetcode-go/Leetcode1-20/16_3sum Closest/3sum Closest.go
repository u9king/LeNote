package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := nums[n-1] + nums[n-2] + nums[n-3] //取三个最大数
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { //避免重复计算提高效率
			continue
		}
		lk, rk := i+1, n-1
		for lk < rk {
			total := nums[i]+nums[lk]+nums[rk]
			if abs(result-target) > abs(total-target) {
				result = total
			}
			if total > target{
				for rk > lk && nums[rk] == nums[rk-1]{  //重复数字就全部跳过因为rk-1在上一个循环已经算过了
					rk--
				}
				rk--
			} else if total < target{
				for rk > lk && nums[lk] == nums[lk+1]{
					lk++
				}
				lk++
			} else if total == target{
				return target
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	//输入数据
	arr := []int{-1, 2, 1, -4}
	target := 1
	//输出内容
	fmt.Println(threeSumClosest(arr, target))
}
