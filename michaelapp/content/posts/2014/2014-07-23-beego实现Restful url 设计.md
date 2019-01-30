---
title: "beego实现Restful url 设计"
date: 2014-07-23T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

1.  安装GO环境

2. 下载beego开发环境(windows 7下开发，Ubuntu 14.04运行）
    主页：http://beego.me/
    go get -u github.com/astaxie/beego 获取
    go install github.com/astaxie/beego 安装

    go get  -u github.com/beego/bee 获取

3.  github上建立源码库，用户存放代码
    go get github.com/panyingyun/gamebox

4.  bee创建代码
    bee create 
    bee run 

5.  restful API编写

6.  独立部署
nohup ./gamebox  > beego.log 2>&1 &

7. 如果要监控访问量，我们可以通过这个实现
//基本配置
appname = gamebox
httpport = 8080
httpaddr= "192.168.1.10"
runmode = dev
autorender = false
copyrequestbody = true

//访问量
EnableAdmin = true
AdminHttpAddr = "192.168.1.10"
AdminHttpPort = 9090

参考这篇文章
http://my.oschina.net/astaxie/blog/143475

基本的HTTP服务器和客户端测试程序
http://blog.csdn.net/xiazdong/article/details/7724349