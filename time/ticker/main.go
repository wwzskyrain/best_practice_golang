package main

import (
	"context"
	"fmt"
	"time"

	"code.byted.org/gopkg/logs"

	"github.com/pratice/golang/util"
)

func main() {
	// testTicker()
	// testTicker2()
	StartCrashMonitorService()
}

func testTicker() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			ctx := context.Background()
			logs.CtxInfo(ctx, "来一次检查并报警, t=%v", t)
		}
	}()
	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("刚刚已经调用ticker.Stop()")
}

func testTicker2() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			ctx := context.Background()
			logs.CtxInfo(ctx, "来一次检查并报警, t=%v")
		}
	}()
	// defer ticker.Stop()
	time.Sleep(10 * time.Second)
	fmt.Println("函数结束!")
}

func StartCrashMonitorService() {

	util.SafeGoRoutine(context.Background(), "crash_detect_service", func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			// ctx := createContext()
			ctx := context.Background()
			logs.CtxInfo(ctx, "来一次检查并报警")
		}
	})

	time.Sleep(20 * time.Second)
}
