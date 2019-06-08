package main

import (
	"fmt"

)

type FuncType func(int int)int

//回调函数，函数参数是函数类型，这个函数就是回调函数了
//计算器，可以进行+-*/
func Calc(a , b int fTest FuncType)(result int){
	fmt.Println("Calc")
	result = fTest(a )
}

func main(){

}
