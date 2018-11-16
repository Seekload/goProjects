package main


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

func main() {

	/**
	函数返回单个值
	 */
	var a int = 1
	var b int = 2
	var maxR int
	maxR = max(a,b)
	println(maxR)

	/**
	函数返回多个值
	 */
	strA,strB := swap("apple","go")
	println(strA,strB)




}


