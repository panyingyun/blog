---
title: "android app混淆问题"
date: 2014-12-24T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---


如果使用到了泛型和反射 务必加上下面两句 
```Java
-keepattributes Annotation 
-keepattributes Signature
```
#### 参考： 
http://charles-tanchao.diandian.com/post/2012-05-24/20118715