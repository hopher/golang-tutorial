---
title: "12-数组"
date: 2019-04-09T21:33:20+08:00
#slug: 20190409213320
weight: 1
---

# 数组

数组声明

```
var balance [10] float32
```

初始化数组

> [...] 为不定长

```
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

访问数组元素

```
var salary float32 = balance[9]
```

示例：

```
package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
```