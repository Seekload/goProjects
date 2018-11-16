package main

import "fmt"

func main(){

	/**
	一句话总结：注意*与谁结合
	如p *[5]int，*与数组结合说明是数组的指针；
	如p [5]*int，*与int结合，说明这个数组都是int类型的指针，是指针数组。
	 */
	arr := [...]int{5:9}
	// 数组指针
	var ptr *[6]int = &arr

	// arrPtr是一个指向arr数组首地址的指针
	fmt.Println(ptr[5])
	fmt.Println((*ptr)[5])

	// 指针数组
	x,y := 1,2
	var arrPtr = [...]*int{&x,&y}
	fmt.Println(arrPtr)


}