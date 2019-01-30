---
title: "Golang go build"
date: 2016-10-27T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

## 命令行使用：
### 1. go build 编译（apiserver.exe amd64/windows7  or apiserver amd64/linux）
``` shell
   go build -x  编译时列出所有编译时的指令（先编译package  最后 link）
```

### 2. go build -race  竞态检查

``` shell
   go build -gcflags "-xxx -xxx"
   比如  go build -gcflags="-N -l"  禁用编译器优化和内联优化
   比如
   把Go程序变小的办法是： go build -ldflags "-s -w" (go install类似)
	-s去掉符号表（然后panic时候的stack trace就没有任何文件名/行号信息了，
	这个等价于普通C/C++程序被strip的效果），
	-w去掉DWARF调试信息，得到的程序就不能用gdb调试了。
```

```shell
其他一些gcflags参数如下： 
  usage: compile [options] file.go...
  -%    debug non-static initializers
  -+    compiling runtime
  -A    for bootstrapping, allow 'any' type
  -B    disable bounds checking
  -D path
        set relative path for local imports
  -E    debug symbol export
  -I directory
        add directory to import search path
  -K    debug missing line numbers
  -M    debug move generation
  -N    disable optimizations
  -P    debug peephole optimizer
  -R    debug register optimizer
  -S    print assembly listing
  -V    print compiler version
  -W    debug parse tree after type checking
  -asmhdr file
        write assembly header to file
  -buildid id
        record id as the build id in the export metadata
  -complete
        compiling complete package (no C or assembly)
  -cpuprofile file
        write cpu profile to file
  -d list
        print debug information about items in list
  -dynlink
        support references to Go symbols defined in other shared libr
  -e    no limit on number of errors reported
  -f    debug stack frames
  -g    debug code generation
  -h    halt on error
  -i    debug line number stack
  -importmap definition
        add definition of the form source=actual to import map
  -installsuffix suffix
        set pkg directory suffix
  -j    debug runtime-initialized variables
  -l    disable inlining
  -largemodel
        generate code that assumes a large memory model
  -linkobj file
        write linker-specific object to file
  -live
        debug liveness analysis
  -m    print optimization decisions
  -memprofile file
        write memory profile to file
  -memprofilerate rate
        set runtime.MemProfileRate to rate
  -msan
        build code compatible with C/C++ memory sanitizer
  -newexport
        use new export format (default true)
  -nolocalimports
        reject local (relative) imports
  -o file
        write output to file
  -p path
        set expected package import path
  -pack
        write package file instead of object file
  -r    debug generated wrappers
  -race
        enable race detector
  -s    warn about composite literals that can be simplified
  -shared
        generate code that can be linked into a shared library
  -ssa
        use SSA backend to generate code (default true)
  -trimpath prefix
        remove prefix from recorded source file paths
  -u    reject unsafe code
  -v    increase debug verbosity
  -w    debug type checking
  -wb
        enable write barrier (default true)
  -x    debug lexer
```

### 参考文献：
	
1. Go的口头禅“不要使用共享数据来通信；使用通信来共享数据”

2. 文档 godoc -http=:6060

    http://localhost:6060/doc/    

3.  性能调试

    http://blog.cyeam.com/golang/2016/08/18/apatternforoptimizinggo
4.  Go语言圣经中文版

	https://docs.ruanjiadeng.com/gopl-zh/ch0/ch0-05.html

5.  文档搜索

	https://godoc.org
	
	https://gowalker.org/
	
	https://golanglibs.com/top?q=mqtt

6.  go 后台任务管理器

	https://www.goworker.org/

7.  fing 探测局域网里面的设备
	
	http://blog.cyeam.com/network/2015/03/16/fing

8.  竞态检查参考：
	
	http://localhost:6060/doc/articles/race_detector.html
	
	https://docs.ruanjiadeng.com/gopl-zh/ch9/ch9-01.html

9. 	竞态的实例
	
   http://localhost:6060/doc/articles/race_detector.html

```golang
	package main
	
	import (
		"fmt"
	)
	
	func main() {
		c := make(chan bool)
		m := make(map[string]string)
		go func() {
			m["1"] = "a" // First conflicting access.
			c <- true
		}()
		m["2"] = "b" // Second conflicting access.
		<-c
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
```

