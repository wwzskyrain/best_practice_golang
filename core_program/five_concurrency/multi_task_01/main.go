package main

import (
	"log"
	"sync"
)

func main() {
	// taskChan 这个是main提供，但是有生产者关闭它。
	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	wait := sync.WaitGroup{}
	go InitTask(taskChan, resultChan, 10000)
	go DistributeTask(taskChan, &wait, resultChan)
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

func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, result chan int) {
	for t := range taskChan {
		wait.Add(1)
		go func() {
			t.doTask()
			wait.Done()
		}()
	}
	wait.Wait()
	close(result) // result的生产方是一个一个的task，而这里就是task的发起方，所以这里算result的总生产方，所以这里负责关闭result通道
}

func ProcessResult(resultChan chan int) int {
	sum := 0
	for r := range resultChan {
		sum += r
	}
	return sum
}
