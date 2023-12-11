package main

//双指针
func maxArea(height []int) int {
	lk,rk := 0,len(height) - 1
	var maxA int
	for lk < rk{
		h := min(height[lk],height[rk])
		area := (rk-lk) * h
		if maxA < area {
			maxA = area
		}
		if height[lk] < height[rk]{
			lk++
		}else{
			rk--
		}
	}
	return maxA
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}


//双循环超时
func maxArea_Diuse(height []int) int {
	maxA := 0
	for lk := 0; lk < len(height); lk++ {  //左指针
		for rk := lk; rk < len(height); rk++ {
			if maxA < (rk - lk) * min(height[rk],height[lk]){
				maxA = (rk - lk) * min(height[rk],height[lk])
			}
		}
	}
	return maxA
}