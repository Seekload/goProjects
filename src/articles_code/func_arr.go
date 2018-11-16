package main

import "fmt"

/**
	函数间传递数组
 */
 /**
 	传递数组的副本
  */
func modify(a [5]int)  {
	a[1] = 1
	fmt.Println(a)
}

/**
	传递数组的指针
 */
func modifyPtr(a *[5]int){
	a[1] = 2
	fmt.Println(*a)
}

func main(){

	arr := [5]int{4:9}
	modifyPtr(&arr)
	fmt.Println(arr)

}