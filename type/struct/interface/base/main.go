package main

import (
	"fmt"
	"reflect"
)

type II interface {
	F1()
	F2()
}

type SS1 struct {
	vv int
	II // 这里嵌套了接口
}

func (ss *SS1) F1() {
	fmt.Println("[SS1] F1 method implement")
}

func (ss *SS1) F2() {
	fmt.Println("[SS1] F2 method implement")
}

type SS2 struct {
	vv int
	II
}

func (ss *SS2) F1() {
	fmt.Println("[SS2] F1 method implement")
}

func (ss *SS2) F2() {
	fmt.Println("[SS2] F2 method implement")
}

func main() {

	ss := SS1{}

	ss.F1()
	ss.F2()
	fmt.Println(reflect.TypeOf(ss))

	var ii II = &ss
	ii.F1()
	ii.F2()
	fmt.Println(reflect.TypeOf(ii)) // ii的类型是 *main.SS1

	// 重新赋值
	ii = &SS2{}
	ii.F1()
	ii.F2()
	fmt.Println(reflect.TypeOf(ii)) // ii的类型是 *main.SS1
}
