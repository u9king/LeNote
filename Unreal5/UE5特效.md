# UE5特效

> | 序号 | 课程                                         | 作者     | 链接                                                         | 备注                                                         |
> | ---- | -------------------------------------------- | -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
> | 1    | Introduction to Materials in Unreal Engine 5 | Mao Mao  | [Udemy](https://www.udemy.com/course/introduction-to-materials-in-unreal-engine-5/) | [b站](https://www.bilibili.com/video/BV1kvDsBAEGz?spm_id_from=333.788.videopod.episodes&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4) |
> | 2    | 虚幻（UE5）角色动画基础                      | 亿峥游戏 | [b站](https://www.bilibili.com/video/BV1ow411U7Eh)           |                                                              |

#### 1.NiagaraSystem

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

#### 2.阶段：Particle Spawn粒子生成

只在粒子出生那一帧执行一次

##### 2.1 模块：Initialize Particle初始化粒子

##### 2.2 模块：Shape Location

##### 2.3 模块：Add Velocity

#### 3.阶段：Particle Update粒子更新

在粒子的生命周期内每一帧都执行，如 Gravity Force







#### 6.阶段：Render渲染

##### 6.1 模块：SpriteRenderer精灵渲染器

- subuv：拆分精灵图







