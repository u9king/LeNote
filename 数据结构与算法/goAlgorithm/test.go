package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var path []string
	dfs(n, n, "", &path)	//生成一颗n
	return path
}
func dfs(lindex, rindex int, str string, path *[]string) {
	if lindex == 0 && rindex == 0 {
		*path = append(*path, str)
		return
	}
	if lindex > 0 {
		dfs(lindex-1, rindex, str+"(", path)
	}
	if rindex > 0 && lindex < rindex {
		dfs(lindex, rindex-1, str+")", path)
	}
}

func main() {
	//输入数据
	//list := []int{3,5,10,16,7,32,83,23,54,29,96}

	//输出内容
	fmt.Println(generateParenthesis(3))

}
