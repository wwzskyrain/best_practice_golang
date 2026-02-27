package main

import (
	"fmt"
	"runtime"
)

func main() {

	testCatchPanic()

	fmt.Println("----------------")

	testPanicWithoutCatch()
	// 不捕捉panic导致main函数退出，下面的输出语句不会执行
	fmt.Println("main over!")
}

func testCatchPanic() {
	defer CatchPanic("我是守护者")
	fmt.Println("开始，要panic嘞")
	panic("我panic来了")
	fmt.Println("我是[testCatchPanic]的over")
}

func testPanicWithoutCatch() {
	fmt.Println("开始，要panic嘞")
	panic("我panic来了")
}

// CatchPanic cover panic from function
func CatchPanic(funcName string) {
	if e := recover(); e != nil {
		buf := make([]byte, 64<<10)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf("funcName: %v, err: %v, stack trace: %v", funcName, e, string(buf))
	}
}
