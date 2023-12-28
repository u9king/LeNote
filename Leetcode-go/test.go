package main

import "fmt"

//func isValid(s string) bool {
//	hashMap := map[byte]byte{')': '(', '}': '{', ']': '['}
//	var stack []byte
//	var value string
//	value,stack = string(stack[len(stack)-1]),stack[:len(stack)-1]
//
//
//	for i := 0; i < len(s); i++ {
//		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
//			stack = append(stack, s[i])
//		} else if len(stack) > 0 && stack[len(stack)-1] == hashMap[s[i]] {
//			value,stack = string(stack[len(stack)-1]),stack[:len(stack)-1]
//		} else {
//			return false
//		}
//	}
//	return len(stack) == 0
//}

func main() {
	//输入数据
	//s := "abc"
	var stack []byte
	var value byte
	stack = append(stack, []byte("abcd")...)
	value,stack = stack[len(stack)-1],stack[:len(stack)-1]
	fmt.Println(value)
	fmt.Println(stack)

	//输出内容
	//fmt.Println(isValid(s))
}
