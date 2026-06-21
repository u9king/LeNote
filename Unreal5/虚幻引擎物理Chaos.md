# 虚幻引擎物理Chaos

> 别称：Chaos物理系统
>
> | 序号 | 课程                                                         | 作者         | 链接                                                         | 备注                                                         |
> | ---- | ------------------------------------------------------------ | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
> | 1    | Introduction to Materials in Unreal Engine 5                 | Mao Mao      | [Udemy](https://www.udemy.com/course/introduction-to-materials-in-unreal-engine-5/) | [b站](https://www.bilibili.com/video/BV1kvDsBAEGz?spm_id_from=333.788.videopod.episodes&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4) |
> | 2    | 虚幻（UE5）角色动画基础                                      | 亿峥游戏     | [b站](https://www.bilibili.com/video/BV1ow411U7Eh)           |                                                              |
> | 3    | UE5 Series: IK Rig Retargeting in UNREAL Engine 5.3 and 5.4  | SARKAMRI     | [Youtube](https://www.youtube.com/watch?v=f6tm7NXEs0I)       |                                                              |
> | 4    | How To Do Animation Retargeting At Runtime In Unreal Engine 5 (Tutorial) | Matt Aspland | [Youtube](https://www.youtube.com/watch?v=iv5mG_cxycs)       |                                                              |

### 0.待分类

##### 1.物理资产

##### 2.刚体

##### 3.约束

##### 4.物理材质 Physical Material

可以分配给刚体，摩擦力弹力

##### 5.模拟

##### 6.物理体

物理体网格体：由基元集合定义的

##### 7.基元组件 Primitive Component

简单碰撞确实是依托于基元组件来存储和发生作用的。

在 UE5 中，只要一个组件要么能被看见，要么能被碰撞，它在底层就一定继承自 UPrimitiveComponent。







### 1.Chaos物理系统

Chaos物理系统是虚幻引擎提供的轻量级物理模拟解决方案。 该系统包括以下主要功能：

- 破坏效果
- 联网物理
- Chaos可视调试器
- 刚体动力学
- 刚体动画节点和物理动画
- 布料物理和机器学习布料模拟
- 布娃娃物理
- 载具
- 物理场
- 流体模拟
- 毛发物理
- 血肉模拟

#### 1.1 帧率

> 别称：子步，迭代
>
> 路径：Physics -> 帧率





物理引擎包括：PhysX，Chaos等



### 2.刚体

> 别称：物理体，Chaos物理体

刚体是指在运动中和受到力的作用后，形状和大小不变，而且内部各点的相对位置不变的物体

Chaos物理系统提供了许多种刚体动力学功能。 这包括碰撞响应、物理约束以及阻尼和摩擦。 此外，它还提供了异步物理模拟和联网物理。

- 质量
- 线性阻尼：摩擦力
- 角阻尼：旋转摩擦力
- 约束
  - 







### 3.碰撞

> 命令行：
>
> - show Collision：显示所有碰撞模型和阻挡体积
> - stat game：显示关于碰撞所用时间的统计信息

原则：最简原则：重叠与碰撞，结果是重叠

#### 2.1 碰撞外形

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

#### 2.2 碰撞通道 

##### 2.2.1 对象通道 Object Channels

> 路径：Edit Menu -> Project Settings -> Collision



| 对象类型     | 翻译     | 说明                                              |
| ------------ | -------- | ------------------------------------------------- |
| WorldStatic  | 世界静态 | 应用于不移动的任意Actor                           |
| WorldDynamic | 世界动态 | 用于受到动画或代码的影响而移动的Actor类型；运动学 |
| Pawn         | 棋子     | 任何由玩家控制的实体的类型都应为Pawn              |
| PhysicsBody  | 物理物体 | 由于物理模拟而移动的任意Actor                     |
| Vehicle      | 载具     | 载具的默认类型                                    |
| Destructible | 可破坏   | 可破坏物网格体的默认类型                          |



##### 2.2.2 追踪通道 Trace Channels

> 路径：Edit Menu-> Project Settings -> Collision 

默认有两个追踪通道可视性Visibility和摄像机Camera。

你可能需要一种特殊激光，它需要能够穿过你无法透视或无法让摄像机穿过的特殊不透明对象。你添加自己的自定义追踪响应通道。







#### 2.9 碰撞复杂度 Collision Complexity

- 简单和复杂 Simple And Complex：默认状态。引擎同时为它创建简单和复杂两套外壳，各司其职。
- 将简单作用于复杂 Use Simple Collision As Complex： 无论任何时候，都只使用那个绿色的简单外壳。最推荐、性能最好。
- 将复杂作用于简单 Use Complex Collision As Simple：全部使用复杂
  - 无法使用物理
  - 仅使用于边界碰撞
  - 精确的射线检测









#### 1.2 Trace 射线检测



#### 1.3 物理约束



#### 1.4 物理组件



### 4.约束

> 组件：PhysicsConstraintActor

基础:距离/角度控制
状态:Rest Target
高级:Soft Plastic Break
校正:Projection



- 









### 2.物理场



### 3.物理材质



### 4.刚体动画节点和物理动画



### 5.布娃娃物理







#### 9.物理资产 Physics Asset

> 缩写：PA，RBAN
>
> 快捷键：Simulate模式下，按住Control + 右键可以控制玩家



#### 10.组件

##### 10.1 物理约束 PhysicsConstraint

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







#### 11.参考文档

##### 11.1 官方文档

物理：https://dev.epicgames.com/documentation/unreal-engine/physics-in-unreal-engine

物理形体参考：https://dev.epicgames.com/documentation/unreal-engine/physics-bodies-reference-for-unreal-engine

物理约束属性：https://dev.epicgames.com/documentation/en-us/unreal-engine/physics-constraint-reference-in-unreal-engine



































































