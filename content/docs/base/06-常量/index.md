---
title: "06-常量"
date: 2019-04-09T21:32:34+08:00
#slug: 20190409213234
weight: 1
---

# 常量

格式：
```
const identifier [type] = value
```
`[type]` 可以省略，使用直行推断

示例:
```
const b string = "abc"
const b = "abc"
```

多个相同类型的声明, 简写
```
const c_name1, c_name2 = value1, value2
```