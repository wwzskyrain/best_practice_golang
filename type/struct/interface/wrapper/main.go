package main

import "fmt"

type Action interface {
	Jump()
	Run()
}

type Somebody struct{}

func (body *Somebody) Jump() {
	fmt.Println("[Somebody] implement (Jump) method!")
}

func (body *Somebody) Run() {
	fmt.Println("[Somebody] implement (Run) method!")
}

type Hero struct {
	Action
	Type int
}

func (h *Hero) Jump() {
	// override（覆盖）实现Jump方法
	fmt.Println("[Hero] implement (Jump) method!")
}

func main() {
	hero := Hero{Action: &Somebody{}, Type: 2}
	hero.Action.Run() // 等价于 hero.Run()
	hero.Run()
	hero.Jump()
}
