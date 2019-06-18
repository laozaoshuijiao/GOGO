/*
GO语言提供了另外一种数据类型即为接口，他把所有具有共性的方法定义在一起，任何其他类型只要实现了这些方法
就是实现了这个接口
*/

package main

import(
	"fmt"
)
type Phone interface{
	call()
}
type NokiaPhone struct {

}
func (nokiaphone NokiaPhone)call() {
	fmt.Println("I am Nokia I can call you！")
}
type IPhone struct{

}
func (iphone IPhone)call(){
	fmt.Println("I am iphone , l can call you!")
}
func main(){
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()
}