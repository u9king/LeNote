# UE5特效

> | 序号 | 课程                                         | 作者     | 链接                                                         | 备注                                                         |
> | ---- | -------------------------------------------- | -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
> | 1    | Introduction to Materials in Unreal Engine 5 | Mao Mao  | [Udemy](https://www.udemy.com/course/introduction-to-materials-in-unreal-engine-5/) | [b站](https://www.bilibili.com/video/BV1kvDsBAEGz?spm_id_from=333.788.videopod.episodes&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4) |
> | 2    | 虚幻（UE5）角色动画基础                      | 亿峥游戏 | [b站](https://www.bilibili.com/video/BV1ow411U7Eh)           |                                                              |

### 0.未分类

#### 0.1 重置世界粒子

快捷键：选中 + /

#### 0.2 游戏模式

快捷键：G

#### 0.3 直线变曲线

快捷键：全选 +  2

#### 0.4 尼亚加拉模块编辑器 Niagara Module Script





### 1.NiagaraSystem

> 定义：粒子系统
>
> 别称：奶爪，出处尼加拉瓜大瀑布
>
> 变量名：NS_XXX



内置模板

|         名称         |   翻译   | 常用于 |             解释              |
| :------------------: | :------: | :----: | :---------------------------: |
|       Fountain       |   喷泉   |        | 循环喷泉 / 持续循环的喷泉水花 |
| Hanging Particulates | 悬浮微粒 |        | 悬浮微粒在矩形区域内盘旋/旋转 |
|                      |          |        |                               |



02.Niagara 堆栈结构逻辑

#### 1.阶段：Emitter Spawn

#### 2.阶段：Particle Spawn 粒子生成

只在粒子出生那一帧执行一次

##### 2.1 模块：Initialize Particle初始化粒子

##### 2.2 模块：Shape Location

##### 2.3 模块：Add Velocity

#### 3.阶段：Particle Update 粒子更新

在粒子的生命周期内每一帧都执行，如 Gravity Force







#### 6.阶段：Render渲染

##### 6.1 模块：SpriteRenderer精灵渲染器

- subuv：拆分精灵图










### 2.模块

### 3.模块 Modules
#### 3.0 系统属性 System Propertites

##### 3.0.1 系统 System

|           属性           |     翻译     | 备注                   |
| :----------------------: | :----------: | ---------------------- |
|       Effect Type        |              |                        |
|       Fixed Bounds       |   固定边界   | 整个系统的固定边界     |
| Initial Streaming Bounds |              |                        |
|       Determinism        |    确定性    |                        |
|       Warmup Time        |   预热时间   |                        |
|  Fixed Tick Delta Time   | 固定时间间隔 | 在总线执行，而不是异步 |

##### 3.0.2 渲染 Rendering

- 阴影 Default Cast Shadows
- 贴花 Default Receives Decals
- 渲染深度 Default Render CustomDepth

##### 3.0.3 性能 Performance

- Bake Rapid Iteration Parameters During Edit
- Bake Rapid Iteration Parameters
- Compress Attributes
- Trim Attributes During Edit
- Trim Attributes
- Ignore Particle Reads for Attribute Trim
- Disable Debug Switches During Edit
- Disable Debug Switches
- Require Current Frame Data
- Allow System State Fast Path
- 最大池数量 Max Pool Size
- Pool Prime Size

##### 3.0.4 系统状态 System State

- Inactive Response
- 循环模式 Loop Behavior
- 循环时长 Loop Duration
- Recalculate Duration Each Loop
- 循环延迟 Loop Delay



#### 3.1 属性 Propertites

##### 3.1.1 本地空间 Local Space

是否跟随发射器移动，是储存在局部空间还是世界空间

##### 3.1.2 确定性 Determinism

粒子生成的确定性

勾选会出现随机种子

可以在时间轴中拖动显示

适用场景

- 电影特效
- 特效制作过程中开启，完成后关闭

##### 3.1.3 插值生成模式 Interpolated Spawn Mode

开启插值生成，会在第一帧同时执行生成和更新

解决粒子在重复利用时生成闪烁的问题，也就是生成第一帧大，然后在更新时变小的问题

##### 3.1.4 模拟目标 Sim Target

选择CPU和GPU的地方

##### 3.0.5 计算边界模式 Calculate Bounds Mode

GPU模拟需要固定边界，超出这个小盒子就会被剔除遮挡

##### 3.0.6 需要持久化ID Requires Persistent IDs

不开ID可重用

##### 3.0.7 发射器依赖关系 Emitter Dependencies

本地空间LocalSpace：可以让粒子跟随发射器移动

##### 3.0.8 阶段按钮 Stage

高级部分

Event Handler时间处理

Generic Simulation Stage 通用模拟阶段



#### 3.1 生成[模块] Spawning [Module]

##### 3.1.1 生成速率 Spawn Rate

- 生成速率 Spawn Rate：单位每秒多少个粒子
- 生成概率Spawn Probability ：这一帧是否生成粒子
- 生成分组 Spawn Group ：

##### 3.1.2 瞬间爆发式生成 Spawn Burst Instantaneous

- 生成数量 Spawn Count：
- 生成时间 Spawn Time：
- 生成概率 Spawn Probability：这一帧是否生成粒子
- 循环次数限制 Loop Count Limit：

##### 3.1.3 每单位生成 Spawn Per Unit

- 生成间距 Spawn Spacing：每移动XX单位/厘米，生成一次粒子
- 最大移动阈值 Max Movement Threshold：超过阈值不生成
- 移动容差 Movement Tolerance：超过这个距离才生成，微小左右移动不生成
- 生成概率 Spawn Probability：这一帧是否生成粒子

##### 3.1.4 每帧生成 Spawn Per Frame

- 生成数量 Spawn Count
- 生成概率 Spawn Probability
- 生成开关 Spawn
- 生成分组 Spawn Group

新功能，可能有些问题，和生成速率基本一样，不过是会依赖帧率，帧率越高生成粒子越多

##### 3.1.5 从其他发射器生成 Spawn Particles from Other Emitter

配合发射器解释

##### 3.1.6 从网格中生成粒子 Spawn Particles in Grid

- X数量 X Count：X轴方向上的数量
- Y数量 Y Count：Y轴方向上的数量
- Z数量 Z Count：Z轴方向上的数量
- 生成时间 Spawn Time：





##### 3.1.7 从顶点动画工具变形目标生成 Spawn MS Vertex Animation Tools Morph Target

高级功能



#### 3.2 位置[模块] Location [Module]

##### 3.2.1 形状位置 Shape Location

- 图元形状 Shape Primitive  
    - 球体 Sphere
        - 球体半径 Sphere Radius
        - 球体分布 Sphere Distribution 
            - 随机 Random：随机分布在球体
            - 直接 Direct：基于物体的UV ,0.5是球的赤道，配合SpawnRate可以生成圆环，
            - 均匀 Uniform：仅对爆发生成器可用，用斐波那契函数生成
        - 球体表面分布 Sphere Surface Distribution：1为全部在表面，0为完整体积
        - 半球分布 Hemisphere Distribution：0~1之间，通过网格体的UV完成的
    - 圆柱体 Cylinder
        - 圆柱体中心 Cylinder Height Midpoint：中心点0为底，1为顶
        - 表面厚度 Surface Only Band Thickness：0为表面，与球体相反
        - 半圆生成X Hemicircle X：在X方向仅生成一半
        - 半圆生成Y Hemicircle Y：在Y方向仅生成一半
        - 开启旋转体轮廓 Enable Lathe Profile：用曲线控制圆柱体宽度
    - 立方体/平面 Box/Plane
        - 表面厚度 Surface Only Band Thickness：0为表面，与球体相反
    - 圆环 Torus
        - 大半径 Large Radius：圆环外半径
        - 控制柄半径 Handle Radius：圆环的管道的粗细
        - 表面分布  Surface Distribution：1为全部在表面，0为完整体积
    - 环 / 圆盘 Ring/Disc
        - 圆盘覆盖率Disc Coverage：0是外边缘，1是全部
    - 圆锥 Cone
        - Cone Surface Distribution 圆锥表面分布：0是外边缘，1是全部

- Distribution 分布
- Transform 变换

##### 3.2.2 网格位置 Grid Location

需要配合Spawn Particles in Grid来使用

- 尺寸定义 Dimensions Definition
    - 每格内边距 Padding Per Cell：
        - XYZ尺寸 XYZ Dimensions：定义网格间距
    - 边界框尺寸 Bounding Box Size：
- 标准化偏移 Normalize Offsets ：开启可以后输入网格偏移的格子数量，不开启输入实际距离
- 网格内随机偏移 Randomize Placement Within Cell
- 偏移 Offset

##### 3.2.3 抖动位置 Jitter Position

- 抖动值 Jitter Amount：专业术语，形容抖动幅度
- 抖动偏移 Jitter Offset：
- 抖动延迟 Jitter Delay：需要在更新阶段才能起作用，每隔一段时间抖动一次

##### 3.2.4 系统位置 System Location

如果没有设置位置，默认就会使用System Location

- 偏移Offset：

##### 3.2.5 静态网格位置 Static Mesh Location

- 在Mesh中
    - 必要：打开允许CPU访问，可以通过Fix按钮自动修复
    - 可选：针对非均匀的静态网格，可以开启均匀分布采样，和GPU采样
        - 支持均匀分布采样 Support Uniformly Distributed Sampling
        - 支持GPU均匀分布采样 Support Gpu Uniformly Distributed Sampling
- 网格采样种类 Mesh Sampling Type：
    - 三角形 Triangles
    - 插槽 Sockets
    - 顶点 Vertices
- 允许材质槽 Allowed Material Slots：需要配合网格采样的过滤模式才能使用，需要网格本身含有不同材质槽，举例Triangle Sampling Mode：Random (Section Filter)
- 源模式 Source Mode
    - 默认 Default
    - 源 Source
    - 附加到父级 Attach Parent：可以根据父级的Mesh来生成粒子
    - 只有默认网格 Default Mesh OnlyMesh
    -  网格参数绑定 Mesh Parameter Binding

##### 3.2.6 骨骼网格位置 Skeletal Mesh Location

- 源模式 Source Mode
    - 默认 Default
    - 源 Source
    - 附加到父级 Attach Parent：可以根据父级的Mesh来生成粒子
- 采样区域Sampling Regions:需要现在SKM上设置SampleRegion这里对应填入名称，还需要开启过滤采样区域
- 网格采样种类 Mesh Sampling Type
    - 骨骼(骨头) Skeleton (Bones)：仅在骨骼处生成粒子
    - 骨骼(插槽) Skeleton (Sockets)：
    - 骨骼(骨头和插槽) Skeleton (Bones and Sockets)：
    - 表面(三角形) Surface (Triangles)：在骨架表面生成粒子
    - 表面(顶点) Surface (Vertices)：

#### 3.3 速度[模块] Velocity [Module]

##### 3.3.1 添加速度 Add Velocity

- Velocity Mode 速度模式
    - 线性 Linear
        - Velocity 速度
        - Velocity Speed Scale 速度缩放
    - 从点 From Point
        - 原点偏移 Origin Offset：调整朝某一方向移动
    - 圆锥中 In Cone

注：速度可以在Update中实时修改，来模拟越来越大

##### 3.3.2 继承速度 Inherit Velocity

- 继承速度值缩放 Inherited Velocity Amount Scale ：描述速度的向量，可以归一化
- 继承速度限制 Inherited Velocity Speed Limit
- 源速度阈值 Source Speed Threshold

##### 3.3.3 缩放速度 Scale Velocity

在Update中可以调用

归一化范围内的速度设置

#### 3.4 力[模块] Force [Module]

##### 3.4.1 应用初始力 Apply Initial Forces

##### 3.4.2 加速度力 Acceleration Force

##### 3.4.3 重力 Gravity Force

与加速度力相同，就是一个特定值的加速度

##### 3.4.4 线性力 Linear Force

- 力 Force
- 坐标空间 Coordinate Space
- 写入参数 Write to Intrinsic Parameters

##### 3.4.5 卷曲噪声力 Curl Noise Force

- 噪声强度 Noise Strength
- 噪声频率 Noise Frequency
- 卷曲噪声遮罩 Mask Curl Noise ： 在角度内才施加力
    - 卷曲噪声锥形遮罩角度 Curl Noise Cone Mask Angle

##### 3.4.6 涡流力 Vortex Force

- 涡流力值 Vortex Force Amount：涡流强度
- 涡流轴 Vortex Axis：
- 涡流轴坐标空间 Vortex Axis Coordinate Space：本地，世界，模拟
- Vortex Influence Position
- Write to Intrinsic Parameters

##### 3.4.7 风力 Wind Force

- 风速 Wind Speed
- 风速缩放 Wind Speed Scale
- 湍流模式 Turbulence Mode ：卷曲噪声
    - 直接设置 Direct Set
    - Speed Range
- 启用地面遮罩 Use Ground Mask：靠地平面和CPU实现









##### 3.4.7 解决力和速度 Solve Forces and Velocity

基于物理力，引擎时间差 算出物体该移动多远

##### 3.4.8 限制力 Limit Force

限制力的最大不会超过





#### 3.4 碰撞 Collision

弹力->Resitution可以设置为0.2





#### 3.5 缩放精灵图大小 Scale Sprite Size

快捷键：

- 鼠标中键添加关键帧
- 全选按住2可以平滑曲线



#### 3.6 发射状态 EmitterState





#### 3.9 粒子状态 Particle State



#### 3.11 空气阻力 Aerodynamic Drag



#### 3.13 缩放精灵大小 Scale Sprite Size



#### 3.14 缩放颜色 Scale Color



#### 3.15 更新网格旋转 Update Mesh Orientation

注：需要在重力模块之后



#### 3.16 阻力 Drag

注：内部实现可能就是减小重力的值



#### 3.17 直接设置新或者已有参数 Set new or existing parameter directly

可以用底层参数复现模块的工作原理

斜体重命名









#### 3.99 渲染器 Render



#### 3.99 渲染器 Render







### 3.后处理PostProcess

#### 3.0 镜头

#### 3.1 动态景深 Mobile Depth of Field

#### 3.2 泛光 Bloom

#### 3.3 曝光 Exposure

#### 3.4 色差 Chromatic Aberration

#### 3.5 脏污遮罩 Dirt Mask

#### 3.6 相机 Camera

#### 3.7 局部曝光 Local Exposure

#### 3.8 镜头光晕 Lens Flares

#### 3.9 图片效果 Image Effects

#### 3.10 景深Depth of Field










##### 3.99.1 精灵渲染器 SpriteRender
##### 3.99.2 网格渲染器 MeshRender



### 4.专有名词

#### 4.1 体积雾

体积雾是指数级大气Exponential Height Fog组件的一种属性。与传统的平面雾不同，它允许光线在三维空间中被散射和吸收，从而产生可见的光束（丁达尔效应/耶稣光）。

- 特效中的作用： 在特效制作中，体积雾能让粒子具有厚度感。例如，当你制作一个发光的魔法阵时，开启体积雾后，光芒会呈现出穿透空气的实体光柱感，而不是单纯的贴图发光。
- 交互性： 粒子系统可以通过Volume 材质与体积雾互动，甚至可以利用 Niagara 粒子写入体积纹理，实现烟雾随风飘动并产生阴影的效果。



意思：真正三维的雾，不是贴在平面上的假雾。

作用：计算空间中每一点的雾气密度与光照，能产生丁达尔效应（光线穿雾可见光束），并与灯光、阴影、Lumen 实时互动。

特效用途：做雾气、尘埃、光束、火焰 / 烟雾的大气感。







#### 4.2 湍流

湍流（通常在Niagara的Curl Noise Force模块中体现）是一种模拟流体或空气不规则运动的物理现象。

- 核心逻辑： 它利用数学算法（如柏林噪声或卷曲噪声）生成一个随机但平滑的方向场。粒子进入这个场后，会像落叶在旋涡中旋转一样，产生自然的摇摆、翻转和无规则路径。
- 应用场景： 它是告别“直线运动”的关键。制作升腾的烟雾、飞舞的萤火虫或魔法能量流时，加入湍流能瞬间增加视觉上的随机感和真实度。



意思：流体（烟 / 火 / 雾）里的不规则漩涡、乱流细节。

作用：让粒子 / 烟雾运动不呆板，有细碎扰动、旋转、撕扯感，更像真实流体。

特效用途：给火焰、浓烟、爆炸加动态紊乱细节，避免平滑得像果冻。







#### 4.3 动态输入

动态输入是 Niagara 系统中极其强大的功能，它允许你使用简单的数学逻辑或变量，实时改变某个参数的值，而无需去写复杂的 C++ 代码。

- 工作原理： 它相当于一个“转换器”。比如，你原本要在“粒子颜色”里填入红色，但你改用动态输入，设置一个 Linear Interpolate线性插值，让颜色根据粒子的 NormalizedAge归一化寿命从红色过渡到蓝色。
- 优势： 它是模块化的。你可以通过动态输入轻松实现：让粒子速度随时间加快、旋转角度随位置偏移、或者根据玩家与特效的距离来缩放粒子大小。





- 意思：特效参数不写死，运行时能随游戏逻辑 / 外部数据实时变化。
- 作用：比如：血量越低→火焰越弱；速度越快→拖尾越长；受击→瞬间爆光。
- 特效用途：让特效和游戏玩法联动，不是固定动画，更有交互感。



#### 4.4 硬编码 Hardcode

含义：把参数值/逻辑直接写死在代码里，而不是使用变量或配置文件。

对比

- 硬编码： 在粒子生成数量里直接写100。如果想改成 200，必须打开文件找到这个数字去改。
- 参数化： 将数量提取为一个变量 SpawnCount。你可以在变量中直接修改它，而不需要进入文件内部。

评价： 在初学阶段，硬编码可以快速看到效果，但在实际商业项目中，硬编码被视为一种“坏习惯”，因为它会导致系统难以维护和复用。





#### 4.5 自定义深度

