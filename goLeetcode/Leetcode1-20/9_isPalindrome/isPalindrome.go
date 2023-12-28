package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return origin == rev
}
