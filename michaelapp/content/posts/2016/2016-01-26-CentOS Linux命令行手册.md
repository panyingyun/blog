---
title: "CentOS Linux命令行手册"
date: 2016-01-26T10:15:35+08:00
draft: false
tags: [
    "linux"
]
categories: [
    "linux",
]
---

转自：http://blog.jobbole.com/97626/

```
#yum install python
#yum install python-pip
#yum install git
#pip install docopt pygments
#git clone https://github.com/chrisallenlane/cheat.git

#cd cheat
#python setup.py install
cheat -v
cheat 2.0.9

```
使用举例1：ls命令使用
```
[root@iZ28s8lydmdZ cheat]# cheat ls
# Displays everything in the target directory
ls path/to/the/target/directory

# Displays everything including hidden files
ls -a

# Displays all files, along with the size (with unit suffixes) and timestamp
ls -lh 

# Display files, sorted by size
ls -S

# Display directories only
ls -d */

# Display directories only, include hidden
ls -d .*/ */
```
使用举例1：tar命令使用
```
[root@iZ28s8lydmdZ cheat]# cheat tar
# To extract an uncompressed archive:
tar -xvf /path/to/foo.tar

# To create an uncompressed archive:
tar -cvf /path/to/foo.tar /path/to/foo/

# To extract a .gz archive:
tar -xzvf /path/to/foo.tgz

# To create a .gz archive:
tar -czvf /path/to/foo.tgz /path/to/foo/

# To list the content of an .gz archive:
tar -ztvf /path/to/foo.tgz

# To extract a .bz2 archive:
tar -xjvf /path/to/foo.tgz

# To create a .bz2 archive:
tar -cjvf /path/to/foo.tgz /path/to/foo/

# To list the content of an .bz2 archive:
tar -jtvf /path/to/foo.tgz

# To create a .gz archive and exclude all jpg,gif,... from the tgz
tar czvf /path/to/foo.tgz --exclude=\*.{jpg,gif,png,wmv,flv,tar.gz,zip} /path/to/foo/

# To use parallel (multi-threaded) implementation of compression algorithms:
tar -z ... -> tar -Ipigz ...
tar -j ... -> tar -Ipbzip2 ...
tar -J ... -> tar -Ipixz ...

```
