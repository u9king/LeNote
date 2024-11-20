# Unity中DialogueSystem插件文档翻译

## Unity的对话系统

### 欢迎来到Unity的对话系统！

对话系统是Unity的一个分支对话系统。它不需要脚本即可编辑并且设计得方便程序员进行拓展。

Unity的对话系统会使你：

- 用逻辑分支编写交互式对话
- 展示对话用一个灵活，完整的定制化UI系统
- 定义问题/人物给玩家来完成
- 播放过长动画序列，例如相机片段和音频
- 保存和加载游戏，依据场景变化保存游戏数据
- 用本地化来用不同语言展示文本和播放动画序列
- 导入和导出从不同格式例如：artity:草稿、聊天映射器和CSV格式

### 在手册中

在本手册中包含下列内容：

- 开始——安装，升级，样例和快速开始
- 对话编辑器——如何编写内容
- 对话管理器——如何设置主要对话系统游戏对象
- 触发和交互——如何设置游戏内的交互
- 对话UIs——如何设置对话UIs
- 逻辑和Lua——如何定义条件和用变量工作
- 过场序列(Cutscene)–如何设计过场序列
- 问题——如何编写和执行问题
- 保存系统——如何保存和加载数据
- 本地化——如何添加语言本地化
- 脚本化——编程参考
- 导入和导出——如何导入和导出不同格式
- 第三方融合——如何用对话系统与别的资产相结合
- 教程——手把手知道一个通用任务（所有视频教程）
- 组件参考
- 快问快答
- 版本说明

### 如何获得帮助

我们在这里提供帮助！如何你被困扰或者有很多问题，无论何时请我们来获得支持在pixelcrushers.com或者访问 https://pixelcrushers.com.我们尽全力24小时回复邮件。如果你没有收到回复在24小时之内，请检查您的垃圾邮件。

## 开始 — Getting Started

### 安装(Setup)

提示：如果要从1.x版本升级，请先阅读从1.x升级中的重要指示。

Unity的对话系统需要Unity2018.4以上版本。再开始前，打开包管理器窗口接着导入Unity的对话系统。将会有以下文件夹添加到你的项目中：

- 插件(Plugins)→像素破碎机(Pixel Crushers)→对话系统(Dialogue System):这就是对话系统
- 插件(Plugins)→像素破碎机(Pixel Crushers)→通用(Common):像素破碎机产品的核心脚本

2D支持：如果你的项目用的是2D物理，参见使用2D物理支持。

#### 从1.X版本升级

如果你不需要从1.X版本升级，跳过这个部分到样例场景即可。（点击这里浏览视频版本）

如果你正在从1.X版本升级，你必须要根据以下步骤来执行：

1.备份你的项目

2.如果你的项目用的是Unity5.x，你必须把先把强制文本序列化（Force Text serialization first）

- 选择菜单下的编辑Edit→项目设置Project Settings→编辑者

- 接着改变项目序列化Asset Serialization到强制文本Force Text和版本控制到可见Meta文件

  1.如果你的项目包含在Unity4.x创建的资源，请阅读下面的注释

  2.删除完整的对话系统文件夹。（如果你正在使用Unity5.x，不要做这些知道你已经转换到强制文本序列化）

  - 如果你已经将任何第三方支持文件夹移出了 Dialogue System 文件夹，也请将它们删除。
  - 确保你的项目依然能够变异。如果你有一些额外的脚本是基于对话系统的，你应该把它们导出作为临时的Unity包，删除他们，接着在重新完成升级之后把他们导回来。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/forceText.png)

  3.从资源商店中导入对话系统给Unity2.x

  4.运行工具→像素破碎者→对话系统→工具→运行1x到2x升级

  注意：如果项目中有任何编译错误，将无法运行更新程序。解决编译错误后，菜单项应会变为可用。

  如果你有任何关于升级过程的问题，请联系我们。
#### 关于 Unity 4.x 资源的注意事项

  对话系统内部脚本参考(GUIDs:Globally Unique Identifiers也叫全局唯一标识符)变化从1.x到2.x.更新工具定位具有过时1.xGUID的资产，并将他们重定向到新的2.xGUID。Unity可以存储资源，泪如场景文件和对话数据库，通过二进制或者文本的形式。这个过程叫做序列化。升级工具必须临时转换资产到文本系列化来升级对话系统的GUIDs。如果你的项目包含资产是通过Unity4进行的二进制序列化，Unity并不会将他们转换成文本序列化。这不仅限于对话系统；这是Unity的32位与64位的限制。在这种情况下，你必须在Unity4中打开他们，将他们转换成文本序列化，导出作为资产，并且在更新之前导入这些资产到你的项目中。

#### 更新注意事项

2.x版本做了如下改动：

- Unity的对话系统转移到了Plugins文件夹中。
- 源代码现在直接在项目中，而不再作为预编译DLL的Unity包文件。
- 旧版的Unity GUI和UnityUI对话界面子界面已被弃用，取而代之的是更强大且灵活的标准UI子系统，旧系统依然可用，但不再显示在组件菜单或检视器面板的添加组件按钮中。
- 触发组件，例如对话除法器和任务除法器，已合并为一个更简单、更优化的对话系统除法器。旧组件依然受支持，但推荐使用的触发器组件是现在对话系统的触发器。

### 下一步(Next Steps)

- 演示场景
- 概述
- 快速开始

### 演示场景(Demo Scene)

演示场景保存在插件→像素破碎者→对话系统→演示。运行它就能看见对话系统是如何工作的了。

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/demoScene.png)

URP/HDRP注释：如果你的项目用了URP或者HDRP，在运行对话系统职前下载并且导入相关的材料包。

#### 如何运行(How to play)

- 用方向键或者WASD键可以移动
- 用鼠标观察场景
- 观察可交互的物体并且用空格键来与之交互
- 鼠标按下左键来进行射击
- 按下Escape键来打开暂停菜单

注释：如果您在现有的项目中使用对话系统，并且玩家穿过地板掉落，可能是因为演示项目使用了Unity默认物理设置。

#### 可以尝试的事项(Things To Try)

你可以在场景中做下面这几件事情

- 获取发射代码(任务):与私人哈特交谈来获取任务。这涉及到找到下一间房间的密码以及和电脑交互。
- 敌人攻击(任务):在下一个房间与NPC对话来获取任务。这展示了如何设置基于击杀的任务。
- 切换场景:与门交互来切换场景。这展示了如何在切换场景时保持数据。
- 保存&加载游戏：使用暂停菜单来保存游戏。

下一个部分展示一个完整的对话系统工作。如果你想要直接跳转到使用你自己的内容，你可以直接跳到快速开始的部分。

### 概述 (Overview)

对话系统包括下面这些部分:

- 对话数据库(Dialogue Database)：包括你的对话，任务和变量等。
- 对话编辑器(Dialogue Editor)：编辑对话数据库。你也可以导入和导出第三方格式。
- 对话管理器(Dialogue Manager)：在游戏运行时管理对话系统。提供了一个预制体Prefab。
- 对话UIs(Dialogue UIs)：展示对话可交互性，警告信息，任务日志等。
- 交互系统(InteractionSystem)：在游戏运行时的触发器例如开启对话。

#### 对话数据库(Dialogue Database)

一个对话数据库是一个资产文件包含你的对话，任务和用户自定义的变量。你将可以传统地使用对话编辑器来编辑他的信息，你也可以通过多种不同的格式。

对话数据库包含设计时的内容，即你在Unity编辑器中编辑的字段。在运行时，对话系统将对话数据库视为只读。它加载对话数据库字段到Lua中，在游戏运行时，值可以被修改。

#### 对话编辑器(Dialogue Editor)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditor.png)

用对话编辑器来编辑你的对话数据库并在播放模式下查看其运行时状态。更多关于对话编辑器的信息，详见对话编辑器。通过不同格式例如articy:draft和Chat Mapper，详见导入&导出。

#### 对话管理器(Dialogue Manager)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueManager.png)

对话管理器是场景中一个游戏物体，用来协调所有对话系统的活动，并持有对话系统运行时的数据。对话系统管理部分会详细说明这部分的内容。

#### 对话UIs(Dialogue UIs)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueUIs.png)

对话系统采用如下UIs

- 对话UI：运行交互对话和警告信息.（例子展示如上）
- 任务日志窗口：显示玩家当前任务和已完成的任务。
- 任务追踪HUD(Heads-Up Display抬头显示器)：显示当前任务的追踪信息。

更多关于对话UIs的信息，参考对话UIs.更多关于任务UIs的信息，参考任务。

#### 交互系统(Interaction System)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueSystemTrigger.png)

对话系统提供组件来与游戏物体进行交互并执行诸如开启对话或更新任务等操作。

#### 本地化(Localization)

对话系统支持数据库内容和通用UI元素的本地化。更多信息，详见本地化。

#### 拓展对话系统(Extending the Dialogue System)

虽然对话系统的默认功能已经非常强大和灵活,但你也可以创建自己的逻辑函数、过场动画动作、UI类型等。这些内容在脚本编写中有详细介绍。

### 快速开始 (Quick Start)

我们在这里为您提供帮助!如果你在任何环节遇到困难，请联系我们。

提示：如果你正在制作2D项目，参考开启2D物理支持。

(点击此处观看视频版本)

本章节将直接进入对话系统来创建和运行对话的过程。您将学习到:

- 创建一个对话数据库和编写对话。
- 设置对话管理器戏物体。
- 开启一系列对话。
- 与物体之间的交互。
- 在对话过程中禁止玩家控制

您可以在紧跟随本快速入门部分之后的章节中找到更详细的步骤。

步骤1：创建一个场景（文件 →新场景）

步骤2：抓取预制体(资源>插件>像素破碎者Pixel Crushers>对话系统>预制体>对话管理者Dialogue Manager)到场景中:

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueManager.png)

步骤3：这个游戏对象已经指向一个基本的对话UI，目前足够使用。但是我们需要创建一个对话数据库。检查该游戏对象。在“初始数据库”字段旁,点击“创建”：

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueManagerCreate.png)

步骤4：在为您的数据库指定名称后，再次检查对话管理器游戏对象。点击编辑按钮或者对话系统的logo横幅，将会打开对话编辑器窗口：

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueEditorOpen.png)

步骤5：点击对话页签。

- 点击“+”按钮来添加一个新的对话
- 右键点击在橘黄色的START按钮上来选择创建新的子节点(Create Child Node)。在检视窗口中，设置对话文字为“Hello”。确保设置的是对话文字字段,而不是别的字段例如标题(Title)。
  - 提示：你也可以双击节点来直接输入对话文字在节点上。
- 右键点击在灰色节点上"Hello"接着再次选择创建子节点。设置对话文本为"Goodbye"
- 灰色节点是NPC们说话的节点,而蓝色节点代表的是玩家。你的对话应该显示如下：![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueEditorConversation.png)

步骤6：创建一个空游戏物体给玩家(游戏物体→创建 空)接着创建一个立方体给NPC（游戏物体→3D 物体→立方体）重命名空游戏对象为“玩家”。重命名立方体为“NPC”接着设置它的位置到(0,0,0)，以便在游戏视图中可见。（在您自己的项目中，您将使用实际的角色代替。）

步骤7：检视这个NPC。添加一个对话系统触发器组件：![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueSystemTrigger.png)

步骤8：点击添加行为并选择开启对话：![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueSystemTriggerConversation.png)

步骤9：从对话下拉菜单中选择您的对话，将“玩家”分配给“对话角色”字段，将“NPC”分配给“对话参与者”。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueSystemTriggerConversationSetup.png)

步骤10：最终，设置(Dialogue System Trigger组件中的)触发下拉菜单为在启动时(On Start)。这将告诉系统触发器在场景开始时执行其操作。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueSystemTriggerOnStart.png)

步骤11：现在运行你的第一个对话！对话系统将显示NPC的台词"Hello"，随后很快出现玩家的回应菜单，其中包含一个选项“Goodbye”。它使用的是基础标准对话UI，这是一个通用模板，您之后可以自定义外观，使其符合您的需求。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartPlay.png)

下一步，我们将会使用交互系统来开启一段对话。

步骤12：设置(Dialogue System Trigger组件中的)触发下拉框回被使用时(On Use)。这将会告诉系统触发器在接收到被使用的消息(On Use)时来执行其操作，通常由玩家的交互系统触发。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueSystemTriggerOnUse.png)

步骤13：添加一个可交互组件(Usable)。这将指示告诉玩家的交互组件该NPC是可交互的。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartUsable.png)

步骤14：检视玩家和添加一个选择器组件(Selector)。将选择在下拉菜单(Select At)设置为鼠标位置。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartSelector.png)

步骤15：现在重新运行场景。把你的鼠标放到方块上，你就能看到黄色的信息(NPC spacebar to interact)![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartSelectorPlay.png)

这意味着玩家的选择器组件(Selector)已经检测到NPC的可交互组件(Usable)。按下空格按钮或者鼠标右键。选择器组件(Selector)将会发送交互信息给NPC，使得NPC的对话系统触发器将响应并开始对话。

选择器组件(Selector)和一个类似的组件，叫做接近选择器(Proximity Selector)，提供了触发对话系统活动的方法。他们非常好自定义。你可以配置它们以不同的方式来选择可交互物品，而不仅仅是通过鼠标指针。使用它们也不是强制的。如果你想，你可以用自己的交互系统或者C#代码来运行对话系统触发器。

步骤16：你可能注意到，选择器组件(Selector)在对话期间保持激活状态。在很多情况下，在对话期间你可能需要禁用选择器组件(Selector)，并且可能还需要禁用其他组件,例如玩家移动和相机控制。有一个方法就是添加对话系统事件组件(Dialogue System Events)。这个组件有运行在各种对话系统活动中触发的事件。配置在对话开启的时候禁用选择器组件(Selector)，而在对话结束的时候重新开启它。![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueSystemEvents.png)

#### 快速开始总结(Quick Start Summary)

本快速开始展示了如何：

- 创建对话数据库并且编写对话内容
- 设置对话管理器游戏对象
- 开启对话
- 与对象交互
- 在对话期间禁用玩家控制

对话系统还有更多功能。继续阅读以获取更多信息。如果有任何问题或需要帮助，请练习我们。

## 对话编辑器 — Dialogue Editor

本章介绍如何使用对话编辑器来编写对话和任务。要导入和导出其他的格式，详见《导入&导出》

### 对话数据库(Dialogue Database)

对话数据库是一个资产文件。要创建对话数据库：

- 点击对话管理器的创建按钮（如果尚未分配对话数据库）
- 或者选择资产**→** 创建→像素破碎者→对话系统→对话数据库。

一个对话数据库包含以下类别：

|    类别    | 描述                                                         |
| :--------: | :----------------------------------------------------------- |
| 数据库属性 | 一些通用设置，如描述、作者和版本信息，以及数据库范围的操作。 |
|    角色    | 对话角色(玩家、NPC、可互动对象)                              |
| 物品/任务  | 物品和任务的定义（物品是可选的数据记录，可供使用）           |
|    地点    | 有关地点的信息（可选的数据记录，可供使用）                   |
|    变量    | 用户定义的变量，用于跟踪对话决策等。                         |
|    对话    | 参与者之间的对话内容。                                       |

下面的对话编辑器部分描述了在编辑内容时每个类别的具体作用。

### 对话编辑器窗口(Dialogue Editor Window)

打开对话编辑器，使用菜单项工具→像素破碎者→对话系统→对话系统编辑器，或者点击对话管理器的标志横幅，或者双击对话数据库的资产，对话编辑器窗口将会与检视器配合使用。

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditor.png)

### 数据库(Database)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorDatabase.png)

数据库类别包含通用对话数据库设置。

数据库属性折叠项包含：

- 作者，版本，描述：供您使用的信息字段。
- 全局用户脚本：在运行时加载的可选Lua代码。
- 强调设置：与可选的标记标签有关的颜色值。

全局搜索和替换(Global Search & Replace)帮助你搜索(选择性地替换)文本在你的对话数据库中。

合并数据库(Merge Database)折叠项允许你将另一个对话数据库的内容合并到当前你正在编辑的数据库中。

导出数据库(Export Database)折叠项让你能够导出你的数据库到不同的格式，如配音脚本和CSV(电子表格)。有关详细信息，请参见《导入&导出》。

本地化导入/导出(Localization Export/Import )折叠项相比于导出数据库(Export Database)让你能够更简单地导出本地化的CSV文件。有关更多信息，请参见《本地化》。

编辑器设置(Editor Settings)折叠项包含一些设置，比如自动备份数据库的位置和频率。

### 角色(Actors)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorActors.png)

角色标签定义了对话的参与者。你可以指定角色的名字(以及在用户界面中使用的可选显示名称)，可选的头像图片等等。头像图片可以是纹理（"编辑器和旧版GUI"）资源或精灵图。如果你给角色使用的是精灵图，仅在头像部分使用它们，并将头像纹理保持未分配状态。同样地，如果你使用头像纹理，请保持头像精灵未分配。

标记为是玩家(Is Player)可以使用玩家响应菜单向玩家呈现对话选项。你还可以自定义数据字段，在所有字段(All Fields)折叠项中能够找到这些字段。

### 任务/物品(Quests/Items)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorQuestsItems.png)

任务物品标签定义了任务和物品。物品只是数据记录，您可以根据自己的需求使用。对话系统不会对物品做任何特殊处理。

任务会被用在任务系统中，任务系统的详细信息参见任务章节。

使用此标签定义任务和可选择的任务条目(子任务)。你的任务文本可以包含标记标签。状态是任务的初始状态。这个值在运行时是不会实时更新；相反，您可以在观察(Watches，只有在陨石时才出现在页签中)标签中查看任务的实时状态。如果你想要任务出现在任务跟踪器HUD中，勾选可跟踪(Trackable)和启动时跟踪(Track on Start)这会在任务激活时立即开始跟踪。

### 位置(Locations)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorLocations.png)

位置只是记录数据，您可以根据自己的需求来使用。对话系统不会做任何特殊的处理。

### 变量(Variables)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorVariables.png)

变量是您在游戏过程中可以设置和检查的值。使用它们来追踪信息，例如玩家做出的选择，任务所需收集或者击败的目标数量等。对话标签和对话组件提供下拉菜单来选择变量和检查或者设置他们。

对话系统在对话开启时会自动创建并设置四个变量。您无需添加他们到您的数据库中，但如果希望在对话中方便地通过下拉菜单引用它们，可以将他们添加到数据库中。

- 角色(Actor)：对话主要角色的显示名称(通常是玩家)。
- 对话者(Conversant)：对话主要角色的显示名称(通常是玩家)。
- 角色序号(ActorIndex)：提供给可选的底层Lua代码的角色在角色表中的索引。
- 对话序号(ConversantIndex)：对话者在角色表中的索引。

如果你的对话对话内容分配给了变量警告(Alert)，对话系统在对话结束时会将这段话作为警告消息显示。您还可以使用ShowAlert()的Lua函数，如 《逻辑&Lua》中所述。

#### **变量组(Variable Groups)**

你可以通过在变量名中使用句号（.）来对变量名进行分组。下拉菜单将把变量组显示在不同的子菜单中，方便您在有大量变量的时候进行选择。

你也可以选择菜单→使用折叠组，将变量折叠成每个变量组的独立折叠栏。

此外，您可以在变量部分顶部的搜索过滤器中输入字符串(例如组名)，以筛选显示的变量。

#### **变量视图窗口(Variable View Window)**

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/variableViewWindow.png)

对话系统为了简便性也提供一个辅助的变量视图工具窗口。你可以在对话编辑器窗口打开时使用它来查看和编辑变量，或者在播放模式下查看变量的运行时的值。你还可以在对话编辑器的观察标签中查看运行时的值。

要打开变量视图窗口，选择菜单项工具→像素破碎者→对话系统→工具→变量查看器。或者也可以，从对话编辑器的变量标签页，选择菜单→变量查看器...

#### **重命名工具(Renamer Tool)**

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/renamerTool.png)

你可以用重命名工具在所有数据库、预制体和场景中重命名变量。在场景和预制体中，它会更新对话系统触发器和销毁时增加的组件中引用的变量，但不会更新自定义脚本中的变量名称。

要打开重命名工具，选择菜单项工具→像素破碎者→对话系统→工具→资源重命名器。

提示：更改角色或任务名称的最佳方法是使用显示名称。这样还可以实现显示名称的本地化。

### 对话(Conversations)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorConversations.png)

这是对话编辑器的核心。你可以使用节点编辑器（如上所示）或大纲编辑器（在本节末尾描述）来创建、编辑和删除对话。在游戏过程中，此标签会显示当前对话的实时视图。当前节点会显示为绿色。

#### **节点编辑器(Node Editor)**

下面的表格总结了您在节点编辑器中可以执行的主要操作：

|     操作      | 步骤                                                         |
| :-----------: | ------------------------------------------------------------ |
|  创建新对话   | 单击对话标题下拉菜单右侧的“+”按钮。                          |
| 编辑对话属性  | 点击画布的空白区域。可用于设置标题、主要参与者，并覆盖对话管理器的显示设置。 |
| 添加/删除节点 | 右键单击节点以打开上下文菜单。（在添加子节点时，按住Shift键可以添加相同角色的对话，而不是交换对话的双方） |
|   编辑节点    | 点击节点                                                     |
| 编辑链接箭头  | 点击链接箭头                                                 |
|    节点组     | 点击并按住鼠标，圈选一组节点。松开鼠标按钮时按住Ctrl键。您可以整体移动节点组。 |
| 更改显示设置  | 使用右上角的菜单                                             |
|   运行对话    | 使用菜单>播放...或右键点击并从上下文菜单中选择播放，在编辑模式下测试您的对话。 |

上下文菜单和菜单下拉框不仅提供了以上这些选项还有更多功能。

编辑器也提供了键盘快捷键。

节点上的图标提供快速指示某些值：

- 蓝色三角：已设置序列字段
- 红色问号：已设置条件字段
- 绿色箭头：已设置脚本字段
- 金色星星：已设置执行OnExecute()事件

一个对话通常有一个主要的对话参与者（通常是玩家）和对话对象（通常是NPC）。你可以分配节点给额外的角色来进行对话。

节点检视器包含下面几个字段：

|     字段     | 描述                                                         |
| :----------: | ------------------------------------------------------------ |
|     编号     | 条目的内部编号。通常不要更改。                               |
|     标题     | 仅供自己参考的可选文本。不会向玩家展示。                     |
|     描述     | 仅供自己参考的可选文本。不会向玩家展示。                     |
|     角色     | 说出这句对话的角色。                                         |
|     对象     | 讲话者所面对的角色                                           |
|      组      | 如果勾选，表示将此节点标记为穿透节点用于分组子节点。         |
|   菜单文本   | 显示在玩家响应菜单中的文本，例如对话文本的简短概述。如果为空，则默认为对话文本。 |
|   对话文本   | 显示在字幕中的文本。如果为空，则默认为菜单文本。<br />注意：您还可以双击节点来直接输入对话文本。 |
|     序列     | 通过下拉菜单手动设置链接，调整它们的优先级和顺序。您还可以再此处添加跨对话链接。详情参见《动画序列》。 |
| 响应菜单序列 | 可选序列，在响应菜单可见时于后台播放。如果玩家在响应菜单序列播放期间做出选择，则该序列将结束，并且下一个对话条目的序列将在同一帧开始。(是否为默认选项?) |
|     条件     | 通过下拉菜单手动设置链接，调整它们的优先级和顺序。您还可以再此处添加跨对话链接。详情参见《条件》部分。 |
|   错误条件   | 在错误条件的时候执行。阻止(Block)表示不考虑从该节点链接的节点。穿透(Passthrough)表示忽略该条，但考虑他的子节点以进行对话。 |
|     脚本     | 节点正在对话时要执行的内容。你可以使用...下拉菜单或直接输入Lua代码。<br />注意：脚本字段运行在DialogueManager.currentConversationState属性设置之前。这是因为脚本字段可能会影响DialogueManager.currentConversationState字段的值。如果脚本中的自定义函数需要访问DialogueManager.currentConversationState，请将他们编写在帧结束时等待执行。 |
|     事件     | 当此节点被触发时要执行的Unity事件。                          |
|    链接到    | 通过下拉菜单手动设置链接，调整它们的优先级和顺序。您还可以再此处添加跨对话链接。详情参见《对话位置堆栈功能》。 |

节点(Nodes)也被称为对话条目(Dialogue entries)。本手册中这两个术语是转换使用。

**对话目录(Conversation Menu)**

对话部分有一个菜单，其中包含多个选项，例如查看会话的主要属性、在节点编辑器中显示/隐藏信息（如ID何头像）、对齐网格以及重新排序内部ID。提示：你可以以深度优先的方式重新排序ID，或者使用另一方法，先对父节点进行排序再对合并的子节点进行排序。

**开始节点(START node)**

每个对话都以一个橙色节点开始，通常应保持不变。右键点击开始节点并且选择创建子节点来从这个入口节点创建实际对话节点。

开始节点有一些独特的属性

- 如果序列字段是空，它将默认为None()，而不是对话管理器的默认序列。这会告知对话系统立即跳转到下一个节点。
- 开始节点永远不会响应菜单项显示。如果将其分配给玩家，它会显示为玩家角色的字幕。如果你想要将开始节点的文本显示为玩家角色的字幕，可以将其序列设置为Delay({{end}})或类似的值，并勾选在行期间显示玩家字幕的选项。
- 条件选项不会应用在开始节点上

**节点链接和优先级(Node Links and Priority)**

你也可以选择在链接的节点之间选择他们的优先级。如果一个节点有不同优先级的外部链接，对话系统将会检查所有链接中优先级最高的。如果在该优先级找到任何有效链接，系统将停止检查并仅使用这些连接。否则，它会检查下一个优先级，以此类推。

链接到另一个对话，检查一个节点并在链接到下拉菜单中选择另一个对话。

**标记标签(Markup Tags)**

你可以在你的对话文本中使用这些标记标签，还可以用于演员显示名称、任务显示名称和任务描述的格式化。

|        标签         | 描述                                                         |
| :-----------------: | ------------------------------------------------------------ |
|         [f]         | （强制）即使这是唯一的响应选择，并且即使对话管理器的始终强制响应菜单复选框未勾选，也会强制显示响应菜单。 |
|       [auto]        | （自动响应）强制自动选择该响应，跳过响应菜单。此标签的优先级高于[f]或对话管理器的始终强制响应菜单复选框。 |
|         [a]         | 将响应菜单中的文本显示为斜体。                               |
|    [nosubtitle]     | 不将此文本显示为字幕。                                       |
|   [em#]...[/em#]    | （强调）对文本应用强调设置。                                 |
|   [var=variable]    | （变量）将标签替换为变量的值。如果变量名包含空格或者连字符，必须将其替换为下划线。例如，如果变量名为My Job-Title，则应使用"I need a doctor, not a [var=My_Job_Title]!"。 |
|   [var=?variable]   | （变量输入）提示玩家输入，并将输入存储到varName的变量中。在内部，这会运行TextInput()序列命令，格式为TextInput(文本字段UI，变量名，变量名)。需要注意的是，文本输入框UI必须命名为Text Field UI。 |
| [autocase=variable] | 就像[var=variable]，但会在句首将第一个字母大写，其余部分为小写 |
|       [pic=#]       | 覆盖角色的肖像图片。该数字是角色在角色标签页中定义的肖像图片列表的索引。 |
|      [pica=#]       | 用主角的图像来覆盖该角色的肖像。                             |
|      [picc=#]       | 覆盖对话者（倾听者）的肖像图片                               |
|    [position=#]     | 使用特定的响应按钮位置。如果条目在响应菜单中显示，它将始终处于指定的位置。例如，如果邪恶选项始终在位置6，可以使用此标签:“[position=6]I hate puppies.”位置编号是从零开始的（即从[posistion=]开始），以便与对话UI的按钮列表对齐。 |
|      [panel=#]      | 在使用标准对话UI系统时，使用特定的字幕面板。                 |
|     [lua(code)]     | 用Lua代码的返回值替换标签例如:"Sorry, that item is [lua(Variable["PlayerFunds"]+1)]gold,out of your price range." |

如果您使用Chat Mapper，这些标签是对话系统新增的，Chat Mapper无法识别它们:[auto]、[position]、[nosubtitle]、[panel]、[lua]。

**拆分对话文本(Splitting Dialogue Text)**

如果节点的对话文本太长，您希望将其拆分成多个节点，如果在希望拆分的位置加入管道字符("|")，然后选择菜单→拆分管道到节点→处理对话。

**对话条目节点组(Dialogue Entry Node Groups)**

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorEntryGroups.png)

要对节点进行分组，可以点击并拖动鼠标圈选节点。在释放鼠标按钮时按住Ctrl键。

要重命名或删除一个分组（不删除节点本身），可以右键单击分组标题。

**键盘快捷键(Keyboard Shortcuts)**

你在节点编辑器中可以用下面这些快捷键：

|   键盘组合   |    替代     | 操作                                                         |
| :----------: | :---------: | ------------------------------------------------------------ |
|     Home     |      —      | 移动到画布顶端                                               |
|  PgUp/PgDn   |      —      | 上下翻页                                                     |
|    Delete    |  Backspace  | 删除节点或者链接                                             |
|  Ctrl+Alt+D  |  Cmd+Alt+D  | 赋值当前节点或链接                                           |
|  Ctrl+Alt+N  |  Cmd+Alt+N  | 创建新节点                                                   |
|  Ctrl+Alt+C  |  Cmd+Alt+C  | 将当前选择复制到剪贴板。                                     |
|  Ctrl+Alt+V  |  Cmd+Alt+V  | 粘贴剪贴板内容。                                             |
|  Alt+Click   |      —      | 创建链接到选择节点。                                         |
|    Shift     |      —      | 在创建子节点时，按住Shift键可以保持相同的角色分配，而不会交换它们。 |
| Ctrl+Shift+< | Cmd+Shift+< | 去往上一个对话                                               |
| Ctrl+Shift+> | Cmd+Shift+> | 去往下一个对话                                               |

#### **条件(Conditions)**

你可以使用点击选择下拉菜单或手动输入，将Lua表达式添加到条件字段中，使得只有在Lua表达式为真时，对话才会使用该条目。例如，假设您希望NPC只有在玩家通过访问巴黎时才说出某句话。您可以定义一个变量“VisitParis”，其初始值为false。当玩家访问巴黎时，将该变量设置为true。然后，向NPC的对话条目中加入：

- 对话文本："所以，你认为巴黎怎么样？"
- 条件：Variable["VisitParis"] == true

**对话评估条件时额外向前推进一级(Conversations Evaluate Conditions One Extra Level Ahead)**

对话系统会在当前对话的基础上额外向前评估一级链接。这样做是为了正确处理某些继续按钮模式。这意味着，如果你在一个节点(例如，节点A)中设置了一个值，你不能在下一个节点（节点B）中检查该值，因为对话系统会在执行节点之前，已经检查过节点B的条件。相反，你需要一个空的间隔节点来延迟条件的评估。为了演示这一点，我们将使用一个抛硬币的示例对话：

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/delayConditionEvaluation.png)

在上面的示例中，抛硬币节点将变量x随机设为1或者2.我们不能立即跟随这两个条件节点（正面或者反面），因为对话系统会提前评估一级—这意味着在抛硬币节点执行之前，系统已经检查过正面和反面的条件。

为了解决这个问题，我们在对话中添加一个中间节点，将标题设置为“延迟评估”以提醒我们它的用途，并将对话文本留空。然后我们将序列设置为None(),这样它会立即进入对话的下一步。如果你使用的是继续按钮，则在序列字段中选择“+”→继续→评估继续按钮点击，而不是使用None()。

在处理条件时，临时将对话管理器的调试级别设置为Info可能会很有帮助。这将向控制台窗口添加大量信息。如果你仔细查看，会找到类似以下几行内容：

```
Dialogue System: Lua(x = math.random(2))
Dialogue System: Referee says 'Flipping coin...'
Dialogue System: Block on False Link (Referee): ID=5:3 'Heads!' Condition='x == 1'
Dialogue System: Add Link (Referee): ID=5:4 'Tails!' (True)
Dialogue System: Referee says 'Tails!'
```

第3—4行表示链接到正面被阻止，因为x==1为假，但它添加了到反面的链接（因为x==2为真）。

注意：如果你不想处理这个问题，可以勾选对话管理器的其他设置>在字幕后面重新评估链接。这将在显示字幕后重新进行评估，因此您无需像上面呢样创建中间节点，但代价是会对子节点进行二次评估。

**模拟仿真(SimStatus)**

如果你想要追踪哪些对话已经被使用(访问过)，例如为了有条件地阻止他们被使用，可以使用模拟仿真。检查对话管理器并勾选包括模拟仿真。这将为每个对话条目添加一个运行时字符串值Dialog[#].SimStatus，其中#是对话条目的ID号。你可以通过检查对话条目来找到ID号。它还会设置一个内部的Lua变量"thisID"，你可以从"..."下拉菜单中选择它来引用当前节点。模拟仿真每个对话条目大约使用20字节的内存，因此如果你有成千上万行对话，除非必要，否则尽量不要勾选这个选项。相反，在需要条件的特定区域使用变量。

模拟仿真可能的值:

- 未涉及Untouched：条目从未被说出或显示在玩家回应菜单中。
- 提供过WasOffered：在玩家回应菜单中显示过，但未被玩家选择。
- 已经展示过WasDisplayed：已经被说出由玩家或者NPC。

你可以像这段对话一样使用它（为了简洁，下面用大纲显示）:

```
[1] 舞者: "侦探，我能为你做什么？"
     [2] 玩家: "你星期六晚上在哪儿？"
         条件: Dialog[3].SimStatus == "Untouched"
          [3] 舞者: "我一直在这里，直到黎明。" [结束]
     [4] 玩家: "你说你一直待在这里直到黎明。那么是谁在晚上9点从你家打的电话？"
         条件: Dialog[3].SimStatus == "WasDisplayed"
          [5] 舞者: "我表弟这周住在我这儿。关你什么事？" [结束]
```

你也可以使用特殊变量值thisID来引用当前节点的ID。例如：

```
Dialog[thisID].SimStatus ~= "WasDisplayed"
```

要访问另一个对话的模拟仿真

```
Conversation[#].Dialog[#].SimStatus == "WasDisplayed"
```

第一个#是对话的ID，第二个#是对话节点的ID。（单独使用 Dialog[#]只是当前对话的快捷方式。）

对于勾选了组框的对话条目，仿真模拟始终是未涉及Untouched。

**在回应菜单中以不同格式呈现旧回答(Formatting Old Responses Differently In Response Menu)**

如果你使用仿真模拟，对话管理器输入设置中有一个下拉菜单"Em Tag"用于旧的回答。你可以将其从“None”更改为应用强调标签，以标记玩家已经选择过的回答条目。有些游戏使用这个功能来提醒玩家已经尝试过那些回答。

OnExecute() UnityEvents

事件折叠框中包含一个名为OnExecute()的Unity事件。你还可以点击添加场景事件来添加第二个OnExecute()事件。这些Unity事件类似于将事件处理程序添加到UnityUI按钮的OnClick事件中。

#### **默认OnExecute() UnityEvent**

你不能将场景对象分配给默认的OnExecute()事件。这是因为对话数据库独立于场景存在。相反，如果你必须分配一个资产文件，如预制体或者脚本化对象。

在下面的示例中，我们将创建一个脚本化对象资产，该资产提供一个方法来播放音频剪辑。然后，我们把这个方法分配给对话条目的OnExecute()事件。

首先，用下面的代码创建一个新的C#脚本叫做测试脚本化对象(TestScriptableObject)。为了让例子足够简洁，代码没有做任何错误检查。

TestScriptableObject.cs

```C#
using UnityEngine;
[CreateAssetMenu(fileName = "TestScriptableObject")]
public class TestScriptableObject : ScriptableObject
{
    public AudioClip[] audioClips;
    public void PlayAudioClip(int i)
    {
        AudioSource.PlayClipAtPoint(audioClips[i], Vector3.zero);
    }
}
```

接着，在项目中选择创建→测试脚本化对象TestScriptableObject。这将会创建如下文件：

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/onExecuteScriptableObject.png)

你可能需要将你的资源移动到一个名为"Resources"的文件夹中。这样它将被包含在构建中。在某些Unity版本中，一些平台会剔除它认为不需要的资源，最常见的情况是未在任何场景中引用的脚本化对象资源会被剔除。

检查你的新资产文件。你将能够将音频剪辑分配给它：

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/onExecuteScriptableObjectEdit.png)

最终，检查对话条目节点的OnExecute()事件，点击"+"按钮，来分配TestScriptableObject。选择你想要执行的方法(播放音乐PlayAudioClip)并指定你想播放的音频剪辑的索引:

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/onExecuteScriptableObjectAssign.png)

提示：默认的OnExecute()不支持数据库导出功能或第三方格式导入器。

**基于场景的OnExecute()Unity事件  (Scene-Based OnExecute() UnityEvent)**

如果你点击添加场景事件，它将会添加第二个OnExecute()事件，你可以分配场景对象给这个事件。实际的UnityEvent存在于当前场景中，因此它与数据库导出功能兼容。

**大纲编辑器Outline Editor**

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorOutlineMode.png)

你还可以像上面所示那样在大纲模式下编辑对话，这对使用过类似Bioware的Aurora工具集。要切换到大纲模式，请选择菜单→大纲。

### 模板(Templates)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorTemplate.png)

使用模板页签来更改添加新演员、任务、对话条目等中的默认字段。你还可以更改对话选项卡中使用的颜色。右上角的菜单允许你保存和加载模板设置，将模板填充为数据库中已存在的自定义字段，并将模板应用于数据库中的所有内容。

要添加新的自定义字段，展开你想要自定义的类别。然后点击灰色的“+”圆圈，并提供字段的详细信息。在下图中，我们已经添加一个自定义岁数字段给演员。我们还选择勾选主要复选框，这意味着对话编辑器会在演员的主检视部分显示该字段，而不仅仅是所在字段折叠区域中的显示。

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorTemplateAdd.png)

#### **自定义字段类型(Custom Field Types)**

如果你熟悉脚本编写，可以定义你自己的字段类型，并为其创建自定义的对话编辑器显示器。要定义一个新的自定义字段类型：

1. 复制模板Templates→脚本Scripts→编辑器Edior→CustomFieldType_TemplateType.cs文件，并根据你的字段类型重命名。例如，你想要定义一个颜色字段类型，将其重命名为CustomFieldType_ColorType.cs。这个文件必须在Editor文件夹或其子文件夹下。
2. 编辑脚本。在注释所示的位置添加你的代码。在内部，字段始终是字符串，但是你可以根据需求解释这些字符串。对于颜色类型，你可以使用标准RGB颜色格式例如“#ff00cc”并使用对话系统Tools.WebColor()和Tools.ToWebColor()方法在颜色和字符串之间进行转换。

你可以在脚本Scripts→编辑器Editor→字段Fields→例子Examples中找到案例。

注意：自定义字段类型是一个设计时的遍历功能，能够让你添加自定义的编辑器绘制器。在运行时，所有自定义字段都被当做文本处理。

### 对话模板(Conversation Templates)

你可以将对话书结构保存为模板，然后在其他结构相似的对话中重新使用它们。

要保存对话模板，选择菜单>模板>保存模板Json...

要从模板创建新对话，选择菜单>模板>从模板新建。你可以萱蕚一个内置模板（例如通用任务对话模板）或者之前保存的Json模板。

### 监视(Watches)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorWatches.png)

在运行时，监视标签将会替代模板标签。在这个标签中，你可以查看变量、任务或者任意Lua表达式的运行时值。你还可以使用底部的代码栏运行Lua表达式。

### 检查问题(Check For Issues)

对话编辑器的数据库部分有一个“检查问题”折叠项，可以检查你的对话数据库是否存在一些常见问题，例如：

- 检查序列：序列字段中的语法错误。
- 找到孤立节点：没有从任何父节点链接的对话条目。
- 找到未定义的变量：引用未在变量部分定义的变量。
- 检查条目标签音频：引用了没有对应音频文件的条目标签，通常用于Audio/AudioWat()序列命令。

### 使用多个数据库(Working With Multiple Databases)

通常，使用单一的对话数据库要容易得多。然后，对话系统支持多个数据库。多个数据库需要你更好地进行管理，以便跟踪它们并根据需要加载和卸载内存，但他们可以用于将内容拆分开来，既可以为组织结构服务，也可以减少任何给定时刻加载到内存中的内容量。

#### **选择单一数据库啊还是多个数据库(Deciding Between Single Or Multiple Databases)**

以下是选择使用单一数据库还是多个数据库的一些指南:

- 一个包含1000个平均长度对话的对话数据库大约使用15MB内存，相当于几个纹理文件大小。
- 如果你正在制作桌游，15MB可以不用考虑。在这种情况下，使用多个数据库的原因是为了方便组织管理。
- 如果你正在构建低端移动设备，15MB可能太多，但是最终还是取决于你的规格和需求。
- 如果你想要使用多个数据库
  - 单个数据库的大小超过你愿意分配的内存预算。
  - 你的数据库包含了太多不同的元素，以至于需要拆分以便进行组织。

无论你使用单个还是多个数据库，记住你可以在对话标题中使用正斜杠来对对话进行分组。例如，你可以将对话命名为：

- 同伴/机器人管家/随时为您服务
- 同伴/机器人管家/故障
- 沙漠/拾荒者
- 沙漠/蝎子牧人
- 丛林/萨满等

这将把它们分组到对话编辑器中的子菜单同伴，沙漠和丛林，使选择对话变得更加容易。同伴子菜单下会有一个机器人管家的子菜单。

**多个数据库的使用技巧(Tips for Multiple Databases)**

如果你选择使用多个数据库，把你所有的全局数据放到全局数据库中，接着分配全局数据库作为你对话管理器的初始数据库。这应该包括保存游戏中需要的一切内容，或无论位置如何都可以访问的内容，例如：

- 玩家角色以及所有可以跨场景的角色，例如跟随玩家的同伴。
- 任务
- 全局变量
- 可以在任何地方播放的对话，例如同伴闲聊。在其他数据库中，使用“从数据库同步”（如下所述）来同步全局数据内容。这样，你就可以在对话编辑器中任何数据库中引用全局数据。

当使用多个数据库，你可以设置数据库的基本ID以减少在测试游戏时运行唯一ID工具的需求。

**从其他数据库同步资产(Sync Assets From Another Database)**

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueEditorSyncDB.png)

如果你使用多个对话数据库，你可能希望在一个数据库中定义所有角色，在另一个数据库中定义所有的任务等。你可以配置一个对话数据库来提取这些数据库中的元素。为了实现这一步，在标签菜单Menu选择中选择从数据库中同步。接着选择你想要同步元素的数据库来源。

你可以添加更多元素，但是记住具有与源数据库中相同ID的元素将会被覆盖。举个例子，说你正在编写一个游戏，玩家带着一群伙伴从一个星球到另一个星球。你可以定义玩家和同伴角色在源数据库中。然后，你可以为每个星球创建一个独立的数据库，该数据库从这个源数据库中同步并添加特定于该星球的角色。

#### **唯一ID工具(Unique ID Tool)**

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/uniqueIDTool.png)

如果你只使用一个对话数据库，你可以跳过这个章节。在内部，每个在对话数据库中的元素（角色，物品，对话，等）拥有一个ID号。在单一数据库中，这些ID通常对每种类型的元素都是唯一的。然而，在一个数据库中使用的ID很可能被用在另一个数据库中。如果你同时加载两个数据库，ID会产生冲突。对话系统无法知道你指的是哪个资源，基于给定的ID。为了阻止这样的事情发生，使用唯一ID工具：工具Tools→像素破碎者Pixel Crushers→对话系统Dialogue System→工具 Tools→唯一ID工具Unique ID Tool。只需添加你希望确保唯一ID的数据库，然后点击处理。如果勾选详细日志记录，工具将在控制台中记录所有ID编号的更改。

提示：如果为每一个数据库设置基础ID，可以大大降低ID冲突的可能性，从而减少运行唯一ID工具的需求。

**当使用唯一ID工具时(When to Use the Unique ID Tool)**

每个加载到内存中的对话数据库都需要为每种元素类型提供唯一的ID—例如，每个角色都有一个唯一角色ID。如果你使用从数据库中同步，那么新元素将被分配与全局资产不同的ID，因此你无须担心唯一ID的问题。然而，你的沙漠数据库中的ID很可能会与丛林数据库的ID重叠。不过，这可能没什么问题。只要你不同时加载沙漠和丛林数据库，内存中就不会有ID冲突。

**基础ID(Base ID)**

在对话系统编辑器的数据库标签，你可以为每个数据库分配值(Base ID)通过数据库属性Database Properties>基础ID BaseID。当你添加角色，变量和对话等，他们的内部ID将会从基础ID开始编号。举个例子，设置全局数据库的基础ID是1，下一个数据的基础ID是10000，再下一个的是20000等等。这有助于减少数据库的ID冲突，使得在进行测试时无需先运行唯一ID工具变得更加方便。然而，在最终测试和发布之前，你仍然应该运行唯一ID工具，以确保完全没有ID冲突。

**额外数据库(Extra Databases)**

在运行时加载额外的数据库到内存中，可以使用额外数据库组件。例如，可以将额外数据库组件添加到沙漠场景中的一个空游戏物体上。将添加触发器设置为开始时，并分配沙漠数据库。将移除触发器设置为销毁时，并分配沙漠数据库。

**脚本注意:DialogueManager.masterDatabase**

如果你加载多个数据库到内存中，无论是使用额外数据库组件还是DialogueManager.AddDatabase(),DialogueManager.masterDatabase属性将包含所有加载到内存中的数据库内容。

## 对话管理器 — Dialogue Manager

本章描述如何设置对话管理器游戏物体到你的场景中

### 对话管理器预制体(Dialogue Manager Prefab)

对话管理器协调运行时活动并维护包含对话系统运行时数据的Lua环境。要添加对话管理器，拖拽预制体资产Asset>插件Plugins>像素破碎者Pixel Crushers>对话系统Dialogue System>预制体Prefabs>对话管理器Dialogue Manager到你的场景中：

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/quickStartDialogueManager.png)

对话管理器的默认设置使其在场景切换时保持持续存在，并确保场景中仅存在一个副本。通常，你会在主场景中使用对话管理器。你可以在你想要的位置场景中放置其他对话管理器，这样便无需从主场景进入即可进行测试。但请注意，在正常游戏进行过程中，主场景中的对话管理器会延续，并销毁位置场景中的测试对话管理器。

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueManager.png)

### 对话系统控制器(Dialogue System Controller)

对话系统控制器组件包含对话系统的设置。

#### 初始化数据库(InitDatabase)

对话管理器在游戏开始时会加载对话数据库。

#### 显示设置(Display Setting)

显示设置部分包含下面这些子部分：

|       子部分       | 描述                                                       |
| :----------------: | ---------------------------------------------------------- |
|     本地化设置     | 控制语言本地化如何实现                                     |
|      字幕设置      | 控制对话UI如何显示字幕                                     |
| 相机和过场动画设置 | 控制过场动画序列如何工作                                   |
|      输入设置      | 控制玩家反馈菜单如何在对话中工作                           |
|      喊话设置      | 控制基本喊话行为，你可以在单个喊话者或喊话UI上覆盖这些设置 |
|      警告设置      | 仅系统UI如何显示警告信息在屏幕上                           |

提示：你可以通过在对话编辑器中检查特定对话来改写这些特定对话的设置，选择菜单Menu→对话属性Conversation Properties→接着勾选覆盖显示设置Override Display Settings。

#### 持久化数据设置(Persistent Data Settings)

持久化数据设置部分允许你指定哪些数据会包含在存档中，并且在场景切换时保存。这些设置会影响持久化数据管理器组件，并且如果你使用完整的保存系统，还会影响对话系统保存组件。

#### 其他设置(Other Setting)

其他设置部分包含一些杂项设置。值得关注的是：

|               其他设置               | 描述                                                         |
| :----------------------------------: | ------------------------------------------------------------ |
|            仅允许一个实例            | 通常应该保持勾选状态。同一时间只能存在一个对话管理器。       |
|             加载时不销毁             | 通常应保持勾选状态。这允许对话管理器及其数据在场景切换时保持持续存在。 |
|            预热对话控制器            | 指定在启动时执行的初始化级别，以防止在低规格设备上启动首次对话时出现卡顿。 |
|         在预热起家不立即关闭         | 勾选，如果你的对话UI始终可见即使在没有对话进行的时候。       |
|             实例化数据库             | 在Unity编辑器的播放模式下，将对话数据库的副本加载到内存中而不是直接使用资产。 |
|             包含模拟状态             | 能够追踪模拟状态                                             |
|      在第一个有效结果时停止评估      | 在对话树中，不评估所有子节点，而是找到第一个有效子节点时停止。 |
|        字幕显示后重新评估链接        | 勾选，如果你更改了直接子节点的评估值。更多信息参阅对话在额外一级前评估条件。 |
|            使用线性组模式            | 如果组节点的条件为真，则不评估剩余兄弟组的节点。             |
|         允许同时进行多个对话         | 通常情况下，如果一个对话处于活跃状态，对话系统会允许另一个对话播放。勾选这个允许多个对话在同一时间播放，每个对话必须使用自己的对话UI，你可以分配给对话系统触发器的覆盖对话UI字段或者通过添加覆盖对话UI组件来设置。 |
| 在启动触发器之前，确保保存数据已应用 | 如果使用存储系统加载已保存的游戏或者更换场景，设置为OnStart的对话触发器将首先等待保存数据应用到场景中，而不是在场景开始时立即运行。 |
|             对话时间模式             | 默认情况下，对话系统运行在实时模式，独立于Time.timeScale。如果你希望对话系统遵从Time.timeScale，请将对话时间模式设置为Gameplay |
|               调试等级               | 设置对话系统在Unity编辑器的控制台和构建后的玩家日志文件中的日志级别。这对于追踪对话系统的活动非常有用。 |

### 其他组件(Other Components)

对话管理器预制体还包含以下组件：

#### 实例化预制体(Instantiate Prefabs)

将基本UI预制体实例化到对话管理器的Canvas中。这些包括选择器元素Selector Elements（参见触发器与互动）、任务追踪HUD Quest Tracker HUD和任务日志窗口Quest Log Window（参见任务）。你可以分配不同的预制体来定制游戏的外观、或者如果不需要，可以将其移除。

#### 输入设备管理器(Input Device Manager)

优雅地检测鼠标，手柄和键盘控制之间的切换，并帮助UI界面在手柄和键盘模式下自动聚焦UI按钮，在鼠标模式下保持按钮未聚焦。兼容Unity内置的输入管理器，Unity的输入系统包和重写，并且还提供挂钩支持其他输入系统。

如果你的场景还没有的话，你可能还需要添加基础Unity的事件系统EventSystem。

在Demo中的对话系统有这些额外组件：

- 保存系统：在场景切换时处理保存游戏和持久化数据。
- 标准场景转换器：用于保存系统在场景切换时淡出和淡入。也可以配置为在加载下一个场景是显示加载场景。

- 玩家首选项存档数据存储器：允许保存系统将数据保存到PlayerPrefs中。
- Json数据序列化器：为保存系统提供Json序列化功能。

这些组件在保存系统部分中有详细说明。

#### 对话管理器设置向导(Dialogue Manager Setup Wizard)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueManagerSetupWizard.png)

可选的对话管理器设置向导将逐步引导你进行配置，并详细说明每个设置。要运行该向导，点击对话管理器的向导按钮或者选择菜单项工具Tools→像素破碎者Pixel Crushers→对话系统Dialogue System→向导Wziards→对话系统向导Dialogue Manager Wizard。

## 触发器和交互 — Triggers&Interaction

本章解释如何在对话系统中事物是如何运作。

提示：如果您正在制作2D项目，参考启动Physics2D支持。

### 对话系统触发器(Dialogue System Trigger)

![](https://www.pixelcrushers.com/dialogue_system/manual2x/html/dialogueSystemTrigger.png)

对话系统触发器组件拥有三个主要部分：

- 触发器Tigger：触发对话系统触发器执行的事件。
- 条件Conditions：状态条件必须为真，才能使对话触发器执行其动作。
- 动作Actions：触发器应该执行的操作，例如启动对话或者设置任务状态。

#### 触发器(Trigger)

你可以将触发器下拉菜单设置为以下值：

|     触发器     |       Trigger        | 描述                                                         |
| :------------: | :------------------: | ------------------------------------------------------------ |
|     使用时     |        OnUse         | 玩家选择器或接近选择器向游戏对象发送了OnUse消息，或者在UnityEvent或脚本中手动调用了对话系统触发器的OnUse()方法。 |
|     开始时     |       OnStart        | 组件在(例如，场景开始时)已启动。参见下方注释。               |
| 应用保存数据时 | On Save Data Applied | 保存系统已将保存的数据应用到场景中，或者已经过等量的帧数。   |
|     启用时     |       OnEnable       | 组件被启动的时候。                                           |
|     禁用时     |      OnDisable       | 组件被禁用的时候。                                           |
|     销毁时     |      OnDestroy       | 组件被销毁的时候。                                           |
|  触发器进入时  |    OnTriggerEnter    | 组件接受到触发器进入的消息。要使用这个触发器，组件的游戏物体应该有一个触发碰撞体。你可能需要设置条件→接受标签，以此将触发器限制在特定标签例如Player的游戏物体中。 |
|  触发器退出时  |    OnTriggerExit     | 组件收到触发器退出信息的时候。                               |
|   碰撞进入时   |   OnCollisionEnter   | 组件接收到碰撞进入信息的时候。                               |
|   碰撞退出时   |   OnCollisionExit    | 组件接收到碰撞退出信息的时候。                               |
|   喊话开始时   |     OnBarkStart      | 游戏物体开始播放一段喊话的时候(一次性的对话台词)。           |
|   喊话结束时   |      OnBarkEnd       | 游戏物体结束播放一段喊话的时候。                             |
|   对话开始时   | OnConversationStart  | 游戏物体刚刚作为主要参与者开启对话的时候。                   |
|   对话结束时   |  OnConversationEnd   | 游戏物体刚刚结束一段对话的时候。                             |
|   序列开始时   |   OnSequenceStart    | 游戏物体刚刚作为过场动画中的主要参与者开始。这一时间不会再对话序列中调用，除非你勾选了对话管理器的字幕设置→通知序列开始和结束的复选框。 |
|   序列结束时   |    OnSequenceEnd     | 游戏物体刚刚结束一个序列。                                   |

注意：当触发器设置为OnStart时，如果你正在使用保存系统，触发器可能会在保存系统应用数据保存之前触发。请改用OnSaveDataApplied，或者在代码中设置DialogueManager.OnStartTriggerWaitForSaveDataApplied = true。

#### 条件(Conditions)

条件允许你指定必须为真才能触发对话系统触发器的状态，包括：

|   条件类型   | 描述                                                         |
| :----------: | ------------------------------------------------------------ |
|   Lua条件    | Lua表达式例如检查变量的值                                    |
|   任务条件   | 需要任务状态                                                 |
|   允许标签   | 对于碰撞时和触发时的事件，另一个游戏物体必须有其中一个标签。如果允许标签是空，则允许所有游戏物体。 |
| 允许游戏物体 | 对于碰撞时和触发时的事件，另一个游戏物体必须在此列表中。如果接受的游戏物体为空，则允许所有游戏物体。 |

#### 行为(Actions)

要添加执行的动作，点击添加动作按钮。你可以添加以下这些动作：

|          动作           | 描述                                                         |
| :---------------------: | ------------------------------------------------------------ |
|      设置任务状态       | 设置任务或者任务条目的状态(参见任务)                         |
|       运行Lua代码       | 运行Lua表达式。Lua变量["Actor"]和变量["ActorIndex"]被设置为交互者的信息。 |
|        播放序列         | 播放一个过场动画序列。序列的发言者被设置为交互者的信息。（参见过场动画序列） |
|        显示警告         | 通过对话UI显示警告信息。                                     |
|        发送消息         | 使用Unity的SendMessage()方法向目标发送消息。                 |
|          喊叫           | 在游戏物体上播放一段喊叫信息(一次性的对话台词)。             |
|        开启对话         | 开启一段对话。                                               |
| 设置游戏物体激活/非激活 | 作用于整个游戏物体。                                         |
|      启用/禁用组件      | 作用于游戏物体的特定组件。                                   |
|      设置动画状态       | 在游戏物体上设置动画器状态。常用于在对话开始时使角色保持静止Idle状态。 |
|   在执行时OnExcute()    | 允许你使用UnityEvent指定其他动作。                           |

#### 角色游戏物体分配(Character GameObject Assignments)















→











