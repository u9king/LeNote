---
title: UnityHub Validation Failed下载编辑器错误，添加模块报错的解决方案
date: 2025-08-10 07:46:35
pic: /images/thumb/ValidationFailed.png
comments: true
toc: true
categories:
- Programming
tags:
- Unity
---

# UnityHub ValidationFailed下载编辑器错误，添加模块报错的解决方案

#### 引言

本文将探讨出现Validation Failed这个错误的原因，以及可能的解决方案。如果接受下载离线版本安装包然后导入UnityHub可以忽略本文内容。

#### 错误现象

UnityHub无法下载unity6，无法安装Android Build Support,WebGL等几乎所有模块。但是，unity2022及以下的版本却可以正常下载。

<img src="/images/Unity/Programming/ValidationFailed/1.png">

#### 解决方案

将代理工具设置为代理所有流量的TUN模式(叫法可能有差别)。如果是这种情况单纯切换全局不行，一定要开启代理所有流量，建议选择一个海外地址，不然一定会被重定向回来。

#### 探寻过程

我尝试过网上几乎所有的方法，类似用管理员权限启用，修改防火墙，确保路径没有中文，删除缓存，删除注册列表，删除Hub下载旧版本，修改默认下载位置等都不起作用。

其中有一篇文章对我产生了启发作用，我们可以查看日志，将日志内容喂给AI来帮我们解读错误。纵使我编码能力有限，好在日志文件并不难找到就在UnityHub的故障排除—>打开日志文件夹，也可以C:\Users\<你的用户名>\AppData\Roaming\UnityHub\logs找到info-log.json

<img src="/images/Unity/Programming/ValidationFailed/2.png">

找到这个文件之后info-log.json，我好像不能直接打开，我是复制出来进行查看的，我将最后几段代码贴在下方

```json
{"timestamp":"2025-07-30T00:31:02.645Z","level":"info","moduleName":"Download Item b1178ac1-5ae1-4d2d-a117-ffdbdfa7be2a","pid":8856,"message":"Transition to state download_validation"}
{"timestamp":"2025-07-30T00:31:02.646Z","level":"info","moduleName":"Disk Validation Strategy","pid":8856,"message":"Executing the Download Validation Strategy..."}
{"timestamp":"2025-07-30T00:31:02.656Z","level":"info","moduleName":"Disk Validation Strategy","pid":8856,"message":"Validating destination path permission for access..."}
{"timestamp":"2025-07-30T00:31:02.658Z","level":"info","moduleName":"Disk Validation Strategy","pid":8856,"message":"Destination Path check: Passed"}
{"timestamp":"2025-07-30T00:31:02.658Z","level":"info","moduleName":"Disk Validation Strategy","pid":8856,"message":"Validating source path (https://download.unity3d.com/download_unity/157d81624ddf/TargetSupportInstaller/UnitySetup-WebGL-Support-for-Editor-6000.0.40f1.exe) availability..."}
{"timestamp":"2025-07-30T00:31:03.387Z","level":"error","moduleName":"Disk Validation Strategy","pid":8856,"message":"Error occured in 'Source Availablity Check'. AxiosError: Request failed with status code 404"}
{"timestamp":"2025-07-30T00:31:03.387Z","level":"info","moduleName":"Disk Validation Strategy","pid":8856,"message":"Source Availability check: Failed"}
{"timestamp":"2025-07-30T00:31:03.388Z","level":"info","moduleName":"Download Item","pid":8856,"message":"Download: b1178ac1-5ae1-4d2d-a117-ffdbdfa7be2a, Exiting from State: Validation, Event: ERROR"}
{"timestamp":"2025-07-30T00:31:03.388Z","level":"info","moduleName":"Download Item","pid":8856,"message":"Download: b1178ac1-5ae1-4d2d-a117-ffdbdfa7be2a, Current State: Failed, Event: ERROR"}
{"timestamp":"2025-07-30T00:31:03.388Z","level":"info","moduleName":"Download Item b1178ac1-5ae1-4d2d-a117-ffdbdfa7be2a","pid":8856,"message":"Transition to state download_failed"}

```

可以很明显看到有一段404Error的代码出现在这里

```json
{"timestamp":"2025-07-30T00:31:02.658Z","level":"info","moduleName":"Disk Validation Strategy","pid":8856,"message":"Validating source path (https://download.unity3d.com/download_unity/157d81624ddf/TargetSupportInstaller/UnitySetup-WebGL-Support-for-Editor-6000.0.40f1.exe) availability..."}
{"timestamp":"2025-07-30T00:31:03.387Z","level":"error","moduleName":"Disk Validation Strategy","pid":8856,"message":"Error occured in 'Source Availablity Check'. AxiosError: Request failed with status code 404"}
```

用人工智能解读这块代码说Unity Hub在校验下载链接时，向以下 URL 发起请求：https://download.unity3d.com/download_unity/157d81624ddf/TargetSupportInstaller/UnitySetup-WebGL-Support-for-Editor-6000.0.40f1.exe，但是返回了 404（找不到资源）。随即我就把这段代码粘到浏览器上了，发现是能够正确下载的可为什么会出现错误呢？
然后给了我一个方法让我去模拟UnityHub整个网络服务请求过程，打开电脑的cmd(在任意位置都行)粘贴这段代码

```
curl -I https://download.unity3d.com/download_unity/157d81624ddf/TargetSupportInstaller/UnitySetup-WebGL-Support-for-Editor-6000.0.40f1.exe
```

<img src="/images/Unity/Programming/ValidationFailed/3.png">

这就非常有意思了，重点关注Location这一行，并不是404而是被重定向为302也就是说我的网址被重定向为：
你访问的地址 https://download.unity3d.com/… 被重定向到了 https://download.unitychina.cn/…
那我大致明白了原因，是因为重定向导致的。Unity中国是从2022开始构建之后没有提供Unity6以及后续的所有安装包导致的。那么问题的解法也大致明白了。

随即我打开自己的代理工具，将连接国际网络的工具切换成最高级别的代理所有流量TUN模式（根据每个代理工具的同步，叫法可能有差别，但是我试过切换全局也不行，一定要开启代理所有流量，选择一个海外地址，不然一定会被重定向回来）。

<img src="/images/Unity/Programming/ValidationFailed/4.png">

至此我的所有Hub组件都能正确下载，花了我将近一天的时间。我还在尝试一些类似Fiddler修改包头的方法，奈何水平不够还没有成功。但是使用完全代理的工具基本可以解决我的问题

<img src="/images/Unity/Programming/ValidationFailed/5.png">

感谢下面的文章对我提供的帮助，如果你没有找到解决办法不妨可以参考以下链接
1.https://blog.csdn.net/weixin_50681225/article/details/141875781
2.https://blog.csdn.net/hsy1842853303/article/details/147857422
3.https://discussions.unity.com/t/unity-hub-amdroid-module-validaition-failed/932130/3
4.https://blog.csdn.net/IT_Lightning/article/details/140692220?spm=1001.2101.3001.6650.5
5.https://developer.unity.cn/projects/63ad4936edbc2a1fd2cffc0d
6.https://zhuanlan.zhihu.com/p/28213789169

