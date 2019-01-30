---
title: "Caddy替换Nginx"
date: 2017-10-19T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---


### 1. 下载和安装Caddy
* 官方网站  https://caddyserver.com/download 下载对应的二进制包,CentOS系统可以使用脚本 
	
	```
	curl https://getcaddy.com | bash -s personal http.authz,http.awse
	s,http.awslambda,http.cache,http.cgi,http.cors,http.datadog,http.expires,http.filemanager,ht
	tp.filter,http.forwardproxy,http.git,http.gopkg,http.grpc,http.hugo,http.ipfilter,http.jekyl
	l,http.jwt,http.login,http.mailout,http.minify,http.nobots,http.prometheus,http.proxyprotoco
	l,http.ratelimit,http.realip,http.reauth,http.restic,http.upload,http.webdav
	```
	
* shell输出下面输出即表示安装成功
	
	```
	Downloading Caddy for linux/amd64 (personal license)...
	Download verification OK
	Extracting...
	Putting caddy in /usr/local/bin (may require password)
	Caddy 0.10.10 (non-commercial use only)
	Successfully installed
	```
	
* 执行下面shell命令

	参考：https://github.com/mholt/caddy/blob/master/dist/init/linux-systemd/README.md
	
	设置完成 /etc/caddy/Caddyfile 和 /etc/ssl/caddy 通过
	
	```shell
	systemctl daemon-reload
	systemctl start caddy.service  //启动
	systemctl enable caddy.service  //开机启动
	systemctl stop caddy.service   //停止
	systemctl restart caddy.service  //重启
	systemctl status caddy.service -l  //查看状态
	```
		
### 2. 配置 Caddyfile 

* 参考官方网站：https://caddyserver.com/tutorial/caddyfile

* nginx to caddyfile https://www.youtube.com/watch?time_continue=28&v=nk4EWHvvZtI

* 反向代理配置 https://caddyserver.com/docs/proxy

	比如本人博客域名到内部服务反向代理：
	
	```shell
	http://www.michaelapp.com {
	   proxy / 127.0.0.1:2000 
	}
	
	https://www.michaelapp.com {
	    proxy / 127.0.0.1:2000
	}
	
	http://books.michaelapp.com {
	   proxy / 127.0.0.1:7000
	}
	```

### 3. 运行

* systemctl restart caddy.service
	
### 4. 参考文献
* https://caddyserver.com/download


