# 虚幻引擎C++

> | 序号 | 课程              | 作者   | 链接                                               | 备注 |
> | ---- | ----------------- | ------ | -------------------------------------------------- | ---- |
> | 1    | UE系列之C++四部曲 | 唐老狮 | [b站](https://www.bilibili.com/video/BV1GQSJYCEXb) |      |
> |      |                   |        |                                                    |      |

### 0.待分类

#### 1. SetupInputComponent

是UE中一个非常经典的基类成员函数Method

```
全称是：APawn::SetupInputComponent() 或 ACharacter::SetupInputComponent()。
```

当玩家Controller在游戏里控制Possess了一个实体（Pawn 或 Character）时，引擎会自动调用这个函数。

作用：把玩家物理键盘/手柄的输入（比如按 W 键、摇动摇杆），与你 C++ 代码里的函数（比如向前移动、跳跃、释放技能）绑定在一起。

在这个函数里，你会看到一个叫 InputComponent 的全局变量。它才是真正的组件（类型是 UInputComponent 或 UEnhancedInputComponent）。

这个组件实际上是随玩家的 APlayerController控制器 一起创建的







#### 2.UEnhancedInputLocalPlayerSubsystem

SubSystem是套什么系统？

#### 3.前向声明

前向声明是什么，为什么能提高性能，相比于头文件。

#### 4.如何快速修改代码放错文件夹的问题，而不导致崩溃。

#### 5.有哪些部分是GAS单独为C++开放而不为蓝图开放的？

#### 6.GAS

- Ability SystemComponent
- Attribute Set
- GameplayAbility
- GameplayEffect
- GameplayCue
- Ability Task
- Gameplay Tag



#### 7.接口中的0是啥意思？

```C++
class IAbilitySystemInterface
{
    GENERATED_IINTERFACE_BODY()

    /** Returns the ability system component to use for this actor. It may live on another actor, such as a Pawn using the PlayerState's component */
    virtual UAbilitySystemComponent* GetAbilitySystemComponent() const = 0;
};
为什么实现的代码必须返回AbilitySystemComponent或者nullptr
```

















### X.长期问题

#### 1.Visual Studio 2022在打开虚幻引擎的C++代码会进行长期的IntelliSense编译，但是全部关掉后完全没有代码提示，如何处理？

答：官方推荐是VisualStudio2022但是在UE5.4~5.7中长期存在这个问题无法解决，

即使参考官方推荐文档设置也无法解决：https://dev.epicgames.com/documentation/unreal-engine/setting-up-visual-studio-development-environment-for-cplusplus-projects-in-unreal-engine?application_version=5.7

目前解决方案是使用Rider，因为Rider在底层做了高度适配。