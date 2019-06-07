/*
当一个函数在其函数体内调用自身，则称之为递归。最经典的例子便是
计算裴波那契数列，即前两个数为1，从第三个数开始每个数均为
前两个数之和
*/
package main

import "fmt"

func fibonacci( n int)(res int){
	if n <=1{
		res = 1
	}else{
		res = fibonacci(n -1) +fibonacci(n -2)
	}	
	return 
}

func main(){
	result := 0
	for i := 0 ; 1 < 10; i ++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d\n", i ,result)
		fmt.Printf("i = %d",i)

	}
}

