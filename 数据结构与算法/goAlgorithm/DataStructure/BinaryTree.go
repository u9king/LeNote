package DataStructure

//todo
//分析该用链式存储结构还是用顺序存储结构如何实现

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var path []string
	dfs(n, n, "", path)
	return path
}
func dfs(lindex, rindex int, str string, path []string) {
	if lindex == 0 && rindex == 0 {
		path = append(path, str)
		return
	}
	if lindex > 0 {
		dfs(lindex-1, rindex, str+"(", path)
	}
	if rindex > 0 && lindex < rindex {
		dfs(lindex, rindex-1, str+")", path)
	}
}

