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
