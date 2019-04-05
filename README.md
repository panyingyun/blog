# blog
Caddy+hugo搭建自己的blog(MichaelApp)

https://www.michaelapp.com/

![效果图](https://www.michaelapp.com/posts/2019/hugo%E6%90%AD%E5%BB%BA%E8%87%AA%E5%B7%B1%E7%9A%84blog/homepage.png)


1、下载安装hugo工具
https://github.com/gohugoio/hugo/releases 

2、下载govendor工具
go get -u github.com/kardianos/govendor

3、下载go-bindata工具
go get -u github.com/jteeuwen/go-bindata/...

发布文章
cd michaelapp
hugo new posts/2019/cmake学习.md

服务器更新文章
cd /home/gopro/src/blog
./all.sh

这样blog就更新完成了

