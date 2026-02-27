package main

import (
	"context"
	"fmt"
	"time"
)

// 如果下游不主动检查ctx.Err()，长时间的操作，也一样根本没法结束的。
// 本例子也没有一个ctx的树状结构呢
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go processOrder(ctx)
	time.Sleep(3 * time.Second)

	fmt.Println("main: Canceling order processing!")
	cancel()

	// give some time for other goroutines to conclude
	time.Sleep(4 * time.Second)
}

func processOrder(ctx context.Context) {

	fmt.Println("Processing order...")
	time.Sleep(1 * time.Second)

	GetOderDetails(ctx)
	// Check if cancellation signal received
	if ctx.Err() != nil {
		fmt.Println("Canceled: Processing order")
		return
	}
}

func GetOderDetails(ctx context.Context) {
	// Perform some work
	fmt.Println("Fetching order details...")

	// Simulate a long-running task
	time.Sleep(1 * time.Second)
	GetInventoryDetails(ctx)

	// Check if cancellation signal received
	if ctx.Err() != nil {
		fmt.Println("Canceled: Fetching order details")
		return
	}

	// continue work after fetching inventory details
}

func GetInventoryDetails(ctx context.Context) {
	// Perform some work
	fmt.Println("Fetching inventory details...")

	// Simulate a long-running task
	time.Sleep(2 * time.Second)

	// Check if cancellation signal received
	if ctx.Err() != nil {
		fmt.Println("Canceled: Fetching inventory details")
		return
	}
}
