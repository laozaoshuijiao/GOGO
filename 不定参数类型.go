package main

import "fmt"

func Myfunc(a , b int){//固定的参数
	fmt.Printf("a = %d , b = %d",a ,b)
}
//...int类型这样的类型，...type不定参数类型
func Myfunc01(args ...int){//传递的实参可以时0或多个
	fmt.Println("len(args = ",len(args))	 //获取用户传递参数的个数
	//返回2个值，第一个是下标，第二个是下标对应的数
	for i ,data := range args{
		fmt.Printf("args[%d] = %d\n",i,args[data])
	}
}
func main(){
	//Myfunc(666,777)
	Myfunc01()
	Myfunc01(1)
	Myfunc01(1,2,3)


}