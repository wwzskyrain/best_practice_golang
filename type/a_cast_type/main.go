package main

import (
	"fmt"

	"github.com/pratice/golang/util/util_log"
)

// processValue 处理不同类型的值
func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		// 处理整数类型
		squared := v * v
		fmt.Printf("传入的是整数 %d，其平方是 %d\n", v, squared)
	case string:
		// 处理字符串类型
		length := len(v)
		fmt.Printf("传入的是字符串 %s，其长度是 %d\n", v, length)
	case []int:
		// 处理整数切片类型
		sum := 0
		for _, num := range v {
			sum += num
		}
		fmt.Printf("传入的是整数切片 %v，其元素之和是 %d\n", v, sum)
	default:
		fmt.Printf("不支持的类型: %T\n", v)
	}
}

func processToInt64(value interface{}) int64 {

	switch v := value.(type) {
	case int:
		util_log.Log(util_log.LevelInfo, "[processToInt64]v is int,  %d", v)
		return int64(v)
	case int32:
		util_log.Log(util_log.LevelInfo, "[processToInt64]v is int32,  %d", v)
		return int64(v)
	case int64:
		util_log.Log(util_log.LevelInfo, "[processToInt64]v is int64,  %d", v)
		return v
	default:
		util_log.Log(util_log.LevelInfo, "[processToInt64]v is not int,  %T", v)
		return 0
	}
}

func main() {
	// testProcessValue()
	fmt.Printf("processToInt64(0) = %d\n", processToInt64(0))
	fmt.Printf("processToInt64(1) = %d\n", processToInt64(int32(1)))
	fmt.Printf("processToInt64(2) = %d\n", processToInt64(int64(2)))
	fmt.Printf("processToInt64(3333333333) = %d\n", processToInt64(3333333333))

}

func testProcessValue() {
	// 传入整数
	processValue(5)
	// 传入字符串
	processValue("hello")
	// 传入整数切片
	processValue([]int{1, 2, 3, 4, 5})
	// 传入不支持的类型
	processValue(true)
}
