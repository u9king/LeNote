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
func slidingWindow(nums []int){
	//视情况使用存储结构，也可以使用hash表
	hashMap := map[byte]bool{}  	//建立hash表
	lk,rk := 0,0					//建立左右指针
	maxLength := 0 					//所求业务
	for rk < len(nums) {			//如果用len(nums)-1，前面rk就得从-1开始并且需要先执行rk++,rk带动，右侧判断窗口
		for hashMap[s[rk]] {		//窗口缩小条件
			hashMap[s[lk]] = false	//移出窗口元素
			lk++					//缩小窗口
		}
		hashMap[s[rk]] = true		//更新窗口
        //处理题目逻辑
		rk++						//扩大窗口
	}
	return maxLength				//返回业务
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

#### 3.动态规划

核心：状态转移方程+memory数组

步骤：

- 设计状态
- 写出状态转移方程
- 设定初始状态
- 执行状态转移
- 返回最终解

```
func fibonacci(n int)int{
	farr := make([]int, n)  //实现memo数组
	farr[0], farr[1] = 1, 1	//设定初始状态
	for i:= 2; i < n; i++ {
		farr[i] = farr[i-1] + farr[i-2]	//状态转移方程
	}
	return farr[n-1]	//返回最终解
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

### 5.longestPalindrome

#### 题目

给你一个字符串 `s`，找到 `s` 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

#### 示例

```
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
```

#### 题目大意

字符串 `s`找到最长的回文子串

#### 解题思路

- 扩大窗口，如果有相同元素则必互为互文，直接扩大即可
- 拿着当前的lk那个元素，使用中心扩散法，找到回文数的边界
- 记录最大符合条件的回文数
- 重置rk,lk至此回文数中心点，并后移一位开始重新搜索

#### 代码

```
//滑动窗口法
func longestPalindrome(s string) string {
	lk, rk := 0, 0
	maxString := ""
	for rk < len(s) {
		for rk < len(s) - 1 && s[lk] == s[rk+1] {
			rk++
		}
		for lk-1 >= 0 && rk < len(s) - 1 && s[lk-1] == s[rk+1] {
			lk--
			rk++
		}
		if rk-lk > len(maxString) - 1 {
			maxString = s[lk : rk+1]
		}
		rk = (lk+rk)/2 + 1	//重置为中心点后的那个元素
		lk = rk
	}
	return maxString
}
```

#### 小结

搜索类题型，考察滑动窗口和中心扩散法的联合使用。

### 6.ZigZag Conversion

#### 题目

将一个给定字符串 `s` 根据给定的行数 `numRows` ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 `"PAYPALISHIRING"` 行数为 `3` 时，排列如下：

```
P   A   H   N
A P L S I I G
Y   I   R
```

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：`"PAHNAPLSIIGYIR"`。

请你实现这个将字符串进行指定行数变换的函数：

```
string convert(string s, int numRows);
```

#### 示例

```
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
```

#### 题目大意

字符串 `s` 根据给定的行数 `numRows` ，以从上往下、从左到右进行 Z 字形排列。

#### 解题思路

- 如果把Z字变换按照填入顺序进行连线，这是一个向下每次一步，到`numRows`时再反向向上每次一步的运动轨迹
- 根据这个逻辑规律可以将其抽象成带方向的step
- 创建一个数组存储`numRows`个字符串，先用初始值`""`来填充
- 用`step`和`rowNum`控制循环向同行字符串添加元素
- 复原字符串

#### 代码

```
func convert(s string, numRows int) string {
	rows := make([]string, numRows)
	rowNum := 0	//行编号
	step := 1 //带方向的步长
	for _, v := range s {
		rows[rowNum] += string(v) //同行字符串添加
		if numRows-1 == 0 { //处理只有一行的问题
			step = 0
		} else if rowNum == numRows-1 {
			step = -1
		} else if rowNum == 0 {
			step = 1
		}
		rowNum += step
	}
	//复原字符串
	var res string
	for _, row := range rows {
		res += row
	}
	return res
}
```

#### 小结

二维数组类型题目，也可以根据逻辑原理进行一维简化。

### 7.Reverse Integer

#### 题目

给你一个32位的有符号整数`x`，返回将`x`中的数字部分反转后的结果。

如果反转后整数超过32位的有符号整数的范围[$-2^{31}$,$2^{31}-1$]，就返回0

#### 示例

```
输入：x = 123
输出：321
```

#### 题目大意

整数反转和Int32最大限制

#### 解题思路

- 整数反转，使用经典方法一位一位取，倒着放回去就行
- int32限制可以用math自带的包，也可以用位运算的方法实现

#### 代码

```
func reverse(x int)  (rev int) {
	for x != 0{
		digit := x % 10
		x /= 10
		rev = rev * 10 + digit
		if rev > 1<<31 - 1 || rev < -1<<31{  //1<<31-1就是2的31次方-1，位运算
			return 0
		}
	}
	return rev
}
```

#### 小结

数字处理类题型，要对数位敏感，并且多训练一些位运算对后续有很大的帮助。

### 8.String to Integer (atoi)

#### 题目

请你来实现一个 `myAtoi(string s)` 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 `atoi` 函数）。

函数 `myAtoi(string s)` 的算法如下：

1. 读入字符串并丢弃无用的前导空格
2. 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
3. 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
4. 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 `0` 。必要时更改符号（从步骤 2 开始）。
5. 如果整数数超过 32 位有符号整数范围 `[−231, 231 − 1]` ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 `−231` 的整数应该被固定为 `−231` ，大于 `231 − 1` 的整数应该被固定为 `231 − 1` 。
6. 返回整数作为最终结果。

#### 示例

```
输入：s = "   -42"
输出：-42
解释：
第 1 步："   -42"（读入前导空格，但忽视掉）
            ^
第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
             ^
第 3 步："   -42"（读入 "42"）
               ^
解析得到整数 -42 。
由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42 。
```

#### 题目大意

实现字符串转数字

#### 解题思路

- 单循环先排除跳过空格
- 保存符号
- 判断当前value是否在字符0到9之间
- 判断相乘后的数字是否大于int32的边界
- 数字与符号相乘返回结果

#### 代码

```
func myAtoi(s string) int {
	result := 0
	sign := 1	//符号位
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	for _, value := range s {
		if value == ' '{	//空就跳过
			continue
		}
		if value == '-'{	//符号保存下来
			sign = -1
		}
		if value >= '0' && value <= '9'{
			result = result*10 + int(value - '0')	//进位并存储新数用ASCII码表示
			if result * sign > MaxInt32{return MaxInt32}	//判断是否大于int32最大数
			if result * sign < MinInt32{return MinInt32}	//判断是否小于int32最小数
		} else{
			break
		}
	}
	return result * sign
}

//待解决问题在哪里？说不符合数据：s := "00000-42a1234" 要求为：“0”，结果为“-42”
func myAtoi_waitTest(s string) int {
	result := 0
	sign := 1        //符号位
	hasSign := false //出现例子：“+-12”
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	for _, value := range s {
		if value == ' ' || value == '0' && result == 0{ //出现空或者左边的0就跳过
			continue
		}
		if value == '-' && hasSign == false { //符号保存下来
			sign = -1
			hasSign = true
			continue
		} else if value == '+' && hasSign == false { //出现例子："+1"
			hasSign = true
			continue
		}

		if value >= '0' && value <= '9' {
			result = result*10 + int(value-'0') //进位并存储新数用ASCII码表示
			if result*sign > MaxInt32 {         //判断是否大于int32最大数
				return MaxInt32
			}
			if result*sign < MinInt32 { //判断是否小于int32最小数
				return MinInt32
			}
		} else {
			break
		}
	}
	return result * sign
}
```

#### 小结

标准库实现类题型，涉及较多编程原理的知识。













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
