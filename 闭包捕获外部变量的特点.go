package main

import "fmt"

func main(){
	a := 10
	str := "make"

	func (){
		//闭包以引用方式捕获外部变量
		a = 66
		str = "go"
		fmt.Printf("a = %d,str = %s\n",a,str)
	}()//()代表直接调用

	fmt.Printf("a = %d,str = %s\n",a,str)

}
