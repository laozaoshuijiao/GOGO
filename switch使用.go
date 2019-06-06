package main

import "fmt"

func main(){
	var num int
	fmt.Println("输入楼层：")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Printf("按下的是%d楼\n",num)

	case 2:
		fmt.Printf("按下的是%d楼\n",num)

	case 3:
		fmt.Printf("按下的是%d楼\n",num)

	case 4:
		fmt.Printf("按下的是%d楼\n",num)

	default:
		fmt.Println("你妈的为什么")
	}
	score := 85
	switch  {
	case score > 90:
		fmt.Println("优秀")
	case score	< 90:
		fmt.Println(("ZZ"))



	}
}
