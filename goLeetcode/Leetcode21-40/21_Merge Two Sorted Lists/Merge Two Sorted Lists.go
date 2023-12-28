package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} //哨兵节点
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil { //补充剩余节点
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

func main() {
	//输入数据
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	//输出内容
	fmt.Println(mergeTwoLists(list1, list2))

}
