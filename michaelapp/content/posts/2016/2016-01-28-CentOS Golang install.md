---
title: "CentOS Golang install"
date: 2016-01-28T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---


1. 下载和安装golang
	
	存放在/home/go目录下
	
	```
	#] wget http://golangtc.com/static/go/go1.5.3/go1.5.3.linux-amd64.tar.gz

	#] 7za x go1.5.3.linux-amd64.tar.gz

	#] 7za x go1.5.3.linux-amd64.tar 
		
	```

2. 配置环境变量
	
	```
	#] vi /etc/profile
	```

	在文件末尾添加：
	
	```
	export PATH=$PATH:/home/go/bin
	export GOROOT=/home/go
	export GOPATH=/home/gopro
	```
	
	使配置生效
	
	```
	#]source /etc/profile
	```

3. 测试一下

	```
	#] go version
	```