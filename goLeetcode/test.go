package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//需要三个变量来实现反转pre,cur,nxt分别代表反转后当前节点需要指向的节点，当前节点，反转前当前节点的Next结点
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{Next:head}
	var pre *ListNode
	cur := dummy.Next
	for cur != nil {
		nxt := cur.Next //nxt保存反转前下一个节点
		cur.Next = pre  //反转cur节点的方向
		pre = cur       //向下一个节点移动
		cur = nxt       //向下一个节点移动
	}
	dummy.Next.Next = cur
	dummy.Next=pre
	return dummy.Next
}

func main() {
	//输入数据
	list1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	//输出内容
	ans := reverseList(list1)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}

}
