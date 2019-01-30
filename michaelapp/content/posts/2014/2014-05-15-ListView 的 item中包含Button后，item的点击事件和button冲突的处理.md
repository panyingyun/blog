---
title: "ListView 的 item中包含Button后，item的点击事件和button冲突的处理"
date: 2014-05-15T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---

解决办法：
在包含button的Listview的Item Layout中加入属性 android:descendantFocusability= "blocksDescendants" 即可