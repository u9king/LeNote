# C#设计模式—23种设计模式

### 第1类 单例模式

#### 第一种 传统单例

别称:饿汉式

缺点：默认加载时就生成了instance

```C#
namespace DesignPattern1
{
    public class Singleton
    {
        private static readonly Singleton instance = new Singleton();
        public static Singleton GetInstance()
        {
            return instance;
        }
        private Singleton() { }  //私有化构造函数
    }
}
```

#### 第二种 静态语句块单例

缺点：本质和第一种，区别只有在第一次访问时才创建

```C#
namespace DesignPattern2
{
    public class Singleton
    {
        private static readonly Singleton instance;
        static Singleton()
        {
            instance = new Singleton();
        }

        public static Singleton GetInstance()
        {
            return instance;
        }
        private Singleton() { }
    }
}
```

#### 第三种 懒加载单例

问题： <span style="color:red;">多线程</span>时会存在，多次创建的问题

```C#
namespace DesignPattern3
{
    public class Singleton : MonoBehaviour
    {
        private static Singleton instance;
        public static Singleton Instance
        {
            get
            {
                if (instance == null)
                {
                    instance = new Singleton();
                }
                return instance;
            }
        }
        private Singleton() { }
    }
}
```

#### 第四种 双重检查锁定

缺点：完美的类型之一，但是写法过于繁琐，特别是对于instance的双重锁定，以及多线程引入的线程锁

```C#
namespace DesignPattern4
{
    public class Singleton
    {
        private static Singleton instance;
        private static readonly object lockObj = new object();  //多线程锁
        public static Singleton Instance
        {
            get
            {
                if (instance == null)
                {
                    lock (lockObj)
                    {
                        if (instance == null)
                        {
                            instance = new Singleton();
                        }
                    }
                }
                return instance;
            }
        }
        private Singleton() { }
    }
}
```













### 第2章 XXX

### 第3章 XXX

### 第4章 XXX

### 第5章 XXX

### 第6章 XXX

### 第7章 XXX

### 第8章 XXX

### 第9章 XXX

### 第10章 XXX

### 第11章 XXX

### 第12章 XXX

### 第13章 XXX

### 第14章 XXX

### 第15章 XXX

### 第16章 XXX

### 









