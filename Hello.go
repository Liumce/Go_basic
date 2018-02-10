package main

/*
 var vname1, vname2, vname3 type
 vname1, vname2, vname3 = v1, v2, v3
 
 var vname1, vname2, vname3 = v1, v2, v3 
 vname1, vname2, vname3 := v1, v2, v3
 
 // 这种因式分解关键字的写法一般用于声明全局变量
 var{
   vname1 v_type1
   vname2 v_type2
 }
 */
 
 // 实例   var()  这里是双括号（），  不是｛｝
 var x, y int
 var(
   a int
   b bool
 )
 var c, d int = 1, 2
 var e, f = 123, "hello" 
 
 // 这种不带声明格式的只能在函数体中出现
 // g, h := 123,"hello"
 

func main() {
 
 g, h := 123,"hello"
 println(x, y, a, b, c, d, e, f, g, h)
 
  
}