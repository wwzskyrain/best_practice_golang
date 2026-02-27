package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/pratice/golang/util/time_util"
)

func main() {
	done := make(chan struct{})
	ch := GenerateInt(done)
	for i := 0; i < 3; i++ {
		// time.Sleep(5 * time.Second)
		time_util.CountDownToZero("接受ch", 3)
		log.Printf("No. %d r = %d\n", i, <-ch)
	}
	time_util.CountDownToZero("关闭done之前", 3)
	done <- struct{}{}
	// time_util.CountDownToZero("关闭done之后", 3)
	log.Printf("should have last one, r = %d\n", <-ch)
	log.Printf("another one ?  r = %d\n", <-ch)
	println("main done")
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	doneToAB := make(chan struct{}) // 这个chan就只是一个提线木偶的提线，一拉那边就关闭了
	resultFromA := GenerateIntA(doneToAB)
	go func() {
	Label:
		for {
			// time_util.CountDownToZero("在 resultFromA 之前", 3)
			select {
			// case i := <-resultFromA: // 这种写法就有一个死锁，而且已经分析出来了。这个分支一定会成功，所以大部分时间，这个select都阻塞在ch<-i那一行。
			// 	log.Printf("从resultFromA 接收到i = %d, 然后要写给ch\n", i)
			// 	ch <- i  //对，阻塞在这里。当主程序要发送done信号时，select阻塞在这里从而不能接受done信号。而外面要写done信号，这样就死锁啦，哈哈哈。
			// 	log.Printf("写好了，写给了ch，i = %d\n", i)
			case ch <- <-resultFromA: // 双箭头的好处，我敢肯定，resultFromA已经写入了（因为那边有打印日志啊），
				// 而且这里其实也已经读取了（因为对于无缓冲的ch，这边不读取，那边就一直阻塞的），但是ch还没有写入哟，所以select的分支还没有确定。
				// 因为外面没有读取ch，所以这个case就不会走进来，而走下面的done的case咯，哈哈哈。
				log.Printf("双箭头，从resultFromA写给了ch\n")
			case <-done:
				// 如果done通道关闭，则跳出循环
				fmt.Printf("接到done信号，GenerateInt，发送doneToAB，然后转到Label\n")
				doneToAB <- struct{}{}
				fmt.Printf("发送了doneToAB信号，转到Label\n")
				break Label
			}
		}
		log.Printf("关闭GenerateInt的ch")
		close(ch)
	}()
	return ch
}

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case <-done:
				log.Printf("接到done信号，GenerateIntA 要done了，转到Label")
				break Label
			case ch <- rand.Int():
				log.Printf("随机写了一个int给GenerateIntAAA的ch")
			}
		}
		log.Printf("关闭GenerateIntAAA的ch")
		close(ch)
	}()
	return ch
}
