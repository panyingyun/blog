---
title: "Golang Context的使用"
date: 2016-10-27T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---

参考：http://www.nljb.net/default/Golang%E4%B9%8BContext%E7%9A%84%E4%BD%BF%E7%94%A8/

### context包介绍：
```
在golang中的创建一个新的线程并不会返回像c语言类似的pid
所有我们不能从外部杀死某个线程，所有我就得让它自己结束
之前我们用channel＋select的方式，来解决这个问题
但是有些场景实现起来比较麻烦，例如由一个请求衍生出多个线程
并且之间需要满足一定的约束关系，以实现一些诸如：
有效期，中止线程树，传递请求全局变量之类的功能。
于是google 就为我们提供一个解决方案，开源了context包。
使用context包来实现上下文功能 .....
约定：需要在你的方法的传入参数的第一个参数是context.Context的变量。
```

### 代码示例：
```golang
package main
import (
    "fmt"
    "time"
    "golang.org/x/net/context"
)
// 模拟一个最小执行时间的阻塞函数
func inc(a int) int {
    res := a + 1                // 虽然我只做了一次简单的 +1 的运算,
    time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,
    // 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
    return res
}
// 向外部提供的阻塞接口
// 计算 a + b, 注意 a, b 均不能为负
// 如果计算被中断, 则返回 -1
func Add(ctx context.Context, a, b int) int {
    res := 0
    for i := 0; i < a; i++ {
        res = inc(res)
        select {
        case <-ctx.Done():
            return -1
        default:
        // 没有结束 ... 执行 ...
        }
    }
    for i := 0; i < b; i++ {
        res = inc(res)
        select {
        case <-ctx.Done():
            return -1
        default:
        // 没有结束 ... 执行 ...
        }
    }
    return res
}
func main() {
    {
        // 使用开放的 API 计算 a+b
        a := 1
        b := 2
        timeout := 2 * time.Second
        ctx, _ := context.WithTimeout(context.Background(), timeout)
        res := Add(ctx, 1, 2)
        fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
    }
    {
        // 手动取消
        a := 1
        b := 2
        ctx, cancel := context.WithCancel(context.Background())
        go func() {
            time.Sleep(2 * time.Second)
            cancel() // 在调用处主动取消
        }()
        res := Add(ctx, 1, 2)
        fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
    }
}
```

