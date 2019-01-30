---
title: "Vagrant和VirtualBox在CentOS7.2上创建多个虚拟机"
date: 2016-10-24T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

1. Install golang
```shell
tar -C /usr/local -xzf go1.7.3.linux-amd64.tar.gz
/etc/profile
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:/usr/local/go/bin
source /etc/profile
```

2. Install GCC/OpenSSL 
```shell
yum install -y openssl-devel
yum groupinstall -y "Development Tools"
```

3. Install goget
```shell
安装goget工具
go get github.com/coderhaoxin/goget
cp goget /usr/local/bin/
```

4. Install  VirtualBox
```shell
VirtualBox安装
cd  /etc/yum.repos.d/
goget http://download.virtualbox.org/virtualbox/rpm/el/virtualbox.repo
yum install -y VirtualBox-5.1
```

5. VirtualBox安装 https://www.virtualbox.org/wiki/Linux_Downloads

6. Vagrant安装 https://www.vagrantup.com/downloads.html

7. 下载Centos Linux Box(http://www.vagrantbox.es/)
```shell
goget https://releases.hashicorp.com/vagrant/1.8.6/vagrant_1.8.6_x86_64.rpm
cd /root/.vagrant.d/boxes
goget https://github.com/CommanderK5/packer-centos-template/releases/download/0.7.2/vagrant-centos-7.2.box
goget https://github.com/kraksoft/vagrant-box-ubuntu/releases/download/14.04/ubuntu-14.04-amd64.box
cd /root/.vagrant.d/boxes
vagrant box add cent1 /root/.vagrant.d/boxes/vagrant-centos-7.2.box
vagrant box add cent2 /root/.vagrant.d/boxes/vagrant-centos-7.2.box
vagrant box add cent3 /root/.vagrant.d/boxes/vagrant-centos-7.2.box
vagrant init cent1
vagrant up
注：
vagrant reload 修改配置后重启
vagrant up
```

8. Vagrantfile配置参考

	```shell
	Vagrant.configure("2") do |config|
	  config.vm.define :cent1 do |cent1|
	    cent1.vm.provider "virtualbox" do |v|
	          v.customize ["modifyvm", :id, "--name", "cent1", "--memory", "512"]
	    end
	    cent1.vm.box = "cent1"
	    cent1.vm.hostname = "cent1"
	    cent1.vm.network :private_network, ip: "192.168.2.110"
	  end
	
	  config.vm.define :cent2 do |cent2|
	    cent2.vm.provider "virtualbox" do |v|
	          v.customize ["modifyvm", :id, "--name", "cent2", "--memory", "512"]
	    end
	    cent2.vm.box = "cent2"
	    cent2.vm.hostname = "cent2"
	    cent2.vm.network :private_network, ip: "192.168.2.120"
	  end
	  
	  config.vm.define :cent3 do |cent3|
	    cent3.vm.provider "virtualbox" do |v|
	          v.customize ["modifyvm", :id, "--name", "cent3", "--memory", "512"]
	    end
	    cent3.vm.box = "cent2"
	    cent3.vm.hostname = "cent3"
	    cent3.vm.network :private_network, ip: "192.168.2.130"
	  end
	end
	```

9. 使用（注意执行下面的命令 一定要在Vagrantfile 同目录下进行）

	```shell
	vagrant ssh cent1
	vagrant ssh cent2
	vagrant ssh cent3
	```
	分别登录三个服务器

	vagrant global-status
	
	ip 命令 替代老的ifconfig\
	
	[https://linuxstory.org/replacing-ifconfig-with-ip/](https://linuxstory.org/replacing-ifconfig-with-ip/)
	
	ip -4 a
	注：虚拟器的 /vagrant目录 映射的就是 主机目录/root/install/ 

10. 更深入的学习 Vagrant

	[https://a358003542.github.io/cross-platform/vagrant-tutorial.html](https://a358003542.github.io/cross-platform/vagrant-tutorial.html)