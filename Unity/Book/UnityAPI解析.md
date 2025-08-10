typora-default-code-lang: csharp

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

GameObject类是Unity场景中所有实体的基类。一个GameObject对象通常由多个组件组成，且至少含有一个Transform组件。本章主要介绍GameObject类的一些实例属性、构造方法、实例方法和静态方法，并在最后对GameObject和Component这两个类之间的关系及其涉及的GetComponent相关方法的使用区别进行了注解。

### 3.1 GameObject类实例属性

在GameObject类中，涉及的实例属性有activeSelf和activeInHierarchy。由于这两个属性功能相似，因此将属性activeInHierarchy作为activeSelf属性的提示内容介绍，下面主要介绍activeSelf属性。

activeSelf属性：GameObject的Active标识

```
基本语法 public bool activeSelf { get; }
```

功能说明：此属性用来返回GameObject对象的Active标识状态。

---

提示：注意此属性与属性activeInHierarchy的区别。activeInHierarchy属性的功能是返回GameObject实例在程序运行时的激活状态，它只有当GameObject实例的状态被激活时才会返回true。而且它会受其父类对象激活状态的影响，如果其父类至最顶层的对象中有一个对象未被激活，activeInHierarchy就会返回false。

---

```C#
using UnityEngine;
using System.Collections;
public class ActiveSelf_ts : MonoBehaviour 
{
    public GameObject cube1, cube2, cube3;
    
    void Start () 
    {
        //对cube2设置为false，其他设置为true
        cube1.SetActive(true);
        cube2.SetActive(false);
        cube3.SetActive(true);
        Debug.Log("activeSelf:");
        //尽管cube2被设置为false，但其子类cube3的activeSelf返回值仍然为true
        Debug.Log("cube1.activeSelf:" + cube1.activeSelf);
        Debug.Log("cube2.activeSelf:" + cube2.activeSelf);
        Debug.Log("cube3.activeSelf:" + cube3.activeSelf);
        Debug.Log("\nactiveInHierarchy:");
        //cube2和cube3的activeInHierarchy返回值都为false
        Debug.Log("cube1.activeInHierarchy:" + cube1.activeInHierarchy);
        Debug.Log("cube2.activeInHierarchy:" + cube2.activeInHierarchy);
        Debug.Log("cube3.activeInHierarchy:" + cube3.activeInHierarchy);
    }
}
```

### 3.2 GameObject构造方法

activeSelf属性：GameObject的Active标识

```
基本语法 (1) public GameObject();
        (2) public GameObject(string name);
        	其中参数name为构造GameObject对象的名字。
        (3) public GameObject(string name, params Type[] components);
        	其中参数name为构造GameObject对象的名字，components为构造对象要添加的组件类型集合，多个组件之间用逗号隔开。
```

功能说明：创建一个GameObject对象。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Constructors_ts : MonoBehaviour 
{
    void Start () 
    {
        //使用构造函数GameObject (name : String)
        GameObject g1 = new GameObject("G1");
        g1.AddComponent<Rigidbody>();
        //使用构造函数GameObject ()
        GameObject g2 = new GameObject();
        g2.AddComponent<FixedJoint>();
        //使用构造函数GameObject (name : String, params components : Type[])
        GameObject g3 = new GameObject("G3",typeof(MeshRenderer),typeof(Rigidbody),typeof(SpringJoint));
        Debug.Log("g1 name:" + g1.name + "\nPosition:" + g1.transform.position);
        Debug.Log("g2 name:" + g2.name + "\nPosition:" + g2.transform.position);
        Debug.Log("g3 name:" + g3.name + "\nPosition:" + g3.transform.position);
    }
}
```

### 3.3 GameObject类实例方法

在GameObject类中，涉及的实例方法有GetComponent、GetComponentInChildren、GetComponents、GetComponentsInChildren 、SendMessage 、BroadcastMessage 和SendMessageUpwards 。

#### 3.3.1 GetComponent方法：获取组件

```
基本语法 (1) public T GetComponent<T>() where T : Component;
	    (2) public Component GetComponent(string type);
			其中参数type为组件名。
		(3) public Component GetComponent(Type type);
			其中参数type为组件类型。
```

功能说明：获取<span style="color:red;">第一个</span>符合Type类型的Component。

---

提示：与此方法功能相似的方法有GetComponentInChildren、GetComponents和GetComponents
InChildren，它们的具体使用请参考实例演示。需要注意以下两点。

- 在使用GetComponents (type : Type))方法时，`Component[] cjs= GetComponents(typeof(ConfigurableJoint)) as Component[];`

    注意不可以这样写：`ConfigurableJoint[] cjs = GetComponents(typeof(ConfigurableJoint)) as Configurable Joint[]`;因为ConfigurableJoint不是Component，而是其子类，建议使用其泛型方式。

- 在使用GetComponentsInChildren (type : Type, includeInactive : boolean = false)方法时，`Component[] cjs = GetComponentsInChildren(typeof(ConfigurableJoint), false) as Component[];`

    注意不可以这样写：`ConfigurableJoint[] cjs =GetComponentsInChildren(typeof(ConfigurableJoint), false) as ConfigurableJoint[];`因为ConfigurableJoint不是Component，而是其子类，建议使用其泛型方式。

---

代码：下面通过实例演示方法GetComponent、GetComponentInChildren、GetComponents和GetComponentsInChildren的使用。

```C#
using UnityEngine;
using System.Collections;
public class GetComponent_ts : MonoBehaviour 
{
	void Start () 
    {
        //以下是GetComponent方法的相关使用代码
        //1.GetComponent (type : Type) | GetComponent.<T>() | GetComponent (type : String)
        Rigidbody rb = GetComponent(typeof(Rigidbody)) as Rigidbody;
        rb = GetComponent<Rigidbody>();
        rb = GetComponent("Rigidbody") as Rigidbody;
        
        //以下是GetComponentInChildren方法的相关使用代码
		//GetComponentInChildren (type : Type) | GetComponentInChildren.<T>()
        rb = GetComponentInChildren(typeof(Rigidbody)) as Rigidbody;
        rb = GetComponentInChildren<Rigidbody>();
        
        //以下是GetComponents方法的相关使用代码
        //GetComponents (type : Type) | GetComponents.<T> ()
        Component[] cjs = GetComponents(typeof(ConfigurableJoint)) as Component[];      
        cjs = GetComponents<ConfigurableJoint>();
        
        //以下是GetComponentsInChildren方法的相关使用代码
        //GetComponentsInChildren(type: Type, includeInactive(包含未激活): boolean = false)  
        // |GetComponentsInChildren.<T> (includeInactive : boolean)
        // |GetComponentsInChildren.<T> ()
        cjs = GetComponentsInChildren(typeof(ConfigurableJoint), false) as Component[];
        cjs = GetComponentsInChildren(typeof(ConfigurableJoint), true) as Component[];      
        cjs = GetComponentsInChildren<ConfigurableJoint>(true);
		cjs = GetComponentsInChildren<ConfigurableJoint>();
	}
}
```

#### 3.3.2 SendMessage方法：发送消息

```
基本语法 (1) public void SendMessage(string methodName);
        (2) public void SendMessage(string methodName, object value);
        (3) public void SendMessage(string methodName, SendMessageOptions options);
        (4) public void SendMessage(string methodName, object value, SendMessageOptions options);
        参数methodName为接受信息的方法名字，参数value为信息的内容，参数options为信息接收的方式，默认为SendMessageOptions.RequireReceiver。
```

功能说明：向GameObject自身发送消息，其作用范围如下。

- 和自身同级的物体不会收到消息，例如cube1和cube2的上一级父类都是cube0，则cube2不会收到cube1发送的消息。
- SendMessageOptions有两个可选方式：SendMessageOptions.RequireReceiver和SendMessageOptions.DontRequireReceiver。前者要求信息的接收方必须有接受信息的方法，否则程序会报错，后者则无此要求。

---

提示：BroadcastMessage和SendMessageUpwards与之类似，如下对其说明。

- BroadcastMessage是向自身及其所有子类发送消息。和自身同级的物体不会收到消息，例如cube1和cube2的上一级父类都是cube0，则cube2不会收到cube1发送的消息。
- SendMessageUpwards是向GameObject 自身及其所有父类发送消息。和自身同级的物体不会收到消息，例如cube1和cube2 的上一级父类都是cube0，则cube2不会收到cube1发送的消息。

---

代码：在Cube1、Cube2和Cube3中分别绑定有脚本BroadcastMessage_ts.cs、SendMessage_ts.cs和SendMessageUpward_ts.cs此处以SendMessage_ts.cs脚本为例说明。

```C#
using UnityEngine;
using System.Collections;
public class SendMessage_ts : MonoBehaviour 
{
    void Start () 
	{
        //向子类及自己发送信息
        //gameObject.BroadcastMessage("GetParentMessage",gameObject.name+":use BroadcastMessage send!");
        //向自己发送信息
        gameObject.SendMessage("GetSelfMessage",gameObject.name+":use SendMessage send!");
        ////向父类及自己发送信息
        //gameObject.SendMessageUpwards("GetChildrenMessage",gameObject.name+":use SendMessageUpwards send!");
    }
    
    //一个接受父类发送信息的方法
    private void GetParentMessage(string str)
    {
    	Debug.Log(gameObject.name + "收到父类发送的消息：" + str);
    }
    
    //一个接受自身发送信息的方法
    private void GetSelfMessage(string str)
    {
    	Debug.Log(gameObject.name + "收到自身发送的消息：" + str);
    }
    
    //一个接受子类发送信息的方法
    private void GetChildrenMessage(string str)
    {
    	Debug.Log(gameObject.name + "收到子类发送的消息：" + str);
    }
}
```

### 3.4 GameObject 类静态方法

静态方法主要有CreatePrimitive。在使用CreatePrimitive创建GameObject对象时，会涉及添加组件（AddComponent）和查找对象（Find）。

#### **3.4.1 CreatePrimitive：创建GameObject对象**

```
基本语法 public static GameObject CreatePrimitive(PrimitiveType type);
		其中参数type为PrimitiveType的类型值。
```

功能说明：创建GameObject对象。

代码：CreatePrimitive方法以及Find方法的使用。

```C#
using UnityEngine;
using System.Collections;
public class CreatePrimitive_ts : MonoBehaviour
{
    void Start()
    {
        //使用GameObject.CreatePrimitive方法创建GameObject
        GameObject g1 = GameObject.CreatePrimitive(PrimitiveType.Sphere);
        g1.name = "G1";
        g1.tag = "sphere_Tag";
        
        //使用 AddComponent (className : String)方法添加组件
        g1.AddComponent("SpringJoint");
        
        //使用AddComponent (componentType : Type)方法添加组件
        g1.AddComponent(typeof(GUITexture));
        
        g1.transform.position = Vector3.zero;
        
        GameObject g2 = GameObject.CreatePrimitive(PrimitiveType.Sphere);
        g2.name = "G2";
        g2.tag = "sphere_Tag";
        
        //使用AddComponent.<T>()方法添加组件
        g2.AddComponent<Rigidbody>();
        
        g2.transform.position = 2.0f * Vector3.right;
        g2.GetComponent<Rigidbody>().useGravity = false;
        
        GameObject g3 = GameObject.CreatePrimitive(PrimitiveType.Sphere);
        g3.name = "G1";
        g3.tag = "sphere_Tag";
        g3.transform.position = 4.0f * Vector3.right;
        
        //使用GameObject.Find类方法获取GameObject，返回符合条件的第一个对象
        Debug.Log("use Find:" + GameObject.Find("G1").transform.position);
        
        //使用GameObject.FindGameObjectWithTag类方法获取GameObject，返回符合条件的第一个对象
        Debug.Log("use FindGameObjectWithTag:" + GameObject.FindGameObjectWithTag("sphere_
        Tag").transform.position);
                                                                                  
        //使用GameObject.FindGameObjectsWithTag类方法获取GameObject，返回符合条件的所有对象
        GameObject[] gos = GameObject.FindGameObjectsWithTag("sphere_Tag");
        foreach (GameObject go in gos)
        {
            Debug.Log("use FindGameObjectsWithTag:" + go.name + ":" + go.transform.position);
        }
        //更改g1、g2和g3的层级关系
        g3.transform.parent = g2.transform;
        g2.transform.parent = g1.transform;
        Debug.Log("use Find again1:" + GameObject.Find("G1").transform.position);
        //使用带"/"限定条件的方式查找GameObject
        //此处返回的对象需其父类为G2，且G2的父类名为G1
        //注意与上面不带"/"限定条件返回的对象的区别
        Debug.Log("use Find again2:" + GameObject.Find("/G1/G2/G1").transform.position);
    }
}
```

### 3.5 关于GameObject类和Component类的使用注解

GameObject和Component是Unity中常用且非常重要的两个类，二者的实例属性和实例方法相似，
只是在使用方法上稍有区别，所以本书不再对Component类作单独介绍，下面对这两个类之间的
关系及其实例方法的使用进行简要说明。

- 通常一个GameObject对象由多个Component组成，而且至少有一个Transform组件。GameObject用来管理工程中的各个物体，而Component用来扩展这些物体自身的功能。
- GameObject类和Component类的属性名称和方法名称基本相同，各个属性和方法的用法也很相近，但它们仍有一些差别，以下以GetComponent方法为例说明。

若要获取当前脚本所在GameObject对象中的某个组件，直接使用GetComponent方法即可，如
`Rigidbody rb = GetComponent<Rigidbody>()`。

代码:

```C#
using UnityEngine;
using System.Collections;
public class GameObjectAndComponent_ts : MonoBehaviour {
    public GameObject sp;
    void Start () {
        //以下3种表达方式功能一样，都返回当前脚本所在GameObject对象中的Rigidbody组件
        Rigidbody rb1 = rigidbody;	//这是个什么写法？内置？
        Rigidbody rb2 = GetComponent<Rigidbody>();
        Rigidbody rb3 = rigidbody.GetComponent<Rigidbody>();
        Debug.Log("rb1的InstanceID：" + rb1.GetInstanceID());
        Debug.Log("rb2的InstanceID：" + rb2.GetInstanceID());
        Debug.Log("rb3的InstanceID：" + rb3.GetInstanceID());
        //使用前置引用获取引用对象的Rigidbody组件
        Debug.Log("前置引用sp对象中Rigidbody的InstanceID："+sp.GetComponent<Rigidbody>().GetInstanceID());
    }
}
```

在这段代码中，首先声明了一个GameObject变量sp，用来指向外部GameObject对象，然后在Start方法中用3种不同的方法来获取当前脚本所在GameObject对象中的Rigidbody组件，并打印出它们的InstanceID，如图3-11所示。最后演示了使用GameObject前置引用来获取外部对象的Rigidbody组件的方法。由打印结果可知，虽然rb1、rb2和rb3的获取方式不一样，但都指向了相同的对象。

## 4.HideFlags类

HideFlags为枚举类，用于控制Object对象的销毁方式及其在检视面板中的可视性。本章将对HideFlags类枚举成员的功能及使用方法进行较为详细的说明。

### 4.1 HideFlags 类枚举成员

枚举类HideFlags涉及的枚举成员有DontSave、HideAndDontSave、HideInHierarchy、HideInInspector、None和NotEditable，下面详细介绍这些枚举成员。

#### 4.1.1 DontSave：保留对象到新场景

功能说明：设置是否将Object对象保留到新的场景中，如果使用HideFlags.DontSave，则Object对象将在新场景中被保留下来。

- 如果GameObject对象被HideFlags.DontSave标识，则在新Scene中GameObject的所有组件将被保留下来，但其子类GameObject对象不会被保留到新Scene中。
- 不可以对GameObject对象的某个组件如Transform进行HideFlags.DontSave标识，否则无效。
- 即使程序已经退出，被HideFlags.DontSave标识的对象也会一直存在于程序中，造成内存泄漏，对HideFlags.DontSave标识的对象，在不需要或程序退出时需要使用DestroyImmediate手动销毁。

代码：演示属性DontSave的使用。本实例工程包含两个场景，下面是场景DontSave_unity中的脚本代码：

```C#
using UnityEngine;
using System.Collections;
public class DontSave_ts : MonoBehaviour 
{
    public GameObject go;
    public Transform t;
    void Start()
    {
        //GameObject对象使用HideFlags.DontSave可以在新scene中被保留
        go.hideFlags = HideFlags.DontSave;
        GameObject P1 = GameObject.CreatePrimitive(PrimitiveType.Plane);
        P1.hideFlags = HideFlags.DontSave;
        //不可以对GameObject的组件设置HideFlags.DontSave，否则无效
        Transform tf = Instantiate(t, go.transform.position + new Vector3(2.0f, 0.0f, 0.0f),Quaternion.identity) as Transform;
        tf.hideFlags = HideFlags.DontSave;
        //导入名为newScene_unity的新scene
        Application.LoadLevel("newScene2_unity");
    }
}
```

在这段代码中，分别对场景中GameObject对象go、新创建的GameObject对象P1和新实例化的Transform实例tf的hideFlags属性设置为HideFlags.DontSave，然后导入名为newScene2_unity的新场景。

代码newScene2_unity中：

```C#
using UnityEngine;
using System.Collections;
public class NewScene2_ts : MonoBehaviour {
    GameObject cube, plane;
    void Start () 
    {
        Debug.Log("这是NewScene2！");
    }
    
    //当程序退出时用DestroyImmediate()销毁被HideFlags.DontSave标识的对象
    //否则即使程序已经退出，被HideFlags.DontSave标识的对象依然在Hierarchy面板中
    //即每运行一次程序就会产生多余对象，造成内存泄漏
    void OnApplicationQuit()
    {
        cube = GameObject.Find("Cube0");
        plane = GameObject.Find("Plane");
        if (cube)
        {
            Debug.Log("Cube0 DestroyImmediate");
            DestroyImmediate(cube);
        }
        if (plane)
        {
            Debug.Log("Plane DestroyImmediate");
            DestroyImmediate(plane);
        }
    }
}
```

在OnApplicationQuit()方法中查找当前场景中是否存在名为cube和plane的对象，如果存在，则在程序退出时将它们立即销毁。

#### 4.1.2 HideAndDontSave：保留对象到新场景

功能说明：设置是否将Object对象保留到新Scene中，如果使用HideFlags.HideAndDontSave，则Object对象将在新Scene中被保留下来，但不会显示在Hierarchy面板中。

#### 4.1.3 HideInHierarchy：在Hierarchy面板中隐藏

功能说明：设置Object对象在Hierarchy面板中是否被隐藏。

- 若要在Hierarchy面板中隐藏不是在脚本中创建的对象，需要在Awake方法中使用HideFlags.HideInHierarchy才能生效。
- 若隐藏父物体，子物体也会被隐藏掉，但隐藏子物体，父物体不会被影响。

代码:

```C#
using UnityEngine;
using System.Collections;
public class HideInHierarchy_ts : MonoBehaviour
{
    public GameObject go, sub;
    public Transform t;
    void Awake()
    {
        //go、sub、gameObject为已存在对象，需在Awake方法中使用HideFlags.HideInHierarchy
        go.hideFlags = HideFlags.HideInHierarchy;
        sub.hideFlags = HideFlags.HideInHierarchy;
        gameObject.hideFlags = HideFlags.HideInHierarchy;
    }
    
    void Start()
    {
        //P1、tf是在代码中创建的对象，可以在任何方法中使用HideFlags.HideInHierarchy
        GameObject P1 = GameObject.CreatePrimitive(PrimitiveType.Plane);
        P1.hideFlags = HideFlags.HideInHierarchy;
        Transform tf = Instantiate(t, go.transform.position + new Vector3(2, 0.0f, 0.0f),Quaternion.identity) as Transform;
        tf.hideFlags = HideFlags.HideInHierarchy;
    }
}
```

在Awake方法中设置hideFlags属性为HideFlags.HideInHierarchy。然后在Start方法中临时创建和实例化了两个GameObject对象P1和tf，并对它们的hideFlags进行设置。从运行结果可以发现，原来的cube对象包括其子类sphere、cube、MainCamera及新创建的两个对象P1和tf都被隐藏了。

#### 4.1.4 HideInInspector：在Inspector面板中隐藏

功能说明：设置Object对象在Inspector面板中是否被隐藏。

- 如果一个GameObject 使用了HideFlags.HideInInspector ， 则其所有组件将在Inspector面板中被隐藏，但其子类对象的组件仍可在Inspector面板中显示。
- 如果只隐藏了GameObject对象的某个组件，如Transform，则并不影响GameObject的其他组件如Renderer、Rigidbody等在Inspector面板中的显示状态。

代码:

```C#
using UnityEngine;
using System.Collections;
public class HideInInspector_td : MonoBehaviour 
{
    public GameObject go;
    public Transform t;
    void Start()
    {
        //go、gameObject、Pl都是GameObject对象，使用HideFlags.HideInInspector后
        //其所有组件将在Inspector面板中隐藏
        //但并不影响其子类在Inspector面板中的显示
        go.hideFlags = HideFlags.HideInInspector;
        gameObject.hideFlags = HideFlags.HideInInspector;
        GameObject P1 = GameObject.CreatePrimitive(PrimitiveType.Plane);
        P1.hideFlags = HideFlags.HideInInspector;
        //tf为Transform对象，使用HideFlags.HideInInspector后
        //tf对应的GameObject的Transform组件将在Inspector面板中隐藏
        //但GameObject的其他组件仍可在Inspector面板中显示
        Transform tf = Instantiate(t, go.transform.position + new Vector3(2.0f, 0.0f, 0.0f),Quaternion.identity) as Transform;
        tf.hideFlags = HideFlags.HideInInspector;
    }
}
```

在Start方法中对go及MainCamera的hideFlags设置为HideFlags.HideInInspector，接着分别创建和实例化了一个GameObject对象P1和一个Transform实例tf，并将它们的hideFlags都设置为HideFlags.HideInInspector。运行程序可以发现，对象go、MainCamera及P1的Inspector面板中的组件都被隐藏，tf组件也被隐藏，但其所在GameObject对象的其他组件却未被隐藏。

#### 4.1.5 None：HideFlags默认值

功能说明：None为HideFlags的默认值，不改变Object对象的可见性。

#### 4.1.6 NotEditable：对象在Inspector面板中的可编辑性

功能说明：在程序运行时Object对象是否可在Inspector面板中被编辑。

- GameObject对象使用HideFlags.NotEditable可以使得GameObject对象的所有组件在Inspector面板中都处于不可编辑状态。但GameObject对象被HideFlags.NotEditable标识并不影响其子类对象的可编辑性。
- 对于GameObject对象的某个组件如Transform单独使用HideFlags.NotEditable，只会使得当前组件不可编辑，但GameObject的其他组件仍可在Inspector面板中编辑。

代码：

```C#
using UnityEngine;
using System.Collections;
public class NotEditable_ts : MonoBehaviour 
{
    public GameObject go;
    public Transform t;
    void Start()
    {
        //GameObject对象使用HideFlags.NotEditable可以使得GameObject的
        //所有组件在Inspector面板中都处于不可编辑状态
        //GameObject对象被HideFlags.NotEditable标识并不影响其子类的可编辑性
        go.hideFlags = HideFlags.NotEditable;
        GameObject P1 = GameObject.CreatePrimitive(PrimitiveType.Plane);
        P1.hideFlags = HideFlags.NotEditable;
        //对于GameObject的某个组件单独使用HideFlags.NotEditable
        //只会使得当前组件不可编辑，但GameObject的其他组件仍可编辑
        t.hideFlags = HideFlags.NotEditable;
        Transform tf = Instantiate(t, go.transform.position + new Vector3(2.0f, 0.0f, 0.0f),Quaternion.identity) as Transform;
        tf.hideFlags = HideFlags.NotEditable;
    }
}
```

在Start方法中对go、新创建的GameObject对象P1及新实例化的Transform实例tf的hideFlags属性设置为HideFlags. NotEditable。对象go和P1的所有组件将不可被编辑，组件tf也不可被编辑，但组件tf所在的GameObject对象的其他组件仍可被编辑。

### 4.2 HideFlags 类使用小结

本章对HideFlags类的枚举成员进行了介绍。在HideFlags类的6个枚举成员中，用得较多的是DontSave，使用它将Object对象保留到新Scene中时，需要注意在合适的时机将Object对象手动销毁。若想将场景中的某个Object对象在Hierarchy面板或Inspector面板中隐藏，可以使用成员HideInHierarchy或HideInInspector，但要注意对象间的层次问题。

## 5.Mathf类

Mathf是Unity中的数学类，只有静态属性和静态方法。在使用时，直接调用其静态属性或静态方法，如Mathf.PI、Mathf.Sin等。

### 5.1 Mathf 类静态属性

在Mathf类中，涉及的静态属性有Deg2Rad、Rad2Deg和Infinity。

#### 5.1.1 Deg2Rad属性：从角度到弧度常量

```
基本语法 public const float Deg2Rad = 0.0174533f;
```

功能说明：表示数学计算中从角度到弧度转变的常量值，其值为`(2 * Mathf.PI) / 360=0.01745329`，此属性只读。

---

提示：Rad2Deg属性与此属性功能相反，是从弧度到角度的转换常量，其值为57.2958f。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class DegAndRad_ts : MonoBehaviour 
{
    void Start () 
    {
        //从角度到弧度转换常量
        Debug.Log("Mathf.Deg2Rad:" + Mathf.Deg2Rad);
        //从弧度到角度转换常量
        Debug.Log("Mathf.Rad2Deg:" + Mathf.Rad2Deg);
    }
}
```

#### 5.1.2 Infinity属性：正无穷大

```
基本语法 public const float Infinity = 1.0f / 0.0f;
```

功能说明：表示数学中的正无穷大，只读。

- `Mathf.Infinity ÷ x = Mathf.Infinity`，其中x为一个具体数值，如10000。
- `Mathf.Infinity ÷ Mathf.Infinity = NaN`，计算结果不是数值（Not a Number）。
- Mathf.Infinity是正无穷大，不要用在计算中。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Infinity_ts : MonoBehaviour
{
    void Start()
    {
        Debug.Log("0:" + Mathf.Infinity);				   //Infinity
        Debug.Log("1:" + Mathf.Infinity / 10000.0f);	   //Infinity
        Debug.Log("2:" + Mathf.Infinity / Mathf.Infinity); //NaN
    }
}
```

### 5.2 Mathf 类静态方法

静态方法有`Clamp`、`ClosestPowerOfTwo`、`DeltaAngle`、`InverseLerp`、`Lerp`、`LerpAngle`、`MoveTowards`、`MoveTowardsAngle`、`PingPong`、`Repeat`、`Round`、`SmoothDamp`、`SmoothDampAngle`和`SmoothStep`。

#### 5.2.1 Clamp：返回有限范围值

```
基本语法 (1) public static float Clamp(float value, float min, float max);
        (2) public static int Clamp(int value, int min, int max);
        	其中min最小值，max最大值。整型/浮点型。
```

功能说明：限制value的值在[min, max]之内。

---

提示：Clamp01只需要一个参数，限制value取值范围为[0,1]之内。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class Clamp_ts : MonoBehaviour
{
    void Start()
    {
        Debug.Log("当value<min时：" + Mathf.Clamp(-1, 0, 10));
        Debug.Log("当min<value<max时：" + Mathf.Clamp(3, 0, 10));
        Debug.Log("当value>max时：" + Mathf.Clamp(11, 0, 10));
        //方法Clamp01的返回值范围为[0,1]
        Debug.Log("当value<0时:" + Mathf.Clamp01(-0.1f));
        Debug.Log("当0<value<1时:" + Mathf.Clamp01(0.5f));
        Debug.Log("当value>1时:" + Mathf.Clamp01(1.1f));
    }
}
```

#### 5.2.2 ClosestPowerOfTwo：返回2 的某次幂

```
基本语法 public static int ClosestPowerOfTwo(int value);
```

功能说明：返回最接近value的2的某次幂的值。中间向上取值：

```C#
f = Mathf.ClosestPowerOfTwo(11)		//f=8；
f = Mathf.ClosestPowerOfTwo(12)		//f=16；
当value值小于0时，返回值为0。
```

代码：

```C#
using UnityEngine;
using System.Collections;
public class ClosestPowerOfTwo_ts : MonoBehaviour
{
    void Start()
    {
        Debug.Log("11与8最接近，输出值为：" + Mathf.ClosestPowerOfTwo(11));
        Debug.Log("12与8和16的差值都为4，输出值为：" + Mathf.ClosestPowerOfTwo(12));
        Debug.Log("当value<0时，输出值为：" + Mathf.ClosestPowerOfTwo(-1));
    }
}
```

#### 5.2.3 DeltaAngle：最小增量角度

```
基本语法 public static float DeltaAngle(float current, float target);
		其中参数current为当前角度，参数target为目标角度。
```

功能说明：返回从current到target的最小增量角度值。计算方法：将current和target按照360度为一周换算到区间（-180,180]中，设current和target换算后的值分别对应c和t，<span style="color:red;">在坐标轴中的夹角为e（0≤e≤180），则若c经过顺时针旋转e度能到达t，则返回值为e；若c经过逆时针旋转e度能到达t，则返回值为-e。</span>

![](https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Book/DeltaAngle%E6%8D%A2%E7%AE%97%E5%9B%BE.png)

代码：

```C#
using UnityEngine;
using System.Collections;
public class DeltaAngle_ts : MonoBehaviour 
{
    void Start () 
    {
        Debug.Log(Mathf.DeltaAngle(1180, 90));
        //100=1180-360*3,即求100和90之间的夹角		//-10
        Debug.Log(Mathf.DeltaAngle(-1130, 90));
        //-50=-1130+360*3,即求-50和90之间的夹角		//140
        Debug.Log(Mathf.DeltaAngle(-1200, 90));
        //-120=-1200+360*3,即求-120和90之间的夹角	//150
    }
}
```

#### 5.2.4 InverseLerp：计算比例值

```
基本语法 public static float InverseLerp(float from, float to, float value);
		其中参数from为起始值，参数to为终点值，参数value为参考值。
```

功能说明：返回value值在从参数from到to中的比例值。设`f = Mathf.InverseLerp(from,to,value)`，$\large{f' = \frac{value-from}{to-from}}$，则若f'∈[0.0f,1.0f]则f = f '；若f ' < 0，则f = 0；若f ' > 1则f = 1。

代码：

```C#
using UnityEngine;
using System.Collections;
public class InverseLerp_ts : MonoBehaviour 
{
    void Start () 
    {
        // Unity 中的实现t = 0.5 (15 在 10~20 的中间)
		float t = Mathf.InverseLerp(10, 20, 15);
       	//***
        float f,from,to,v;
        from = -10.0f;
        to = 30.0f;
        v = 10.0f;
        f = Mathf.InverseLerp(from,to,v);
        Debug.Log("当0<f'<1时："+f);
        v = -20.0f;
        f = Mathf.InverseLerp(from, to, v);
        Debug.Log("当f'<0时：" + f);
        v = 40.0f;
        f = Mathf.InverseLerp(from, to, v);
        Debug.Log("当f'>1时：" + f);
        ***//
    }
}
```

#### 5.2.5 Lerp：线性插值

```
基本语法 public static float Lerp(float from, float to, float t);
		from为起始值，to为结束值，t为插值系数。
```

功能说明：返回一个从from到to的线性插值。返回值的计算方法为(to-from)*t'+from

- t的取值范围为[0,1]，当t<0时有效值t'=0，当t>1时有效值t'=1；

代码：

```C#
using UnityEngine;
using System.Collections;
public class Lerp_ts : MonoBehaviour 
{
    float r, g, b;
    void FixedUpdate () 
    {
        r = Mathf.Lerp(0.0f,1.0f,Time.time*0.2f);
        g = Mathf.Lerp(0.0f, 1.0f, -1.0f + Time.time * 0.2f);
        b = Mathf.Lerp(0.0f, 1.0f, -2.0f + Time.time * 0.2f);
        light.color = new Color(r, g, b);
    }
}
```

变量r、g和b记录Color的RGB值。在FixedUpdate方法中使用Lerp使r、g、b随着时间依次递增。运
行程序，light会由黑变红接着变黄最后变成白色。

#### 5.2.6 LerpAngle：角度插值

```
基本语法 public static float LerpAngle(float a, float b, float t);
		a为起始角度，b为结束角度，参数t为插值系数。
```

功能说明：返回从a到b之间的角度插值。

- 插值系数t的取值范围为[0,1]，当t<0时其有效值t'=0，当t>1时其有效值t'=1。
- 插值计算之前需要先对a、b进行规范化，以确定需要插值的大小，对a、b规范化规
    则如下（以参数a为例，b与此相同）：

$$
a' = 360k + a\text{,  其中k∈Z,求k使得a'∈[0, 360]}
$$

对a、b规范化为a'、b'。设a'和b'之间的差值为c，并且c∈[0，180]。

当a'沿着顺时针方向旋转c度与b'重合时，则插值计算方式为：`f=a-c*t',t'为t的有效值`；
当a'沿着逆时针方向旋转c度与b'重合时，则插值计算方式为：`f=a+c*t',t'为t的有效值`；

```C#
using UnityEngine;
using System.Collections;
public class LerpAngle_ts : MonoBehaviour 
{
    void Start () 
    {
        float a, b;
        a = -50.0f;	  //a'=360-50=310
        b = 400.0f;	  //b'=-360+400=40
        Debug.Log("test1:"+Mathf.LerpAngle(a,b,0.3f));
        //从a'到b'可以逆时针旋转90°，故返回值test1 = a+c*t = -50+90*0.3 = -23
        
        a = 400.0f;	   //a'=-360+400=40
        b = -50.0f;	   //b'=360-50=310
        Debug.Log("test2:" + Mathf.LerpAngle(a, b, 0.3f));
        //从a'到b'可以顺时针旋转90°，故返回值test2 = a-c*t= 400-90*0.3 = 373
    }
}
```

#### 5.2.7 MoveTowards：选择性插值

```
基本语法 public static float MoveTowards(float current, float target, float maxDelta);
		current为当前值，target为目标值，maxDelta为步长。
```

功能说明：返回一个从current到target之间的插值。

代码：

```C#
using UnityEngine;
using System.Collections;
public class MoveTowards_ts : MonoBehaviour
{
    void Start()
    {
        float current, target, maxDelta;
        current = 10.0f;
        target = -10.0f;
        maxDelta = 5.0f;
        Debug.Log(Mathf.MoveTowards(current, target, maxDelta)); 
        //5 = 10 - 5 
        
        maxDelta = 50.0f;
        Debug.Log(Mathf.MoveTowards(current, target, maxDelta));
        //-10
        
        current = 10.0f;
        target = 50.0f;
        maxDelta = 5.0f;
        Debug.Log(Mathf.MoveTowards(current, target, maxDelta));
        //15 = 10 + 5
        
        maxDelta = 50.0f;
        Debug.Log(Mathf.MoveTowards(current, target, maxDelta));
        //50
    }
}
```

#### 5.2.8 MoveTowardsAngle：角度的选择性插值

```
基本语法 public static float MoveTowardsAngle(float current, float target, float maxDelta);
		curren为当前角度，target为目标角度，maxDelta为步长。
```

功能说明：返回一个从当前角度current向目标角度target旋转的插值。

代码：

```C#
using UnityEngine;
using System.Collections;
public class MoveTowardsAngle_ts : MonoBehaviour
{
    float targets = 0.0f;
    float speed = 40.0f;
    void Update()
    {
        //每帧不超过speed * Time.deltaTime度
        float angle = Mathf.MoveTowardsAngle(transform.eulerAngles.y, targets, speed * Time.deltaTime);
        transform.eulerAngles = new Vector3(0, angle, 0);
    }
    
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "顺时针旋转90度"))
        {
            targets += 90.0f;
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "逆时针旋转90度"))
        {
            targets -= 90.0f;
        }
    }
}
```

targets用于记录物体旋转的目标角度，可在OnGUI方法中设置，变量speed用于控制物体每帧旋转的最大角度。在Update方法中调用MoveTowardsAngle，返回一个从物体当前角度到目标角度的一个插值。将这个插值赋给transform的eulerAngles。

#### 5.2.9 PingPong：往复运动

```
基本语法 public static float PingPong(float t, float length);
```

功能说明：让数值在 0 和 length之间来回往返，形成循环运动。
$$
\text{PingPong}(t, l) = l - \left| (t\%2l) - l \right|
\text{其中t为时间参数，l为运动区间长度}
$$
代码：

```C#
using UnityEngine;
using System.Collections;
public class PingPong_ts : MonoBehaviour 
{
    void Start () 
    {
        float f, t, l;
        t = 11.0f;
        l = 5.0f;
        f = Mathf.PingPong(t,l);	//1 = 5 - |(11 % (2*5) - 5)|
        t = 17.0f;
        l = 5.0f;
        f = Mathf.PingPong(t, l);	//3 = 5 - |(17 % (2*5)) - 5|
    }
}
```

#### 5.2.10 Repeat：取模运算

```
基本语法 public static float Repeat(float t, float length);
```

功能说明：浮点数的取模运算。
$$
\text{Repeat}(t, l) = t \: \% \: l
$$
代码：

```C#
using UnityEngine;
using System.Collections;
public class Repeat_ts : MonoBehaviour 
{
    void Start () 
    {
        float f, t, l;
        t = 12.5f;
        l = 5.3f;
        f = Mathf.Repeat(t,l); // 1.9 = 12.5 - 5.3 * 2
        t = -12.5f;
        l = -5.3f;
        f = Mathf.Repeat(t, l);// -1.9 = -12.5f - (-5.3) * 2
        //特殊
        t = -12.5f;
        l = 0.0f;
        f = Mathf.Repeat(t, l);// Nan
    }
}
```

#### 5.2.11 Round：浮点数的整型值

```
基本语法 public static float Round(float f);
```

功能说明：返回离f最近的整型浮点值。设a为整数部分和b为小数部分即f=a+b计算规则（银行家舍入法（四舍六入五成双））。

- |小数部分|< 0.5，Round(f)返回整数部分；
- |小数部分|> 0.5，若f为正数返回整数部分+1；若f为负数返回整数部分-1。（上下取整）
- |小数部分|= 0.5，若整数部分为偶数，返回值为整数部分；若整数部分为奇数，如果f是负数返回整数部分-1，如果f是正数返回整数部分+1。

$$
\text{Round}(x) = 
\begin{cases} 
\lfloor x \rfloor      & \text{if } x - \lfloor x \rfloor < 0.5, \\
\lceil x \rceil        & \text{if } x - \lfloor x \rfloor > 0.5, \\
\lfloor x \rfloor      & \text{if } x - \lfloor x \rfloor = 0.5 \text{ 且 } \lfloor x \rfloor \text{ 为偶数}, \\
\lfloor x \rfloor + 1  & \text{if } x - \lfloor x \rfloor = 0.5 \text{ 且 } \lfloor x \rfloor \text{ 为奇数}.
\end{cases}
$$

---

提示：RoundToInt类似，返回整型值。

tips:银行家舍入法可以减少统计偏差通过「五成双」规则（向偶数取整），使舍入误差在统计上更均衡。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class Round_ts : MonoBehaviour
{
    void Start()
    {
        //设Round(f)中f=a.b
        Debug.Log("b<0.5,f>0：" + Mathf.Round(2.49f));	//2
        Debug.Log("b<0.5,f<0：" + Mathf.Round(-2.49f));	//-2
        Debug.Log("b>0.5,f>0：" + Mathf.Round(2.61f));	//3
        Debug.Log("b>0.5,f<0：" + Mathf.Round(-2.61f));	//-3
        Debug.Log("b=0.5,a为偶数,f>0：" + Mathf.Round(6.5f));	//6
        Debug.Log("b=0.5,a为偶数,f<0：" + Mathf.Round(-6.5f));	//-6
        Debug.Log("b=0.5,a为奇数,f>0：" + Mathf.Round(7.5f));	//8
        Debug.Log("b=0.5,a为奇数,f<0：" + Mathf.Round(-7.5f));	//-8
    }
}
```

#### 5.2.12 SmoothDamp：模拟阻尼运动

```
基本语法 (1) public static float SmoothDamp(float current, float target, ref float currentVelocity, float smoothTime);
        (2) public static float SmoothDamp(float current, float target, ref float currentVelocity, float smoothTime, float maxSpeed);
        (3) public static float SmoothDamp(float current, float target, ref float currentVelocity, float smoothTime, float maxSpeed, float deltaTime);
        current为起始值；target为目标值；currentVelocity为当前帧速度；smoothTime为预计平滑时间；maxSpeed为当前帧最大速度值，默认值为Mathf.Infinity；deltaTime为平滑时间，值越大返回值也相对越大，一般用Time.deltaTime计算。
```

功能说明：模拟平滑阻尼运动，并返回模拟插值。smoothTime预计平滑时间，物体越靠近目标，加速度的绝对值越小。实际到达目标的时间往往要比预计时间大很多，smoothTime∈(0.0f,1.0f)。

代码：

```C#
using UnityEngine;
using System.Collections;
public class SmoothDamp_ts : MonoBehaviour
{
    float targets = 110.0f;//目标值
    float cv1 = 0.0f, cv2 = 0.0f; //输出值
    float maxSpeeds = 50.0f;//每帧最大值
    float f1 = 10.0f, f2 = 10.0f;//起始值
    void FixedUpdate()
    {
        //maxSpeed取默认值
        f1 = Mathf.SmoothDamp(f1, targets, ref cv1, 0.5f);
        Debug.Log("f1:" + f1);
        Debug.Log("cv1:" + cv1);
        //maxSpeed取有限值50.0f
        f2 = Mathf.SmoothDamp(f2, targets, ref cv2, 0.5f, maxSpeeds);
        Debug.Log("f2:" + f2);
        Debug.Log("cv2:" + cv2);
    }
}
```

在FixedUpdate中调用了两次SmoothDamp方法，第一次调用时取maxSpeed值为默认值，即无穷大，第二次调用时取maxSpeed值为有限值。图中分别是对输出值cv1和cv2随时间变化的可视化显示，起始时输出速度提升很快，结束时速度下降却很平缓，在移动距离相同的情况下，有最大速度限制的花费时间也更多。

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Book/SmoothDamp%E6%A8%A1%E6%8B%9F%E9%98%BB%E5%B0%BC%E5%9B%BE.png" style="zoom:80%;" />

#### 5.2.13 SmoothDampAngle：阻尼旋转

```
基本语法 (1) public static float SmoothDampAngle(float current, float target, ref float
currentVelocity, float smoothTime);
		(2) public static float SmoothDampAngle(float current, float target, ref float
currentVelocity, float smoothTime, float maxSpeed);
		(3) public static float SmoothDampAngle(float current, float target, ref float
currentVelocity, float smoothTime, float maxSpeed, float deltaTime);
        current为起始值；target为目标值；currentVelocity为当前帧速度；smoothTime为预计平滑时间；maxSpeed为当前帧最大速度值，默认值为Mathf.Infinity；参数deltaTime为平滑时间。
```

功能说明：模拟角度的平滑阻尼旋转，并返回模拟插值。

代码：

```C#
using UnityEngine;
using System.Collections;
public class SmoothDampAngle_ts : MonoBehaviour
{
    public Transform targets;
    float smoothTime = 0.3f;
    float distance = 5.0f;
    float yVelocity = 0.0f;
    void Update()
    {
        //返回平滑阻尼角度值
        float yAngle = Mathf.SmoothDampAngle(transform.eulerAngles.y, targets.eulerAngles.y,ref yVelocity, smoothTime);
        Vector3 positions = targets.position;
        //由于使用transform.LookAt，此处计算targets的-z轴方向距离targets为distance
        //欧拉角为摄像机绕target的y轴旋转yAngle的坐标位置
        positions += Quaternion.Euler(0, yAngle, 0) * new Vector3(0, 0, -distance);
        //向上偏移2个单位
        transform.position = positions + new Vector3(0.0f, 2.0f, 0.0f);
        transform.LookAt(targets);
    }
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "将targets旋转60度"))
        {
            //更改targets的eulerAngles
            targets.eulerAngles += new Vector3(0.0f, 60.0f, 0.0f);
        }
    }
}
```

#### 5.2.14 SmoothStep：平滑插值

```
基本语法 public static float SmoothStep(float from, float to, float t);
		from为起始值，to为结束值，t为插值系数。
```

功能说明：返回一个从from到to的平滑插值。t∈[0.0f,1.0f]

代码：

```C#
using UnityEngine;
using System.Collections;
public class SmoothStep_ts : MonoBehaviour 
{
    float min = 10.0f;
    float max = 110.0f;
    float f1, f2 = 0.0f;
    void FixedUpdate () 
    {
        //f1为SmoothStep插值返回值
        f1 = Mathf.SmoothStep(min,max,Time.time);
        //计算相邻两帧插值的变化
        f2 = f1 - f2;
        Debug.Log("f1:" + f1);
        Debug.Log("f2:" + f2);
        f2 = f1;
    }
}
```

在FixedUpdate方法中调用方法SmoothStep，并将返回值赋给f1，接着计算与前一帧插值的差，并分别打印出f1和f2的值。

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Book/SmoothStep%E5%B9%B3%E6%BB%91%E6%8F%92%E5%80%BC%E5%9B%BE.png" style="zoom: 80%;" />

## 6.Matrix4X4类

在Unity常用Vector3、Quaternion、Transform等来对物体进行变换，Matrix4x4类通常用在一些比较特殊的地方，如对摄像机的非标准投影变换等。本章主要介绍了Matrix4x4类的一些实例方法和静态方法。

### 6.1 Matrix4x4 类实例方法

在Matrix4x4类中实例方法有`MultiplyPoint`、`MultiplyPoint3x4`、`MultiplyVector`和`SetTRS`。

#### 6.1.1 MultiplyPoint：投影矩阵变换

```
基本语法 public Vector3 MultiplyPoint(Vector3 v);
```

功能说明：对点v进行投影矩阵变换。例如，设m1为Matrix4x4实例，v1为Vector3实例，Vector3 v2=m1. MultiplyPoint(v1)，则v2值的变换过程如下：v2=v1·m1·M，系统在进行变换时会给v1增加一个w的分量，扩充为四维向量，w默认值为1，而M为投影变换矩阵：
$$
M = \begin{vmatrix} 
\frac{\cotθ}{Aspect} & 0 & 0 & 0 \\ 
0 & \cotθ & 0 & 0 \\ 
0 & 0 & \frac{f}{f-n} & 1 \\ 
0 & 0 & \frac{n·f}{n-f} & 0 \\ 
\end{vmatrix} \\
\text{f: 远视口距离} \quad \text{θ: 视口夹角} \quad \text{Aspect: 纵横比} \quad \text{n: 近视口距离}
$$

---

提示：MultiplyPoint主要用于Camera的投影变换，对于一般物体的矩阵变换用MultiplyPoint3x4，不涉及投影变换，计算速度也更快。

---

代码：参考Camera类中的cameraToWorldMatrix代码功能。

#### 6.1.2 MultiplyPoint3X4：矩阵变换

```
基本语法 public Vector3 MultiplyPoint3x4(Vector3 v);
```

功能说明：对参数值点v进行矩阵变换。因为矩阵变换中，不涉及投影变换，所以速度比MultiplyPoint快。例如，设m1为Matrix4x4实例，v1为Vector3实例，则`Vector3 v2=m1. MultiplyPoint3x4(v1)`，即为v2=v3*m1，其中v3=(v1,w)，w默认值为1。

#### 6.1.3 MultiplyVector方法：矩阵变换

```
基本语法 public Vector3 MultiplyVector(Vector3 v);
```

功能说明：对向量v进行矩阵变换。把v当做方向向量而非坐标点进行变换，当用矩阵与v进行变换时，只是对v的方向进行转换，即系统会对变换Matrix4x4进行特殊处理：设M为参与变换的Matrix4x4实例，其值为：
$$
M = \begin{vmatrix} 
m00 & m01 & m02 & m03 \\ 
m10 & m11 & m12 & m13 \\ 
m20 & m21 & m22 & m23 \\
m30 & m31 & m32 & m33 
\end{vmatrix}
$$
则系统处理后的M值为：
$$
M' = \begin{vmatrix} 
n00 & n01 & n02 & 0 \\ 
n10 & n11 & n12 & 0 \\ 
n20 & n21 & n22 & 0 \\
0 & 0 & 0 & 1
\end{vmatrix} \\
其中n00^2 + n10^2 + n20^2 = 1, \quad n01^2 + n11^2 + n21^2 = 1, \quad n02^2 + n12^2 + n22^2 = 1
$$
代码：此处以DirectX为例。在DirectX中向量变换是v·M，而在OpenGL中向量变换形式是M·v。

```C#
using UnityEngine;
using System.Collections;
public class MultiplyVector_ts : MonoBehaviour
{
    public Transform tr;
    Matrix4x4 mv0 = Matrix4x4.identity;
    Matrix4x4 mv1 = Matrix4x4.identity;
    void Start()
    {
        //分别设置变换矩阵mv0和mv1的位置变换和角度变换都不为0
        mv0.SetTRS(Vector3.one * 10.0f, Quaternion.Euler(new Vector3(0.0f, 30.0f, 0.0f)),Vector3.one);
        mv1.SetTRS(Vector3.one * 10.0f, Quaternion.Euler(new Vector3(0.0f, 0.6f, 0.0f)),Vector3.one);
    }
    void Update()
    {
        tr.position = mv1.MultiplyVector(tr.position); //用tr来定位变换后的向量
    }
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 120.0f, 45.0f), "方向旋转30度"))
        {
            Vector3 v = mv0.MultiplyVector(new Vector3(10.0f, 0.0f, 0.0f));
            Debug.Log("变换后向量："+v); //(8.7,0.0,-5.0)
            Debug.Log("变换后向量模长：" + v.magnitude);  //10
            //尽管mv0的位置变换不为0，但变换后向量的长度应与变换前相同
        }
    }
}
```

#### 6.1.4 SetTRS方法：重设Matrix4x4变换矩阵

```
基本语法 public void SetTRS(Vector3 pos, Quaternion q, Vector3 s);
		pos为位置向量，q为旋转角，s为放缩向量。
```

功能说明：此方法用来重设Matrix4x4变换矩阵。设有如下代码：

```C#
Matrix4x4 m1 = Matrix4x4.identity;
m1.SetTRS(pos,q,s);
Vector3 v2 = m1.MultiplyPoint3x4(v1);
则v2的值等于将v1的position增加pos，rotation旋转q，scale放缩s后的值。
```

代码：

```C#
using UnityEngine;
using System.Collections;
public class SetTRS_ts : MonoBehaviour
{
    Vector3 v1 = Vector3.one;
    Vector3 v2 = Vector3.zero;
    void Start()
    {
        Matrix4x4 m1 = Matrix4x4.identity;
        //Position沿y轴增加5个单位，绕y轴旋转45度，放缩2倍
        m1.SetTRS(Vector3.up * 5, Quaternion.Euler(Vector3.up * 45.0f), Vector3.one * 2.0f);
        //也可以使用如下静态方法设置m1变换
        //m1 = Matrix4x4.TRS(Vector3.up * 5, Quaternion.Euler(Vector3.up * 45.0f), Vector3.one* 2.0f);
        v2 = m1.MultiplyPoint3x4(v1);
        Debug.Log("v1的值：" + v1);  //(1.0,1.0,1.0)
        Debug.Log("v2的值：" + v2); //(2.8,7.0,0.0)
    }
    
    void FixedUpdate()
    {
        Debug.DrawLine(Vector3.zero, v1, Color.green);
        Debug.DrawLine(Vector3.zero, v2, Color.red);
    }
}
```

在Start中初始化m1，并调用SetTRS重置m1，接着调用方法MultiplyPoint3x4对向量v1进行变换，并将变换后的值赋给v2，最后在FixedUpdate方法中根据v1和v2的值绘制了两条直线。v1向v2变换顺序如下:

- v1绕y轴旋转45度后变为(1.414,1.0,0.0)，即x轴的分量变为了原来x和z分量的长度值，y值不变。
- 对v1扩大2倍后v1变为(2.8,2.0,0.0)。
- v1沿着y轴增加5个单位后变为(2.8,7.0,0.0)。

### 6.2 Matrix4x4 类静态方法

在Matrix4x4类中静态方法有`Ortho`、`Perspective`和`TRS`。

#### 6.2.1 Ortho：创建正交投影矩阵

```
基本语法 public static Matrix4x4 Ortho(float left, float right, float bottom, float top,
float zNear, float zFar);
        left:正交视口左边边长 right:正交视口右边边长 bottom:正交视口下部边长  
        top:正交视口上部边长  zNear:近视口距离      zFar:远视口距离。
```

功能说明：创建一个正交投影矩阵。

---

提示：

- left、right、bottom和top分正负方向，一般以right和top为正数，left和bottom为负数。
- left与right的值不能相等，bottom与top的值也不能相等，否则程序会报错。
- 为防止视图变形，参数的设定通常需要和Camera的aspect结合使用。

---

代码：参考Perspective

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Book/Ortho%E6%AD%A3%E4%BA%A4%E6%8A%95%E5%BD%B1%E7%9F%A9%E9%98%B5%E5%9B%BE.png" style="zoom:80%;" />

#### 6.2.2 Perspective方法：创建透视投影矩阵

```
基本语法 public static Matrix4x4 Perspective(float fov,float aspect,float zNear,float zFar);
        fov:视口夹角 aspect:视口纵横比例 zNear:近视口距离 zFar:远视口距离
```

功能说明：创建一个透视投影矩阵。若要更改摄像机的透视投影矩阵，可以用如下代码：

```
Camera.main. projectionMatrix=Matrix4x4.Perspective(fov,aspect,zNerar,zFar);
若要重置其投影矩阵，需要使用代码：
Camera.main. ResetProjectionMatrix ();
其中aspect=width/height
```

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Book/Perspective%E9%80%8F%E8%A7%86%E6%8A%95%E5%BD%B1%E7%9F%A9%E9%98%B5%E5%9B%BE.png" style="zoom:80%;" />

代码：

```C#
using UnityEngine;
using System.Collections;
public class OrthoAndPerspective_ts : MonoBehaviour
{
    Matrix4x4 Perspective = Matrix4x4.identity;//透视投影变量
    Matrix4x4 ortho = Matrix4x4.identity;//正交投影变量
    //声明变量，用于记录正交视口的左、右、下、上的值
    float l, r, b, t;
    void Start()
    {
        //设置透视投影矩阵
        Perspective = Matrix4x4.Perspective(65.0f, 1.5f, 0.1f, 500.0f);
        t = 10.0f;
        b = -t;
        //为防止视图变形需要与 Camera.main.aspect相乘
        l = b * Camera.main.aspect;
        r = t * Camera.main.aspect;
        //设置正交投影矩阵
        ortho = Matrix4x4.Ortho(l, r, b, t, 0.1f, 100.0f);
    }
    void OnGUI()
    {
        //使用默认正交投影
        if (GUI.Button(new Rect(10.0f, 8.0f, 150.0f, 20.0f), "Reset Ortho"))
        {
            Camera.main.orthographic = true;
            Camera.main.ResetProjectionMatrix();
            Camera.main.orthographicSize = 5.1f;
    	}
        //使用自定义正交投影
        if (GUI.Button(new Rect(10.0f, 38.0f, 150.0f, 20.0f), "use Ortho"))
        {
            ortho = Matrix4x4.Ortho(l, r, b, t, 0.1f, 100.0f);
            Camera.main.orthographic = true;
            Camera.main.ResetProjectionMatrix();
            Camera.main.projectionMatrix = ortho;
            Camera.main.orthographicSize = 5.1f;
        }
        //使用自定义透视投影
        if (GUI.Button(new Rect(10.0f, 68.0f, 150.0f, 20.0f), "use Perspective"))
        {
            Camera.main.orthographic = false;
            Camera.main.projectionMatrix = Perspective;
        }
        //恢复系统默认透视投影
        if (GUI.Button(new Rect(10.0f, 98.0f, 150.0f, 20.0f), "Reset Perspective"))
        {
            Camera.main.orthographic = false;
            Camera.main.ResetProjectionMatrix();
        }
    }
}
```

在Start中创建一个透视投影矩阵和一个正交投影矩阵，在OnGUI方法中编写了4种不同的投影方式：正交投影、自定义正交投影、自定义透视投影和系统默认透视投影。

#### 6.2.3 TRS：返回Matrix4x4 实例

```
基本语法 public static Matrix4x4 TRS(Vector3 pos, Quaternion q, Vector3 s);
		pos:位置向量 q:旋转角 s:放缩向量。
```

功能说明：使用pos、q和s作为变换参数返回一个Matrix4x4实例

```C#
Matrix4x4 m1 = Matrix4x4.TRS(pos,q,s);
Vector v2 = m1.MultiplyPoint3x4(v1);
v2等于将v1的position增加pos，rotation旋转q，scale放缩s
```

代码：参考SetTRS

## 7.Object类

Object类是Unity中所有对象的基类，例如GameObject、Component、Material、Shader、Texture、
Mesh、Font等都是Object的子类。本章主要介绍了Object类的实例方法和静态方法。

### 7.1 Object类实例方法

#### 7.1.1 GetInstanceID：Object对象ID

```
基本语法 public int GetInstanceID();
```

功能说明：返回Object对象的实例化ID，每个实例都有唯一的ID（int类型）

代码：

```C#
using UnityEngine;
using System.Collections;
public class GetInstanceID_ts : MonoBehaviour 
{
    void Start () 
    {
        Debug.Log("gameObject的ID："+ gameObject.GetInstanceID());
        Debug.Log("transform的ID："+ transform.GetInstanceID());
        GameObject g1, g2;
        //从GameObject创建一个对象
        g1 = GameObject.CreatePrimitive(PrimitiveType.Cube);
        //克隆对象
        g2 = Instantiate(g1,Vector3.zero,Quaternion.identity) as GameObject;
        Debug.Log("GameObject g1的ID："+ g1.GetInstanceID());
        Debug.Log("Transform g1的ID："+ g1.transform.GetInstanceID());
        Debug.Log("GameObject g2的ID：" + g2.GetInstanceID());
        Debug.Log("Transform g2的ID：" + g2.transform.GetInstanceID());
    }
}
```

打印出了gameObject和transform的InstanceID，然后创建和实例化两个新对象g1和g2，打印g1和g2的ID。

### 7.2 Object类静态方法

在Object类中的静态方法有`Destroy`、`DontDestroyOnLoad`、`FindObjectOfType`、
`FindObjectsOfType`和`Instantiate`。

#### 7.2.1 Destroy：销毁对象

```
基本语法 (1) public static void Destroy(Object obj);
        (2) public static void Destroy(Object obj, float t);
        obj:待销毁的对象，t:销毁延迟时间，默认为0。
```

功能说明：执行方法t秒后销毁obj对象。Destroy也可以销毁GameObject对象中的某个组件如Rigidbody、脚本等。

---

提示：相近的方法DestroyImmediate立即销毁某个Object对象及其在Assets中的资源文件，编程中慎用，推荐使用Destroy方法。

```
基本语法 DestroyImmediate(Object obj, bool allowDestroyingAssets = false);
```

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class Destroy_ts : MonoBehaviour 
{
    public GameObject GO,Cube;
    void Start () 
    {
        //5秒后销毁GO对象的Rigidbody组件
        Destroy(GO.rigidbody,5.0f);
        //7秒后销毁GO对象中的Destroy_ts脚本
        Destroy(GO.GetComponent<Destroy_ts>(),7.0f);
        //10秒后销毁Cube对象，同时Cube对象的所有组件及子类将一并销毁
        Destroy(Cube, 10.0f);
    }
}
```

#### 7.2.2 DontDestroyOnLoad：新场景中保留对象

```
基本语法 public static void DontDestroyOnLoad(Object target);
		target:被保留的对象。
```

功能说明：设置参数target指向的对象是否在新Scene中被保留下来。

- 如果target为父物体其子物体都会被导入到新Scene中。

代码：

```C#
using UnityEngine;
using System.Collections;
public class DontDestoryOnLoad_ts : MonoBehaviour
{
    public GameObject g1, g2;
    public Renderer re1, re2;
    void Start()
    {
        //g1指向一个顶层父物体对象,在导入新Scene时g1被保存
        DontDestroyOnLoad(g1);
        //g2指向一个子类对象,在导入新Scene时会发现g2没有被保存
        DontDestroyOnLoad(g2);
        //re1指向一个顶层父物体的Renderer组件,在导入新Scene时re1被保存
        DontDestroyOnLoad(re1);
        //re2指向一个子类对象的renderer组件，在导入新Scene时会发现re2指向的对象及组件没有被保存
        DontDestroyOnLoad(re2);
        Application.LoadLevel("FindObjectsOfType_unity");
    }
}
```

#### 7.2.3 FindObjectsOfType：获取对象

```
基本语法 (1) public static T[] FindObjectsOfType<T>() where T : Object;
		(2) public static Object[] FindObjectsOfType(Type type);
		type：对象类型
```

功能说明：获取工程中所有符合参数类型的对象。遍历整个工程，执行速度较慢，不适宜在每帧中调用。

---

提示：FindObjectOfType相近，获取工程中符合type类型的第一个对象，多用于检测工程中是否含有某种类型的对象。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class FindObjectOfType_ts : MonoBehaviour 
{
    void Start () 
    {
        GameObject[] gos = FindObjectsOfType(typeof(GameObject)) as GameObject[];
        foreach(GameObject go in gos)
        {
            //1.5秒后销毁除摄像机外的所有GameObject
            if (go.name != "Main Camera") Destroy(go, 1.5f);
        }
        
        Rigidbody[] rbs = FindObjectsOfType(typeof(Rigidbody))as Rigidbody[];
        foreach(Rigidbody rb in rbs)
        {
            //启用除球体外的所有刚体的重力感应
            if(rb.name!="Sphere") rb.useGravity = true;
        }
    }
}
```

调用FindObjectsOfType查找游戏中所有的GameObject对象，并将查找结果赋给数组gos，然后遍历数组gos，在1.5秒后销毁除摄像机外的所有GameObject对象。

#### 7.2.4 Instantiate：实例化对象

```
基本语法 (1) public static Object Instantiate(Object original);
        (2) public static Object Instantiate(Object original, Vector3 position,Quaternion rotation);
        original：类型，position：位置，rotation：旋转角度。
```

功能说明：实例化一个Object对象。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Instantiate_ts : MonoBehaviour 
{
    public GameObject A;
    public Transform B;
    public Rigidbody C;
    void Start () 
    {
        Object g1 = Instantiate(A,Vector3.zero,Quaternion.identity) as Object;
        Debug.Log("克隆一个Object对象g1:"+g1);
        GameObject g2 = Instantiate(A, Vector3.zero, Quaternion.identity) as GameObject;
        Debug.Log("克隆一个GameObject对象g2:" + g2);
        Transform t1 = Instantiate(B, Vector3.zero, Quaternion.identity) as Transform;
        Debug.Log("克隆一个Transform对象t1:" + t1);
        Rigidbody r1 = Instantiate(C, Vector3.zero, Quaternion.identity) as Rigidbody;
        Debug.Log("克隆一个Rigidbody对象r1:" + r1);
    }
}
```

在Start中调用Instantiate分别实例化了4种不同类型的对象，并将实例化的对象打印出来。

## 8.Quaternaion类

Quaternion又称四元数，由x、y、z和w这4个分量组成，属于struct类型。在Unity中，用Quaternion来存储和表示对象的旋转角度。Quaternion的变换比较复杂，对于GameObject一般的旋转及移动，可以用Transform中的相关方法实现。本章主要介绍了Quaternion类的实例属性、静态方法和运算符`*`.

### 8.1 Quaternion类实例属性

#### 8.1.1 eulerAngles：欧拉角

```
基本语法 public Vector3 eulerAngles { get; set; }
```

功能说明：返回Quaternion实例对应的欧拉角。<span style="color:red;">注意不同的旋转次序得到的最终状态是不同的。</span>

- 对Transform的欧拉角变换次序是
  1. 先绕z轴旋转相应的角度
  2. 再绕x轴旋转相应的角度
  3. 最后绕y轴旋转相应的角度

代码：

```C#
using UnityEngine;
using System.Collections;
public class EulerAngle_ts : MonoBehaviour
{
    public Transform A, B;
    Quaternion rotations=Quaternion.identity;
    Vector3 eulerAngle = Vector3.zero;
    float speed = 10.0f;
    void Update()
    {
        //第一种方式：将Quaternion赋值给transform的rotation
        rotations.eulerAngles = new Vector3(0.0f, speed * Time.time, 0.0f);
        A.rotation = rotations;
        //第二种方式：将三维向量代表的欧拉角直接赋值给transform的eulerAngles
        eulerAngle = new Vector3(0.0f, speed * Time.time, 0.0f);
        B.eulerAngles = eulerAngle;
    }
}
```

### 8.2 Quaternion类实例方法

涉及的实例方法有`SetFromToRotation`、`SetLookRotation`和`ToAngleAxis`，静态方法AngleAxis和实例方法ToAngleAxis一起介绍。

#### 8.2.1 SetFromToRotation：创建rotation实例

```
基本语法 public void SetFromToRotation(Vector3 fromDirection, Vector3 toDirection);
```

功能说明：从fromDirection到toDirection的旋转。

```C#
Quaternion q1 = Quaternion.identity;  //新建
q1.SetFromToRotation(v1,v2);		  //旋转
transform.rotation = q1;		      //赋值
```

将GameObject对象自身坐标系中向量v1指向的方向旋转到v2方向。

---

提示：不可直接使用transform.rotation.SetFromToRotation(v1,v2)方式进行设置，只能将实例化的Quaternion赋值给transform.rotation。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class SetFromToRotation_ts : MonoBehaviour 
{
    public Transform A, B, C;
    Quaternion q1 = Quaternion.identity;
    void Update () 
    {
        //不可直接使用C.rotation.SetFromToRotation(A.position,B.position);
        q1.SetFromToRotation(A.position,B.position);
        C.rotation = q1;
        //在Scene面板中绘制直线
        Debug.DrawLine(Vector3.zero,A.position,Color.red);
        Debug.DrawLine(Vector3.zero, B.position, Color.green);
        Debug.DrawLine(C.position, C.position+new Vector3(0.0f,1.0f,0.0f), Color.black);
        Debug.DrawLine(C.position, C.TransformPoint(Vector3.up*1.5f), Color.yellow);
    }
}
```

A为起始坐标，B为结束坐标，C为旋转物体，在Update方法中对新建的q1调用方法SetFromToRotation，并将改变后的q1赋给C.rotation。在Scene场景中绘制直线以观察，可以更改A或B的位置查看物体C的状态变化。

#### 8.2.2 SetLookRotation：设置Quaternion实例的朝向

```
基本语法 (1) public void SetLookRotation(Vector3 view);
		(2) public void SetLookRotation(Vector3 view, Vector3 up);
```

功能说明：设置Quaternion的朝向

```C#
Quaternion q1 = Quaternion.identity;
q1.SetLookRotation(v1, v2);
transform.rotation = q1;
```

- 物体`transform.forward`与v1方向相同。

---

提示：不可以直接使用transform.rotation.SetLookRotation(v1, v2)的方式来使用SetLookRotation
方法。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class SetLookRotation_ts : MonoBehaviour
{
    public Transform A, B, C;
    Quaternion q1 = Quaternion.identity;
    void Update()
    {
        q1.SetLookRotation(A.position, B.position);
        C.rotation = q1;
        //分别绘制A、B和C.right的朝向线
        //请在Scene视图中查看
        Debug.DrawLine(Vector3.zero, A.position, Color.red);
        Debug.DrawLine(Vector3.zero, B.position, Color.green);
        Debug.DrawLine(C.position, C.TransformPoint(Vector3.right * 2.5f), Color.yellow);
        Debug.DrawLine(C.position, C.TransformPoint(Vector3.forward * 2.5f), Color.gray);
        //分别打印C.right与A、B的夹角
        Debug.Log("C.right与A的夹角:" + Vector3.Angle(C.right, A.position));
        Debug.Log("C.right与B的夹角:" + Vector3.Angle(C.right, B.position));
        //C.up与B的夹角
        Debug.Log("C.up与B的夹角:" + Vector3.Angle(C.up, B.position));
    }
}
```

A为起始坐标，B为结束坐标，C为旋转物体，在Update方法中对新建的q1调用方法SetLookRotation，并将改变后的q1赋给C.rotation。在Scene场景中绘制直线以观察，可以更改A或B的位置查看物体C的状态变化。

#### 8.2.3 ToAngleAxis：Quaternion实例的角轴表示

```
基本语法 public void ToAngleAxis(out float angle, out Vector3 axis);
		参数angle:旋转角 参数axis:轴向量
```

功能说明：将Quaternion实例转换为角轴表示。在transform.rotation.ToAngleAxis(out angle, out axis)中，输出值angle和axis的含义为：要将GameObject对象的rotation从Quaternion.Identity状态变换到当前状态，只需要将GameObject对象绕着axis的轴向（指世界坐标系中）旋转angle角度即可。

---

提示：此方法通常和静态方法AngleAxis (angle : float, axis : Vector3)联合使用，使得一个物体的rotation始终和另一个物体的rotation保持一致。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class ToAngleAxis_ts : MonoBehaviour 
{
    public Transform A, B;
    float angle;
    Vector3 axis = Vector3.zero;
    void Update () 
    {
        //使用ToAngleAxis获取A的Rotation的旋转轴和角度
        A.rotation.ToAngleAxis(out angle, out axis);
        //使用AngleAxis设置B的rotation，使得B的rotation状态的和A相同
        //当然也可以只使得B与A的axis相同，而angle不同
        //可以在程序运行时修改A的rotation查看B的状态
        B.rotation = Quaternion.AngleAxis(angle,axis);
    }
}
```

声明两个GameObject变量A和B，用于指向场景中的物体，然后在Update方法中，首先调用ToAngleAxis方法将A的rotation转换为角轴angle和axis，调用方法AngleAxis将角轴代表的Quaternion值赋给B.rotation。修改A的旋转角查看物体B的相应旋转状态。

### 8.3 Quaternion类静态方法

在Quaternion类中涉及的静态方法有`Angle`、`Dot`、`Euler`、`FromToRotation`、`Inverse`、`Lerp`、`LookRotation`、`RotateTowards`和`Slerp`。

#### 8.3.1 Angle：Quaternion实例间夹角

```
基本语法 public static float Angle(Quaternion a, Quaternion b);
```

功能说明：返回从参数a到参数b变换的夹角。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Angle_ts : MonoBehaviour 
{
    void Start()
    {
        Quaternion q1 = Quaternion.identity;
        Quaternion q2 = Quaternion.identity;
        q1.eulerAngles = new Vector3(10.0f, 20.0f, 30.0f);
        float f1 = Quaternion.Angle(q1,q2);
        float f2 = 0.0f;
        Vector3 v1 = Vector3.zero;
        q1.ToAngleAxis(out f2, out v1);
        Debug.Log("f1:" + f1);
        Debug.Log("f2:" + f2);
        Debug.Log("q1的欧拉角：" + q1.eulerAngles + " q1的rotation：" + q1);
        Debug.Log("q2的欧拉角：" + q2.eulerAngles + " q2的rotation：" + q2);
    }
}
```

对q1的欧拉角进行赋值，调用方法Angle求q1和q2之间的夹角，并将返回值赋给f1，最后调用方法ToAngleAxis，求解从当前q1状态转换到Quaternion.identity状态需要旋转的最小角度值f2。

#### 8.3.2 Dot：点乘

```
基本语法 public static float Dot(Quaternion a, Quaternion b);
```

功能说明：a和b的点乘。

```
q1(x1,y1,z1,w1)	q2(x2,y2,z2,w2)
float f = Quaternion.Dot(q1,q2); 
等价于f=x1*x2+y1*y2+z1*z2+w1*w2，结果值f的范围为[-1,1]
```

代码：

```C#
using UnityEngine;
using System.Collections;
public class Dot_ts : MonoBehaviour 
{
    public Transform A, B;
    Quaternion q1=Quaternion.identity;
    Quaternion q2=Quaternion.identity;
    float f;
    void Start () 
    {
        A.eulerAngles = new Vector3(0.0f,40.0f,0.0f);
        //B比A绕y轴多转360度
        B.eulerAngles = new Vector3(0.0f, 360.0f+40.0f, 0.0f);
        q1 = A.rotation;
        q2 = B.rotation;
        f = Quaternion.Dot(q1,q2);
        Debug.Log("q1的rotation:"+q1);
        Debug.Log("q2的rotation:" + q2);
        Debug.Log("q1的欧拉角:" + q1.eulerAngles);
        Debug.Log("q2的欧拉角:" + q2.eulerAngles);
        Debug.Log("Dot(q1,q2):"+f);
    }
}
```

将A和B的rotation分别赋予q1和q2，调用Dot方法将q1和q2的点乘值赋给f。从输出结果可知q1和q2的欧拉角相等，但它们的值却相反。当Dot返回值为-1时，两个参数的角度差值相差一周。

#### 8.3.3 Euler：欧拉角对应的四元数

```
基本语法 (1) public static Quaternion Euler(Vector3 euler);
		(2) public static Quaternion Euler(float x, float y, float z);
```

功能说明：返回欧拉角Vector3(x,y,z) 对应的四元数Quaternion 实例。四元数
Quaternion(qx,qy,qz,qw)与其欧拉角eulerAngles(ex,ey,ez)之间的对应关系如下。$PIover180 = \frac{3.141592}{180} = 0.0174532925$是计算机图形学中的一个常量，变换过程如下。
$$
\begin{aligned}
e_x &=  \frac{e_x \times PIover180}{2.0f} \\
e_y &=  \frac{e_y \times PIover180}{2.0f} \\
e_z &=  \frac{e_z \times PIover180}{2.0f}
\end{aligned}
$$

$$
\begin{aligned}
q_x &= \sin(e_x)\cos(e_y)\cos(e_z) + \cos(e_x)\sin(e_y)\sin(e_z) \\
q_y &= \cos(e_x)\sin(e_y)\cos(e_z) - \sin(e_x)\cos(e_y)\sin(e_z) \\
q_z &= \cos(e_x)\cos(e_y)\sin(e_z) - \sin(e_x)\sin(e_y)\cos(e_z) \\
q_w &= \cos(e_x)\cos(e_y)\cos(e_z) + \sin(e_x)\sin(e_y)\sin(e_z)
\end{aligned}
$$

代码：

```C#
using UnityEngine;
using System.Collections;
public class Euler_ts : MonoBehaviour
{
    //记录欧拉角，单位为角度，可以在Inspector面板中设置
    public float ex, ey, ez;
    //用于记录计算结果
    float qx, qy, qz, qw;
    float PIover180 = 0.0174532925f;//常量
    Quaternion Q = Quaternion.identity;
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 100.0f, 45.0f), "计算"))
        {
            Debug.Log("欧拉角：" + " ex：" + ex + " ey：" + ey + " ez：" + ez);
            //欧拉角:ex:30 ey:40 ez:50
            Q = Quaternion.Euler(ex, ey, ez);
            Debug.Log("Q.x:" + Q.x + " Q.y:" + Q.y + " Q.z:" + Q.z + " Q.w:" + Q.w);
            //Q.x:0.3600422 Q.y:0.1966282 Q.z:0.3033718 Q.w:0.8600422
            ex = ex * PIover180 / 2.0f;
            ey = ey * PIover180 / 2.0f;
            ez = ez * PIover180 / 2.0f;
            qx = Mathf.Sin(ex) * Mathf.Cos(ey) * Mathf.Cos(ez) + Mathf.Cos(ex) *
            Mathf.Sin(ey) * Mathf.Sin(ez);
            qy = Mathf.Cos(ex) * Mathf.Sin(ey) * Mathf.Cos(ez) - Mathf.Sin(ex) *
            Mathf.Cos(ey) * Mathf.Sin(ez);
            qz = Mathf.Cos(ex) * Mathf.Cos(ey) * Mathf.Sin(ez) - Mathf.Sin(ex) *
            Mathf.Sin(ey) * Mathf.Cos(ez);
            qw = Mathf.Cos(ex) * Mathf.Cos(ey) * Mathf.Cos(ez) + Mathf.Sin(ex) *
            Mathf.Sin(ey) * Mathf.Sin(ez);
            Debug.Log("qx:" + qx + " qy:" + qy + " qz:" + qz + " qw:" + qw);
            //qx:0.3600422 qy:0.3033718 qw:0.8600422
        }
    }
}
```

#### 8.3.4 FromToRotation：Quaternion变换

```
基本语法 public static Quaternion FromToRotation(Vector3 fromDirection, Vector3 toDirection);
```

功能说明：从参数fromDirection到toDirection的Quaternion变换。

代码：

```C#
using UnityEngine;
using System.Collections;
public class FromToRotation_ts : MonoBehaviour
{
    public Transform A, B, C, D;
    Quaternion q1 = Quaternion.identity;
    void Update()
    {
        //使用实例方法
        q1.SetFromToRotation(A.position, B.position);
        C.rotation = q1;
        //使用类方法
        D.rotation = Quaternion.FromToRotation(A.position, B.position);
        //在Scene视图中绘制直线
        Debug.DrawLine(Vector3.zero, A.position, Color.white);
        Debug.DrawLine(Vector3.zero, B.position, Color.white);
        Debug.DrawLine(C.position, C.position + new Vector3(0.0f, 1.0f, 0.0f),Color.white);
        Debug.DrawLine(C.position, C.TransformPoint(Vector3.up * 1.5f), Color.white);
        Debug.DrawLine(D.position, D.position + new Vector3(0.0f, 1.0f, 0.0f),Color.white);
        Debug.DrawLine(D.position, D.TransformPoint(Vector3.up * 1.5f), Color.white);
    }
}
```

4个Transform变量指向场景中的对象。在Update方法中调用方法SetFromToRotation，使得C的旋转角度与向量A和B之间的夹角相同，接着使用方法FromToRotation，使得D的旋转角度与向量A和B之间的夹角相同，在Scene视图中拖动A或B的位置观察C和D的变化。

<img src="https://gitee.com/u9king/ImageHostingService/raw/master/Unity/Book/FromToRotation%E6%BC%94%E7%A4%BA.png" style="zoom:80%;" />

#### 8.3.5 Inverse：逆向Quaternion值

```
基本语法 public static Quaternion Inverse(Quaternion rotation);
```

功能说明：返回rotation的Quaternion逆向值。例如，设有实例rotation=(x,y,z,w)，则Inverse(rotation)=(-x,-y,-z,w)。从效果上说，设rotation. eulerAngles=(a,b,c)，则transform.rotation=Inverse(rotation)相当于transform依次绕自身坐标系的z轴、x轴和y轴分别旋转-c度、-a度和-b度。由于是局部坐标系内的变换，最后transform的欧拉角的各个分量值并不一定等于-a、-b或-c。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Inverse_ts : MonoBehaviour
{
    public Transform A, B;
    void Start()
    {
        Quaternion q1 = Quaternion.identity;
        Quaternion q2 = Quaternion.identity;
        q1.eulerAngles = new Vector3(10.0f, 20.0f, 30.0f);
        q2 = Quaternion.Inverse(q1);
        A.rotation = q1;
        B.rotation = q2;
        Debug.Log("q1的欧拉角：" + q1.eulerAngles + " q1的rotation：" + q1);
        Debug.Log("q2的欧拉角：" + q2.eulerAngles + " q2的rotation：" + q2);
    }
}
```

#### 8.3.6 Lerp：线性插值

```
基本语法 public static Quaternion Lerp(Quaternion from, Quaternion to, float t);
```

功能说明：返回从参数from到to的线性插值。当参数t≤0时返回值为from，当参数t≥1时返回值为to。此方法执行速度比Slerp方法快，一般情况下可代替Slerp方法。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Slerp_ts : MonoBehaviour
{
    public Transform A, B, C, D;
    float speed = 0.2f;
    //分别演示方法Slerp和Lerp的使用
    void Update()
    {
        C.rotation = Quaternion.Slerp(A.rotation, B.rotation, Time.time * speed);
        D.rotation = Quaternion.Lerp(A.rotation, B.rotation, Time.time * speed);
    }
}
```

4个Transform变量A、B、C和D，分别指向场景中不同的物体对象，在Update方法中分别演示了方法Slerp和Lerp的使用。

#### 8.3.7 LookRotation：设置Quaternion的朝向

```
基本语法 (1) public static Quaternion LookRotation(Vector3 forward);
        (2) public static Quaternion LookRotation(Vector3 forward, Vector3 upwards);
        参数forward：返回Quaternion的forward朝向。
```

功能说明：返回一个Quaternion实例，使GameObject对象的z轴朝向参数forward方向。

代码：

```C#
using UnityEngine;
using System.Collections;
public class LookRotation_ts : MonoBehaviour
{
    public Transform A, B, C, D;
    Quaternion q1 = Quaternion.identity;
    void Update()
    {
        //使用实例方法
        q1.SetLookRotation(A.position, B.position);
        C.rotation = q1;
        //使用类方法
        D.rotation = Quaternion.LookRotation(A.position, B.position);
        //绘制直线，请在Scene视图中查看
        Debug.DrawLine(Vector3.zero, A.position, Color.white);
        Debug.DrawLine(Vector3.zero, B.position, Color.white);
        Debug.DrawLine(C.position, C.TransformPoint(Vector3.up * 2.5f), Color.white);
        Debug.DrawLine(C.position, C.TransformPoint(Vector3.forward * 2.5f),Color.white);
        Debug.DrawLine(D.position, D.TransformPoint(Vector3.up * 2.5f), Color.white);
        Debug.DrawLine(D.position, D.TransformPoint(Vector3.forward * 2.5f),Color.white);
	}
}
```

#### 8.3.8 RotateTowards：Quaternion插值

```
基本语法 public static Quaternion RotateTowards(Quaternion from, Quaternion to, float
maxDegreesDelta);
        from：起始Quaternion to：结束Quaternion maxDegreesDelta：每帧最大角度值。
```

功能说明：返回从参数from到to的插值，且返回值的最大角度不超过maxDegreesDelta。此方法功能与方法Slerp相似，只是maxDegreesDelta指的是角度值，不是插值系数。当maxDegreesDelta<0时，将沿着从to到from的方向插值计算。

代码：

```C#
using UnityEngine;
using System.Collections;
public class RotateTowards_ts : MonoBehaviour 
{
    public Transform A, B, C;
    float speed = 10.0f;
    void Update()
    {
        //调用方法RotateTowards，并将其返回值赋给C.rotation
        C.rotation = Quaternion.RotateTowards(A.rotation, B.rotation, Time.time *
        speed-40.0f);
        Debug.Log("C与A的欧拉角的差值：" + (C.eulerAngles-A.eulerAngles) + " maxDegreesDelta: +" (Time.time * speed - 40.0f));
    }
}
```

#### 8.3.9 Slerp：球面插值

```
基本语法 public static Quaternion Slerp(Quaternion from, Quaternion to, float t);
```

功能说明：返回从参数from到to的球面插值。当参数t≤0时返回值为from，当参数t≥1时返回值为to。一般情况下可用方法Lerp代替。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Slerp_ts : MonoBehaviour
{
    public Transform A, B, C, D;
    float speed = 0.2f;
    //分别演示方法Slerp和Lerp的使用
    void Update()
    {
        C.rotation = Quaternion.Slerp(A.rotation, B.rotation, Time.time * speed);
        D.rotation = Quaternion.Lerp(A.rotation, B.rotation, Time.time * speed);
    }
}
```

### 8.4 Quaternion类运算符

在Quaternion类中涉及的运算符运算有两个Quaternion相乘、Quaternion和一个Vector3相乘，下面介绍这两种不同的运算。

#### 8.4.1 operator * (lhs : Quaternion, rhs : Quaternion)

功能说明：返回两个Quaternion实例相乘后的结果。A 和B均为GameObject对象代码如下：

```
B.rotation *= A.rotation;
```

- 代码每执行一次，B都会绕着B的局部坐标系的z、x、y轴分别旋转A.eulerAngles.z度、A.eulerAngles.x度和A.eulerAngles.y度，注意它们的旋转次序一定是先饶z轴再绕x轴最后绕y轴进行相应的旋转。另外由于是绕着局部坐标系旋转，故而当绕着其中一个轴旋转时，很可能会影响其余两个坐标轴方向的欧拉角（除非其余两轴的欧拉角都为0才不受影响）。
- 设A的欧拉角为euler_a(ax,ay,az)，则沿着B的初始局部坐标系的euler_a方向向下看，会发现B在绕euler_a顺时针旋转。B的旋转状况还受其初始状态的影响，可以在运行示例程序时在Inspector面板中更改B的初始欧拉角，从而查看运行状态的不同。

代码：

```C#
using UnityEngine;
using System.Collections;
public class QxQ_ts : MonoBehaviour 
{
    public Transform A, B;
    void Start () 
    {
        //设置A的欧拉角//试着更改各个分量查看B的不同旋转状态
        A.eulerAngles = new Vector3(1.0f,1.5f,2.0f);
    }
    void Update () 
    {
        B.rotation *= A.rotation;
        //输出B的欧拉角，注意观察B的欧拉角变化
        Debug.Log(B.eulerAngles);
    }
}
```

在Update方法中将B.rotation与A.rotation相乘的值赋给B.rotation，从而实现B对象的不断旋转。注意观察B的欧拉角的变化，B绕着其自身坐标系的Vector3(1.0f,1.5f,2.0f)方向旋转。虽然每次绕这个轴向旋转的角度相同，但角度的旋转在3个坐标轴上的值都不为零，其中一轴的旋转会影响其他两轴的角度，故而B的欧拉角的各个分量的每次递增值是不固定的。

#### 8.4.2 operator * (rotation : Quaternion, point : Vector3)

功能说明：对参数坐标点point进行rotation变换。A为Vector3：

```
transform.position += transform.rotation * A;
```

则每执行一次代码，transform对应的对象便会沿着自身坐标系中向量A的方向移动A的模长的距离。transform.rotation与A相乘主要来确定移动的方向和距离。

```C#
using UnityEngine;
using System.Collections;
public class QxV_ts : MonoBehaviour
{
    public Transform A;
    float speed = 0.1f;
    //初始化A的position和eulerAngles
    void Start()
    {
        A.position = Vector3.zero;
        A.eulerAngles = new Vector3(0.0f, 45.0f, 0.0f);
    }
    void Update()
    {
        //沿着A自身坐标系的forward方向每帧前进speed距离
        A.position += A.rotation * (Vector3.forward * speed);
        Debug.Log(A.position);
    }
}
```

在Update方法中，使用Quaternion与Vector3相乘，使得物体A沿着自身坐标系的forward方向每帧前进speed距离。

### 8.5 关于Quaternion类中相乘运算符的两种重载方式的注解

在Quaternion类中，相乘运算符（*）有两种重载方式，分别为operator * (lhs : Quaternion, rhs :
Quaternion)和operator * (rotation : Quaternion, point : Vector3)。

设A为两个Quaternion实例的乘积，结果为Quaternion类型，设B为Quaternion实例和Vector3的乘
积，结果为Vector3类型，则二者主要有以下异同。

- A与B的相似处是它们都通过自身坐标系的“相乘”方式来实现在世界坐标系中的变换；
- A主要用来实现transform绕着自身坐标系中的某个轴进行旋转，而B主要用来实现
    transform沿着自身坐标系的某个方向进行移动；
- B的相乘顺序只有Quaternion*Vector3一种形式，而没有Vector3*Quaternion的形式；
- 由于它们都是相对于自身坐标系进行的旋转或移动，故而当自身坐标系的轴向和世界坐标
    系的轴向不一致时，它们绕着自身坐标系中某个轴向的变动都会影响到物体在世界坐标系
    中各个坐标轴的变动。

## 9.Random类

Random类是Unity中用于产生随机数的类，不可实例化，只有静态属性和静态方法。

### 9.1 Random类静态属性

在Random类中，静态属性有`insideUnitCircle`、`insideUnitSphere`、`onUnitSphere`、`rotationUniform`、`rotation`和`seed`。

#### 9.1.1 insideUnitCircle：圆内随机点

```
基本语法 public static Vector2 insideUnitCircle { get; }
```

功能说明：返回一个半径为1的圆内的随机点坐标，返回值为Vector2类型。

---

提示：类似属性还有

- insideUnitSphere属性：返回一个半径为1 的球内的随机点坐标，返回值为Vector3类型；
- onUnitSphere属性：返回一个半径为1 的球表面的随机点坐标，返回值为Vector3类型。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class insideUnitCircle_ts : MonoBehaviour
{
    public GameObject go;
    void Start()
    {
        //每隔0.4秒执行一次use_rotationUniform方法
        InvokeRepeating("use_rotationUniform", 1.0f, 0.4f);
    }
    void use_rotationUniform()
    {
        //在半径为5的圆内随机位置实例化一个GameObject对象
        //Vector2实例转为Vector3时，z轴分量默认为0
        Instantiate(go, Random.insideUnitCircle * 5.0f, Quaternion.identity);
        //在半径为5的球内随机位置实例化一个GameObject对象
        Instantiate(go, Vector3.forward * 15.0f + 5.0f * Random.insideUnitSphere,
        Quaternion.identity);
        //在半径为5的球表面随机位置实例化一个GameObject对象
        Instantiate(go, Vector3.forward * 30.0f + 5.0f * Random.onUnitSphere,
        Quaternion.identity);
    }
}
```

在这段代码中，首先声明了一个公共的GameObject变量go，然后在方法use_rotationUniform中分别使用Random属性insideUnitCircle、insideUnitSphere和onUnitSphere的返回值作为实例化对象的参考位置，最后在Start方法中调用InvokeRepeating方法，每隔0.4秒执行一次use_rotationUniform方法。

![](https://gitee.com/u9king/ImageHostingService/raw/92ecee9b86de398bd49e723226ed730e44a42932/Unity/Book/insideUnitCircle%E9%9A%8F%E6%9C%BA%E7%82%B9%E7%94%9F%E6%88%90%E5%9B%BE.png)

#### 9.1.2 rotationUniform：均匀分布特征

```
基本语法 public static Quaternion rotationUniform { get; }
```

功能说明：返回一个随机且符合均匀分布特征的rotation值。

---

提示：Random类的rotation属性与此属性功能相近，不同之处在于rotation属性只是产生一个随机的rotation值，不符合均匀分布特征，但其计算速度比rotationUniform快，一般情况下可用Random类的rotation属性产生一个随机的rotation值。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class rotationUniform_ts : MonoBehaviour
{
    public GameObject go;
    GameObject cb, sp;
    GameObject cb1, sp1;
    void Start()
    {
        //分别获取cb、sp、cb1和sp1对象
        cb = GameObject.Find("Cube");
        sp = GameObject.Find("Cube/Sphere");
        cb1 = GameObject.Find("Cube1");
        sp1 = GameObject.Find("Cube1/Sphere1");
        //每隔0.4秒执行一次use_rotationUniform方法
        InvokeRepeating("use_rotationUniform", 1.0f, 0.4f);
    }
    void use_rotationUniform()
    {
        //使用rotationUniform产生符合均匀分布特征的rotation
        cb.transform.rotation = Random.rotationUniform;
        //使用rotation产生一个随机rotation
        cb1.transform.rotation = Random.rotation;
        //分别在sp和sp1的位置实例化一个GameObject对象
        Instantiate(go, sp.transform.position, Quaternion.identity);
        Instantiate(go, sp1.transform.position, Quaternion.identity);
    }
}
```

在方法use_rotationUniform中分别使用属rotationUniform和rotation参数两个随机的rotation值，并将它们分别赋予cb和cb1。最后分别在sp和sp1的位置实例化一个GameObject对象。

#### 9.1.3 seed：随机数种子

```
public static int seed { get; set; }
```

功能说明：设置随机数的种子。在计算机中产生随机数的方法有很多，但每种方法都需要一个种子，例如经典的伪随机数产生函数：f(x)=f(x-1)*a+b，其中a、b为已知的固定数值，那么只要知道某个x对应的f值，就可以推算出所有的值。通常情况下，会把f(0)当做随机数产生的种子，即只要知道了f(0)的值就可以推算出f(1)、f(2)…的值。总之，相同的Random.seed值对应着相同的随机数序列，如果不人为设定其值，Unity会根据某种算法自动产生一个种子。

```C#
using UnityEngine;
using System.Collections;
public class Seed_ts : MonoBehaviour
{
    void Start()
    {
        //设置随机数的种子
        //不同的种子产生不同的随机数序列
        //对于相同的种子，在程序每次启动时其序列是相同的
        Random.seed = 1;
    }
    void Update()
    {
        Debug.Log(Random.value);
        //[0.7323, 0.6478, 0.2391, 0.9382, 0.1523, ...]
        //UnityRandom的实现以前主要基于Xorshift需要2^31 - 1次才会出现重复
    }
}
```

#### 9.2 Random类其他常用静态属性功能简介

- insideUnitSphere：返回一个半径为1的球内的随机点坐标，返回值为Vector3类型；
- onUnitSphere：返回一个半径为1的球表面的随机点坐标，返回值为Vector3类型；
- rotation：用于返回一个随机的rotation值，返回值为Quaternion类型；
- value：用于返回一个[0.0f,1.0f]区间内的随机数。

## 10.RigidBody类

Rigidbody类的功能是用来模拟GameObject对象在现实世界中的物理特性，包括重力、阻力、质量、速度等。对Rigidbody对象属性的赋值代码通常放在脚本的OnFixedUpdate()方法中。

### 10.1 Rigidbody类实例属性

在Rigidbody类中涉及的实例属性有`collisionDetectionMode`、`drag`、`inertiaTensor`、`mass`和`velocity`。

#### 10.1.1 collisionDetectionMode：碰撞检测模式

```
基本语法 public CollisionDetectionMode collisionDetectionMode { get; set; }
```

功能说明：设置刚体的碰撞检测模式。刚体的碰撞检测模式有3种，即枚举类型CollisionDetectionMode的３个值。

- Discrete：静态离散检测模式，为系统的默认设置。在此模式下，只有在某一帧中两物体的碰撞器发生重叠时才能被检测到，这样就有可能导致某物体的前一帧在另一个刚体的上方，而下一帧移动到了另一个刚体的下方，这样就会发生穿越现象。
- Continuous：静态连续检测模式，一般用在高速运动刚体的目标碰撞体上，防止被穿越，检测强度比Discrete强。
- ContinuousDynamic：最强的连续动态检测模式，一般用在两个高速运动的物体上，防止互相穿越。其计算消耗最大，一般情况下慎用。

无论哪种检测模式都有可能被穿越，为了防止穿越现象的发生，除了设置其碰撞检测模式外，还要适当增加两物体碰撞器的厚度，一般不要小于0.1，同时尽量降低两物体碰撞时的相对速度。

代码：

```C#
using UnityEngine;
using System.Collections;
public class CollisionDetectionMode_ts : MonoBehaviour
{
    public Rigidbody A, B;
    Vector3 v1, v2;
    void Start()
    {
        A.useGravity = false;
        B.useGravity = false;
        v1 = A.position;
        v2 = B.position;
    }
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "Discrete模式不被穿越"))
        {
            inists();
            A.collisionDetectionMode = CollisionDetectionMode.Discrete;
            B.collisionDetectionMode = CollisionDetectionMode.Discrete;
            A.velocity = new Vector3(0.0f,-10.0f,0.0f);
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "Discrete模式被穿越"))
        {
            inists();
            A.collisionDetectionMode = CollisionDetectionMode.Discrete;
            B.collisionDetectionMode = CollisionDetectionMode.Discrete;
            A.velocity = new Vector3(0.0f, -40.0f, 0.0f);
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "Continuous模式不被穿越"))
        {
            inists();
            A.collisionDetectionMode = CollisionDetectionMode.Continuous;
            B.collisionDetectionMode = CollisionDetectionMode.Continuous;
            A.velocity = new Vector3(0.0f, -20.0f, 0.0f);
        }
        if (GUI.Button(new Rect(10.0f, 160.0f, 200.0f, 45.0f), "Continuous模式被穿越"))
        {
            inists();
            A.collisionDetectionMode = CollisionDetectionMode.Continuous;
            B.collisionDetectionMode = CollisionDetectionMode.Continuous;
            A.velocity = new Vector3(0.0f, -15.0f, 0.0f);
            B.velocity = new Vector3(0.0f, 15.0f, 0.0f);
        }
        if (GUI.Button(new Rect(10.0f, 210.0f, 200.0f, 45.0f), "ContinuousDynamic模式"))
        {
            inists();
            A.collisionDetectionMode = CollisionDetectionMode.ContinuousDynamic;
            B.collisionDetectionMode = CollisionDetectionMode.ContinuousDynamic;
            A.velocity = new Vector3(0.0f, -200.0f, 0.0f);
            B.velocity = new Vector3(0.0f, 200.0f, 0.0f);
        }
        if (GUI.Button(new Rect(10.0f, 260.0f, 200.0f, 45.0f), "重置"))
        {
            inists();
        }
    }
    //初始化A、B
    void inists() 
    {
        A.position = v1;
        A.rotation = Quaternion.identity;
        A.velocity = Vector3.zero;
        A.angularVelocity = Vector3.zero;
        B.position = v2;
        B.rotation = Quaternion.identity;
        B.velocity = Vector3.zero;
        B.angularVelocity = Vector3.zero;
    }
}
```

声明了两个Rigidbody变量A、B和两个Vector3变量v1、v2。在Start方法中设置A、B的useGravity属性为false，并将A、B的Position赋值给v1和v2，然后在OnGUI 方法中定义了多个Button ， 用来演示Discrete 、Continuous和ContinuousDynamic的功能。最后定义了一个inists方法，用于重置变量A、B对应刚体的状态。

#### 10.1.2 drag：刚体阻力

```
基本语法 public float drag { get; set; }
```

功能说明：给刚体添加一个阻力。drag值越大刚体速度减慢得越快，当drag>0时，刚体在增加到一定速度后会匀速移动。

---

提示：刚体在自由落体运动中的最大速度值只与Gravity和drag值有关，与质量Mass无关。

---

代码：

```C#
using UnityEngine;
using System.Collections;
public class Drag_ts : MonoBehaviour
{
    public Rigidbody R;
    float drags = 20.0f;
    string str = "200";
    //初始化R.drag
    void Start()
    {
        R.drag = drags;
    }
    void OnGUI()
    {
        GUI.Label(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "Gravity:" + Physics.gravity);
        str = (GUI.TextField(new Rect(10.0f, 60.0f, 200.0f, 45.0f), str));
        if (GUI.Button(new Rect(10.0f, 210.0f, 200.0f, 45.0f), "compute"))
        {
            drags = float.Parse(str);
            R.drag = drags;
        }
        GUI.Label(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "Velocity:" + R.velocity.y);
        GUI.Label(new Rect(10.0f, 160.0f, 200.0f, 45.0f), "drag:" + drags);
    }
}
```

在OnGUI方法中绘制了一些GUI组件，用于测试不同的drag值下，物体下落的最终速度。

| drag  | 0.1  | 0.2  | 0.3  | 0.4  | 0.5   | 0.6   | 0.7   | 0.8  | 0.9  | 1.0  |
| ----- | ---- | ---- | ---- | ---- | ----- | ----- | ----- | ---- | ---- | ---- |
| max_v | 97.9 | 48.9 | 32.5 | 24.3 | 19.42 | 16.15 | 13.82 | 12.1 | 10.7 | 9.67 |
| drag  | 1.5  | 2    | 5    | 10   | 20    |       |       |      |      |      |
| max_v | 6.34 | 4.7  | 1.76 | 0.78 | 0.29  |       |       |      |      |      |

#### 10.1.3 inertiaTensor：惯性张量

功能说明：用于设置刚体的惯性张量。在距离重心同等的条件下，刚体会向张量值大的一边倾斜。例如，设有如下代码：

```C#
rigidbody.inertiaTensor = new Vector3(5.0f,10.0f,1.0f);
```

起始状态如图10-1所示，则刚体会向y轴所在的方向倾斜翻转，如图10-2所示。

<img src="https://gitee.com/u9king/ImageHostingService/raw/7f57eb279b52cf593f010a94635806cef66f3ab6/Unity/Book/inertiaTensor%E6%83%AF%E6%80%A7%E5%BC%A0%E9%87%8F.png" style="zoom: 80%;" />

代码：

```C#
using UnityEngine;
using System.Collections;
public class inertiaTensor_ts : MonoBehaviour
{
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 160.0f, 45.0f), "x轴惯性张量大于y轴"))
        {
            transform.position = new Vector3(0, 4, 0);
            //transform绕z轴旋转45度
            transform.rotation = Quaternion.Euler(0.0f, 0.0f, 45.0f);
            //设置rigidbody的惯性张量
            //x轴分量值大于y轴，则刚体会向x轴方向倾斜
            rigidbody.inertiaTensor = new Vector3(15.0f, 10.0f, 1.0f);
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 160.0f, 45.0f), "y轴惯性张量大于x轴"))
        {
            transform.position = new Vector3(0, 4, 0);
            transform.rotation = Quaternion.Euler(0.0f, 0.0f, 45.0f);
            //设置rigidbody的惯性张量
            //x轴分量值小于y轴，则刚体会向y轴方向倾斜
            rigidbody.inertiaTensor = new Vector3(5.0f, 10.0f, 1.0f);
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 160.0f, 45.0f), "x轴和y轴惯性张量相同"))
        {
            transform.position = new Vector3(0, 4, 0);
            transform.rotation = Quaternion.Euler(0.0f, 0.0f, 45.0f);
            //设置rigidbody的惯性张量
            //x轴和y轴惯性张量相同，则刚体会保持静止
            rigidbody.inertiaTensor = new Vector3(10.0f, 10.0f, 1.0f);
        }
	}
}
```

定义了3个Button，用来模拟不同惯性张量情况下刚体的运动情况。

#### 10.1.4 mass：刚体质量

```
基本语法 public float mass { get; set; }
```

功能说明：设置或返回刚体的质量。一般刚体质量取值在0.1附近模拟最佳，最大不要超过10，否则容易出现模拟不稳定的情况。

- 对于自由落体运动，物体的速度只与重力加速度Gravity和空气阻力drag有关，与质量mass无关。
- mass的主要作用是在物体发生碰撞时计算碰撞后物体的速度。当一个物体分别去撞击mass大的物体和mass小的物体时，根据动量守恒定律，较重的物体被撞后的速度要慢于较轻的物体。
- 代码

```C#
using UnityEngine;
using System.Collections;
public class Mass_ts : MonoBehaviour
{
    public Rigidbody r1, r2, r3, r4, r5;
    Vector3 v3 = Vector3.zero;
    void Start()
    {
        //r1和r2质量不同，但它们的速度始终相同
        //r4和r5质量不同，当r3以同样的速度撞r4和r5后速度明显不同
        r1.mass = 0.1f;
        r2.mass = 5.0f;
        r3.mass = 2.0f;
        r4.mass = 0.1f;
        r5.mass = 4.0f;
        r3.useGravity = false;
        r4.useGravity = false;
        r5.useGravity = false;
        v3 = r3.position;
    }
    
    void FixedUpdate()
    {
        Debug.Log(Time.time + " R1的速度：" + r1.velocity);
        Debug.Log(Time.time + " R2的速度：" + r2.velocity);
        Debug.Log(Time.time + " R3的速度：" + r3.velocity);
        Debug.Log(Time.time + " R4的速度：" + r4.velocity);
        Debug.Log(Time.time + " R5的速度：" + r5.velocity);
    }
    
    void OnGUI()
    {
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "用R3撞R4"))
        {
            r3.position = v3;
            r3.rotation = Quaternion.identity;
            r3.velocity = new Vector3(4.0f, 0.0f, 0.0f);
        }
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "用R3撞R5"))
        {
            r3.position = v3;
            r3.rotation = Quaternion.identity;
            r3.velocity = new Vector3(0.0f, 0.0f, 4.0f);
        }
    }
}
```

在OnGUI方法中定义了两个不同功能的Button，最后在FixedUpdate方法中打印出每帧各个刚体的移动速度。运行程序发现，刚体r1和r2质量虽然不同，但速度却始终相同；当r3以同样的速度分别去撞r4和r5后，r4和r5的速度明显不同。

#### 10.1.5 velocity：刚体速度

```
基本语法 public Vector3 velocity { get; set; }
```

功能说明：设置或返回刚体的速度值。

- 在脚本中无论是给刚体赋予一个Vector3类型的速度向量v，还是获取当前刚体的速度v，v的方向都是相对世界坐标系而言的。
- velocity的单位是米每秒，而不是帧每秒，其中米是Unity中默认的长度单位。

代码：

```C#
using UnityEngine;
using System.Collections;
public class Velocity_ts : MonoBehaviour 
{
    public Rigidbody r1,r2;
    // Use this for initialization
    void Start () 
    {
        //给父物体r1一个向-z轴的速度，给子物体一个+z轴的速度
        r1.velocity = new Vector3(0.0f,0.0f,-15.0f);
        r2.velocity = new Vector3(0.0f, 0.0f, 10.0f);
    }
    void OnGUI() 
    {
        GUI.Label(new Rect(10.0f,8.0f,300.0f,40.0f),"R1当前速度："+r1.velocity);
        GUI.Label(new Rect(10.0f,58.0f,300.0f,40.0f),"R2当前速度："+r2.velocity);
        Debug.Log("R1当前速度：" + r1.velocity);  //R1当前速度:(0.0,0.0,-15.0)
        Debug.Log("R2当前速度：" + r2.velocity);  //R2当前速度:(0.0,0.0,10.0)
    }
}
```

### 10.2 Rigidbody类实例方法

在Rigidbody类中的实例方法有`AddExplosionForce`、`AddForceAtPosition`、`AddTorque`、`ClosestPointOnBounds`、`GetPointVelocity`、`MovePosition`、`Sleep`、`SweepTest`、`SweepTestAll`和`WakeUp`。

#### 10.2.1 AddExplosionForce：模拟爆炸力

```
基本语法 (1) public void AddExplosionForce(float explosionForce, Vector3 explosionPosition,float explosionRadius);
        (2) public void AddExplosionForce(float explosionForce, Vector3 explosionPosition,
        float explosionRadius, float upwardsModifier);
        (3) public void AddExplosionForce(float explosionForce, Vector3 explosionPosition,
        float explosionRadius, float upwardsModifier, ForceMode mode);
  参数 ———— explosionForce:爆炸点施加的力的大小
            explosionPosition:爆炸点坐标（相对世界坐标系）
            explosionRadius:爆炸作用力有效半径
            upwardsModifier:爆炸力作用点在y轴方向上的偏移
            mode:爆炸力的作用模式，默认为ForceMode.Force。
```

功能说明：用于对刚体添加一个模拟爆炸效果的作用力。设爆炸力大小为F，爆炸点坐标为E，有效半径为R，y轴的偏离量为y_m，刚体A的坐标为P，A受到的爆炸作用力为：A. AddExplosionForce(F,E,R,y_m);

- 爆炸点E作用到A上的力的大小由E点到A表面的最近距离决定。尽管E到刚体A的空间距离为模长|EP|，但爆炸力的作用距离却是E到A的左下角B的距离|EB|，即爆炸力作用到A的大小为：
    $$
    \begin{aligned}
    FA &=  \frac{R - |EB|}{R} × F \\
    \end{aligned}
    $$
    当|EB|>R时，FA=0。由上可知，有可能刚体A发生了移动，但作用到A上的力的大小却没有改变，当刚体从粗实线的位置移动到细实线再移动到虚线的位置时，爆炸点到刚体的有效作用力距离始终是|Ep|。

- 爆炸力作用在A上的方向由爆炸点E、刚体A的位置P和y轴偏移量y_m共同决定。

    - 当y_m为默认值0时，作用力的方向即为从爆炸点到刚体A最近表面的方向。

- 当y_m<0时，作用点向y轴正向移动y_m，从e1点移动到e2点，然后再以e2为爆炸原点求e2到刚体的作用力方向，从e2到B的方向。需要注意的是，无论爆炸点被偏移到什么地方，作用于刚体的力的大小都是由e1点到刚体的最短距离决定的，与偏移后的爆炸点位置无关。

- 当y_m>0时，作用点向y轴负向移动y_m，从e1点移动到e2点，然后再以e2为爆炸原点求e2到刚体的作用力方向，从e2到B的方向。需要注意的是，无论爆炸点被偏移到什么地方，作用于刚体的力的大小都是由e1点到刚体的最短距离决定的，与偏移后的爆炸点位置无关。

综上可知：

- 爆炸力作用在刚体上的力的大小和方向是分开计算的；
- 当爆炸点E固定时，刚体在某个范围移动时受到的爆炸力的大小可能不变；
- 当作用力半径R=0时，所有接受爆炸点E作用的刚体受到的作用力大小都为F。

代码：

```C#
using UnityEngine;
using System.Collections;
public class AddExplosionForce_ts : MonoBehaviour
{
    public Rigidbody A;
    public Transform Z;//Scene视图中显示爆炸点坐标
    Vector3 E = Vector3.zero;//爆炸点坐标
    float F, R, y_m;
    bool is_change = false;
    void Start()
    {
        //初始位置使得爆炸点和A的x、y轴坐标值相等
        //可以更改F及R的大小查看运行时效果
        E = A.position - new Vector3(0.0f, 0.0f, 3.0f);
        F = 40.0f;
        R = 10.0f;
        y_m = 0.0f;
        A.transform.localScale = Vector3.one * 2.0f;
        Z.position = E;
    }
    void FixedUpdate()
    {
        if (is_change)
        {
            A.AddExplosionForce(F, E, R, y_m);
            is_change = false;
    	}
	}
    void OnGUI()
    {
    	//当爆炸点和A的重心的两个坐标轴的值相等时，A将平移不旋转
        if (GUI.Button(new Rect(10.0f, 10.0f, 200.0f, 45.0f), "刚体移动不旋转"))
        {
            is_change = true;
            inits();
        }
        //虽然受力大小不变，但产生扭矩发生旋转
        if (GUI.Button(new Rect(10.0f, 60.0f, 200.0f, 45.0f), "刚体发生移动但受力大小不变"))
        {
            inits();
            A.position += new Vector3(0.5f, -0.5f, 0.0f);
            is_change = true;
        }
        if (GUI.Button(new Rect(10.0f, 110.0f, 200.0f, 45.0f), "按最近表面距离计算力的大小"))
        {
            inits();
            A.position += new Vector3(0.0f, 2.0f, 0.0f);
            is_change = true;
        }
        //y轴的偏移改变了A的原始方向
        //可以更改y_m的值查看不同的效果
        if (GUI.Button(new Rect(10.0f, 160.0f, 200.0f, 45.0f), "y轴发生偏移"))
        {
            inits();
            is_change = true;A.position += new Vector3(0.0f, 2.0f, 0.0f);
            y_m = -2.0f;
        }
    }
    //初始化数据
    void inits()
    {
        A.velocity = Vector3.zero;
        A.angularVelocity = Vector3.zero;
        A.position = E + new Vector3(0.0f, 0.0f, 3.0f);
        A.transform.rotation = Quaternion.identity;
        y_m = 0.0f;
    }
}
```

在OnGUI方法中定义了多个功能不同的Button，用于演示方法AddExplosionForce 在不同条件下的作用情况。

#### 10.2.2 AddForceAtPosition：增加刚体点作用力

```
基本语法 (1) public void AddForceAtPosition(Vector3 force, Vector3 position);
		(2) public void AddForceAtPosition(Vector3 force, Vector3 position, ForceMode mode);
		force为扭矩向量  position为作用点坐标值  mode为力的作用方式
```

功能说明：为position点增加一个力force，其参考坐标系为世界坐标系，作用方式为mode，默认值为ForceMode.Force。此方法与方法AddForce不同，AddForce方法对刚体施加力时不会产生扭矩使物体发生旋转，而AddForceAtPosition方法是在某个坐标点对刚体施加力，这样很可能会产生扭矩使刚体产生旋转。当力的作用点不在刚体重心时，由于作用点的扭矩会使刚体发生旋转。

代码：

```C#
using UnityEngine;
using System.Collections;
public class AddForceAtPosition_ts : MonoBehaviour
{
    public Rigidbody A, B, C;
    Vector3 m_force = new Vector3(0.0f, 0.0f, 10.0f);
    void FixedUpdate()
    {
        //当力的作用点在刚体重心时，刚体不发生旋转
        A.AddForceAtPosition(m_force, A.transform.position, ForceMode.Force);
        //当力的作用点不在刚体重心时，由于作用点的扭矩会使刚体发生旋转
        B.AddForceAtPosition(m_force, B.transform.position + new Vector3(0.0f, 0.3f, 0.0f),
        ForceMode.Force);
        //但是，当力的作用点和刚体重心坐标的差向量与作用力的方向同向时不发生旋转
        C.AddForceAtPosition(m_force, C.transform.position + new Vector3(0.0f, 0.0f, 0.3f),
        ForceMode.Force);
        Debug.Log("A的欧拉角：" + A.transform.eulerAngles); //A的欧拉角：(0.0,30.0,0.0)
        Debug.Log("B的欧拉角：" + B.transform.eulerAngles); 
        //B的欧拉角：(0.0,315.0,0.0)  //B的欧拉角：(0.0,315.0,359.7)
        Debug.Log("C的欧拉角：" + C.transform.eulerAngles); //C的欧拉角：(0.0,50.0,0.0)
    }
}
```

声明了３个Rigidbody变量和一个Vector3变量，然后在方法FixedUpdate中分别对刚体A、B、C施加作用力，最后打印出刚体A、B、C的欧拉角。只有刚体B发生了旋转，刚体A和C的角度未发生变化。



## 11.Time类

## 12.Transform类

## 13.Vector2类

## 14.Vector3类

## 15.游戏实例——坚守阵地

