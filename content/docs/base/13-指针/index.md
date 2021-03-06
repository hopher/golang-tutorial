---
title: "13-指针"
date: 2019-04-09T21:33:29+08:00
#slug: 20190409213329
weight: 1
---

## 指针

指针，返回的是内存地址空间 (十六进制)

格式：`0xc000070010`, 开头的“0”令解析器更易辨认数，而“x”则代表十六进制

> **二进制**
> 数据在计算机中的表示，最终是以二进制的形式存在
> - 优点：更直观地解决
> - 缺点：太长，不容易思考，可以用更短的16进制表示


指针一般用 `ptr` 表示

```
package main

import "fmt"

func main() {
   var a int = 10
   fmt.Printf("变量的地址: %v\n", &a  )
   fmt.Printf("变量的地址: %#x\n", &a  )
}
```


值方法 vs 指针方法区别

golang隐式传递指针，但是不隐式定义指针

```
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}


func printBook( book Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}
```