package main

import (
	"fmt"
	"time"
	"unsafe"
)

type Person struct {
	age  int
	name string
}

func main() {
	var p Person
	if &p != nil { // 这里必须要用&p，不能用p的，否则编译报错：invalid operation: p == nil (mismatched types Person and untyped nil)
		fmt.Println("p is not nil")
	}
	fmt.Println("p：", p)
	testEmptyStruct()
	testSetUsage()
	testEmptyChannelUsage()
}

// EST 定义一个空结构体
type EST struct {
}

func testEmptyStruct() {
	var b EST
	var c struct{}
	fmt.Printf("b 地址：%p size:%d\n", &b, unsafe.Sizeof(b))
	fmt.Printf("c 地址：%p size:%d\n", &c, unsafe.Sizeof(c))
}

func testSetUsage() {
	students := make(map[string]struct{}, 10)
	students["zhangsan"] = EST{}
	students["lisi"] = struct{}{}
	students["lisi"] = EST{} // 再加一个李四
	fmt.Println(len(students))
}

func testEmptyChannelUsage() {
	c := make(chan struct{}, 0)
	fmt.Println("turn on a channel")
	go func(c chan<- struct{}) {
		time.Sleep(3 * time.Second)
		fmt.Println("sub process sleep over!")
		c <- EST{}
	}(c)
	fmt.Println("main thread still alive")
	<-c
	fmt.Println("main thread over!")
}
