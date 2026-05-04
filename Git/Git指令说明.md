# Git指令说明

#### 1.提交代码

```
git add .
git commit -m "xxx"
git push
```

#### 2.ssh查询提交

```
git remote -v
ssh -T git@github.com
ssh-add C:\\Users\\Administrator\\.ssh\\id_rsa
ssh -T git@github.com

eval "$(ssh-agent -s)"
```

#### 3.ssh连接失败

```
//在~/.ssh下创建config输入以下内容
Host github.com
  Hostname ssh.github.com
  Port 443
```

#### 4.查询config是否正确

```
git config --global --list
```

#### 5.文件打包

```
//查看文件占用情况
git count-objects -vH
//将松散对象打包压缩
git gc --prune=now --aggressive
```

#### 6.查询状态

```
//查看当前分支
git status
```

#### 7.克隆仓库

```
git clone XXXXX
```

#### 8.查看分支

```
//-r 表示远程仓库所有分支  //-a 表示远程和本地的所有分支
git branch -r
git branch -a
```

#### 9.创建分支

```
git checkout -b <branch-name>
```

#### 10.切换分支

```
git checkout <branch-name>
```

























