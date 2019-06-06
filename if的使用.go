package main

import "fmt"

func main(){
	s := "王思聪"
	//if就是关系条件运算符
	if s == "王思聪"{ //左括号和 if在同一行
		fmt.Println("左手一个妹子右手一个大妈")
	}
	//if支持一个初始化语句
	if a := 10;a == 10{//条件为真就运行大括号的东西
		fmt.Println("a == 10")
	}
}
