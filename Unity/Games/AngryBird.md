---
title: AngryBird
date: 2024-01-10 19:40:56
sticky: 10
pic: /images/thumb/AngryBird.png
comments: true
toc: true
categories:
- Games
tags:
- Unity
---

# AngryBird

#### 1. 游戏背景

​	《AngryBird》是一款由芬兰游戏开发公司Rovio Entertainment开发的独特而火爆的物理学游戏。该游戏于2009年12月首次推出，迅速在全球范围内引起轰动，成为一种文化现象。背景故事发生在一个美丽的岛屿群中，这里的宁静被一群狡猾的绿色猪所打破。这些猪偷走了愤怒的小鸟们的蛋，激起了小鸟们愤怒的怒火。为了夺回蛋并向猪们复仇，小鸟们决定以自己独特的方式对付猪群。

​		这款游戏在全球范围内都取得了极大的成功，无论是在亚洲、欧洲、北美还是南美，都有大量的玩家。由于其简单而有趣的玩法，它成功地突破了语言和文化的障碍，成为一种国际化的娱乐现象。

#### 2. 游戏内容

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Games/AngryBird/1.png">

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Games/AngryBird/2.png">

#### 3. 核心机制

- **演示视频**

<iframe src="//player.bilibili.com/player.html?bvid=BV1qb411c76x" width="580px" height="320px"></iframe>

​	游戏采用了简单直观的物理学原理，让玩家通过拉动弹弓并释放小鸟，以精准的角度和力度来摧毁猪群所在的各种结构，从而解救被偷走的蛋。
#### 4. **基本组件**

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Games/AngryBird/Basic%20Components.png" style="zoom:60%">



| Animations | 动画控制器              |
| ---------- | ----------------------- |
| Images     | 精灵（图片）            |
| Materials  | 材料（胜利粒子效果）    |
| Music      | 音乐                    |
| Prefabs    | 预制体                  |
| Resources  | 包含所有关卡(stage)内容 |
| Scenes     | 场景                    |
| Scripts    | C#脚本                  |
| TileMap    | 瓦片地图                |

#### 5.设计和编码

RedBirdControl的Coding，其中包括拉弹弓点击事件，相机跟随，小鸟销毁，使用技能，受伤切换等内容。

```C#
using System.Collections;
using System.Collections.Generic;
//using TreeEditor;
using UnityEngine;

public class RedBirdControl : MonoBehaviour
{
    //获取状态
    //获取点击情况
    private bool isClick;
    //获取飞行状态
    private bool isFlying;
    //限制发射后不能进行选择
    private bool canSelected = true;

    [HideInInspector]
    public SpringJoint2D sp;
    protected Rigidbody2D rg;

    //获取爆炸特效
    public GameObject boomEffect;

    //调用弹弓控制
    private SlingshotControl slingshot;

    //平滑相机移动
    private float smooth = 3;

    //小鸟音效
    public AudioClip selectSound;
    public AudioClip flySound;

    //技能使用判断
    private bool canUseSkill;

    //受伤图片
    public Sprite hurtIMG;
    protected SpriteRenderer render;

    //小鸟初始位置获取
    private Vector3 startPosition;

    //Awake和Start的区别
    private void Awake()
    {
        //默认初始赋值区域
        sp = GetComponent<SpringJoint2D>();
        rg = GetComponent<Rigidbody2D>();
        render = GetComponent<SpriteRenderer>();
    }

    private void Start()
    {
        //默认初始赋值区域
        slingshot = GameObject.Find("Slingshot").GetComponent<SlingshotControl>();
    }

    //鼠标按下
    private void OnMouseDown()
    {
        if (canSelected)
        {
            //播放选中音效
            AudioSource.PlayClipAtPoint(selectSound, transform.position);
            isClick = true;
            //todo
            //动力学开启关闭的原因？
            rg.isKinematic = true;
            //记录初始小鸟位置
            startPosition = transform.position;
        }
    }

    //鼠标抬起
    private void OnMouseUp()
    {
        if (canSelected)
        {
            //todo开启重力//即可解决++1
            isClick = false;
            rg.isKinematic = false;
            //过0.1秒后飞出？0.1秒吗
            Invoke("FlyOut", 0.1f);
            //禁用划线组件
            slingshot.rightRender.enabled = false;
            slingshot.leftRender.enabled = false;
            canSelected = false;
            //开启小鸟旋转
            rg.freezeRotation = false;
        }

    }

    private void Update()
    {
        //小鸟跟随鼠标按下位置移动
        if (isClick)
        {
            transform.position = Camera.main.ScreenToWorldPoint(Input.mousePosition);
            transform.position -= new Vector3(0, 0, Camera.main.transform.position.z);
            if (Vector3.Distance(transform.position, slingshot.rightPos.position) > slingshot.maxBandDis)
            {
                //限制最大移动向量
                Vector3 directionVector = (transform.position - slingshot.rightPos.position).normalized * slingshot.maxBandDis;
                //重新赋值最大位置
                transform.position = directionVector + slingshot.rightPos.position;
            }
            if (slingshot != null)
            {
                slingshot.DrawBands(transform);
            }
        }

        //相机追随,数字大于10就会出现瓦片重载的问题
        float posX = transform.position.x;
        //相机跟随的最大距离可以通过调整Mathf.Clamp(posX, 0, 5)中的5来实现
        Camera.main.transform.position = Vector3.Lerp(Camera.main.transform.position, new Vector3(Mathf.Clamp(posX, 0, 6), Camera.main.transform.position.y,
            Camera.main.transform.position.z), smooth * Time.deltaTime);

        //飞行状态判定
        if (isFlying)
        {
            // 检查小鸟是否静止落地或者超出最大距离后，1秒开始销毁
            if (rg.velocity.magnitude < 0.2f || Mathf.Abs(transform.position.x - startPosition.x) > 6)
            {
                isFlying = false;
                Invoke("Ruin", 2);
            }
        }


        //技能使用判断
        if (canUseSkill)
        {
            //如果点击鼠标,就可以使用技能
            if(Input.GetMouseButtonDown(0))
            {
                UseSkill();
            }

        }

    }

    //小鸟飞出动作
    void FlyOut()
    {
        //小鸟技能使用开启
        canUseSkill = true;
        //飞出时关闭弹簧
        sp.enabled = false;
        //播放飞行音效
        AudioSource.PlayClipAtPoint(flySound, transform.position);
        //飞出3秒后执行销毁
        //todo应该更改为落地/静止后3秒销毁更加合理
        //Invoke("Ruin", 3);
        // 标记小鸟正在飞行
        isFlying = true;

    }

    //使用技能，供各类小鸟继承
    public virtual void UseSkill()
    {
        //小鸟技能使用后关闭
        canUseSkill = false;
    }

    //落地销毁小鸟
    protected virtual void Ruin()
    {
        //在游戏管理器中删除自身
        GameManagerControl._instance.birds.Remove(this);
        //销毁自身
        Destroy(gameObject);
        //爆炸特效
        Instantiate(boomEffect, transform.position, Quaternion.identity);
        //每只鸟消失时判断一下当前胜负环境
        GameManagerControl._instance.JudgeGameResult();
    }

    //触碰物体后动作
    private void OnCollisionEnter2D(Collision2D collision)
    {
        //todo？开始就落地就受伤变化了
        //解决方案，点击释放后再开启碰撞（重力）
        //碰撞到猪或者木头发生受伤状态
        if (collision.gameObject.tag == "Enemys"  || collision.gameObject.tag == "Blocks")
        {
            //切换受伤精灵
            render.sprite = hurtIMG;
        }
        //碰撞到物体后关闭技能使用
        canUseSkill = false;
    }
}

```

