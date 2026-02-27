package main

import (
	"log"

	"github.com/pratice/golang/util/time_util"
)

func main() {
	toBeAppend := []int{100, 101}
	go func() {
		slice := getSlice(1)
		slice = append(slice, toBeAppend...)
		log.Printf("slice = %+v", slice)
	}()

	go func() {
		slice := getSlice(2)
		slice = append(slice, toBeAppend...)
		log.Printf("slice = %+v", slice)
	}()
	time_util.CountDownToZero("main", 5)
}

func getSlice(i int) []int {
	var ret []int
	if i%2 == 0 {
		ret = append(ret, 1)
	}
	return ret
}
