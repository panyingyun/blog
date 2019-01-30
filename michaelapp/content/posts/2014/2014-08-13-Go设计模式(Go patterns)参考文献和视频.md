---
title: "Go设计模式(Go patterns)参考文献和视频"
date: 2014-08-13T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

并发是Golang中比较复杂也是困难的部分，设计上需要考虑是否死锁和goroutine是否未推出导致一直占有资源等问题
这里给出了一些优秀的文章链接供参考。

#### 参考文献
1. RobPike在Google IO 2012大会 上 关于并发的介绍（DOC链接和视频链接） 

（1）https://talks.golang.org/2012/concurrency.slide#1 

（2）https://www.youtube.com/watch?v=f6kdp27TYZs 

（3）优酷视频   http://v.youku.com/v_show/id_XNDI1NjgxMTAw.html


#### 2. Advanced Go Concurrency Patterns（DOC链接和视频链接） 

（1）http://talks.golang.org/2013/advconc.slide#1

（2）http://blog.golang.org/advanced-go-concurrency-patterns

（3） 优酷视频   http://v.youku.com/v_show/id_XNTcyMTA4MTM2.html


#### 3. Go Concurrency Patterns: Pipelines and cancellation

http://blog.golang.org/pipelines

中文译

http://air.googol.im/2014/03/15/go-concurrency-patterns-pipelines-and-cancellation.html

#### 4. Go Concurrency Patterns: Context

http://blog.golang.org/context

src:https://github.com/gorilla/context


#### 5. Douglas McIlroy的论文《一窥级数数列》展示了Go使用的这类并发技术是如何优雅地支持复杂计算

文章地址：http://swtch.com/~rsc/thread/squint.pdf