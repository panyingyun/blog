---
title: "Android app的文件缓存目录"
date: 2014-12-19T10:15:35+08:00
draft: false
tags: [
    "android"
]
categories: [
    "android",
]
---

Android app的文件缓存目录
可以是app内置私有的目录，
当然也可以选择外置sdcard目录

#### 内置私有目录

1. /data/data/[packagename]/files 文件缓存目录,一般存小的文件缓存，如果是图片，不建议放这里，一般放到外置卡。
    
    File file = getFilesDir(); 返回该目录
   写文件到该目录下可以像这样
```java
FileOutputStream fos = null;
		try {
			fos = appontext.openFileOutput(name, Context.MODE_PRIVATE);
			fos.write(content.getBytes());
		} catch (Exception e) {
		}finally{
			try {
				if(null != fos){
					fos.close();
				} 
			}catch (Exception e) {
			}
		}
```
2.  /data/data/[packagename]/cache目录，存放一些其他缓存  File cache = getCacheDir(); 

3.  /data/data/[packagename]/databases，存放数据库

4.   /data/data/[packagename]/lib，应用的so目录

5.   /data/data/[packagename]/shared_prefs 应用的SharedPreferences保存

可以自己创建其他目录吗？ 可以的
使用 File ownDataPath = getDir("service",Context.MODE_PRIVATE);
使用它可以创建app_service目录，放什么自己定义
		
#### 外置SDCARD目录

1. 外置缓存目录（File sdcache = getExternalCacheDir();）
 /storage/emulated/0/Android/data/[packagename]/cache
一些重要性不高的cache或者大文件放到这里，比如图片缓存

2. 外置文件缓存目录（File sdfile = getExternalFilesDir(null);），
 /storage/emulated/0/Android/data/[packagename]/files
一些重要性不高的file cache或者大文件放到这里
注意： /storage/emulated/0/Android/data/[packagename] 在android2.2之后，在应用卸载后也会一并卸载。所以不需要用什么清理缓存的软件清理的。

#### PS: 现在手机支持外置挂载T卡的，如何访问外置T卡目录呢？

1. 可以参考 [!http://stackoverflow.com/questions/5694933/find-an-external-sd-card-location](http://stackoverflow.com/questions/5694933/find-an-external-sd-card-location)