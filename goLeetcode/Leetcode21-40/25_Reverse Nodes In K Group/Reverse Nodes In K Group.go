package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head ; cur != nil;cur = cur.Next{
		n++
	}
	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre *ListNode
	cur := p0.Next
	for ; n>= k;n -= k{
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
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
