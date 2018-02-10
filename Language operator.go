package main

import "fmt"

func main() {
    
  /* 算术运算符 加减乘除余 自增自减*/  
  // var a int = 21
  // var b int = 10
  // var c int
  
  // c = a + b
  // fmt.Println("第一行 - c 的值为 ",c)
  // c = a - b
  // fmt.Println("第二行 - c 的值为 %d\n",c)
  // c = a * b
  // fmt.Println("第三行 - c 的值为 %d\n",c)
  // c = a / b
  // fmt.Println("第四行 - c 的值为 %d\n",c)
  // c = a % b
  // fmt.Println("第五行 - c 的值为 %d\n",c)
  // a++
  // fmt.Println("第六行 - c 的值为 %d\n",a)
  // a = 21
  // a--
  // fmt.Println("第七行 - c 的值为 %d\n",a)
  // fmt.Println("Hello World")
  
  /* 关系运算符 */
  
  // var a int = 21
  // var b int = 10
  
  // if(a == b){
  //   fmt.Println("第一行 - a 等于 b\n")
  // }else {
  //   fmt.Println("第一行 - a 不等于 b\n")
  // }
  // if(a < b){
  //   fmt.Println("第二行 - a 小于 b\n")
  // }else{
  //   fmt.Println("第二行 - a 不小于 b\n")
  // }
  // if(a > b){
  //   fmt.Println("第三行 - a 大于 b\n")
  // }else{
  //   fmt.Println("第三行 - a 不大于 b\n")
  // }
  
  // a = 5
  // b = 20
  // if(a <= b){
  //   fmt.Println("第四行 - a 小于等于 b\n")
  // }
  // if(b>=a){
  //   fmt.Println("第五行 - a 大于等于 b\n")
  // }
  
  /* 逻辑运算符 */
  // var a bool  = true
  // var b bool = false
  // if(a && b){
  //   fmt.Println("第一行 -条件为 true\n")
  // }
  // if(a || b){
  //   fmt.Println("第二行 -条件为 true\n")
  // }
  // /* 修改a 和 b 的值 */
  // a = false
  // b = true
  // if(a && b) {
  //   fmt.Println("第三行 -条件为 true\n")
  // }else{
  //   fmt.Println("第三行 -条件为 false\n")
  // }
  // if(!(a&&b)) {
  //   fmt.Println("第四行 -条件为 true\n")
  // }
  
  /* 逻辑运算符 */
  // var a uint = 60   /* 60 = 0011 1100*/
  // var b uint = 13   /* 13 = 0000 1101*/
  // var c uint = 0
  
  // c = a & b         /* 12 = 0000 1100*/
  // fmt.Println("第一行 - c的值为",c)
  
  // c = a | b        /* 61 = 0011 1101*/
  // fmt.Println("第二行 - c的值为",c)
  
  // c = a ^ b        /* 49 = 0011 0001*/
  // fmt.Println("第三行 - c的值为",c)
  
  // c = a << 2       /* 240 = 1111 0000*/
  // fmt.Println("第四行 - c的值为",c)
  
  // c = a >> 2       /* 15 = 0000 1111 */
  // fmt.Println("第五行 - c的值为",c)
  
  /* 赋值运算符 */
  // var a int = 21
  // var c int

  // c =  a
  // fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

  // c +=  a
  // fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

  // c -=  a
  // fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

  // c *=  a
  // fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

  // c /=  a
  // fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

  // c  = 200; 

  // c <<=  2
  // fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

  // c >>=  2
  // fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

  // c &=  2
  // fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

  // c ^=  2
  // fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

  // c |=  2
  // fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
   
   /* 其他运算符 */
   
  // var a int= 4
  // var b int32
  // var c float32
  // var ptr *int
   
   /* 运算符实例 */
  // fmt.Println("第 1 行 a变量类型为 =%T\n", a)
  // fmt.Println("第 2 行 b变量类型为 =%T\n", b)
  // fmt.Println("第 3 行 c变量类型为 =%T\n", c)
   
  // /* & 和 * 运算符实例 */
  // ptr = &a  /* 'ptr' 包含了 'a' 变量的地址 */
  // fmt.Println("a 的值为 %d\n", a)
  // fmt.Println("ptr 为 %d\n", *ptr)
   
   /* 运算优先级 */
  
   var a int = 20
   var b int = 10
   var c int = 15
   var d int = 5
   var e int;

   e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );  
  
  
  
  
}