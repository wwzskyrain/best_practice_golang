package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 原始字符串
	str := "iOS  14.2   (18B92)"

	// 定义正则表达式，匹配一个或多个空格
	re := regexp.MustCompile(`\s+`)

	// 使用 Split 方法按照正则表达式分割字符串
	parts := re.Split(str, -1)
	for i := range parts {
		fmt.Printf("No.%d = %s \n", i, parts[i])
	}
	// 输出分割后的结果
	fmt.Println(parts) // 输出: [iOS 14.2 (18B92)]

}
