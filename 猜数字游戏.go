package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int){
	rand.Seed(time.Now().UnixNano())
	var num int

	for {
		num =rand.Intn(10000)
		if num >= 1000{
			break

}
func main(){
	var randNum int

	//随机产生一个四位的随机数
	CreatNum(&randNum)
	fmt.Println("randNum = ",randNum)
}
