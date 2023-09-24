package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	// 奇数
	if totalLength%2 == 1 {
		midIndex := totalLength/2 + 1
		return float64(getKthElement(nums1, nums2, midIndex))
	} else {
		// 偶数
		midIndex1, midIndex2 := totalLength/2, totalLength/2+1
		return float64(getKthElement(nums1, nums2, midIndex1)+getKthElement(nums1, nums2, midIndex2)) / 2.0
	}
}

// 二分法找到第k项
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		// 如果nums1数组用完
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k = k + index1 - newIndex1 - 1
			index1 = newIndex1 + 1
		} else {
			k = k + index2 - newIndex2 - 1
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
