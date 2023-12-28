package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true

	for i := 0; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			switch p[j-1] {
			case '.':
				if i > 0 {
					f[i][j] = f[i-1][j-1]
				}
			case '*':
				if i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					f[i][j] = f[i-1][j] || f[i][j-2]
				} else {

				}
			default:
				if i > 0 && (s[i-1] == p[j-1]) {
					f[i][j] = f[i-1][j-1]
				} else {
					f[i][j] = false
				}
			}
		}
	}
	return f[m][n]
}

func main() {
	//输入数据
	s := "abbbbc"
	p := "ab*d*c"

	//输出内容
	fmt.Println(isMatch(s, p))
}