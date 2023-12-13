# Golang学习笔记

#### 1.配置GOPATH

步骤1.在电脑上新建一个目录Go_environment

步骤2.在环境变量里，新建一项:GOPATH:D:\Go_environment

步骤3.在D:\Go_environment下新建三个文件夹,分别是:bin、src、pkg

步骤4.把D:\Go_environment\bin这个目录添加到Path这个环境变量中

#### 2.数组

```
a := []int{x,y,z}
//数组指定位置插入
arr = append(arr[:index], append([]T{value}, arr[index:]...)...)
```

#### 3.地址

```
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

```
判断key是否存在
_,status := data["name"]
字典删除元素
delete(hashMap,s[lk])  //前面是字典，后面是字典的键
```

#### 5.if语句

```
if条件是可以赋值的，赋值变量的作用域在这个if语句中
if contents,err := ioutil.ReadFile(filename);err!=nil{}
```

#### 6.for循环

```
只使用index的简化写法
for index := range s{	
}
```

#### 7.四则运算

对于除号/,它的整数除和小数除是有区别的；整数之间做除法时，只保留整数时部分而舍弃小数部分。不会进行四舍五入等。例：x:=50/9 结果是5

```
x := 50 / 9 	//其结果是5（整数之间做除法时，只保留整数）

//指数运算，只有math.Pow(base,exponent)
result := math.Pow(2,31)
```

#### 8.位运算

```
<< 左移
//取int32中的最大数，直接用左移的方法
MaxInt32  = 1<<31 - 1  
//解释：（1<<31）表示1左移31位，使得符号位为1，其他位为0，再将左移后的结果减一就使得符号位为0，其他位都为1，即为int表示的有符号数的最大整数。
```

#### 9.内置排序库

```
如果按升序对切片进行排序，可以使用sort.Ints()、sort.Float64()和sort.Strings()等方法。
nums := []int{6, 4, 5, 2, 8, 1}
sort.Ints(nums)
fmt.Println(nums)

自定义排序规则
nums := []int{3, 7, 8, 1, 5, 4}
sort.Slice(nums, func(i, j int) bool {
  return nums[i] > nums[j]
})
fmt.Println(nums)
```

