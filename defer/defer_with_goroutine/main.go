package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("main 方法开始")
	logWithMethodName("main", "start--->")
	testWhenDeferRunWithGoroutine()
	time.Sleep(10 * time.Second)
	// 注释掉这一句，则main函数会立刻结束，
	// 然后go-routine就只说一句“这是goroutine函数=睡一会”就没有下一句了
	logWithMethodName("main", "end!!!")
}

func testWhenDeferRunWithGoroutine() {
	methodName := "testWhenDeferRunWithGoroutine"
	logWithMethodName(methodName, "start")
	defer func() {
		log.Println("这是defer函数.")
	}()
	go func() {
		log.Println("这是goroutine函数=睡一会")
		time.Sleep(1 * time.Millisecond)
		log.Println("这是goroutine函数=睡好了，结束！")
	}()
	logWithMethodName(methodName, "end")
}

func logWithMethodName(methodName string, logMessage string) {
	// fmt.Printf("[%s] %s\n", methodName, logMessage)
	log.Default().Printf("[%s] %s", methodName, logMessage)
}
