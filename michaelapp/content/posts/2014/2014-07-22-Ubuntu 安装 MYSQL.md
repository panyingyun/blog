---
title: "Ubuntu 安装 MYSQL"
date: 2014-07-22T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---

#### 1. 安装MySQL

sudo apt-get install mysql-server

#### 2. 修改MYSQL  root用户密码
默认的MySQL安装之后根用户是没有密码的，所以首先用根用户进入：
$mysql -u root -p
mysql>use mysql; update user set password=PASSWORD('panpan') where user='root';flush privileges;
修改密码为 panpan

#### 3. 修改MYSQL root用户远程访问权限
首先修改配置文件 sudo vim /etc/mysql/my.cnf
bind-address=127.0.0.1 => bind-address= 你机器的IP
重启mysql  
sudo /etc/init.d/mysql restart

本机登陆mysql：mysql -u root -p 
然后输入密码 panpan
use mysql;
grant all privileges on *.* to root@"%" identified by "panpan" with grant option;flush privileges;
进mysql库查看host为%的数据是否添加：use mysql; select * from user;

这样就可以使用navicat 进行访问登录了 3306 /192.168.1.21 /root /panpan 

#### 4. 增加新用户panyingyun 并创建数据库mame，让它只能对mame数据库进行访问
本机登陆mysql：mysql -u root -p  
mysql>create database mame;
mysql>grant all privileges on mame.* to panyingyun@"%" identified by "panpan" with grant option;flush privileges;
这样就可以使用navicat 进行访问登录了 3306 /192.168.1.21 /panyingyun /panpan 

#### 5. 参考文献：

1. http://wiki.ubuntu.org.cn/MySQL%E5%AE%89%E8%A3%85%E6%8C%87%E5%8D%97
2. http://www.cnblogs.com/wuhou/archive/2008/09/28/1301071.html
3. http://tech.sina.com.cn/s/s/2008-12-24/09322685698.shtml
4. http://www.cnblogs.com/mr-wid/archive/2013/05/09/3068229.html#c1（21分钟 MySQL 入门教程）
