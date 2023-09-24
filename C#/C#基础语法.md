# C#基础语法

#### 1.十种技术编写Hello world

Console

```
namespace ConsoleHelloWorld
{
    internal class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello, World!");
        }
    }
}
```

#### 2.类，名称空间，类库



#### 3.类的三大成员

属性，方法，事件

#### 4.C#语言的基本元素

#### 5.关键字Keyword（保留字）





#### 6.操作符

#### 7.标识符

#### 8.函数编写

```
class Calculator
{
	pulic int Add(int a, int b)
	{
		int result = a + b;
		return result;
	}
}
```

#### 9.for循环

```
class Calculator
{
    public void Print1ToX(int x)
    {
        for (int i = 0; i < x; i++)
        {
            Console.WriteLine(i);
        }
    }
}
```

#### 10.递归

```
class Calculator
{
    public void PrintXTo1(int x)
    {
        if (x == 1) 
        {
            Console.WriteLine(x);
        }
        else
        {
            Console.WriteLine(x);
            PrintXTo1(x - 1);
        }
    }
}
```











#### 100.快捷键

ctrl+alt+L 打开资源管理器

#### 问题

1.pulic是指公用方法吗?

2.void是func的意思吗?

3.权限命名