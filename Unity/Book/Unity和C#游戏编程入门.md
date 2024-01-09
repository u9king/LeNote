# Unity和C#游戏编程入门

#### 第1章 了解开发环境

Unity下载地址：https://www.unity.com

Unity开发文档：file:///D:/Unity/2022.3.9f1/Editor/Data/Documentation/en/Manual/index.html

UnityHub应用程序资源：https://docs.unity3d.com/Manual/index.html

Unity窗口功能：https://docs.unity3d.com/Manual/UsingTheEditor.html

Unity社区解决方案：

Unity Forums：https://forum.unity.com

Unity Answers:https://answer.unity.com/index.html

Unity Discord: https://discord.com/invite/unity

C#技术文档:https://docs.microsoft.com/en-us/dotnet/csharp/programming-guide/index

#### 第2章 编程的构成要素

Unity类

```
public class LearningCurve : MonoBehaviour
其中MonoBehaviour表示该类可以附加到Unity场景中的GameObject上。在C#中，类可以独立存在。
```

C#代码注释

```
鼠标悬停显示
在代码前输入：///+Enter
/// <summary>
/// 
/// </summary>
```

来自MonoBehavior的助力

```
C#脚本是类，那么Unity又是如何知道该把哪些脚本视为组件而哪些不呢？
答案很简单，LearningCurve（以及在Unity中创建的任何脚本）都继承自名称为MonoBehavior的另一个类。这便告诉Unity，可以将这个C#类转换为组件。
```



#### 第3章 深入研究变量、类型和方法

大括号用法

```
传统C#是让每个大括号单独占一行
public void MethodName()
{
    
}
而Unity习惯将第一个大括号与声明处于同一行
public void MethodName(){
    
}
```

任何未标记为public的变量的访问修饰符将默认为private，并且不会显示在Unity编辑器的Inspector面板中。

C#访问修饰符accessModifier

```
public:公有变量,所有人都能用
private：私有变量，所属类内用
protected:受保护变量，所属及派生类内用
internal:内部变量，程序集中使用
```

创建内插字符串

```
可以通过$头，将变量插入文本中，同时将变量用花括号括起来
void Start()
{
    Debug.Log($"A string can have variables like {firstName} inserted directly!");
}
```

类型转换

Convert类可以处理更多复杂的转换

#### 第4章 控制流和集合类型

数组创建

```
int[] topPlayerScores = new int[3]{713,549,984};
```

列表

```
List<string>questPartyMembers = new List<string>()
{"Grim the Barbarian", "Merlin the Wise", "Sterling the Knight"};
//尾部插入
questPartyMembers.Add("Carven the Necromancer");
//指定序号插入
questPartyMembers.Insert(1,"Tanis the Thief");
//删除指定位置
questPartyMembers.RemoveAt(0);
//删除指定名称
questPartyMembers.Remove("Grim the Barbarian");
```

字典

```
Dictionary<string, int>itemInventory = new Dictionary<string,int>()
{
	{“Potion”,5},
	{“Antidote”,7},
	{“Aspirin”,1},
}
```

迭代语句

```
//for语句
for(int i = 0; i < questPartyMembers.Count;i++)
{
	Debug.LogFormat("Index:{0} - {1}", i, questPartyMembers[i]);
}

//foreach语句
foreach(string partyMember in questPartyMembers)
{
	Debug.LogFormat("{0} - Here!",partyMember);
}

//while语句
while(playerLives > 0)
{
	Debug.Log("Still alive!");
	playerLives--;
}
```

#### 第5章 类、结构体和OOP

class是引用类型，结构体是值类型

```
区别在hero2 = hero时
修改hero2.name
当类型为class时，hero也对应修改
当类型为struct时,hero不修改
```

class对象都为引用类型，所以要有重置方法来还原

面向对象多态

```
父元素  virtual标记为虚拟方法
public virtual void PrintStatsInfo() 
{
    Debug.LogFormat("Hero :{0} - {1} EXP", name, exp);
}
子元素  派生类可以通过override进行改写
public override void PrintStatsInfo()
{
    Debug.LogFormat("Hail :{0} - take up your {1}", name, weapon.name);
}
```

#### 第6章 亲手实践Unity

#### 第7章 角色移动、摄像机及碰撞

Character Controller:https://docs.unity3d.com/ScriptReference/CharacterController.html

#### 第8章 游戏机制脚本编写

#### 第9章 AI基础与敌人行为

#### 第10章 再谈类型、方法和类

#### 第11章 栈、队列和HashSet

#### 第12章 探索泛型、委托等

#### 第13章 旅程继续



#### 课后作业

1.翻译MonoBehavior类的解释文档https://docs.unity3d.com/ScriptReference/MonoBehaviour.html

