# World Redo 开发日志

#### 1.游戏类型

银河恶魔城系列（类地下城），（可选像素风），boss战+地下城探索+剧情向(完成遗愿向)游戏，

欧洲女巫为主题，成长经历为主线，核心玩法为魔法，能力收集和剧情探索。

#### 1.1 核心玩法

地图关卡设计，boss战斗，剧情展示，人物能力

#### 2.技术内容

战斗系统 √

存储系统 √

背包系统 √

购买系统

角色升级系统

武器系统

日夜切换系统

宠物系统

剧情内容系统

特效shader系统

#### 2.1 人物能力

一级：走，跑，跳，攻击

二级：爬墙（1），滑铲（2），冲刺（4），射箭（6），滑翔（8），游泳（10）

三级：宠物（3）（仅限捡拾道具），法术攻击（12），替身攻击（14）

#### 2.2 场地能力

机关：风扇（1），弹簧（1），齿轮，，单向平台，移动平台，地刺，激光，念咒语，传送门，密室，隐藏通道，爬梯子，掉落重物，链球，移动刺头



#### 2.3 场地种类

森林1，沙漠2，雪地3，沼泽4，烈日5，海边6，地下7，石头8，木质建筑9，古城10，雷雨11，悬崖12，天空13，水底14，学院15，寺庙16，遗迹17，熔岩18，丧尸19，星海20，战争21，地下城22，奇幻23，魔法24，科技25，赛博26，黑白27，萨满28，傣族29，酒会30，原野31，黑帮32，君主33，法老34，三星堆35，血月36，爵士37



#### 2.4 boss

boss战斗大致分为20-30个左右

boss战更加侧重艺术表达，画面一定要华丽



每个boss房间前方需要有10-15分钟的地图探索

期间每个地图需要有2-3种不同的小怪，侧重点是小怪的攻击行为，出现地点，出现个数略有不同



20*2.5 = 50种小怪

20个boss

20张地图

#### 2.5 特色风格

商店小推车运送前线效果（随时购买）

跺脚从树上掉下来东西效果

以水车为目标设计一个关卡

（较多的魔法效果）

场景传送（打响指召唤）

#### 2.6 人物线

魔法导师：伟大，疯狂，提前为主角铺好道路

挚友：结局1

情人1：若即若离

情人2：求而不得

追随者1：心生仰慕，弱小无助，后续立碑纪念，凡人

死敌：对抗，结局2

邪恶组织：

低智：心善，率真，无心机

宠物：协助



#### 3.关键技术：

1.有限状态机√

2.行为树√

3.实现多平台，PC和手柄要都能支持

#### 3.1 对于场景比例的研究

人物建议采用：64*64的 按照每单位像素16进行放入

建筑瓦片建议采用：32*32的 按照每单位像素32进行放入

人物比地图大致为：8.2(8~9)    空洞：8.88   莫莫多拉：9.4  勇敢的哈克：7.5



#### 4.阶段划分

1阶段 5月份前

搭建5关左右的demo实现除宠物、剧情、特效外的所有内容



2阶段  5月份后 至 8月份实现第一关和第二关的demo工作



3阶段  寻找画师和制作过场动画（12月份之前）

实现宠物、剧情、特效、优化美术工作

可以尝试搭建到20关内容



4阶段

实现打包（PC端和Switch(尝试)），平台测试、宣发工作（制作宣传片和教程）、商讨发布工作，投稿平台（Steam、Taptap等）



#### 4.时间表

| 日期 | 工作 |                             备注                             |
| :--: | :--: | :----------------------------------------------------------: |
| 6/13 |      |                                                              |
| 6/12 |      |                编写游戏策划案搭建地图关卡demo                |
| 6/11 |      |                    学习AIGC绘画任务素材图                    |
| 6/10 |      |                      制作MoonGraveyard1                      |
| 6/7  |      | 制作LavaVolcano1和PastelWorld1<br />建议后续不要采用垂直地图的创作思路，部分背景会存在一定的问题<br />学习部分2D水系统的内容 |
| 6/6  |      | 制作GrimRock1和GothicCemetery1<br />其中GrimRock采用大场景制作，细节内容丰富<br />后续设计也需要考虑到环境因素 |
| 6/5  |      | 制作Alien1地图WinterForest1<br />对于WinterForest1的素材处理存在较大问题 |
| 6/4  |      | 制作GothicCemetery1和WarpCave1<br />使用Grid设置0.99的方式解决WarpCave图片中瓦片存在缝隙漏出紫色光的问题 |
| 6/3  |      |   制作两张地图Mondstadt1和WarpCave1<br />看Tilemap相关视频   |
| 6/2  |      |                      学习Tiled软件使用                       |
| 5/30 |      |           制作两张地图Cyberpunk1和MedievalCastle1            |
| 5/29 |      |              制作两张地图PrisonCity1和DarkCave1              |
| 5/28 |      |         制作两段动画BeginTimeline和小骷髅穿越传送门          |
| 5/27 |      | 制作两张地图OakWoods1和Mossy cavern1<br />如何解决Sprite需要模糊的问题？<br />看书<br />编写一定程度的人物关系 |
| 5/24 |      | 制作两张地图DarkForest1和StarForest1<br />如何解决素材切割完过小的问题？<br />如何能够让不同地图的尺寸保持一致？ |
| 5/23 |      |         学习空洞骑士关卡设计，蔚蓝关卡设计<br />看书         |
| 5/22 |      |              研究Timeline对Dialog系统的使用情况              |
| 5/21 |      | 制作两张地图Moonlight1和SwampForest1<br />如何用Rule Tile实现有限定的随机，通常三块需要连在一起而不是全随机？ |
| 5/20 |      |            以女巫为题材设计地图和Timeline动画内容            |
| 4/29 |      |                         骨骼动画学习                         |
| 4/28 |      |      替换人物为Knight插图<br />增加翻滚，下蹲，冲刺能力      |
| 4/27 |      | 牛头人受伤，死亡，击退效果完成<br />主项目决定去除PlayPersist场景，转而使用ScriptObject来管理短暂存储数据 |
| 4/26 |      |           研究牛头人行为树<br />完成小范围施法动作           |
| 4/24 |      |                     Pixel Adventure学完                      |
| 4/23 |      |                Pixel Adventure从30集开始学习                 |
| 4/22 |      |                       学习行为树Boss战                       |
| 4/19 |      |                Pixel Adventure从18集开始学习                 |
| 4/18 |      | Pixel Adventure从10集开始学习<br />学会下雨效果<br />练习简单像素画 |
| 4/17 |      |                 Pixel Adventure从0集开始学习                 |
| 4/16 |      |        学习冲刺跑项目案例*2，对象池和粒子系统实现方法        |
| 4/15 |      |                         跑步脚底灰尘                         |
| 4/12 |      |                      修改动画状态机内容                      |
| 4/11 |      |       对话系统<br />任务系统略看一遍<br />动画系统查询       |
| 4/10 |      |    添加小地图内容<br />添加渐变天空效果<br />修改金币系统    |
| 4/9  |      |                   学习小地图项目Roguelike                    |
| 4/8  |      |                      学习部分小地图项目                      |
| 4/7  |      |   修改背包系统内容<br />研究ScriptObject的存储方法，未完成   |
| 4/3  |      |      增加了Forest0和Forest2的场景<br />增加背包系统内容      |
| 4/1  |      |         增加了游戏存档系统<br />增加游戏音效调节功能         |
| 3/27 |      |     Menu场景搭建<br />优化场景转移代码<br />渐入渐出效果     |
| 3/26 |      |            实现场景加载转换<br />替换场景加载逻辑            |
| 3/24 |      |           优化相机工作，省去不需要的代码工作<br />           |
| 3/22 |      |  Awake启动的时候isGround不是true应该怎么办<br />UI重写绘画   |
| 3/21 |      | World Redo的白猪简单有限状态机实现，实现Idle，Patrol和Chase的逻辑 |
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

#### 2.AI模型参数

模型:高像素 像素画

关键词1:((Side view)),pixel,pixel art,pixelart,xiangsu,xiang su,8bit,16bit,32bit,64bit,128bit,256bit,1girl,full body,Solid background,masterpiece,best,quality,purple hat,witch,cute,full body,with staff,attack,muti-player

关键词2:((Side view)),pixel,pixel art,pixelart,xiangsu,xiang su,8bit,16bit,32bit,64bit,128bit,256bit,1girl,full body,Solid background,masterpiece,best,quality,warrior,cute,full body,muti-player,yellow hair,

采样方法:Euler a

迭代步数:20

高清修复:

宽高1024*1024

Lora:2D pixel Toolkit