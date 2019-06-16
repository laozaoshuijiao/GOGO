/*
map是给予散列表来实现的，每次迭代map的时候，打印的key和value是无序的，每次迭代的都不一样，即使按照一定
顺序也不行。
Map的散列表包含一组桶，每次储存和查找键值对的时候，都要先选择一个桶。如何选择桶，就是把指定的键传给
散列函数，就可以索引到相应的桶，进而找到对应的键值。
这种方法好处在于，储存的数据越多，索引分布越均匀，所以访问键值对的速度也就越快
*/
package main

import "fmt"

func main(){
	//Map的创建有make函数，Map字面量，make函数我们用它创建过切片，除此之外，它还可以用来创建Map
	//创建一个键类型为string的，值为int的map
		//dict := make(map[string]int)
		//map创建好后，但里面是空的，我们给储存一个键值对
		//dict["张三"] = 43
		//储存了一个key为张三的键，value为43的的键值数据
		//此外还有一种使用map字面量的方式创建和初始化map，对于上面的例子，我们可以完美的等同实现
		//使用一个大括号进行初始化，键值对通过用：分开，如果要同时初始化多个键值对，用逗号分开
		dict2 := map[string]int{"张三":43,"李四":50}
		dict2["xx"] = 18
		/*
		当然可以不指定任何键值，也就是一个空的map
		dict := map[string]int{}
		不管怎么样，使用map的字面量创建一定腰带有大括号。如果我们要船舰一个nil的map
		*/

		dict := map[string]int{"张三":43}

		fmt.Println(dict)

		age := dict["张三"]
		fmt.Println(age)

		age1 ,exists := dict["李四"]

}