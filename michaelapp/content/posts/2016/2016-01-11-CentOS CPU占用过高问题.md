---
title: "CentOS CPU占用过高问题"
date: 2016-01-08T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---

转自：http://blog.chinaunix.net/uid-10449864-id-3463151.html

问题描述：
生产环境下的某台tomcat7服务器，在刚发布时的时候一切都很正常，在运行一段时间后就出现CPU占用很高的问题，基本上是负载一天比一天高。

问题分析：

1，程序属于CPU密集型，和开发沟通过，排除此类情况。

2，程序代码有问题，出现死循环，可能性极大。


问题解决：

1，开发那边无法排查代码某个模块有问题，从日志上也无法分析得出。

2，记得原来通过strace跟踪的方法解决了一台PHP服务器CPU占用高的问题，但是通过这种方法无效，经过google搜索，发现可以通过下面的方法进行解决，那就尝试下吧。

解决过程：

1，根据top命令，发现PID为2633的Java进程占用CPU高达300%，出现故障。

2，找到该进程后，如何定位具体线程或代码呢，首先显示线程列表,并按照CPU占用高的线程排序：

```
[root@localhost logs]# ps -mp 2633 -o THREAD,tid,time | sort -rn
```

显示结果如下：

```
USER     %CPU PRI SCNT WCHAN  USER SYSTEM   TID     TIME
root     10.5  19    - -         -      -  3626 00:12:48
root     10.1  19    - -         -      -  3593 00:12:16
```

找到了耗时最高的线程3626，占用CPU时间有12分钟了！

将需要的线程ID转换为16进制格式：
```
[root@localhost logs]# printf "%x\n" 3626
e18
```

最后打印线程的堆栈信息：
```
[root@localhost logs]# jstack 2633 |grep e18 -A 30
```

将输出的信息发给开发部进行确认，这样就能找出有问题的代码。
通过最近几天的监控，CPU已经安静下来了。
