package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/pratice/golang/util/time_util"
)

func main() {
	// Test01NoBufferChannel()
	// Test01NoBufferChannelWillSleepWithoutRead()
	// Test02WriteIntoClosedChan()
	// TestWriteIntoChanWithoutConsumer()
	// TestSelectWhichCase()
	// TestSelectClosedChan()
	TestRangeChan()
	time_util.CountDownToZero("main 函数要结束了", 4)
}

func TestRangeChan() {
	collection := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			collection <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(collection) // 如果不关闭collection这个chan的话，range会死锁的.
	}()
	for i := range collection {
		log.Printf("i = %d \n", i)
	}
	time_util.CountDownToZero("TestRangeChan 函数要结束了", 4)
}

func TestSelectClosedChan() {
	ch := make(chan int)
	close(ch)
Label:
	for {
		select {
		case i := <-ch:
			log.Printf("read from ch, i = %d \n", i)
		case <-time.After(3 * time.Second):
			break Label

		}
	}
	time_util.CountDownToZero("TestSelectClosedChan 函数要结束了", 4)

}

// Test01NoBufferChannelWillSleepWithoutRead 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
func TestWriteIntoChanWithoutConsumer() {
	done := make(chan struct{})
	i1 := make(chan int)
	i2 := make(chan int)
	i3 := make(chan int)
	i4 := make(chan int)
	i5 := make(chan int)
	i6 := make(chan int)
	go TestSelect(done, i1, i2, i3, i4, i5, i6)
	i1 <- 1
	i2 <- 2
	i3 <- 3
	i4 <- 4
	done <- struct{}{} // 关闭了‘TestSelect’这个goroutine，从此没有了读取i5和i6的了，当然，i1和i2也没有了
	// i2 <- 2 	// 一样会panic-死锁
	// 当没有接受者后，再往ch中写数据，fatal error: all goroutines are asleep - deadlock!
	i5 <- 5
	i6 <- 6
	time_util.CountDownToZero("TestWriteIntoChanWithoutConsumer 函数要结束了", 4)
}

func TestSelectWhichCase() {
	done := make(chan struct{})
	i1 := make(chan int, 10)
	i2 := make(chan int, 10)
	i3 := make(chan int, 10)
	i4 := make(chan int, 10)
	i5 := make(chan int, 10)
	i6 := make(chan int, 10)
	go TestSelect(done, i1, i2, i3, i4, i5, i6)
	i1 <- 1
	i2 <- 2
	i3 <- 3
	i4 <- 4
	done <- struct{}{}
	i5 <- 5
	i6 <- 6
	// close(i6) // 如果关了i6，则i6的case分支就是可读的，而且一直可读，哈哈哈。
	time_util.CountDownToZero("TestWriteIntoChanWithoutConsumer 函数要结束了", 1)
}

func TestSelect(done chan struct{}, i1, i2, i3, i4, i5, i6 chan int) {
	time_util.CountDownToZero("在进入for循环之前", 1)
	for i := 0; i < 100; i++ {
		// 如果多个chan都OK，那么select分支会随机选一个；如果想全部都select到，那就在for循环中使用select。
		// 如果其中某个chan已经closed，那么这个分支就是OK的，但是如果真的选中了这个分支，那么从这个关闭的chan中读出来的数据是零值哦，所以要注意。
		var j = i
		select {
		case <-done:
			log.Printf("选择了done分之 i = %d ,j=%d\n", i, j)
			return // 这里可以调控的.
		case ii := <-i1:
			log.Printf("i1 接受到ii=%d i=%d, j=%d\n", ii, i, j)
		case ii := <-i2:
			log.Printf("i2 接受到ii=%d i=%d, j=%d\n", ii, i, j)
		case ii := <-i3:
			log.Printf("i3 接受到ii=%d i=%d, j=%d\n", ii, i, j)
		case ii := <-i4:
			log.Printf("i4 接受到ii=%d i=%d, j=%d\n", ii, i, j)
		case ii := <-i5:
			log.Printf("i5 接受到ii=%d i=%d, j=%d\n", ii, i, j)
		case ii := <-i6:
			log.Printf("i6 接受到ii=%d i=%d, j=%d\n", ii, i, j)
		}
	}
	log.Printf("for循环出来了")
}

func Test02WriteIntoClosedChan() {
	ch := make(chan int)
	close(ch)
	log.Printf("read from ch, i = %d\n", <-ch)
	log.Printf("read from ch, i = %d\n", <-ch)
	ch <- 19 // panic: send on closed channel
}

// Test01NoBufferChannel 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
// 一共两个变量，一个是chan的capacity，另一个是write动作是单写还是双写
func Test01NoBufferChannel() {
	// var ch = make(chan int) // 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
	var ch = make(chan int, 1) // 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
	// go writeToChV2(ch, rand.Intn(100))
	go writeToCh(ch, rand.Intn(100))
	time.Sleep(2 * time.Second)
	i := <-ch
	log.Printf("TestNoBufferChannel read from ch, i = %d \n", i)
}

func writeToCh(ch chan int, i int) {
	log.Printf("<<<<before write to ch, i = %d \n", i)
	ch <- i
	log.Printf(">>>>after write to ch, i = %d \n", i)
}

func writeToChV2(ch chan int, i int) {
	log.Printf("<<<<第一次写之前, i = %d \n", i)
	ch <- i
	log.Printf(">>>>第一次写完，第二次写之前 , i = %d \n", i)
	ch <- i
	log.Printf(">>>>第二次写完, i = %d \n", i)
}

// Test01NoBufferChannelWillSleepWithoutRead 如果不读就直接退出，不会导致 死锁
func Test01NoBufferChannelWillSleepWithoutRead() {

	var ch = make(chan int) // 无缓冲通道，写入动作是被阻塞的，直到下游消费者读取了。
	go writeToCh(ch, rand.Intn(100))
	time.Sleep(2 * time.Second)
	// close(ch) // panic: send on closed channel
	log.Printf("TestNoBufferChannel read from ch, i = %d \n", 2)
	time.Sleep(3 * time.Second)
}
