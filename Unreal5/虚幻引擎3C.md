# UE5动作

> | 序号 | 课程                                         | 作者     | 链接                                                         | 备注                                                         |
> | ---- | -------------------------------------------- | -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
> | 1    | Introduction to Materials in Unreal Engine 5 | Mao Mao  | [Udemy](https://www.udemy.com/course/introduction-to-materials-in-unreal-engine-5/) | [b站](https://www.bilibili.com/video/BV1kvDsBAEGz?spm_id_from=333.788.videopod.episodes&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4) |
> | 2    | 虚幻（UE5）角色动画基础                      | 亿峥游戏 | [b站](https://www.bilibili.com/video/BV1ow411U7Eh)           |                                                              |

### 1.Move移动

#### 1.1 初版

![](https://gitee.com/u9king/ImageHostingService/raw/master/Unreal5/UE5%E5%8A%A8%E4%BD%9C/Move1.png)



Get Forward Vector

$$Forward.X = \cos(Yaw) \times \cos(Pitch)$$

$$Forward.Y = \sin(Yaw) \times \cos(Pitch)$$

$$Forward.Z = \sin(Pitch)$$

球面坐标系





对于绕 $Z$ 轴（Yaw）旋转 $\theta$ 角的矩阵公式为：
$$
\vec{V}= R_z(\theta) \times \vec{V} \\
当\theta=90°，即绕Z轴旋转90° \quad
\vec{V} = \begin{bmatrix} 0 & -1 & 0 \\ 1 & 0 & 0 \\ 0 & 0 & 1 \end{bmatrix} \times \begin{pmatrix} 1 \\ 0 \\ 0 \end{pmatrix} = \begin{pmatrix} 0\\ 1 \\ 0 \end{pmatrix}
$$




#### 1.2 目前![](https://gitee.com/u9king/ImageHostingService/raw/master/Unreal5/UE5%E5%8A%A8%E4%BD%9C/Move2.png)

