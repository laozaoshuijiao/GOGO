package main

import "fmt"

func main(){
	//break
	//continue
	//goto 可以用在任何地方，但是不能跨函数使用
	fmt.Println("11111111111111111")
	goto End    //goto是关键字 ，End是用户起的名字,他叫标签
	fmt.Println("zzzzzzzzzzzzzzzzz")

End:
	fmt.Println("33333333333")

}
