package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lk, rk := head, head       //建立左右指针(快慢指针)
	for i := 0; i < n+1; i++ { //右指针先走n位
		if rk == nil { //剔除删除第一个元素的问题
			return head.Next
		}
		rk = rk.Next
	}
	for rk != nil { //左右指针同时走直到右指针到底，左指针就离右指针正好差n位
		rk = rk.Next
		lk = lk.Next
	}
	lk.Next = lk.Next.Next
	return head
}

func main() {
	//输入数据
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 5

	//输出内容
	ans := removeNthFromEnd(head, n)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}
