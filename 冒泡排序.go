package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [10] int
	n := len(a)

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100) //获取100以内的随机数
		fmt.Printf("%d", a[i])

	}
	fmt.Printf("\n")

	//冒泡排序，挨着哥的2个元素比较，升序
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Printf("\n排序后:\n")
	for i := 0; i < n; i++ {

		fmt.Printf("%d ", a[i])
	}
}