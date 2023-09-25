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





#### MFP1项目问题

1.欧拉角

2.向量的插值是什么意思？比例计算