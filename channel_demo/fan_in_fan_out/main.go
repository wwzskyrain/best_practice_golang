package main

import (
	"fmt"
	"math"
	"sync"
)

func makeRange(min, max int) []int {
	diff := max - min
	// 创建arr数组
	arr := make([]int, diff+1)
	for i := 0; i < diff; i++ {
		arr[i] = min + i
	}
	return arr
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// 对输入的一个通道做filter（注意：它不负责fan_out），当然同时也返回一个新的通道
func prime(in <-chan int) <-chan int {
	out := make(chan int)
	// 开启 goroutine
	go func() {
		for i := range in {
			if isPrime(i) {
				out <- i
			}
		}
		// !!! 谁开的channel谁负责close
		close(out)
	}()
	return out
}

// 合并多个channel到一个channel，这个过程不做任何计算，就单纯的搬运
// 注意：它负责fan_in了
func merge(channels []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		// 这两步一定要放到goroutine中，不然这个函数就阻塞了
		wg.Wait()
		close(out)
	}()
	return out
}

// 把数组a导入到通道
func echo(a []int) <-chan int {
	out := make(chan int)
	go func() {
		// 注意：要在goroutine中，不然就没哪味儿了~
		for n := range a {
			out <- n
		}
		// 还是那句话，谁开的谁关（注意：close是非阻塞的）
		close(out)
	}()
	return out
}

// 对给定一个channel计算出一个sum值，但是以chan的形式返回而不是真的具体的一个值
func sum(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		s := 0
		for n := range c {
			s += n
		}
		out <- s
		close(out)
	}()

	return out
}

func main() {
	arr := makeRange(0, 1999)
	numChan := echo(arr)
	channels := make([]<-chan int, 5)
	//var channels [5]<-chan int // 与上一句等价
	for i, _ := range channels {
		// 在这里被分成了5个通道
		//（也就是开了5个goroutine来争着处理最原始的numChan通道中的数据，
		//	并返回了5个通道）
		// 	提示：sum其实已经把每个通道中的和计算好了，而且这个计算的和是
		//	以chan给出的，但是不能理解为“流式reduce”，它其实内部也是做完
		//	了sum动作的。
		channels[i] = sum(prime(numChan))
	}
	for n := range sum(merge(channels)) {
		// 只会打印出一个数据，再次印证了sum其实是一个只有一个元素的通道。
		fmt.Println(n)
	}
}
