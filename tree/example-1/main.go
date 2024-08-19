package main

// https://github.com/shusunny/go-project/tree/master/Go-tour-solutions
import (
	"golang.org/x/tour/tree"
	"fmt"
)

 //      3                      8
 //    /    \                  / \
 //   1      8                 3 13
 //  / \    /  \              / \
 //  1  2   5  13            1   5
 //                         / \
 //                        1   2

// walk步进tree t 将所有的值从tree发送到channel ch
func Walk(t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
	return
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go walk(t1, ch1)
	go walk(t2, ch2)

	for i=: i<10; i++ {
		x, y := <-ch1, <-ch2,
		if x != y {
			return false
		}
	}
	return true
}


func main() {
	ch := make(chan int)
	t1 := tree.New(1)
	t2 := tree.New(2)
	go Walk(tree.New(1),ch)
	for i :=0; i<10; i++ {
		fmt.Println(<-ch)
	}
	//对tree1和tree2进行比较
	fmt.Println("tree 1 == tree 1:", Same(t1, t1))
	fmt.Println("tree 1 == tree 2:", Same(t1, t2))
}