# Unity实现有限状态机

#### 1.实现接口

```C#
public interface IState //状态类型接口
{
    void OnEnter(); //状态进入

    void OnUpdate(); //状态执行

    void OnExit();  //状态退出
}
```

#### 2.实现FSM

```c#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;


public enum StateType   //枚举变量包含所有状态类型
{
    Idle, Patrol, Chase,React, Attack
}

[Serializable]
public class Parameter
{
    public int health;  //怪物血量
    public int moveSpeed;   //怪物移动速度
    public float chaseSpeed;    //怪物追击速度
    public float IdleTime;  //怪物闲置时间
    public Transform[] patrolPoints;    //怪物巡逻范围
    public Transform[] chasePoints; //怪物追击范围
    public Transform target;    //攻击目标
    public LayerMask targetLayer;
    public Transform attackPoint;   //攻击圆心
    public float attackArea;    //攻击圆半径
    public Animator animator;   //动画控制器
}

public class FSM : MonoBehaviour
{
    public Parameter parameter;
    private IState currentState;    //当前状态
    private Dictionary<StateType, IState> states = new Dictionary<StateType, IState>();   //创建状态
    // Start is called before the first frame update
    void Start()
    {
        //向字典中注册状态
        states.Add(StateType.Idle, new IdleState(this));
        states.Add(StateType.Patrol, new PatrolState(this));
        states.Add(StateType.Chase, new ChaseState(this));
        states.Add(StateType.React, new ReactState(this));
        states.Add(StateType.Attack, new AttackState(this));

        parameter.animator = GetComponent<Animator>();

        TransitionState(StateType.Idle);
    }

    // Update is called once per frame
    void Update()
    {
        currentState.OnUpdate();
    }

    public void TransitionState(StateType type)
    {
        if (currentState! != null)
            currentState.OnExit();
        currentState = states[type];
        currentState.OnEnter();

    }

    public void FlipTo(Transform target)
    {
        if (target != null)
        {
            if ( transform.position.x > target.position.x)
            {
                transform.localScale = new Vector3(-1,1,1);
            } else if (transform.position.x < target.position.x) 
            {
                transform.localScale = new Vector3(1,1,1);
            }
        }
    }

    private void OnTriggerEnter2D(Collider2D collision)
    {
        if (collision.CompareTag("Player") && collision.GetType().ToString() == "UnityEngine.CapsuleCollider2D")
        {
            parameter.target = collision.transform;
        }
    }

    private void OnTriggerExit2D(Collider2D collision)
    {
        if (collision.CompareTag("Player") && collision.GetType().ToString() == "UnityEngine.CapsuleCollider2D")
        {
            parameter.target = null;
        }
    }

    private void OnDrawGizmos()  //绘制攻击范围
    {
        Gizmos.DrawWireSphere(parameter.attackPoint.position, parameter.attackArea);
    }
}
```

#### 3.实现每个状态的功能（Idle）

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class IdleState : IState
{
    private FSM manager;
    private Parameter parameter;
    private float timer;

    public IdleState(FSM manager)
    {
        this.manager = manager;
        this.parameter = manager.parameter;
    }

    public void OnEnter()  //状态进入
    {
        parameter.animator.Play("SnakeIdle");
    }
    public void OnUpdate()  //状态执行
    {
        timer += Time.deltaTime;
        if (parameter.target != null &&
          manager.transform.position.x >= parameter.chasePoints[0].position.x &&
          manager.transform.position.x <= parameter.chasePoints[1].position.x)
        {
            manager.TransitionState(StateType.Attack);
        }
        if (timer >= parameter.IdleTime)
        {
            manager.TransitionState(StateType.Patrol);
        }
    }
    public void OnExit() //状态退出
    {
        timer = 0;
    }
}
```

#### 4.实现每个状态的功能（Patrol）

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PatrolState : IState
{
    private FSM manager;
    private Parameter parameter;
    private int patrolPosition;     //巡逻点序号
    public PatrolState(FSM manager)
    {
        this.manager = manager;
        this.parameter = manager.parameter;
    }

    public void OnEnter()  //状态进入
    {
        parameter.animator.Play("Walk");
    }
    public void OnUpdate()  //状态执行
    {
        manager.FlipTo(parameter.patrolPoints[patrolPosition]);

        manager.transform.position = Vector2.MoveTowards(manager.transform.position, parameter.patrolPoints[patrolPosition].position, parameter.moveSpeed * Time.deltaTime);

        if (parameter.target != null &&
            manager.transform.position.x >= parameter.chasePoints[0].position.x &&
            manager.transform.position.x <= parameter.chasePoints[1].position.x)
        {
            manager.TransitionState(StateType.React);
        }

        if (Vector2.Distance(manager.transform.position, parameter.patrolPoints[patrolPosition].position) < .1f)
        {
            manager.TransitionState(StateType.Idle);
        }
    }
    public void OnExit() //状态退出
    {
        patrolPosition++;
        if (patrolPosition >= parameter.patrolPoints.Length)
        {
            patrolPosition = 0;
        }
    }
}
```

