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

