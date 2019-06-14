package main

import "fmt"

func main(){
	a := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println("a = ",a)


	//新切片
	s1 := a [2 : 5] //从2到5

	fmt.Println("s1 =",s1)
}