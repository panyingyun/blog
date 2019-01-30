---
title: "Android Image Loader 第三方库对比测试"
date: 2015-12-04T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---

Android Image Loader 第三方库对比测试 <https://www.zybuluo.com/linux1s1s/note/135758>

通过对比，我们发现Fresco的内存模型和其他不同，加载图片的内存分配在native heap
Fresco+OkHttp 加载速度优势较为明显，值得研究

```Shell
Fresco: 168ms
Glide: 321ms
Universal Image Loader:64ms
Volley: 137ms
Fresco+OkHttp:  71ms
Picasso: 103ms
```

