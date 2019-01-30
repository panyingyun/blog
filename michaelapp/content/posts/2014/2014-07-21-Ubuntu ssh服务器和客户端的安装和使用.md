---
title: "Ubuntu ssh服务器和客户端的安装和使用"
date: 2014-05-15T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---

#### 1. 安装服务器端
   sudo apt-get install openssh-server
   确认是否开启ssh   ps -e| grep ssh
   如果有sshd进程表示启动了。
   配置文件  /etc/ssh/sshd_config

#### 2. 启动、停止、重启
   启动    sudo /etc/init.d/ssh  start
   停止    sudo /etc/init.d/ssh   stop
   重启    sudo /etc/init.d/ssh   restart
   退出ssh exit

#### 3. 客户端登录
   SecureCRT 5.1或者SSHSecureShellClient 都可以登录，主机名输入ip地址，选择ssh2 端口22连接后，输入ubuntu用户名和密码即可登录。

#### 4. 参考文献：

   1. http://blog.csdn.net/netwalk/article/details/12952051