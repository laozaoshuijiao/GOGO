/*
函数和方法，虽然概念不同，但是定义非常相似。函数的定义声明没有接收者，所以我们直接在go文件里，
go包之下定义声明即可
*/

package main

import "fmt"

func add(a , b int)int{

	return a+b
}

func main(){
	sum := add(1,2)
	fmt.Println("sum = ",sum)
}
/*
例子中，我们定义了add就是一个函授，它的函数签名是func add(a , b int)int,没有接收者，直接定义在go的
一个包之下，可以直接调用，比如例子中的main函数调用了add函数。
例子中的这个函数名称是小写开头的add，所以它的作用域只属于所声明的包内使用，不能被其他包使用
*/