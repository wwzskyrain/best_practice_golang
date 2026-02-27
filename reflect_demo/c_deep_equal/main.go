package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []string{"1", "2"}
	b := []string{"2", "1"}
	aa := []string{"1", "2"}
	fmt.Printf("equal %b\n", reflect.DeepEqual(a, b))
	fmt.Printf("equal %b\n", reflect.DeepEqual(a, aa))
}
