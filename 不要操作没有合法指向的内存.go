package main

import"fmt"

func main(){
	var p *int
	fmt.Println("p = ",p)

	//*p = 6666,没有合法指向

	var a int
	p = &a //p 指向 a.
	*p = 666
	fmt.Println("a =",a)
}
