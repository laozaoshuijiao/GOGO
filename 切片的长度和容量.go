package main

import "fmt"

func main(){
	a := []int{1,2,3,0,0}
	s := a[0:3:5]
	fmt.Println("s =",cap(s))    //长度
	fmt.Println("len(s) = ",len(s))//容量

	p := a[1:4:5]
	fmt.Println("s =",cap(p))    //长度
	fmt.Println("len(s) = ",len(p))//容量
}

