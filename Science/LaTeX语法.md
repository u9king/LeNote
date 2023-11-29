# LaTeX语法

#### 1.论文排版常用语句

```
\title{文章标题}

\begin{document}

\maketitle		创建标题

\begin{abstract} 	摘要

\keywords{关键词1\quad 关键词2}

\end{abstract}

\tableofcontents		目录

\newpage		新的一页

\section{一级标题}

\subsection{二级标题}

\end{document}
```

#### 2.字体

```
字体选择（字体声明，用于声明后面的字体为罗马字体，无衬线字体，打字机字体，大括号用于分组，限定字体限制范围）
\rmfamily Roman Family{\sffamily Scans Serif Family}{\ttfamily Typewriter Family}
加粗
\textbf{Boldface Series}
字体切换
{\songti 宋体}	{\heiti 黑体}	{\kaishu 楷书}
字体大小
{\tiny Hello}
{\large Hello}
{\Huge Hello}
{\zihao{4} 你好！}
```

#### 3.公式

```
$ W > W_0 $ 行内公式
$$ W >W_0 $$ 行间公式
```

#### 4.特殊符号

符号前需要加上转义符号\

| $\div$    |   `\div`   | $\times$  |  ` \times`    |  $\approx$  |  `\approx`  |
| --------------- | ---- | ------------------ | ---- | --------------- | --------------- |
| $\cdot$  |   `\cdot`   | $\neq$      |   `\neq`    |   $\gt$   |   `\gt`   |
| $\geq$    |   `\geq`   | $\leq$       |   `\leq`   |   $\lt$   |   `\lt`   |
| $\pm$     |   `\pm`    | $\pi$         |   `\pi`   |   $\infty$   |   `\infty`   |
| $\geqslant$ | `\geqslant` | $\quad$ | `\quad` | $\LaTeX$ | `\LaTeX` |
| $\#$ | `\#` | $\%$ | `\%` | $\$$ | `\$` |
| $\triangle$ | `\triangle` | $\bigcirc$ | `\bigcirc` | $\theta$ | `\theta` |
| $\overbrace{1+2+\cdots+100}$ | `\overbrace{1+2+\cdots+100}` | $\because$ | `\because` | $\therefore$ | `\therefore` |
| $\overbrace{a+b+c}^{n}$ | `\overbrace{a+b+c}^{n}` | $\alpha$ | `\alpha` | $\beta$ | `\beta` |
| $\ce{CO2 + C -> 2 CO}$ | `\ce{CO2 + C -> 2 CO}` | $\ce{h2o}$ | `\ce{h2o}` |  |  |

#### 5.排版

| $\frac{2}{3}$         | 默认     |
| --------------- | ---- |
| $\large{\frac{2}{3}}$ | `\large` |
| $\Large{\frac{2}{3}}$ | `\Large` |
| $\huge{\frac{2}{3}}$  | `\huge`  |
| $\Huge{\frac{2}{3}}$  | `\Huge`  |

#### 6.方程组

方程组写法
$$
\begin{cases}
2x-y&=5\\
3x+4y&=2
\end{cases}
$$
列式计算写法
$$
\begin{align}
&am+an+bm+bn\\
=&a(m+n)+b(m+n)\\
=&(m+n)(a+b)
\end{align}
$$


#### 98.论文排版

```
\begin{document}
	\chapter{绪论}
	\section{研究的目的和意义}
	\section{国内外研究现状}
	\subsection{国外研究现状}
    \subsection{国内研究现状}
    \section{研究内容}
    \section{研究方法和技术路线}
    \subsection{研究方法}
    \subsection{研究路线}
    \chapter{实验与结果分析}
    \section{引言}
    \section{实验方法}
    \section{实验结果}
    \subsection{数据}
    \subsection{图表}
    \subsubsection{实验条件}
    \subsubsection{实验过程}
    \subsection{结果分析}
    \section{结论}
    \section{致谢}
\end{document}
```







#### 99.参考网站

1. LaTeX语法（https://superior-leo.gitee.io/2021/01/30/latex-xue-xi-bi-ji/）

