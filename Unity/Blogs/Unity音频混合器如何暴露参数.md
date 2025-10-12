---
title: Unity音频混合器如何暴露参数
date: 2025-03-20 14:07:35
pic: /images/thumb/AudioMixExposed.png
comments: true
toc: true
categories:
- Engine
tags:
- Unity
---

# Unity音频混合器如何暴露参数

音频混合器是Unity推荐管理音效混音的工具，那么如何使用代码对它进行管理呢？
首先我在AudioMixer的Master组中创建了BGM和SFX的分组，你也可以直接用Master没有问题。
这里我以BGM为例，如果要在代码中进行使用就需要将参数暴露出去，问题是这个按钮并不在Exposed Parameters的位置，也没有任何+按钮与之关联就很恼火。

<img src="/images/Unity/Programming/AudioMixExposed/1.png">

如果要将BGM的音量暴露出去就需要在inspector窗口，右键Volumn(体积)这个变量名(不是右边的滑轨),然后选择弹出框的第一条Expose “Volume(of BGM)” to script即可

<img src="/images/Unity/Programming/AudioMixExposed/2.png">

在Exposed Parameters中就能看到MyExposedParam(Volumn of BGM)就是BGM的音量

<img src="/images/Unity/Programming/AudioMixExposed/3.png">

可以双击这个变量来自定义暴露的变量名，这里我就叫BGM了

<img src="/images/Unity/Programming/AudioMixExposed/4.png">

最后让我们来到代码中进行使用

```
public AudioMixer audioMix;  //挂载音频混合器

public void SetAudioMix(float percentage)
{
    //int step = 10;  //步长，自定义可以省略
    //float maxVolumn = 20f;  //最大音量，对应mix上定义的最大20db，也可以自定义
    //float minVolumn = -80f;  //最小音量，对应mix上定义的最小-80db，也可以自定义
    //currentBGMVolumn = minVolumn + Mathf.Log(1f + percentage * (step - 1), step) * (maxVolumn - minVolumn);  //对数对应
    audioMix.SetFloat("BGM", 20f);//设置音频混合器音量,这里的BGM就是我们对外暴露所设置的变量名
}
```


