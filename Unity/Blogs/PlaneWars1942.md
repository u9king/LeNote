---
title: PlaneWars1942
date: 2024-01-10 19:40:56
pic: /images/thumb/PlaneWars1942.png
comments: true
toc: true
categories:
- Games
tags:
- Unity
---

# PlaneWars1942

#### 1. 游戏背景

​	《Plane Wars 1942》是一款经典的飞行射击游戏，最初于1984年由日本的卡普空（Capcom）公司推出。这款游戏的最初版本名为《1942》街机游戏，成为了飞行射击游戏领域的里程碑。

​	故事背景设定在第二次世界大战时期，玩家将扮演一名精英飞行员，操控各种经典的战斗机，投身于激烈的空战中。敌方军队在太平洋战场展开侵略，而玩家的任务是独自对抗庞大的敌军，争取取得制空权，保卫自己的领土。该游戏在北美、亚洲和欧洲等多个地区都取得了火爆的销售成绩。

#### 2. 游戏内容

<img src="/images/Unity/Games/PlaneWars1942/1.png">

#### 3. 核心机制

- **演示视频**

<iframe src="//player.bilibili.com/player.html?bvid=BV1xj411g76B" width="580px" height="320px"></iframe>

-. **飞机控制：** 玩家操控一架战斗机，使用操纵杆或方向键进行上下左右的移动，以及进行飞行机动。
-. **射击系统：** 游戏以垂直卷轴的形式进行，敌机和敌方设施从上方滚动而下。玩家可以使用按钮进行射击，向前发射子弹以摧毁敌机、敌方设施和其他敌人。
-. **敌人和关卡设计：** 游戏中包含多样化的敌人，包括敌机、轰炸机和敌方设施。每个关卡都有独特的地形和敌军编队，玩家需要适应不同的战术来取得胜利。
-. **能量和武器升级：** 在游戏中，玩家可以获得能量和武器升级，增强他们的飞机性能和火力。这可能包括改进子弹、增加射程、提高飞机速度等。
-. **Boss战：** 游戏的某些关卡会有强大的敌军首领（Boss），玩家需要利用游戏中获得的技能和升级，战胜这些庞大的敌人。
-. **生命和续命系统：** 玩家通常有有限的生命数，当生命耗尽时，玩家可以使用游戏内的续命机制（如果有的话），继续挑战。


#### 4. **基本组件**

| Resource | 包含音乐、预制体 |
| -------- | ---------------- |
| Scripts  | C#脚本           |
| Sprites  | 精灵（图片）     |

#### 5.设计和编码

PlayerControl的Coding，其中包括玩家血量，子弹数量，散射角度，死亡判定等内容。

```C#
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
[RequireComponent(typeof(BoxCollider2D))]  //请求添加2D的盒型碰撞体

public class PlayerControl : MonoBehaviour
{
    private GameObject bulletPrefab;
    private Transform firePos;
    public AudioClip clips;
    public int hp;  //玩家血量
    public int bulletNumber;  //子弹数量
    public float angle;  //角度

    void Start()
    {
        bulletPrefab = Resources.Load<GameObject>("Bullet");
        firePos = transform.GetChild(0);
        GetComponent<BoxCollider2D>().size = new Vector2(1.0f, 1.0f);
        InvokeRepeating("Attack", 0, 0.5f);

    }

    void Update()
    {
        //键盘移动
        OnKeyMove();

        //特殊攻击使用空格键，todo后续修改为吃道具修改为时效性
        if (Input.GetKeyDown(KeyCode.Space))
        {
            RangeAttack();
        };
    }


    void OnMouseDrag()  //鼠标按下即为移动，内置函数自动调用
    {
        Vector3 targetPos = Camera.main.ScreenToWorldPoint(Input.mousePosition);
        targetPos.z = 1;
        targetPos.x = Mathf.Clamp(targetPos.x, -2.3f, 2.3f);
        targetPos.y = Mathf.Clamp(targetPos.y, -4.5f, 4.5f);
        transform.position = targetPos;
    }

    void OnKeyMove() //键盘移动
    {
        // 设置移动速度
        float moveSpeed = 10f; 

        // 获取水平和垂直输入（A和D键控制水平，W和S键控制垂直）
        float horizontalInput = Input.GetAxis("Horizontal");
        float verticalInput = Input.GetAxis("Vertical");

        // 计算移动方向
        Vector3 movement = new Vector3(horizontalInput, verticalInput, 0f);

        // 根据输入调整位置
        transform.Translate(movement * moveSpeed * Time.deltaTime);

        // 限制玩家的移动范围
        Vector3 clampedPosition = transform.position;
        clampedPosition.x = Mathf.Clamp(clampedPosition.x, -2.3f, 2.3f);
        clampedPosition.y = Mathf.Clamp(clampedPosition.y, -4.5f, 4.5f);
        transform.position = clampedPosition;
    }

    private void Attack()
    {
        GameObject tempBullet = Instantiate(bulletPrefab, firePos.position, firePos.rotation);
        tempBullet.AddComponent<BulletControl>();
        tempBullet.name = "PlayerBullet";
         AudioSource.PlayClipAtPoint(clips, firePos.position, 0.5f);
    }

    private void RangeAttack() //特殊攻击
    {
        for (int i = -bulletNumber / 2; i < bulletNumber / 2 + 1; i++)
        {
            GameObject tempBullet = Instantiate(bulletPrefab, firePos.position, firePos.rotation);
            tempBullet.transform.Rotate(0, 0, angle * i);
            tempBullet.AddComponent<BulletControl>();
            tempBullet.name = "PlayerBullet";
             AudioSource.PlayClipAtPoint(clips, firePos.position, 0.5f);
        }
    }

    private void Damage(int damage)
    {
        if (hp > 0)
        {
            hp -= damage;
            if (hp <= 0)
            {
                hp = 0;
                //死亡
                Debug.Log("死亡");
                //UnityEditor.EditorApplication.isPlaying = false;
            }
        }
        else { }
    }

}

```

