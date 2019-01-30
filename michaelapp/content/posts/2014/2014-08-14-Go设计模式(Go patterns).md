---
title: "Go 设计模式(Go patterns)"
date: 2014-08-14T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

#### 1.Generator(发生器)
在Google IO 2012大会中提到的Go pattern，记录如下，以便加深理解。
Go patterns 可以理解为Go的设计模式，这个往往是在实践中遇到的一些典型场景而总结出来的通用的方法论。
Generator可以理解为发生器

```Golang
//golang partens
//Generator: function that returns a channel
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	google := boring("google")
	apple := boring("apple")
	for i := 0; i < 5; i++ {
		fmt.Println(<-google)
		fmt.Println(<-apple)
	}
}
```
它返回一个只能接受的channel，内部起一个goruntine并发产生string（当然也可以是其他事物，这里只是用string举例）
事实上可以将boring看出一个generator服务，而且很容易生产多个实例，只要是不同的输入参数就可以了。
每个服务又是在不同的goruntine中独立运行的。
这里的管道（channel）很像服务器的handle，把主goruntine和服务的goruntine联系起来了。比起Android里面的handle 这里的channel更加直观和容易理解。

完整代码：

1. https://github.com/panyingyun/gostudy/blob/master/generator.go

#### 2.Multiplexing(多路复合)
```Golang

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("google"), boring("apple"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're boing; I'm leaving.")
}
```
完整代码：

1. https://github.com/panyingyun/gostudy/blob/master/multiplexing.go

#### 3.多路选择(select 实现的fanIn)
fanIn的select语法的实现方式
```Golang
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
```
完整代码：

1. https://github.com/panyingyun/gostudy/blob/master/select.go

#### 4.timeout(超时)

超时channel有数据则返回，超时往往都是这么弄的，特别是在网络请求的时候，超时是必要的选择。
```Golang
func timeout(ch <-chan string) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much!")
			return
		}
	}
}
```
完整代码：

1. https://github.com/panyingyun/gostudy/blob/master/timeout.go

#### 5.退出(quit)

```Golang
func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	c := boring("goo", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("goo says: %q\n", <-quit)

}
```
主goroutine通过往quit channel中写数据和boring进行通信，告诉它退出。

完整代码：

1. https://github.com/panyingyun/gostudy/blob/master/quit.go

#### 6.Google Search的模拟实现
这里提供了Google Search模拟实现的4个版本，慢慢体会其中的道理(主要是利用前面的这些知识)

```Golang
var (
	Web    = fakeSearch("web")
	Image  = fakeSearch("image")
	Video  = fakeSearch("video")
	Web1   = fakeSearch("web1")
	Image1 = fakeSearch("image1")
	Video1 = fakeSearch("video1")
	Web2   = fakeSearch("web2")
	Image2 = fakeSearch("image2")
	Video2 = fakeSearch("video2")
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google10(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func Google20(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func Google21(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func Google30(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google30("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// test first function
// func main() {

// 	rand.Seed(time.Now().UnixNano())
// 	start := time.Now()
// 	result := First("golang",
// 		fakeSearch("replica 1"),
// 		fakeSearch("replica 2"))
// 	elapsed := time.Since(start)
// 	fmt.Println(result)
// 	fmt.Println(elapsed)
// }
```

#### 部分代码注解：
本人对“函数”和“方法”叫法不做什么区别，所以看到“函数”等名字可以和“方法”等同。

(1) type Result string 和type Search func(query string) Result
分别自定义一个Result类型（实际上等同一个字符串）
自定义一个函数类型 Search 
定义Search的好处显而易见，代码会简化不少，而且意思表达也清晰

(2) fakeSearch是一个闭包，用于模拟搜索引擎的搜索状态
其中使用time.Sleep来模拟等待服务器返回的时间

(3) Web、Image、Video、Web1等9个var变量可以认为是9台搜索服务器，
 可以认为是9台分布式服务器

(4) Google1.0搜索版本 func Google10(query string) (results []Result)
比较浅显，是将3台服务器Web、Image、Video 返回的结果进行整合返回
对这个函数中results变量的写法和return的写法是非常的Golang style,
对返回值定义了变量名，运行时会自动创建0值变量，返回的时候也就可以省略掉
这个返回参数。
从Google1.0看，它实际上是三个搜索查询是串行的，在主goroutine中运行

(5)  Google2.0 搜索版本func Google20(query string) (results []Result) 
 将三个搜索放到不同的goroutine中，并且用channel进行连接整合，从实际运行
 时间看，并发后总的搜索时间要少很多。

(6) Google2.1 搜索版本 func Google21(query string) (results []Result) 加上了
超时处理，这样可以防止搜索时间太久，在规定的80ms内必须响应用户查询
请求，返回搜索结果。
从结果看，Google2.1比Google1.0版本用好，但是搜索结果没有能最优，所以我们有了Google3.0版本

(7) Google3.0办法 func Google30(query string) (results []Result)
对于同类服务提供了多路复用，web服务器增加为两台，那么我们给两台服务器发相同
的搜索请求，谁返回结果快，我们就用谁的结果，这样比Google2.1的搜索结果要更加快，
结果也更加优秀，因为多路都超时的概率毕竟下降很多。
这里的多路复用在函数func First(query string, replicas ...Search) Result 中完成的，其中replicas
即是提供多个相同服务的服务器，输入包括搜索词和可变参数replicas，它从多服务器中响应那个最快返回结果
丢弃其他的服务器结果。
   
#### 完整代码：

1. https://github.com/panyingyun/gostudy/blob/master/googlesearch.go

#### 参考文献：

1. https://talks.golang.org/2012/concurrency.slide#25