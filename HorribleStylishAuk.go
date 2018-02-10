package main

import "fmt"
//import "unsafe"

// Go 语言常量
func main() {
  // 显式类型定义
  // const b string = "abc"
  // 隐式类型定义
  // const b = "abc"
  // fmt.Println("Hello World")
  
  // const LENGTH int = 10;
  // const WIDTH int = 5;
  // var area int
  // const a, b, c = 1, false, "str" //多重赋值
  
  // area  =LENGTH * WIDTH
  // fmt.Println("面积为 : ", area)
  // println()
  // println(a, b, c)
  
  // const (
  //   a = "abc"
  //   b = len(a)
  //   c = unsafe.Sizeof(a)
  //   )
    
  
  // fmt.Println(a, b, c)
  
  
  // iota 用法一
  // const (
  //   a = iota  //0
  //   b         //1
  //   c         //2
  //   d = "ha"  // 独立值，iota += 1
  //   e         // "ha" iota += 1
  //   f = 100   // iota += 1
  //   g         //100 iota+=1
  //   h = iota  //7,恢复计数
  //   i         //8
  //   )
  //   fmt.Println(a,b,c,d,e,f,g,h,i)
    
  // iota 用法二 
  /*
     iota表示从0开始自动加1，所以i=1<<0,j=3<<1（<<表示左移的意思），即：i=1,j=6，这没问题，关键在k和l，从输出结果看，k=3<<2，l=3<<3。
  */
    const(
        i=1<<iota
        j=3<<iota
        k
        l
  )
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
  
  
}