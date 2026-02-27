package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// TestSlowTaskSeeWhichSelectBranch()
	TestSlowTaskV2HowManyBranchIsPassed()
}

func TestSlowTaskSeeWhichSelectBranch() {
	ctx := context.Background()
	// 调整这里的过期时间，可以控制slowTask是从哪个分支来结束
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	go slowTask(ctx)
	// give some time for the slowTask goroutine to conclude
	time.Sleep(5 * time.Second)
	println("TestSlowTaskSeeWhichSelectBranch() stop")
}
func slowTask(ctx context.Context) {
	select {
	// simulate a time-consuming task using time.After
	case <-time.After(time.Second * 3):
		fmt.Println("Operation complete.")
	case <-ctx.Done():
		// print the cancellation message with the reason using ctx.Err().
		fmt.Println("Canceled:", ctx.Err())
	}
}

func TestSlowTaskV2HowManyBranchIsPassed() {
	ctx := context.Background()
	// 调整这里的过期时间，可以控制slowTask是从哪个分支来结束
	duration, err := time.ParseDuration("+5s")
	if err != nil {
		fmt.Println("err = ", err)
	}
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	go slowTaskV2(ctx)
	// give some time for the slowTask goroutine to conclude
	time.Sleep(20 * time.Second)
	println("TestSlowTaskV2HowManyBranchIsPassed() stop")
}

func slowTaskV2(ctx context.Context) {
RETURN:
	for true {
		select {
		// simulate a time-consuming task using time.After
		case <-time.After(time.Second * 2):
			fmt.Println("Operation complete.")
		case <-ctx.Done():
			// print the cancellation message with the reason using ctx.Err().
			fmt.Println("Canceled:", ctx.Err())
			break RETURN
		}
	}
}
