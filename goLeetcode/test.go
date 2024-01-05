package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse2Group(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil { //有两个点
		value := cur.Val       //暂存奇数点的值
		cur.Val = cur.Next.Val //值变为偶数点
		cur.Next.Val = value   //交换值
		cur = cur.Next.Next //移动两格
	}
	return head
}


//分治
func reverseKGroup(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 2 {
		cur := lists[0]
		for cur != nil && cur.Next != nil { //有两个点
			value := cur.Val       //暂存奇数点的值
			cur.Val = cur.Next.Val //值变为偶数点
			cur.Next.Val = value   //交换值
			cur = cur.Next.Next //移动两格
		}
		return head
	}
	num := length / 2
	left := reverseKGroup(lists[:num])
	right := reverseKGroup(lists[num:])
	return reverseKGroup(left, right)
}

func main() {
	//输入数据
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	//输出内容
	ans := reverseKGroup(list1,2)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}

}
