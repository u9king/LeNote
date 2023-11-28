# Golang学习笔记

#### 1.配置GOPATH

步骤1.在电脑上新建一个目录Go_environment

步骤2.在环境变量里，新建一项:GOPATH:D:\Go_environment

步骤3.在D:\Go_environment下新建三个文件夹,分别是:bin、src、pkg

步骤4.把D:\Go_environment\bin这个目录添加到Path这个环境变量中

#### 2.数组

```
a := []int{x,y,z}
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

