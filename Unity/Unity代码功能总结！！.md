# Unity代码功能总结

#### 1.点击按钮缩放

```C#
private void OnMouseDown()
{
    transform.localScale = transform.localScale * 0.8f;
}

private void OnMouseUp()
{
    transform.localScale = transform.localScale / 0.8f;
}
```

#### 2.切换场景

```C#
using UnityEngine.SceneManagement;
SceneManager.LoadScene("xxx");
```

#### 3.相机跟随

```C#
public class CamerFollow1 : MonoBehaviour
{
    //跟随的游戏人物
    private Transform Player;
    //相机与人物之间的距离
    Vector3 offset;
    // Start is called before the first frame update
    void Start()
    {
        Player = GameObject.Find("JohnLemon").transform;
        offset = transform.position - Player.position;
    }

    // Update is called once per frame
    void Update()
    {
        transform.position = offset + Player.position;
    }
}
```

#### 4.鼠标点击跟随

```C#
private void Update()
{
        transform.position = Camera.main.ScreenToWorldPoint(Input.mousePosition);
        transform.position -= new Vector3(0,0,Camera.main.transform.position.z);
}
```

#### 5.获取当前物体名称

```C#
gameObject.name
```

#### 6.游戏对象旋转

```c#
//关闭z轴旋转
private void Start()
{
    //开启旋转
    rg.freezeRotation = false;
}
rg.freezeRotation = false;
```

