# Unity中Timeline文档翻译

------

#### 目录

关于时间线——About Timeline

时间线概述——Timeline overview

使用时间线窗口——Using the Timeline window

​		创建时间线资源和时间线实例——Creating a Timeline Asset and Timeline instance

​		使用无限片段记录基础动画——Recording basic animation with an infinite clip

​		将无限片段转化为动画片段——Converting an Infinite clip to an Animation clip

​		制作人形机器人的动画——Animating a humanoid

​		用一个动画重写轨道和头像遮罩——Using an Animation Override track and Avatar masking

​		嵌套时间线实例——Nesting Timeline instances

时间线窗口——Timeline window

​		时间线预览和时间线选择器——Timeline Preview and Timeline Selector

​		时间线播放控件——Timeline Playback Controls

​		轨道列表和轨道标题——Track List and Track Headers

​				添加轨道——Adding Tracks

​				选择轨道——Selecting Tracks

​				重复轨道——Duplicateing Tracks

​				删除轨道——Deleting Tracks

​				锁定轨道——Locking Tracks

​				静音轨道——Muting Tracks

​				重排轨道和渲染优先级——Reordering Tracks and Rendering Priority

​		使用轨道组——Using Track Groups

​				折叠和展开轨道组——Collapsing and Expanding Track Groups

​				锁定轨道组——Locking Track groups

​		片段编辑模式和片段视图——Clip Edit modes and the Clips view

​				平移和缩放片段视图——Panning and Zooming the Clips View

​				添加片段——Adding Clips

​				插入片段——Inserting Clips

​				选择片段——Selecting Clips

​				定位片段——Positioning Clips

​				平铺片段——Tiling Clips

​				重复片段——Duplicating Clips

​				剪辑片段——Trimming Clips

​				分离片段——Splitting Clips

​				重置片段——Resetting Clips

​				改变片段播放速度——Changing Clip Play Speed

​				设置间隙外推法——Setting Gap Extrapolation

​				片段的淡入淡出——Easing-in and Easing-out Clips

​				绑定片段——Blending Clips

​				匹配片段偏移——Matching clip offsets

​		曲线视图——Curves View

​		时间线设置——Timeline Settings

时间线检视器——Timeline Inspector

可播放导演组件——Playable Director Component

示例——Samples

自定义USS样式——Defining custom USS styles

时间线词汇表——Timeline Glossary

------

#### 行为设计——Behaviour Designer

------

#### 概述——Overview

​	Behaviour Designer插件是一个行为树设计给每个人—编程开发者，艺术家，设计者。Behaviour Designer插件提供一个强大的API能让你非常容易地创建新任务。它通过大量的第三方整合提供直观可视化编辑器使其能够不写一行代码也能够完成复杂AI的创建工作。

​	这本手册是要给你一个Behaviour Designer全方面完整的概览。如果你只是刚刚开始使用行为树，我们提供了一套叫做Behavior Tree Basics的系列视频。这一页也能够快速概览行为树。用Behaviour Designer你不需要知道行为树的底层实现，但还是建议你知道一些核心概念像是任务类型（行为，混合，条件和装饰器）。

​	当你第一次打开Behaviour Designer你将会看到以下界面窗口：

​	这里有Behaviour Designer的四个部分。从下面的截图中，第一部分是图形区域。这是将会是你创建行为树的地方。第二部分是一个属性面板。属性面板将会是你编辑行为树的具体属性，增加新的任务，创建新的变量或者编辑任务参数的地方。第三部分是一个行为树操作工具栏。你可以用下拉框来选择已有的行为树或者增加/减少行为树。最后一个部分，第四部分，是一个调试工具栏。你可以通过这个面板开始/停止，步进，或者暂停Unity。此外，在执行行为树之前，你将会看见行为树所抛出的错误的数量。——4

---

​	第一部分是Behaviour Designer主要的部分，你将会在这里面工作。在这个部分中，你可以创建新任务并且安排那些任务到行为树中。在开始之前，你首先需要添加行为树组件。在你刚刚开始创建的时候，行为树组件将会扮演行为树的管理者。你可以通过右键点击图形区域，紧接着选择"Add Beahvior Tree"；或者通过点击第三部分中“Lock”边上的加号按钮来创建一个新的行为树组件。

​	一旦行为树被添加好后，你就可以开始添加任务。增加任务通过右键点击图形区域，或者在第二部分中的属性面板中点击在Tasks标签窗口。新的任务也可以通过按下空格键并且打开快速搜索窗口来添加：

​	一旦一个任务被添加，你将会看到以下内容	——5

---

​	除了你添加的任务外，入口任务(entry task)还将被(自动地)添加进来。入口任务扮演着树的根节点的角色。入口任务的唯一目的就是扮演根节点。序列任务(sequence task)有一个错误因为它没有孩子(叶子节点)。一旦你添加了一个孩子节点，错误就会消除。现在我们已经添加了我们的第一个任务，让我们再增加更多的一些内容：

​	你可以连接序列任务和选择任务通过抓取sequence(序列任务)的底部到selector(选择任务)的顶部。对其余的任务重复这一过程。如果你遇到错误，你可以选择一个连接线并且通过delete键删除它。你   ——6

---

可以通过点击某个任务或者拖动它来重新安排这些任务。

​	Behaviour Designer将会按照深度优先算法来执行任务。你通过拖动它们到兄弟节点(同级任务)的左侧或者右侧来改变任务的执行顺序。从上面的截图来看，任务将会被按照以下顺序来执行：

​	SequenceA->SelectorA->SequenceB->ActionA->ActionB->ActionC->SelectorB->ActionD->ActionE

​	现在我们已经创建了一个基础的行为树，让我们一起来修改其中一个任务的参数。选择ActionC这个节点来显示属性面板中的检查器。正如你所见，我们可以在这里给任务重命名，设置任务是否即刻生效，或者进入任务的评论区。除此之外，我们可以修改在task类中存在的所有公共变量。这将会包括分配在Behaviour Designer中已经被创建的变量。在我们的例子中，唯一公共变量是旋转速度(Rotation Speed)。我们在属性中设置这个值将会被用在行为树里。

​	在属性面板中有另外三个标签页分别是Variables, Tasks和Behavior。变量面板允许你去创建在任务之间共享的变量。想要获取更多的信息，可以去看一看Variables topic这一章的内容。这个任务面板列举了所有你可以能用到的任务来提供给你使用。通过右键点击并且添加任务得到的也是这个列表。这个列表是通过搜索动作(action),复合(composite), 条件(conditional)或者装饰(decorator)任务类及其其他任何派生来创建的。最后一个面板(标签页)，行为树面板，展示你一次创建行为树通过行为树组件的检查器。有关每个选项的作用的更多详细信息，请参考Behavior Component Overview page这一章的内容。——7

---



