---
title: 虚幻引擎UE5.6编译卡在rc2.exe长时间未响应的解决方法
date: 2025-12-30 07:46:35
sticky: 13
pic: /images/thumb/UE5rc2Issue.png
comments: true
toc: true
categories:
- Engine
tags:
- UE5
---

# 虚幻引擎UE5.6编译卡在rc2.exe长时间未响应的解决方法

#### 引言

本文将探讨出现Resource Default.rc2并且长时间未响应的解决办法。我将分享我的解决经验，帮助同样卡在这一步的小伙伴少走弯路，作为虚幻新手小白，文中若有理解偏差，欢迎各位前辈交流指正，与大家共同探讨学习。

#### 错误现象

项目编译过长中卡在Resource Default.rc2并且长时间未响应，甚至在Visual Studio2022中重新编译也会同样卡在这个.rc2上，导致UE项目完全无法打开。本人实测至少30分钟以上，没有进展，并且电脑也会变得异常卡顿。

<img src="/images/UE5/Programming/UE5rc2Issue/1.png">

#### 解决方案

修改你的项目下的Saved\UnrealBuildTool\BuildConfiguration.xml的文件，添加如下内容

<img src="/images/UE5/Programming/UE5rc2Issue/2.png">

```xml
<?xml version="1.0" encoding="utf-8" ?>
<Configuration xmlns="https://www.unrealengine.com/BuildConfiguration">
    <BuildConfiguration>
        <bAllowUBAExecutor>false</bAllowUBAExecutor>
    </BuildConfiguration>
</Configuration>
```

#### 知识储备

在提及探寻过程前，先铺垫两个关键知识。
1.电脑是同时存在至少两个BuildConfiguration.xml，引擎的和项目的。并且项目的优先级会大于引擎，也就是说如果项目内存在，项目的这个是最终执行的。貌似是右键.uproject生成VS项目文件的时候生成的。

```
引擎默认启用：%APPDATA%\Unreal Engine\UnrealBuildTool\BuildConfiguration.xml
项目默认启用：你的项目\Saved\UnrealBuildTool\BuildConfiguration.xml
优先级：项目>引擎
```

2.虚幻每个版本都有自己对应的编辑器，特别是MSVC和WindowsSDK，如果与虚幻版本不相同时常会出现意料之外的错误

[UE5.x对应VisualStudio的编译版本网址](https://dev.epicgames.com/documentation/zh-cn/unreal-engine/setting-up-visual-studio-development-environment-for-cplusplus-projects-in-unreal-engine)

#### 探寻过程

以下是我曾尝试过但无效的方案（多为gpt和gemini建议），为了避免大家少走弯路，我记录一下结果。各位也可以亲自尝试一下，毕竟错误原因各不相同。

1.修改%APPDATA%\Unreal Engine\UnrealBuildTool\BuildConfiguration.xml为以下内容,这是我最先看到的方案，没想到是最接近结果的方案

```xml
<?xml version="1.0" encoding="utf-8" ?>
<Configuration xmlns="https://www.unrealengine.com/BuildConfiguration">
    <BuildConfiguration>
        <bAllowUBAExecutor>false</bAllowUBAExecutor>
    </BuildConfiguration>
</Configuration>
```

结果:尝试并重启无效，猜测是UE5引擎有对这个文件的缓存（GPT告诉我的），暂不确定。



2.指定编译器的版本WindowsSdkVersion为10.0.22621.0对于UE5.6，这个方案同样也是两个地方一个%APPDATA%下的BuildConfiguration.xml还有一个项目的Saved\UnrealBuildTool下的BuildConfiguration.xml我都尝试过

```xml
<?xml version="1.0" encoding="utf-8" ?>
<Configuration xmlns="https://www.unrealengine.com/BuildConfiguration">
    <WindowsPlatform>
    	<WindowsSdkVersion>10.0.22621.0</WindowsSdkVersion>
	</WindowsPlatform>
</Configuration>
```

结果:尝试并重启无效，由于我是多个虚幻版本同时存在UE5.0,UE5.3,UE5.6所以修改%APPDATA%通常不是一个特别好的选择，并且如果你的开在Configuration中间是空白，按照虚幻的说法它会自动找最适合自己版本的编译器版本

3.重新手动构建项目，也就是删除除了Content,Config和Source还有.uproject外的所有文件，并且右键.uproject先Generate再双击打开项目Rebuild。

<img src="/images/UE5/Programming/UE5rc2Issue/3.png">



4.修改项目的Config/DefaultEngine.ini文件,添加以下内容

<img src="/images/UE5/Programming/UE5rc2Issue/4.png">

```
[BuildSettings] 
bUseUnrealBuildAccelerator=false 
bUseUBTMakefiles=false 
bAllowFASTBuild=false
```

结果：尝试无效，我其实觉得这个是有可能正确的，奈何水平不够，暂定为无效。

5.添加引擎路径和项目路径到Windows的杀毒软件白名单中，Windows安全中心->病毒与防护->文件夹限制访问->添加引擎和项目路劲
结果:尝试无效，我听闻rc2.exe是会与杀毒软件有关导致卡死，但是我这边实测没有效果

6.验证虚幻引擎完整性和Visual Studio2022的完整性
结果：尝试无效，这个非常痛苦且无效，虚幻引擎还好，VS2022验证完整性等于完全重装，好多设置都没有了，不推荐并且无效



结果：尝试并重启且无效，相信我，我尝试了不下20遍这个操作，这是所有问题的万能(无效)解法，并不是你编译的次数多问题就能解决，虽然他们都会告诉你应该这么做。

7.物理阻断UBA执行文件，找到引擎目录下的将UbaAgent.exe修改为UbaAgent.exe.bak
结果：尝试无效

8.在VS的解决方案资源管理器里，右键点击你的项目 -> 属性 -> NMake ->生成命令行。在现有的命令后面手动加上 -NoUBA
结果：尝试无效



感谢下面的文章对我提供的帮助，如果你没有找到解决办法不妨可以参考以下链接
1.https://forums.unrealengine.com/t/ue-5-6-project-creation-fails-on-default-rc2-rc-exe-exit-code-1/2595642/4
2.https://forums.unrealengine.com/t/ue-5-6-project-stuck-at-rc2-exe-when-building/2687891/3
3.https://gamedev.tv/courses/unreal-c-dev/lecture/7985/thread/34317
4.https://www.bilibili.com/video/BV1nnQWYME5T/?spm_id_from=333.1391.0.0&vd_source=9a146b8fa39d5ea05ce3a524dcff45d4
5.https://dev.epicgames.com/documentation/zh-cn/unreal-engine/setting-up-visual-studio-development-environment-for-cplusplus-projects-in-unreal-engine