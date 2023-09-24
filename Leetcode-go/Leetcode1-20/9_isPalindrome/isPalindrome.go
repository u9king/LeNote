package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	redirect := 0
	for x != 0 {
		redirect = redirect*10 + x%10
		x /= 10
	}
	return origin == redirect
}
