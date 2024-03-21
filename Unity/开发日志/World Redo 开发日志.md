# World Redo 开发日志

#### 1.游戏类型

类地下城，像素风，战斗剧情向游戏

#### 2.技术内容

背包系统

战斗系统

购买系统

角色升级系统

武器系统

日夜切换系统

宠物系统

剧情内容系统

特效shader系统

#### 3.关键技术：

1.有限状态机

2.行为树



#### 4.阶段划分

1阶段 5月份前

搭建5关左右的demo实现除宠物、剧情、特效外的所有内容



2阶段  6月份后 至 8月份

实现宠物、剧情、特效、优化美术工作

可以尝试搭建到30关内容

拍摄部分开发阶段内容



3阶段

实现打包，平台测试、商讨发布工作，投稿平台

#### 4.时间表

| 日期 | 工作 |                             备注                             |
| :--: | :--: | :----------------------------------------------------------: |
|      |      |                                                              |
|      |      |                                                              |
|      |      |                                                              |
|      |      |                                                              |
| 3/20 |      |                         switch语法糖                         |
| 3/19 |      | 三段攻击的攻击间隔存在问题<br />动画器存在问题<br />是否需要将动画逻辑剥离<br />写在动画上的逻辑在代码中无法查看，应该怎么办<br />新的地面存在卡墙的情况 |
| 3/18 |      |                优化跳跃手感<br />替换地形内容                |
| 3/15 |      |                      学习勇士传说<br />                      |
| 3/14 |      | 动画从Attack到Jump存在问题，在Fall的情况下不能变回正常的Idle情况<br />角色攻击判定被锁死所有物体都需要使用Polygon的问题<br />字体大小会随窗口大小变化而移动<br />蝙蝠的流血效果不够平滑，没有跟着蝙蝠移动，而是停留在原地流血<br />待优化：敌人死亡时，巡逻点和目标点也应该被删除掉，本质上需要解决leftPos挂在身上随本体变化的问题<br />动画旋转如何不使用差值，我需要将角色旋转到合适位置 |
| 3/13 |      | 单向平台边沿存在卡角色物体的情况<br />向上爬行存在自动的情况<br /> |
| 3/12 |      | 动画系统存在问题<br />瓦片网格大小不一致怎么办<br /><br />少屋顶素材<br />~~写文字一定要用Canvas吗，可以用TextRender~~ |
| 3/11 |      |           天空盒使用遇见问题<br />混合树动画有问题           |
|      |      |                                                              |



每关大概3-5个怪

每关2-3个机关

每章5+N个场景



### 1.核心代码总结

人物移动

```
public void Run()   //跑步
{
    float moveX = Input.GetAxisRaw("Horizontal");
    myRigidBody.velocity = new Vector2(moveX * runSpeed, myRigidBody.velocity.y);
    bool xAxisHasSpeed = Mathf.Abs(myRigidBody.velocity.x) > Mathf.Epsilon;    //检测速度
    if (xAxisHasSpeed)  //翻转
    {
        if (myRigidBody.velocity.x > 0.1f)
        {
            transform.localRotation = Quaternion.Euler(0, 0, 0);
        }
        if (myRigidBody.velocity.x < -0.1f)  //向左移动
        {
            transform.localRotation = Quaternion.Euler(0, 180, 0);  //图像翻转
        }
    }
    myAnim.SetBool("isRun", xAxisHasSpeed);    //动画设置布尔值
}
```

单向平台

```
void PlatformCheck()
{
    isPlatform = myFeet.IsTouchingLayers(LayerMask.GetMask("Platform"));
    if (isPlatform && Input.GetKeyDown("s"))
    {
        gameObject.layer = LayerMask.NameToLayer("Platform");
        Invoke("RestorePlayerLayer", 0.5f);
    }
}
void RestorePlayerLayer()
{
    gameObject.layer = LayerMask.NameToLayer("Player");
}
```

