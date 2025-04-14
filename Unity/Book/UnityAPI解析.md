---
typora-default-code-lang: csharp
---
# UnityAPI解析

## 1.Application类

Application类常用来控制程序的运行时数据，如场景的管理、数据的加载等。

### 1.1 Applicaiton类静态属性

#### 1.1.1 Application.dataPath

```
基本语法 public static string dataPath{get;}
```

功能说明：程序的数据文件所在文件夹的路径(只读)。返回路径为相对路径，不同游戏平台的数据文件保存路径不同，具体如下所示。

| Unity Editor  | <工程文件夹所在路径>/Assets                        |
| ------------- | -------------------------------------------------- |
| Mac player    | <应用程序路径>/Contents                            |
| IPhone player | <应用程序路径>/<APPName.app>/Data                  |
| Win player    | <包含可执行文件的文件夹路径>/Data                  |
| Web player    | 播放器数据文件夹的绝对路径(没有实际的数据文件名称) |
| Flash         | 播放器数据文件夹的绝对路径(没有实际的数据文件名称) |

---

与此属性功能相近的属性有`persistentDataPath`、`streamingAssetsPath`和`temporaryCachePath`，它们的具体功能如下所示。

- `persistentDataPath`:返回一个持久化数据存储目录的路径(只读)，可以在此路径下存储一些持久化的数据文件。
- `streamingAssetsPath`:返回流数据的缓存目录，返回路径为相对路径适合设置一些外部数据文件的路径。
- `temporaryCachePath`:返回一个临时数据的缓存目录(只读)。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class DataPath_ts : MonoBehaviour
{
    void Start()
    {
        //4种不同的路径，都为只读
        //dataPath和streamingAssetsPath的路径位置一般是相对程序的安装目录位置
        //persistentDataPath和temporaryCachePath的路径位置一般是相对所在系统的固定位置
        Debug.Log("dataPath:" + Application.dataPath);
        Debug.Log("persistentDataPath:" + Application.persistentDataPath);
        Debug.Log("streamingAssetsPath:" + Application.streamingAssetsPath);
        Debug.Log("temporaryCachePath:" + Application.temporaryCachePath);
    }
}
```

#### 1.1.2 Application.loadedLevel

```
基本语法 public static int loadedLevel {get;}
```

功能说明：最后加载Scene的索引值(只读)。

代码：

```C#
using UnityEngine;
using System.Collections;
public class LoadedLevel_ts : MonoBehaviour
{
    void Start()
    {
    //返回当前场景的索引值
    Debug.Log("loadedLevel:" + Application.loadedLevel);
    //返回当前场景的名字
    Debug.Log("loadedLevelName:" + Application.loadedLevelName);
    //是否有场景正在被加载
    //在使用Application类的静态方法LoadLevel或LoadLevelAdditive加载一个新的场景时，
    //常常需要持续一段时间才能加载完毕，当场景加载完毕时，isLoadingLevel返回true，
    //否则返回false
    Debug.Log("isLoadingLevel:" + Application.isLoadingLevel);
    //返回游戏中可被加载的场景数量
    Debug.Log("levelCount:" + Application.levelCount);
    //返回当前游戏的运行平台
    //游戏的运行平台有很多种，例如手机、电脑、游戏机等，具体类型可在枚举类
    //RuntimePlatform中查看
    Debug.Log("platform:" + Application.platform);
    //当前游戏是否正在运行
    Debug.Log("isPlaying:" + Application.isPlaying);
    //当前游戏是否处于Unity编辑模式
    Debug.Log("isEditor:" + Application.isEditor);
    }
}
```

### 1.2 Applicaiton类静态方法

#### 1.2.1 CaptureScreenshot方法：截屏

```
基本语法 
(1) public static void CaptureScreenshot(string filename)
(2) public static void CaptureScreenshot(string filename,int superSize)
其中filename为截屏文件名称,superSize为放大系数，默认为0
```

功能说明：截图为PNG格式。文件保存在根目录下，如果根目录下已存在同名文件，将会被替换。当supersize大于1时，截屏文件的宽度和高度将同时被放大supersize倍。

---

提示：

- 此方法在Web模式下无效
- 当放大系数小于0时，按默认值0处理。
- 每帧只执行一次最后一次生效

---

```C#
using UnityEngine;
using System.Collections;
public class CaptureScreenshot_ts : MonoBehaviour
{
    void Update() {
    Application.CaptureScreenshot("test01.png", 0);
    Application.CaptureScreenshot("test02.png", 1);
    Application.CaptureScreenshot("test03.png", 2);
    }
}
	在Update方法中，依次调用了3次CaptureScreenshot方法，试图生成3个不同放大系统
的图片，但是在同一帧中，只有最后一次调用才能生效，故此段程序的运行结果只能
生成一张PNG图片，即test03.png。
```

#### 1.2.2 LoadLevelAdditiveAsync方法：异步加载关卡

```
基本语法 
(1) public static AsyncOperation LoadLevelAdditiveAsync(int index)
(2) public static AsyncOperation LoadLevelAdditiveAsync(string levelName)
其中index为关卡索引值，levelName是被加载的关卡名字
```

功能说明：按照关卡名字在后台异步加载关卡到当前场景中，当前场景的原有内容不会被销毁。

```
IEnumerator Start()
{
	AsyncOperation async = Application.LoadLevelAdditiveAsync("Game02");
    //异步加载中
    Debug.Log("1:" + async.isDone);//是否加载完成
    Debug.Log("2:" + async.progress);//加载进度，范围0～1
    yield return async;
    //加载完成后
    Debug.Log("3:" + async.isDone);
    Debug.Log("4:" + async.progress);
    is_load_over = async.isDone;
}
```

#### 1.2.3 RegisterLogCallback方法：注册委托

```
基本语法 
(1) public static void RegisterLogCallback(Application.LogCallback handler);
其中handle是委托方法的名字
```

功能说明：注册一个委托来调用日志信息

---

提示：RegisterLogcallbackThreaded方法与此方法的功能相似，不同之处在于RegisterlogCallbackThreaded方法是在一个新的线程中调用委托。

---

```
using UnityEngine;
using System.Collections;
public class RegisterLogCallback_ts : MonoBehaviour
{
    string output = "";//日志输出信息
    string stack = "";//堆栈跟踪信息
    string logType = "";//日志类型
    int tp = 0;
    //打印日志信息
    void Update()
    {
        Debug.Log("output:" + output);
        Debug.Log("stack:" + stack);
        Debug.Log("logType:" + logType);
        Debug.Log("tp:"+(tp++));
    }
    void OnEnable()
    {
        //注册委托
        Application.RegisterLogCallback(MyCallback);
    }
    void OnDisable()
    {
        //取消委托
        Application.RegisterLogCallback(null);
    }
    //委托方法
    //方法名字可以自定义，但方法的参数类型要符合Application.LogCallback中的参数类型
    void MyCallback(string logString, string stackTrace, LogType type)
    {
        output = logString;
        stack = stackTrace;
        logType = type.ToString();
    }
}
```

## 2.Camera类

Camera类用来控制游戏中虚拟场景的展示，以左下角为屏幕的(0.0)点坐标，以右上角为屏幕的(camera.pixelwidth，camera.pixelHeight)点坐标。如果用单位化方式表示，则左下角为(0.0)点右上角为(1.1)点。本章主要介绍了camera类的一些实例属性和实例方法，最后对摄像机的视口与aspect、pixelRect及rect之间的关系进行了注解。

### 2.1 Camera类实例属性

在Camera类中，涉及的实例属性有aspect、cameraToWorldMatrix、cullingMask、eventMask、layerCullDistances、layerCullSpherical、orthographic、pixelRect、projectionMatrix、rect、renderingPath、targetTexture和worldToCameraMatrix，下面详细介绍这些属性。

#### 2.1.1 aspect：摄像机视口比例

```
基本语法 public float aspect{get;set};
```

功能说明：获取或设置Camera视口的宽高比例值。例如，设camera.aspect=2.0f，则camera视口的宽度/高度=2.0f，但是当硬件显示器屏幕的宽度与高度比例不为2.0f时视图的显示将会发生变形。aspect只处理摄像机camera可以看到的视图的宽高比例，而硬件显示屏的作用只是把摄像机camera看到的内容显示出来，当硬件显示屏的宽高比例与aspect的比例值不同时，视图将发生变形。补充1920*1080的比例为1.777;

代码:

```C#
using UnityEngine;
using System.Collections;
public class Aspect_ts : MonoBehaviour
{
    void Start()
    {
        //camera.aspect的默认值即为当前硬件的aspect值
        Debug.Log("camera.aspect的默认值：" + camera.aspect);
    }
    void OnGUI()
    {
    	if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "aspect=1.0f"))
   		{
            //重置当前摄像机的aspect值
            camera.ResetAspect();
            camera.aspect = 1.0f;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "aspect=2.0f"))
        {
            //重置当前摄像机的aspect值
            camera.ResetAspect();
            camera.aspect = 2.0f;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "aspect还原默认值"))
        {
            //重置当前摄像机的aspect值
            camera.ResetAspect();
        }
	}
}
```

#### 2.1.2 cameraToWorldMatrix属性：变换矩阵

```
基本语法 public Matrix4×4 cameraToWorldMatrix{get};
```

功能说明：返回从摄像机的局部坐标系到世界坐标系的变换矩阵（只读）。

---

提示:Camera中的forward方向为其自身坐标系的-z轴方向，GameObject对象的forward方向为自身坐标系的z轴方向。

---

代码:

```C#
using UnityEngine;
using System.Collections;
public class CameraToWorldMatrix_ts : MonoBehaviour
{
    void Start()
    {
        Debug.Log("Camera旋转前位置：" + transform.position);
        Matrix4x4 m = camera.cameraToWorldMatrix;
        //v3的值为沿着Camera局部坐标系的-z轴方向前移5个单位的位置在世界坐标系中的位置
        Vector3 v3 = m.MultiplyPoint(Vector3.forward * 5.0f);
        //v4的值为沿着Camera世界坐标系的-z轴方向前移5个单位的位置在世界坐标系中的位置
        Vector3 v4 = m.MultiplyPoint(transform.forward * 5.0f);
        //打印v3、v4
        Debug.Log("旋转前，v3坐标值：" + v3);
        Debug.Log("旋转前，v4坐标值：" + v4);
        transform.Rotate(Vector3.up * 90.0f);
        Debug.Log("Camera旋转后位置：" + transform.position);
        m = camera.cameraToWorldMatrix;
        //v3的值为沿着Camera局部坐标系的-z轴方向前移5个单位的位置在世界坐标系中的位置
        v3 = m.MultiplyPoint(Vector3.forward * 5.0f);
        //v3的值为沿着Camera世界坐标系的-z轴方向前移5个单位的位置在世界坐标系中的位置
        v4 = m.MultiplyPoint(transform.forward * 5.0f);
        //打印v3、v4
        Debug.Log("旋转后，v3坐标值：" + v3);
        Debug.Log("旋转后，v4坐标值：" + v4);
	}
}
```

#### 2.1.3 cullingMask属性：摄像机按层渲染

```
基本语法 public int cullingMask{get;set;}
```

功能说明：按层(即GameObject.layer)有选择性地渲染场景中的物体。通过cullingMask可以使得当前摄像机有选择性地渲染场景中的部分物体，默认cu1lingMask=-1即渲染场景中任何层物体，当cullingMask=0时不渲染场景中任何层。若只渲染分别位于2、3、4层的物体，则可以使用代码cullingMask=(1<<2)+(1<<3)+(1<<4)来实现。

代码：

```
using UnityEngine;
using System.Collections;
public class CullingMask_ts : MonoBehaviour
{
    void OnGUI()
    {
        //默认CullingMask=-1，即渲染任何层
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "CullingMask=-1"))
        {
        	camera.cullingMask = -1;
        }
        //不渲染任何层
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "CullingMask=0"))
        {
        	camera.cullingMask = 0;
        }
        //仅渲染第0层
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "CullingMask=1<<0"))
        {
       		camera.cullingMask = 1 << 0;
        }
        //仅渲染第8层
        if (GUI.Button(new Rect(10.0f, 160.0f, 200.0f, 45.0f), "CullingMask=1<<8"))
        {
        	camera.cullingMask = 1 << 8;
        }
        //渲染第8层与第0层
        if (GUI.Button(new Rect(10.0f, 210.0f, 200.0f, 45.0f), "CullingMask=0&&8"))
        {
            //注：不可大意写成camera.cullingMask = 1 << 8+1;或
            //camera.cullingMask = 1+1<<8 ;因为根据运算符优先次序，其分别等价于
            //camera.cullingMask = 1 << (8+1)和camera.cullingMask = (1+1)<<8
            camera.cullingMask = (1 << 8) + 1;
        }
    }
}
```

#### 2.1.4 eventMask属性：按层响应事件

```
基本语法 public int eventMask{get;set;}
```

功能说明：选择某层(layer)的物体响应鼠标事件。

- 第一，物体在摄像机的视野范围内。
    第二，在2的layer次方的值与eventMask进行与运算(&)后结果为仍为2的1ayer次方的值，例如当前物体的层为Default，即laver值为0，则2的0次方为1，如果1与eventMask进行与运算后结果仍为1,则此物体便会响应鼠标事件。由于当eventMask为奇数时，与1的与运算结果都为1，所以若物体的层为Default并且eventMask为奇数时物体便会响应鼠标事件。
- 如果想要多个不同层的物体都响应鼠标事件，则需要把所有层的2的layer次方值相加，再与eventMask做与运算。例如，现有两个物体，它们的layer值分别为1和3，则当eventMask与9(因为2^0^+2^3^=9)进行与运算后若结果仍为9，则这两个物体都会响应鼠标事件。
- 此属性有一个特殊情况，当物体的layer选择IgnoreRaycast(其为系统内置，值为2)时，无论eventMask值为多少，物体都无法响应鼠标事件。原因很简单，因为这个层的物体都忽略了射线碰撞，鼠标就探测不到物体的存在,因而也就无法响应鼠标事件了。

---

提示：此属性的计算比较消耗资源，在不使用鼠标事件时建议将值设为零。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class EventMask_ts : MonoBehaviour
{
    bool is_rotate = false;//控制物体旋转
    public Camera c;//指向场景中的摄像机
    //记录摄像机的eventMask值，可以在程序运行时在Inspector面板中修改值的大小
    public int eventMask_now = -1;
    //记录当前物体的层
    int layer_now;
    int tp;//记录2的layer次方的值
    int ad;//记录与运算（&）的结果string str = null;
    void Update()
    {
        //记录当前对象的层，可以在程序运行时在Inspector面板中选择不同的层
        layer_now = gameObject.layer;
        //求2的layer_now次方的值
        tp = (int)Mathf.Pow(2.0f, layer_now);
        //与运算（&）
        ad = eventMask_now & tp;
        c.eventMask = eventMask_now;
        //当is_rotate为true时旋转物体
        if (is_rotate)
        {
        	transform.Rotate(Vector3.up * 15.0f * Time.deltaTime);
        }
    }
    //当鼠标左键按下时，物体开始旋转
    void OnMouseDown()
    {
        is_rotate = true;
    }
    //当鼠标左键抬起时，物体结束旋转
    void OnMouseUp()
    {
        is_rotate = false;
    }
    void OnGUI()
    {
        GUI.Label(new Rect(10.0f, 10.0f, 300.0f, 45.0f), "当前对象的layer值为：" + layer_now
        + " , 2的layer次方的值为" + tp);
        GUI.Label(new Rect(10.0f, 60.0f, 300.0f, 45.0f), "当前摄像机eventMask的值为：" +
        eventMask_now);
        GUI.Label(new Rect(10.0f, 110.0f, 500.0f, 45.0f), "根据算法，当eventMask的值与" + tp
        + "进行与运算（&）后， 若结果为" + tp + "，则物体响应OnMousexxx方法，否则不响应！
        ");
        if (ad == tp)
        {
        	str = " ,所以物体会响应OnMouseXXX方法！";
        }
        else
        {
       		str = " ,所以物体不会响应OnMouseXXX方法！";
        }
        GUI.Label(new Rect(10.0f, 160.0f, 500.0f, 45.0f), "而当前eventMask与" + tp + "进行与运算（&）的结果为" + ad + str);
    }
}
```

#### 2.1.5 layerCullDistances属性：层消隐的距离

```
基本语法 public float[] layerCullDistances{get;set;}
```

功能说明：设置摄像机基于层的消隐距离。摄像机可以通过基于层(GameObject.layer)的方式来设置不同层物体的消隐距离，但这个距离必须小于或等于摄像机的farClipPlane才有效。

代码：

```c#
using UnityEngine;
using System.Collections;
public class LayerCullDistances_ts : MonoBehaviour
{
    public Transform cb1;
    void Start()
    {
        //定义大小为32的一维数组，用来存储所有层的剔除距离
        float[] distances = new float[32];
        //设置第9层的剔除距离
        distances[8] = Vector3.Distance(transform.position, cb1.position);
        //将数组赋给摄像机的layerCullDistances
        camera.layerCullDistances = distances;
    }
    void Update()
    {
        //摄像机远离物体
        transform.Translate(transform.right * Time.deltaTime);
    }
}
```

#### 2.1.6 layerCullSpherical属性：基于球面距离剔除

```
基本语法 public bool layerCullSpherical{get;set;}
```

功能说明：设置摄像机在基于层剔除物体时,是否采用基于球面距离的剔除方式。默认值为false，即不使用球面剔除方式，此时，只要物体表面上有一点没有超出物体所在层的远视口平面，物体就是可见的。当layercullSpherical为true时，只要物体的世界坐标点position与摄像机的距离大于所在层的剔除距离，物体就是不可见的，这和值为false时的计算方式不同。

代码：

```C#
using UnityEngine;
using System.Collections;
public class layerCullSpherical_ts : MonoBehaviour
{
    public Transform cb1, cb2, cb3;
    void Start()
    {
        //定义大小为32的一维数组，用来存储所有层的剔除距离
        float[] distances = new float[32];
        //设置第9层的剔除距离
        distances[8] = Vector3.Distance(transform.position, cb1.position);
        //将数组赋给摄像机的layerCullDistances
        camera.layerCullDistances = distances;
        //打印出3个物体距离摄像机的距离
        Debug.Log("Cube1距离摄像机的距离：" + Vector3.Distance(transform.position, cb1.position));
        Debug.Log("Cube2距离摄像机的距离：" + Vector3.Distance(transform.position, cb2.position));
        Debug.Log("Cube3距离摄像机的距离：" + Vector3.Distance(transform.position, cb3.position));
    }
    
    void OnGUI()
    {
        //使用球形距离剔除
        if (GUI.Button(new Rect(10.0f, 10.0f, 180.0f, 45.0f), "use layerCullSpherical"))
        {
            camera.layerCullSpherical = true;
        }
        //取消球形距离剔除
        if (GUI.Button(new Rect(10.0f, 60.0f, 180.0f, 45.0f), "unuse layerCullSpherical"))
        {
       		camera.layerCullSpherical = false;
    	}
    }
}
```

#### 2.1.7 orthgraphic属性：摄像机投影模式

```
基本语法 public bool orthographic{get;set;}
```

功能说明：获取或设置当前摄像机的投影模式，投影模式包括正交投影模式(orthographic)和透视投影模式(perspective)。若值为true则为正交投影，false为透视投影。

代码：

```C#
using UnityEngine;
using System.Collections;
public class orthographic_ts : MonoBehaviour
{
    float len = 5.5f;
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 120.0f, 45.0f), "正交投影"))
        {
            camera.orthographic = true;
            len = 5.5f;
        }
        if (GUI.Button(new Rect(150.0f, 10.0f, 120.0f, 45.0f), "透视投影"))
        {
            camera.orthographic = false;
            len = 60.0f;
        }
        if (camera.orthographic)
        {
            //正交投影模式下，物体没有远大近小的效果，
            //orthographicSize的大小无限制，当orthographicSize为负数时视口的内容会颠倒，
            //orthographicSize的绝对值为摄像机视口的高度值，即上下两条边之间的距离
            len = GUI.HorizontalSlider(new Rect(10.0f, 60.0f, 300.0f, 45.0f), len, -20.0f,
            20.0f);
            camera.orthographicSize = len;
        }
        else
        {
            //透视投影模式下，物体有远大近小的效果
            //fieldOfView的取值范围为1.0-179.0
            len = GUI.HorizontalSlider(new Rect(10.0f, 60.0f, 300.0f, 45.0f), len, 1.0f,
            179.0f);
            camera.fieldOfView = len;
        }
        //实时显示len大小
        GUI.Label(new Rect(320.0f, 60.0f, 120.0f, 45.0f), len.ToString());
    }
}
```

#### 2.1.8 pixelRect属性：摄像机渲染区间

```
基本语法 public Rect pixelRect{get;set;}
```

功能说明：设置camera被渲染到屏幕中的坐标位置。pixelRect与属性rect功能类似，不同的是pixelRect以实际像素大小来设置显示视口的位置，而rect以单位化方式设置显示视口的位置。设camera.pixelRect的值为(x~0~,y~0~,w,h)，如图2-5所示，A为原始平面大小，B为变换后的视口大小，则x~0~的值为视口右移的像素大小，y~0~的值为视口上移的像素大小，w的值为camera.pixelWidth，h的值为camera.pixelHeight。

代码：

```C#
using UnityEngine;
using System.Collections;
public class PixelRect_ts : MonoBehaviour
{
    int which_change = -1;
    float temp_x = 0.0f, temp_y = 0.0f;
    void Update()
    {
        //Screen.width和Screen.height为模拟硬件屏幕的宽高值
        //其返回值不随camera.pixelWidth和camera.pixelHeight的改变而改变
        Debug.Log("Screen.width:" + Screen.width);
        Debug.Log("Screen.height:" + Screen.height);
        Debug.Log("pixelWidth:" + camera.pixelWidth);
        Debug.Log("pixelHeight:" + camera.pixelHeight);
        //通过改变Camera的坐标位置而改变视口的区间
        if (which_change == 0)
        {
            if (camera.pixelWidth > 1.0f)
            {
            temp_x += Time.deltaTime * 20.0f;
            }
            //取消以下注释查看平移状况
            //if (camera.pixelHeight > 1.0f)
            //{
            // temp_y += Time.deltaTime * 20.0f;
            //}
            camera.pixelRect = new Rect(temp_x, temp_y, camera.pixelWidth,
            camera.pixelHeight);
        }
        //通过改变Camera的视口宽度和高度来改变视口的区间
        else if (which_change == 1)
        {
            if (camera.pixelWidth > 1.0f)
            {
            	temp_x = camera.pixelWidth - Time.deltaTime * 20.0f;
            }
            //取消以下注释查看平移状况
            //if (camera.pixelHeight > 1.0f)
            //{
            // temp_y = camera.pixelHeight - Time.deltaTime * 20.0f
            //}
        camera.pixelRect = new Rect(0, 0, temp_x, temp_y);
        }
    }
    void OnGUI()
    {
    	if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "视口改变方式1"))
        {
            camera.rect = new Rect(0.0f, 0.0f, 1.0f, 1.0f);
            which_change = 0;
            temp_x = 0.0f;
            temp_y = 0.0f;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "视口改变方式2"))
        {
            camera.rect = new Rect(0.0f, 0.0f, 1.0f, 1.0f);
            which_change = 1;
            temp_x = 0.0f;
            temp_y = camera.pixelHeight;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "视口还原"))
        {
            camera.rect = new Rect(0.0f, 0.0f, 1.0f, 1.0f);
            which_change = -1;
        }
    }
}
```

#### 2.1.9 projectionMatrix属性：自定义投影矩阵

```
基本语法 public Matrix4X4 projectionMatrix{get;set;}
```

功能说明：设置摄像机的自定义投影矩阵。常在一些特效场景下用到，在切换变换矩阵时通常需要先用camera.ResetProiectionMatrix()重置camera的变换矩阵。

```C#
using UnityEngine;
using System.Collections;
public class ProjectionMatrix_ts : MonoBehaviour
{
    public Transform sp, cb;
    public Matrix4x4 originalProjection;
    float q = 0.1f;//晃动振幅
    float p = 1.5f;//晃动频率
    int which_change = -1;
    void Start()
    {
        //记录原始投影矩阵
        originalProjection = camera.projectionMatrix;
    }
    
    void Update()
    {
        sp.RotateAround(cb.position, cb.up, 45.0f * Time.deltaTime);
        Matrix4x4 pr = originalProjection;
        switch (which_change)
        {
        	case -1:
        		break;
        	case 0:
                //绕摄像机x轴晃动
                pr.m11 += Mathf.Sin(Time.time * p) * q;
                break;
        	case 1:
                //绕摄像机y轴晃动
                pr.m01 += Mathf.Sin(Time.time * p) * q;
                break;
        	case 2:
                //绕摄像机z轴晃动
                pr.m10 += Mathf.Sin(Time.time * p) * q;
                break;
        	case 3:
                //绕摄像机左右移动
                pr.m02 += Mathf.Sin(Time.time * p) * q;
                break;
        	case 4:
                //摄像机视口放缩运动
                pr.m00 += Mathf.Sin(Time.time * p) * q;
                break;
        }
        //设置Camera的变换矩阵
        camera.projectionMatrix = pr;
    }
    
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "绕摄像机x轴晃动"))
        {
            //重置摄像机的投影矩阵
            camera.ResetProjectionMatrix();
            which_change = 0;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "绕摄像机y轴晃动"))
        {
            camera.ResetProjectionMatrix();which_change = 1;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "绕摄像机z轴晃动"))
        {
            camera.ResetProjectionMatrix();
            which_change = 2;
        }
        if (GUI.Button(new Rect(10.0f, 160.0f, 200.0f, 45.0f), "绕摄像机左右移动"))
        {
            camera.ResetProjectionMatrix();
            which_change = 3;
        }
        if (GUI.Button(new Rect(10.0f, 210.0f, 200.0f, 45.0f), "视口放缩运动"))
        {
            camera.ResetProjectionMatrix();
            which_change = 4;
        }
    }
}
```

#### 2.1.10 rect属性：摄像机视图的位置和大小

```
基本语法 public Rect rect{get;set;}
```

功能说明：使用单位化坐标系的方式来设置当前摄像机的视图位置和大小。如图2-6所示，A为原始视口大小，B为移位及放缩后的视口位置和大小。设camera.rect = x~0~,y~0~,w~0~.h~0~，则

- X~0~的值为A物体向右移动了A的总宽度的m,倍，例如X~0~=0.1f则表示A向右移动的距离为x~1~=0.1*w~1~。y~0~与此类似，不再赘述。通过设定x~0~和y~0~的值可以设置移动后B的左下角的位置。
- w~0~和h~0~的值分别为w~0~=$\frac{w_2}{w_1}$ ,h~0~=$\frac{h_2}{h_1}$ ，通过设定w~0~和h~0~的值可以设置A物体被放缩的比例。当w~0~与h~0~相等时才能实现对A物体的等比例放缩。
- x~0~和y~0~的有效范围为[-1,1]；w~0~和h~0~的有效范围为[0,1]。

代码：

```
using UnityEngine;
using System.Collections;
public class Rect_ts : MonoBehaviour
{
    int which_change = -1;
    float temp_x = 0.0f, temp_y = 0.0f;
    void Update()
    {
    //视口平移
    if (which_change == 0)
    {
        if (camera.rect.x < 1.0f)
        {
            //沿着x轴平移
            temp_x = camera.rect.x + Time.deltaTime * 0.2f;
        }
        //取消下面注释查看平移的变化
        //if (camera.rect.y< 1.0f)
        //{
            //沿着y轴平移
            // temp_y = camera.rect.y + Time.deltaTime * 0.2f
        //}
        camera.rect = new Rect(temp_x, temp_y, camera.rect.width, camera.rect.height);
    }
    //视口放缩
    else if (which_change == 1)
    {
        if (camera.rect.width > 0.0f)
        {
            //沿着x轴放缩
            temp_x = camera.rect.width - Time.deltaTime * 0.2f;
            }
            if (camera.rect.height > 0.0f)
            {
            	//沿着y轴放缩
                temp_y = camera.rect.height - Time.deltaTime * 0.2f;
            }
            camera.rect = new Rect(camera.rect.x, camera.rect.y, temp_x, temp_y);
        }
    }

    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "视口平移"))
        {
            //重置视口
            camera.rect = new Rect(0.0f, 0.0f, 1.0f, 1.0f);
            which_change = 0;
            temp_y = 0.0f;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "视口放缩"))
        {
            camera.rect = new Rect(0.0f, 0.0f, 1.0f, 1.0f);
            which_change = 1;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "视口还原"))
        {
            camera.rect = new Rect(0.0f, 0.0f, 1.0f, 1.0f);
            which_change = -1;
        }
    }
}
```

#### 2.1.11 renderingPath属性：渲染路径

```
基本语法 public RenderingPath renderingPath{get;set;}
```

功能说明：获取或设置摄像机的渲染路径。Unity中渲染路径RenderingPath为枚举类型，共有以下4种设置方式。

- UsePlayersettings:使用工程中的设置，即Edit→Projectsettings→Player中各个平台下的设置，如图2-7所示。
- VertexLit:使用顶点光照，最低消耗的渲染路径，不支持实时阴影，适用于移动及老式设备。
- Forward:使用正向光照，基于着色器的渲染路径，支持逐像素计算光照(包括法线贴图和灯光Cookies)和来自一个平行光的实时阴影，具体请参考官方文档。
- DeferredLighting:使用延迟光照，支持实时阴影，计算消耗大，对硬件要求高，不支持移动设备，仅专业版可用。

代码：

```C#
using UnityEngine;
using System.Collections;
public class renderingPath_ts : MonoBehaviour
{
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 120.0f, 45.0f), "UsePlayerSettings"))
        {
        	camera.renderingPath = RenderingPath.UsePlayerSettings;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 120.0f, 45.0f), "VertexLit"))
        {
            //无凹凸纹理，无投影
            camera.renderingPath = RenderingPath.VertexLit;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 120.0f, 45.0f), "Forward"))
        {
            //有凹凸纹理，只能显示一个投影
            camera.renderingPath = RenderingPath.Forward;
        }
        if (GUI.Button(new Rect(10.0f, 160.0f, 120.0f, 45.0f), "DeferredLighting"))
        {
            //有凹凸纹理，可以显示多个投影
            camera.renderingPath = RenderingPath.DeferredLighting;
        }
    }
}
```

#### 2.1.12 targetTexture属性：目标渲染纹理

```
基本语法 public RenderTexture targetTexture{get;set;}
```

功能说明：调用可以生成目标渲染纹理。把某个摄像机的视图作为Renderer Texture,然后添加到一个Material，再把这个material赋给一个游戏对象的Renderer组件，就可以在这个游戏对象中实时地看到摄像机A中的视图，可以用来做一些实时跟踪类的功能。

```C#
using UnityEngine;
using System.Collections;
public class Target_ts : MonoBehaviour {
    public Transform ts;
    void Update () 
    {
    	transform.RotateAround(ts.position,ts.up,30.0f * Time.deltaTime);
    }
}
```

#### 2.1.13 worldCameraMatrix属性：变换矩阵

```
基本语法 public Matrix4x4 worldToCameraMatrix{get;set;}
```

功能说明：返回或设置从世界坐标系到当前camera自身坐标系的变换矩阵。当用camera.worldToCameraMatrix重设摄像机的转换矩阵时，摄像机对应的Transform组件数据不会同步更新，如果想回到Transformm的可控状态，需要调用ResetWorldToCameraMatrix方法重置摄像机的转换矩阵。

```C#
using UnityEngine;
using System.Collections;
public class WorldToCameraMatrix_ts : MonoBehaviour
{
    public Camera c_test;
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "更改变换矩阵"))
        {
            //使用c_test的变换矩阵
            camera.worldToCameraMatrix = c_test.worldToCameraMatrix;
            //也可使用如下代码实现同样功能
            // camera.CopyFrom(c_test)
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "重置变换矩阵"))
        {
            //重置摄像机的转换矩阵，常和属性worldToCameraMatrix配合使用
            camera.ResetWorldToCameraMatrix();
        }
    }
}
```

### 2.2 Camera类实例方法

在Camera类中，涉及的实例方法有`RenderToCubemap`、`RenderWithShader`、`ScreenPointToRay`、`ScreenToViewportPoint `、`ScreenToWorldPoint `、`SetTargetBuffers` 、`ViewportPointToRay `、
`ViewportToWorldPoint`、`WorldToScreenPoint`和`WorldToViewportPoint`，下面详细介绍这些方法。

#### 2.2.1 RenderToCubemap方法：生成Cubemap静态贴图

```
基本语法 （1）public bool RenderToCubemap(Cubemap cubemap)
			其中参数cubemap为Cubemap静态贴图
    	（2）public bool RenderToCubemap(RenderTexture cubemap);
    		其中参数cubemap为RenderTexture静态贴图
    	（3）public bool RenderToCubemap(Cubemap cubemap,int faceMask)
    		其中参数cubemap为Cubemap静态贴图,faceMask为反射面数量，默认值为63
    	（4）public bool RenderToCubemap(RenderTexture cubemap,int faceMask)
    		其中参数cubemap为RenderTexture静态贴图，faceMask为反射面数量，默认值为63
```

功能说明：此方法的作用是使用摄像机生成一个Cubemap静态贴图。当faceMask值为63时，表示Cubemap的上下左右前后6个面全部反射，这种情况下系统计算耗费也最大。该值是一个二进制计算的参数，默认值为63即111111，表示6个面全开，如果不需要全部反射，则需要修改faceMask的值。

代码：

```C#
using UnityEngine;
using System.Collections;
public class RenderToCubemap_ts : MonoBehaviour
{
    public Cubemap cp;
    void Start()
    {
        camera.RenderToCubemap(cp);
    }
}
```

#### 2.2.2 RenderWithShader方法：使用其他shader渲染

```
基本语法 public void RenderWithShader(Shader shader,string replacementTag)
		其中参数shader为要使用的shader,replacementTage为shader的Tag标示
```

功能说明：使用指定shader来代替当前物体的shader渲染一帧。当replacementTag为空时会替换视口中所有物体的shader。

---

提示：SetReplacementShader方法与此方法功能相近，不同之处是，SetReplacementshader方法使用指定的shader来替换物体当前的shader,被替换后每一帧都会用替换的shader来渲染物体，而不是只渲染一帧。

提示：方法RenderWithShader每调用一次只渲染一帧，故不可直接将其放到GUI的Button中，
否则看不出效果。

---

代码：

```C#
using UnityEngine;
using System.Collections;

public class RenderWithShader_ts : MonoBehaviour
{
    bool is_use = false;
    void OnGUI()
    {
        if (is_use)
        {
            //使用高光shader：Specular来渲染Camera
            camera.RenderWithShader(Shader.Find("Specular"), "RenderType");
        }
        if (GUI.Button(new Rect(10.0f, 10.0f, 300.0f, 45.0f), "使用RenderWithShader启用高光"))
        {
            //RenderWithShader每调用一次只渲染一帧，所以不可将其直接放到这儿
            //camera.RenderWithShader(Shader.Find("Specular"), "RenderType")
            is_use = true;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 300.0f, 45.0f), "使用SetReplacementShader
        启用高光"))
        {
            //SetReplacementShader方法用来替换已有shader，调用一次即可
            camera.SetReplacementShader(Shader.Find("Specular"), "RenderType");
            is_use = false;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 300.0f, 45.0f), "关闭高光"))
        {
            //重置摄像机的shader渲染模式
            camera.ResetReplacementShader();
            is_use = false;
        }
    }
}
```

#### 2.2.3 ScreenPointToRay方法：近视口到屏幕的射线

```
基本语法 public void ScreenPointToRay(Vector3 poisiton)
		其中参数position为屏幕位置参考点
```

功能说明：从Camera的近视口nearClip向前发射一条射线到屏幕上的position点。Position用实际像素值的方式来决定Ray到屏幕的位置。Position的x/y轴从0增长到最大值时，Ray从屏幕一边移动到另一边。当Ray未能碰撞到物体时，hit.point返回值为Vector3(0.0.0)。Position的z轴值无效。

```C#
using UnityEngine;
using System.Collections;
public class ScreenPointToRay_ts : MonoBehaviour
{
    Ray ray;
    RaycastHit hit;
    Vector3 v3 = new Vector3(Screen.width / 2.0f, Screen.height / 2.0f, 0.0f);
    Vector3 hitpoint = Vector3.zero;
    void Update()
    {
        //射线沿着屏幕x轴从左向右循环扫描
        v3.x = v3.x >= Screen.width ? 0.0f : v3.x + 1.0f;
        //生成射线
        ray = camera.ScreenPointToRay(v3);
        if (Physics.Raycast(ray, out hit, 100.0f))
        {
            //绘制线，在Scene视图中可见
            Debug.DrawLine(ray.origin, hit.point, Color.green);
            //输出射线探测到的物体的名称
            Debug.Log("射线探测到的物体名称：" + hit.transform.name);
        }
    }
}
```

#### 2.2.4 ScreenToViewportPoint方法: 坐标系转换

```
基本语法 public Vector3 ScreenToViewportPoint(Vector3 postion);
		其中参数position为屏幕参考点
```

功能说明：实现坐标点position从屏幕坐标系向摄像机视口的单位化坐标系转换。参考点position的x和y分量为屏幕的实际坐标值，单位为像素，z值无效。

代码：

```
using UnityEngine;
using System.Collections;
public class ScreenToViewportPoint_ts : MonoBehaviour
{
    void Start()
    {
        transform.position = new Vector3(0.0f, 0.0f, 1.0f);
        transform.rotation = Quaternion.identity;
        //从屏幕的实际坐标点向视口的单位化比例值转换
        Debug.Log("1:" + camera.ScreenToViewportPoint(new Vector3(Screen.width / 2.0f,
        Screen.height / 2.0f, 100.0f)));
        //从视口的单位化比例值向屏幕的实际坐标点转换
        Debug.Log("2:" + camera.ViewportToScreenPoint(new Vector3(0.5f, 0.5f, 100.0f)));
        Debug.Log("屏幕宽：" + Screen.width + " 屏幕高：" + Screen.height);
    }
}
```

#### 2.2.5 ScreenToWorldPoint方法: 坐标系转换

```
基本语法 public Vector3 ScreenToWorldPoint(Vector3 postion);
		其中参数position为屏幕参考点
```

功能说明：将参考点position从屏幕坐标系转换到世界坐标系。此方法与方法ViewportToWorldPoint功能类似，只是此方法的参考点position中各个分量值都为实际单位像素值，而非比例值。

```
例如执行如下代码后，
    Vector v3= camera. ScreenToWorldPoint (pos);//ps为已知参考点
v3的各个分量值为：
    v3.x= camera.transform.position.x+ pos.z* asp * tan(e/2);
    v3.y= camera.transform.position.y+ pos.z* tan(e/2);
    v3.z= camera.transform.position.z +pos.z;
其中pos为已知参考值，e为摄像机的视口夹角fieldOfView的值，asp为摄像机视口的宽高比例值aspect。更多内容请参考方法ViewportToWorldPoint的功能说明。
```

代码：

```C#
using UnityEngine;
using System.Collections;
public class ScreenToViewportPoint_ts : MonoBehaviour
{
    void Start()
    {
        transform.position = new Vector3(0.0f, 0.0f, 1.0f);
        transform.rotation = Quaternion.identity;
        //从屏幕的实际坐标点向视口的单位化比例值转换
        Debug.Log("1:" + camera.ScreenToViewportPoint(new Vector3(Screen.width / 2.0f,
        Screen.height / 2.0f, 100.0f)));
        //从视口的单位化比例值向屏幕的实际坐标点转换
        Debug.Log("2:" + camera.ViewportToScreenPoint(new Vector3(0.5f, 0.5f, 100.0f)));
        Debug.Log("屏幕宽：" + Screen.width + " 屏幕高：" + Screen.height);
    }
}
```

#### 2.2.6 SetTargetBuffers方法:重设摄像机到TargetTexture的渲染

```
基本语法(1)public void SetTargetBuffers(RenderBuffer colorBuffer,RenderBuffer depth Buffer);
		其中参数colorBuffer为纹理的颜色缓存,depthBuffer为纹理的深度缓存
	   (2)public void SetTargetBuffers(RenderBuffer[] colorBuffer,RenderBuffer depth Buffer);
	    其中参数colorBuffer为纹理的颜色缓存,depthBuffer为纹理的深度缓存。此重载方法可以将摄像机的渲染一次赋给多个colorBuffer
```

功能说明：此方法用于将camera的渲染赋给RenderTexture的colorBuffer和depthBuffer

```C#
using UnityEngine;
using System.Collections;
public class SetTargetBuffers_ts : MonoBehaviour
{
    //声明两个RendererTexture变量
    public RenderTexture RT_1, RT_2;
    public Camera c;//指定Camera
    void OnGUI()
    {
        //设置RT_1的buffer为摄像机c的渲染
        if (GUI.Button(new Rect(10.0f, 10.0f, 180.0f, 45.0f), "set target buffers"))
        {
            c.SetTargetBuffers(RT_1.colorBuffer, RT_1.depthBuffer);
        }
        //设置RT_2的buffer为摄像机c的渲染，此时RT_1的buffer变为场景中Camera1的渲染
        if (GUI.Button(new Rect(10.0f, 60.0f, 180.0f, 45.0f), "Reset target buffers"))
        {
            c.SetTargetBuffers(RT_2.colorBuffer, RT_2.depthBuffer);
        }
    }
}
```

#### 2.2.7 ViewportPointToRay方法：近视口到屏幕的射线

```
基本语法 public Ray ViewportPointToRay(Vector3 position);
		其中参数position为单位化坐标中的参考点
```

功能说明：从Camera的近视口nearclip向前发射一条射线到屏幕上的position点。position用单位化的方式来决定Ray到屏幕的位置。position的x轴或y轴从0到1增长时，Ray从屏幕一边移动到另一边。当Ray未能碰撞到物体时，hit.point的返回值为Vector3(0.0.0)。position的z轴值无效。

代码：

```
using UnityEngine;
using System.Collections;
public class ViewportPointToRay_ts : MonoBehaviour
{
    Ray ray;//射线
    RaycastHit hit;
    Vector3 v3 = new Vector3(0.5f, 0.5f, 0.0f);
    Vector3 hitpoint = Vector3.zero;
    void Update()
    {
        //射线沿着屏幕x轴从左向右循环扫描
        v3.x = v3.x >= 1.0f ? 0.0f : v3.x + 0.002f;
        //生成射线
        ray = camera.ViewportPointToRay(v3);
        if (Physics.Raycast(ray, out hit, 100.0f))
        {
            //绘制线，在Scene视图中可见
            Debug.DrawLine(ray.origin, hit.point, Color.green);
            //输出射线探测到的物体的名称
            Debug.Log("射线探测到的物体名称：" + hit.transform.name);
        }
    }
}
```

#### 2.2.8 ViewportToWorldPoint方法：坐标点的坐标系转换

```
基本语法 public Vector3 ViewportToWorldPoint(Vector3 position);
		其中参数position为单位化坐标中的参考点
```

功能说明：从Camera视口坐标点向世界坐标点转换，此方法与WorldToViewportPoint正好相反。方法的返回值大小受当前camera在世界坐标系中的位置Camera的fieldofView值以及参考点position的共同影响。其中position的x和y的有效范围为[0.0.1.0]为比例值。而z值为实际单位值，非比例值。对此方法的算法说明如下

​	设o为摄像机坐标点，asp为摄像机的视口宽高比例值aspect，e为摄像机的视口夹角fieldofView值，oA的模长为position的z值，position的y值为y~0~，则此方法执行后的返回值的z值大小为:o点的z轴分量值加上oA的长度;返回值的y分量值为o点的y轴分量值加上K~1~，其中K~1~的计算方法为:

$$
K_1 = |oA| \times \left( \frac{y_0 - 0.5}{0.5} \right) \times \tan\left( \frac{e}{2} \right)
$$
返回值的x分量值为o点的x轴分量值加上K2，其中K2的计算方法为：


$$
K_2 = |oA| \times asp \times \left( \frac{y_0 - 0.5}{0.5} \right) \times \tan\left( \frac{e}{2} \right)
$$
例如执行如下代码后，

Vector v3= camera.ViewportToWorldPoint(ps);//ps为已知参考点

v3的各个分量值为：

v3.x= camera.transform.position.x+ ps.z * asp * (( ps.x-0.5)/0.5)* tan(e/2);
v3.y= camera.transform.position.y+ ps.z * (( ps.y-0.5)/0.5)* tan(e/2);
v3.z= camera.transform.position.z +ps.z;

其中ps为已知参考值，e为摄像机的视口夹角fieldOfView的值，asp为摄像机视口的宽
高比例值aspect。

代码：

```C#
using UnityEngine;
using System.Collections;
public class ViewportToWorldPoint_ts : MonoBehaviour
{
    void Start()
    {
        transform.position = new Vector3(1.0f, 0.0f, 1.0f);
        camera.fieldOfView = 60.0f;
        camera.aspect = 16.0f / 10.0f;
        //屏幕左下角
        Debug.Log("1:" + camera.ViewportToWorldPoint(new Vector3(0.0f, 0.0f, 100.0f)));
        //屏幕中间
        Debug.Log("2:" + camera.ViewportToWorldPoint(new Vector3(0.5f, 0.5f, 100.0f)));
        //屏幕右上角
        Debug.Log("3:" + camera.ViewportToWorldPoint(new Vector3(1.0f, 1.0f, 100.0f)));
    }
}
```

#### 2.2.9 WorldToScreenPoint方法：坐标点的坐标系转换

```
基本语法 public Vector3 WorldToScreenPoint(Vector3 position);
		其中参数position为单位化坐标中的参考点
```

功能说明：从世界坐标点向屏幕坐标点转换，即position投射到屏幕上的坐标值。返回值的x和y分量是以屏幕左下角为(0,0)点，以向上为y轴、向右为x轴来计算的。

代码：

```C#
using UnityEngine;
using System.Collections;
public class WorldToScreenPoint_ts : MonoBehaviour
{
    public Transform cb, sp;
    public Texture2D t2;
    Vector3 v3 = Vector3.zero;
    float sg;
    void Start()
    {
        //记录屏幕高度
        sg = Screen.height;
    }
    void Update()
    {
        //sp绕着cb的y轴旋转
        sp.RotateAround(cb.position, cb.up, 30.0f * Time.deltaTime);
        //获取sp在屏幕上的坐标点
        v3 = camera.WorldToScreenPoint(sp.position);
    }
    void OnGUI()
    {
        //绘制纹理
        GUI.DrawTexture(new Rect(0.0f, sg - v3.y, v3.x, sg), t2);
    }
}
```

#### 2.2.10 WorldToViewportPoint方法：坐标点的坐标系转换

```
基本语法 public Vector3 WorldToViewportPoint(Vector3 position);
		其中参数position为待转换的世界坐标系中的坐标点
```

功能说明：把三维坐标点position从世界坐标系转换到屏幕的单位化坐标系中即世界坐标点position投射到屏幕上的坐标点的x、y分量所占屏幕宽高的比例大小。此方法与WorldToScreenPoint功能类似,不同的是其返回值的x和y分量是比例值以屏幕的总宽度和总高度分别为x和y分量的最大值1。返回值的x和y分量是以屏幕左下角为(0.0)点，以向上为y轴、向右为x轴来计算的。

代码：

```C#
using UnityEngine;
using System.Collections;
public class WorldToViewportPoint_ts : MonoBehaviour
{
    public Transform cb, sp;
    public Texture2D t2;
    Vector3 v3 = Vector3.zero;
    float sw, sh;
    void Start()
    {
        //记录屏幕的宽度和高度
        sw = Screen.width;
        sh = Screen.height;
    }
    void Update()
    {
        //物体sp始终绕cb的y轴旋转
        sp.RotateAround(cb.position, cb.up, 30.0f * Time.deltaTime);
        //记录sp映射到屏幕上的比例值
        v3 = camera.WorldToViewportPoint(sp.position);
    }
    void OnGUI()
    {
        //绘制纹理，由于方法WorldToViewportPoint的返回值的y分量是从屏幕下方向上方递增的,
        //所以需要先计算1.0f - v3.y的值，然后再和sh相乘
        GUI.DrawTexture(new Rect(0.0f, sh * (1.0f - v3.y), sw * v3.x, sh), t2);
    }
}
```

### 2.3 关于Camera视口、aspect、pixelRect及rect的关系注解

搞清楚摄像机的视口与aspect、pixelRect及rect之间的关系有助于对camera类的理解和使用，下面对它们之间的关系进行说明。

- Camera视口用来记录当前摄像机能看到场景中的哪些内容，其大小及位置是可以改变的。而屏幕视口是指当前硬件的屏幕,对于一个固定的硬件(如手机),它的屏幕视口大小(即分辨率)是固定的。Camera视口的内容不一定可以完全显示在屏幕上，屏幕可能只显示了一部分视口内容，也可能对视口内容进行了放缩。可以简单理解为，camera视口是一张二维图片，而屏幕是用来显示这张图片的，图片可能被剪切，也可能被压缩。Camera视口的内容显示到屏幕上的方式由很多因素决定。
- Unity的Game面板中的aspect选项(如图2-19所示)是用来模拟硬件屏幕的，可分为3类:全屏显示、固定比例显示和固定分辨率显示。全屏方式即以当前Game屏幕的大小来模拟硬件屏幕分辨率，其camera视口即为当前摄像机的默认状态。而在固定比例方式则会改变amera视口的宽高比例，其大小不固定。而在固定分辨率方式下，视口的最大宽度和高度是固定的，当Game视口的宽度和高度大于固定分辨率时，其有效显示区间将保持固定分辨率的大小(如图2-20所示)。
- 在Camera.aspect固定的情况下，无论选择Game视图中哪种屏幕模拟方式，它们的显示内容都是相同的(请查看实例演示)。不同的屏幕模拟方式只会对显示的内容进行放缩。决定屏幕视口显示内容的是Camera.aspect的值和Camera的Transform，至于屏幕要如何显示amera视口的内容，那就是硬件显示屏要处理的事情了。
- PixelRect和Rect功能类似,都是决定硬件显示屏如何显示camera视口提供的内容的。不同的是，PixelRect以实际像素来展示显示内容，而Rect以单位化形式来展示显示内容(请查看PixelRect和Rect的相关实例演示)。

```C#
using UnityEngine;
using System.Collections;
public class Compare_ts : MonoBehaviour
{
    Vector3 v1;
    void Start()
    {
        v1 = transform.position;
    }
    void OnGUI()
    {
        //设置Camera视口宽高比例为2∶1，点击此按钮后更改Game视图中不同的aspect值
        //尽管Scene视图中Camera视口的宽高比会跟着改变
        //但实际Camera视口的内容依然是按照2∶1的比例获取的
        //不同的屏幕显示相同的内容会发生变形
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "Camera视口宽高比例2∶1"))
        {
            camera.aspect = 2.0f;
            transform.position = v1;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "Camera视口宽高比例1∶2"))
        {
            camera.aspect = 0.5f;
            //更改Camera坐标，使被拉伸后的物体显示出来
            transform.position = new Vector3(v1.x, v1.y, 333.2f);
        }
        //恢复aspect的默认设置，即屏幕比例和Camera视口比例保持一致
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "使用Game面板中aspect的选择"))
        {
            camera.ResetAspect();
            transform.position = v1;
        }
    }
}
```

​	在这段代码中,用3个Button来设置3种不同的camera视口比例,在实例工程中设置Game视图的aspect值为16:10。如图2-21所示，按钮“使用Game面板中aspect的选择”的作用即为使视口的宽高比为Game视图中的设置，在这种状态下，场景中物体不会变形。

​	当设置Camera视口宽高比例为2:1时，相当于将camera的视口变宽(2:1>16:10)Camera的x轴方向上的视野将相对更大，要将更大的视野放到相同大小的屏幕上，物体自然会被压缩，如图2-22所示。同理，当设置camera视口宽高比例为1:2时，相当于将Camera的视口变窄(1:2<16:10)，要将更小的视野放到相同大小的屏幕上，物体自然会被拉伸，如图2-23所示。









## 3.GameObject类

## 4.HideFlags类

## 5.Mathf类

## 6.Matrix4X4类

## 7.Object类

## 8.Quaternaion类

## 9.Random类

## 10.RigidBody类

## 11.Time类

## 12.Transform类

## 13.Vector2类

## 14.Vector3类

## 15.游戏实例——坚守阵地

