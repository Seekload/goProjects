package main

import "fmt"

func swap(stra,strb string) (string,string) {
	return strb,stra
}

func max(a,b int) int {
	if a>=b {
		return a
	}else {
		return b
	}
}

type name string

func main() {

	var strA string
	var strB name = "hello world"
	var strC string = "go"

	strA = strC

	fmt.Println(strA,strB,strC)






}


