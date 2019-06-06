package main

import "fmt"

func main(){
	var ch byte
	ch = 97
	fmt.Printf("%c,%d\n",ch,ch)
	ch = 'a'
	fmt.Printf("大写：%d,小写：%d\n",'A','a')
}