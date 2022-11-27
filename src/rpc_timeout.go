package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息

		isSuccess := true
		if url == "http://rpc_2_url" {
			isSuccess = false
		} else {
			time.Sleep(5 * time.Second)
		}

		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
	case <-ctx.Done():
		// 其他RPC调用调用失败
		return ctx.Err()
	case e := <-err:
		// 本RPC调用失败，返回错误信息
		return e
	case <-result:
		// 本RPC调用成功，不返回错误信息
		return nil
	}
}

// 使用context 通知多个并发请求操作
//  rpc 2，,3，4 并行，2失败后使3,4也一块结束
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// RPC1调用
	err := Rpc(ctx, "http://rpc_1_url")
	if err != nil {
		fmt.Println("1", err.Error())
		return
	}

	wg := sync.WaitGroup{}

	// RPC2调用
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_2_url")
		if err != nil {
			fmt.Println("2", err.Error())
			cancel()
		}
	}()

	// RPC3调用
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_3_url")
		if err != nil {
			cancel()
			fmt.Println("3", err.Error())
		}
	}()

	// RPC4调用
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_4_url")
		if err != nil {
			cancel()
			fmt.Println("4", err.Error())
		}
	}()

	wg.Wait()
}

// 使用context 控制请求其他服务的超时
func main2() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // 超时
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 完成请求
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
