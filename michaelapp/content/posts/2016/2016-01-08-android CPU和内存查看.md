---
title: "android CPU和内存查看"
date: 2016-01-08T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---

### 1. CPU查看

查看CPU排行
```
top -m 5 
```

```shell
  PID PR CPU% S  #THR     VSS     RSS PCY UID      Name
12695  0  33% S    67 925692K 119092K  bg u0_a79   com.dangbeimarket
21660  3   2% S    53 927132K 115780K  fg u0_a89   cn.com.wasu.main
 1987  1   0% S    43 867652K  53132K  bg u0_a82   com.starcor.mango
22753  3   0% R     1   1312K    488K  fg shell    top
```
或者
```
dumpsys cpuinfo cn.com.wasu.main
```

### 2. 内存查看
```
dumpsys meminfo cn.com.wasu.main
```

```shell
shell@MagicBox2:/ $ dumpsys meminfo cn.com.wasu.main
Applications Memory Usage (kB):
Uptime: 43600223 Realtime: 188044672

** MEMINFO in pid 21660 [cn.com.wasu.main] **
                   Pss  Private  Private  Swapped     Heap     Heap     Heap
                 Total    Dirty    Clean    Dirty     Size    Alloc     Free
                ------   ------   ------   ------   ------   ------   ------
  Native Heap        0        0        0        0    10260     7753      494
  Dalvik Heap        0        0        0        0   112335   108928     3407
 Dalvik Other     1359     1356        0        0
        Stack      352      352        0        0
       Ashmem   108746   108556        0        0
    Other dev        9        0        8        0
     .so mmap     3011      708     1628        0
    .apk mmap      158        0       68        0
    .ttf mmap      851        0      384        0
    .dex mmap     2565        0     1952        0
   Other mmap       41        4        8        0
      Unknown     7027     7020        0        0
        TOTAL   124119   117996     4048        0   122595   116681     3901
```
