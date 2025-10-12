---
title: UnityWebGL打包后本地启动方法
date: 2025-08-10 08:24:35
pic: /images/thumb/WebGLServer.png
comments: true
toc: true
categories:
- Engine
tags:
- Unity
---

# UnityWebGL打包后本地启动方法

#### 引言

常见WebGL开启方法常需要重新打包点击`Build and Run`或者将游戏放到Unity的云服务器上，作为开发者而言这两个方案一个为了开启再次打包，另一个直接放到了公开环境都不太合适。所以我们需要一个能在本地开启测试的WebGL的方法。

#### 解决方案

使用的是nodejs中的一个轻量级命令行工具http-server只有100kB左右。熟悉npm的可以在cmd中用命令`npm install -g http-server`安装，然后用命令`http-server -p 8000`打开服务器Url为http://127.0.0.1:8000。http-server的github地址:https://github.com/http-party/http-server，熟悉github的可以自行翻阅。

### 手把手教学

下面我将从以下几个步骤来阐述WebGL本地开启的方法，
1.找到你所在的打包WebGL的位置，能看到3个游戏相关的文件，如果没有还请移步Unity的WebGL先行打包，确定存在这三个文件的当前目录下在这里打开cmd窗口

<img src="/images/Unity/Programming/WebGLServer/1.png">

2.输入cmd按下回车，直接在当前目录下打开Windows命令处理程序窗口

<img src="/images/Unity/Programming/WebGLServer/2.png">

3.检查node是否存在,node版本没有太大的关系用LTS都可以，下载地址:https://nodejs.org/zh-cn/download/，需要配置一下node的环境，具体方法可以在网上搜索(参考:https://blog.csdn.net/web15185420056/article/details/144251082)，相信我如果会接触程序，配置node环境是这辈子最简单的事情之一。

<img src="/images/Unity/Programming/WebGLServer/3.png">

4.在确保node已经正常运行的情况下安装http-server
`npm install -g http-server`简单的一行代码就能搞定，能看到added没报错就已经成功了，这里我又添加了一行http-server -v查看当前版本，确保安装成功

<img src="/images/Unity/Programming/WebGLServer/4.png">

5.当前所有条件就准备齐全了，下面就可以打开自己的webgl项目了，同样的在WebGL文件所在目录下打开Windows命令处理程序窗口，输入`http-server -p 8000`，可以看到Available on:
http://192.168.1.107:8000
http://127.0.0.1:8000这两个就是本地网络地址
同时为大家解释一下-p 8000就是端口号8000，基本上你用什么端口号都行就是别用有特殊含义的端口号就行，什么80,443,22,25,1这些别用，我们常用的就是8000，或者8080
<img src="/images/Unity/Programming/WebGLServer/5.png">

此时将这个http://127.0.0.1:8000粘贴到自己浏览器中就可以，切记你看最后一句hit Ctrl-c是关闭服务器，所有喜欢用CtrlC，CtrlV粘贴的小伙伴要当心。

<img src="/images/Unity/Programming/WebGLServer/6.png">

此时就能够运行自己的WebGL游戏了在浏览器上。

最后给大家解释一下，按需获取
1.本地服务器端口
http://127.0.0.1:8000这个地址端口8000对应的就是你电脑上正在运行的本地服务器（Windows里就是打开的命令行窗口启动的程序）。
如果你关掉那个命令行窗口，服务器就停止了，网页就打不开了。
重新开一个命令行窗口，重复启动命令，就能新建一个本地服务器。
2.127.0.0.1是本地回环地址，是你的电脑自己用的地址，只能在你自己电脑上访问。别人用别的电脑访问不了这个地址，确保你本地测试安全又方便。