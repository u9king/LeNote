# Hexo基本指令

#### 1.介绍

Hexo是一款基于Node.js的静态博客框架，依赖少易于安装使用，可以方便的生成静态网页托管在GitHub和Heroku上。因为Hexo的创建者是台湾人，对中文的支持很友好，可以选择中文进行查看。

#### 2.Hexo一句提交

```bash
hexo clean && hexo g -d
```


#### 3.基本指令介绍

|        指令         |     作用     |       简写       |                             参数                             |
| :-----------------: | :----------: | :--------------: | :----------------------------------------------------------: |
|     hexo clean      |   清除缓存   |                  |                                                              |
|    hexo generate    |     生成     |      hexo g      |                                                              |
|     hexo deploy     |  部署到远端  |      hexo d      |                                                              |
|     hexo server     | 启动服务预览 |      hexo s      | hexo s -p 5000：更改端口<br />hexo s -i 192.168.1.1自定义IP<br />hexo s -s 静态模式 |
|      hexo init      |  初始化博客  |                  |                                                              |
| hexo new "我的博客" |   新建文章   | hexo n "我的博客 |                                                              |
| npm update hexo -g  |     升级     |                  |                                                              |

#### 4.HTTP提交失败

可以改用ssh来提交

```bash
hexo config deploy.type git
hexo config deploy.repo git@github.com:u9king/u9king.github.io.git
```

#### 5.创建Hexo博客步骤

- 在github上创建仓库`github用户名.http://github.io`

- 搭建git与github的ssh

- 安装nodejs

- 安装hexo

  ```
  npm install -g hexo-cli 
  ```

- 创建文件夹在其中

  ```
  hexo init blog
  ```

- 在文件夹中初始化博客项目

  ```
  hexo new [博客名称]
  hexo generate #生成
  hexo server #启动服务预览
  然后就能在本地浏览器http://localhost:4000/
  ```

- 关联github项目

  ```
  在blog根目录里的_config.yml也就是站点配置文件中
  找到# Deployment并对应修改为以下内容
  deploy:
      type: git
      repo: https://github.com/[github用户名]/[github用户名].github.io.git
      branch: main
  ```

- 安装Hexo的Git部署插件

  ```
  npm install hexo-deployer-git --save
  ```

- 部署到远端

  ```
  hexo clean 
  hexo generate #生成
  hexo deploy #部署
  ```

#### 6.目录

在项目的`._config.yml`文件中Directory对应参数的解释

| 参数           | 描述                                                         | 默认值           |
| :------------- | :----------------------------------------------------------- | :--------------- |
| `source_dir`   | 资源文件夹，这个文件夹用来存放内容。                         | `source`         |
| `public_dir`   | 公共文件夹，这个文件夹用于存放生成的站点文件。               | `public`         |
| `tag_dir`      | 标签文件夹                                                   | `tags`           |
| `archive_dir`  | 归档文件夹                                                   | `archives`       |
| `category_dir` | 分类文件夹                                                   | `categories`     |
| `code_dir`     | Include code 文件夹，`source_dir` 下的子目录                 | `downloads/code` |
| `i18n_dir`     | 国际化（i18n）文件夹                                         | `:lang`          |
| `skip_render`  | 跳过指定文件的渲染。匹配到的文件将会被不做改动地复制到 `public` 目录中。您可使用 [glob 表达式](https://github.com/micromatch/micromatch#extended-globbing)来匹配路径。 |                  |

#### 7.国际化（i18n）

若要让您的网站以不同语言呈现，您可使用国际化（internationalization）功能。请先在 `_config.yml` 中调整 `language` 设定，这代表的是预设语言，您也可设定多个语言来调整预设语言的顺位。

```
language: zh-tw

language: 
- zh-tw
- en
```

语言文件

语言文件可以使用 YAML 或 JSON 编写，并放在主题文件夹中的 `languages` 文件夹。您可以在语言文件中使用 [printf 格式](https://github.com/alexei/sprintf.js)。

路径

您可在 front-matter 中指定该页面的语言，也可在 `_config.yml` 中修改 `i18n_dir` 设定，让 Hexo 自动侦测。

```
i18n_dir: :lang
```

`i18n_dir` 的预设值是 `:lang`，也就是说 Hexo 会捕获网址中的第一段以检测语言，举例来说：

```
/index.html => en
/archives/index.html => en
/zh-tw/index.html => zh-tw
```

捕获到的字符串唯有在语言文件存在的情况下，才会被当作是语言，因此例二 `/archives/index.html` 中的 `archives` 就不被当成是语言。



#### 疑问：

1.Hexo对于标签tag的定义是什么？和分类之间有什么关系

2.`skip_render`不渲染是指Markdown文件不渲染吗？

3.Hexo只支持markdown的文件吗？

4.如何修改网站图标icon