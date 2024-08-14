package main

import "fmt"
// go支持函数递归
// 阶乘
func fact(n int) int{
	if n ==0 {
		return 1
	}
	return n * fact(n-1)
}

func main(){
	fmt.Println(fact(4))
	// 斐波那契数列

	var fib func(n int) int

	fib = func(n int) int {
		if n <2 {
			return 2
		} else {
			return fib(n-1) + fib(n-2)
		}

	}
	fmt.Println(fib(3))
}