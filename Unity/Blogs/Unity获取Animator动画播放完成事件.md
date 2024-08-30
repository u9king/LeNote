---
title: Unity获取Animator动画播放完成事件
date: 2024-08-30 20:13:34
pic: /images/thumb/GetAnimator.png
comments: true
toc: true
categories:
- Programming
tags:
- Unity
---

# Unity获取Animator动画播放完成事件

整理了一些在日常经验中处理动画播放完成事件的方法
方法:
1.Dotween配合异步实现
2.状态机计时方法实现
3.原生动画行为方法实现

#### 方法一：Dotween异步方法

```C#
using UnityEngine;
using System.Threading.Tasks;
using DG.Tweening;

public class PlayerAnimAsync : MonoBehaviour
{
    private Animator animator;
    private bool isAnimPlaying = false;

    void Start()
    {
        animator = GetComponent<Animator>();
    }

    void Update()
    {
        animator = GetComponent<Animator>();

        // 开始动画
        if (!isAnimPlaying)
        {
            StartAttack();
        }
    }

    private async void StartAttack()
    {
        isAnimPlaying = true;
        await AnimationFinish("Attack", 0f);   //等待Attack播放完
        Debug.Log("Attack播放完了，可以执行Idle");
        animator.Play("Idle");
        isAnimPlaying = false;
    }

    //AnimationFinish 异步播放动画
    public async Task AnimationFinish(string animName, float extreTime = 0f)
    {
        await DOTween.Sequence()
            .AppendCallback(() => animator.Play(animName))
            .AppendInterval(GetAnimationClipLength(animName) + extreTime)
            .AsyncWaitForCompletion();
    }

    //GetAnimationClipLength 获取动画片段时长
    private float GetAnimationClipLength(string animName)
    {
        RuntimeAnimatorController ac = animator.runtimeAnimatorController;
        foreach (AnimationClip clip in ac.animationClips)
        {
            if (clip.name == animName)
            {
                return clip.length;
            }
        }
        return 0f;
    }
}
```

#### 方法二：状态机计时方法

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

#### 方法三：原生动画行为方法实现

这里需要用到两个脚本PlayerAnimSM和AttackFinish来实现，此处isAnimPlaying借助原生动画行为来复原

```C#
using UnityEngine;

public class PlayerAnimSM : MonoBehaviour
{
    private Animator animator;
    public bool isAnimPlaying = false;

    void Start()
    {
        animator = GetComponent<Animator>();
    }

    void Update()
    {
        animator = GetComponent<Animator>();

        // 开始动画
        if (!isAnimPlaying)
        {
            StartAttack();
        }
    }

    private void StartAttack()
    {
        isAnimPlaying = true;
        animator.Play("Attack");
    }
}
```

AttackFinish脚本借助界面创建步骤如下：
1.在动画器中点击需要传递动画完成事件的动画，点击右下角的Add Behaviour(添加行为)，可以添加Unity预制的脚本
2.使用这个方法需要有动画过渡的方式(此处为AnyState到Idle)，供后续代码中的OnStateExit使用（举例：如果希望有攻击结束到闲置动画的过渡，就需要从攻击动画连线到闲置动画，重点！！！一定要有退出时间，设置0s也没事，但一定要勾选，这里我就使用AnyState过渡过去了，使用AnyState时也一定要勾选退出时间

<img src="/images/Unity/Programming/GetAnimator/1.png">

<img src="/images/Unity/Programming/GetAnimator/2.png">

<img src="/images/Unity/Programming/GetAnimator/3.png">

创建名为AttackFinish 的脚本(这里Unity叫行为)
双击点开这个脚本

```C#
using UnityEngine;

public class AttackFinish : StateMachineBehaviour
{
    // OnStateEnter is called when a transition starts and the state machine starts to evaluate this state
    //override public void OnStateEnter(Animator animator, AnimatorStateInfo stateInfo, int layerIndex)
    //{
    //}

    // OnStateUpdate is called on each Update frame between OnStateEnter and OnStateExit callbacks
    //override public void OnStateUpdate(Animator animator, AnimatorStateInfo stateInfo, int layerIndex)
    //{

    //}

    //OnStateExit is called when a transition ends and the state machine finishes evaluating this state
    override public void OnStateExit(Animator animator, AnimatorStateInfo stateInfo, int layerIndex)
    {
        //在此处将isAnimPlaying重置回false
        animator.GetComponent<PlayerAnimSM>().isAnimPlaying = false;
        Debug.Log("Attack播放完了");
        //播放动画结束后的默认动画，我这里设置为idle你可以设置为任意动画但是一定要有过渡，从Attack到Idle的过渡
        animator.Play("Idle");
    }

    // OnStateMove is called right after Animator.OnAnimatorMove()
    //override public void OnStateMove(Animator animator, AnimatorStateInfo stateInfo, int layerIndex)
    //{
    //    // Implement code that processes and affects root motion
    //}

    // OnStateIK is called right after Animator.OnAnimatorIK()
    //override public void OnStateIK(Animator animator, AnimatorStateInfo stateInfo, int layerIndex)
    //{
    //    // Implement code that sets up animation IK (inverse kinematics)
    //}
}
```

#### 后谈

​	还有一些诸如动画帧事件的方法没有收录在内。作者总感觉无论哪一个方法都不是特别合适或者顺手，当然无论是用原生还是自己去写，寻找适合自己项目的方法才是最好的。所以作者也会在今后的开发道路上继续学习。