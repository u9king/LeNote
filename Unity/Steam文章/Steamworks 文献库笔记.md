# Steamworks 文献库笔记
## 入门指南

### 登记筹备

### Steam Direct 费

### 管理您的Steamworks账户

#### 管理用户

要让其他用户帐户访问 Steamworks 中管理您游戏的任何环节，您需要先在添加/管理用户页面向该用户发送邀请，以添加该用户。 邀请流程如下：

- 输入新用户的电子邮件地址。 邮件无需为任何现有的 Steam 帐户所持有。
- 受邀人将收到一封电子邮件，通知其收到邀请，并提供接受链接。
- 受邀人接受邀请后，您将收到一封电子邮件，告知您对方已接收。 您或您 Steamworks 合作伙伴帐户中的其他管理员需要在七天内批准或拒绝。
- 在确认后，受邀人将收到最后一封电子邮件，欢迎其加入 Steamworks，而您将能把该用户添加入任何管理组，并提供游戏组的权限。

  添加用户并选择特定权限将授予该用户相应权限，这些权限将涵盖那些当前包含在“所有人”组中的所有应用程序。 该“所有人”组为每个 Steamworks 合作伙伴帐户自动创建。 如果希望通过组来管理特定游戏的更具体的访问权限，拥有管理权限的用户可以从“所有人”组中移除特定的应用程序。

#### 实际授权

实际授权是一种特殊的管理权限，与其他类型的权限不同。 实际授权为用户提供了以下权限：

- 查看并修改银行账户信息
- 查看并修改税务信息
- 查看并修改最低发款额
- 签署法律协议
- 发起并接收应用程序转移

相应的，实际授权也受到限制。 若要为一名用户申请实际授权，请联系 Steam 发行团队。 我们建议您定期审核谁拥有此权限。

#### 用户权限

##### 合作伙伴权限

合作伙伴级权限适用于整个 Steamworks 合作伙伴帐户。 当前权限包括：

| ![icon_star.png](https://steamcdn-a.akamaihd.net/steamcommunity/public/images/steamworks_docs/schinese/icon_star.png)管理员 | 允许用户管理合作伙伴帐户中的用户与其权限。 只有拥有管理员权限的用户才能支付 Steam 应用费并激活应用额度。 请参阅 Steam 应用费文档，了解相关信息。 |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ![icon_actual_authority.png](https://steamcdn-a.akamaihd.net/steamcommunity/public/images/steamworks_docs/schinese/icon_actual_authority.png)实际授权 | 此权限只被分配给负责代表您的机构签署文件的帐户，且除 Valve 外任何人都无法将其授予某个用户。 具有“实际授权”权限的用户可以填写并修改银行和税务信息、代表您的机构签署法律协议，其中包括发起和接受转让的权利。 如果您的 Steamworks 帐户中没有人具有“实际授权”权限，或您需要指定新用户具有“实际授权”权限，请联系 Steam 发行团队，让我们知晓。 |

##### 组权限

| 仅限可见性| 没有管理应用程序的权限，但作为 Steamworks 合作伙伴组的成员，可以通过 Developer Comp 程序包访问应用程序，并被授予该组中任何应用程序 Steam 社区的“版主”权限。 |
| ------------ | ------------- |
| 编辑应用元数据| 授予上传 depot，测试 Steamworks 功能，定义程序包、成就、排行榜，发布集换式卡牌等权限。 同时授予组内任何应用程序 Steam 社区的管理员访问权限和身份。 |
|编辑应用的营销数据| 授予编辑和发布商店页面数据，例如游戏描述、品牌形象和视频的权限。 还包括对营销工具（如曝光轮次）的访问。 |
| 编辑商店本地化数据| 授予在商店设置选项卡上编辑任何数据的权限，选项卡包含一个或多个可进行本地化的字段，例如产品描述和图像资产。 不授予发布的权限或访问曝光轮次等营销工具的权限。 *注意：这是属于“编辑应用的营销数据”下的功能。* |
|查看市场营销流量数据| 授予在应用登陆页的“营销与曝光”栏目中查看流量数据的权限。 可以在该栏目中查看内部 Steam 商店流量数据、外部商店流量数据、发行及更新曝光轮次数据和 UTM 分析数据。 这些栏目的某些数据还可以下载为 csv 文件。 |
|在 Steam 上发布应用更改| 授予权限，以发布在您游戏 Steamworks 开发者站点的“技术工具” > “编辑 Steamworks 设置”栏目中做的更改。 |
|查看错误数据| 授予查看您产品错误报告的权限（前提是产品已经与 Steamworks 的错误报告系统整合）。 |
| 生成序列号| 可以为您的应用程序申请批量序列号或上传第三方序列号。|
|查看财务信息| 可以查看您应用程序和程序包的销售与下载数据。|
|管理签名| 可以为您的产品管理数字签名。这将为一款应用程序启用其反作弊与SDK 验证设置的管理。|
| 支持经济体/创意工坊| 允许用户在合作伙伴站点经济体支持工具中查看/授予经济体物品。|
| 经济体/创意工坊主管| 当前具有与上述“支持”同样的功能。|
|管理定价与折扣| 允许用户提出定价变更建议并选择加入折扣与促销活动。 参见定价文档，了解更多信息。 注意：必须有此权限才能完成发行流程，才能发布即将推出页面或发行产品。|
|实况直播| 允许用户播放您应用程序的官方直播流。|

### 管理应用程序

#### 概述

##### 方法一：单个Steam账户/组

##### 方法二：多个Steamworks合作伙伴账户

#### 常见问题

问题：如果我的游戏是由其他公司发行的，那么我如何在不需创建新应用且不支付应用费的情况下创建新的 Steamworks 帐户？

答：您可以遵循此链接中的说明，注册为 Steamworks 受限帐户。 您只需填写《NDA 保密协议》以及公司或个人的相关联系信息即可完成注册流程。 如果您还需要填写银行与税务信息以接收付款，那么请勿使用此链接。

#### 转让应用程序

#### 应用管理共享

问： 如果另一方需要获得应用程序的收入，该怎么办？

答： 您需要将应用程序转让给另一个 Steamworks 帐户。 应用共享不会修改支付规则。 如果您需要转让应用程序，以便另一个 Steamworks 合作伙伴收到款项，但您仍然需要访问该应用程序以进行相关工作，则转让的接收人可以重新与您共享该应用程序，以便您仍旧具备访问权限。 请参见有关应用转让的文档，了解更多信息。（理解为谁收钱谁拿应用，再分享给别人共同工作）

#### 转让应用程序

游戏发生所有权变更，或个人成立了公司，需要将游戏转移至新公司进行分销等情况下会用到转让功能。

##### 必要权限和配置

##### 开始转让

请访问转让应用程序工具，开始转让。您也可以在 Steamworks 网站的“工具”菜单下找到此工具。 您需要具备您的 Steamworks 合作伙伴帐户的实际授权人资格，方能完成应用转让。

##### 转让步骤

转让包括五个步骤。 整个流程所需时间视接收人和转让人确认转让所需时间而定。 一旦双方均完全同意转让应用，Valve 需要对转让进行批准。 Valve 会尽快批准，通常在 2-7 个工作日内完成。

1. 转让人发起转让。应用的当前管理者需要...
2. 接收人批准转让。接收人将会...
3. 转让人接受转让。
4. Valve 对转让进行审查。 如果对转让的原因有疑问，我们会对具有实际授权的用户发送电子邮件，获取更多信息。
5. 转让完成。 一经批准，所有管理权限会转让至新的合作伙伴。

##### 生效日期

转让的生效日期为开始向新合作伙伴付款的日期。

生效日期并非管理权限转让至新合作伙伴的日期。 一经 Valve 批准，管理权限立即进行转让。

系统默认自转让发起日起 7 天后为生效日期。

##### 转让随附数据

##### 常见问题

### 内容调查

#### 概览

#### 内容调查的构成

##### 一般内容

##### 成人内容

##### 生成式人工智能内容

游戏中的 AI 瞬息万变，我们希望能够采取恰当的做法。 许多艺术家和创作者对内容的盗取和滥用表达了合理的担忧。 但我们也意识到科技和工具将会逐步发生巨大的变化，而许多富有创意的游戏开发者们正在探索通过正当使用 AI 来创造优秀娱乐产品的办法。 如果您计划在 Steam 上发行产品，且认为我们忽视了关于 AI 的一些重要方面，我们随时恭候您的反馈。

#### 常见问题

**问： 为什么使用实时生成 AI 的仅限成人的色情内容得到不同的对待？**
答：关于仅限成人的色情内容，Valve 试图通过内容调查和我们自己的发布前审核实现三个目标：遵守并实施法律限制、尊重游戏创作者的创作自由，以及满足顾客的期待。 在游戏过程中加入由 AI 实时生成的仅限成人色情内容将不利于我们实现这三个目标。 这类内容会带来较大的法律风险，也很有可能冒犯顾客，因此我们目前不希望发布由 AI 实时生成的仅限成人色情内容。

#### 德国的强制性年龄评级

##### 审核您所有的游戏

##### 常见问题

### 将手机游戏移植到Steam

这几年，我们看到许多原为手机开发的成功游戏，通过 Steam 转向个人电脑发展。 有些游戏成功转型，而有些则以失败告终。 其中关键不同之处似乎在于，开发者对电脑受众的考虑及应此计划对游戏进行的调整有多深入。 以下并非什么硬性规定，而只是我们从 Steam 上发行的游戏当中观察到的现象。

除这些建议外，您也许还想听一听 James Vaughan 在 GDC （游戏开发者大会）上的发言。 他谈到了把游戏《瘟疫公司》（Plague Inc.）从手机上移植到电脑的情况：http://www.gdcvault.com/play/1021963/Evolving-Plague-Inc-Taking-a

### 常见问题

##### 公司信息

**问： 我能以合伙企业（Partnership）的身份完成数字文书吗？**

答： 能。 如果您的合伙企业处于美国境外，且是以各个合伙人的名义而非以企业名义纳税，那么您需要额外提供每位合伙人的税务信息。

**问： 如果我是独资个人（Sole proprietor），应该怎么填写我的公司名称？**

答： 请勿填写您的经营别称（DBA）、易记名称、别名或昵称。 请输入您的合法全名。 您在此填写的名称必须与您的银行和美国国税局（IRS）税务文件或外国税务文件（如适用）之官方文件上的名称一致。

##### 税务信息

**问： 我刚刚搬了家， 需要重新进行税务认证吗？**

答： 如果您并没有搬离您在税务认证中申明的国家/地区，那么您不需要重要进行税务认证。 如果您搬至了新的国家/地区，则需要重新进行税务认证。

## 商店状态

本节深入介绍了管理您 Steam 商店状态的各种要素。

| 应用程序                                                     | 应用程序（或应用）是 Steam 上产品的主要表现形式。 一个应用通常有自己的商店页面、自己的社区中心，并且会出现在客户的“库”中。 |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| 开发者和发行商主页 |
| 即将推出 | 您可以在产品发行日期之前很早便推出一个“即将推出”页面。 这样，您可以让顾客知道您的产品即将上架 Steam，并围绕该产品开始吸睛造势，积累人气。 |
| 抢先体验 | Steam 抢先体验让您能在游戏尚处于开发阶段时就在 Steam 上销售，并让顾客了解该产品应被视为“未完成”。抢先体验适用于处于可玩的 alpha 或 beta 版本、其当前可玩生成版本物有所值，且计划继续开发以供发行的游戏。 |
| store/editing | 本文将说明发行产品前如何准备商店页面，以及发行后如何更新商店页面。 |
| 产品系列页面 | 为各游戏、电影、软件或 VR 体验定义并设置一个系列，然后将其链接至产品系列主页，从而将相关的游戏或产品关联在一起。 |
| 图像资产 - 概览 | 作为 Steam 发行流程的一部分，您需要为自己的产品提供各类图像资产。 这些图像将会在 Steam 商店和顾客的库中等各种地方显示。 |
| 直播   | 进一步了解如何在 Steam 上的产品页面或开发者主页直播游戏。    |
| 本地化和语言 | Steam 是一个跨国平台，其许多功能支持多种语言。 超过六成 Steam 用户以英语以外的语言使用 Steam，因此为这些用户量身打造使用体验十分重要。 Steam 尽可能多地支持多种语言、货币、支付方式，以便尽量为全世界的顾客提供最佳体验。 本文将阐述 Steam 是如何支持多种语言的。 |
| store/accolades | 赞誉包括您游戏收到的测评和奖项。 此文档将告诉您如何在产品页面添加这些赞誉。 |
| 在 Steam 上进行预购| 关于预购，我们的经验是除非一件作品倍受期待或受到大力推广，否则实质效益不大。即便您的产品符合上述条件，我们仍建议您的产品在 Steam 上的预购期较短为佳。 |
| 定价| Steam 合作伙伴负责为自己的产品设置并管理定价。 Steamworks 开发者网站提供了在 Steam 支持的所有币种下配置定价与折扣的工具。 |
|发行选项| 您在点击“发行应用”时，可以按您的产品与发行时间，从三种商店曝光方式中自由选择。 |
|发行流程| 发行产品大部分的控制权都在您的手中，但在推出之前仍然需要通过 Valve 审核并批准。 您可以自行决定发行日期，也能自行配置绝大部分的功能。 在 Steam 上发行产品的大致流程可在 Steamworks 的产品登陆页面上方找到。 |
|从Steam上移除产品 | 如果出于某些原因，您需要停止销售您的产品，Valve 可以为您进行必要的变更。 如果您要将一款游戏撤销或退场，或是因失去了发行权而必须停止销售某款作品，此文档将助您了解相关信息。 |
|审核流程| 在完成核对清单后，您需要将商店状态与产品生成版本交付 Valve 审核，审核通过后方可发行游戏或软件。 请点击“标记为就绪待审”，让 Valve 知道您已完成各项工作、建好商店页面并提出建议定价，准备接受审核。 |
|Steam 标签| 标签为应用于您的游戏并在您的商店页面可见的术语。 标签为一组重要的元数据，有助于向顾客描述游戏并让 Steam 确定如何能最好地推荐您的游戏。 |
|宣传片| 您必须为自己的产品上传宣传片，这是 Steam 发行流程的一部分。 宣传片将会显示在您的产品商店页面的顶端，通常是您的潜在顾客第一眼会看见的内容。 |
|热销商品榜单| Steam 在整个商店范围内有各种热销商品榜单，最醒目的莫过于 Steam 主页上的榜单了。 您也可以在浏览单个标签、主题、类型时找到针对某个游戏类别的热销商品榜单。 |
|更新您的游戏——最佳实践| 进行更新时可以考虑的一般实践，还有 Valve 在进行自己产品的重大更新时以及从 Steam 上的许多其他产品学习到的最佳实践。 |
|用户评测| 只要是 Steam 上拥有您的产品的游戏时间记录的用户，均可为您的产品撰写评测，并表明是否向他人推荐您的产品。 这些评测可能会出现在您的产品商店页面和 Steam 社区中，这取决于有多少其他用户觉得评测有用。 正面与负面的评测汇总用来计算评测分数，这个分数也显示在您的商店页面中，说明最近 30 天内及自产品上架以来，顾客对这款产品的看法如何。 |
|免费开玩游戏| 您可以将自己的游戏作为免费开玩游戏发行，或者，在某些情况下，也可以将您已有的游戏转换为免费开玩游戏。 此文档详细介绍了如何设置免费开玩游戏、商店可见性以及免费开玩游戏开发者可以采取的最佳实践。 |

### Steam标签



#### 产品系列页面

#### 从Steam上移除产品

#### 促销活动工具、规则及指引

#### 免费开玩游戏

#### 即将推出

#### 发行日期

#### 发行流程

#### 发行选项

#### 商店页面及其建立和编辑

#### 图像资产 - 概览

#### 在Steam上进行测试

#### 在Steam上进行预购

#### 季票

#### 定价

#### 审核流程

#### 宣传片

#### 应用程序

#### 开发者和发行商主页

#### 抢先体验

#### 更新您的游戏——最佳实践

#### 本地化和语言

#### 热销商品榜单

#### 用户测评

#### 直播

#### 管理版本

#### 蒸汽平台















