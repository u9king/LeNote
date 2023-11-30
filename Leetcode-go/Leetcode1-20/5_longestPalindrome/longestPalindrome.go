package main

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	begin := 0
	maxLength := 1
	dp := make([][]bool, n)
	for i := range s {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for L := 2; L < n+1; L++ {
		for i := 0; i < n; i++ {
			j := i + L - 1
			if j >= n {
				break
			}
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == true && maxLength < L {
				maxLength = L
				begin = i
			}
		}
	}
	return s[begin : begin+maxLength]
}

//longestPalindrome_v1 滑动视窗的方法
func longestPalindrome_v1(s string) string {
	lk, rk := 0, 0
	maxString := ""
	for rk < len(s) {
		for rk < len(s) - 1 && s[lk] == s[rk+1] {
			rk++
		}
		for lk-1 >= 0 && rk < len(s) - 1 && s[lk-1] == s[rk+1] {
			lk--
			rk++
		}
		if rk-lk > len(maxString) - 1 {
			maxString = s[lk : rk+1]
		}
		rk = (lk+rk)/2 + 1
		lk = rk
	}
	return maxString
}
