---
title: Brick Breaker
date: 2024-01-10 19:40:56
sticky: 10
pic: /images/thumb/Brick Breaker.png
comments: true
toc: true
categories:
- Games
tags:
- Unity
---

# Brick Breaker

#### 1. 游戏背景

​	打砖块" Brick Breaker" 游戏诞生于1976年，初版 "Breakout"由Atari公司制作，是一款经典的电子游戏，由于其玩法简单：这款游戏因其简单直观的玩法和持久的挑战性而备受欢迎。

#### 2. 游戏内容

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Games/Brick%20Breaker/1.png">

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Games/Brick%20Breaker/2.png">

#### 3. 核心机制

​	玩家操控一个滑板接住反弹的球，而屏幕上方有一堵由砖块组成的墙。玩家的目标是用球击中墙上的砖块，每次击中一个砖块，就会让它消失，直到整堵墙都被击破。游戏的难度会随着关卡的进行而增加，砖块的排列和形状也会变得更加复杂。

​	游戏中会出现一些特殊砖块，给予玩家额外的分数、增加滑板的宽度、减缓球的速度等。

- **演示视频**

<iframe src="//player.bilibili.com/player.html?bvid=BV19M411v7Ni" width="580px" height="320px"></iframe>

#### 4. **基本组件**

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Games/Brick%20Breaker/Basic%20Components.png" style="zoom:40%">

| Physics | 物理材质     |
| ------- | ------------ |
| Prefabs | 预制体       |
| Scenes  | 场景         |
| Scripts | C#脚本       |
| Sprite  | 精灵（图片） |

#### 5.设计和编码

GameManager的Coding，其中包括游戏分数，存活次数，游戏胜负的判定，还包括加载场景，重置物体等代码内容。

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using UnityEngine.UI;

public class GameManager : MonoBehaviour
{
    public Ball ball { get; private set; }
    public Paddle paddle { get; private set; }
    public Brick[] bricks { get; private set; }

    public int level = 1;
    public int score = 0;    //分数
    public int lives = 3;  //复活次数

    //右下角显示文字
    public Text scoreNum;
    public Text liveNum;
    private void Awake()
    {
        DontDestroyOnLoad(this.gameObject);

        SceneManager.sceneLoaded += OnLevelLoaded;
    }
    private void Start()
    {
        NewGame();
    }

    private void NewGame()
    {
        this.score = 0;
        this.lives = 3;
        LoadLevel(1);       //todo多余否
    }

    private void LoadLevel(int level)
    {
        this.level = level;
        if (level > 1)
        {
            SceneManager.LoadScene("WinScreen");
        }
        else
        {
            SceneManager.LoadScene("Level" + level);
        }
    }

    private void OnLevelLoaded(Scene scene, LoadSceneMode mode)
    {
        this.ball = FindObjectOfType<Ball>();
        this.paddle = FindObjectOfType<Paddle>();
        this.bricks = FindObjectsOfType<Brick>();
        this.scoreNum = GameObject.Find("scoreNum").GetComponent<Text>();
        this.liveNum = GameObject.Find("liveNum").GetComponent<Text>();
    }

    private void Update()
    {
        scoreNum.text = score.ToString();
        liveNum.text = lives.ToString();
    }

    private void ResetLevel()
    {
        //用掉一条命重置内容
        this.ball.ResetBall();
        this.paddle.ResetPaddle();
        //for (int i = 0; i < this.bricks.Length; i++)
        //{
        //    this.bricks[i].ResetBrick();
        //}

    }

    private void GameOver()
    {
        //游戏结束  待完成
        //SceneManager.LoadScene("GameOver");

        NewGame();
    }


    public void Miss()
    {
        this.lives--;
        if (this.lives > 0)
        {
            ResetLevel();
        }
        else
        {
            GameOver();
        }
    }

    public void Hit(Brick brick)
    {
        this.score += brick.points;

        if (Cleared())
        {
            LoadLevel(this.level + 1);
        }
    }

    private bool Cleared()
    {
        for (int i = 0; i < this.bricks.Length; i++)
        {
            if (this.bricks[i].gameObject.activeInHierarchy && !this.bricks[i].unbreakable)
            {
                return false;
            }
        }
        return true;
    }
}

```

