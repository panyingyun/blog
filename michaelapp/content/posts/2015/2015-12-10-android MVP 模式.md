---
title: "android MVP 模式"
date: 2015-12-10T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---

### MVP模式的思考

MVP模式和MVC模式没有本质的区别，Android的Activity就是C不是V  V主要是layout XML实现的，M这个没有争议。

所以P主要是C的延伸用于解决复杂度的问题，解决Activity的设计过大的问题，有了P带来一些好处，比如原来的Activity过大，导致业务逻辑太过复杂，
所以有了P将原来的Activity的代码分为 控制显示 和 与Model对接连个部分， Activity和P之间通过接口联系，P只通过控制接口进行业务代码的编写。

很多文章提到解耦，这是不对的，其实基本上解耦的效果没有达到，只是C的代码做了拆分，在书写上跟有优势，同样引入了新的问题，源码比之前没有减少，
而且多了大量的接口。

当然还带来另外的好处就是M只与P进行交互，不再与Activity进行交互，简洁了？？ 不见得，复杂了，也么有。感觉算不上好处
第二点，P可以无需持有Activity的context,理论上可以长期存在，起到缓存的作用。在该层实现内存缓存的，减少Model层的厚度。
其实应该在该层实现业务逻辑的缓存，在model层实现是不对的。


所以P层的出现主要是解决了C模块的复杂度高的问题，解耦的效果有，但是没多少，缓存的作用有，也没多少。

其实杀手锏是在P->M的业务，见参考文献里面“最重要的老外文章”，P 和 M 做到最大的解耦，两层通过EventBut进行解耦。
由于之前拆分出来的P仅仅它和M之间通过一个热插拔的方式进行数据流交换，通过几个简单的回调实现了数据流交换的本质。

“最重要的老外文章”里面的M 有一个很大的DataManager，通过RxJava实现异步请求的级联调用，很强大，一行代码完成一个业务，个人觉得不用这个也可以，不会影响框架的质量，RxJava确实很帅，但是隐藏了逻辑的细节,而且它对大量使用常规Http请求的应用很合适，对待特殊的就没办法。



参考文献
### 1.android MVP 模式3个文章：

(1)安卓中的Model-View-Presenter模式介绍

http://www.jcodecraeer.com/a/anzhuokaifa/androidkaifa/2015/0425/2782.html

(2)我为什么不主张使用Fragment

http://www.jcodecraeer.com/a/anzhuokaifa/androidkaifa/2015/0427/2806.html

(3)最重要的老外文章

(翻译：http://my.oschina.net/mengshuai/blog/541314)

原文：https://medium.com/ribot-labs/android-application-architecture-8b6e34acda65


### 2.框架的作用：

易维护 Easy to maintain

易测试 Easy to test.

高内聚 Very cohesive.

低耦合 Decoupled.

### 3.相关库

Eventbus: https://github.com/greenrobot/EventBus

OKhttp: http://square.github.io/okhttp/



