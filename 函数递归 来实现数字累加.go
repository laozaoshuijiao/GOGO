//递归，就是在运行的过程中调用自己
/*
语法格式如下：
func 递归(){
	递归() //函数调用自身
}

func main(){
	递归
}
*/
/*
GO语言支持递归，但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。
递归函数对于解决数学上的问题是非常有用的，像是计算乘积，生成裴波那契数列等
*/
//以下实例通过Go语言递归函数实例阶乘:
package main

import "fmt"

func Factorial(n int)(result int){
	if (n > 0){
		result = n * Factorial(n - 1)
		return result
	}
	return 1
}
func main(){
	i := 15
	fmt.Printf("%d 的乘积是 %d\n",i , Factorial(int(i)))
}