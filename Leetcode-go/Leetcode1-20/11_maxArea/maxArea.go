package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ret := 0
	for left < right {
		h := height[left]
		if h > height[right] {
			h = height[right]
		}
		area := (right - left) * h
		if ret < area {
			ret = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return ret
}
