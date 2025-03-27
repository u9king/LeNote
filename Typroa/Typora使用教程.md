# Typora使用教程

#### 1.标题

```
#		一级标题
##		二级标题
###		三级标题
####	四级标题
#####	五级标题
######	六级标题
```

#### 2.字体

```
**加粗**
*斜体*
***斜体加粗***
~~删除线~~~
$\textcolor{颜色}{内容}$	字体颜色
<span style="color:颜色;">内容</span>
```

#### 3.表格

快捷键：ctrl+T

```
合并表单元格
<table>
    <tr>
        <th>表头1</th>
        <th>表头2</th>
    </tr>
    <tr>
        <td>行1，列1</td>
        <td rowspan="2">合并单元格</td>
    </tr>
    <tr>
        <td>行2，列1</td>
    </tr>
</table>
```

<table>
    <tr>
        <th>表头1</th>
        <th>表头2</th>
    </tr>
    <tr>
        <td>行1，列1</td>
        <td rowspan="2">合并单元格</td>
    </tr>
    <tr>
        <td>行2，列1</td>
    </tr>
</table>
#### 4.列表

```
快捷键
有序列表: 1.空格
无序列表：-空格
无序二级列表：-空格-空格
```

#### 5.引用

```
快捷键: > 空格
```

> a a1 a2
>
> b b1 b2

#### 6.排版

```
`<center>文字</center>`	居中
```

#### 7.图片改变大小，位置

<img src="xxx" style="zoom:50%" align="left">  或者直接右键就能缩放

<img src="xxx" style="width:100vw; height: 20vw;">

其中vw是Viewport Width的缩写，100vw让图片宽度等于整个视口宽度，这样图片就能够根据视窗调整大小

参考:https://www.typora.net/460.html

#### 8.公式

```
公式左对齐
使用LaTex中的内联公式
```

例如：
$$
\lim_{x \to 0} \frac{\sin x}{x} = 1
$$
重要极限2:	$$\lim_{x \to \infty}({1+\frac{1}{x}})^x = e$$

#### 9.Emoji 表情的语法

:smile::smiling_imp::raising_hand::sunny::earth_asia::santa::blue_book::bicyclist::house::arrows_counterclockwise::pisces:

```
:单词:
```

详细参见:https://doc.fly2you.cn/en-US/study/markdown/mdemoji.html

#### 10.下划线

---

```
--- + enter
```



#### 疑问

1.没有表头的表



#### 参考

1.https://zhuanlan.zhihu.com/p/261750408