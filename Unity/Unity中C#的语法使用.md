# Unity中C#的语法使用

#### 1.程序基础结构

```C#
using System —— 引用命名空间 —— 引用工具包
namespace —— 命名空间 —— 工具包
class —— 类 ——工具
func —— 函数 ——工具能做的事情
```

#### 2.折叠代码

```
#region MyRegion(折叠是显示的内容)
#endregion
只在编辑时起作用，发布时会被自动删除
```

#### 3.浮点数

```
float 存储7/8为有效数字 根据编译器不同 有效数字也可能不一样
有效数字 是从左到右非0数开始算有效数字的
float f = 0.1234567890f;  (需要加上f，C#默认小数为double类型)
实际打印:0.12345679
float f = 0.01234567890f;
实际打印:0.012345679
float f = 1.01234567890f;  (需要加上f，C#默认小数为double类型)
实际打印:1.0123457
Console.WriteLine(f);
```

#### 4.多变量同时声明

```
int a1 = 1, b1 = 2, c1 = 3, d1 = 4;
```

#### 5.取消转义字符

```
string str = @"哈哈\哈哈"
```

#### 6.隐式转换

```
隐式转换，系统自动转换
大范围可以装下小范围
浮点数可以装在任何类型的整数(超过部分会使用科学计数法换算)
示例:long l = 1;			 |    float f2;
	int i = 1;			  |	   long l = 1;	
	l = i;   //隐式转换    |     f2 = l;
```

#### 7.显式转换

```
方法1:强制转换法
short s = 1;
int i = 1;
s = (short)i
方法2:Parse法
int i4 = int.Parse("123");
方法3:Convert法
int a = Convert.ToInt32("12");
方法4:ToString法
string str6 = 1.ToString();
```

#### 8.异常捕获

```C#
//必备部分
try
{

}
catch(Exception e)	//e是具体错误信息
{
	//出错执行
}
//可选部分
finally
{
	//不管出错与否都会执行
}
```

#### 9.自增运算符

```
int a = 1;
a++;	先用再加
++a;	先加再用
```

#### 10.字符串拼接

```
string str = string.Format("a:{0},b:{1}","A","B");
```

#### 11.三目运算符

```
bool ? 真:假;
```

#### 12.do while语句

```
do
{
	//先执行后判断
}while(bool);	//使用较少
```

#### 13.委托,事件,Action,Func的关系

```
1.Action是delegate的简写
2.event约束是限制功能的delegate
3.Func是带返回值(最后一个参数)的delegate——注多播情况下返回最后挂的一个委托的返回值
```

#### 14.字符串转枚举

```
playerType = (E_PlayerType)Enum.Parse(typeof(E_PlayerType),"Other");
```

#### 15.数组

```
变量类型[] 数组名 = new 变量类型[]{内容1，内容2，内容3，......};
int[] arr = new int[]{ 1, 2, 3, ......};
简化写法：
int[] arr = { 1, 2, 3, ......};
//建立定长空数组
int[] arr = new int[5];
注释：
数组初始化后是不能添加新的元素的！！！
所以需要开新数组用循环进行搬家
搬家代码：
int[] array2 = new int[6];	//建新房
for (int i = 0; i < array.Length; i++)
{
	array2[i] = array[i];	//搬数值
}
array = array2;//搬地址
```

#### 16.二维数组

```
变量类型[,] 数组名 = new 变量类型[,]{{内容1，内容2，内容3},{},{}......};
int[] arr = new int[,]{{1,2,3},
					   {4,5,6},
					   {7,8,9}};
简化写法：
int[] arr = {{1,2,3},
			 {4,5,6},
			 {7,8,9}};
```

#### 17.交错数组

```
变量类型[][] 数组名 = new 变量类型[][]{new 类型[]{},new 类型[]{},......};
int[][] arr = new int[][]{new int[]{1,2,3},
					   	  new int[]{1,2},
					      new int[]{1}};
简化写法：
int[] arr = {new int[]{1,2,3},
			 new int[]{1,2},
			 new int[]{1}};
```

#### 18.引用类型

```
应用类型：数组，class，string
表现：会随引用的变化而变化
	 值类型它变我不变，引用类型它变我也变
	 值类型存储在栈空间上	—— 系统分配，自动回收，小而快
	 引用类型存储在堆空间上，用栈存地址  —— 手动申请和释放，大而慢
int[] arr =new int[]{1}
int[] arr2 = arr
arr2[0] = 2;
此时如果打印 arr[0]的值为 2;
```

#### 19.ref和out

```
函数参数修饰符
ref:当传入的值类型参数在内部修改（引用类型参数在内部声明）时，外部的值产生变化
static void ChangeValueRef(ref int value)
{
	value = 3;
}
使用: ChangeValueRef(ref a);
out:
static void ChangeValueOut(out int value)
{
	value = 3;
}
使用: ChangeValueOut(out a);

ref和out的区别
1.ref传入的变量必须初始化 out不用
2.out传入的变量必须在内部赋值  ref不用
说白了就是，ref是引用，你得先有；out是输出，我帮你留了位置，你自己定
```

#### 20.变长参数

```
params关键字 一个函数只有一个
static int Sum(string name, params int[] arr)
{
	int sum = 0;
	for (int i = 0; i < arr.Length; i++)
	{
		sum += arr[i];
	}
	return sum;
}
```

#### 21.结构体

```
struct Student
{
	//结果提中申明变量不能直接初始化
	int age;
	bool sex;
}
```

#### 22.类型默认值

```
default(int);  //值类型的都能查看，引用类型出来是空
```

#### 23.构造函数

```
class Person
{
	public Person()
	{
	
	}
}

特殊继承:
class Person
{
	private int a;
    private int b;
	public Person()
	{
		a = 1;
	}
	public Person(int b):this()	特殊继承，类似贯穿:
	{
		b = 2;
		Debug.Log(a,b);  //a = 1 b = 2
	}
}
```

#### 24.析构函数

```
class Person
{
	//作为垃圾回收的时候触发
	~ Person()
	{
	
	}
}
```

#### 25.垃圾回收GC

```
Garbage Collector
垃圾回收有很多种算法，比如
//代级回收（0代，1代，2代）  新来的进入0代
//引用计数(Reference Counting）
//标记清除(Mark Sweep)
//标记整理(Mark Compact)
//复制集合(Copy Collection)

GC只负责堆内存的垃圾回收，栈有自己的生命周期GC管不着
System.GC.Collect();
```

#### 26.成员属性(get,set)

```
class Person
{
	private string name;
	public string Name
	{
		get
		{
			return name;
		}
		set
		{
			name = value;  //其中value只在set中有用
		}
	}
}

//简写
public float Height
{
	get;
	private set;
}
```

#### 27.索引器

```
class Person
{
	private string name;
	private int age;
	private Person[] friends;
	
	public Person this[int index]	//索引器,this[]
	{
		get
		{
			return friends[index];
		}
		set
		{
			friends[index] = value;
		}
	}
}
```

#### 28.静态类

```
public static class Tools
{
	//静态成员变量
	public static int a = 0;
	public static void TestFun()
	{
	
	}
}
```

#### 29.静态类拓展方法

```
慎用！！！  可能会破坏原本类的封装
格式：
//访问修饰符 static 返回值(this 拓展类名 参数名，参数类型 参数名 ...)
static class Tools{
	public static void SpeakValue(this int value){}；
	这就是给int这个类拓展SpeakValue的方法
}

示例：
static class Tools
{
	public static void SpeakValue(this int a)
	{
		Console.WriteLine("1:" + a);
	}
}
class Main
{
	int i = 10;
	Debug.Log(i.SpeakValue);
}
```

#### 30.运算符重载

```
关键字：operator
可重载符号：+ - * / % ++ -- ！(> <:成对出现) (>= <=) == !=
格式：public static 返回值 operator 运算符(参数列表)
class Point
{
	public int x;
	public int y
	public static Point operator +(Point p1, Point p2)
	{
		Point p = new Point();
		p.x = p1.x + p2.x;
		p.y = p1.y + p2.y
	}
}
class Main
{
	Point p1 = new Point(1, 1);
	Point p2 = new Point(2, 2);
	Point p3 = p1 + p2;
}
```

#### 31.分部类

```
关键词：partial
partial class Student
{
	public bool sex;
	public string name;
}
partial class Student
{
	public int number;
}
```

#### 32.里氏替换原则(is/as)

```
重要性：面向对象七大原则中最重要的原则
基本概念：任何父类出现的地方，子类都可以替代
class GameObject{}
Class Player : GameObject
{
 	public void PlayAtk()
 	{
 		Console.WriteLine("玩家攻击");
	 }
}
Class Monster : GameObject
{
 	public void MonsterAtk()
 	{
 		Console.WriteLine("怪物攻击");
	 }
}
class Main
{
	//里氏替换原则  用父类容器  装载子类对象
	GameObject player = new Player();
    GameObject monster = new Monster();
    GameObject[] objects = new GameObject[]{new Player(),new Monster()}
    
    #region is和as
    //is：判断一个对象是否是指定类对象
    //返回值 bool 是为真 不是为假
    if(player is Player)
    {
    
    }
    //as: 将一个对象转换为指定类对象
    //返回值：指定类型对象 失败返回null
    #endregion
}
```

#### 33.继承中的构造

```
1.默认子类会自动执行父类的构造函数一次
2.父类的无参构造很重要（会影响子类继承）/可以用base找到某个构造函数
Class Father
{
	public Father(int i)
	{
		Console.WriteLine("父类构造");
	}
}
class Son : Father
{
	public Son(int i): base(i)
	{	
	}
}
class Main
{
	Son s = new Son();
}
```

#### 34.装箱拆箱

```
装箱：值类型用引用类型存储，栈内存迁移到堆内存中
例：object o = 1f;
使用时 float f = (float)o; (强转)
```

