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

- 提示：上下标$3x^{22}还有a_{11}$需要添加{}，`3x^{22}`和`a_{11}`
- 组合数表示法$\mathbf{C^{22}_{11}}$可以通过`\mathbf{C^{22}_{11}}`

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

- 放大缩小
| $\frac{2}{3}$         | 默认     |
| --------------- | ---- |
| $\large{\frac{2}{3}}$ | `\large` |
| $\Large{\frac{2}{3}}$ | `\Large` |
| $\huge{\frac{2}{3}}$  | `\huge`  |
| $\Huge{\frac{2}{3}}$  | `\Huge`  |
#### 6.方程组

- 方程组写法

$$
\begin{cases}
2x-y&=5\\
3x+4y&=2
\end{cases}
$$
- 列式计算写法

$$
\begin{align}
&am+an+bm+bn\\
=&a(m+n)+b(m+n)\\
=&(m+n)(a+b)
\end{align}
$$

- 带编号的公式

  ```
  \begin{equation}		%编号为(1)
  a + b = b + a
  \end{equation}
  
  \begin{equation}		%编号为(2)
  a + b = b + a
  \end{equation}
  
  \begin{equation*}		%不带编号
  a + b = b + a
  \end{equation}
  ```

#### 7.表格

- 普通表格（居中使用嵌套`\begin{center}`）

  ```
  \begin{tabular}{|l|c|c|c|p{1.5cm|}	    	%|：竖线 l:左对齐 c:居中 p:{自定宽度}
   \hline										%hline:横线
   姓名 & 语文 & 数学 & 外语 & 备注\\			  %&:对齐符号
   \hline
   张三 & 87 & 100 & 93 & 优秀\\ 
   \hline
   李四 & 75 & 64 & 52 & 补考另行通知\\
    \hline
  \end{tabular}
  ```

- 三线表

```
\begin{tabular}{cc}		%cc中间没有线														
 \toprule[1.5pt]		%上横线								
 \makebox[0.3\textwidth][c]{符号}	&  \makebox[0.4\textwidth][c]{意义} \\
 \midrule[1pt]
 $ W $	    	& 某一小时内该路段运行总收益-总成本   \\ 
 $ W_0 $	    & 区分高峰和低峰的一个临界值  \\ 
 $ P $	    	& 线路在一小时内所有站的总上车人数 \\ 
\bottomrule[1.5pt]
\end{tabular}
```

#### 8.列表

```
\begin{itemize}
\item 假设该线路上运行的是同一种类型的公交车；
\item 假设运行的是不同种类型的该线路；
\item 该线路上运行的是可能钟类型的公交车；
\end{itemize}
```

#### 9.图片

- 插入图片

```
%引用图片
\usepackage{float}					%用于排版图片位置
\includegraphics{排队论模型.png}
\begin{figure}[H]					%[H]是浮动体的类型
	\centering
	\includegraphics[scale=0.4]{排队论模型.png}
    \caption{问题一模型示意图}		   %标题
    \label{paiduimx}				%为了引用取名，后续使用`图\ref{paiduimx}`生成图1语句
\end{figure}
```

- 图片排版（四宫格）

```
\begin{figure}[H]
	\caption{高峰区到低峰区的调度示意图}
	\label{diaoduhou1}
	\subfigure
	{
		\begin{minipage}[b]{.3\linewidth}
			\centering
			\includegraphics[scale=0.25]{调度1.png}
		\end{minipage}
	} \quad \quad \quad \quad \quad \quad \quad 
	\subfigure
	{
		\begin{minipage}[b]{.3\linewidth}
			\includegraphics[scale=0.25]{调度2.png}
		\end{minipage}
	}

	\subfigure
	{
		\begin{minipage}[b]{.3\linewidth}
			
			\includegraphics[scale=0.25]{调度3.png}
		\end{minipage}
	} \quad \quad \quad \quad \quad \quad \quad
	\subfigure
	{
	\begin{minipage}[b]{.3\linewidth}
		\includegraphics[scale=0.25]{调度4.png}
	\end{minipage}
	}
\end{figure}
```

#### 10.矩阵

```
a = \left(				%这里可以任意修改成(,{,[
\matrix{
	\alpha_1 & test1\\
	\alpha_2 & test2
}
\right)
```

$$
a = \left(
\matrix{
	\alpha_1 & test1\\
	\alpha_2 & test2
}
\right) \\
\text{a: 矩阵}
$$

- 也可以使用`matrix`,小括号矩阵`pmatrix`,方括号矩阵`bmatrix`,中括号矩阵`Bmatrix`,单竖线矩阵`vmatrix`,双竖线矩阵`Vmatrix`

$$
\begin{matrix}
    0 & 1 \\
    1 & 0
\end{matrix}
\qquad
\begin{pmatrix}
    0 & 1 \\
    1 & 0
\end{pmatrix}
\qquad
\begin{bmatrix}
    0 & 1 \\
    1 & 0
\end{bmatrix}
\qquad
\begin{Bmatrix}
    0 & 1 \\
    1 & 0
\end{Bmatrix}
\qquad
\begin{vmatrix}
    0 & 1 \\
    1 & 0
\end{vmatrix}
\qquad
\begin{Vmatrix}
    0 & 1 \\
    1 & 0
\end{Vmatrix}
$$

- 多项矩阵，斜点
  $$
  \begin{pmatrix}
  a_{11} & a_{12} & \dots & a_{1n} \\
  a_{21} & a_{22} & \dots & a_{2n} \\
  \vdots & \vdots & \ddots & \vdots \\
  a_{n1} & a_{n2} & \dots & a_{nn} \\
  \end{pmatrix}
  $$

- 分块矩阵
  $$
  \begin{pmatrix}
  	\begin{matrix}
  	1 & 0 \\
  	0 & 1 \\
  	\end{matrix} & \Large 0 \\
  	 \Large 0 & \begin{matrix}
  	1 & 0 \\
  	0 & 1 \\
  	\end{matrix}
  \end{pmatrix}
  $$

- 三角矩阵
  $$
  \begin{pmatrix}
  	a_{11} & a_{12} & \dots & a_{1n} \\
  	& a_{22} & \cdots & a_{2n} \\
  	& & \ddots & \vdots \\
  	& & & a_{nn}
  \end{pmatrix}
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





#### 疑问

1.图片存放在哪里才能索引到

#### 99.参考网站

1. LaTeX语法（https://superior-leo.gitee.io/2021/01/30/latex-xue-xi-bi-ji/）
2. 表格生成网站：https://tablesgenerator.com/
3. 拍照公式网站：https://mathpix.com/

