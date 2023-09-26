# Unity学习笔记

#### 1.生命周期方法

| 方法名称    | 调用时间                                                     |
| ----------- | ------------------------------------------------------------ |
| Awake       | 最早调用，所以一般可以在此实现单例模式                       |
| OnEnable    | 组件激活后调用，在Awake后会调用一次                          |
| Start       | 在Update之前调用一次,在OnEnable之后调用，可以在此设置一些初始值 |
| FixedUpdate | 固定频率调用方法，每次调用与上次调用的时间间隔相等           |
| Update      | 帧率调用方法，每帧调用一次，每次调用与上次调用的时间间隔不相同 |
| LateUpdate  | 在Update每调用完一次后，紧跟着调用一次                       |
| OnDisable   | 与OnEnable相反，组件未激活时调用                             |
| OnDestory   | 被销毁后调用一次                                             |

#### 2.预设体与预设体变体

#### 3.向量编程

```
public class VectorTest : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        //向量,位置，旋转，缩放
        Vector3 v = new Vector3(1,1,0.5f);
        v = Vector3.zero;
        v = Vector3.right;
        Vector3 v2 = Vector3.forward;
        //计算两个向量夹角
        Debug.Log(Vector3.Angle(v, v2));
        //计算两点之间的距离
        Debug.Log(Vector3.Distance(v, v2));
        //点乘
        Debug.Log(Vector3.Dot(v, v2));
        //叉乘
        Debug.Log(Vector3.Cross(v, v2));
        //插值
        Debug.Log(Vector3.Lerp(Vector3.zero, Vector3.one,0.5f));
        //向量的模
        Debug.Log(v.magnitude);
        //规范化的向量
        Debug.Log(v.normalized);
    }
    // Update is called once per frame
    void Update()
    {
        
    }
}
```

#### 4.旋转

欧拉角，四元数

```
void Start()
{
    //旋转：欧拉角，四元数
    Vector3 rotate = new Vector3(0, 30, 0);
    Quaternion quaternion = Quaternion.identity;
    //欧拉角转四元数
    quaternion = Quaternion.Euler(rotate);
    //四元数转为欧拉角
    rotate = quaternion.eulerAngles;
    //看向一个物体
    quaternion = Quaternion.LookRotation(new Vector3(0, 0, 0));
}
```

#### 5.Debug调试

```
public class DebugTest : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        Debug.Log("test");
        Debug.LogWarning("test2");
        Debug.LogError("test3");
    }

    // Update is called once per frame
    void Update()
    {
        //绘制一条线
        Debug.DrawLine(new Vector3(1,0,0), new Vector3(1, 1, 0), Color.blue);
        //绘制一条射线
        Debug.DrawRay(new Vector3(1, 0, 0), new Vector3(1, 1, 0), Color.red);

    }
}
```

#### 6.操作游戏物体

```
public class EmptyTest : MonoBehaviour
{
    public GameObject Cube;
    //获取预设体
    public GameObject Prefab;
    // Start is called before the first frame update
    void Start()
    {
        //拿到当前脚本所挂载的游戏物体
        //GameObject go = this.gameObject;
        //游戏物体名称
        Debug.Log(gameObject.name);
        //标签tag
        Debug.Log(gameObject.tag);
        //layer
        Debug.Log(gameObject.layer);
        //立方体名称
        Debug.Log(Cube.name);
        //真正激活状态
        Debug.Log(Cube.activeInHierarchy);
        //自身激活状态
        Debug.Log(Cube.activeSelf);
        //获取Transform组件
        //Transform trans = this.transform;
        //打印当前位置
        Debug.Log(transform.position);
        //获取其他组件
        BoxCollider bc = GetComponent<BoxCollider>();
        //获取当前物体的子物体上的某个组件
        //GetComponentInChildren<CapsuleCollider>(bc);
        //获取当前物体的父物体身上的某个组件
        //GetComponentInParent<BoxCollider>(bc);
        //添加一个组件
        gameObject.AddComponent<AudioSource>();
        //通过游戏物体的名称获取游戏物体本身
        GameObject test = GameObject.Find("Test");
 		//确定游戏物体名称
        Debug.Log(test.name);
        //通过游戏标签来获取游戏物体
        test = GameObject.FindWithTag("GameController");
        Debug.Log(test.name);
        //设置物体激活状态
        test.SetActive(false);
        //通过预设体来实例化一个游戏物体
        GameObject clone = Instantiate(Prefab,transform);
        //销毁
        Destroy(clone);
    }
    // Update is called once per frame
    void Update()
    {
        
    }
}
```

#### 7.时间

```
public class TimeTest : MonoBehaviour
{
    float timer = 0;
    void Start()
    {
        //游戏开始到现在所花的时间
        Debug.Log(Time.time);
        //时间缩放值
        Debug.Log(Time.timeScale);
        //固定时间间隔
        Debug.Log(Time.fixedDeltaTime);
    }

    // 60帧
    void Update()
    {
        timer = timer + Time.deltaTime;
        //上一帧到这一帧所用的游戏时间1/60调一次
        //Debug.Log(Time.deltaTime);
        //如果大于3秒
        if(timer > 3)
        {
            Debug.Log("大于3秒");
        }
    }
}
```

#### 8.路径

```
void Start()
{
    //游戏数据文件夹路径(只读，加密压缩)
    Debug.Log(Application.dataPath + "/新建文本文档.txt");
    //持久化文件夹路径
    Debug.Log(Application.persistentDataPath);
    //StreamingAssets文件夹路径（只读,配置打包不压缩）
    Debug.Log(Application.streamingAssetsPath);
    //临时文件夹
    Debug.Log(Application.temporaryCachePath);
    //控制是否在后台运行
    Debug.Log(Application.runInBackground);
    //浏览器打开url
    //Application.OpenURL("https://space.bilibili.com/67744423");
    //退出游戏
    Application.Quit();
}
```

#### 9.场景加载

```
void Start()
{
    //两个类，场景类，场景管理类
    //场景跳转
    //SceneManager.LoadScene("MyScene");
    //获取当前场景
    Scene scene = SceneManager.GetActiveScene();
    //场景名称
    Debug.Log(scene.name);
    //场景是否已经加载了
    Debug.Log(scene.isLoaded);
    //场景路径
    Debug.Log(scene.path);
    //场景索引
    Debug.Log(scene.buildIndex);
    //场景内所有的游戏物体
    GameObject[] gos =  scene.GetRootGameObjects();
    Debug.Log(gos.Length);

    //场景管理类
    Debug.Log(SceneManager.sceneCount);
    //创建新场景
    Scene newScene = SceneManager.CreateScene("newScene");
    //卸载场景
    SceneManager.UnloadSceneAsync(newScene);
    //加载场景替换
    SceneManager.LoadScene("MyScene", LoadSceneMode.Single);
    //加载场景叠加
    //SceneManager.LoadScene("MyScene", LoadSceneMode.Additive);
}
```

#### 10.异步跳转

```
public class AsyncTest : MonoBehaviour
{
    AsyncOperation operation;
    void Start()
    {
        StartCoroutine(loadScene());
    }
    //协程方法用来异步加载场景
    IEnumerator loadScene() {
        operation = SceneManager.LoadSceneAsync(1);
        //加载完场景后不要自动跳转
        operation.allowSceneActivation = false;
        yield return operation;
    }
    float timer = 0;
    void Update()
    {
        //输出加载进度 0-0.9
        Debug.Log(operation.progress);
        timer += Time.deltaTime;
        //如果到达5秒，再跳转
        if (timer > 5) {
            operation.allowSceneActivation = true;
        }
    }
}
```





#### MFP1项目问题

1.欧拉角

2.向量的插值是什么意思？比例计算