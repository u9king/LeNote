package main

import "fmt"

//最优解
//空间换时间
//提问：为什么使用map会优化速度呢？
func twoSum(nums []int, target int) []int {
	resMap := map[int]int{}
	for k,v := range nums{
		last := target - v
		if val,ok := resMap[last];ok{
			return []int{k,val}
		}
		resMap[v] = k
	}
	return nil
}

//方法1：最直接的逻辑
//二层循环
func twoSum_v1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}

func main(){
	//输入数据
	nums := []int{2,7,11,15}
	target := 9
	//调用方法
	ans := twoSum(nums,target)
	//输出结果
	fmt.Println(ans)
}