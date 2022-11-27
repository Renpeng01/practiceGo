package main

import (
	"context"
	"fmt"
	"time"
)

// 使用context 控制请求其他服务的超时
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)  // 超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 完成请求
	defer cancel()

	done := make(chan int, 0)

	go func() {
		//  这里模拟发送请求
		time.Sleep(5 * time.Second)
		done <- 1
	}()

	select {
	case <-done:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("cancle", ctx.Err().Error())
	}
}
