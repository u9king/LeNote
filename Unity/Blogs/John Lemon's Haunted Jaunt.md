---
title: John Lemon's Haunted Jaunt
date: 2024-01-10 19:40:56
pic: /images/thumb/John Lemon's Haunted Jaunt.png
comments: true
toc: true
categories:
- Games
tags:
- Unity
---

# John Lemon's Haunted Jaunt

#### 1. 游戏背景

​	John Lemon 的 Haunted Jaunt3D潜行游戏，玩家扮演John Lemon是一名冒险者。有一天，他接到了一家调查机构的委托，任务是前往一个废弃的城堡大厅，调查那里关于一个失踪家族的传说。

​	这个城堡曾是一个富裕且神秘的家族的住所，直到某天，家族成员在进行一场神秘的仪式时突然失踪。传说他们的灵魂被困在城堡中，而城堡因此受到了诅咒。约翰·莱蒙的任务是解开这个家族的神秘和庄园的诅咒，带着冒险的心情，他踏上了这次 "鬼魂之旅"。

​	进入城堡大厅后，约翰发现这里的环境异常阴森，大厅内散发着陈旧的气息，家具散乱，墙壁上的画作变得模糊而诡异。在石像的激光扫射和鬼的探测下，约翰必须小心翼翼地穿越大厅，寻找隐藏的谜团，解开庄园的秘密。

​	约翰知道，一旦他能够成功抵达大厅中的最终安全区域，就能解除这个家族的诅咒，让失踪的灵魂得到安宁。他的勇气将会揭示庄园的过去，同时也可能让他发现一些关于自己的意外真相。

#### 2. 游戏内容

<img src="/images/Unity/Games/John%20Lemon's%20Haunted%20Jaunt/1.png">

<img src="/images/Unity/Games/John%20Lemon's%20Haunted%20Jaunt/2.png">

#### 3. 核心机制

- **演示视频**

<iframe src="//player.bilibili.com/player.html?bvid=BV18T411P7Ec" width="580px" height="320px"></iframe>

- **潜行与躲藏：**
   - 玩家需要运用潜行技巧，避开石像的激光扫射和鬼的探测。
   - 在废弃的城堡大厅中，玩家可以利用环境中的柱子、家具和墙壁来隐藏自己，避免被敌人发现。
- **观察和计划：**
   - 玩家需要仔细观察石像的激光运动模式和鬼的探测范围，制定穿越计划。
   - 在探索过程中，玩家可以找到提示和线索，帮助解谜和规避危险。
- **时机把握：**
   - 石像的激光和鬼的探测都有特定的时机，玩家需要巧妙选择合适的时机穿越。
   - 在大厅中找到安全区域的最佳时机，避免被追击的鬼所捕捉。
- **灵活的移动与快速反应：**
   - 玩家可以利用游戏中的灵活移动机制，躲过激光，躲避追逐的鬼。
   - 快速反应和迅速改变行动计划是成功穿越的关键。
- **解谜与故事发现：**
   - 大厅中布满了谜题和隐藏的机关，玩家需要解谜以进一步深入城堡。
   - 解开谜题可能会揭示庄园的过去和关于失踪家族的秘密。
- **安全区域的抵达：**
   - 游戏的终极目标是成功抵达大厅中的安全区域，解除家族的诅咒。
   - 到达安全区域后，玩家完成关卡，可能解锁下一个章节，同时可能获得游戏内奖励。

#### 4. **基本组件**

| MyAnimator | 动画控制器   |
| ---------- | ------------ |
| Prefabs    | 预制体       |
| Scenes     | 场景         |
| Scripts    | C#脚本       |
| Sprite     | 精灵（图片） |
| Art        | 场景美术     |
| Shaders    | 渲染器       |

#### 5.设计和编码

PlayerMovement的Coding，其中包括角色移动，旋转过渡，动画挂载，刚体旋转等内容。

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerMovement : MonoBehaviour
{
    //定义一个旋转的速度
    public float turnspeed = 20f;

    //定义游戏人物上的组件
    Animator m_Animator;
    Rigidbody m_Rigidbody;

    //游戏人物移动的矢量
    Vector3 m_Movement;
    //游戏人物旋转的角度
    Quaternion m_Quaternion = Quaternion.identity;

    void Start()
    {
        //获取游戏人物上的刚体，动画控制机组件
        m_Animator = GetComponent<Animator>();
        m_Rigidbody = GetComponent<Rigidbody>();
    }

    // Update is called once per frame
    void Update()
    {
        
    }
    //固定帧长度，与机器性能就无关
    private void FixedUpdate()
    {
        //获取水平、竖直（z轴方向）是否有键值输入
        float horizontal = Input.GetAxis("Horizontal");
        float vertical = Input.GetAxis("Vertical");

        //设置游戏人物移动的方向
        m_Movement.Set(horizontal,0f, vertical);
        m_Movement.Normalize();

        //定义游戏人物是否移动的bool值
        bool hasHorizontalInput = !Mathf.Approximately(horizontal, 0f);
        bool hasVerticalInput = !Mathf.Approximately(vertical, 0f);
        //控制人物动画的转换
        bool isWalking = hasHorizontalInput || hasVerticalInput;
        m_Animator.SetBool("IsWalking", isWalking);

        //旋转的过渡
        Vector3 desirForward = Vector3.RotateTowards(transform.forward,m_Movement,turnspeed*Time.deltaTime,0f);
        m_Quaternion = Quaternion.LookRotation(desirForward);
    }

    //游戏人物的旋转和移动
    private void OnAnimatorMove()
    {
        m_Rigidbody.MovePosition(m_Rigidbody.position+m_Movement*m_Animator.deltaPosition.magnitude);
        m_Rigidbody.MoveRotation(m_Quaternion);
    }



}

```

