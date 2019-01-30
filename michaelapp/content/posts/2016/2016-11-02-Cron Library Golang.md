---
title: "cron library"
date: 2016-11-02T10:15:35+08:00
draft: false
tags: [
    "golang"
]
categories: [
    "golang",
]
---


# cron_test
cron is a crontab in golang(just a library)

Source code: [github.com/robfig/cron](github.com/robfig/cron)


```go
//every 5 second run a job
var spec5s = "0-59/5 * * * * *"

st := &stat{
	mac: 1000,
}

c.AddJob(spec5s, st)
```

```go
//every day(00:00:00) run a job
var specday = "0 0 0 1-31 * *"

day := &day{
	id: 9000,
}

c.AddJob(specday, day)
```

### Run as a daemon service
nohup cron_test >test.log 2>&1 &

### Reference Paper

1.[http://blog.studygolang.com/tag/cron/](http://blog.studygolang.com/tag/cron/)

2.[http://ju.outofmemory.cn/entry/65356](http://ju.outofmemory.cn/entry/65356)




