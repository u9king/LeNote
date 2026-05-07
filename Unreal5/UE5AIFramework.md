# UE5AIFramework

> | 序号 | 课程                                         | 作者       | 链接                                                         | 备注                                                         |
> | ---- | -------------------------------------------- | ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
> | 1    | Introduction to Materials in Unreal Engine 5 | Mao Mao    | [Udemy](https://www.udemy.com/course/introduction-to-materials-in-unreal-engine-5/) | [b站](https://www.bilibili.com/video/BV1kvDsBAEGz?spm_id_from=333.788.videopod.episodes&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4) |
> | 2    | PBR材质底层原理讲解                          | 嗜睡的猫咪 | [b站](https://www.bilibili.com/video/BV1oY4y1t717)           |                                                              |

### 0.AIController控制器

需要创建一个AIController为父级的子蓝图





在角色的Pawn中进行设置，如果是生成角色需要更改Pawn类型



#### 0.1 自动放置Auto Possess Al

|            类型            |         翻译         |
| :------------------------: | :------------------: |
|          Disabled          |        不开启        |
|      Placed in World       |    AI已放置在世界    |
|          Spawned           |    AI生成在世界中    |
| Placed in World or Spawned | AI放置盒生成在世界中 |



### 1.行为树Behavior Tree

> 前缀：BT_

流控制



职能：用节流程图定义AI行为逻辑



执行顺序

- 从左到右
- 从上到下



核心定位：AI的决策流程树，用节点式的流程图定义 AI 行为逻辑（巡逻→发现敌人→追击→攻击 / 撤退）。

核心特点：按优先级执行节点，自带失败 / 成功分支，适合复杂 AI 的分层决策，是传统 AI 的主流方案。

依赖关系：必须绑定一个 Blackboard 来读写数据。



角色：决策者 / 逻辑树 行为树是 AI 的大脑逻辑”。它按照从左到右、从上到下的顺序，决定 AI 当前应该执行什么动作。

- 核心逻辑： 通过选择器（Selector）、序列（Sequence）和条件判断（Decorator）来控制流程。
- 通俗理解： 如果看到玩家，就追上去；如果追不到，就去巡逻；如果累了，就休息。” 这种逻辑分支全部写在这里。





#### 1.1 合成节点Composites



##### 1.1 选择节点 Select



##### 1.2 顺序节点 Sequence

特点：顺序执行，在无法执行处打断重启





##### 1.3 平行合成节点 Simple Parallel

直接/延迟











### 2.黑板

> 前缀：BB_



核心定位：AI 的全局数据黑板，存储 AI 需要的所有变量（目标位置、状态、是否看见敌人等）。

核心特点：作为 AI 逻辑的 共享内存”，BT、EQS、AI 控制器都能读写，是 AI 状态的统一数据源。

依赖关系：被 BT、EQS、AI 控制器引用，无它 BT 无法传递状态。



角色：内存 / 记忆库 黑板是专门用来存储数据的地方，行为树本身不记事，它需要去黑板里读写数据。

- 核心内容： 存放变量，如 `SelfActor`（自己）、`TargetActor`（目标玩家）、`MoveToLocation`（目的地坐标）、`IsAlerted`（是否警觉）。
- 通俗理解： 它是大脑的便利贴”。行为树发现玩家后，会在黑板上写下玩家的位置，然后下一个任务会从黑板读取这个位置并走过去。







为什么黑板没法直接获取变量内容啊？

LiteralName字面名称是什么？





### 10.Environment Query

- 核心定位：AI 的环境探测工具，帮 AI 在场景中自动筛选、评估出符合条件的位置（比如找掩体、找离玩家最近的射击位）。
- 核心特点：可配置生成点 + 评分规则，比如 找离玩家 10 米内、无遮挡、能躲子弹的位置”，结果可直接写入 Blackboard 给 BT 用。
- 依赖关系：常被 BT 节点调用，结果写入 Blackboard。





角色：感官 / 空间分析 EQS 负责处理复杂的物理空间计算。它会向周围发射大量探测点”，并根据你设定的规则给这些点打分。

- 应用场景： 找一个既能看到玩家，又离我最近，且有掩体挡住玩家子弹的位置。”
- 通俗理解： 它是 AI 的眼睛和直觉”。它不负责决定干什么，只负责在地图上成百上千个点中，选出那个最符合要求的黄金位置”。

### 

### 11.状态树StateTree

核心定位：UE5 新出的状态机工具，是 Behavior Tree 的轻量化替代方案，用状态和过渡条件来定义 AI 逻辑。

核心特点：比 BT 更轻量、性能更好，逻辑更直观（适合简单到中等复杂度的 AI），支持蓝图 / 原生混合开发，部分场景可替代 BT+Blackboard 的组合。

依赖关系：可独立运行，也能和 Blackboard、EQS 配合使用。





角色：高效状态机 / 下一代决策框架 这是 UE5 重点推行的新特性。它结合了传统状态机（FSM）的清晰和行为树的层次感，但性能更高。

- 核心优势： 比行为树更轻量，逻辑更紧凑。它非常适合处理状态切换”，比如从闲逛状态”切换到战斗状态”。
- 通俗理解： 如果说行为树是在森林里找路，状态树就像是一个高性能的调度台。在处理大量简单的 AI（如群体 NPC）或需要极高逻辑严密性的动作切换时，StateTree 往往比 BT 更优雅。









### 3.组件AIPerception

OnTargetPerceptionUpdated

##### 3.1 Senses Config感官配置

- AI视觉配置AI Sight Config
  - 视线半径Sight Radius
  - 丢失视野半径Lose Sight Radius
  - 周边视觉半角PeripheralVisionHalfAngle
  - 按阵营检测Dectection By Affiliation
    - 检测敌人Detect Enemies
    - 检测中立单位Detect Neutrals
    - 检测友军Detect Friendlies





### 4.组件AIperceptionTimuliSource

### 5.任务Task

> 前缀：BTT_ / BTTask_







每个任务都要单独创建？

### 6.装饰器Decoration





#### 7.服务Service

Task 只负责「流程、条件、触发」；

服务只负责「改属性、改状态」，各司其职。



为什么不写在 Task 里？（核心区别）

如果你把“寻找目标”的逻辑写在 `Task_Attack` 里，会产生以下几个问题：

A. 生命周期与并发性

- Task： 只有当行为树运行到这个 Task 时，逻辑才会执行。如果 Task 结束了（比如攻击动画播完了），逻辑就停止了。
- Service： 只要 Service 所在的分支（Branch）是激活状态，它就会一直运行。
    - *场景：* 当 AI 在执行“移动到目标”这个 Task 时，挂在父节点上的 Service 可以持续更新目标位置。如果目标跑了，Service 更新黑板，左边的装饰器（条件）就能立刻发现并中止当前的移动 Task。

B. 逻辑复用（DRY原则）

如果你的 AI 有“远程攻击”、“近战攻击”、“追逐”三个 Task，而它们都需要实时获取目标位置：

- 写在 Task 里： 你得在三个 Task 类里写三遍同样的获取逻辑。
- 写在 Service 里： 你只需要在一个 Service 里写一次，然后把它挂在这些 Task 共同的父节点上。

性能优化（性能平衡）

行为树的 Task 往往是每帧执行（Tick）的，如果每个 Task 都去执行复杂的搜索逻辑（比如 `GetAllActorsOfClass`），性能消耗会很大。

- Service 的优势： 它可以设置时间间隔（Interval）和\偏差（Random Deviation）。
    - 你可以设置 Service 每 0.5 秒更新一次数据，而不是每帧更新。这在处理大量 AI 角色时，能极大地平滑 CPU 的压力。











### 7.AIController

自带的PathFollowingComponent是干什么的？



### 8.显示AIPerception范围

> 按下：‘ + 4

可以显示AIPerception的视野范围



感官丢失需要的是在On Target Perception Updated输出的AlStimulus中Bbreak找到Successfully Sensed，这里面能拿到当前感觉是否起作用









### 9.导航网格边界体积NavMeshBoundsVolume

> 快捷键：P 开关显示



























