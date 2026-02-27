package main

import (
	"fmt"
	"time"
)

func safeTask(id int, ch chan<- error) {
	defer func() {
		if r := recover(); r != nil {
			// 捕获 panic，并将错误传递给通道
			ch <- fmt.Errorf("Goroutine %d panicked: %v", id, r)
		}
	}()

	fmt.Printf("Goroutine %d started\n", id)
	time.Sleep(time.Second * 1)

	// 模拟 panic
	if id == 2 {
		panic("Something went wrong")
	}

	// 无错误时返回 nil
	ch <- nil
}

func main() {
	numGoroutines := 3
	ch := make(chan error, numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		go safeTask(i, ch)
	}

	for i := 1; i <= numGoroutines; i++ {
		err := <-ch
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Goroutine %d completed successfully\n", i)
		}
	}
}
