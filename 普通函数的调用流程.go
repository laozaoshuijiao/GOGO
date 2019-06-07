package main

import "fmt"

func funb()(b int){
	b = 22222
	fmt.Println("b = ",b)
	return
}

func funca()(a int){
	a = 111
	b := funb()
	fmt.Println("funcb = ",b)
	//调用另外一个函数
	fmt.Println("funca a =",a)
	return
}

func main(){
	fmt.Println("main func")
	a := funca()
	fmt.Printf("main a = %d",a)

}
