---
title: Unity有限状态机，SimpleFSM
date: 2024-08-30 20:07:35
pic: /images/thumb/SimpleFSM.png
comments: true
toc: true
categories:
- Engine
tags:
- Unity
---

# Unity有限状态机，SimpleFSM

​	SimpleFSM是一个简单的敌人有限状态机（FSM）实现，用于管理敌人的行为。

​	SimpleFSM适用于开发2D或3D游戏中的敌人AI系统，特别是那些需要基本状态管理的情况。它可以作为一个简单的起点，用于创建更复杂的敌人行为模式。 这里定义了四种状态：Idle（闲置）、Patrol（巡逻）、Chase（追击）、Dead（死亡），并根据当前状态执行相应的行为。

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

​	进一步优化代码：针对状态中所需要的单次执行事件，考虑采用状态锁的方式实现引入变量stateLock,下面将以Idle为例仅展示修改部分的代码

```C#
using UnityEngine;
public class PlayerAnimFSM : MonoBehaviour
{
    private enum AnimationState {Idle, Attack}
    private AnimationState currentState;
    private Animator animator;
    private float waitAnimTime = 0f;    //动画计时器

    void Start()
    {
        animator = GetComponent<Animator>();
    }

    void Update()
    {
        CheckState();
        currentState = currentState switch
        {
            AnimationState.Idle => IdleState(),
            AnimationState.Attack => AttackState(),
            _ => currentState
        };
    }

    private AnimationState IdleState()
    {
        animator.Play("Idle");
        return AnimationState.Idle;
    }

    private AnimationState AttackState()
    {
        //播放动画
        animator.Play("Attack");
        waitAnimTime += Time.deltaTime;
        if (waitAnimTime >= animator.GetCurrentAnimatorStateInfo(0).length)
        {
            //当动画记录时间大于当前正在播放动画的时间时
            //todo：这里有一个BUG，在animator.Play()的动画需要在下一帧animator.GetCurrentAnimatorStateInfo(0).length才能获取到正确的时间
            //在这里默认动画长度都会大于1帧所以没太大的问题
            //正确的做法是参考方法一种的GetAnimationClipLength来获取动画时间
            waitAnimTime = 0;   //重置动画时间
            Debug.Log("Attack播放完了，可以执行Idle");
            return AnimationState.Idle; //转换Idle状态
        }
        return AnimationState.Attack; //维持攻击状态
    }

    private void CheckState()    //检测状态转换
    {
        if (currentState == AnimationState.Attack)  //维持攻击动画不被打断
            return;
        if (Input.GetKeyDown(KeyCode.Space))
        {
            currentState = AnimationState.Attack;
            return;
        }
        currentState = AnimationState.Idle;
    }
}

```

进一步优化代码：针对状态中所需要的单次执行事件，考虑采用状态锁的方式实现引入变量stateLock,下面将以Idle为例仅展示修改部分的代码

```C#
public class EnemySimpleFSM : MonoBehaviour
{
    private bool stateLock = false;  //引入状态锁

    private void ChangeState(State nextState)  //修改ChangeState函数引入stateLock
    {
        if (currentState != nextState)
        {
            stateLock = true;   //状态发生变化的时候开启状态锁
            currentState = nextState;
        }
        return currentState;
    }

    private State UpdateIdleState()  //Idle状态
    {
        if (stateLock)
        {
            stateLock = false;   //第一次进入状态后关闭锁
            Debug.Log("只执行一次的 Idle State");   //引入状态锁,只在状态变化执行一次
            return State.Idle;  //或者根据情况转出为其他状态
        }
        return State.Idle;
    }
}
```
