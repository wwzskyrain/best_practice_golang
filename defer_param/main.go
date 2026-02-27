package main

import (
	"strconv"
	"time"
)

func main() {

	testDeferInFor()
	time.Sleep(100 * time.Second)
}

// 本来想看看，surroundingFunc 结束后，name这个变量还能不能把defer函数访问到呢.
func surroundingFunc(i int) {
	name := strconv.Itoa(i)
	println("name = ", name)
	defer func() {
		println("defer name = ", name)
	}()
	println("func 【name】结束")
	if i == 30 {
		panic("i == 30")
	}
}

func testDeferInFor() {
	for i := 0; i < 100; i++ {
		lockName := "name" + strconv.Itoa(i)
		println("lockName=", lockName)
		defer func() {
			println("defer_lockName=", lockName)
		}()
	}
}
