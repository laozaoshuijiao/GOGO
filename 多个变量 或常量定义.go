package main

import "fmt"

func main(){
	//不同类型变量的声明
	var (
		a int
		b float64
	)
	const(
		i int = 10
		j float64 = 3.14
	)


	a , b = 10 , 3.14

	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	fmt.Println("i = ",i)
	fmt.Println("j = ",j)
}
