---
title: "11-变量作用域"
date: 2019-04-09T21:33:11+08:00
#slug: 20190409213311
weight: 1
---

## 3种变量作用域

- 局部变量  
  作用域只在函数体内
- 全局变量  
  全局变量可以在整个包 && 外部包（被导出后）使用
  > 外部包引用，变量要大写 
  ```
  import "outside"
  fmt.Printf("%s", outside.Name)
  ```
- 形式参数
  形式参数会作为函数的局部变量来使用。  
  **示例**
  ```
  package main

  import "fmt"

  /* 声明全局变量 */
  var a int = 20;

  func main() {
    /* main 函数中声明局部变量 */
    var a int = 10
    var b int = 20
    var c int = 0

    fmt.Printf("main()函数中 a = %d\n",  a);
    c = sum( a, b);
    fmt.Printf("main()函数中 c = %d\n",  c);
  }

  /* 函数定义-两数相加 */
  func sum(a, b int) int {
    fmt.Printf("sum() 函数中 a = %d\n",  a);
    fmt.Printf("sum() 函数中 b = %d\n",  b);

    return a + b;
  }  
  ```