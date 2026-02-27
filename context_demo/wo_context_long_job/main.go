package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// 调整超时时间，可以<5或者>5，可以看到不同的结果
	ctx, _ = context.WithTimeout(ctx, 6*time.Second)
	go queryWithCancel(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("[main] 结束")
}

func queryWithCancel(ctx context.Context) {
	resultCh := make(chan int)
	// 必须新开一个goroutine
	go queryThroughRPC(5*time.Second, resultCh)
	select {
	case <-ctx.Done():
		fmt.Println("【queryWithCancel】超时了.") // 到点了就不等下面的resultCh了
	case r := <-resultCh:
		{
			fmt.Printf("【queryWithCancel】结果回来了 r = %d \n", r)
		}
	}
	fmt.Println("【queryWithCancel】函数结束")
}

func queryThroughRPC(t time.Duration, resultCh chan<- int) {
	time.Sleep(t)
	result := 100
	fmt.Printf("【queryThroughRPC】 result = %d\n", result)
	resultCh <- result
	return
}
