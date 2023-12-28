package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var res []string	//结果集
	dfs(n, n, "", &res) //三个左括号三个右括号
	return res
}

//dfs回溯法，二叉搜索树
func dfs(lindex int, rindex int, path string, res *[]string) {
	if lindex == 0 && rindex == 0 { //递归跳出条件.左右括号全部用完
		*res = append(*res, path) //收集结果path中包含符合条件的字符串结果
		return
	}
	if lindex > 0 {
		dfs(lindex-1, rindex, path+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		dfs(lindex, rindex-1, path+")", res)
	}
}

func main() {
	//输入数据

	//输出内容
	fmt.Println(generateParenthesis(3))

}
