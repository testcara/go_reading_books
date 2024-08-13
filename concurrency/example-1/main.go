package main

import "fmt"

func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println("recv:", x)
	}
	done <- true // 通知main, 消费结束
}

func producer(data chan int) {
	for i :=0; i <4; i++ {
		data <- i
	}
	close(data) // 生产结束，关闭通道
}

func main() {
	done := make(chan bool) // 用来接收消费者结束信号
	data := make(chan int) // 数据管道

	// 启动消费者
    go consumer(data, done)
	// 启动生产者
	go producer(data)

	// 阻塞，直到消费者发回结束信号
	<-done
}