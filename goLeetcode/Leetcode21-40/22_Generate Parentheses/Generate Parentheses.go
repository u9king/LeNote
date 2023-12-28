package main

import (
	"fmt"
)

func generateParenthesis(n int) (ans []string) {
	m := n*2  //全二叉树
	path := make([]byte,m)  //生成树
	var dfs func(int,int)  //深度优先搜索算法
	dfs = func(i,open int){
		if i == m{
			ans = append(ans,string(path))
			return
		}
		if open < n{  //可以填左括号
			path[i] = '('
			dfs(i+1,open+1)
		}
		if i-open < open{ //可以填右括号
			path[i] = ')'
			dfs(i+1,open)
		}
	}
	dfs(0,0)
	return
}



func main() {
	//输入数据
	//list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	//list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	//输出内容
	fmt.Println(generateParenthesis(3))

}
