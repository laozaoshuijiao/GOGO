package main

import "fmt"

func main(){
	var a int
	var p *int
	a = 10
	fmt.Println("a = ",a)
	fmt.Printf("&a = %v\n",&a)

	p = &a
	fmt.Printf("p = %v , &a = %v\n",p,&a)

	*p = 666 //*p操作的不是p的内存，是p所指向的内存 p指向 a， *p = 6666 == a =666
	fmt.Printf("*p = %v a = %v\n",*p,a)
}