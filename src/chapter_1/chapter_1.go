package main

import "fmt"

//import "fmt"

func p() *int{
	v := 1
	return &v
}

func a(p *int){
	*p++
}


func main(){

	//var i int = 10
	//var ptr *int = &i
	//fmt.Println(ptr,*ptr)  // 0xc000018060 10  指针变量ptr存储的是i地址，*ptr对应指针指向的变量的值
	//*ptr = 12  				// i=12  更新指针指向的变量的值
	//fmt.Println(*ptr,i)		// 12 12

	//var i int = 10
	//var i int = 10    // 1、指针声明
	//var ip *int
	//ip = &i           // 2、赋值
	//fmt.Println(*ip)     // 访问
	//pstr = &str
	//fmt.Println(*ip,*pstr)

	//var i int = 10
	//var x int = 10
	//var y = 10

	p := new(int)
	q := new(int)
	fmt.Println(p,q)





}