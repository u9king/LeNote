package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		var length int
		if len(prefix) < len(strs[i]) {
			length = len(prefix)
		} else {
			length = len(strs[i])
		}
		index := 0
		for index < length && prefix[index] == strs[i][index] {
			index++
		}
		prefix = prefix[:index]
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}
