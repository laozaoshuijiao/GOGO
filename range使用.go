package main

import "fmt"

func main(){
	str := "abc"

	//通过for 打印每隔字符
	for i := 0 ;i < len(str);i++{
		fmt.Printf("str[%d] = %c\n",i,str[i])
	}

	//迭代打印每个元素，默认返回两个值:一个是元素的位置，一个是元素的本身
	for i, date := range str{
		fmt.Printf("str[%d] = %c\n",i,date)
	}
}
