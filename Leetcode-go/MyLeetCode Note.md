# MyLeetCode Note

[TOC]

## 第一章	数据结构知识

#### 1.单链表结构

```
type ListNode struct {
      Val int
      Next *ListNode
 }
```

## 第二章	算法专题

#### 1.滑动窗口

核心：左右指针+hashMap，提升嵌套循环问题执行效率。

```
func slidingWindow(){
	hashMap := map[byte]bool{}  //建立hash表
	lk := 0
	maxLength := 0
	for rk := range s {
		for hashMap[s[rk]] {	//从左指针开始剔除元素直至没有与右指针相同的元素存在
			hashMap[s[lk]] = false
			lk++
		}
		hashMap[s[rk]] = true
		//处理题目逻辑
	}
	return maxLength
}
```

<img src="https://img-blog.csdnimg.cn/20af5e0ee7cf4bdfb514822bb23ad062.gif#pic_center">

#### 2.二分查找

核心：左右指针+对半删除条件

```
func binarySearch(arr []int, target int) int{
	start := 0
	end := len(arr) - 1
	for start <= end{
		mid := (start + end) / 2
		if arr[mid] > target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return  -1
}
```



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

链表需要先引入头结点`head`和指针`cur`，然后顺序读取两个链表中的数据，数据不足用0补齐，直到两个链表都没有数据，处理最后一位的进位。

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

### 3.lengthOfLongestSubstring

#### 题目

给定一个字符串 `s` ，请你找出其中不含有重复字符的 **最长子串** 的长度。

#### 示例

```
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

#### 题目大意

字符串找不重复的 **最长子串** 的长度

#### 解题思路

使用滑动窗口，可以优化嵌套循环问题。第一步判断右指针所指内容是否存在于hash表中，存在则删除所有右指针所指元素。第二步记录右指针所指元素。第三步计算最大长度。重复上面过程直至右指针到达数组最大长度

#### 代码

```
func lengthOfLongestSubstring(s string) int {
	hashMap := map[byte]bool{}
	lk := 0
	maxLength := 0
	for rk := range s {
		for hashMap[s[rk]] {	//判断右指针所指元素是否存在
			hashMap[s[lk]] = false	//存在就从左指针开始，删除直至右指针所指内容不存在
			lk++
		}
		hashMap[s[rk]] = true  //记录右指针所指元素
		if maxLength < rk-lk+1 {
			maxLength = rk - lk + 1		//更新最大长度
		}
	}
	return maxLength
}
```

#### 小结

查找类题型，考察滑动窗口。左指针和右指针分别会遍历整个字符串一次，所以时间复杂度为O(n)。

### 4.findMedianSortedArrays

#### 题目

给定两个大小分别为 `m` 和 `n` 的正序（从小到大）数组 `nums1` 和 `nums2`。请你找出并返回这两个正序数组的 **中位数** 。

算法的时间复杂度应该为 `O(log (m+n))` 。

#### 示例

```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

#### 题目大意

两正序数组找中位数，由于时间复杂度是`O(log (m+n))`可能为二分查找

#### 解题思路

- 先解决和为奇数个中位数是midIndex ,和为偶数个中位数是中间两数的平均数
- 再用二分法查找中位数，建立两个指针，循环退出条件是k为1即找到最后一个数字
- 处理边界条件即两数组过界问题

#### 代码

```
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength/2 + 1
		return float64(getKthElement(nums1, nums2, midIndex))
	} else {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex) + getKthElement(nums1, nums2, midIndex + 1)) / 2.0
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0	//二分法寻找中位数
	for {
		if index1 == len(nums1) {	//nums1元素全部用完
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {	//查到最后一个元素，判断输出较小的是中位数
			return min(nums1[index1], nums2[index2])
		}
		half := k/2
		newIndex1 := min(index1+half, len(nums1)) - 1	//避免过界，half和len(nums)都为序号，序号变为下标-1
		newIndex2 := min(index2+half, len(nums2)) - 1	
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]	//提取两数组对应新指针的目标元素
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1	//newIndex1 - index1实际是half/2，但由于有界限限制,所以需要用newIndex代替
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1	//下标变为序号+1
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```

#### 小结

查找类题型，考察二分查找的代码。













### 格式

### X.XXX

#### 题目



#### 示例

```

```

#### 题目大意



#### 解题思路



#### 代码

```

```

#### 小结