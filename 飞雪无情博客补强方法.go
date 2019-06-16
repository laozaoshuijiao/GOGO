/*
方法的声明和函数类似，他们的区别是：方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者
这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法
*/
package main

import "fmt"
type person struct {
	name string
}
/*func 和方法名之间增加的参数（p person），这个就是接收者。类型person有了一个String方法*/
func (p person) String() string{
	/*使用类型的变量进行调用即可，类型变量和方法之前是一个 . 操作符，表示要调用这个类型变量的某个意思
	Go语言有两种类型的接收者：值接收和指针接收*/
	return "the person name is "+p.name
}
func main(){
	p := person{name : "张三"}
	fmt.Println(p.String())
}
