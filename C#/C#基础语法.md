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

#### 8.方法（函数）

函数名称的首字母大写，取名用动词短语

```
//方法头
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

#### 11.数据类型

| C#类型关键字 | Range                                       | 字节 |
| ------------ | ------------------------------------------- | ---- |
| sbyte        | -128 to 127                                 | 1    |
| byte         | 0 to 255                                    | 1    |
| char         | U+0000 to U+ffff                            | -    |
| short        | -327648 to 32767                            | 2    |
| ushort       | 0 to 65535                                  | 2    |
| int          | -2147483648 to 2147483647                   | 4    |
| uint         | 0 to 4294967295                             | 4    |
| long         | -9223372036854775808 to 9223372036854775807 | 8    |
| ulong        | 0 to 18446744073709551615                   | 8    |
| float        | ±1.5e-45 to ± 3.4e38                        | 4    |
| double       | ±5.0e-324 to 1.7e308                        | 8    |

#### 12.C#五大数据类型

类Classes，如Windows，Form，Console

结构体Structures：如Int32，Int64，single,double

枚举Enumerations：如HorizontalAlignment Visibility

接口Interfaces

委托Delegates

#### 13.C#类型的派生谱系

![image-20230925111816132](C:\Users\u9king\AppData\Roaming\Typora\typora-user-images\image-20230925111816132.png)

#### 14.内存存储（补码反码）

```
internal class Program
{
    static void Main(string[] args)
    {
        short s;
        s = -1000;
        string str = Convert.ToString(s, 2);
        Console.WriteLine(str);
    }
}
```

注意：存入时为反码，补码，注意高高低低原则，高位存高处

![image-20230925125247350](C:\Users\u9king\AppData\Roaming\Typora\typora-user-images\image-20230925125247350.png)

![image-20230925125315698](C:\Users\u9king\AppData\Roaming\Typora\typora-user-images\image-20230925125315698.png)

#### 15.引用类型

首先占用四个字节

创建实例后将实例的首个内存地址编号存入四个字节中

创建实例按照内部要求请求地址

#### 16.自顶向下算法

面向过程算法

#### 17.





#### 100.快捷键

ctrl+alt+L 打开资源管理器

#### 问题 

1.pulic是指公用方法吗?

2.void是func的意思吗?

3.权限命名

4.浮点型到底是什么，为什么有浮点型

5.内存泄漏