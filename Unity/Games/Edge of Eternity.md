---
title: Edge of Eternity
date: 2024-07-17 9:44:34
sticky: 13
pic: /images/thumb/Edge of Eternity.png
comments: true
toc: true
categories:
- Games
tags:
- Unity
---

# Edge of Eternity

#### 1. 游戏背景

​	《Edge of Eternity》是一款半即时回合制奇幻风格雨世界剧情向游戏。本作品为CIGA48小时极限开发作品的后续开发版本。故事设定在人类灭亡四万年之后，核废水污染了地球的每一寸，由于上层海域污染越来越严重，部分生物发生变异，深海中出现了新的生态文明。主角是诞生在浅海但生活在深海中的生物，由于外观与深海中其他生物不同(发光且可爱)，被其他生物排挤，它的心中一直埋藏着一个寻找自己同类的愿望，偶然间听闻自己的身世，它决定离开深海，向上寻找归属，(战斗开始)在经过重重难后，它发现浅海区的同类们其实因为核污染早已变异(boss战)，存活下来的仅有它自己。

#### 2. 团队成员

程序：u9king,Lam1ngton,韦德西斯widsith

策划：小咖

美术：Celice泥,Toko,小魚

永恒工作室

<img src="/images/Unity/Games/Edge of Eternity/L1.jpg" style="width: 50px; height: auto;display: inline-block;margin-left: 20px;">

#### 3. 游戏内容

<img src="/images/Unity/Games/Edge of Eternity/1.jpg">

<img src="/images/Unity/Games/Edge of Eternity/2.jpg">

<img src="/images/Unity/Games/Edge of Eternity/3.jpg">

#### 4. 核心机制
<img src="/images/Unity/Games/Edge of Eternity/4.png">

​	游戏采用了半即时回合制原理，在充能条达到满的时候，玩家通过按下攻击键，防御键或者背包键根据对手的选择来做出自己的策略。

- **演示视频**

<iframe src="//player.bilibili.com/player.html?bvid=BV1o4beeWENw" width="580px" height="320px"></iframe>

#### 5.核心代码

采用有限状态机SimpleFSM的方式编写,参考文章:https://developer.unity.cn/projects/661e1cc2edbc2a001e0b9e31

```C#
using System.Collections;
using System.Net.NetworkInformation;
using UnityEngine;

public class EnemySimpleFSM : MonoBehaviour
{
    //规范四种状态，闲置，巡逻，追击，死亡
    public enum State { Idle, Patrol, Chase, Dead }

    [Header("参数")]
    public int health;  //仅供参考，非必要
    public State currentState;  //当前执行状态

    private void Awake()
    {
        //仅供参考，唤醒实例化变量
    }

    private void Start()
    {
        // 初始化状态，此处为初始化闲置状态
        currentState = State.Idle;
    }

    private void Update()
    {
        //检测敌人状态变化，非必要
        //CheckState();
        //注册并轮询状态，初始为闲置Idle状态
        currentState = currentState switch
        {
            State.Idle => UpdateIdleState(),
            State.Patrol => UpdatePatrolState(),
            State.Chase => UpdateChaseState(),
            State.Dead => UpdateDeadState(),
            _ => currentState
        };
    }

    private void ChangeState(State nextState)  //切换状态函数
    {
        currentState = nextState;
    }

    private void CheckState()
    {
        //检查状态变化，例如接触地面等发生变化时调用，非必要
        Debug.Log("Check State");
    }

    private State UpdateIdleState()  //Idle状态
    {
        Debug.Log("Idle State");
        //仅供参考,从闲置状态转入巡逻状态的代码逻辑
        //ChangeState(State.Patrol);
        return State.Idle;  //无变化，则维持闲置状态
    }

    private State UpdatePatrolState()  //巡逻状态
    {
        Debug.Log("Patrol State");
        return State.Patrol; //维持巡逻状态
    }

    private State UpdateChaseState()  //追击状态
    {
        Debug.Log("Chase State");
        return State.Chase;
    }

    private State UpdateDeadState()  //死亡状态
    {
        Debug.Log("Dead State");
        return State.Dead;
    }
}
```

