package main

import "fmt"

// Go 支持匿名函数，匿名函数可以形成闭包。
// 此函数 intSeq 返回另一个函数，我们在 intSeq 的主体中匿名定义该函数。返回的函数对变量 i 进行封闭以形成闭包。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	// 输出为1 2 3 4 5
	for i := 0; i < 5; i++ {
		fmt.Println(nextInt())
	}

	// 闭包中的变量仅受闭包影响。
	// 该函数输出为1
	nextInt1 := intSeq()
	fmt.Println(nextInt1())
}
