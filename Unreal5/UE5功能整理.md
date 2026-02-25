# Unreal5功能整理

## 1.场景

### 1.1 天空

#### 1.1.1 无太阳(Sun Brightness)

> 解释:关闭太阳贴图
>
> 组件:BP_Sky_Sphere -> Sun Brightness

#### 1.1.2 多云(Cloud Opacity)

> 解释:控制云朵数量
>
> 组件:BP_Sky_Sphere -> Cloud Opacity

#### 1.1.3 星星(Stars Brightness)

> 解释:控制星星数量
>
> 组件:BP_Sky_Sphere -> Stars Brightness



### 1.2 环境

##### 1.2.1 镜头光晕(Lens Flares)

> 解释:光打进镜头里产生的炫光
>
> 组件:GlobalPostProcessVolume -> Lens Flares

<img src="https://bkimg.cdn.bcebos.com/pic/43a7d933c895d1437c05c01a7af082025aaf0793?x-bce-process=image/format,f_auto/watermark,image_d2F0ZXIvYmFpa2UyNzI,g_7,xp_5,yp_5,P_20/resize,m_lfit,limit_1,h_1080" align="left">



## 2.角色





## 3.游戏性



## 4.引擎

### 3.1 通用

#### 3.1.1 重定向器(Redirectors)

> 解释:当你移动或者重命名资源时UE不会直接断开旧引用，而是在原位置生成一个 Redirector，把旧路径指向新路径。
>
> 组件:右键文件夹

修复方案：复制完成后，右键文件夹点击Fix Up Redirectors in Folder后仍需要手动删除















## 99.疑问整理

1.为什么直射光需要勾选大气光才能正常看见天空颜色?

2.BP_Sky_Sphere为什么要挂载直射光才能产生对应的颜色变化

3.为什么skylight天空光对地面反光的影响这么大？

4.很多素材为什么是单面的，单面的最合理的处理办法是啥？

> 答:游戏引擎默认启用背面剔除Backface Culling,材质开启 Two Sided

5.Map自带的BuildData是个什么数据?（FP1.7）

> 答:关卡烘焙计算后的缓存数据。静态光照（Lightmass 烘焙结果）,光照贴图（Lightmaps）,阴影贴图,反射捕捉数据（Reflection Capture）,预计算可见性（Precomputed Visibility）,预计算体积光照（Volumetric Lightmap）

6.为什么Paragon文件的材质上来是灰的，需要手动修改为XXX_1才能显示颜色,为什么要这么设计?（FP1.7）

> 答:灰色是因为它默认使用的是基础材质实例，真正的颜色信息在对应的 XXX_1材质实例里。Paragon 采用工业级“主材质 + 多实例”架构，默认实例没填贴图，所以显示灰色，必须换成带贴图参数的实例才有颜色。

7.Material中多材料的情况是什么，模型是怎么区分多材质的？

8.组合物体移动的时候没法显示坐标，好蠢啊？





























