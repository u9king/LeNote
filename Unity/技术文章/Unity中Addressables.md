# Unity中Addressables

前置内容:

- Assets Bundle



API接口 挂载

```
using UnityEngine.AddressableAssets;

//AssetReference	通用资源标识类可以用来加载任意类型资源
//AssetReferenceAtlasedSprite 图集资源标识类
//AssetReferenceGameObject 游戏对象资源标识类
//AssetReferenceSprite 精灵图片资源标识类
//AssetReferenceTexture 贴图资源标识类
//AssetReferenceTexture2D
//AssetReferenceTexture3D
//AssetReferenceT<> 指定类型标识类
 

//通过不同类型标识类对象的申明 我们可以在Inspector窗口中筛选关联的资源对象
```



API接口 加载

```
using UnityEngine.ResourceManagement.AsyncOperations;  //异步加载

AsyncOperationHandle<GameObject> handle =assetReference.LoadAssetAsync<Game0bject>();

handle.Completed += Handle_Completed;

private void Handle_Completed(AsyncOperationHandle<GameObject> obj){
	if(handle.Status == AsyncOperationStatus.Succeeded){
		Instantiate(handle.Result);
	}
}
```



标签加载

```
[AssetReferenceUILabelRestriction("Xx")]
public AssetReference assetReference;
```



tips：

1.Resources文件夹下资源如果变为可寻址资源，会移入Resources_moved文件夹中

2.



可寻址包构建在基于Unity的资源包系统和提供一个用户接口来管理你的资源包。当你做制作一个可寻址的资产，你可以在任何地方用这个资产地址加载。可寻址系统定位并且返回资产，无论它是在本地应用程序可用还是存储在远程的网络上。

你可以用这个包来远程发布内容

