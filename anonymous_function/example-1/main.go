package main

import "fmt"

func test1(f func()) {
	f() // 作为参数
}

func test2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
} // 作为返回值
func main() {
	func(s string) {
		fmt.Println(s)
	}(`hello world`) //匿名函数可以直接调用

	add := func(x, y int) int {
		
		return x + y
	} // 赋值给变量
	fmt.Println(add(1, 2))

	test1(func() { fmt.Println(`hello, world`) })
	fmt.Println(test2()(1, 2))

}
