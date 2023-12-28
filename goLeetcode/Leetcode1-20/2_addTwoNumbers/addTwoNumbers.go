package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//最优解
//缺0补齐，进位计算，最后一次进位是否添加
//考察链表能力
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: l1.Val + l2.Val}
	cur := head
	for l1.Next != nil || l2.Next != nil {
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = &ListNode{Val: 0}
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = &ListNode{Val: 0}
		}
		cur.Next = &ListNode{Val: l1.Val + l2.Val + cur.Val/10}
		cur.Val = cur.Val % 10
		cur = cur.Next
	}
	if cur.Val >= 10 {
		cur.Next = &ListNode{Val: cur.Val / 10}
		cur.Val = cur.Val % 10
	}
	return head
}

//有问题
//好像链表的读取是定死的，只能用上一种方法
func addTwoNumbers_v1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: l1.Val + l2.Val}
	carry := head.Val / 10
	head.Val = head.Val % 10
	cur := head
	for l1.Next != nil || l2.Next != nil {
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = &ListNode{Val: 0}
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = &ListNode{Val: 0}
		}
		cur.Next = &ListNode{Val: l1.Val + l2.Val + carry}
		carry = cur.Next.Val / 10
		cur.Next.Val = cur.Next.Val % 10
	}
	if cur.Val >= 10 {
		cur.Next = &ListNode{Val: cur.Val / 10}
		cur.Val = cur.Val % 10
	}
	return head
}

func main() {
	//输入数据
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	//调用方法
	ans := addTwoNumbers_v1(l1, l2)
	//输出结果
	fmt.Println(ans.Val, ans.Next.Val, ans.Next.Next.Val)
}
