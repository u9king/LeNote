package main

//给你两个按非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

//伪代码
//从大数组的最后一个位置开始，
//依次选取两个数组中大的数，从第一个数组的尾巴开始往头放
//疑问：如何返回，是否需要用到指针

func merge(nums1 []int, m int, nums2 []int, n int) {
	for rk := m + n; m > 0 && n > 0; rk-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[rk-1] = nums2[n-1]
			n--
		} else {
			nums1[rk-1] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	//nums1 := []int{2,0}
	//m := 1
	//nums2 := []int{1}
	//n := 1

	merge(nums1, m, nums2, n)
}
