---
title: "Windows 7安装ST3, 搭建Go开发环境，Markdown编辑环境"
date: 2014-04-14T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---

#### 第一步：Go环境的安装 直接使用官方的安装包就可以了  

#### 第二步： 下载ST3
[下载ST3 http://www.sublimetext.com/3](http://www.sublimetext.com/3)  

#### 第三步： 安装插件  

首先，我们需要安装 Sublime Text 的 Package Control 功能，在打开软件后，按下快捷键 Ctrl+\`，（\`这个符号为英文半角模式下，按下 Tab 键上方、数字键1左边的那个按键），此时会打开一个命令窗口，复制并输入以下内容，最后回车：  

```python
	import urllib.request,os,hashlib; h = '7183a2d3e96f11eeadd761d777e62404' + 'e330c659d4bb41d3bdf022e94cab3cd0'; pf = 'Package Control.sublime-package'; ipp = sublime.installed_packages_path(); urllib.request.install_opener( urllib.request.build_opener( urllib.request.ProxyHandler()) ); by = urllib.request.urlopen( 'http://sublime.wbond.net/' + pf.replace(' ', '%20')).read(); dh = hashlib.sha256(by).hexdigest(); print('Error validating download (got %s instead of %s), please try manual install' % (dh, h)) if dh != h else open(os.path.join( ipp, pf), 'wb' ).write(by)
```

注：如果出现运行错误，可以参考 [https://sublime.wbond.net/installation#st3](https://sublime.wbond.net/installation#st3)**右侧的说明**来解决。

#### 第四步：安装插件

按快捷键 Ctrl+Shift+p 键入 pci，选择第一个Install Package:

```python
1. GoSublime  

2. ConvertToUTF8  

3. SideBarEnhancements  

4. Markdown Preview  
```
等几个插件

如果还有不清楚的地方请参考：  

1. [Ubuntu 配置 Go 语言开发环境(Sublime Text+GoSublime)](http://my.oschina.net/Obahua/blog/110767)
2. [Windows搭建Sublime Text 3 + Go开发环境](http://blog.csdn.net/cyxcw1/article/details/10329481)
3. [Ubuntu 将 Sublime Text 添加到 Launcher 和其它方式](http://my.oschina.net/Obahua/blog/110612)
4. [Ubuntu 使用 Sublime Text 作为Go语言源码的默认启动程序](http://my.oschina.net/Obahua/blog/111079)
