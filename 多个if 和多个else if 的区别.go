package main

import "fmt"

func main(){
	a := 10
	if a == 10{
		fmt.Println("a==10")

	}else if a > 10{
		fmt.Println("a > 10")
	}else if a < 10{
		fmt.Println("a < 10")
	}
	b := 10
	if b == 10{
		fmt.Println("b==10")

	}
	if b >= 10{
		fmt.Println("b > 10")
	}
	if b <= 10{
		fmt.Println("b < 10")
	}


}
