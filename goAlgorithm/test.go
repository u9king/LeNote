package main

import (
	. "./Heap"
	. "./LinkList"
	"container/heap"
	"fmt"
)

func mergeKLists(lists []*ListNode) *ListNode {
	h := Heap{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	//heap.Init(&h) //堆化  //本来就是最小序不需要再排序

	dummy := &ListNode{} // 哨兵节点，作为合并后链表头节点的前一个节点
	cur := dummy
	for len(h) > 0 { // 循环直到堆为空
		node := heap.Pop(&h).(*ListNode) // 剩余节点中的最小节点
		if node.Next != nil {            // 下一个节点不为空
			heap.Push(&h, node.Next) // 下一个节点有可能是最小节点，入堆
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next // 哨兵节点的下一个节点就是新链表的头节点
}



func main() {
	//输入数据
	list1 := CreateList([]int{1,4,5})
	list2 := CreateList([]int{1,3,4})
	list3 := CreateList([]int{2,6})
	lists := []*ListNode{list1, list2, list3}

	//输出内容
	ans := mergeKLists(lists)
	for ans != nil {
		fmt.Print(ans.Val)
		if ans.Next != nil {
			fmt.Print("->")
		}
		ans = ans.Next
	}

}


