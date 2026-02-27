package main

import (
	"log"
)

func main() {
	// taskChan 这个是main提供，但是有生产者关闭它。
	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	done := make(chan struct{}, 10)
	workers := 3
	go InitTask(taskChan, resultChan, 10000)
	go DistributeTask(taskChan, workers, done)
	go closeResult(done, resultChan, workers)
	sum := ProcessResult(resultChan)
	log.Printf("sum = %d\n", sum)
}

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) doTask() {
	sum := 0
	for i := t.begin; i < t.end; i++ {
		sum += i
	}
	t.result <- sum
}
func InitTask(taskChan chan<- task, r chan int, p int) {
	// 第一个入参taskChan必须有协调者创建，然后有这里来使用它，而不能有这里(使用者)来创建(然后返回给协调方)
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		t := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- t
	}
	if mod != 0 {
		t := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- t
	}
	close(taskChan) // 就这么多task了，所以关闭taskChan就行了
}

func DistributeTask(taskChan <-chan task, works int, done chan struct{}) {
	for i := 0; i < works; i++ {
		go processTask(taskChan, done)
	}
}

func processTask(taskChan <-chan task, done chan struct{}) {
	for t := range taskChan {
		t.doTask()
	}
	done <- struct{}{}
}

func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		// 一直累加，直到resultChan关闭
		sum += r
	}
	return sum
}

func closeResult(done chan struct{}, resultChan chan int, works int) {
	for i := 0; i < works; i++ {
		<-done
	}
	// 所有的work的结果都出来了，可以关闭resultChan了
	close(done) // 这个done关不关无所谓了。
	close(resultChan)
}
