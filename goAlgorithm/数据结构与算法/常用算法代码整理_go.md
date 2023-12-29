### 常用代码整理

#### 1.递增数组原地去重

```
func removeDuplicates(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	slow := 1
	for fast := 1;fast < n ; fast++{
		if nums[fast] != nums[fast-1]{
			nums[slow] = nums[fast]
			slow++
		}
	}
	return nums[:slow]
}
```



### 算法代码整理

#### 1.快慢指针法 

作用：数组去重，双指针

前提：？有序，？单数组

```
func removeDuplicates(nums []int) int {
    n := len(nums)
    if n <= 0 {
        return n
    }
    slow, fast := 1, 1
    for fast < n {
        if nums[slow-1] != nums[fast] {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}
```





### 方法

1.指针法

2.双指针法

3.倒序法