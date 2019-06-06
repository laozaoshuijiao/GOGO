package main

import "fmt"

func main(){
	type bigint int64
	var a  bigint  //等价于int64 a

	fmt.Printf("%T",a)


}
