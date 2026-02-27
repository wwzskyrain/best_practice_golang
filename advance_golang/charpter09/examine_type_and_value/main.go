package main

import (
	"fmt"
	"reflect"
)

func main() {
	// chapter9/sources/go-reflect/examine_value_and_type.go

	// 简单原生类型
	var b = true // 布尔型
	val := reflect.ValueOf(b)
	typ := reflect.TypeOf(b)
	fmt.Println(typ.Name(), val.Bool()) // bool true

	var i = 23 // 整型
	val = reflect.ValueOf(i)
	typ = reflect.TypeOf(i)
	fmt.Println(typ.Name(), val.Int()) // int 23

	var f = 3.14 // 浮点型
	val = reflect.ValueOf(f)
	typ = reflect.TypeOf(f)
	fmt.Println(typ.Name(), val.Float()) // float64 3.14

	var s = "hello, reflection" // 字符串
	val = reflect.ValueOf(s)
	typ = reflect.TypeOf(s)
	fmt.Println(typ.Name(), val.String()) // string hello, reflection

	var fn = func(a, b int) int { // 函数(一等公民)
		return a + b
	}
	val = reflect.ValueOf(fn)
	typ = reflect.TypeOf(fn)
	fmt.Println(typ.Kind(), typ.String()) // func func(int, int) int
	fmt.Println(typ.Name(), typ.String()) // "" func(int, int) int
}
