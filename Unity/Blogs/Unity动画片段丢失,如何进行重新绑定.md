---
title: Unity中动画片段丢失(AnimationClip),如何进行重新绑定
date: 2025-03-20 11:43:35
pic: /images/thumb/AnimatorClipMiss.png
comments: true
toc: true
categories:
- Engine
tags:
- Unity
---

# Unity动画片段丢失(AnimationClip),如何进行重新绑定

从外部导入的AnimationClip存在黄色丢失的missing提示，这时候不需要重新制作动画，只需要重新绑定动画即可。

<img src="/images/Unity/Programming/AnimatorClipMiss/1.png">

我们以第一条Intro1:Anchored Position(缺失!)为例
第一步：双击动画中的Intro1条目，可以查看片段存储该动画的对应路径(不太好点，多点几次)
可以看到动画中的路径设置的是Intro/Intro1，而实际在层级窗口中却是Intro/intro1，有一个大小写的区别，这就是造成物体丢失的原因

<img src="/images/Unity/Programming/AnimatorClipMiss/2.png">

第二步，将动画中的路径改为实际路径即可(当然反过来改游戏物体名称也是可以的)

<img src="/images/Unity/Programming/AnimatorClipMiss/3.png">

可以看到intro1:Anchored Position已经能够正常显示为灰色，在动画片段中也运行正常。
回想起曾经重新制作动画的痛苦和心酸，希望能够帮助到大家。

