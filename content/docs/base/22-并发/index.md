---
title: "21-并发"
date: 2019-04-09T21:34:44+08:00
#slug: 20190409213444
weight: 1
---

## 知识列表 

- goroutine
- channel

## goroutine 

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

语法格式

```
go 函数名( 参数列表 )
go f(x, y, z)
```

## channel

> Don’t communicate by sharing memory; share memory by communicating.   
>（不要通过共享内存来通信，而应该通过通信来共享内存。）

![img](1.png)