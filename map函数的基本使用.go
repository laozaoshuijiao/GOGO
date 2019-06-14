package main

import "fmt"

func main(){
	//定义一个变量,类型为map[int]string
	var m1 map[int]string

	fmt.Println("m1 =",m1)
	//对于map只有len没有cap
	fmt.Println("len =",len(m1))

	//可以通过make创建s
	m2 := make(map[int]string)
	fmt.Println("m2 = ",m2)
	fmt.Println("len m2 = ",len(m2))

	m3 := make(map[int]string,10)
	m3[0] = "mike"
	m3[1] = "c++"
	m3[2] = "python"

	fmt.Println("m3 = ",m3)
	fmt.Println("len m3 = ",len(m3))
	fmt.Println("m3[0]",m3[0])

	m4 := map[int]string{1:"mike",2:"go",3:"c++"}
	fmt.Println("m4 :",m4)

	fmt.Println(m4[1])
}
