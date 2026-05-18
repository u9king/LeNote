# UE5特效

> | 序号 | 课程                                         | 作者     | 链接                                                         | 备注                                                         |
> | ---- | -------------------------------------------- | -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
> | 1    | Introduction to Materials in Unreal Engine 5 | Mao Mao  | [Udemy](https://www.udemy.com/course/introduction-to-materials-in-unreal-engine-5/) | [b站](https://www.bilibili.com/video/BV1kvDsBAEGz?spm_id_from=333.788.videopod.episodes&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4) |
> | 2    | 虚幻（UE5）角色动画基础                      | 亿峥游戏 | [b站](https://www.bilibili.com/video/BV1ow411U7Eh)           |                                                              |

### 0.未分类

#### 0.1 重置世界粒子

快捷键：选中 + /





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





### 









### 2.模块

#### 3.0 属性 Propertites

本地空间LocalSpace：可以让粒子跟随发射器移动





#### 3.1 生成速率 Spawn Rate



#### 3.2 添加速度 Add Velocity



#### 3.3 重力 Gravity Force



#### 3.4 碰撞 Collision

弹力->Resitution可以设置为0.2





#### 3.5 缩放精灵图大小 Scale Sprite Size

快捷键：

- 鼠标中键添加关键帧
- 全选按住2可以平滑曲线



#### 3.6 发射状态 EmitterState



#### 3.7 精灵渲染器 SpriteRender



#### 3.8 形状位置 Shape Location













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













### 4.专有名词

#### 4.1 体积雾

#### 4.2 湍流





























