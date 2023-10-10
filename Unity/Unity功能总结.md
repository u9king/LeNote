# Unity功能总结

#### 1.点击按钮缩放

```
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

```
using UnityEngine.SceneManagement;
SceneManager.LoadScene("xxx");
```

