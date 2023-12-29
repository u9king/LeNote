# Algorithm and Data Structure

导言：**程序= 数据+结构算法**————尼古拉斯•威茨（Niklaus Wirth）瑞士计算机科学家

## 数据结构

### 1.树

分类：

- 满二叉树：叶子全部铺满的树，一颗深度为k且有2^k^-1个结点的二叉树
- 完全二叉树：叶子按照满二叉树的顺序产生，二叉树有n个结点，其中编号1~n都与满二叉树的编号对应

#### 1.1 二叉树

定义：每个结点最多有两个孩子，分别称为左子树和右子树，可以是空子树

适用范围：解决只有两种情况的问题



#### 1.8 哈夫曼树

适用范围：哈夫曼编码









## 算法

### 1.排序

定义：将一组无序的数据按一定规律排列为有序

适用范围：排行榜

分类：按相同元素排序次序变化与否可分为——稳定排序和不稳定排序

#### 1.1 插入排序

思想：选择一个未排序的值，比较值的大小插入到对应为止，直到对象全部插完为止

时间复杂度：O(n^2^) ——比较+移动的次数

空间复杂度：O(1) ——只需要一个backup变量

```
//InsertSort 插入排序(前小后大)
func InsertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ { 		   //从第二个位置开始排
		backup := arr[i]                       //备份待插入的数据
		j := i - 1                             //从前一个已经排过序的开始找
		for ; j >= 0 && backup < arr[j]; j-- { //从后往前检索插入位置
			arr[j+1] = arr[j] 		//一个一个后移
		}
		arr[j+1] = backup
	}
}
```

演示图片：

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/InsertSort.gif" style="zoom:75%"  align="left">

### 2.遍历

#### 2.1 深度优先搜索(Depth_First Search——DFS)

```
func generateParenthesis(n int) (ans []string) {
	m := n*2  //全二叉树
	path := make([]byte,m)  //生成树
	var dfs func(int,int)  //深度优先搜索算法
	dfs = func(i,open int){
		if i == m{
			ans = append(ans,string(path))
			return
		}
		if open < n{  //可以填左括号
			path[i] = '('
			dfs(i+1,open+1)
		}
		if i-open < open{ //可以填右括号
			path[i] = ')'
			dfs(i+1,open)
		}
	}
	dfs(0,0)
	return
}
```

#### 2.2 广度优先搜索(Breadth_First Search——BFS)





### 3.搜索

#### 3.1 深度优先搜索树





#### 1.线性表

顺序存储结构

顺序存储：把逻辑上相邻的元素存储在物理上相邻的存储但愿中的存储结构

链式存储结构

#### 2.平均查找长度ASL

#### 3.串的模式匹配算法

算法目的:确定主串中所含子串（模式串）第一次出现的位置（定位）

算法应用：搜索引擎、拼写检查、语言翻译、数据压缩

算法种类：

1. BF算法（Brute-Force，暴力破解法，穷举法）
2. KMP算法

#### 4.哈夫曼编码

#### 5.遍历二叉树的三种序

1. DLR——先(根)序遍历
2. LDR——中(根)序遍历
3. LRD——后(根)序遍历

#### 1.数组

定义：按一定格式排列起来的，具有相同类型的数据元素的集合

结构形式：

<img src=".\数据结构picture\数组.png"  style="zoom:50%" align="left">

推广：二维数组 

#### 2.链表







#### 方法

1.指针法

2.双指针法

3.倒序法

