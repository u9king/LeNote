package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		last := target - nums[i]
		if index, status := mMap[last]; status {
			return []int{index, i}
		}
		mMap[nums[i]] = i
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: l1.Val + l2.Val}
	cur := head
	for l1.Next != nil || l2.Next != nil {
		if l1.Next == nil {
			l1 = &ListNode{Val: 0}
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2 = &ListNode{Val: 0}
		} else {
			l2 = l2.Next
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

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(l1, l2))
}
