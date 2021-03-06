---
title: "15-切片"
date: 2019-04-09T21:33:49+08:00
#slug: 20190409213349
weight: 1
---

# 切片

切片扩容时，会重新建立新的底层数组，内存重新分配内存储存空间, 然后把旧的数据复制到新的数组里, 时间复杂度O(n)
切片长度 >= 1024时, 以1.25倍进行扩容, 小于1024时则2倍扩容

golang 官方文档在关于 slice 的 append 一节有说明：
> If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array.

也就是说，如果原来的底层数组足够大，能放的下 append 的数据，就不会新建底层数组。而如果不够的话，则会分配一个新的数组。也因此是 O(n) 的时间复杂度


## 切片 vs 数组

声明时，数组指定长度大小, 切片没有

```
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 数组
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // 数组

var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}  // 切片
```


## 相关函数

- len() 获取切片长度 
- cap() 获取切片容量

