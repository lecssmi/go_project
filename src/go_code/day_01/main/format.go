package main

import (
	"fmt"
)

//要求严格的格式，比如，下面的大括号 ，是不能分开写的
/*func main()
 {
	fmt.Println("hello  world !")
}*/

func main() {
	fmt.Println("hello world!")
}

//使用gofmt -w xxx.go，直接格式化代码,不然就只能手动处理，否则，没法提交且通过编译.
