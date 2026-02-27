package main

import (
	"fmt"
	"time"
)

func main() {
	rangeAlwaysWithYou()
	//testSpend()
}

func rangeAlwaysWithYou() {
	c := make(chan int, 1)
	go putIntIntoChannel(c, 3)
	go putIntIntoChannel(c, 5)
	for v := range c {
		//当range发现没有生产者的时候，会抛出死锁： fatal error: all goroutines are asleep - deadlock!
		//为啥会死锁呢？因为主线程作为receive了。那么如果是在一个goroutine中呢，试试吧

		fmt.Printf("从chan中读取数据啦. v=%d\n", v)
	}
	println("main func over")
}

func putIntIntoChannel(c chan int, second int) {
	i := 7
	for i > 0 {
		c <- i * second
		i--
		time.Sleep(time.Duration(second) * time.Second)
	}
	fmt.Printf("[putIntIntoChannel] is over with s = %d\n", second)
}

func testSpend() {
	//哈哈，好像也没测出来啥，就测试出来了chan的基本操作：阻塞队列
	//空了卡读。满了卡写。通过调节读写的休眠时间，可以看出来，差别越大效果越好
	c := make(chan int)
	go writeDataIntoChan(c, 9)
	go readDataFromChan(c, 1)
	time.Sleep(1 * time.Minute)
	println("[test_spend]收工")
}

func readDataFromChan(c chan int, sleepSecond int) {
	i := 10
	for i > 0 {
		fmt.Printf("[read]我要读数据，快给我\n")
		v := <-c
		fmt.Printf("[read]读到了，v=%d, 睡1秒\n", v)
		time.Sleep(time.Duration(sleepSecond) * time.Second)
		i--
	}
}

func writeDataIntoChan(c chan int, sleepSecond int) {
	i := 10
	for i > 0 {
		fmt.Printf("[write]我要写数据了，快点腾出空间让我写 v=%d \n", i)
		c <- i
		fmt.Printf("[write]写完了，v=%d \n", i)
		i--
		time.Sleep(time.Duration(sleepSecond) * time.Second)
	}
}
