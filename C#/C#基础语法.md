# C#基础语法

#### 1.十种技术编写Hello world

Console

```c#
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

```c#
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

#### 9.循环

```c#
//for循环
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

```c#
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

#### 11.数据类型（Data Type）

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

区别：JavaScript为弱类型语言，C#为强类型语言。比较方法C语言示例：if条件。C#语言可以使用dynamic对弱类型进行模仿。

#### 12.C#五大数据类型

类（Classes），如Windows，Form，Console

结构体（Structures）：如Int32，Int64，single,double

枚举（Enumerations）：如HorizontalAlignment Visibility

接口（Interfaces）

委托（Delegates）

#### 13.C#类型的派生谱系

![image-20230925111816132](C:\Users\u9king\AppData\Roaming\Typora\typora-user-images\image-20230925111816132.png)

引用类型：object，string，class等

值类型：bool，byte，char，struct等

#### 14.值类型变量与引用类型变量的区别

值类型变量：

- 根据值类型申请字节长度（最高位可能为符号位）

引用类型变量：

- 申请占用四个字节
- 创建实例后将实例的首个内存地址编号存入四个字节中
- 创建实例按照内部要求请求地址

#### 15.内存存储（补码反码）

```c#
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

注：左侧$\large\textcolor{orange}{橙色}$区域为内存保留地址，$\large\textcolor{yellow}{黄色}$区域表示已经被占的区域

#### 16.自顶向下算法

面向过程算法

#### 17.构造器

快捷键:ctor+tab+tab

```c#
namespace ConstructorExample
{
    class Program
    {
        static void Main(string[] args)
        {
            Student stu = new Student();
            Console.WriteLine(stu.ID);
        }
    }
    class Student
    {
        public Student()
        {
            this.ID = 1;
            this.Name = "No name";
        }
        public int ID;
        public string Name;
    }
}
```

#### 18.方法重载Overload

方法签名：由方法的名称，类型参数的个数和它的每一个形参（从左到右）的类型和种类（值、引用或输出）组成。方法签名不包括返回类型。

#### 19.程序调用过程（内存中的栈和堆）

程序存储在硬盘中，运行程序后，装载到内存中，内存会被分为两个空间（Stack和Heap）。

栈（Stack）：用来函数调用（占用较小），会有栈爆的风险（栈溢出Stackoveflow）

```C#
//栈溢出实现
namespace StackOverflow
{
    internal class Program
    {
        static void Main(string[] args)
        {
            BadGuy bg = new BadGuy();
            bg.BadMethod();
        }
    }

    class BadGuy
    {
        public void BadMethod()
        {
            int x;
            this.BadMethod();
        }
    }
}

```

堆（Heap）：用来存储对象（占用较大），不会爆，但是不回收会有内存泄漏的风险，使用Performance Mointor查看进程的堆内存使用情况

控制台输入perfmon打开性能监视器,需要调整。

```

```

<img src="E:\ImageHostingService\C#\性能监视器.jpg">

#### 20.变量

在C#中，变量一共有7种。

​		静态变量，实例变量，数组元素，值参数，引用参数，输出形参，局部变量

静态变量是隶属于类而不是类的实例的

```c#
namespace TypeInCSharp
{ 
    class Program
    {
        static void Main(string[] args)
        {
            Student.Amount = 1;		//静态变量
        }
    }

    class Student
    {
        public static int Amount;
        public double Add(ref double a, double b)  //引用参数变量
        {
            return a+b;
        }
        
        public double Addd(out double a, double b)  //输出参数变量
        {
            return 2a+b;
        }
    }

}
```

#### 21.数组

```C#
//数组申请
int[] array = new int[100];
```









#### 100.快捷键

截图快捷键 fn+R

#### 问题 

1.pulic是指公用方法吗?

2.void是func的意思吗?

3.权限命名

4.浮点型到底是什么，为什么有浮点型

5.内存泄漏

6.函数重载

7.类型形参

9.为什么会产生命名空间这个概念？

11.WPF没法用，C语言没法用，C#中生成这么多文件是为什么？每个是什么意思

12.离线文档如何实现