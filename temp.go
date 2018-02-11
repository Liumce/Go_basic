package main

import "fmt"

func main() {
 /* 值传递 */
 var a int = 100
 var b int = 200
 
 fmt.Println("交换前 a 的值为：\n",a)
 fmt.Println("交换前 b 的值为：\n",b)
 
 /* 通过调用函数来交换值 */
 swap(a,b)
 
 fmt.Println("交换后 a 的值为：\n",a)
 fmt.Println("交换后 b 的值为：\n",b)
}

/* 定义相互交换值的函数 */
func swap(x,y int) int{
  var temp int
  
  temp = x /* 保存x的值 */
  x = y /* 将y值赋给x */
  y = temp /* 将temp 值赋给y */
  
  return temp;
}
