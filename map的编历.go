package main

import "fmt"

func main(){
	m := map[int]string{1 : "mike" , 2 : "yoyo" , 3 : "go"}
	fmt.Println("m =" , m)

	for i := 0 ; i < 100000000000; i ++ {
		fmt.Println("m = ",m)
	}
}
