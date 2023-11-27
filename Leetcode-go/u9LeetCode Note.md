# MyLeetCode Note

[TOC]

## 第一章	数据结构知识

#### 单链表结构

```
type ListNode struct {
      Val int
      Next *ListNode
 }
```

## 第二章	算法专题

## 第三章	题型模板

## 第四章	Leetcode题解

<img src="https://img2.baidu.com/it/u=2643091948,2534409172&fm=253&fmt=auto&app=138&f=JPG?w=952&h=500" style="zoom:50%" align="left">

### 1.twoSum

#### 题目

给定一个整数数组 `nums`和一个整数目标值` target`，请你在该数组中找出**和**为目标值 `target`的那两个整数，并**返回**它们的数组**下标**。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

#### 示例

```
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
```

#### 题目大意

数组内寻找两数之和等于目标值，返回两者下标。

#### 解题思路

顺序扫描数组，如果在字典中找到相加等于`target`的数，直接返回二者下标，没有找到就将这个数和下标记录在字典中。

#### 代码

```
func twoSum(nums []int, target int) []int {
	resMap := map[int]int{}
	for k, v := range nums {
		last := target - v
		if val, ok := resMap[last]; ok {
			return []int{k, val}
		}
		resMap[v] = k  //记录值和下标，留作下次last查询
	}
	return nil
}
```

#### 小结

查找类题型，可用空间换时间增加执行速率。一次for循环，时间复杂度为O(n)。

### 2.addTowNumbers

#### 题目

给你两个非空的**链表**，表示两个非负的整数。他们每位数字都是按照**逆序**的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字0之外，这两个数都不会以0开头。

#### 示例

```
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807
```

#### 题目大意

链表逆序相加，返回表头，注意进位

#### 解题思路

先引入头结点`head`和指针`cur`，然后顺序读取两个链表中的数据，数据不足用0补齐，直到两个链表都没有数据，处理最后一位的进位。

#### 代码

```
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
```

#### 小结

链表类题型。一次for循环，时间复杂度为O(n)。