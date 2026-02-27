package main

import (
	"fmt"
	"time"
)

func main() {
	read_some_one_randomly()
}

func read_some_one_randomly() {
	var c1 = make(chan int)
	var c2 = make(chan int)
	go put_some_int_into_chan(c1, 5)
	go put_some_int_into_chan(c2, 11)
	count := 5
	for {
		select {
		// 这两个case谁来了谁先返回，前提是没有default分支哈
		case e := <-c1:
			fmt.Printf("从c1中读出数据 %d \n", e)
		case e := <-c2:
			fmt.Printf("从c2中读出数据 %d \n", e)
			//default: // 这个default是个关键，有它，就不阻塞了
			//	println("no element in c1 and c2")
			//	time.Sleep(200 * time.Millisecond)
		}
		if count <= 0 {
			break
		}
		println("count=", count, "继续for")
		count--
	}
	println("主函数退出啦")
}

func put_some_int_into_chan(c chan int, sleep int) {
	i := 0
	for {
		c <- i * sleep
		i++
		fmt.Printf("[put_some_int_into_chan]睡[%d]秒哈，i=%d\n", sleep, i)
		time.Sleep(time.Duration(sleep) * time.Second)
		if i > 5 {
			break
		}
	}
	println("[put_some_int_into_chan]函数结束")
}
