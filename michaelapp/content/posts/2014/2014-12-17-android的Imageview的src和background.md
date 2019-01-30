---
title: "android的Imageview的src和background"
date: 2014-12-17T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---


ImageView中XML属性src和background的区别：

background会根据ImageView组件给定的长宽进行拉伸，而src就存放的是原图的大小，不会进行拉伸。src是图片内容（前景），bg是背景，可以同时使用。

此外：scaleType只对src起作用；bg可设置透明度，比如在ImageButton中就可以用android:scaleType控制图片的缩放方式，示例代码如下：

```XML
   <ImageView android:id="@+id/img" 
    android:src="@drawable/logo"
    android:scaleType="centerInside"
    android:layout_width="60dip"
    android:layout_height="60dip"
    android:layout_cen_terVertical="true"/>
```

说明：centerInside表示按比例缩放图片，使得图片长 (宽)的小于等于视图的相应维度。

注意：控制的图片为资源而不是背景，即android:src="@drawable/logo"，而非android:background="@drawable/logo"。程序中动态加载图片也类似，如：应该imgView.setImageResource(R.drawable.*);而非imgView.setBackgroundResource(R.drawable.*);



附：更详细的scaleType说明：
```XML
CENTER /center 在视图中心显示图片，并且不缩放图片

CENTER_CROP / centerCrop 按比例缩放图片，使得图片长 (宽)的大于等于视图的相应维度

CENTER_INSIDE / centerInside 按比例缩放图片，使得图片长 (宽)的小于等于视图的相应维度

FIT_CENTER / fitCenter 按比例缩放图片到视图的最小边，居中显示

FIT_END / fitEnd 按比例缩放图片到视图的最小边，显示在视图的下部分位置

FIT_START / fitStart 把图片按比例扩大/缩小到视图的最小边，显示在视图的上部分位置

FIT_XY / fitXY 把图片不按比例缩放到视图的大小显示

MATRIX / matrix 用矩阵来绘制
```