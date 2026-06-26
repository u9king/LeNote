# 虚幻引擎物理Chaos

> 别称：Chaos物理系统
>
> | 序号 | 课程                                                         | 作者          | 链接                                                         | 备注                                               |
> | ---- | ------------------------------------------------------------ | ------------- | ------------------------------------------------------------ | -------------------------------------------------- |
> | 1    | Unreal Engine 5: Physics                                     | David Nixon   | [Udemy](https://www.udemy.com/course/unreal-engine-5-for-games-physics/) | [b站](https://www.bilibili.com/video/BV1nn9iYCE2W) |
> | 2    | 19-物理模拟 \| Epic 肖月                                     | 肖月          | [b站](https://www.bilibili.com/video/BV1S54y1N7Fk)           |                                                    |
> | 3    | UE物理引擎-入门教程                                          | 王瑨Artist    | [b站](https://www.bilibili.com/video/BV1U1421C7sZ)           |                                                    |
> | 4    | [技术分享]《UE中的世界物理场及应用》 \| 李文磊 Epic          | 李文磊        | [Youtube](https://www.bilibili.com/video/BV1dP411s7Mn)       |                                                    |
> | 5    | An Introduction to the Mover Plugin \|Unreal Fest2024        | Unreal Enigne | [Youtube](https://www.youtube.com/watch?v=P4IKS5k47Wg&t=716s) |                                                    |
> | 6    | A Dive Into Networked Physics-Based Character Movement \| Unreal Fest Stockholm 2025 | Unreal Enigne | [Youtube](https://www.youtube.com/watch?v=_jRLlTDqoGI&t=225s) |                                                    |



### 0.待分类

#### 1.物理网格体

物理体网格体：由基元集合定义的

#### 2.基元组件 Primitive Component

简单碰撞确实是依托于基元组件来存储和发生作用的。

在 UE5 中，只要一个组件要么能被看见，要么能被碰撞，它在底层就一定继承自 UPrimitiveComponent。











### 1.Chaos物理系统

Chaos物理系统是虚幻引擎提供的轻量级物理模拟解决方案。 

计算在物理体而不是在网格体上。物理体也可以叫做刚体。

该系统包括以下主要功能：

| 刚体动力学             | 破坏效果        | 布料物理和机器学习布料模拟 |
| ---------------------- | --------------- | -------------------------- |
| 物理场                 | Chaos可视调试器 | 毛发物理                   |
| 刚体动画节点和物理动画 | 联网物理        | 血肉模拟                   |
| 载具                   | 布娃娃物理      | 流体模拟                   |



#### 1.1 帧率

> 别称：子步，迭代
>
> 路径：Physics -> 帧率

物理引擎包括：PhysX，Chaos等



### 2.刚体动力学

> 别称：物理体，Chaos物理体

<img src="https://dev.epicgames.com/community/api/documentation/image/a89b1a8d-039e-41f7-80c4-eecd932532a8?resizing_type=fit" style="zoom: 50%;" />

Chaos物理系统提供了许多种刚体动力学功能。 这包括碰撞响应、物理约束以及阻尼和摩擦。 此外，它还提供了异步物理模拟和联网物理。



刚体是指在运动中和受到力的作用后，形状和大小不变，而且内部各点的相对位置不变的物体

Chaos物理系统提供了许多种刚体动力学功能。 这包括碰撞响应、物理约束以及阻尼和摩擦。 此外，它还提供了异步物理模拟和联网物理。

- 质量
- 线性阻尼：摩擦力
- 角阻尼：旋转摩擦力
- 约束
  - 



#### 2.1 物理组件 Physics Component

##### 2.1.1 模拟物理

##### 2.1.2 质量

不勾选自动计算，假定一个平均密度

##### 2.1.3 线性阻尼

> 别称：空气阻力 / 介质粘滞力
>
> 计算：仅与物体速度成正比，速度越大阻力越大
>
> 特别：非常类似摩擦力，但不是摩擦力，摩擦力需要两个面接触，线性阻尼不需要

用来定义应该对物体的运动施加多大的摩擦力或阻力的。

##### 2.1.4 角阻尼

> 计算：仅与物体速度成正比，速度越大阻力越大

物体旋转运动的摩擦力

##### 2.1.5 启用重力

默认开启

##### 2.1.6 约束

锁定平移XYZ，和旋转轴XYZ











#### 2.2 碰撞组件 Collision Component

> 命令行：
>
> - show Collision：显示所有碰撞模型和阻挡体积
> - stat game：显示关于碰撞所用时间的统计信息

原则：最简原则：重叠与碰撞，结果是重叠

##### 2.2.1 碰撞外形

- 凸包Convex Hull：包裹最小形状
  - 常见的凸包求解方法：扫描法 Graham，步进法 Jarvis，快包法 QuickHull
- 简单碰撞 Simple Collision：
  - 复杂模型拿去算碰撞，计算量太大。 为了优化，我们用基元组件的几何形状去包裹这个模型。这些由简单数学公式构成的隐式几何体，就叫做简单碰撞。
- KDOP：K离散导向多面体，是包围体的一种，它抓取轴对齐的平面，将其尽力推向离网格体最近的位置。结果形态用作碰撞凸包。
  - 10 - 方块有 4 条边形成斜角 - 可选择 X、Y 或 Z 轴对齐的边。
  - 18 -方块中所有边均形成斜角。
  - 26 - 方块中所有边和角均形成斜角。
- 复杂碰撞体：Complex Collision：复杂碰撞体是完全精确，但是计算量很大，无法使用于碰撞，必须选择将复杂作用于简单。



简单碰撞

复杂碰撞



因为凸包没有内凹，物理引擎可以使用类似 GJK 算法在极短时间内判断两个凸体是否碰撞。

##### 2.2.2 碰撞类型

Trace Responses 射线反应



ObjectType 物体类型

- WorldStatic：世界静态
- WorldDynamic：世界动态
- Pawn：棋子
- PhysicsBody：物理体
- Vehicle：车辆
- Destructbile：破碎
- Projectile：发射体

##### 2.2.3 碰撞通道 

###### 2.2.3.1 对象通道 Object Channels

> 路径：Edit Menu -> Project Settings -> Collision

| 对象类型     | 翻译     | 说明                                              |
| ------------ | -------- | ------------------------------------------------- |
| WorldStatic  | 世界静态 | 应用于不移动的任意Actor                           |
| WorldDynamic | 世界动态 | 用于受到动画或代码的影响而移动的Actor类型；运动学 |
| Pawn         | 棋子     | 任何由玩家控制的实体的类型都应为Pawn              |
| PhysicsBody  | 物理物体 | 由于物理模拟而移动的任意Actor                     |
| Vehicle      | 载具     | 载具的默认类型                                    |
| Destructible | 可破坏   | 可破坏物网格体的默认类型                          |

###### 2.2.3.2 追踪通道 Trace Channels

> 路径：Edit Menu-> Project Settings -> Collision 

默认有两个追踪通道可视性Visibility和摄像机Camera。

你可能需要一种特殊激光，它需要能够穿过你无法透视或无法让摄像机穿过的特殊不透明对象。你添加自己的自定义追踪响应通道。



##### 2.2.4 碰撞复杂度 Collision Complexity

- 简单和复杂 Simple And Complex：默认状态。引擎同时为它创建简单和复杂两套外壳，各司其职。
- 将简单作用于复杂 Use Simple Collision As Complex： 无论任何时候，都只使用那个绿色的简单外壳。最推荐、性能最好。
- 将复杂作用于简单 Use Complex Collision As Simple：全部使用复杂
  - 无法使用物理
  - 仅使用于边界碰撞
  - 精确的射线检测



**最简原则**：重叠与碰撞，结果是重叠

忽略，重叠，阻挡

##### 2.2.5 碰撞预设

预设是一堆已经设置好的类型

##### 2.2.6 自定义

自定义预设

自定义通道

自定义类型



#### 2.3 射线检测 Trace

![](https://pica.zhimg.com/v2-6bf20bddcb43794f8c5b612293e8a11e_1440w.jpg)

| 序号 | 名称            | 翻译               | 作用                                                         |
| ---- | --------------- | ------------------ | ------------------------------------------------------------ |
| 1    | Blocking Hit    | 阻挡击中           | 是否发生了碰撞                                               |
| 2    | Initial Overlap | 初始重叠           | 射线形状是否发生重叠，当Initial Overlap为true时，也就意味着Distance的值为零。 |
| 3    | Time            | 碰撞时间           | 碰撞时间，值的范围是（0~1），当发生了初始重叠时 Time为0，当未发生碰撞时 Time为1（即无穷远） |
| 4    | Distance        | 起点与碰撞点间距离 | 射线起点到碰撞点之间的距离。注意：当Distance为0时可能有两种情况，发生初始重叠或未发生碰撞，需要通过Initial Overlap来判断。 |
| 5    | Location        | 球体(射线位置)位置 | 触发碰撞时射线形状的中心点位置                               |
| 6    | Impact Point    | 碰撞点             | 触发碰撞的实际点位置                                         |
| 7    | Normal          | 法线               |                                                              |
| 8    | Impact Normal   | 碰撞法线           | 被碰撞表面的法线方向，很多时候Normal就等于Impact Normal，我实验了几种情况没有发现它们的区别，此处如有错误敬请指正！ |
| 9    | PhysMat         | 物理材质           | 被射线碰撞对象的物理材质                                     |
| 10   | HitActor        | 碰撞对象           | 被射线碰撞的Actor                                            |
| 11   | HitComponent    | 碰撞图元组件       | 被射线碰撞的Primitive Component                              |
| 12   | HitBoneName     | 碰撞骨骼名         | 如果被射线碰撞对象是一个SkeletalMesh，那么将会返回被碰撞的骨头名称 |
| 13   | HitItem         | 碰撞图元索引       | 被碰撞的Primitive索引                                        |
| 14   | ElementIndex    | 碰撞元素索引       | 如果一个Primitive是由多个部分组成的，那么这里返回被碰撞元素的索引 |
| 15   | FaceIndex       | 面索引             |                                                              |
| 16   | TraceStart      | 射线起点           |                                                              |
| 17   | TraceEnd        | 射线终点           |                                                              |



| <img src="https://picx.zhimg.com/v2-64228e34ea58343d77ea36ba5d4090b3_1440w.jpg" alt="img" style="zoom:25%;" /> | <img src="https://pic2.zhimg.com/v2-aa0b72e3ba74716de0ab0658bf5e3023_1440w.jpg" style="zoom:25%;" /> | <img src="https://pica.zhimg.com/v2-010fea7bf996e16f7cb23ba1780df666_1440w.jpg" style="zoom: 50%;" /> | <img src="https://pic3.zhimg.com/v2-d00f837ee97bac937fc736c1948ffe9c_1440w.jpg" style="zoom:25%;" /> |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |                                                              |                                                              |



- 在ALS V4的CapsuleHasRoomCheck方法中，作者就借助Initial Overlap来判断一个空间是否有足够的位置容纳角色的胶囊体。

#### 2.4 物理约束 Physics Constraint

定义了刚体与刚体之间的关系



基础:距离/角度控制
状态:Rest Target
高级:Soft Plastic Break
校正:Projection



- Linear Limits 线性约束
    - Free 不控制
    - Limited 约束范围
    - Locked 锁住,没有相对运动
- Angular Limits 角度约束
- Rest Target
    - 线性马达
    - 角发动机
- Soft Constraint 软约束
- Plasticity 塑性约束
- Breakable 断裂约束
- Projection 投影



#### 2.5.物理材质 Physics Material

> 前缀：PM_XXX

可以分配给刚体，摩擦力弹力

- 摩擦力
- 静态摩擦
- 摩擦组合模式
- 覆盖摩擦组合模式
- 恢复力
- 恢复力组合模式
- 覆盖弹性恢复组合模式
- 密度
- 睡眠线性速度阈值
- 睡眠角速度阈值
- 睡眠计数阈值












#### 2.6.物理资产 Physics Asset

> 缩写：PA，RBAN
>
> 前缀：PA_
>
> 快捷键：Simulate(骨架按钮边上)模式下，按住Control + 右键可以控制玩家
>
> 切换：SKM骨骼网格体-> Asset Detail资产细节 ->Physics物理 进行切换

作用：骨架网格体与物理模拟的交互方式

一套SKM同时使用一个PhysicsAsset



约束之间可以旋转

创建，右键骨骼创建

















#### 2.7 其他组件

##### 2.7.1 物理动画组件

<img src="https://d1iv7db44yhgxn.cloudfront.net/documentation/images/5ee04273-523a-45c5-8de7-3a7fea0ec4cf/ue5_1-physical-animation.png" style="zoom:25%;" />

物理动画组件（Physical Animation Component） 在 骨骼网格体（Skeletal Mesh） 动画顶部应用 物理模拟 。通过使用该组件，你可以在播放动画的同时将真实的物理模拟应用到骨骼网格体中的特定骨骼组。

##### 2.7.2 物理约束组件

<img src="https://d1iv7db44yhgxn.cloudfront.net/documentation/images/585e0e58-b4c1-40ff-bab8-77e672d8dae1/ue5_1-constraint.png" style="zoom:25%;" />

​	物理约束组件（PhysicsConstraintComponent） 是一种能连接两个刚性物体的接合点。你可以借助该组件的各类参数来创建不同类型的接合点。

​	借助 PhysicsConstraintComponent 和两个StaticMeshComponents ，你可以创建悬摆型对象，如秋千、重沙袋或标牌。它们可以对世界中的物理作用做出响应，让玩家与之互动。

##### 2.7.3 物理抓柄组件

<img src="https://d1iv7db44yhgxn.cloudfront.net/documentation/images/bab4254c-e88d-4aa8-ad2d-9718ee92a21a/ue5_1-handle.png" style="zoom:25%;" />

物理抓柄组件（PhysicsHandleComponent） 用于"抓取"和移动物理对象，同时允许抓取对象继续使用物理效果。案例包括"重力枪"——你可以拾取和掉落物理对象

##### 2.7.4 物理推进器组件

<img src="https://d1iv7db44yhgxn.cloudfront.net/documentation/images/493f0cfe-917a-4e2f-b868-a23b419bc1bc/ue5_1-thruster.png" style="zoom:25%;" />

物理推进器组件（PhysicsThrusterComponent） 可以沿着 X 轴的负方向施加特定作用力。推力组件属于连续作用力，而且能通过脚本来自动激活、一般激活或取消激活。

推力组件的用途包括火箭（见下图）。它将持续施加作用力，将火箭向上推（因为推力部分位于火箭下方）。你可以用 阻挡体积（Blocking Volumes） ，限制受推力影响的组件的动作。

##### 2.7.5 径向力组件

​	径向力组件（RadialForceComponent） 用于发出径向力或脉冲来影响物理对象或可摧毁对象。与 PhysicsThrusterComponent 不同，这类组件只施加"发射后不管"类型的作用力，而且并不持续。

​	你可以使用这类组件来推动被摧毁对象（如爆炸物）的碎片。使用 RadialForceComponent 指定作用力和方向，当对象被摧毁时，你可以像下面的图示那样，沿着特定方向将碎片向外"推"

<img src="https://d1iv7db44yhgxn.cloudfront.net/documentation/images/65989139-2a64-44f4-82a1-d223ea4c8486/ue5_1-radial_force.png" style="zoom:25%;" />











### 3.物理场

物理场系统（Physics Field System）提供了一款工具，能在运行时直接影响特定空间中的Chaos物理系统的模拟效果。 你可以配置这些场，从而以各种方式影响物理模拟，例如在刚体上施加力、打破几何体集合群集、锚定或禁用破裂的刚体等。

此外，物理场系统还可以与其他引擎系统通信，例如Niagara和材质。 这些系统可以通过内置功能对特定地点的物理场求值，从而对物理场进行取样。

要设置场，请先创建场系统组件（Field System Component）蓝图，然后指定允许哪些系统查询你的场。 你可以将场配置为世界场（World Field），以允许材质和Niagara系统对该场进行取样。 此外，你还可以将场配置为Chaos场（Chaos Field），以允许几何体集合和其他物理对象与其交互。



### 4.破坏效果

<img src="https://dev.epicgames.com/community/api/documentation/image/80a5387b-c2aa-4f5a-8eac-b3179e3b543a?resizing_type=fit" style="zoom:50%;" />

Chaos破坏系统是虚幻引擎的一套工具集，可用于达到电影级质量的实时破坏效果。 除了炫目的视觉效果，该系统还优化了性能，让美术和设计师能够更好地控制内容创作。

Chaos破坏系统利用了几何体集合，这是一种根据一个或多个静态网格体编译的资产，包括蓝图中嵌套的静态网格体。 这些几何体集合可以破裂，从而达到实时破坏效果。

该系统使用直观的非线性工作流对破裂过程进行前所未有的控制。 用户可以在几何体集合的多个部分上创建多个破裂级别和选择性破裂，更好地进行美术控制。 用户还可以为各群集分别定义触发破裂所需的破坏阈值。

你可以使用连接图表进一步修改结构在受到破坏时的坍塌效果，以模拟真实世界的结构完整性。 Chaos破坏系统提供了"世界支持"，让你可以将几何体集合的某些部分设为符合运动学，不再需要使用锚点场。

Chaos破坏系统还提供了一套缓存系统，该系统能够追踪事件和变换。 该系统能逐帧捕获变换和事件，减少了破坏数据量，从而实现高效存储和播放功能。 该系统可在运行时流畅地播放复杂的破坏效果，最大程度地降低对性能的影响。 请务必注意，缓存模拟仍然可以交互，因为模拟的某些部分可能在交互时"激活"。

Chaos破坏系统还深度集成了虚幻引擎的Niagara粒子系统。 Niagara可读取Chaos破坏系统的分解事件（Break Events）和碰撞事件（Collision Events），从而在运行时生成粒子或修改现有粒子系统。

此外，Chaos破坏系统还集成了物理场系统。用户可使用物理场系统在运行时对指定空间区域内的Chaos物理系统的模拟施加影响。 物理场可以用多种方式来影响物理模拟，例如在刚体上施加力、打破几何体集合群集、锚定或禁用破裂的刚体。

### 5.载具

<img src="https://dev.epicgames.com/community/api/documentation/image/d029461f-47c6-4ebe-8aa2-1bd9c22235cd?resizing_type=fit" style="zoom:50%;" />

Chaos载具是虚幻引擎用来执行载具物理模拟的轻量级系统。

该系统为用户提供了更大的灵活性，因为它能逐车模拟任意数量的车轮。 你还可以配置任意数量的前进挡和后退挡，以进一步实现自定义效果。

Chaos载具支持复杂的载具模拟配置。 你可以添加任意数量的翼面，而这类翼面可在底盘的特定位置提供下行压力或升力。 这些翼面还可以模拟载具扰流板，甚至模拟飞行器的机翼或方向舵。 你可以通过滚动、俯仰和偏转等方法操纵这些控制表面。

在模拟运行时，你甚至可以完全关闭轮胎摩擦和发动机模拟。 此外，用户还可以配置任意数量的推力，这些推力可以被施加在特定位置，以推动和控制车辆的方向。

该系统支持虚幻引擎5中的异步物理模式，从而提高了模拟的确定性，确保了每次模拟都能获得可预测的结果。

搭配使用骨架网格体和一个或多个轮子蓝图，即可在引擎中呈现载具。 系统使用物理资产在物理编辑器中生成相应的碰撞几何体，使用动画蓝图自动生成轮子的转向和旋转动画。

最终，创建出包含骨架网格体、动画蓝图和载具控制点的轮式载具Pawn蓝图。

### 6.刚体动画节点和物理动画

<img src="https://dev.epicgames.com/community/api/documentation/image/f92b4b2b-3394-4830-960e-6533797cfed4?resizing_type=fit" style="zoom: 50%;" />

Chaos物理系统能在运行时为角色提供刚体模拟和物理动画。 该系统会使用物理资产编辑器为角色的骨架网格体设置可随角色动画而模拟的刚体。 与静态动画相比，这种方式能为被模拟的物体实现更逼真的运动效果。



### 7.布娃娃物理

<img src="https://dev.epicgames.com/community/api/documentation/image/9c908508-b7a4-497c-a5e3-0dba6957a547?resizing_type=fit" style="zoom:50%;" />

​	Chaos物理系统自带布娃娃物理系统，即一种为连接到骨架网格体的刚体实时制作动画（模拟）的系统。 此类系统常被用于模拟人形角色坠落时的动作。 下方示例为MetaHuman示例项目，该项目为玩家角色设置了布娃娃物理效果。







### 8.联机物理

![](https://dev.epicgames.com/community/api/documentation/image/bb085b8b-0adf-4b05-b789-92777f49966c?resizing_type=fit)

​	游戏中的联网或复制是指通过互联网连接在多台设备之间传递Gameplay信息的能力。 虚幻引擎具有健壮的网络框架，可帮助开发者简化多人游戏的创建。

​	联网物理是网络框架的一部分，使物理驱动型模拟能够在多人游戏环境中运行。 在虚幻引擎中，物理复制是指具有复制的移动效果且能模拟物理效果的Actor。 在Gameplay过程中，这些模拟在本地客户端（玩家的设备）内运行。

​	虚幻引擎提供了三种复制模式：

​	默认（Default）复制模式是虚幻引擎的旧版物理复制模式。 此模式在复制其移动且其根组件被设为模拟物理的Actor上激活。

​	预测性插值（Predictive Interpolation）复制模式专为服务器授权的Actor而设计。 工作原理是改变每个对象在客户端上的速度，以匹配当时对象在服务器上的速度，类似于默认模式。 然而，此模式旨在更好地处理客户端上的交互和局部物理改变，条件是该操作按预测的方式完成（预计在服务器和客户端上统一应用）。

​	重新模拟（Resimulation）复制模式专为服务器授权的Pawn和Actor而设计。 这种模式旨在通过客户端对物理效果的预测，在多人游戏中实现物理Pawn，并更精确地处理各种交互。 当客户端从服务器接收状态信息时，客户端会将其与历史记录中相应物理帧的缓存物理状态进行比较。 如果状态信息的差异足够大，就会触发物理重模拟。



### 9.Chaos可视调试器

> 简称：CVD

<img src="https://dev.epicgames.com/community/api/documentation/image/477e7097-8e10-4777-b488-70c2f04af544?resizing_type=fit" style="zoom:50%;" />

Chaos视觉调试器（CVD）是用于Chaos物理系统模拟的视觉调试工具。 该调试器提供了Chaos物理场景的图形视图，并配备了用于直观显示数据和分析模拟结果的各种工具。

CVD作为编辑器工具和运行时系统包含在虚幻引擎中，用于记录Gameplay过程中的物理模拟状态。 然后，它可以在工具内部重放这些模拟，并检查模拟的任何给定帧或子步骤的数据。



### 10.布料物理和机器学习布料模拟

Chaos布料可为游戏和实时体验提供精确、高性能的布料模拟。 该系统提供了丰富的用户控制功能和物理反应（比如风），可实现特定的美术构想。 此外，Chaos布料还配备了强大的动画驱动系统，可以使布料网格体变形，以匹配其父项的带动画骨架网格体。

Chaos布料参数对蓝图公开，从而在运行时实现了对布料模拟前所未有的精密控制。 用户可以根据特定用例的Gameplay条件来修改模拟参数。 例如，潜入到水下时，角色的服装将会作出不同的反应。

Chaos布料配备了一套强大的调试工具，例如可视化工具和调试控制台命令等，可简化调试并缩短迭代时间。 用户可以直观地看到主动布料模拟、运动学和动态粒子，甚至是模拟时间。

Chaos布料还提供机器学习布料模拟功能。 相比基于物理的传统模型，该系统能带来更高的保真度，并使用经过训练的、可实时使用的数据集，从而生成之前只能通过离线模拟来实现的结果。

Chaos布料面板节点编辑器是虚幻引擎5.3中新引入的工作流程。 该工作流程侧重于缩短迭代时间，让你能在引擎中以更灵活、非破坏性的方式创作Chaos布料。 此系统使用一种布料资产，该资产能存储在运行时生成和模拟布料所需的所有信息。 编译布料资产时，你可以从外部基于面板的数字内容创建（DCC）包中导入静态网格体，并转移蒙皮权重和遮罩。 创建资产后，你就可以通过Chaos布料组件将其应用于骨架网格体或静态网格体。



### 11.流体模拟

<img src="https://dev.epicgames.com/community/api/documentation/image/e7fdbb91-71f3-49f9-842a-b813a4e8513c?resizing_type=fit" style="zoom:50%;" />

​	虚幻引擎5提供了一组用于实时模拟2D和3D流体效果的工具。 这些系统使用基于物理的模拟方法，来产生逼真的效果，例如火焰、烟雾、云层、河流、飞溅物和海滩上的波浪。

这套工具旨在让美术师尽可能容易上手，通过利用模拟阶段、可复用的模块，和强大的数据接口，成为一个开放实验平台。

美术师只需修改几个参数，就能实时实现想要的效果。高级用户和研发工程师则可以深入其中，拆分模拟器来尝试新的算法。





### 12.毛发物理

<img src="https://dev.epicgames.com/community/api/documentation/image/82bae221-d858-46fe-b2e2-f741441ddecb?resizing_type=fit" style="zoom:50%;" />

虚幻引擎的毛发渲染和模拟系统使用基于发束的工作流程，以物理准确的运动渲染每一束毛发。 它支持美术师针对DDC包中创建的groom实时模拟和渲染数十万根甚至更多逼真的毛发。



### 13.Chaos Flesh

​	Chaos Flesh系统可在虚幻引擎中对可变形物体（软体）提供高质量的实时模拟。 与刚体模拟不同，软体的形状可以在模拟过程中根据软体的属性发生变化。

​	该系统支持使用各种参数模拟静态和骨骼网格体，使美术师能够对最终结果进行前所未有的控制。 我们设计该系统的目的主要专注于模拟骨骼动画期间角色的肌肉变形。

​	Chaos Flesh系统通过在运行时模拟低分辨率几何体，并利用高分辨率电影品质几何体离线模拟得到的缓存结果来实现高性能。



### 14.数据流图表系统

![](https://dev.epicgames.com/community/api/documentation/image/c1e3815d-6320-4417-b935-738dd5ce30d8?resizing_type=fit)

数据流图表系统是虚幻引擎编辑器内的一套基于节点的程序化资产生成环境。

创建数据流是为了在引擎中创建某些资产类型时优化迭代时间。 同一数据流图表可被多个资产使用，而且图表本身可根据源资产提供的输入而产出不同的结果。

数据流是一套通用系统，可适配各种物理资产类型，如Chaos布料、Chaos Flesh和几何体集合破裂等。 该系统被设计为可由C++开发者扩展。 开发者可以根据具体需求进一步调整系统。





### 11.参考文档

#### 11.1 官方文档

虚幻引擎物理系统总览：https://dev.epicgames.com/documentation/unreal-engine/physics-in-unreal-engine

物理形体参考指南：https://dev.epicgames.com/documentation/unreal-engine/physics-bodies-reference-for-unreal-engine

物理约束属性参考：https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-constraint-reference-in-unreal-engine
































