package main

import "fmt"

func main() {
  // 判断最大值
  /* 定义局部变量 */
//   var a int  = 100
//   var b int = 200
//   var ret int
  
//   /* 调用函数并返回最大值 */
//   ret = max(a,b)
  
//   fmt.Println("最大值是：\n",ret)
 //}

// /* 函数返回两个数的最大值 */
// func max(num1,num2 int) int {
//   /* 定义局部变量 */
//   var result int
  
//   if(num1>num2) {
//     result = num1
//   }else {
//     result = num2
//   }
//   return result

// 返回多个值
a,b := swap("Mahesh","Kumar")
fmt.Println(a,b)
 }
 
 func swap(x,y string)(string,string) {
   return y,x
 }