# Golang学习笔记

#### 1.配置GOPATH

步骤1.在电脑上新建一个目录Go_environment

步骤2.在环境变量里，新建一项:GOPATH:D:\Go_environment

步骤3.在D:\Go_environment下新建三个文件夹,分别是:bin、src、pkg

步骤4.把D:\Go_environment\bin这个目录添加到Path这个环境变量中

#### 2.数组

```go
a := []int{x,y,z}
//数组指定位置插入
arr = append(arr[:index], append([]T{value}, arr[index:]...)...)
//切片因为前闭后开
a[:3]  	//在0,1,2三个元素的时候是可以的
```

#### 3.地址

```go
func changeValue(p *int){  //指针类型
	*p = 10  //位置赋值
}

func main(){
	var a int = 1
	changeValue(&a)      //取地址
	fmt.Println("a = ",a)
}
```

#### 4.字典的使用

```go
判断key是否存在
_,status := data["name"]
字典删除元素
delete(hashMap,s[lk])  //前面是字典，后面是字典的键
```

#### 5.if语句

```go
//if条件是可以赋值的，赋值变量的作用域在这个if语句中
if contents,err := ioutil.ReadFile(filename);err!=nil{}
```

#### 6.for循环

```go
//只使用index的简化写法
for index := range s{	
}
```

#### 7.四则运算

对于除号/,它的整数除和小数除是有区别的；整数之间做除法时，只保留整数时部分而舍弃小数部分。不会进行四舍五入等。例：x:=50/9 结果是5

```go
x := 50 / 9 	//其结果是5（整数之间做除法时，只保留整数）

//指数运算，只有math.Pow(base,exponent)
result := math.Pow(2,31)
```

#### 8.位运算

```go
//<< 左移
//取int32中的最大数，直接用左移的方法
MaxInt32  = 1<<31 - 1  
//解释：（1<<31）表示1左移31位，使得符号位为1，其他位为0，再将左移后的结果减一就使得符号位为0，其他位都为1，即为int表示的有符号数的最大整数。
```

#### 9.内置排序库

```go
//如果按升序对切片进行排序，可以使用sort.Ints()、sort.Float64()和sort.Strings()等方法。
nums := []int{6, 4, 5, 2, 8, 1}
sort.Ints(nums)
fmt.Println(nums)

//自定义排序规则
nums := []int{3, 7, 8, 1, 5, 4}
sort.Slice(nums, func(i, j int) bool {
  return nums[i] > nums[j]
})
fmt.Println(nums)
```

#### 10.类型转换

```go
//ASCII转int标准解法
digit := '5'
digitInt := int(digit - '0')

//字符串转int标准解法
str := "5"
number, err := strconv.Atoi(str)

//rune转string标准解法
combination = string(letters[i])
```

#### 11.切片的引用特性

在 Go 语言中，当你传递一个数组作为参数时，实际上是传递了数组的副本，而不是指针。这意味着函数内对传递的数组进行修改不会影响到原始数组。但是在使用切片时，副本和原始切片都指向相同的底层数组。所以尽管传递的是切片而非指针，原切片内容还是会被修改。

```go
//参考在排序过程中，传入的是切片而非指针，但是结构遭到了修改
list := []int{3,5,10,16,7,32,83,23,54,29,96}
sort.Ints(list)
```

#### 12.变量

golang使用驼峰法给变量命名:myVar

#### 13.位运算

```go
&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空（AND NOT）
<<     左移
>>     右移

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
```

