package main

import "fmt"

func main(){

	// 数组声明
	// 1、
	var balance [10]float32
	for i:=0;i<10 ;i++  {
		balance[i] = 100
	}

	// 2、
	var arr  = [...]int{5:9}
	fmt.Println(arr)

	// 3、
	array := [5]int{1:1,4:4}
	fmt.Println(array)

	for i,v := range array{
		fmt.Println(i,v)
	}


}