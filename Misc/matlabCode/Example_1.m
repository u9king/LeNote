%% MATLAB数据类型 
%数字、字符与字符串、矩阵、元胞数组、结构体
%字符与字符串
s = 'a';
abs(s)        %取ASCII码
char(65)      %ASCII码转字符
num2str(65)   %数字转字符串
doc num2str   %查询函数作用
%矩阵
A = [1 2 3;4 5 2;3 2 7];
B = A';      %矩阵转置
C = A(:);    %拉伸为列向量
D = inv(A);  %逆矩阵
A * D        %A矩阵乘以逆矩阵等于单位矩阵

E = zeros(10,5,3);         %10行5列3通道全0矩阵
E(:,:,1) = rand(10,5)      %rand生成0到1之间的10行5列均匀分布的伪随机数
E(:,:,2) = randi(5,10,5)   %randi生成均匀分布的伪随机整数，randi(iMax,m,n)，生成0到iMax之间的整数
E(:,:,3) = randn(10,5)     %randn生成标准正太分布的伪随机数

%元胞数组
F = cell(1,6)    %定义1行6列的元胞数组  ！！MATLAB是从1开始的
F{2} = eye(3)    %eye生成3*3的单位矩阵
F{5} = magic(5)  %magic生成5阶幻方
G = F{5}

%结构体  类字典Python中
books = struct('name',{{'Machine Learning','Data Mining'}},'price',[30,40])
books.name      %用键获取值
books.name(1)   %取出的是cell
books.name{1}   %取出的是string

%% MATLAB矩阵操作
%矩阵的定义与构造
A = [1 2 3 5 8 5 4 6];
B = 1:2:9;                  %生成1到9之间步长为2的矩阵
C = repmat(B,3,1);          %repmat生成B矩阵重复3行1列
D = ones(2,4);              %ondes生成2行4列的全1矩阵
%矩阵的四则运算
A = [1 2 3 4; 5 6 7 8]
B = [1 1 2 2; 2 2 1 1]
C = A + B     %对应项相加
D = A - B     %对应项相减
E = A * B'    %A乘B转置
F = A .* B    %对应项相乘，所有用.的都是对应项
G = A / B     %A乘B的逆
H = A ./ B    %对应项相除
%矩阵的下标
A = magic(5)
B = A(2,3)    %取第2行第3列的值
C = A(3,:)    %取第3行的所有值
D = A(:,4)    %取第4列的所有值
[m,n] = find(A > 20)  %找大于20的序号值/矩阵

%% MATLAB逻辑与流程控制
%1.if...else...end
a = 100;
b = 20;
if a > b
    "成立"
else
    "不成立"
end
%2.for...end
sum = 0;
for n = 1:5
    sum = sum + n^2;
end
%3.while...end
s = 0;
n = 1;
while n <= 10
    s = s + n;
    n = n + 1;
end
%4.switch...case...end
a = 100;
switch a 
    case 1
        "One"
    case 10
        "Ten"
    case 100
        "Hundred"
    otherwise
        "Unknow"
end

%% MATLAB基本绘图操作
%二维平面绘图
x = 0:0.01:2*pi;    %生成0到2pi差为0.01的等差数列
y = sin(x);
figure              %生成画板
plot(x,y)           %绘制x,y
title('y=sin(x)')   %图片标题
xlabel('x')         %x轴标签
ylabel('sin(x)')    %y轴标签
xlim([0 2*pi])      %x轴限制范围为[0,2pi],解决matlab右侧空白问题

x = 0:0.01:20;
y1 = 200*exp(-0.05*x).*sin(x);
y2 = 0.8*exp(-0.5*x).*sin(10*x);
figure
[AX,H1,H2] = plotyy(x,y1,x,y2,'plot');              %plotyy公用一个x坐标系
set(get(AX(1),'Ylabel'),'String','Slow Decay')      %在y轴上设置左标签
set(get(AX(2),'Ylabel'),'String','Fast Decay')      %在y轴上设置右标签
xlabel('Time (\musec)')
title('Multiple Decay Rates')
set(H1,'LineStyle','--')    %设置虚线                            
set(H1,'LineStyle',':')     %设置冒号线

%三维立体绘图
t = 0:pi/50:10*pi;
plot3(sin(t),cos(t),t)  %plot3三维绘图x,y,z
xlabel('sin(t)')
ylabel('cos(t)')
zlabel('t')
grid on         %开启网格
axis square     %正交视图

%图形窗口的分割
x = linspace(0.2*pi,60);
subplot(2,2,1)      %将画板分为2行2列，子图在其中的第1个格
plot(x,sin(x)-1);
title('sin(x)-1');
axis([0,2*pi,-2,0])
subplot(2,1,2)      %将画板分为2行2列，子图在其中的第2个格
plot(x,cos(x)+1)
title('cos(x)+1');
axis([0,2*pi,0,2])
subplot(4,4,3)      %将画板分为4行4列，子图在其中的第3个格
plot(x,tan(x));
title('tan(x)');
axis([0,2*pi,-40,40])
subplot(4,4,8);      %将画板分为4行4列，子图在其中的第8个格
plot(x,cot(x));
title('cot(x)');
axis([0,2*pi,-35,35])

%% 图形的保存与导出
%(1) Figure视窗下→编辑→复制图形(窗)
%(2) Figure视窗下→文件→导出设置

%% 补充内容
[x,y,z] = peaks(30);    %peaks生成双峰函数
mesh(x,y,z)
grid


%% 清空环境变量及命令
clear all  % 清除Workspace中的所有变量
clc        % 清除Command Window中的所有命令