package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//swapPairs 标准交换节点的做法，需要知道4个点的交换顺序
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil { //奇数偶数退出条件
		node1 := cur.Next	//需要申请2个变量
		node3 := cur.Next.Next.Next
		cur.Next = cur.Next.Next  //链接node2
		cur.Next.Next = node1
		node1.Next = node3
		cur = cur.Next.Next
	}
	return dummy.Next
}

//swapPairs_v1 Cookbook的写法好叼啊
func swapPairs_v1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur.Next != nil && cur.Next.Next != nil;{
		cur,cur.Next,cur.Next.Next,cur.Next.Next.Next = cur.Next,cur.Next.Next,cur.Next.Next.Next,cur.Next
	}
	return dummy.Next
}

//swapPairs_pb problem版本，原题目不允许直接交换值，但是可以参考一下
func swapPairs_pb(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil { //有两个点
		value := cur.Val       //暂存奇数点的值
		cur.Val = cur.Next.Val //值变为偶数点
		cur.Next.Val = value   //交换值
		cur = cur.Next.Next //移动两格
	}
	return head
}






func main() {
	//输入数据
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	//输出内容
	ans := swapPairs(list1)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}

}
