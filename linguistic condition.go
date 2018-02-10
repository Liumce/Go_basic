package main

import "fmt"

func main() {
  /* 局部变量定义 */
  // var a int = 100;
  
  // /* 判断布尔表达式 */
  // if a < 20 {
  //   /* 如果条件为true 则执行以下语句 */
  //   fmt.Println("a 小于 20\n ");
  // }else {
  //   fmt.Println("a 不小于 20\n ")
  // }
  // fmt.Println("a 的值为：\n ",a);
  
  
  
  // 寻找到100以内的所有素数
  // 定义了变量如果不用也会报错
  var count int
  var flag bool
  count = 1
  // while(count<100){   // go没有while
  for count <100 {
    count++
    flag  = true;
    // 注意tmp变量 :=
    for tmp:=2;tmp<count;tmp++ {
      if count%tmp==0{
        flag = false
      }
    }
    if flag ==true{
      fmt.Println(count,"素数")
    }else{
      continue
    }
  }
}