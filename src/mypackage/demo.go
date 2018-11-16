package mypackage

import "fmt"

var I int

/**
这是自定义的一个包
如何使用：

import (
	"fmt"
	"mypackage"
)

func main(){
	fmt.Println("Hello go,I=",mypackage.I)
}
 */




func init(){
	I = 0
	fmt.Println("call package init1")
}

func init()  {
	I = 1
	fmt.Println("call package init2")
}



