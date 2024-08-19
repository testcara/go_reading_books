package main

import (
	"fmt"
	"time"
)

// select 用来监听，那1个有数据，就执行哪里的对应逻辑
func main() {
	tick := time.Tick(1 * time.Millisecond)
	boom := time.After(5 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
		default:
			fmt.Println("...")
		}
	}
}
